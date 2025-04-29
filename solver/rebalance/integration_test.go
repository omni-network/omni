package rebalance_test

import (
	"context"
	"flag"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"testing"
	"time"

	"github.com/omni-network/omni/lib/anvil"
	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/cctp"
	"github.com/omni-network/omni/lib/cctp/testutil"
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
	"github.com/omni-network/omni/solver/rebalance"

	"github.com/ethereum/go-ethereum/common"

	"github.com/stretchr/testify/require"
)

var (
	integration = flag.Bool("integration", false, "run integration tests")
)

//go:generate go test . -integration -v -run=TestIntegration

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
	tutil.RequireNoError(t, err)

	// Handle interrupts
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer cancel()

	rpcs := getRPCs(t)
	chains := getChains(t)
	network := makeNetwork(t, chains)

	clients, stop := testutil.StartAnvilForks(t, ctx, rpcs, chains)
	defer stop()

	// Stop anvil on interrupt
	go func() {
		<-ctx.Done()
		stop()
	}()

	solverPk, solver := testutil.NewAccount(t)
	fundETH(t, ctx, clients, solver)

	backends, err := ethbackend.BackendsFromClients(clients, solverPk)
	tutil.RequireNoError(t, err)
	xprov := xprovider.New(network, clients, nil)

	// Start attesting CCTP messages
	cctpClient := cctp.StartTestClient(ctx, t, xprov, chains, clients)

	l1USDC := mustToken(t, evmchain.IDEthereum, tokens.USDC)
	l1WSTETH := mustToken(t, evmchain.IDEthereum, tokens.WSTETH)
	baseWSTETH := mustToken(t, evmchain.IDBase, tokens.WSTETH)
	baseUSDC := mustToken(t, evmchain.IDBase, tokens.USDC)

	// Fund L1 USDC above suprluse threshold, so we have some to swap after briding from base
	err = anvil.FundUSDC(ctx, clients[evmchain.IDEthereum], l1USDC.Address, rebalance.GetFundThreshold(l1USDC).Surplus(), solver)
	tutil.RequireNoError(t, err)

	// Fund 10 base wstETH above surplus
	surplusBaseWSETH := bi.Ether(10)
	err = anvil.FundERC20(ctx, clients[evmchain.IDBase], baseWSTETH.Address,
		bi.Add(rebalance.GetFundThreshold(baseWSTETH).Surplus(), surplusBaseWSETH), // fund above surplus
		solver,
		anvil.WithSlotIdx(1)) // wstETH balance map at slot 1
	tutil.RequireNoError(t, err)

	// Fund base USDC at surplus threshold, so we can send some after swapping wstETH
	err = anvil.FundUSDC(ctx, clients[evmchain.IDBase], baseUSDC.Address, rebalance.GetFundThreshold(baseUSDC).Surplus(), solver)
	tutil.RequireNoError(t, err)

	// Start rebalancing
	cfg := rebalance.Config{Interval: 5 * time.Second} // fast interval for testing
	dbDir := ""                                        // use in-mem db
	err = rebalance.Start(ctx, cfg, network, cctpClient, backends, solver, dbDir)
	tutil.RequireNoError(t, err)

	// Wait for rebalance
	tutil.RequireEventually(t, ctx, func() bool {
		baseWSTETHBal, err := tokenutil.BalanceOf(ctx, clients[evmchain.IDBase], baseWSTETH, solver)
		tutil.RequireNoError(t, err)

		l1USDCBal, err := tokenutil.BalanceOf(ctx, clients[evmchain.IDEthereum], l1USDC, solver)
		tutil.RequireNoError(t, err)

		l1WSTETHBal, err := tokenutil.BalanceOf(ctx, clients[evmchain.IDEthereum], l1WSTETH, solver)
		tutil.RequireNoError(t, err)

		// All base WSTETH should be moved
		if !bi.IsZero(baseWSTETHBal) {
			return false
		}

		// All base WSTETH should be on L1 (almost) - some lost to swap fees and uni pool price differences
		if !bi.GT(l1WSTETHBal, bi.Sub(surplusBaseWSETH, bi.Ether(0.1))) {
			return false
		}

		log.Info(ctx, "Rebalance complete",
			"base_wsteth", baseWSTETH.FormatAmt(baseWSTETHBal),
			"l1_usdc", l1USDC.FormatAmt(l1USDCBal),
			"l1_usdc", l1WSTETH.FormatAmt(l1WSTETHBal))

		return true
	}, 2*time.Minute, 5*time.Second)
}

func fundETH(t *testing.T, ctx context.Context, clients map[uint64]ethclient.Client, account common.Address) {
	t.Helper()

	// Fund ETH
	for chainID, client := range clients {
		amount := bi.Ether(1) // 1 ETH
		err := anvil.FundAccounts(ctx, client, amount, account)
		require.NoError(t, err)
		log.Info(ctx, "Funded ETH", "chain", chainID, "amount", amount, "account", account)
	}
}

// getRPCs returns mainnet rpcs urls from env vars.
func getRPCs(t *testing.T) map[uint64]string {
	t.Helper()

	notEmpty := func(env string) string {
		v := os.Getenv(env)
		require.NotEmpty(t, v, "%s must be set", env)

		return v
	}

	return map[uint64]string{
		evmchain.IDEthereum:    notEmpty("ETH_RPC"),
		evmchain.IDBase:        notEmpty("BASE_RPC"),
		evmchain.IDArbitrumOne: notEmpty("ARB_RPC"),
		evmchain.IDOptimism:    notEmpty("OP_RPC"),
	}
}

func getChains(t *testing.T) []evmchain.Metadata {
	t.Helper()

	return []evmchain.Metadata{
		mustMeta(t, evmchain.IDEthereum),
		mustMeta(t, evmchain.IDBase),
	}
}

func makeNetwork(t *testing.T, chains []evmchain.Metadata) netconf.Network {
	t.Helper()

	network := netconf.Network{ID: netconf.Mainnet}
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

func mustToken(t *testing.T, chainID uint64, asset tokens.Asset) tokens.Token {
	t.Helper()
	token, ok := tokens.ByAsset(chainID, asset)
	require.True(t, ok)

	return token
}
