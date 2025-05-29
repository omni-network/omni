package usdt0_test

import (
	"context"
	"flag"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/errors"
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

	// Private key needs at least 1 USDT on Ethereum to run this test, plus ETH for gas
	pkHex := os.Getenv("TEST_PRIVATE_KEY")
	if pkHex == "" {
		t.Skip("TEST_PRIVATE_KEY environment variable not set, skipping integration test")
	}

	pk, err := crypto.HexToECDSA(strings.TrimPrefix(pkHex, "0x"))
	require.NoError(t, err)

	user := crypto.PubkeyToAddress(pk.PublicKey)

	// Get Ethereum client
	ethClient, err := ethclient.Dial(evmchain.Name(evmchain.IDEthereum), "https://eth.llamarpc.com")
	require.NoError(t, err)

	// Get HyperEVM client
	hyperClient, err := ethclient.Dial(evmchain.Name(evmchain.IDHyperEVM), "https://rpc.hyperevm.com")
	require.NoError(t, err)

	// Create backend for Ethereum
	ethBackend, err := ethbackend.NewBackend(
		evmchain.Name(evmchain.IDEthereum),
		evmchain.IDEthereum,
		12, // 12 second block period
		ethClient,
		pk,
	)
	require.NoError(t, err)

	// Check USDT balance on Ethereum
	usdt, ok := tokens.ByAsset(evmchain.IDEthereum, tokens.USDT)
	require.True(t, ok)

	balance, err := tokenutil.BalanceOf(ctx, ethClient, usdt, user)
	require.NoError(t, err)

	amount := bi.Dec6(1)
	tutil.RequireGTE(t, balance, amount)

	receipt, err := usdt0.Send(ctx, ethBackend, user, evmchain.IDEthereum, evmchain.IDHyperEVM, amount)
	require.NoError(t, err)

	log.Info(ctx, "Sent USDT0 from Ethereum to HyperEVM", "tx_hash", receipt.TxHash.Hex())

	// Check USDT0 balance on HyperEVM
	usdt0, ok := tokens.ByAsset(evmchain.IDHyperEVM, tokens.USDT0)
	require.True(t, ok)

	// Create LayerZero client to check message status
	lzClient := layerzero.NewClient(layerzero.MainnetAPI)

	// Wait for message to be delivered, checking every minute for up to 10 minutes
	ctx, cancel := context.WithTimeout(ctx, 10*time.Minute)
	defer cancel()

	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	getMsg := func() (layerzero.Message, string) {
		messages, err := lzClient.GetMessagesByTx(ctx, receipt.TxHash.Hex())
		require.NoError(t, err)
		require.NotEmpty(t, messages, "No messages found for transaction")

		if len(messages) != 1 {
			log.Warn(ctx, "Expected 1 message",
				errors.New("unexpected number of messages"),
				"expected", 1,
				"got", len(messages))
		}

		return messages[0], messages[0].Status.Name
	}

	for {
		select {
		case <-ctx.Done():
			_, status := getMsg()
			t.Fatalf("Timeout waiting for message to be delivered. Last status: %s", status)
		case <-ticker.C:
			msg, status := getMsg()

			if msg.IsFailed() {
				t.Fatalf("Message failed: %s", status)
			}

			if msg.IsDelivered() {
				// Message delivered, check balance
				balance, err = tokenutil.BalanceOf(ctx, hyperClient, usdt0, user)
				require.NoError(t, err)
				tutil.RequireEQ(t, balance, amount)

				return
			}

			log.Info(ctx, "Message status", "status", status)
		}
	}
}
