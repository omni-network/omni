package app

import (
	"context"
	"log/slog"
	"math/big"
	"os"
	"os/signal"
	"syscall"
	"testing"

	"github.com/omni-network/omni/lib/anvil"
	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/cctp/testutil"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/tokens"
	"github.com/omni-network/omni/lib/tokens/tokenutil"
	"github.com/omni-network/omni/lib/tutil"

	"github.com/ethereum/go-ethereum/common"

	"github.com/stretchr/testify/require"
)

//go:generate go test -run=TestMoveFundsConfig -golden

// TestMoveFundsConfig ensures moveFunds configuration does not change without explicit golden update.
func TestMoveFundsConfig(t *testing.T) {
	t.Parallel()

	golden := map[string]any{
		"moveFundsTo":     moveFundsTo.Hex(),
		"moveFundsChains": moveFundsChains,
	}

	tutil.RequireGoldenJSON(t, golden)
}

// Usage: go test . -integration -v -run=TestMoveFundsIntegration

func TestMoveFundsIntegration(t *testing.T) {
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

	// Pick a chain to test on (using Base as example)
	chainID := evmchain.IDBase
	chainMeta := mustMeta(t, chainID)

	// Get RPC URL
	rpcURL := os.Getenv("BASE_RPC")
	require.NotEmpty(t, rpcURL, "BASE_RPC must be set")

	// Start anvil fork
	ethCl, stop, err := anvil.Start(ctx, tutil.TempDir(t), chainID,
		anvil.WithFork(rpcURL),
		anvil.WithAutoImpersonate(),
		anvil.WithBlockTime(1),
		anvil.WithSlotsInEpoch(2),
	)
	require.NoError(t, err)
	defer stop()

	log.Info(ctx, "Started anvil fork", "chain", chainMeta.Name)

	// Stop anvil on interrupt
	go func() {
		<-ctx.Done()
		stop()
	}()

	// Create test accounts
	solverPk, solver := testutil.NewAccount(t)
	_, targetAddr := testutil.NewAccount(t)

	// Create backends
	clients := map[uint64]ethclient.Client{chainID: ethCl}
	backends, err := ethbackend.BackendsFromClients(clients, solverPk)
	tutil.RequireNoError(t, err)
	backend, err := backends.Backend(chainID)
	tutil.RequireNoError(t, err)

	// Get all tokens for this chain
	allTokens := tokens.ByChain(chainID)
	require.NotEmpty(t, allTokens, "chain should have tokens")

	// Separate tokens into native and ERC20
	var nativeToken tokens.Token
	var erc20Tokens []tokens.Token
	hasNative := false

	for _, token := range allTokens {
		if token.IsNative() {
			nativeToken = token
			hasNative = true
		} else {
			erc20Tokens = append(erc20Tokens, token)
		}
	}

	log.Info(ctx, "Funding solver account", "address", solver.Hex())

	// Fund native token
	if hasNative {
		initialNativeBalance := bi.Ether(10)
		err := anvil.FundAccounts(ctx, ethCl, initialNativeBalance, solver)
		tutil.RequireNoError(t, err)

		balance, err := tokenutil.BalanceOf(ctx, backend, nativeToken, solver)
		tutil.RequireNoError(t, err)

		log.Info(ctx, "Funded native token",
			"token", nativeToken.Symbol,
			"amount", nativeToken.FormatAmt(balance))
	}

	// Fund ERC20 tokens - test with a few tokens
	var fundedERC20s []tokens.Token
	for i, token := range erc20Tokens {
		if i >= 3 { // Fund up to 3 ERC20 tokens for testing
			break
		}

		var amount *big.Int
		if token.Decimals == 6 {
			amount = bi.Dec6(10_000) // 10k for 6 decimal tokens
		} else {
			amount = bi.Ether(10) // 10 for 18 decimal tokens
		}

		err := fundToken(t, ctx, ethCl, token, solver, amount)
		if err != nil {
			log.Warn(ctx, "Failed to fund token", err, "token", token.Symbol)
			continue
		}

		balance, err := tokenutil.BalanceOf(ctx, backend, token, solver)
		tutil.RequireNoError(t, err)

		log.Info(ctx, "Funded ERC20 token",
			"token", token.Symbol,
			"amount", token.FormatAmt(balance))

		fundedERC20s = append(fundedERC20s, token)
	}

	// Record initial balances
	initialBalances := make(map[string]*big.Int)
	for _, token := range fundedERC20s {
		balance, err := tokenutil.BalanceOf(ctx, backend, token, solver)
		tutil.RequireNoError(t, err)
		initialBalances[token.Symbol] = balance
	}

	if hasNative {
		balance, err := tokenutil.BalanceOf(ctx, backend, nativeToken, solver)
		tutil.RequireNoError(t, err)
		initialBalances[nativeToken.Symbol] = balance
	}

	// Run moveFundsToOn
	log.Info(ctx, "Moving funds to target address", "target", targetAddr.Hex())
	err = moveFundsToOn(ctx, backends, solver, targetAddr, []uint64{chainID})
	tutil.RequireNoError(t, err)

	// Assert balances transferred correctly
	log.Info(ctx, "Verifying balances")

	// Check ERC20 tokens - should be fully transferred
	for _, token := range fundedERC20s {
		// Solver should have zero balance
		solverBalance, err := tokenutil.BalanceOf(ctx, backend, token, solver)
		tutil.RequireNoError(t, err)
		require.Equal(t, 0, solverBalance.Sign(),
			"solver should have zero %s balance, got %s", token.Symbol, token.FormatAmt(solverBalance))

		// Target should have the initial balance
		targetBalance, err := tokenutil.BalanceOf(ctx, backend, token, targetAddr)
		tutil.RequireNoError(t, err)
		require.True(t, bi.EQ(targetBalance, initialBalances[token.Symbol]),
			"target should have received %s %s, got %s",
			token.FormatAmt(initialBalances[token.Symbol]), token.Symbol, token.FormatAmt(targetBalance))

		log.Info(ctx, "ERC20 transfer verified",
			"token", token.Symbol,
			"amount", token.FormatAmt(targetBalance))
	}

	// Check native token - should have transferred at least 99% of initial balance
	if hasNative {
		targetNativeBalance, err := tokenutil.BalanceOf(ctx, backend, nativeToken, targetAddr)
		tutil.RequireNoError(t, err)

		// Target should have received at least 99% of the initial native balance
		minExpected := bi.MulF64(initialBalances[nativeToken.Symbol], 0.99) // 99% of initial

		require.True(t, bi.GTE(targetNativeBalance, minExpected),
			"target should have received at least 99%% of initial native balance (%s), got %s",
			nativeToken.FormatAmt(minExpected), nativeToken.FormatAmt(targetNativeBalance))

		solverNativeBalance, err := tokenutil.BalanceOf(ctx, backend, nativeToken, solver)
		tutil.RequireNoError(t, err)

		log.Info(ctx, "Native transfer verified",
			"token", nativeToken.Symbol,
			"initial_balance", nativeToken.FormatAmt(initialBalances[nativeToken.Symbol]),
			"target_received", nativeToken.FormatAmt(targetNativeBalance),
			"solver_remaining", nativeToken.FormatAmt(solverNativeBalance))
	}

	log.Info(ctx, "Integration test passed")
}

