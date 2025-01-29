package e2e_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/omni-network/omni/e2e/solve"
	"github.com/omni-network/omni/e2e/types"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/stretchr/testify/require"
)

// TestSolver submits deposits to the solve inbox and waits for them to be processed.
func TestSolver(t *testing.T) {
	t.Parallel()
	skipFunc := func(manifest types.Manifest) bool {
		return !manifest.DeploySolve
	}
	maybeTestNetwork(t, skipFunc, func(ctx context.Context, t *testing.T, network netconf.Network, endpoints xchain.RPCEndpoints) {
		t.Helper()

		ensureSolverAPILive(t)
		err := solve.TestV2(ctx, network, endpoints)
		require.NoError(t, err)
	})
}

//nolint:noctx // Not an issue in tests
func ensureSolverAPILive(t *testing.T) {
	t.Helper()

	resp, err := http.Get("http://localhost:26661/live")
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode)
	require.NoError(t, resp.Body.Close())
}
