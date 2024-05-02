package app

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"

	"github.com/cometbft/cometbft/test/e2e/pkg/exec"
)

const (
	EnvInfraType       = "INFRASTRUCTURE_TYPE"
	EnvInfraFile       = "INFRASTRUCTURE_FILE"
	EnvE2EManifest     = "E2E_MANIFEST"
	EnvE2ENode         = "E2E_NODE"
	EnvE2ERPCEndpoints = "E2E_RPC_ENDPOINTS"
	EnvE2EDeployInfo   = "E2E_DEPLOY_INFO"
)

// Test runs test cases under tests/.
func Test(ctx context.Context, def Definition, verbose bool) error {
	log.Info(ctx, "Running tests in ./test/...")
	endpoints := externalEndpoints(def)

	networkDir, err := os.MkdirTemp("", "omni-e2e")
	if err != nil {
		return errors.Wrap(err, "creating temp dir")
	}

	endpointsFile := filepath.Join(networkDir, "endpoints.json")
	if endopintsBytes, err := json.Marshal(endpoints); err != nil {
		return errors.Wrap(err, "marshaling endpoints")
	} else if err := os.WriteFile(endpointsFile, endopintsBytes, 0644); err != nil {
		return errors.Wrap(err, "writing endpoints")
	} else if err = os.Setenv(EnvE2ERPCEndpoints, endpointsFile); err != nil {
		return errors.Wrap(err, "setting env ar")
	}

	manifestFile, err := filepath.Abs(def.Testnet.File)
	if err != nil {
		return errors.Wrap(err, "absolute manifest path")
	}

	if err = os.Setenv(EnvE2EManifest, manifestFile); err != nil {
		return errors.Wrap(err, "setting env var")
	}

	infd := def.Infra.GetInfrastructureData()
	if infd.Path != "" {
		infdPath, err := filepath.Abs(infd.Path)
		if err != nil {
			return errors.Wrap(err, "absolute infrastructure path")
		}
		err = os.Setenv(EnvInfraFile, infdPath)
		if err != nil {
			return errors.Wrap(err, "setting env var")
		}
	}

	if err = os.Setenv(EnvInfraType, infd.Provider); err != nil {
		return errors.Wrap(err, "setting env var")
	}

	deployInfoFile := filepath.Join(networkDir, "deployinfo.json")
	if err := def.DeployInfos().Save(deployInfoFile); err != nil {
		return errors.Wrap(err, "saving deployinfo")
	}
	if err = os.Setenv(EnvE2EDeployInfo, deployInfoFile); err != nil {
		return errors.Wrap(err, "setting E2E_DEPLOY_INFO")
	}

	log.Debug(ctx, "Env files",
		EnvE2EManifest, manifestFile,
		EnvInfraType, infd.Provider,
		EnvInfraFile, infd.Path,
		EnvE2EDeployInfo, deployInfoFile,
		EnvE2ERPCEndpoints, endpointsFile,
	)

	args := []string{"go", "test", "-timeout", "60s", "-count", "1"}
	if verbose {
		args = append(args, "-v")
	}
	args = append(args, "github.com/omni-network/omni/e2e/test")
	log.Debug(ctx, "Test command", "args", args)

	err = exec.CommandVerbose(ctx, args...)
	if err != nil {
		return errors.Wrap(err, "go tests failed")
	}

	return nil
}
