// Package evmchain provides static metadata about supported evm chains.
//
// This package should only contain public well-known metadata and
// should not be Omni-network specific, since multiple omni networks
// can be deployed to the same evm chain.
package evmchain

import (
	"fmt"
	"time"

	"github.com/omni-network/omni/lib/tokens"
)

// Source chain IDs as per https://chainlist.org/

const (
	// Mainnets.
	IDEthereum    uint64 = 1
	IDOptimism    uint64 = 10
	IDBSC         uint64 = 56
	IDPolygon     uint64 = 137
	IDOmniMainnet uint64 = 166
	IDSolana      uint64 = 350
	IDHyperEVM    uint64 = 999
	IDMantle      uint64 = 5000
	IDBase        uint64 = 8453
	IDArbitrumOne uint64 = 42161
	IDBerachain   uint64 = 80094
	IDPlume       uint64 = 98866

	// Testnets.
	IDBSCTestnet uint64 = 97
	IDOmniOmega  uint64 = 164

	IDHyperEVMTestnet uint64 = 998
	IDHolesky         uint64 = 17000
	IDPolygonAmoy     uint64 = 80002
	IDBaseSepolia     uint64 = 84532
	IDPlumeTestnet    uint64 = 98867
	IDArbSepolia      uint64 = 421614
	IDSepolia         uint64 = 11155111
	IDOpSepolia       uint64 = 11155420
	IDSolanaDevnet    uint64 = 351 // Note we use solana devnet, not testnet which is for performance testing only.

	// Ephemeral.
	IDOmniStaging uint64 = 1650
	IDOmniDevnet  uint64 = 1651
	IDMockL1      uint64 = 1652
	IDMockL2      uint64 = 1654
	IDMockOp      uint64 = 1655
	IDMockArb     uint64 = 1656
	IDSolanaLocal uint64 = 352

	omniEVMName        = "omni_evm"
	omniEVMBlockPeriod = time.Second * 2
)

type Metadata struct {
	ChainID     uint64
	PostsTo     uint64 // chain id to which tx data is posted
	Name        string
	PrettyName  string
	BlockPeriod time.Duration
	NativeToken tokens.Asset
	Reorgs      bool // Only if chain actually reorgs, e.g. L2s don't
}

func MetadataByID(chainID uint64) (Metadata, bool) {
	resp, ok := static[chainID]
	return resp, ok
}

func MetadataByName(name string) (Metadata, bool) {
	for _, metadata := range static {
		if metadata.Name == name {
			return metadata, true
		}
	}

	return Metadata{}, false
}

