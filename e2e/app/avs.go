package app

import (
	"context"

	"github.com/omni-network/omni/lib/contracts/avs"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
)

// DeployAVSAndCreate3 deploys a create3 contract and the Omni AVS contract.
func DeployAVSAndCreate3(ctx context.Context, def Definition) error {
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

	return deployAVS(ctx, def)
}

func deployAVS(ctx context.Context, def Definition) error {
	chain, err := def.Testnet.AVSChain()
	if err != nil {
		return errors.Wrap(err, "avs chain")
	}

	backend, err := def.Backends().Backend(chain.ID)
	if err != nil {
		return errors.Wrap(err, "backend")
	}

	addr, receipt, err := avs.DeployIfNeeded(ctx, def.Testnet.Network, backend)
	if err != nil {
		return errors.Wrap(err, "deploy")
	}

	if receipt == nil {
		log.Info(ctx, "AVS contract already deployed", "chain", chain.Name, "addr", addr.Hex())
	} else {
		log.Info(ctx, "AVS contract deployed", "chain", chain.Name, "addr", addr.Hex(), "block", receipt.BlockNumber)
	}

	return nil
}
