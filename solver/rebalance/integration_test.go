package rebalance_test

import (
	"context"
	"flag"
	"log/slog"
	"math/big"
	"math/rand"
	"os"
	"os/signal"
	"slices"
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
	pricer := newPricer(ctx)

	clients, stop := testutil.StartAnvilForks(t, ctx, rpcs, chains)
	defer stop()

	// Stop anvil on interrupt
	go func() {
		<-ctx.Done()
		stop()
	}()

	solverPk, solver := testutil.NewAccount(t)

	backends, err := ethbackend.BackendsFromClients(clients, solverPk)
	tutil.RequireNoError(t, err)
	xprov := xprovider.New(network, clients, nil)

	// Start attesting CCTP messages
	cctpClient := cctp.StartTestClient(ctx, t, xprov, chains, clients)

	// Unbalance solver
	fundUnbalanced(t, ctx, pricer, clients, solver)

	// sumDeficits returns the token deficits in
	sumDeficits := func() *big.Int {
		sum := bi.Zero()

		for _, tkn := range rebalance.Tokens() {
			d, err := rebalance.GetUSDDeficit(ctx, clients[tkn.ChainID], pricer, tkn, solver)
			tutil.RequireNoError(t, err)
			sum = bi.Add(sum, d)
		}

		return sum
	}

	// Confirm unbalance
	deficit := sumDeficits()
	log.Info(ctx, "Starting deficit", "deficit", formatUSD(deficit))
	tutil.RequireGT(t, deficit, bi.Dec6(50_000)) // Require at least 50k deficit

	// Start rebalancing
	interval := 5 * time.Second // Fast interval for testing
	dbDir := ""                 // Use in-memory db
	err = rebalance.Start(ctx, network, cctpClient, pricer, backends, solver, dbDir, rebalance.WithInterval(interval))
	tutil.RequireNoError(t, err)

	// Wait for rebalance
	tutil.RequireEventually(t, ctx, func() bool {
		deficit := sumDeficits()

		// Consider < 10k deficit as "rebalanced"
		// Wiggle room accounts for min swaps & sends
		if bi.GT(deficit, bi.Dec6(10000)) {
			log.Info(ctx, "Rebalance not complete", "deficit", formatUSD(deficit))
			return false
		}

		log.Info(ctx, "Rebalance complete", "deficit", formatUSD(deficit))

		return true
	}, 2*time.Minute, 5*time.Second)
}

// fundUnbalanced funds the solver w/ unbalanced tokens (based on threshold values).
func fundUnbalanced(t *testing.T, ctx context.Context, pricer tokenpricer.Pricer, clients map[uint64]ethclient.Client, solver common.Address) {
	t.Helper()

	// Make sure we have enough ETH for gas on all chains
	for _, client := range clients {
		err := anvil.FundAccounts(ctx, client, bi.Ether(1), solver)
		tutil.RequireNoError(t, err)
	}

	// return list of tokens to defict and surplus
	toDeficit, toSurplus := func() ([]tokens.Token, []tokens.Token) {
		for { // retry until we can return
			var toDeficit []tokens.Token
			var toSurplus []tokens.Token

			for i, token := range shuffle(rebalance.Tokens()) {
				if i%2 == 0 {
					toDeficit = append(toDeficit, token)
				}
				if i%2 == 1 {
					toSurplus = append(toSurplus, token)
				}
			}

			// filter out tokens we can never surplus
			toSurplus = filter(toSurplus, func(t tokens.Token) bool {
				return !rebalance.GetFundThreshold(t).NeverSurplus()
			})

			// retry, should generally not happen
			// Need at least one token to surplus
			if len(toSurplus) == 0 {
				continue
			}

			return toDeficit, toSurplus
		}
	}()

	// Goal total deficit / surplus (should match, so we can rebalance)
	// NOTE: target thresholds may prever us from matching defict target
	totalDeficitUSD := float64(100_000)
	totalSurplusUSD := float64(100_000)

	// Fund deficit tokens
	for _, token := range toDeficit {
		toDeficitUSD := totalDeficitUSD / float64(len(toDeficit))

		price, err := pricer.USDPrice(ctx, token.Asset)
		tutil.RequireNoError(t, err)

		thresh := rebalance.GetFundThreshold(token)
		toDeficit := bi.MulF64(oneOf(token), toDeficitUSD/price)
		toFund := bi.Sub(thresh.Target(), toDeficit)

		// Target threshold < desired deficit - don't fund and move on
		if bi.LTE(toFund, bi.Zero()) {
			continue
		}

		fundToken(t, ctx, clients[token.ChainID], token, solver, toFund)

		log.Info(ctx, "Funded deficit",
			"chain", evmchain.Name(token.ChainID),
			"token", token.Asset,
			"amount", token.FormatAmt(toFund),
			"thresh_target", token.FormatAmt(thresh.Target()),
			"deficit", token.FormatAmt(toDeficit))
	}

	// Fund surplus tokens
	for _, token := range toSurplus {
		thresh := rebalance.GetFundThreshold(token)
		require.False(t, thresh.NeverSurplus())

		toSurplusUSD := totalSurplusUSD / float64(len(toSurplus))

		price, err := pricer.USDPrice(ctx, token.Asset)
		tutil.RequireNoError(t, err)

		toSurplus := bi.MulF64(oneOf(token), toSurplusUSD/price)
		toFund := bi.Add(thresh.Surplus(), toSurplus)

		fundToken(t, ctx, clients[token.ChainID], token, solver, toFund)

		log.Info(ctx, "Funded surplus",
			"chain", evmchain.Name(token.ChainID),
			"token", token.Asset,
			"amount", token.FormatAmt(toFund),
			"thresh_surplus", token.FormatAmt(thresh.Surplus()),
			"surplus", token.FormatAmt(toSurplus))
	}
}

