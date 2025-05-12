package rebalance

import (
	"os"
	"testing"

	"github.com/omni-network/omni/lib/anvil"
	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/cctp/testutil"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/tokens"
	"github.com/omni-network/omni/lib/tokens/tokenutil"
	"github.com/omni-network/omni/lib/tutil"

	"github.com/stretchr/testify/require"
)

func TestNoMoreSTETH(t *testing.T) {
	t.Skip() // Comment to run.

	t.Parallel()

	ctx := t.Context()

	rpcs := map[uint64]string{evmchain.IDEthereum: mustEnv(t, "ETH_RPC")}
	chains := []evmchain.Metadata{mustMeta(t, evmchain.IDEthereum)}

	clients, stop := testutil.StartAnvilForks(t, ctx, rpcs, chains)
	defer stop()

	client := clients[evmchain.IDEthereum]
	wsteth := mustToken(evmchain.IDEthereum, tokens.WSTETH)
	steth := mustToken(evmchain.IDEthereum, tokens.STETH)

	solverPk, solver := testutil.NewAccount(t)

	backends, err := ethbackend.BackendsFromClients(clients, solverPk)
	tutil.RequireNoError(t, err)

	backend, err := backends.Backend(evmchain.IDEthereum)
	tutil.RequireNoError(t, err)

	contract, err := NewWSTETH(wsteth.Address, backend)
	tutil.RequireNoError(t, err)

	// First, fund solver with wsteth and unwrap.
	// This is more straightforward than funding steth, which does not have a
	// normal balance storage mapping.
	err = anvil.FundERC20(ctx, client, wsteth.Address, bi.Ether(10), solver, anvil.WithSlotIdx(0))
	tutil.RequireNoError(t, err)

	// Also, fund with gas
	err = anvil.FundAccounts(ctx, client, bi.Ether(1), solver)
	tutil.RequireNoError(t, err)

	// Then, unwrap wsteth to steth.
	txOpts, err := backend.BindOpts(ctx, solver)
	tutil.RequireNoError(t, err)
	tx, err := contract.Unwrap(txOpts, bi.Ether(10))
	tutil.RequireNoError(t, err)

	_, err = backend.WaitMined(ctx, tx)
	tutil.RequireNoError(t, err)

	stethBalance, err := tokenutil.BalanceOf(ctx, backend, steth, solver)
	tutil.RequireNoError(t, err)
	tutil.RequireGT(t, stethBalance, bi.Zero())

	log.Debug(ctx, "STETH balance", "balance", steth.FormatAmt(stethBalance))

	// Now, run nowMoreSTETH.
	err = wrapSTETH(ctx, backends, solver)
	tutil.RequireNoError(t, err)

	// Require that the steth balance is now zero.
	stethBalance, err = tokenutil.BalanceOf(ctx, backend, steth, solver)
	require.NoError(t, err)
	tutil.RequireGTE(t, bi.N(1), stethBalance) // One wei steth may be left over.
}

func mustEnv(t *testing.T, name string) string {
	t.Helper()

	v := os.Getenv(name)
	require.NotEmpty(t, v, "missing env var %s", name)

	return v
}

func mustMeta(t *testing.T, chainID uint64) evmchain.Metadata {
	t.Helper()

	meta, ok := evmchain.MetadataByID(chainID)
	require.True(t, ok)

	return meta
}
