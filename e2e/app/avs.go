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
	chain, err := def.Testnet.AVSChain()
	if err != nil {
		return errors.Wrap(err, "avs chain")
	}

	factory, receipt, err := deployCreate3(ctx, def, chain.ID)
	if err != nil {
		return errors.Wrap(err, "deploy create3")
	} else if receipt != nil {
		log.Info(ctx, "Deployed create3 factory", "chain", chain.Name, "addr", factory.Hex(), "block", receipt.BlockNumber)
	}

	log.Info(ctx, "Deploying via create3 factory", "chain", chain.Name, "addr", factory.Hex())

	avs, receipt, err := deployAVS(ctx, def, chain.ID)
	if err != nil {
		return err
	}

	logAVSDeployment(ctx, chain.Name, avs, receipt)

	return nil
}

// deployAVSWithExport deploys the Omni AVS contracts and exports the deployment information.
func deployAVSWithExport(ctx context.Context, def Definition, deployInfo types.DeployInfos) error {
	chain, err := def.Testnet.AVSChain()
	if err != nil {
		return errors.Wrap(err, "avs chain")
	}

	addr, receipt, err := deployAVS(ctx, def, chain.ID)
	if err != nil {
		return err
	}

	// If receipt is nil, the avs has already been deployed, in which case we
	// don't know the deploy height. This is fine for the AVS contract - we only
	// need deploy heights for Portal contracts.
	//
	// It may be worth refactoring DeployInfos to allow for explicitly nil deploy heights.
	// Or, do not track AVS in e2e run deploy info - it does not look like it is used.

	blockNumber := uint64(0)
	if receipt != nil {
		blockNumber = receipt.BlockNumber.Uint64()
	}

	deployInfo.Set(chain.ID, types.ContractOmniAVS, addr, blockNumber)

	logAVSDeployment(ctx, chain.Name, addr, receipt)

	return nil
}

func deployAVS(ctx context.Context, def Definition, chainID uint64) (common.Address, *ethtypes.Receipt, error) {
	backend, err := def.Backends().Backend(chainID)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "backend")
	}

	addr, receipt, err := avs.DeployIfNeeded(ctx, def.Testnet.Network, backend)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "deploy")
	}

	return addr, receipt, nil
}

func logAVSDeployment(ctx context.Context, chain string, addr common.Address, receipt *ethtypes.Receipt) {
	if receipt == nil {
		log.Info(ctx, "AVS contract already deployed", "chain", chain, "addr", addr.Hex())
	} else {
		log.Info(ctx, "AVS contract deployed", "chain", chain, "addr", addr.Hex(), "block", receipt.BlockNumber)
	}
}
