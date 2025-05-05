package rebalance_test

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
	"github.com/omni-network/omni/lib/cctp/testutil"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/tokenpricer"
	"github.com/omni-network/omni/lib/tokenpricer/coingecko"
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

	// Fund L1 USDC at suprlus threshold, so we have some to swap after briding from base
	err = anvil.FundUSDC(ctx, clients[evmchain.IDEthereum], l1USDC.Address, rebalance.GetFundThreshold(l1USDC).Surplus(), solver)
	tutil.RequireNoError(t, err)

	// Fund 10 base wstETH above surplus
	err = anvil.FundERC20(ctx, clients[evmchain.IDBase], baseWSTETH.Address,
		bi.Add(rebalance.GetFundThreshold(baseWSTETH).Surplus(), bi.Ether(10)),
		solver,
		anvil.WithSlotIdx(1))
	tutil.RequireNoError(t, err)

	// Fund L1 wstETH at 8 below target
	// We use a deficit < base surplus, because not all surplus will be moved (due to min / max swap limits)
	err = anvil.FundERC20(ctx, clients[evmchain.IDEthereum], l1WSTETH.Address,
		bi.Sub(rebalance.GetFundThreshold(l1WSTETH).Target(), bi.Ether(8)),
		solver)
	tutil.RequireNoError(t, err)

	// Fund base USDC at surplus threshold, so we can send some after swapping wstETH
	err = anvil.FundUSDC(ctx, clients[evmchain.IDBase], baseUSDC.Address, rebalance.GetFundThreshold(baseUSDC).Surplus(), solver)
	tutil.RequireNoError(t, err)

	// Start rebalancing
	interval := 5 * time.Second // fast interval for testing
	dbDir := ""                 // use in-mem db
	err = rebalance.Start(ctx, network, cctpClient, newPricer(ctx), backends, solver, dbDir, rebalance.WithInterval(interval))
	tutil.RequireNoError(t, err)

	must := func(amt *big.Int, err error) *big.Int {
		tutil.RequireNoError(t, err)
		return amt
	}

	balance := func(chainID uint64, token tokens.Token) string {
		return token.FormatAmt(must(tokenutil.BalanceOf(ctx, clients[chainID], token, solver)))
	}

	// Wait for rebalance
	tutil.RequireEventually(t, ctx, func() bool {
		// Surplus base WSTETH should < min swap
		if bi.LT(
			must(rebalance.GetSurplus(ctx, clients[evmchain.IDBase], baseWSTETH, solver)),
			rebalance.GetFundThreshold(baseWSTETH).MinSwap(),
		) {
			return false
		}

		// L1 wsteth should not be in deficit
		deficit := must(rebalance.GetDeficit(ctx, clients[evmchain.IDEthereum], l1WSTETH, solver))
		if !bi.IsZero(deficit) {
			log.Info(ctx, "L1 wstETH deficit",
				"deficit", l1WSTETH.FormatAmt(deficit),
				"balance", balance(evmchain.IDEthereum, l1WSTETH))

			return false
		}

		// Log results
		log.Info(ctx, "Rebalance complete",
			"base_wsteth", balance(evmchain.IDBase, baseWSTETH),
			"l1_usdc", balance(evmchain.IDEthereum, l1USDC),
			"l1_usdc", balance(evmchain.IDEthereum, l1WSTETH))

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

func newPricer(ctx context.Context) tokenpricer.Pricer {
	apiKey := os.Getenv("COINGECKO_API_KEY")
	pricer := tokenpricer.NewCached(coingecko.New(coingecko.WithAPIKey(apiKey)))

	// use cached pricer avoid spamming coingecko public api
	const priceCacheEvictInterval = time.Minute * 10
	go pricer.ClearCacheForever(ctx, priceCacheEvictInterval)

	return pricer
}
