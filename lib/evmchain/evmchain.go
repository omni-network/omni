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
	IDArbSepolia  uint64 = 421614
	IDOpSepolia   uint64 = 11155420
	IDBaseSepolia uint64 = 84532

	// Ephemeral.
	IDOmniEphemeral uint64 = 1651
	IDMockL1Fast    uint64 = 1652
	IDMockL1Slow    uint64 = 1653
	IDMockL2        uint64 = 1654
	IDMockOp        uint64 = 1655
	IDMockArb       uint64 = 1656

	omniEVMName        = "omni_evm"
	omniEVMBlockPeriod = time.Second * 2
)

type Metadata struct {
	ChainID     uint64
	IsL2        bool
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

var static = map[uint64]Metadata{
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
	IDArbSepolia: {
		ChainID:     IDArbSepolia,
		Name:        "arb_sepolia",
		BlockPeriod: 300 * time.Millisecond,
		NativeToken: tokens.ETH,
		IsL2:        true,
	},
	IDOpSepolia: {
		ChainID:     IDOpSepolia,
		Name:        "op_sepolia",
		BlockPeriod: 2 * time.Second,
		NativeToken: tokens.ETH,
		IsL2:        true,
	},
	IDBaseSepolia: {
		ChainID:     IDBaseSepolia,
		Name:        "base_sepolia",
		BlockPeriod: 2 * time.Second,
		NativeToken: tokens.ETH,
		IsL2:        true,
	},
	IDOmniEphemeral: {
		ChainID:     IDOmniEphemeral,
		Name:        omniEVMName,
		BlockPeriod: omniEVMBlockPeriod,
		NativeToken: tokens.OMNI,
	},
	IDMockL1Fast: {
		ChainID:     IDMockL1Fast,
		Name:        "mock_l1",
		BlockPeriod: time.Second,
		NativeToken: tokens.ETH,
	},
	IDMockL1Slow: {
		ChainID:     IDMockL1Slow,
		Name:        "slow_l1",
		BlockPeriod: time.Second * 12,
		NativeToken: tokens.ETH,
	},
	IDMockL2: {
		ChainID:     IDMockL2,
		Name:        "mock_l2",
		BlockPeriod: time.Second,
		NativeToken: tokens.ETH,
		IsL2:        true,
	},
	IDMockOp: {
		ChainID:     IDMockOp,
		Name:        "mock_op",
		BlockPeriod: time.Second * 2,
		NativeToken: tokens.ETH,
		IsL2:        true,
	},
	IDMockArb: {
		ChainID:     IDMockArb,
		Name:        "mock_arb",
		BlockPeriod: time.Second,
		NativeToken: tokens.ETH,
		IsL2:        true,
	},
}
