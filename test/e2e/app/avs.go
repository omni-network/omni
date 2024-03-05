package app

import (
	"context"
	"encoding/json"

	"github.com/omni-network/omni/lib/avs"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/test/e2e/app/static"
	"github.com/omni-network/omni/test/e2e/types"

	_ "embed"
)

func deployAVS(ctx context.Context, def Definition, cfg DeployConfig, deployInfo types.DeployInfos) error {
	var (
		elDeps avs.EigenDeployments
		err    error
	)
	if cfg.EigenFile != "" {
		elDeps, err = avs.LoadDeployments(cfg.EigenFile)
	} else {
		log.Debug(ctx, "Using default eigen deployments")
		err = json.Unmarshal(static.GetElDeployments(), &elDeps)
	}
	if err != nil {
		return errors.Wrap(err, "load eigen deployments")
	}

	chain, err := def.Testnet.AVSChain()
	if err != nil {
		return err
	}

	portal := def.Netman.Portals()[chain.ID]

	// just use the first OmniEVM chain for now
	omniChainID := def.Testnet.OmniEVMs[0].Chain.ID

	avsDeploy := avs.NewDeployer(
		avs.DefaultTestAVSConfig(elDeps),
		elDeps,
		portal.DeployInfo.PortalAddress,
		omniChainID,
	)

	deployer, _, backend, err := def.Backends.BindOpts(ctx, chain.ID)
	if err != nil {
		return errors.Wrap(err, "bind opts")
	}

	// Use the deployer key as owner of avs
	if err := avsDeploy.Deploy(ctx, backend, deployer); err != nil {
		return errors.Wrap(err, "deploy avsDeploy")
	}

	avsDeploy.ExportDeployInfo(deployInfo)

	return nil
}
