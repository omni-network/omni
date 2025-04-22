package cctp

import (
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/netconf"
)

var (
	//  map chain IDs to CCTP domain IDs (reference: https://developers.circle.com/stablecoins/evm-smart-contracts)
	mainnetDomains = map[uint64]uint32{
		evmchain.IDEthereum:    0,
		evmchain.IDOptimism:    2,
		evmchain.IDArbitrumOne: 3,
		evmchain.IDBase:        6,
	}

	// same as mainnet.
	testnetDomains = map[uint64]uint32{
		evmchain.IDSepolia:     0,
		evmchain.IDOpSepolia:   2,
		evmchain.IDArbSepolia:  3,
		evmchain.IDBaseSepolia: 6,
	}

	// mainnet and testnet merged.
	domains = mustMergeDomains(mainnetDomains, testnetDomains)
)

func chainIDForDomain(networkID netconf.ID, domainID uint32) (uint64, bool) {
	findIn := func(ds map[uint64]uint32) (uint64, bool) {
		for chainID, d := range ds {
			if d == domainID {
				return chainID, true
			}
		}

		return 0, false
	}

	switch networkID {
	case netconf.Mainnet:
		return findIn(mainnetDomains)
	case netconf.Omega, netconf.Staging:
		return findIn(testnetDomains)
	default:
		return 0, false
	}
}

func mustMergeDomains(mainnet, testnet map[uint64]uint32) map[uint64]uint32 {
	merged := make(map[uint64]uint32)
	for k, v := range mainnet {
		merged[k] = v
	}

	for k, v := range testnet {
		if _, ok := merged[k]; ok {
			panic("duplicate chain id")
		}

		merged[k] = v
	}

	return merged
}
