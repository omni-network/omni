package app

import (
	"context"
	"os"
	"path/filepath"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"

	e2e "github.com/cometbft/cometbft/test/e2e/pkg"
	"github.com/cometbft/cometbft/test/e2e/pkg/exec"
)

// Test runs test cases under tests/.
func Test(ctx context.Context, testnet *e2e.Testnet, infra Provider) error {
	log.Info(ctx, "Running tests in ./tests/...")

	network := infra.ExternalNetwork()

	networkDir, err := os.MkdirTemp("", "omni-e2e")
	if err != nil {
		return errors.Wrap(err, "creating temp dir")
	}
	networkFile := filepath.Join(networkDir, "network.json")
	if err := netconf.Save(network, networkFile); err != nil {
		return errors.Wrap(err, "saving network")
	}

	if err = os.Setenv("E2E_NETWORK", networkFile); err != nil {
		return errors.Wrap(err, "setting E2E_MANIFEST")
	}

	manifestFile, err := filepath.Abs(testnet.File)
	if err != nil {
		return errors.Wrap(err, "absolute manifest path")
	}

	if err = os.Setenv("E2E_MANIFEST", manifestFile); err != nil {
		return errors.Wrap(err, "setting E2E_MANIFEST")
	}

	infraData := infra.GetInfrastructureData()

	if p := infraData.Path; p != "" {
		err = os.Setenv("INFRASTRUCTURE_FILE", p)
		if err != nil {
			return errors.Wrap(err, "setting INFRASTRUCTURE_FILE")
		}
	}

	if err = os.Setenv("INFRASTRUCTURE_TYPE", infraData.Provider); err != nil {
		return errors.Wrap(err, "setting INFRASTRUCTURE_TYPE")
	}

	err = exec.CommandVerbose(ctx, "go", "test", "-count", "1", "github.com/omni-network/omni/test/e2e/tests")
	if err != nil {
		return errors.Wrap(err, "go tests failed")
	}

	return nil
}
