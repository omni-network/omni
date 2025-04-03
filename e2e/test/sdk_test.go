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

		solverContractsPath, err := filepath.Abs("../../contracts/solve")
		require.NoError(t, err)

		sdkPath, err := filepath.Abs("../../sdk")
		require.NoError(t, err)

		writeABI := func(ctx context.Context, contract string) {
			outfile, err := os.Create(filepath.Join(sdkPath, "integration-tests/assets", contract+".json"))
			require.NoError(t, err, "failed to create ABI file for contract: %s", contract)
			defer outfile.Close()

			cmd := exec.CommandContext(ctx, "forge", "inspect", "--json", "src/"+contract+".sol:"+contract, "abi")
			cmd.Dir = solverContractsPath
			cmd.Stdout = outfile
			cmd.Stderr = os.Stderr

			err = cmd.Run()
			require.NoError(t, err, "failed to write ABI file for contract: %s", contract)
		}

		exe := func(ctx context.Context, dir string, name string, args ...string) {
			cmd := exec.CommandContext(ctx, name, args...)
			cmd.Dir = dir
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr

			err := cmd.Run()
			require.NoError(t, err, "failed to run command: %s %v", name, args)
		}

		exe(ctx, solverContractsPath, "pnpm", "install")
		writeABI(ctx, "SolverNetInbox")
		writeABI(ctx, "SolverNetMiddleman")
		writeABI(ctx, "SolverNetOutbox")

		exe(ctx, sdkPath, "pnpm", "install")
		exe(ctx, sdkPath, "pnpm", "run", "build")
		exe(ctx, sdkPath, "pnpm", "run", "test:unit")
		exe(ctx, sdkPath, "pnpm", "run", "test:integration")
	})
}
