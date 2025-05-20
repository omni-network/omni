package e2e_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/e2e/solve"
	"github.com/omni-network/omni/e2e/types"
	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/contracts"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/tokens"
	"github.com/omni-network/omni/lib/tutil"
	"github.com/omni-network/omni/lib/umath"

	"github.com/ethereum/go-ethereum/common"

	"github.com/stretchr/testify/require"
)

// TestSolver submits deposits to the solve inbox and waits for them to be processed.
func TestSolver(t *testing.T) {
	t.Parallel()
	skipFunc := func(manifest types.Manifest) bool {
		return !manifest.AllE2ETests
	}
	maybeTestNetwork(t, skipFunc, func(ctx context.Context, t *testing.T, deps NetworkDeps) {
		t.Helper()

		ensureSolverAPILive(t, deps.SolverAddr)
		testContractsAPI(ctx, t, deps.SolverAddr)
		testSolverApprovals(ctx, t, deps)

		err := solve.Test(ctx, deps.Network, deps.Backends, deps.SolverAddr)
		tutil.RequireNoError(t, err)
	})
}

func testContractsAPI(ctx context.Context, t *testing.T, solverAddr string) {
	t.Helper()

	addrs, err := contracts.GetAddresses(ctx, netconf.Devnet)
	require.NoError(t, err)

	uri, err := url.JoinPath(solverAddr, "/api/v1/contracts")
	require.NoError(t, err)
	resp, err := http.Get(uri)
	require.NoError(t, err)

	body := make(map[string]any)
	require.NoError(t, json.NewDecoder(resp.Body).Decode(&body))
	require.NoError(t, resp.Body.Close())

	require.Equal(t, http.StatusOK, resp.StatusCode)
	addrEqual := func(addr common.Address, name string) {
		// Golang common.Address marshalls to lower case (not EIP55).
		require.Equal(t, strings.ToLower(addr.Hex()), body[name], name)
	}
	addrEqual(addrs.Portal, "portal")
	addrEqual(addrs.SolverNetInbox, "inbox")
	addrEqual(addrs.SolverNetOutbox, "outbox")
	// Middleman deprecated and logic moved to executor, temporarily retained for backwards compatibility.
	addrEqual(addrs.SolverNetExecutor, "middleman")
	addrEqual(addrs.SolverNetExecutor, "executor")
}

func testSolverApprovals(ctx context.Context, t *testing.T, deps NetworkDeps) {
	t.Helper()

	network := deps.Network.ID

	addrs, err := contracts.GetAddresses(ctx, network)
	require.NoError(t, err)

	solverAddr := eoa.MustAddress(network, eoa.RoleSolver)

	for _, tkn := range tokens.All() {
		backend, err := deps.Backends.Backend(tkn.ChainID)
		if err != nil {
			continue // Tokens are defines for all networks.
		}

		isDeployed, err := contracts.IsDeployed(ctx, backend, tkn.Address)
		require.NoError(t, err)

		if !isDeployed {
			continue
		}

		erc20, err := bindings.NewIERC20(tkn.Address, backend)
		require.NoError(t, err)

		allowance, err := erc20.Allowance(nil, solverAddr, addrs.SolverNetOutbox)
		require.NoError(t, err)

		// must be max allowance
		expected := bi.MulF64(umath.MaxUint256, 0.9)
		tutil.RequireGT(t, allowance, expected, "allowance missing: chain=%s, allowance=%s", backend.Name(), tkn.FormatAmt(allowance))
	}
}

func ensureSolverAPILive(t *testing.T, solverAddr string) {
	t.Helper()

	uri, err := url.JoinPath(solverAddr, "/live")
	require.NoError(t, err)
	resp, err := http.Get(uri)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode)
	require.NoError(t, resp.Body.Close())
}
