package app

import (
	"context"

	"github.com/omni-network/omni/lib/contracts/feeoraclev2"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
)

func DeployFeeOracleV2(ctx context.Context, def Definition) error {
	if err := deployFeeOracleV2(ctx, def); err != nil {
		return errors.Wrap(err, "deploy fee oracle v2")
	}

	return nil
}

func deployFeeOracleV2(ctx context.Context, def Definition) error {
	_, ok := def.Testnet.OmniEVMChain()
	if !ok {
		return errors.New("no omni evm chain")
	}

	allChains := def.Testnet.EVMChains()
	for _, chain := range allChains {
		backends := def.Backends()

		destChainIDs := make([]uint64, 0, len(allChains)-1)
		for _, destChain := range allChains {
			if destChain.ChainID != chain.ChainID {
				destChainIDs = append(destChainIDs, destChain.ChainID)
			}
		}

		addr, receipt, err := feeoraclev2.DeployIfNeeded(ctx, def.Testnet.Network, chain.ChainID, destChainIDs, backends)
		if err != nil {
			return errors.Wrap(err, "deploy", "chain", chain.Name, "tx", maybeTxHash(receipt))
		}

		log.Info(ctx, "FeeOracleV2 deployed", "chain", chain.Name, "address", addr.Hex(), "tx", maybeTxHash(receipt))
	}

	return nil
}
