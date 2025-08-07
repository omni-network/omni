package app

import (
	"context"

	"github.com/omni-network/omni/lib/contracts/createx"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/log"

	"github.com/ethereum/go-ethereum/common"
)

func deployAllCreateX(ctx context.Context, def Definition) error {
	for _, chain := range def.Testnet.EVMChains() {
		_, _, err := deployCreateX(ctx, def, chain.ChainID)
		if err != nil {
			return errors.Wrap(err, "deploy createx", "chain", chain.Name)
		}
	}

	return nil
}

func deployCreateX(ctx context.Context, def Definition, chainID uint64) (common.Address, *ethclient.Receipt, error) {
	backend, err := def.Backends().Backend(chainID)
	if err != nil {
		return common.Address{}, nil, err
	}

	log.Debug(ctx, "Deploying CreateX", "network", def.Testnet.Network, "chain", chainID)

	return createx.DeployIfNeeded(ctx, def.Testnet.Network, backend)
}
