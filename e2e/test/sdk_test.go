package e2e_test

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/omni-network/omni/e2e/types"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/cometbft/cometbft/test/e2e/pkg/exec"

	"github.com/stretchr/testify/require"
)

func TestSDK(t *testing.T) {
	t.Parallel()
	skipFunc := func(manifest types.Manifest) bool {
		return !manifest.DeploySolve
	}
	maybeTestNetwork(t, skipFunc, func(ctx context.Context, t *testing.T, network netconf.Network, endpoints xchain.RPCEndpoints) {
		t.Helper()

		cwd, err := os.Getwd()
		require.NoError(t, err)

		sdkPath := filepath.Join(cwd, "../../sdk")
		err = os.Chdir(sdkPath)
		require.NoError(t, err)

		err = exec.CommandVerbose(ctx, "pnpm", "install")
		require.NoError(t, err)

		err = exec.CommandVerbose(ctx, "pnpm", "run", "build")
		require.NoError(t, err)

		err = exec.CommandVerbose(ctx, "pnpm", "run", "test:unit")
		require.NoError(t, err)

		err = exec.CommandVerbose(ctx, "pnpm", "run", "test:integration")
		require.NoError(t, err)
	})
}
