// Package evmchain provides static metadata about supported evm chains.
//
// This package should only contain public well-known metadata and
// should not be Omni-network specific, since multiple omni networks
// can be deployed to the same evm chain.
package evmchain

import (
	"sort"
	"time"

	"github.com/omni-network/omni/lib/tokens"
)

const (
	// Mainnets.
	IDEthereum    uint64 = 1
	IDOmniMainnet uint64 = 166

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

	forkChainIDOffset = 1100000000
)

type Metadata struct {
	ChainID     uint64
	PostsTo     uint64 // chain id to which tx data is posted
	Name        string
	BlockPeriod time.Duration
	NativeToken tokens.Token
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

func addForks(mds map[uint64]Metadata) map[uint64]Metadata {
	withForks := make(map[uint64]Metadata)

	for _, md := range mds {
		withForks[md.ChainID] = md

		if shouldFork(md.ChainID) {
			// omit PostsTo in fork meta - no need for proper data fees
			withForks[ForkChainID(md.ChainID)] = Metadata{
				ChainID:     ForkChainID(md.ChainID),
				Name:        md.Name + "_fork",
				BlockPeriod: md.BlockPeriod,
				NativeToken: md.NativeToken,
			}
		}
	}

	return withForks
}

func ForkChainID(chainID uint64) uint64 {
	return chainID + forkChainIDOffset
}

func shouldFork(chainID uint64) bool {
	// add fork metadata for all non-omni and all non-mock chains
	return !(chainID == IDOmniMainnet ||
		chainID == IDOmniOmega ||
		chainID == IDOmniDevnet ||
		chainID == IDOmniStaging ||
		chainID == IDMockL1 ||
		chainID == IDMockL2 ||
		chainID == IDMockOp ||
		chainID == IDMockArb)
}

var static = addForks(map[uint64]Metadata{
	IDEthereum: {
		ChainID:     IDEthereum,
		Name:        "ethereum",
		BlockPeriod: 12 * time.Second,
		NativeToken: tokens.ETH,
	},
	IDOmniMainnet: {
		ChainID:     IDOmniMainnet,
		Name:        omniEVMName,
		BlockPeriod: omniEVMBlockPeriod,
		NativeToken: tokens.OMNI,
	},
	IDOmniOmega: {
		ChainID:     IDOmniOmega,
		Name:        omniEVMName,
		BlockPeriod: omniEVMBlockPeriod,
		NativeToken: tokens.OMNI,
	},
	IDHolesky: {
		ChainID:     IDHolesky,
		Name:        "holesky",
		BlockPeriod: 12 * time.Second,
		NativeToken: tokens.ETH,
	},
	IDSepolia: {
		ChainID:     IDSepolia,
		Name:        "sepolia",
		BlockPeriod: 12 * time.Second,
		NativeToken: tokens.ETH,
	},
	IDArbSepolia: {
		ChainID:     IDArbSepolia,
		Name:        "arb_sepolia",
		BlockPeriod: 300 * time.Millisecond,
		NativeToken: tokens.ETH,
		PostsTo:     IDSepolia,
	},
	IDOpSepolia: {
		ChainID:     IDOpSepolia,
		Name:        "op_sepolia",
		BlockPeriod: 2 * time.Second,
		NativeToken: tokens.ETH,
		PostsTo:     IDSepolia,
	},
	IDBaseSepolia: {
		ChainID:     IDBaseSepolia,
		Name:        "base_sepolia",
		BlockPeriod: 2 * time.Second,
		NativeToken: tokens.ETH,
		PostsTo:     IDSepolia,
	},
	IDOmniDevnet: {
		ChainID:     IDOmniDevnet,
		Name:        omniEVMName,
		BlockPeriod: omniEVMBlockPeriod,
		NativeToken: tokens.OMNI,
	},
	IDOmniStaging: {
		ChainID:     IDOmniStaging,
		Name:        omniEVMName,
		BlockPeriod: omniEVMBlockPeriod,
		NativeToken: tokens.OMNI,
	},
	IDMockL1: {
		ChainID:     IDMockL1,
		Name:        "mock_l1",
		BlockPeriod: time.Second,
		NativeToken: tokens.ETH,
	},
	IDMockL2: {
		ChainID:     IDMockL2,
		Name:        "mock_l2",
		BlockPeriod: time.Second,
		NativeToken: tokens.ETH,
	},
	IDMockOp: {
		ChainID:     IDMockOp,
		Name:        "mock_op",
		BlockPeriod: time.Second * 2,
		NativeToken: tokens.ETH,
	},
	IDMockArb: {
		ChainID:     IDMockArb,
		Name:        "mock_arb",
		BlockPeriod: time.Second,
		NativeToken: tokens.ETH,
	},
})
