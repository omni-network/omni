package bridging

import (
	"math/big"

	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/netconf"
)

type flowConfig struct {
	srcChain uint64
	dstChain uint64

	minOrderSize *big.Int
	maxOrderSize *big.Int
}

var config = map[netconf.ID]flowConfig{
	netconf.Devnet: {
		srcChain:     evmchain.IDMockL1,
		dstChain:     evmchain.IDMockL2,
		minOrderSize: bi.Ether(0.2),
		maxOrderSize: bi.Ether(3),
	},

	netconf.Staging: {
		srcChain:     evmchain.IDBaseSepolia,
		dstChain:     evmchain.IDOpSepolia,
		minOrderSize: bi.Ether(0.2),
		maxOrderSize: bi.Ether(3),
	},

	netconf.Omega: {
		srcChain:     evmchain.IDOpSepolia,
		dstChain:     evmchain.IDArbSepolia,
		minOrderSize: bi.Ether(0.2),
		maxOrderSize: bi.Ether(3),
	},

	netconf.Mainnet: {
		srcChain:     evmchain.IDOptimism,
		dstChain:     evmchain.IDArbitrumOne,
		minOrderSize: bi.Ether(0.2),
		maxOrderSize: bi.Ether(3),
	},
}
