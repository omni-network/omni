package app

import (
	"context"
	"encoding/json"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/test/e2e/netman/avs"
	"github.com/omni-network/omni/test/e2e/types"

	_ "embed"
)

//go:embed static/el-deployments.json
var defaultEigenDeps []byte

func deployAVS(ctx context.Context, def Definition, cfg DeployConfig, deployInfo types.DeployInfos) error {
	var (
		elDeps avs.EigenDeployments
		err    error
	)
	if cfg.EigenFile != "" {
		elDeps, err = avs.LoadDeployments(cfg.EigenFile)
	} else {
		log.Debug(ctx, "Using default eigen deployments")
		err = json.Unmarshal(defaultEigenDeps, &elDeps)
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

	xdapp := avs.New(
		avs.DefaultTestAVSConfig(elDeps),
		elDeps,
		portal.DeployInfo.PortalAddress,
		chain,
		omniChainID,
		def.Backends,
	)

	if err := xdapp.Deploy(ctx); err != nil {
		return errors.Wrap(err, "deploy xdapp")
	}

	xdapp.ExportDeployInfo(deployInfo)

	return nil
}
