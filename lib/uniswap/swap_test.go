package uniswap_test

import (
	"crypto/ecdsa"
	"flag"
	"log/slog"
	"math/big"
	"os"
	"testing"
	"time"

	"github.com/omni-network/omni/lib/anvil"
	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/tokens"
	"github.com/omni-network/omni/lib/tokens/tokenutil"
	"github.com/omni-network/omni/lib/tutil"
	"github.com/omni-network/omni/lib/uniswap"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/stretchr/testify/require"
)

var integration = flag.Bool("integration", false, "run integration tests")

//go:generate go test . -integration -v -run=TestSwapToUSDC

// TestSwapUSDC tests the SwapToUSDC and SwapUSDCTo functions.
func TestSwapUSDC(t *testing.T) {
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

	meta, ok := evmchain.MetadataByID(evmchain.IDEthereum)
	require.True(t, ok, "ethereum metadata not found")

	swapperPk, swapper := newAccount(t)

	rpcURL := os.Getenv("ETH_RPC")
	require.NotEmpty(t, rpcURL, "ETH_RPC required")

	ethCl, stop, err := anvil.Start(ctx, tutil.TempDir(t), evmchain.IDEthereum, anvil.WithFork(rpcURL))
	tutil.RequireNoError(t, err)
	defer stop()

	backend, err := ethbackend.NewBackend(meta.Name, meta.ChainID, time.Second, ethCl, swapperPk)
	tutil.RequireNoError(t, err)

	wstETH, ok := tokens.ByAsset(evmchain.IDEthereum, tokens.WSTETH)
	require.True(t, ok, "WSTETH token not found")

	usdc, ok := tokens.ByAsset(evmchain.IDEthereum, tokens.USDC)
	require.True(t, ok, "USDC token not found")

	usdt, ok := tokens.ByAsset(evmchain.IDEthereum, tokens.USDT)
	require.True(t, ok, "USDT token not found")

	eth1k := bi.Ether(1_000)
	tutil.RequireNoError(t, anvil.FundAccounts(ctx, ethCl, eth1k, swapper))
	tutil.RequireNoError(t, anvil.FundERC20(ctx, ethCl, wstETH.Address, eth1k, swapper))
	tutil.RequireNoError(t, anvil.FundUSDT(ctx, ethCl, usdt.Address, bi.Dec6(1000), swapper))

	tests := []struct {
		name     string
		asset    tokens.Asset
		amountIn *big.Int
	}{
		{
			name:     "ETH to USDC",
			asset:    tokens.ETH,
			amountIn: bi.Ether(1),
		},
		{
			name:     "WSTETH to USDC",
			asset:    tokens.WSTETH,
			amountIn: bi.Ether(1),
		},
		{
			name:     "USDT to USDC",
			asset:    tokens.USDT,
			amountIn: bi.Dec6(1),
		},
	}

	// Run tests synchronously, to avoid swaps in same block (reverts with 'Too little received"')
	for _, tt := range tests {
		t.Logf("Running test: %s", tt.name)

		token, ok := tokens.ByAsset(evmchain.IDEthereum, tt.asset)
		require.True(t, ok, "%s token not found", tt.asset)

		// Token balance pre-swap
		balanceIn, err := tokenutil.BalanceOf(ctx, backend, token, swapper)
		tutil.RequireNoError(t, err)

		// Swap to USDC
		usdcOutMin, err := uniswap.SwapToUSDC(ctx, backend, swapper, token, tt.amountIn)
		tutil.RequireNoError(t, err)
		tutil.RequireIsPositive(t, usdcOutMin)

		// Make sure we actually swapped for USDC
		// We can assert out == min received, because no other swaps in same block
		tutil.RequireEQ(t, usdcOutMin, balanceOf(t, backend, usdc, swapper))

		// Check we lost some input token
		balanceOut := balanceOf(t, backend, token, swapper)
		tutil.RequireGT(t, balanceIn, balanceOut)

		// Swap back to original token (re-use amountIn)
		usdcInMax, err := uniswap.SwapUSDCTo(ctx, backend, swapper, token, bi.Div(tt.amountIn, big.NewInt(2))) // divide by 2, to make sure we don't exceed balance we just got
		tutil.RequireNoError(t, err)
		tutil.RequireIsPositive(t, usdcInMax)

		// USDC balance should be back to zero
		tutil.RequireEQ(t, big.NewInt(0), balanceOf(t, backend, usdc, swapper))

		// Check we got back the original token
		tutil.RequireGT(t, balanceOf(t, backend, token, swapper), balanceOut)
	}
}

func balanceOf(t *testing.T, backend *ethbackend.Backend, token tokens.Token, addr common.Address) *big.Int {
	t.Helper()

	balance, err := tokenutil.BalanceOf(t.Context(), backend, token, addr)
	require.NoError(t, err)

	return balance
}

func newAccount(t *testing.T) (*ecdsa.PrivateKey, common.Address) {
	t.Helper()

	pk, err := crypto.GenerateKey()
	require.NoError(t, err)
	addr := crypto.PubkeyToAddress(pk.PublicKey)

	return pk, addr
}
