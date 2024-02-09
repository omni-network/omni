package app

import (
	"context"
	"os"
	"path/filepath"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/cometbft/cometbft/test/e2e/pkg/exec"
)

// Test runs test cases under tests/.
func Test(ctx context.Context, def Definition, verbose bool) error {
	log.Info(ctx, "Running tests in ./tests/...")

	extNetwork := externalNetwork(def.Testnet, def.Netman.DeployInfo())

	networkDir, err := os.MkdirTemp("", "omni-e2e")
	if err != nil {
		return errors.Wrap(err, "creating temp dir")
	}
	networkFile := filepath.Join(networkDir, "network.json")
	if err := netconf.Save(extNetwork, networkFile); err != nil {
		return errors.Wrap(err, "saving network")
	}

	if err = os.Setenv("E2E_NETWORK", networkFile); err != nil {
		return errors.Wrap(err, "setting E2E_MANIFEST")
	}

	manifestFile, err := filepath.Abs(def.Testnet.File)
	if err != nil {
		return errors.Wrap(err, "absolute manifest path")
	}

	if err = os.Setenv("E2E_MANIFEST", manifestFile); err != nil {
		return errors.Wrap(err, "setting E2E_MANIFEST")
	}

	infd := def.Infra.GetInfrastructureData()
	if infd.Path != "" {
		infdPath, err := filepath.Abs(infd.Path)
		if err != nil {
			return errors.Wrap(err, "absolute infrastructure path")
		}
		err = os.Setenv("INFRASTRUCTURE_FILE", infdPath)
		if err != nil {
			return errors.Wrap(err, "setting INFRASTRUCTURE_FILE")
		}
	}

	if err = os.Setenv("INFRASTRUCTURE_TYPE", infd.Provider); err != nil {
		return errors.Wrap(err, "setting INFRASTRUCTURE_TYPE")
	}

	args := []string{"go", "test", "-timeout", "15s", "-count", "1"}
	if verbose {
		args = append(args, "-v")
	}
	args = append(args, "github.com/omni-network/omni/test/e2e/tests")

	err = exec.CommandVerbose(ctx, args...)
	if err != nil {
		return errors.Wrap(err, "go tests failed")
	}

	return nil
}
