package app

import (
	"context"
	"encoding/json"

	"github.com/omni-network/omni/lib/avs"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/test/e2e/app/static"
	"github.com/omni-network/omni/test/e2e/types"
)

func DefaultAVSDeployConfig() AVSDeployConfig {
	return AVSDeployConfig{}
}

type AVSDeployConfig struct {
	EigenFile string
}

// AVSDeploy deploys the Omni AVS contracts.
func AVSDeploy(ctx context.Context, def Definition, cfg AVSDeployConfig) error {
	info := make(types.DeployInfos)
	err := deployAVS(ctx, def, cfg, info)
	if err != nil {
		return err
	}

	chain, err := def.Testnet.AVSChain()
	if err != nil {
		return err
	}

	addr, found := info.Addr(chain.ID, types.ContractOmniAVS)
	if !found {
		return errors.New("avs contract not found")
	}

	log.Info(ctx, "AVS contract deployed", "addr", addr.Hex())

	return nil
}

func deployAVS(ctx context.Context, def Definition, cfg AVSDeployConfig, deployInfo types.DeployInfos) error {
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
		log.Warn(ctx, "Not deploying AVS Contract", err)
		return nil
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
