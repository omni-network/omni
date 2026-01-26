package usdt0_test

import (
	"context"
	"flag"
	"log/slog"
	"os"
	"strings"
	"testing"
	"time"

	e2e "github.com/omni-network/omni/e2e/types"
	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/layerzero"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/tokens"
	"github.com/omni-network/omni/lib/tokens/tokenutil"
	"github.com/omni-network/omni/lib/tutil"
	"github.com/omni-network/omni/lib/usdt0"

	"github.com/ethereum/go-ethereum/crypto"

	cosmosdb "github.com/cosmos/cosmos-db"
	"github.com/stretchr/testify/require"
)

var integration = flag.Bool("integration", false, "run integration tests")

// TestSendUSDT0 tests sending USDT0 from Ethereum to HyperEVM, on mainnet.
func TestSendUSDT0(t *testing.T) {
	t.Parallel()

	if !*integration {
		t.Skip("Skipping integration test. Use -integration flag to run")
	}

	ctx := t.Context()

	logCfg := log.DefaultConfig()
	logCfg.Level = slog.LevelDebug.String()
	logCfg.Color = log.ColorForce

	ctx, err := log.Init(ctx, logCfg)
	tutil.RequireNoError(t, err)

	// Private key needs at least 1 USDT on Ethereum to run this test, plus ETH for gas
	pkHex := os.Getenv("TEST_PRIVATE_KEY")
	if pkHex == "" {
		t.Skip("TEST_PRIVATE_KEY environment variable not set, skipping integration test")
	}

	pk, err := crypto.HexToECDSA(strings.TrimPrefix(pkHex, "0x"))
	require.NoError(t, err)

	user := crypto.PubkeyToAddress(pk.PublicKey)

	srcChain := mustChain(t, evmchain.IDEthereum)
	dstChain := mustChain(t, evmchain.IDHyperEVM)
	srcChainID := srcChain.ChainID
	dstChainID := dstChain.ChainID
	srcToken := mustUSDT0Token(t, srcChainID)
	dstToken := mustUSDT0Token(t, dstChainID)

	srcClient, err := ethclient.Dial(srcChain.Name, e2e.PublicRPCByName(srcChain.Name))
	require.NoError(t, err)

	dstClient, err := ethclient.Dial(dstChain.Name, e2e.PublicRPCByName(dstChain.Name))
	require.NoError(t, err)

	// Create backend for Ethereum
	srcBackend, err := ethbackend.NewBackend(
		srcChain.Name,
		srcChain.ChainID,
		srcChain.BlockPeriod,
		srcClient,
		pk,
	)
	require.NoError(t, err)

	// Check balance on Ethereum
	balance, err := tokenutil.BalanceOf(ctx, srcClient, srcToken, user)
	require.NoError(t, err)

	amount := bi.Dec6(1)
	tutil.RequireGTE(t, balance, amount)

	// Create new database
	memDB := cosmosdb.NewMemDB()
	db, err := usdt0.NewDB(memDB)
	require.NoError(t, err)

	// Create LayerZero client
	lzClient := layerzero.NewClient(layerzero.MainnetAPI)

	// Add 10 minute timeout to context
	ctx, cancel := context.WithTimeout(ctx, 20*time.Minute)
	defer cancel()

	usdt0.MonitorSendsForever(ctx, db, lzClient, []uint64{srcChainID, dstChainID})

	// Send USDT0
	receipt, err := usdt0.Send(ctx, srcBackend, user, srcChainID, dstChainID, amount, db)
	require.NoError(t, err)

	// Wait for message to be deleted from db (indicating delivery)
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			t.Fatal("Timeout waiting for message to be delivered")
		case <-ticker.C:
			// Check messages in db
			msgs, err := db.GetMsgs(ctx)
			require.NoError(t, err)

			log.Info(ctx, "Checking for message delivery...", "msgs_in_db", len(msgs))
			for _, m := range msgs {
				log.Info(ctx, "Message in db", "tx_hash", m.TxHash, "status", m.Status)
			}

			if len(msgs) > 1 {
				t.Fatalf("Expected 1 message in db, got %d", len(msgs))
			}

			if len(msgs) == 1 {
				require.Equal(t, receipt.TxHash, msgs[0].TxHash, "Message in database has wrong tx hash")
				log.Info(ctx, "Message still in db, waiting for delivery")

				continue
			}

			// No messages in db, verify delivery with LayerZero
			messages, err := lzClient.GetMessagesByTx(ctx, receipt.TxHash.Hex())
			require.NoError(t, err)
			require.NotEmpty(t, messages, "No messages found for transaction")
			require.Len(t, messages, 1, "Expected 1 message")
			require.True(t, messages[0].IsDelivered(), "Message not delivered")

			// Check balance on HyperEVM
			balance, err = tokenutil.BalanceOf(ctx, dstClient, dstToken, user)
			require.NoError(t, err)
			tutil.RequireGTE(t, balance, amount)

			log.Info(ctx, "Message delivered")

			return
		}
	}
}

func mustUSDT0Token(t *testing.T, chainID uint64) tokens.Token {
	t.Helper()

	// On L1, use canonical USDT token
	if chainID == evmchain.IDEthereum {
		return mustToken(t, chainID, tokens.USDT)
	}

	return mustToken(t, chainID, tokens.USDT0)
}

func mustToken(t *testing.T, chainID uint64, asset tokens.Asset) tokens.Token {
	t.Helper()

	token, ok := tokens.ByAsset(chainID, asset)
	require.True(t, ok)

	return token
}

func mustChain(t *testing.T, id uint64) evmchain.Metadata {
	t.Helper()

	chain, ok := evmchain.MetadataByID(id)
	require.True(t, ok)

	return chain
}
