package solvernet

import (
	"github.com/omni-network/omni/lib/errors"

	"github.com/ethereum/go-ethereum/common"
)

const (
	None      uint8 = iota // 0
	OmniCore  uint8 = 1    // 1
	Hyperlane uint8 = 2    // 2
)

var (
	mailbox = map[uint64]common.Address{
		// Mainnet
		1:     common.HexToAddress("0xc005dc82818d67AF737725bD4bf75435d065D239"), // Ethereum
		10:    common.HexToAddress("0xd4C1905BB1D26BC93DAC913e13CaCC278CdCC80D"), // Optimism
		56:    common.HexToAddress("0x2971b9Aec44bE4eb673DF1B88cDB57b96eefe8a4"), // BSC
		137:   common.HexToAddress("0x5d934f4e2f797775e53561bB72aca21ba36B96BB"), // Polygon
		999:   common.HexToAddress("0x3a464f746D23Ab22155710f44dB16dcA53e0775E"), // HyperEVM
		5000:  common.HexToAddress("0x398633D19f4371e1DB5a8EFE90468eB70B1176AA"), // Mantle
		8453:  common.HexToAddress("0xeA87ae93Fa0019a82A727bfd3eBd1cFCa8f64f1D"), // Base
		42161: common.HexToAddress("0x979Ca5202784112f4738403dBec5D0F3B9daabB9"), // Arbitrum
		80094: common.HexToAddress("0x7f50C5776722630a0024fAE05fDe8b47571D7B39"), // Berachain
		98866: common.HexToAddress("0x3a464f746D23Ab22155710f44dB16dcA53e0775E"), // Plume

		// Testnet
		17000:    common.HexToAddress("0x46f7C5D896bbeC89bE1B19e4485e59b4Be49e9Cc"), // Ethereum Holesky
		84532:    common.HexToAddress("0x6966b0E55883d49BFB24539356a2f8A673E02039"), // Base Sepolia
		421614:   common.HexToAddress("0x598facE78a4302f11E3de0bee1894Da0b2Cb71F8"), // Arbitrum Sepolia
		11155420: common.HexToAddress("0x6966b0E55883d49BFB24539356a2f8A673E02039"), // Optimism Sepolia
	}

	provider = map[uint64]uint8{
		// Mainnet
		1:     OmniCore,  // Ethereum
		10:    OmniCore,  // Optimism
		56:    Hyperlane, // BSC
		137:   Hyperlane, // Polygon
		999:   Hyperlane, // HyperEVM
		5000:  Hyperlane, // Mantle
		8453:  OmniCore,  // Base
		42161: OmniCore,  // Arbitrum
		80094: Hyperlane, // Berachain
		98866: Hyperlane, // Plume

		// Testnet
		17000:    OmniCore, // Ethereum Holesky
		84532:    OmniCore, // Base Sepolia
		421614:   OmniCore, // Arbitrum Sepolia
		11155420: OmniCore, // Optimism Sepolia
	}
)

func HyperlaneMailbox(chainID uint64) (common.Address, error) {
	addr, ok := mailbox[chainID]
	if !ok {
		return common.Address{}, errors.New("hyperlane mailbox not found for chainID", "chain_id", chainID)
	}

	return addr, nil
}

func Provider(chainID uint64) (uint8, error) {
	provider, ok := provider[chainID]
	if !ok {
		return None, errors.New("provider not found for chainID", "chain_id", chainID)
	}

	return provider, nil
}
