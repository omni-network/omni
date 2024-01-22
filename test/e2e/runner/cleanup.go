package main

import (
	"context"
	"fmt"
	"os"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"

	e2e "github.com/cometbft/cometbft/test/e2e/pkg"
	"github.com/cometbft/cometbft/test/e2e/pkg/exec"
)

// Cleanup removes the Docker Compose containers and testnet directory.
func Cleanup(ctx context.Context, testnet *e2e.Testnet) error {
	err := cleanupDocker(ctx)
	if err != nil {
		return err
	}
	err = cleanupDir(ctx, testnet.Dir)
	if err != nil {
		return err
	}

	return nil
}

// cleanupDocker removes all E2E resources (with label e2e=True), regardless
// of testnet.
func cleanupDocker(ctx context.Context) error {
	log.Info(ctx, "Removing Docker containers and networks")

	// GNU xargs requires the -r flag to not run when input is empty, macOS
	// does this by default. Ugly, but works.
	xargsR := `$(if [[ $OSTYPE == "linux-gnu"* ]]; then echo -n "-r"; fi)`

	err := exec.Command(ctx, "bash", "-c", fmt.Sprintf(
		"docker container ls -qa --filter label=e2e | xargs %v docker container rm -f", xargsR))
	if err != nil {
		return errors.Wrap(err, "remove docker containers")
	}

	err = exec.Command(ctx, "bash", "-c", fmt.Sprintf(
		"docker network ls -q --filter label=e2e | xargs %v docker network rm", xargsR))
	if err != nil {
		return errors.Wrap(err, "remove docker networks")
	}

	return nil
}

// cleanupDir cleans up a testnet directory.
func cleanupDir(ctx context.Context, dir string) error {
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

	err = os.RemoveAll(dir)
	if err != nil {
		return errors.Wrap(err, "remove dir")
	}

	return nil
}
