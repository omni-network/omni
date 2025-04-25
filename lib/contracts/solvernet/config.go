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
		// Hyperlane doesn't support Omni EVM

		// Testnet
		97:       common.HexToAddress("0xF9F6F5646F478d5ab4e20B0F910C92F1CCC9Cc6D"), // BSC Testnet
		998:      common.HexToAddress("0x589C201a07c26b4725A4A829d772f24423da480B"), // HyperEVM Testnet
		17000:    common.HexToAddress("0x46f7C5D896bbeC89bE1B19e4485e59b4Be49e9Cc"), // Ethereum Holesky
		80002:    common.HexToAddress("0x54148470292C24345fb828B003461a9444414517"), // Polygon Amoy
		80084:    common.HexToAddress("0xDDcFEcF17586D08A5740B7D91735fcCE3dfe3eeD"), // Berachain bArtio
		84532:    common.HexToAddress("0x6966b0E55883d49BFB24539356a2f8A673E02039"), // Base Sepolia
		98867:    common.HexToAddress("0xDDcFEcF17586D08A5740B7D91735fcCE3dfe3eeD"), // Plume Testnet
		421614:   common.HexToAddress("0x598facE78a4302f11E3de0bee1894Da0b2Cb71F8"), // Arbitrum Sepolia
		11155111: common.HexToAddress("0xfFAEF09B3cd11D9b20d1a19bECca54EEC2884766"), // Ethereum Sepolia
		11155420: common.HexToAddress("0x6966b0E55883d49BFB24539356a2f8A673E02039"), // Optimism Sepolia
		// Hyperlane doesn't support Omni EVM or Mantle Testnet
	}

	provider = map[uint64]uint8{
		// Mainnet
		1:     OmniCore,  // Ethereum
		10:    OmniCore,  // Optimism
		56:    Hyperlane, // BSC
		137:   Hyperlane, // Polygon
		166:   OmniCore,  // Omni
		999:   Hyperlane, // HyperEVM
		5000:  Hyperlane, // Mantle
		8453:  OmniCore,  // Base
		42161: OmniCore,  // Arbitrum
		80094: Hyperlane, // Berachain
		98866: Hyperlane, // Plume

		// Testnet
		97:       Hyperlane, // BSC Testnet
		164:      OmniCore,  // Omni Omega
		998:      Hyperlane, // HyperEVM Testnet
		17000:    OmniCore,  // Ethereum Holesky
		80002:    Hyperlane, // Polygon Amoy
		80084:    Hyperlane, // Berachain bArtio
		84532:    OmniCore,  // Base Sepolia
		98867:    Hyperlane, // Plume Testnet
		421614:   OmniCore,  // Arbitrum Sepolia
		11155111: Hyperlane, // Ethereum Sepolia
		11155420: OmniCore,  // Optimism Sepolia

		// Devnet
		1650: OmniCore, // Omni Staging
		1651: OmniCore, // Omni Devnet
		1652: OmniCore, // Mock L1
		1654: OmniCore, // Mock L2
		1655: OmniCore, // Mock Op
		1656: OmniCore, // Mock Arb
	}
)

func HyperlaneMailbox(chainID uint64) (common.Address, bool) {
	addr, ok := mailbox[chainID]
	if !ok {
		return common.Address{}, false
	}

	return addr, true
}

func Provider(chainID uint64) (uint8, error) {
	provider, ok := provider[chainID]
	if !ok {
		return None, errors.New("provider not found for chainID", "chain_id", chainID)
	}

	return provider, nil
}
