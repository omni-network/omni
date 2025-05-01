package solvernet

import (
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/evmchain"

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
		evmchain.IDEthereum:    common.HexToAddress("0xc005dc82818d67AF737725bD4bf75435d065D239"),
		evmchain.IDOptimism:    common.HexToAddress("0xd4C1905BB1D26BC93DAC913e13CaCC278CdCC80D"),
		evmchain.IDBSC:         common.HexToAddress("0x2971b9Aec44bE4eb673DF1B88cDB57b96eefe8a4"),
		evmchain.IDPolygon:     common.HexToAddress("0x5d934f4e2f797775e53561bB72aca21ba36B96BB"),
		evmchain.IDHyperEVM:    common.HexToAddress("0x3a464f746D23Ab22155710f44dB16dcA53e0775E"),
		evmchain.IDMantle:      common.HexToAddress("0x398633D19f4371e1DB5a8EFE90468eB70B1176AA"),
		evmchain.IDBase:        common.HexToAddress("0xeA87ae93Fa0019a82A727bfd3eBd1cFCa8f64f1D"),
		evmchain.IDArbitrumOne: common.HexToAddress("0x979Ca5202784112f4738403dBec5D0F3B9daabB9"),
		evmchain.IDBerachain:   common.HexToAddress("0x7f50C5776722630a0024fAE05fDe8b47571D7B39"),
		evmchain.IDPlume:       common.HexToAddress("0x3a464f746D23Ab22155710f44dB16dcA53e0775E"),
		// Hyperlane doesn't support Omni EVM

		// Testnet
		evmchain.IDBSCTestnet:      common.HexToAddress("0xF9F6F5646F478d5ab4e20B0F910C92F1CCC9Cc6D"),
		evmchain.IDHyperEVMTestnet: common.HexToAddress("0x589C201a07c26b4725A4A829d772f24423da480B"),
		evmchain.IDHolesky:         common.HexToAddress("0x46f7C5D896bbeC89bE1B19e4485e59b4Be49e9Cc"),
		evmchain.IDPolygonAmoy:     common.HexToAddress("0x54148470292C24345fb828B003461a9444414517"),
		evmchain.IDBaseSepolia:     common.HexToAddress("0x6966b0E55883d49BFB24539356a2f8A673E02039"),
		evmchain.IDPlumeTestnet:    common.HexToAddress("0xDDcFEcF17586D08A5740B7D91735fcCE3dfe3eeD"),
		evmchain.IDArbSepolia:      common.HexToAddress("0x598facE78a4302f11E3de0bee1894Da0b2Cb71F8"),
		evmchain.IDSepolia:         common.HexToAddress("0xfFAEF09B3cd11D9b20d1a19bECca54EEC2884766"),
		evmchain.IDOpSepolia:       common.HexToAddress("0x6966b0E55883d49BFB24539356a2f8A673E02039"),
		// Hyperlane doesn't support Omni EVM or Mantle Testnet
	}

	provider = map[uint64]uint8{
		// Mainnet
		evmchain.IDEthereum:    OmniCore,
		evmchain.IDOptimism:    OmniCore,
		evmchain.IDBSC:         Hyperlane,
		evmchain.IDPolygon:     Hyperlane,
		evmchain.IDOmniMainnet: OmniCore,
		evmchain.IDHyperEVM:    Hyperlane,
		evmchain.IDMantle:      Hyperlane,
		evmchain.IDBase:        OmniCore,
		evmchain.IDArbitrumOne: OmniCore,
		evmchain.IDBerachain:   Hyperlane,
		evmchain.IDPlume:       Hyperlane,

		// Testnet
		evmchain.IDBSCTestnet:      Hyperlane,
		evmchain.IDOmniOmega:       OmniCore,
		evmchain.IDHyperEVMTestnet: Hyperlane,
		evmchain.IDHolesky:         OmniCore,
		evmchain.IDPolygonAmoy:     Hyperlane,
		evmchain.IDBaseSepolia:     OmniCore,
		evmchain.IDPlumeTestnet:    Hyperlane,
		evmchain.IDArbSepolia:      OmniCore,
		evmchain.IDSepolia:         Hyperlane,
		evmchain.IDOpSepolia:       OmniCore,

		// Devnet
		evmchain.IDOmniStaging: OmniCore,
		evmchain.IDOmniDevnet:  OmniCore,
		evmchain.IDMockL1:      OmniCore,
		evmchain.IDMockL2:      OmniCore,
		evmchain.IDMockOp:      OmniCore,
		evmchain.IDMockArb:     OmniCore,
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
