package targets

import (
	"github.com/omni-network/omni/lib/evmchain"

	"github.com/ethereum/go-ethereum/common"
)

const NameGearbox = "Gearbox"

var (
	gearbox = Target{
		Name: NameGearbox,
		Addresses: func(chainID uint64) map[common.Address]bool {
			if chainID == evmchain.IDEthereum {
				return set(
					addr("0xfdBB83182078767dB0D41Aa7C5b06bA118495fC8"), // WETHDepositZapper
				)
			}

			if chainID == evmchain.IDArbitrumOne {
				return set(
					addr("0x6D6bA570Bb02f95AaF1C6e6eC1aCc2119e2E9a76"), // WETHDepositZapper
					addr("0xbe7b59A8a00b2D25bF95177dEe6A941C76A8C2F5"), // UnderlyingDepositZapper
				)
			}

			if chainID == evmchain.IDOptimism {
				return set(
					addr("0xC26952f027f6884391F8c43cAe5cca375C85E262"), // WETHDepositZapper
				)
			}

			return map[common.Address]bool{}
		},
	}
)
