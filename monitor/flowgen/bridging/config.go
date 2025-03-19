package bridging

import (
	"math/big"

	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/netconf"
)

type flowConfig struct {
	srcChain  uint64
	dstChain  uint64
	orderSize *big.Int
}

var config = map[netconf.ID]flowConfig{
	netconf.Devnet: {
		srcChain:  evmchain.IDMockL1,
		dstChain:  evmchain.IDMockL2,
		orderSize: bi.Ether(0.02),
	},

	netconf.Staging: {
		srcChain:  evmchain.IDBaseSepolia,
		dstChain:  evmchain.IDOpSepolia,
		orderSize: bi.Ether(0.02),
	},

	netconf.Omega: {
		srcChain:  evmchain.IDOpSepolia,
		dstChain:  evmchain.IDArbSepolia,
		orderSize: bi.Ether(0.02),
	},

	netconf.Mainnet: {
		srcChain:  evmchain.IDOptimism,
		dstChain:  evmchain.IDArbitrumOne,
		orderSize: bi.Ether(0.02),
	},
}
