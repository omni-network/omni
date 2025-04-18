package uniswap_test

import (
	"flag"
	"math/big"
	"os"
	"testing"
	"time"

	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/e2e/types"
	"github.com/omni-network/omni/lib/anvil"
	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/tokens"
	"github.com/omni-network/omni/lib/tokens/tokenutil"
	"github.com/omni-network/omni/lib/tutil"
	"github.com/omni-network/omni/lib/uniswap"

	"github.com/stretchr/testify/require"
)

var integration = flag.Bool("integration", false, "run integration tests")

//go:generate go test . -integration -v -run=TestSwapToUSDC

func TestSwapToUSDC(t *testing.T) {
	t.Parallel()
	if !*integration {
		t.Skip("skipping integration test")
	}

	ctx := t.Context()

	meta, ok := evmchain.MetadataByID(evmchain.IDEthereum)
	require.True(t, ok, "ethereum metadata not found")

	solver := eoa.MustAddress(netconf.Devnet, eoa.RoleSolver)

	// Prefer rpc from env, default to public rpc
	rpcURL := os.Getenv("ETHEREUM_RPC_URL")
	if rpcURL == "" {
		rpcURL = types.PublicRPCByName(meta.Name)
	}

	ethCl, stop, err := anvil.Start(ctx, t.TempDir(), evmchain.IDEthereum, anvil.WithFork(rpcURL))
	require.NoError(t, err)
	defer stop()

	backend, err := ethbackend.NewDevBackend(meta.Name, meta.ChainID, time.Second, ethCl)
	require.NoError(t, err)

	wstETH, ok := tokens.ByAsset(evmchain.IDEthereum, tokens.WSTETH)
	require.True(t, ok, "WSTETH token not found")

	usdc, ok := tokens.ByAsset(evmchain.IDEthereum, tokens.USDC)
	require.True(t, ok, "USDC token not found")

	eth1k := bi.Ether(1_000)
	require.NoError(t, anvil.FundAccounts(ctx, ethCl, eth1k, solver))
	require.NoError(t, anvil.FundERC20(ctx, ethCl, wstETH.Address, eth1k, solver))

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
	}

	// Total USDC out
	total := big.NewInt(0)

	// Run tests synchronously, to avoid swaps in same block (reverts with 'Too little received"')
	// TODO(kevin): consider decreasing amountOutMinimum to avoid reverts
	for _, tt := range tests {
		t.Logf("Running test: %s", tt.name)

		token, ok := tokens.ByAsset(evmchain.IDEthereum, tt.asset)
		require.True(t, ok, "%s token not found", tt.asset)

		// Swap to USDC, assert amount out is positive
		amountOut, err := uniswap.SwapToUSDC(ctx, backend, solver, token, tt.amountIn)
		tutil.RequireNoError(t, err)
		tutil.RequireIsPositive(t, amountOut, "amount out should be positive")

		// Add amount out to total
		total.Add(total, amountOut)
	}

	// Assert USDC balance is sum of swaps
	balance, err := tokenutil.BalanceOf(ctx, backend, usdc, solver)
	tutil.RequireNoError(t, err)
	tutil.RequireEQ(t, balance, total, "usdc balance should be equal to amount out")
}
