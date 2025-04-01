package e2e_test

import (
	"context"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/omni-network/omni/e2e/types"

	"github.com/stretchr/testify/require"
)

func TestSDK(t *testing.T) {
	t.Parallel()
	skipFunc := func(manifest types.Manifest) bool {
		return !manifest.DeploySolve
	}
	maybeTestNetwork(t, skipFunc, func(ctx context.Context, t *testing.T, _ NetworkDeps) {
		t.Helper()

		sdkPath, err := filepath.Abs("../../sdk")
		require.NoError(t, err)

		exe := func(ctx context.Context, name string, args ...string) {
			cmd := exec.CommandContext(ctx, name, args...)
			cmd.Dir = sdkPath
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr

			err := cmd.Run()
			require.NoError(t, err, "failed to run command: %s %v", name, args)
		}

		exe(ctx, "pnpm", "install")
		exe(ctx, "pnpm", "run", "build")
		exe(ctx, "pnpm", "run", "test:unit")
		exe(ctx, "pnpm", "run", "test:integration")
	})
}
