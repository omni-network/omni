package usdt0

import (
	"github.com/omni-network/omni/lib/evmchain"
)

var (
	// eidByChain maps chain ID to LayerZero's Endpoint ID (EID).
	eidByChain = map[uint64]uint32{
		evmchain.IDEthereum:    30101,
		evmchain.IDArbitrumOne: 30110,
		evmchain.IDOptimism:    30111,
		evmchain.IDHyperEVM:    30367,
	}
)
