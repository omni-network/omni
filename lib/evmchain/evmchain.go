// Package evmchain provides static metadata about supported evm chains.
//
// This package should only contain public well-known metadata and
// should not be Omni-network specific, since multiple omni networks
// can be deployed to the same evm chain.
package evmchain

import (
	"fmt"
	"sort"
	"time"

	"github.com/omni-network/omni/lib/tokens"
)

// Source chain IDs as per https://chainlist.org/

const (
	// Mainnets.
	IDEthereum    uint64 = 1
	IDOmniMainnet uint64 = 166
	IDArbitrumOne uint64 = 42161
	IDBase        uint64 = 8453
	IDOptimism    uint64 = 10

	// Testnets.
	IDOmniOmega   uint64 = 164
	IDHolesky     uint64 = 17000
	IDSepolia     uint64 = 11155111
	IDArbSepolia  uint64 = 421614
	IDOpSepolia   uint64 = 11155420
	IDBaseSepolia uint64 = 84532

	// Ephemeral.
	IDOmniStaging uint64 = 1650
	IDOmniDevnet  uint64 = 1651
	IDMockL1      uint64 = 1652
	IDMockL2      uint64 = 1654
	IDMockOp      uint64 = 1655
	IDMockArb     uint64 = 1656

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

func IsOmniEVM(name string) bool {
	return name == omniEVMName
}

// All returns all evmchain metadatas ordered by chain ID.
func All() []Metadata {
	var resp []Metadata
	for _, metadata := range static {
		resp = append(resp, metadata)
	}

	sort.Slice(resp, func(i, j int) bool {
		return resp[i].ChainID < resp[j].ChainID
	})

	return resp
}

var static = map[uint64]Metadata{
	IDEthereum: {
		ChainID:     IDEthereum,
		Name:        "ethereum",
		PrettyName:  "Ethereum",
		BlockPeriod: 12 * time.Second,
		NativeToken: tokens.ETH,
		Reorgs:      true,
	},
	IDOmniMainnet: {
		ChainID:     IDOmniMainnet,
		Name:        omniEVMName,
		PrettyName:  "Omni Mainnet",
		BlockPeriod: omniEVMBlockPeriod,
		NativeToken: tokens.OMNI,
	},
	IDArbitrumOne: {
		ChainID:     IDArbitrumOne,
		Name:        "arbitrum_one",
		PrettyName:  "Arbitrum One",
		BlockPeriod: 300 * time.Millisecond,
		NativeToken: tokens.ETH,
		PostsTo:     IDEthereum,
	},
	IDOptimism: {
		ChainID:     IDOptimism,
		Name:        "optimism",
		PrettyName:  "Optimism",
		BlockPeriod: 2 * time.Second,
		NativeToken: tokens.ETH,
		PostsTo:     IDEthereum,
	},
	IDBase: {
		ChainID:     IDBase,
		Name:        "base",
		PrettyName:  "Base",
		BlockPeriod: 2 * time.Second,
		NativeToken: tokens.ETH,
		PostsTo:     IDEthereum,
	},
	IDOmniOmega: {
		ChainID:     IDOmniOmega,
		Name:        omniEVMName,
		PrettyName:  "Omni Omega",
		BlockPeriod: omniEVMBlockPeriod,
		NativeToken: tokens.OMNI,
	},
	IDHolesky: {
		ChainID:     IDHolesky,
		Name:        "holesky",
		PrettyName:  "Holesky",
		BlockPeriod: 12 * time.Second,
		NativeToken: tokens.ETH,
		Reorgs:      true,
	},
	IDSepolia: {
		ChainID:     IDSepolia,
		Name:        "sepolia",
		PrettyName:  "Sepolia",
		BlockPeriod: 12 * time.Second,
		NativeToken: tokens.ETH,
		Reorgs:      true,
	},
	IDArbSepolia: {
		ChainID:     IDArbSepolia,
		Name:        "arb_sepolia",
		PrettyName:  "Arb Sepolia",
		BlockPeriod: 300 * time.Millisecond,
		NativeToken: tokens.ETH,
		PostsTo:     IDSepolia,
	},
	IDOpSepolia: {
		ChainID:     IDOpSepolia,
		Name:        "op_sepolia",
		PrettyName:  "OP Sepolia",
		BlockPeriod: 2 * time.Second,
		NativeToken: tokens.ETH,
		PostsTo:     IDSepolia,
	},
	IDBaseSepolia: {
		ChainID:     IDBaseSepolia,
		Name:        "base_sepolia",
		PrettyName:  "Base Sepolia",
		BlockPeriod: 2 * time.Second,
		NativeToken: tokens.ETH,
		PostsTo:     IDSepolia,
	},
	IDOmniDevnet: {
		ChainID:     IDOmniDevnet,
		Name:        omniEVMName,
		PrettyName:  "Omni Devnet",
		BlockPeriod: omniEVMBlockPeriod,
		NativeToken: tokens.OMNI,
	},
	IDOmniStaging: {
		ChainID:     IDOmniStaging,
		Name:        omniEVMName,
		PrettyName:  "Omni Staging",
		BlockPeriod: omniEVMBlockPeriod,
		NativeToken: tokens.OMNI,
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
