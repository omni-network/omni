package cctp_test

import (
	"context"
	"flag"
	"log/slog"
	"math/big"
	"os"
	"os/signal"
	"syscall"
	"testing"
	"time"

	"github.com/omni-network/omni/lib/anvil"
	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/cctp"
	cctpdb "github.com/omni-network/omni/lib/cctp/db"
	"github.com/omni-network/omni/lib/cctp/testutil"
	"github.com/omni-network/omni/lib/cctp/types"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/tokens"
	"github.com/omni-network/omni/lib/tokens/tokenutil"
	"github.com/omni-network/omni/lib/tutil"
	"github.com/omni-network/omni/lib/xchain"
	xprovider "github.com/omni-network/omni/lib/xchain/provider"

	"github.com/ethereum/go-ethereum/common"

	cosmosdb "github.com/cosmos/cosmos-db"
	"github.com/stretchr/testify/require"
)

var (
	integration = flag.Bool("integration", false, "run integration tests")
)

//go:generate go test . -integration -v -run=TestIntegration

// TestIntegration runs a CCTP integration test. It:
//   - runs anvil forks of mainnet chains
//   - attests and signs CCTP messages w/ a devnet attester
//   - bridges USDC from chain to chain, asserting success
func TestIntegration(t *testing.T) {
	t.Parallel()

	if !*integration {
		t.Skip("skipping integration test")
	}

	ctx := t.Context()

	logCfg := log.DefaultConfig()
	logCfg.Level = slog.LevelDebug.String()
	logCfg.Color = log.ColorForce

	ctx, err := log.Init(ctx, logCfg)
	require.NoError(t, err)

	// Handle interrupts
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer cancel()

	rpcs := getForkRPCs(t)
	chains := getChains(t)
	network := makeNetwork(t, chains)

	clients, stop := testutil.StartAnvilForks(t, ctx, rpcs, chains)
	defer stop()

	// Stop anvil on interrupt
	go func() {
		<-ctx.Done()
		stop()
	}()

	devPk, devAddr := testutil.NewAccount(t) // Test user

	backends, err := ethbackend.BackendsFromClients(clients, devPk)
	require.NoError(t, err)

	// Fund dev account with USDC & ETH
	fund(t, ctx, clients, devAddr)

	// Create xchain provider
	xprov := xprovider.New(network, clients, nil)

	// Start attesting
	cctpClient := cctp.StartTestClient(ctx, t, xprov, chains, clients)

	// In-mem CCTP DB
	memDB := cosmosdb.NewMemDB()
	db, err := cctpdb.New(memDB)
	require.NoError(t, err)

	// Mint forever
	err = cctp.MintForever(ctx, db, cctpClient, backends, chains, devAddr,
		cctp.WithMintInterval(1*time.Second),
		cctp.WithPurgeInterval(10*time.Second))
	require.NoError(t, err)

	// Audit forever
	err = cctp.AuditForever(ctx, db, netconf.Mainnet, xprov, clients, chains, devAddr)
	require.NoError(t, err)

	// Track initial balances
	initialBalances := make(map[uint64]*big.Int)
	for _, chain := range chains {
		usdc := mustUSDC(t, chain.ChainID)
		backend, err := backends.Backend(chain.ChainID)
		require.NoError(t, err)

		balance, err := tokenutil.BalanceOf(ctx, backend, usdc, devAddr)
		require.NoError(t, err)

		initialBalances[chain.ChainID] = balance
	}

	// Track expected final balances
	expectedBalances := make(map[uint64]*big.Int)
	for chainID, balance := range initialBalances {
		expectedBalances[chainID] = new(big.Int).Set(balance)
	}

	// wrongDb is used to simulate messages missed, and not stored in the db.
	// these should be caught by audit.
	wrongCosmosDB := cosmosdb.NewMemDB()
	wrongDB, err := cctpdb.New(wrongCosmosDB)
	require.NoError(t, err)

	// Define sends
	sends := []struct {
		srcChainID  uint64
		destChainID uint64
		amount      *big.Int
		wrongDB     bool
	}{
		{evmchain.IDEthereum, evmchain.IDOptimism, bi.Dec6(50), false},
		{evmchain.IDArbitrumOne, evmchain.IDBase, bi.Dec6(50), false},
		{evmchain.IDOptimism, evmchain.IDBase, bi.Dec6(25), true}, // wrong db
		{evmchain.IDBase, evmchain.IDEthereum, bi.Dec6(10), false},
		{evmchain.IDEthereum, evmchain.IDOptimism, bi.Dec6(5), false},
		{evmchain.IDOptimism, evmchain.IDBase, bi.Dec6(2), false},
		{evmchain.IDBase, evmchain.IDArbitrumOne, bi.Dec6(1), true}, // wrong db
		{evmchain.IDArbitrumOne, evmchain.IDEthereum, bi.Dec6(0.5), false},
		{evmchain.IDEthereum, evmchain.IDBase, bi.Dec6(0.1), true}, // wrong db
	}

	// Do sends
	msgs := make([]types.MsgSendUSDC, len(sends))
	for i, send := range sends {
		backend, err := backends.Backend(send.srcChainID)
		require.NoError(t, err)

		sendDB := db
		if send.wrongDB {
			sendDB = wrongDB
		}

		msg, err := cctp.SendUSDC(ctx, sendDB, netconf.Mainnet, backend, cctp.SendUSDCArgs{
			Sender:      devAddr,
			Recipient:   devAddr,
			SrcChainID:  send.srcChainID,
			DestChainID: send.destChainID,
			Amount:      send.amount,
		})
		require.NoError(t, err)
		msgs[i] = msg

		// Update expected balances
		expectedBalances[send.srcChainID] = bi.Sub(expectedBalances[send.srcChainID], send.amount)
		expectedBalances[send.destChainID] = bi.Add(expectedBalances[send.destChainID], send.amount)
	}

	// Wait for all sends
	tutil.RequireEventually(t, ctx, func() bool {
		for chainID, expectedBalance := range expectedBalances {
			usdc := mustUSDC(t, chainID)

			backend, err := backends.Backend(chainID)
			tutil.RequireNoError(t, err)

			balance, err := tokenutil.BalanceOf(ctx, backend, usdc, devAddr)
			tutil.RequireNoError(t, err)

			if !bi.EQ(balance, expectedBalance) {
				return false
			}
		}

		log.Info(ctx, "All sends completed")

		return true
	}, 2*time.Minute, 1*time.Second)

	// Confirm all messages received
	for _, msg := range msgs {
		received, err := cctp.DidReceive(ctx, clients[msg.DestChainID], msg, nil)
		require.NoError(t, err)
		require.True(t, received, "message not received on dest chain %d", msg.DestChainID)
	}

	// Wait for all purged (confirmed and deleted)
	tutil.RequireEventually(t, ctx, func() bool {
		msgs, err := db.GetMsgs(ctx)
		require.NoError(t, err)

		if len(msgs) > 0 {
			return false
		}

		log.Info(ctx, "All messages purged")

		return true
	}, 2*time.Minute, 1*time.Second)
}

