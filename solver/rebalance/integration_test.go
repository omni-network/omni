//nolint:unused // Unused code left for reference
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
	"github.com/omni-network/omni/lib/tokens/tokenutil"
	"github.com/omni-network/omni/lib/tutil"
	"github.com/omni-network/omni/lib/xchain"
	xprovider "github.com/omni-network/omni/lib/xchain/provider"
	"github.com/omni-network/omni/solver/fundthresh"
	"github.com/omni-network/omni/solver/rebalance"

	"github.com/ethereum/go-ethereum/common"

	"github.com/stretchr/testify/require"
)

var (
	integration = flag.Bool("integration", false, "run integration tests")
)

// Usage: go test . -integration -v -run=TestIntegration

// TestIntegration tests integration rebalance package, currently configured
// to move USDC back to Ethereum L1. Original rebalance tests left commented
// out below for reference.
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

	// Fund all tokens except USDC on Ethereum L1
	fundAllExceptEthL1USDC(t, ctx, clients, solver)

	// getEthL1USDCBalance returns USDC balance on Ethereum L1
	getEthL1USDCBalance := func() *big.Int {
		usdc, ok := tokens.ByAsset(evmchain.IDEthereum, tokens.USDC)
		require.True(t, ok)

		client, ok := clients[evmchain.IDEthereum]
		require.True(t, ok)

		balance, err := tokenutil.BalanceOf(ctx, client, usdc, solver)
		tutil.RequireNoError(t, err)

		return balance
	}

	logSnapshot := func() {
		for _, tkn := range rebalance.SwappableTokens() {
			if !cctp.IsSupportedChain(tkn.ChainID) {
				continue
			}

			client, ok := clients[tkn.ChainID]
			require.True(t, ok)

			b, err := tokenutil.BalanceOf(ctx, client, tkn, solver)
			tutil.RequireNoError(t, err)

			bUSD, err := rebalance.AmtToUSD(ctx, pricer, tkn, b)
			tutil.RequireNoError(t, err)

			log.Info(ctx, "Token balance",
				"chain", evmchain.Name(tkn.ChainID),
				"token", tkn.Asset,
				"balance", tkn.FormatAmt(b),
				"usd_value", formatUSD(bUSD))
		}
	}

	// Log initial state
	initialBalance := getEthL1USDCBalance()
	log.Info(ctx, "Starting Ethereum L1 USDC balance", "balance", formatUSDC(initialBalance))

	// Start rebalancing
	interval := 5 * time.Second // Fast interval for testing
	dbDir := ""                 // Use in-memory db
	err = rebalance.Start(ctx, network, cctpClient, pricer, backends, solver, dbDir, rebalance.WithInterval(interval))
	tutil.RequireNoError(t, err)

	// Wait for USDC to return to Ethereum L1
	tutil.RequireEventually(t, ctx, func() bool {
		logSnapshot()

		currentBalance := getEthL1USDCBalance()
		log.Info(ctx, "Current Ethereum L1 USDC balance", "balance", formatUSDC(currentBalance))

		// Target: accumulate most USDC on Ethereum L1
		// Consider > 90% of target as success to account for gas costs and rounding
		targetUSD := float64(150_000) // Should match funding amount
		targetBalance := bi.Dec6(int64(targetUSD))
		threshold := bi.MulF64(targetBalance, 0.9)

		if bi.LT(currentBalance, threshold) {
			log.Info(ctx, "Rebalance not complete",
				"current", formatUSDC(currentBalance),
				"target", formatUSDC(targetBalance),
				"threshold", formatUSDC(threshold))

			return false
		}

		log.Info(ctx, "Rebalance complete - USDC returned to Ethereum L1",
			"balance", formatUSDC(currentBalance),
			"threshold", formatUSDC(threshold))

		return true
	}, 4*time.Minute, 10*time.Second)
}

/*
Comment out original test. Leaving for reference.
Rewriting above to test returning USDC to eth l1 w/ new fund thresholds.

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

		// Log deficits
		for _, tkn := range rebalance.SwappableTokens() {
			if !cctp.IsSupportedChain(tkn.ChainID) {
				continue
			}

			client, ok := clients[tkn.ChainID]
			require.True(t, ok)

			d, err := rebalance.GetUSDDeficit(ctx, client, pricer, tkn, solver)
			tutil.RequireNoError(t, err)

			sum = bi.Add(sum, d)
		}

		return sum
	}

	logSnapshot := func() {
		for _, tkn := range rebalance.SwappableTokens() {
			if !cctp.IsSupportedChain(tkn.ChainID) {
				continue
			}

			client, ok := clients[tkn.ChainID]
			require.True(t, ok)

			d, err := rebalance.GetDeficit(ctx, client, tkn, solver)
			tutil.RequireNoError(t, err)

			dUSD, err := rebalance.AmtToUSD(ctx, pricer, tkn, d)
			tutil.RequireNoError(t, err)

			s, err := rebalance.GetSurplus(ctx, client, tkn, solver)
			tutil.RequireNoError(t, err)

			sUSD, err := rebalance.AmtToUSD(ctx, pricer, tkn, s)
			tutil.RequireNoError(t, err)

			log.Info(ctx, "Token snapshot",
				"chain", evmchain.Name(tkn.ChainID),
				"token", tkn.Asset,
				"deficit", tkn.FormatAmt(d),
				"deficit_usd", formatUSD(dUSD),
				"surplus", tkn.FormatAmt(s),
				"surplus_usd", formatUSD(sUSD))
		}
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
		logSnapshot()

		deficit := sumDeficits()

		// Consider < 20k deficit as "rebalanced" to reduce flaps.
		// TODO(kevin): fix flaps
		if bi.GT(deficit, bi.Dec6(20000)) {
			log.Info(ctx, "Rebalance not complete", "deficit", formatUSD(deficit))
			return false
		}

		log.Info(ctx, "Rebalance complete", "deficit", formatUSD(deficit))

		return true
	}, 4*time.Minute, 10*time.Second)
}

*/

