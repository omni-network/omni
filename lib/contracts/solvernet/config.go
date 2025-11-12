package solvernet

import (
	"slices"

	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/ethereum/go-ethereum/common"
)

const (
	ProviderNone    uint8 = iota // No provider
	ProviderCore                 // Omni Core
	ProviderHL                   // Hyperlane
	ProviderTrusted              // Trusted (e.g. Solana)
)

var (
	// mailbox maps chain ID to Hyperlane mailbox address.
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

	// isCore maps chain ID to whether it is a core Omni chain.
	isCore = map[uint64]bool{
		// Mainnet
		evmchain.IDEthereum:    true,
		evmchain.IDOptimism:    true,
		evmchain.IDArbitrumOne: true,
		evmchain.IDBase:        true,
		evmchain.IDOmniMainnet: true,

		// Omega / Staging
		evmchain.IDHolesky:     true,
		evmchain.IDArbSepolia:  true,
		evmchain.IDBaseSepolia: true,
		evmchain.IDOpSepolia:   true,
		evmchain.IDOmniOmega:   true,
		evmchain.IDOmniStaging: true,

		// Devnet
		evmchain.IDMockL1:     true,
		evmchain.IDMockL2:     true,
		evmchain.IDOmniDevnet: true,
	}
)

// HyperlaneMailbox returns the Hyperlane mailbox address for `chainID`.
func HyperlaneMailbox(chainID uint64) (common.Address, bool) {
	addr, ok := mailbox[chainID]
	return addr, ok
}

// IsCore returns true if the chain supports Omni Core interop.
func IsCore(chainID uint64) bool {
	return isCore[chainID]
}

// IsHL returns true if the chain supports Hyperlane interop.
func IsHL(chainID uint64) bool {
	_, ok := mailbox[chainID]
	return ok
}

// IsTrusted returns true if the chain ID is solver-trusted chain, like solana.
func IsTrusted(chainID uint64) bool {
	for _, chains := range trustedChains {
		if slices.Contains(chains, chainID) {
			return true
		}
	}

	return false
}

// IsHLOnly returns true if the chain only supports Hyperlane interop.
func IsHLOnly(chainID uint64) bool {
	return IsHL(chainID) && !IsCore(chainID) && !IsTrusted(chainID)
}

// IsSupported returns true if the chain ID is supported.
func IsSupported(chainID uint64) bool {
	return IsCore(chainID) || IsHL(chainID) || IsTrusted(chainID)
}

func SkipRole(chainID uint64, role eoa.Role) bool {
	// Only skip the role if the chain is a solvernet chain.
	if !IsSolverOnly(chainID) {
		return false
	}

	// Skip non-HL roles on HL-only chains.
	if IsHLOnly(chainID) && !isHLRole(role) {
		return true
	}

	if IsTrusted(chainID) && !isTrustedRole(role) {
		return true // Skip non-trusted roles on trusted chains.
	}

	return false
}

// OnlyCoreEndpoints filters the given RPC endpoints to only include core endpoints.
// Necessary prereq for netconf.AwaitOnExecutionChain, which expects all
// endopints to have portal registrations.
func OnlyCoreEndpoints(endpoints xchain.RPCEndpoints) xchain.RPCEndpoints {
	out := make(xchain.RPCEndpoints)

	for name, rpc := range endpoints {
		meta, ok := evmchain.MetadataByName(name)
		if !ok {
			continue
		}

		if IsCore(meta.ChainID) {
			out[name] = rpc
		}
	}

	return out
}

// Provider returns the provider between a source and destination chain.
// It returns false if the route is not supported.
func Provider(srcChainID, destChainID uint64) (uint8, bool) {
	if !IsSupported(srcChainID) || !IsSupported(destChainID) {
		// Unsupported chain: not supported
		return ProviderNone, false
	}

	if IsTrusted(srcChainID) || IsTrusted(destChainID) {
		return ProviderTrusted, true
	}

	if srcChainID == destChainID {
		// Same chain: no provider needed
		return ProviderNone, true
	}

	if IsHLOnly(srcChainID) && !IsHL(destChainID) {
		// HL-only -> No-HL: not supported
		return ProviderNone, false
	}

	if !IsHL(srcChainID) && IsHLOnly(destChainID) {
		// No-HL -> HL-only: not supported
		return ProviderNone, false
	}

	if IsCore(srcChainID) && IsCore(destChainID) {
		// Core -> Core: use Core provider
		return ProviderCore, true
	}

	// Default: use HL provider
	return ProviderHL, true
}
