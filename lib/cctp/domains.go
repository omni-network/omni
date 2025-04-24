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
)

func domainIDForChain(networkID netconf.ID, chainID uint64) (uint32, bool) {
	switch networkID {
	case netconf.Mainnet:
		d, ok := mainnetDomains[chainID]
		return d, ok
	case netconf.Omega, netconf.Staging:
		d, ok := testnetDomains[chainID]
		return d, ok
	default:
		return 0, false
	}
}

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
