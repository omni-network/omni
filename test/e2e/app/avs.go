package app

import (
	"context"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/test/e2e/netman/avs"
	"github.com/omni-network/omni/test/e2e/types"
)

func deployAVS(ctx context.Context, def Definition, cfg DeployConfig, deployInfo types.DeployInfos) error {
	if cfg.EigenFile == "" {
		log.Warn(ctx, "No eigen deployments file provided, skipping AVS deployment", nil)
		return nil
	}

	elDeps, err := avs.LoadDeployments(cfg.EigenFile)
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

	xdapp := avs.New(
		avs.DefaultTestAVSConfig(elDeps),
		elDeps,
		portal.DeployInfo.PortalAddress,
		chain,
		omniChainID,
		portal.Client,
		portal.TxOpts(ctx, nil),
	)

	if err := xdapp.Deploy(ctx); err != nil {
		return errors.Wrap(err, "deploy xdapp")
	}

	xdapp.ExportDeployInfo(deployInfo)

	return nil
}
