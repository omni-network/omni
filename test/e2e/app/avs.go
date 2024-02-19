package app

import (
	"context"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/test/e2e/netman/avs"
)

func deployAVS(ctx context.Context, def Definition, cfg DeployConfig) error {
	if cfg.EigenFile == "" {
		log.Warn(ctx, "No eigen file provided, skipping AVS deployment", nil)
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

	xdapp := avs.New(
		avs.DefaultTestAVSConfig(elDeps),
		elDeps,
		portal.DeployInfo.PortalAddress,
		chain,
		portal.Client,
		portal.TxOpts(ctx, nil),
	)

	return xdapp.Deploy(ctx)
}
