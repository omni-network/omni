package e2e_test

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/cometbft/cometbft/test/e2e/pkg/exec"
	"github.com/stretchr/testify/require"
)

func TestSDK(t *testing.T) {
	t.Parallel()
	testNetwork(t, func(ctx context.Context, t *testing.T, network netconf.Network, endpoints xchain.RPCEndpoints) {
		t.Helper()

		cwd, err := os.Getwd()
		require.NoError(t, err)

		sdkPath := filepath.Join(cwd, "../../sdk")
		os.Chdir(sdkPath)

		err = exec.CommandVerbose(ctx, "pnpm", "install")
		require.NoError(t, err)

		err = exec.CommandVerbose(ctx, "pnpm", "run", "test:unit")
		require.NoError(t, err)

		// TODO: inject RPC endpoints as environment variables
		err = exec.CommandVerbose(ctx, "pnpm", "run", "test:integration")
		require.NoError(t, err)
	})
}
