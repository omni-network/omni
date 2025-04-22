package cctp

import (
	"github.com/omni-network/omni/lib/evmchain"
)

var (
	// domains is a map of chain IDs to CCTP domain IDs (reference: https://developers.circle.com/stablecoins/evm-smart-contracts)
	domains = map[uint64]uint32{
		evmchain.IDEthereum:    0,
		evmchain.IDOptimism:    2,
		evmchain.IDArbitrumOne: 3,
		evmchain.IDBase:        6,

		// same for testnets
		evmchain.IDSepolia:     0,
		evmchain.IDOpSepolia:   2,
		evmchain.IDArbSepolia:  3,
		evmchain.IDBaseSepolia: 6,
	}
)
