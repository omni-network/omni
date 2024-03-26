package app

import (
	"context"

	"github.com/omni-network/omni/e2e/types"
	"github.com/omni-network/omni/lib/contracts"
	"github.com/omni-network/omni/lib/contracts/avs"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
)

// AVSDeploy deploys the Omni AVS contracts.
func AVSDeploy(ctx context.Context, def Definition) error {
	_, _, err := deployAVS(ctx, def)
	if err != nil {
		return err
	}

	return nil
}

// deployAVSWithExport deploys the Omni AVS contracts and exports the deployment information.
func deployAVSWithExport(ctx context.Context, def Definition, deployInfo types.DeployInfos) error {
	chain, deployment, err := deployAVS(ctx, def)
	if err != nil {
		return err
	}

	deployInfo.Set(chain.ID, types.ContractOmniAVS, deployment.Address, deployment.BlockHeight)

	return nil
}

func deployAVS(ctx context.Context, def Definition) (types.EVMChain, contracts.Deployment, error) {
	chain, err := def.Testnet.AVSChain()
	if err != nil {
		return types.EVMChain{}, contracts.Deployment{}, errors.Wrap(err, "avs chain")
	}

	backend, err := def.Backends.Backend(chain.ID)
	if err != nil {
		return types.EVMChain{}, contracts.Deployment{}, errors.Wrap(err, "backend")
	}

	deployment, err := avs.DeployIfNeeded(ctx, def.Testnet.Network, backend)
	if err != nil {
		if !deployment.IsEmpty() {
			log.Warn(ctx, "AVS deployed with error", err, "chain", chain, "addr", deployment.Address.Hex(), "block", deployment.BlockHeight)
		}

		return types.EVMChain{}, deployment, errors.Wrap(err, "deploy")
	}

	log.Info(ctx, "AVS contract deployed", "chain", chain, "addr", deployment.Address.Hex(), "block", deployment.BlockHeight)

	return chain, deployment, nil
}
