package e2e_test

import (
	"context"
	"testing"

	"github.com/omni-network/omni/e2e/solve/devapp"
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
	maybeTestNetwork(t, skipFunc, func(t *testing.T, network netconf.Network, endpoints xchain.RPCEndpoints) {
		t.Helper()
		err := devapp.TestFlow(context.Background(), network, endpoints)
		require.NoError(t, err)
	})
}
