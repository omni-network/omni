package app

import (
	"context"

	"github.com/omni-network/omni/lib/contracts/feeoraclev2"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/log"
)

// DeployFeeOracleV2 deploys fee oracle v2.
func DeployFeeOracleV2(ctx context.Context, def Definition) error {
	_, ok := def.Testnet.OmniEVMChain()
	if !ok {
		return errors.New("no omni evm chain")
	}

	allChains := def.Testnet.EVMChains()
	chainIDs := getChainIDs(def)
	backends := def.Backends()

	for _, chain := range allChains {
		addr, receipt, err := feeoraclev2.DeployIfNeeded(ctx, def.Testnet.Network, chain.ChainID, chainIDs, backends)
		if err != nil {
			return errors.Wrap(err, "deploy", "chain", chain.Name, "tx", maybeTxHash(receipt))
		}

		log.Info(ctx, "FeeOracleV2 deployed", "chain", chain.Name, "address", addr.Hex(), "tx", maybeTxHash(receipt))
	}

	return nil
}

func getChainIDs(def Definition) []uint64 {
	allChains := def.Testnet.EVMChains()

	chainIDs := make([]uint64, 0, len(allChains))
	for _, chain := range allChains {
		chainIDs = append(chainIDs, chain.ChainID)
	}

	return chainIDs
}

func maybeTxHash(receipt *ethclient.Receipt) string {
	if receipt != nil {
		return receipt.TxHash.Hex()
	}

	return "nil"
}