var static = map[uint64]Metadata{
	// Mainnets.
	IDEthereum: {
		ChainID:     IDEthereum,
		Name:        "ethereum",
		PrettyName:  "Ethereum",
		BlockPeriod: 12 * time.Second,
		NativeToken: tokens.ETH,
		Reorgs:      true,
	},
	IDOptimism: {
		ChainID:     IDOptimism,
		Name:        "optimism",
		PrettyName:  "Optimism",
		BlockPeriod: 2 * time.Second,
		NativeToken: tokens.ETH,
		PostsTo:     IDEthereum,
	},
	IDBSC: {
		ChainID:     IDBSC,
		Name:        "bsc",
		PrettyName:  "Binance Smart Chain",
		BlockPeriod: 3 * time.Second,
		NativeToken: tokens.BNB,
		Reorgs:      true,
	},
	IDPolygon: {
		ChainID:     IDPolygon,
		Name:        "polygon",
		PrettyName:  "Polygon",
		BlockPeriod: 2 * time.Second,
		NativeToken: tokens.POL,
		Reorgs:      true,
	},
	IDOmniMainnet: {
		ChainID:     IDOmniMainnet,
		Name:        omniEVMName,
		PrettyName:  "Omni Mainnet",
		BlockPeriod: omniEVMBlockPeriod,
		NativeToken: tokens.OMNI,
	},
	IDSolana: {
		ChainID:     IDSolana,
		Name:        "solana",
		PrettyName:  "Solana Mainnet Beta",
		BlockPeriod: 400 * time.Millisecond,
		NativeToken: tokens.SOL,
	},
	IDHyperEVM: {
		ChainID:     IDHyperEVM,
		Name:        "hyper_evm",
		PrettyName:  "HyperEVM",
		BlockPeriod: 2 * time.Second,
		NativeToken: tokens.HYPE,
		Reorgs:      true, // Hyperlane docs don't mention reorgs at all, neither do explorers
	},
	IDMantle: {
		ChainID:     IDMantle,
		Name:        "mantle",
		PrettyName:  "Mantle",
		BlockPeriod: 2 * time.Second,
		NativeToken: tokens.MNT,
		// PostsTo: EigenDA
	},
	IDBase: {
		ChainID:     IDBase,
		Name:        "base",
		PrettyName:  "Base",
		BlockPeriod: 2 * time.Second,
		NativeToken: tokens.ETH,
		PostsTo:     IDEthereum,
	},
	IDArbitrumOne: {
		ChainID:     IDArbitrumOne,
		Name:        "arbitrum_one",
		PrettyName:  "Arbitrum One",
		BlockPeriod: 300 * time.Millisecond,
		NativeToken: tokens.ETH,
		PostsTo:     IDEthereum,
	},
	IDBerachain: {
		ChainID:     IDBerachain,
		Name:        "berachain",
		PrettyName:  "Berachain",
		BlockPeriod: 2 * time.Second,
		NativeToken: tokens.BERA,
		Reorgs:      true, // Berachain docs don't mention reorgs at all, neither do explorers
	},
	IDPlume: {
		ChainID:     IDPlume,
		Name:        "plume",
		PrettyName:  "Plume",
		BlockPeriod: 1 * time.Second,
		NativeToken: tokens.PLUME,
	},

	// Testnets.
	IDBSCTestnet: {
		ChainID:     IDBSCTestnet,
		Name:        "bsc_testnet",
		PrettyName:  "BSC Testnet",
		BlockPeriod: 2 * time.Millisecond,
		NativeToken: tokens.BNB,
		Reorgs:      true,
	},
	IDOmniOmega: {
		ChainID:     IDOmniOmega,
		Name:        omniEVMName,
		PrettyName:  "Omni Omega",
		BlockPeriod: omniEVMBlockPeriod,
		NativeToken: tokens.NOM,
	},
	IDHyperEVMTestnet: {
		ChainID:     IDHyperEVMTestnet,
		Name:        "hyper_evm_testnet",
		PrettyName:  "HyperEVM Testnet",
		BlockPeriod: 2 * time.Second,
		NativeToken: tokens.HYPE,
		Reorgs:      true, // Hyperlane docs don't mention reorgs at all, neither do explorers
	},
	IDHolesky: {
		ChainID:     IDHolesky,
		Name:        "holesky",
		PrettyName:  "Holesky",
		BlockPeriod: 12 * time.Second,
		NativeToken: tokens.ETH,
		Reorgs:      true,
	},
	IDPolygonAmoy: {
		ChainID:     IDPolygonAmoy,
		Name:        "polygon_amoy",
		PrettyName:  "Polygon Amoy",
		BlockPeriod: 2 * time.Second,
		NativeToken: tokens.POL,
		Reorgs:      true,
	},
	IDBaseSepolia: {
		ChainID:     IDBaseSepolia,
		Name:        "base_sepolia",
		PrettyName:  "Base Sepolia",
		BlockPeriod: 2 * time.Second,
		NativeToken: tokens.ETH,
		PostsTo:     IDSepolia,
	},
	IDPlumeTestnet: {
		ChainID:     IDPlumeTestnet,
		Name:        "plume_testnet",
		PrettyName:  "Plume Testnet",
		BlockPeriod: 1 * time.Second,
		NativeToken: tokens.PLUME,
	},
	IDArbSepolia: {
		ChainID:     IDArbSepolia,
		Name:        "arb_sepolia",
		PrettyName:  "Arb Sepolia",
		BlockPeriod: 300 * time.Millisecond,
		NativeToken: tokens.ETH,
		PostsTo:     IDSepolia,
	},
	IDSepolia: {
		ChainID:     IDSepolia,
		Name:        "sepolia",
		PrettyName:  "Sepolia",
		BlockPeriod: 12 * time.Second,
		NativeToken: tokens.ETH,
		Reorgs:      true,
	},
	IDOpSepolia: {
		ChainID:     IDOpSepolia,
		Name:        "op_sepolia",
		PrettyName:  "OP Sepolia",
		BlockPeriod: 2 * time.Second,
		NativeToken: tokens.ETH,
		PostsTo:     IDSepolia,
	},
	IDSolanaDevnet: {
		ChainID:     IDSolanaDevnet,
		Name:        "solana_testnet",
		PrettyName:  "Solana Testnet",
		BlockPeriod: 400 * time.Millisecond,
		NativeToken: tokens.SOL,
	},

	// Ephemeral.
	IDOmniStaging: {
		ChainID:     IDOmniStaging,
		Name:        omniEVMName,
		PrettyName:  "Omni Staging",
		BlockPeriod: omniEVMBlockPeriod,
		NativeToken: tokens.NOM,
	},
	IDOmniDevnet: {
		ChainID:     IDOmniDevnet,
		Name:        omniEVMName,
		PrettyName:  "Omni Devnet",
		BlockPeriod: omniEVMBlockPeriod,
		NativeToken: tokens.NOM,
	},
	IDMockL1: {
		ChainID:     IDMockL1,
		Name:        "mock_l1",
		PrettyName:  "Mock L1",
		BlockPeriod: time.Second,
		NativeToken: tokens.ETH,
		Reorgs:      true,
	},
	IDMockL2: {
		ChainID:     IDMockL2,
		Name:        "mock_l2",
		PrettyName:  "Mock L2",
		BlockPeriod: time.Second,
		NativeToken: tokens.ETH,
	},
	IDMockOp: {
		ChainID:     IDMockOp,
		Name:        "mock_op",
		PrettyName:  "Mock OP",
		BlockPeriod: time.Second * 2,
		NativeToken: tokens.ETH,
	},
	IDMockArb: {
		ChainID:     IDMockArb,
		Name:        "mock_arb",
		PrettyName:  "Mock ARB",
		BlockPeriod: time.Second,
		NativeToken: tokens.ETH,
	},
	IDSolanaLocal: {
		ChainID:     IDSolanaLocal,
		Name:        "solana_localnet",
		PrettyName:  "Solana Localnet",
		BlockPeriod: 400 * time.Millisecond,
		NativeToken: tokens.SOL,
	},
}

// IsDisabled returns true if the chain is disabled.
// TODO(corver): Remove once holesky issue resolved.
func IsDisabled(_ uint64) bool {
	return false // id == IDHolesky
}

// Name returns the name of the chain by its ID.
func Name(id uint64) string {
	metadata, ok := MetadataByID(id)
	if !ok {
		return fmt.Sprintf("unknown(%d)", id)
	}

	return metadata.Name
}

// IsSVM returns true if the chain is a Solana Virtual Machine (SVM) chain.
func IsSVM(chainID uint64) bool {
	return chainID == IDSolana || chainID == IDSolanaDevnet || chainID == IDSolanaLocal
}