// fundToken funds the solver with a specific token.
func fundToken(t *testing.T, ctx context.Context, client ethclient.Client, token tokens.Token, account common.Address, amt *big.Int) {
	t.Helper()

	if token.Is(tokens.ETH) {
		err := anvil.FundAccounts(ctx, client, amt, account)
		tutil.RequireNoError(t, err)

		return
	}

	if token.Is(tokens.USDC) {
		err := anvil.FundUSDC(ctx, client, token.Address, amt, account)
		tutil.RequireNoError(t, err)

		return
	}

	// L1 WSTETH has _balance map at slot
	if token.Is(tokens.WSTETH) && token.ChainID == evmchain.IDEthereum {
		err := anvil.FundERC20(ctx, client, token.Address, amt, account, anvil.WithSlotIdx(0))
		tutil.RequireNoError(t, err)

		return
	}

	// Bridged WSTETH has _balance map at slot
	if token.Is(tokens.WSTETH) {
		err := anvil.FundERC20(ctx, client, token.Address, amt, account, anvil.WithSlotIdx(1))
		tutil.RequireNoError(t, err)

		return
	}

	if token.Is(tokens.USDT) {
		err := anvil.FundUSDT(ctx, client, token.Address, amt, account)
		tutil.RequireNoError(t, err)

		return
	}

	t.Fatalf("unsupported token %s", token)
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
		mustMeta(t, evmchain.IDArbitrumOne),
		mustMeta(t, evmchain.IDOptimism),
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

func newPricer(ctx context.Context) tokenpricer.Pricer {
	apiKey := os.Getenv("COINGECKO_API_KEY")
	pricer := tokenpricer.NewCached(coingecko.New(coingecko.WithAPIKey(apiKey)))

	// use cached pricer avoid spamming coingecko public api
	const priceCacheEvictInterval = time.Minute * 10
	go pricer.ClearCacheForever(ctx, priceCacheEvictInterval)

	return pricer
}

func shuffle[T any](xs []T) []T {
	clone := slices.Clone(xs)

	rand.Shuffle(len(clone), func(i, j int) {
		clone[i], clone[j] = clone[j], clone[i]
	})

	return clone
}

func filter[T any](xs []T, f func(T) bool) []T {
	var out []T
	for _, x := range xs {
		if f(x) {
			out = append(out, x)
		}
	}

	return out
}

func oneOf(t tokens.Token) *big.Int {
	if t.Decimals == 6 {
		return bi.Dec6(1)
	}

	return bi.Ether(1)
}
