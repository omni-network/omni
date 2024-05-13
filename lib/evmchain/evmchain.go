// Package evmchain provides static metadata about supported evm chains.
package evmchain

import (
	"time"

	"github.com/omni-network/omni/lib/tokens"
)

const (
	// Mainnets.
	IDEthereum    uint64 = 1
	IDOmniMainnet uint64 = 166

	// Testnets.
	IDOmniTestnet uint64 = 165
	IDHolesky     uint64 = 17000
	IDArbSepolia  uint64 = 421614
	IDOpSepolia   uint64 = 11155420

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
	IDOmniTestnet: {
		ChainID:     IDOmniTestnet,
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
	},
	IDOpSepolia: {
		ChainID:     IDOpSepolia,
		Name:        "op_sepolia",
		BlockPeriod: 2 * time.Second,
		NativeToken: tokens.ETH,
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
		BlockPeriod: time.Second / 4,
		NativeToken: tokens.ETH,
	},
}