// fundToken funds the solver with a specific token.
func fundToken(t *testing.T, ctx context.Context, client ethclient.Client, token tokens.Token, account common.Address, amt *big.Int) error {
	t.Helper()

	if token.Is(tokens.ETH) {
		err := anvil.FundAccounts(ctx, client, amt, account)
		if err != nil {
			return err
		}

		return nil
	}

	if token.Is(tokens.USDC) {
		err := anvil.FundUSDC(ctx, client, token.Address, amt, account)
		if err != nil {
			return err
		}

		return nil
	}

	// L1 WSTETH has _balance map at slot 0
	if token.Is(tokens.WSTETH) && token.ChainID == evmchain.IDEthereum {
		err := anvil.FundERC20(ctx, client, token.Address, amt, account, anvil.WithSlotIdx(0))
		if err != nil {
			return err
		}

		return nil
	}

	// Bridged WSTETH has _balance map at slot 1
	if token.Is(tokens.WSTETH) {
		err := anvil.FundERC20(ctx, client, token.Address, amt, account, anvil.WithSlotIdx(1))
		if err != nil {
			return err
		}

		return nil
	}

	if token.Is(tokens.USDT) && token.ChainID == evmchain.IDEthereum {
		err := anvil.FundL1USDT(ctx, client, token.Address, amt, account)
		if err != nil {
			return err
		}

		return nil
	}

	if token.Is(tokens.USDT) && token.ChainID == evmchain.IDArbitrumOne {
		err := anvil.FundArbUSDT(ctx, client, token.Address, amt, account)
		if err != nil {
			return err
		}

		return nil
	}

	if token.Is(tokens.USDT) && token.ChainID == evmchain.IDOptimism {
		err := anvil.FundOPUSDT(ctx, client, token.Address, amt, account)
		if err != nil {
			return err
		}

		return nil
	}

	// For other standard ERC20 tokens, try default slot 0
	err := anvil.FundERC20(ctx, client, token.Address, amt, account)
	if err != nil {
		return err
	}

	return nil
}

func mustMeta(t *testing.T, chainID uint64) evmchain.Metadata {
	t.Helper()

	meta, ok := evmchain.MetadataByID(chainID)
	require.True(t, ok)

	return meta
}
