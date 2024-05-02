package app

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"

	"github.com/cometbft/cometbft/test/e2e/pkg/infra/docker"
)

// CleanInfra stops and removes the infra containers.
func CleanInfra(ctx context.Context, def Definition) error {
	if err := def.Infra.Clean(ctx); err != nil {
		return errors.Wrap(err, "cleaning infrastructure")
	}

	return nil
}

// CleanupDir cleans up a testnet directory.
func CleanupDir(ctx context.Context, dir string) error {
	if dir == "" {
		return errors.New("no directory set")
	}

	_, err := os.Stat(dir)
	if os.IsNotExist(err) {
		return nil
	} else if err != nil {
		return errors.Wrap(err, "stat")
	}

	log.Info(ctx, "Cleanup dir", "dir", dir)

	// On Linux, some local files in the volume will be owned by root since CometBFT
	// runs as root inside the container, so we need to clean them up from within a
	// container running as root too.
	if runtime.GOOS == "linux" {
		absDir, err := filepath.Abs(dir)
		if err != nil {
			return errors.Wrap(err, "abs dir")
		}
		err = docker.Exec(ctx, "run",
			"--rm",             // Remove the container after it exits
			"--entrypoint", "", // Clear the entrypoint so we can run a shell command
			"-v", fmt.Sprintf("%v:/mount", absDir), // Mount the testnet dir into the container
			"ethereum/client-go:latest",    // Use the latest geth image (which runs as root)
			"sh", "-c", "rm -rf /mount/*/") // Remove all files in the mounted testnet dir
		if err != nil {
			return errors.Wrap(err, "exec rm dir")
		}
	}

	err = os.RemoveAll(dir)
	if err != nil {
		return errors.Wrap(err, "remove dir")
	}

	return nil
}