// getForkRPCs returns mainnet rpcs urls from env vars.
func getForkRPCs(t *testing.T) map[uint64]string {
	t.Helper()

	notEmpty := func(env string) string {
		v := os.Getenv(env)
		require.NotEmpty(t, v, "%s must be set", env)

		return v
	}

	return map[uint64]string{
		evmchain.IDEthereum:    notEmpty("ETH_RPC"),
		evmchain.IDArbitrumOne: notEmpty("ARB_RPC"),
		evmchain.IDOptimism:    notEmpty("OP_RPC"),
		evmchain.IDBase:        notEmpty("BASE_RPC"),
	}
}

// fund funds accounts with USDC on each chain.
func fund(t *testing.T, ctx context.Context, clients map[uint64]ethclient.Client, account common.Address) {
	t.Helper()

	// Fund USDC
	for chainID, client := range clients {
		amount := bi.Dec6(1000) // 1000 USDC
		usdc := mustUSDC(t, chainID)
		err := anvil.FundUSDC(ctx, client, usdc.Address, amount, account)
		require.NoError(t, err)
		log.Info(ctx, "Funded USDC", "chain", chainID, "amount", amount, "account", account)
	}

	// Fund ETH
	for chainID, client := range clients {
		amount := bi.Ether(1) // 1 ETH
		err := anvil.FundAccounts(ctx, client, amount, account)
		require.NoError(t, err)
		log.Info(ctx, "Funded ETH", "chain", chainID, "amount", amount, "account", account)
	}
}

func getChains(t *testing.T) []evmchain.Metadata {
	t.Helper()

	return []evmchain.Metadata{
		mustMeta(t, evmchain.IDEthereum),
		mustMeta(t, evmchain.IDArbitrumOne),
		mustMeta(t, evmchain.IDOptimism),
		mustMeta(t, evmchain.IDBase),
	}
}

func makeNetwork(t *testing.T, chains []evmchain.Metadata) netconf.Network {
	t.Helper()

	network := netconf.Network{ID: netconf.Devnet}
	network.Chains = make([]netconf.Chain, len(chains))
	for i, chain := range chains {
		network.Chains[i] = netconf.Chain{
			ID:             chain.ChainID,
			Name:           chain.Name,
			BlockPeriod:    chain.BlockPeriod,
			Shards:         []xchain.ShardID{xchain.ShardFinalized0, xchain.ShardLatest0},
			AttestInterval: 10000,
		}
	}

	return network
}

func mustMeta(t *testing.T, chainID uint64) evmchain.Metadata {
	t.Helper()

	meta, ok := evmchain.MetadataByID(chainID)
	require.True(t, ok)

	return meta
}

func mustUSDC(t *testing.T, chainID uint64) tokens.Token {
	t.Helper()
	usdc, ok := tokens.ByAsset(chainID, tokens.USDC)
	require.True(t, ok)

	return usdc
}
