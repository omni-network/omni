package app

import (
	"context"

	"github.com/omni-network/omni/e2e/types"
	"github.com/omni-network/omni/lib/contracts/avs"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"

	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

// AVSDeploy deploys the Omni AVS contracts.
func AVSDeploy(ctx context.Context, def Definition) error {
	chain, addr, receipt, err := deployAVS(ctx, def)
	if err != nil {
		return err
	}

	log.Info(ctx, "AVS contract deployed", "chain", chain.Name, "addr", addr.Hex(), "block", receipt.BlockNumber)

	return nil
}

// deployAVSWithExport deploys the Omni AVS contracts and exports the deployment information.
func deployAVSWithExport(ctx context.Context, def Definition, deployInfo types.DeployInfos) error {
	log.Info(ctx, "Deploying AVS contract")

	chain, addr, receipt, err := deployAVS(ctx, def)
	if err != nil {
		return err
	}

	deployInfo.Set(chain.ID, types.ContractOmniAVS, addr, receipt.BlockNumber.Uint64())

	log.Info(ctx, "AVS contract deployed", "chain", chain.Name, "addr", addr.Hex(), "block", receipt.BlockNumber)

	return nil
}

func deployAVS(ctx context.Context, def Definition) (types.EVMChain, common.Address, *ethtypes.Receipt, error) {
	chain, err := def.Testnet.AVSChain()
	if err != nil {
		return types.EVMChain{}, common.Address{}, nil, errors.Wrap(err, "avs chain")
	}

	backend, err := def.Backends.Backend(chain.ID)
	if err != nil {
		return types.EVMChain{}, common.Address{}, nil, errors.Wrap(err, "backend")
	}

	addr, receipt, err := avs.Deploy(ctx, def.Testnet.Network, backend)
	if err != nil {
		return types.EVMChain{}, common.Address{}, nil, errors.Wrap(err, "deploy")
	}

	return chain, addr, receipt, nil
}