// fundAllExceptEthL1USDC funds all tokens except USDC on Ethereum L1.
// The goal is to have USDC distributed on other chains, and rebalance should move it to Ethereum L1.
func fundAllExceptEthL1USDC(t *testing.T, ctx context.Context, clients map[uint64]ethclient.Client, solver common.Address) {
	t.Helper()

	// Make sure we have enough ETH for gas on all chains
	for _, client := range clients {
		err := anvil.FundAccounts(ctx, client, bi.Ether(1), solver)
		tutil.RequireNoError(t, err)
	}

	totalFundingUSD := float64(150_000) // Total USDC to distribute across chains (except Ethereum L1)

	// Get all USDC tokens except Ethereum L1
	var usdcTokens []tokens.Token
	for _, token := range rebalance.SwappableTokens() {
		if !token.Is(tokens.USDC) {
			continue
		}
		if token.ChainID == evmchain.IDEthereum {
			continue
		}
		if !cctp.IsSupportedChain(token.ChainID) {
			continue
		}
		usdcTokens = append(usdcTokens, token)
	}

	require.NotEmpty(t, usdcTokens, "need at least one USDC token to fund")

	// Distribute USDC evenly across other chains
	perChainUSD := totalFundingUSD / float64(len(usdcTokens))

	for _, token := range usdcTokens {
		// Fund with target threshold + surplus amount
		thresh := fundthresh.Get(token)
		surplusAmt := bi.Dec6(int64(perChainUSD))
		toFund := bi.Add(thresh.Target(), surplusAmt)

		fundToken(t, ctx, clients[token.ChainID], token, solver, toFund)

		log.Info(ctx, "Funded USDC surplus",
			"chain", evmchain.Name(token.ChainID),
			"token", token.Asset,
			"amount", token.FormatAmt(toFund),
			"thresh_target", token.FormatAmt(thresh.Target()),
			"surplus", token.FormatAmt(surplusAmt))
	}

	// Fund Ethereum L1 with all other tokens at their target thresholds
	// This ensures USDC is the only token with surplus that needs rebalancing
	for _, token := range rebalance.SwappableTokens() {
		if !cctp.IsSupportedChain(token.ChainID) {
			continue
		}
		if token.ChainID != evmchain.IDEthereum {
			continue
		}
		if token.Is(tokens.USDC) {
			// Don't fund Ethereum L1 USDC - we want it to accumulate there
			continue
		}

		thresh := fundthresh.Get(token)
		toFund := thresh.Target()

		if bi.LTE(toFund, bi.Zero()) {
			continue
		}

		fundToken(t, ctx, clients[token.ChainID], token, solver, toFund)

		log.Info(ctx, "Funded Ethereum L1 token at target",
			"chain", evmchain.Name(token.ChainID),
			"token", token.Asset,
			"amount", token.FormatAmt(toFund),
			"thresh_target", token.FormatAmt(thresh.Target()))
	}
}

func formatUSDC(amt *big.Int) string {
	return formatUSD(amt)
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

			for i, token := range shuffle(rebalance.SwappableTokens()) {
				if !cctp.IsSupportedChain(token.ChainID) {
					continue
				}

				if i%2 == 0 {
					toDeficit = append(toDeficit, token)
				}
				if i%2 == 1 {
					toSurplus = append(toSurplus, token)
				}
			}

			// filter out tokens we can never surplus
			toSurplus = filter(toSurplus, func(t tokens.Token) bool {
				return !fundthresh.Get(t).NeverSurplus()
			})

			// retry, should generally not happen
			// Need at least one token to surplus
			if len(toSurplus) == 0 {
				continue
			}

			return toDeficit, toSurplus
		}
	}()

	// Goal total deficit / surplus
	//
	// Surplus more than deficit, so that avoid case in which surpluses are
	// swapped to USDC, but USDC remains in deficit, and is not used to
	// rebalance.
	totalDeficitUSD := float64(100_000)
	totalSurplusUSD := float64(150_000)

	// Fund deficit tokens
	for _, token := range toDeficit {
		toDeficitUSD := totalDeficitUSD / float64(len(toDeficit))

		price, err := pricer.USDPrice(ctx, token.Asset)
		tutil.RequireNoError(t, err)

		thresh := fundthresh.Get(token)
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
		thresh := fundthresh.Get(token)
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

	if token.Is(tokens.USDT) && token.ChainID == evmchain.IDEthereum {
		err := anvil.FundL1USDT(ctx, client, token.Address, amt, account)
		tutil.RequireNoError(t, err)

		return
	}

	if token.Is(tokens.USDT) && token.ChainID == evmchain.IDArbitrumOne {
		err := anvil.FundArbUSDT(ctx, client, token.Address, amt, account)
		tutil.RequireNoError(t, err)

		return
	}

	if token.Is(tokens.USDT) && token.ChainID == evmchain.IDOptimism {
		err := anvil.FundOPUSDT(ctx, client, token.Address, amt, account)
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
		evmchain.IDMantle:      notEmpty("MANTLE_RPC"),
		evmchain.IDHyperEVM:    notEmpty("HYPER_EVM_RPC"),
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
