package rebalance

import (
	"github.com/omni-network/omni/lib/evmchain"
)

var (
	// dependents maps a chain ID to a list of chain IDs that depend on it.
	dependents = map[uint64][]uint64{
		evmchain.IDEthereum: {
			evmchain.IDMantle,   // Ethereum USDC refills Mantle USDC
			evmchain.IDHyperEVM, // Ethereum USDT refills HyperEVM USDT0
		},
	}
)

func GetDependents(chainID uint64) []uint64 {
	if deps, ok := dependents[chainID]; ok {
		return deps
	}

	return []uint64{}
}
