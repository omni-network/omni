package bridging

import (
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/netconf"
)

type flowConfig struct {
	srcChain uint64
	dstChain uint64
}

func (c flowConfig) Flip() flowConfig {
	return flowConfig{
		srcChain: c.dstChain,
		dstChain: c.srcChain,
	}
}

var config = map[netconf.ID]flowConfig{
	netconf.Devnet: {
		srcChain: evmchain.IDMockL1,
		dstChain: evmchain.IDMockL2,
	},

	netconf.Staging: {
		srcChain: evmchain.IDBaseSepolia,
		dstChain: evmchain.IDOpSepolia,
	},

	netconf.Omega: {
		srcChain: evmchain.IDOpSepolia,
		dstChain: evmchain.IDArbSepolia,
	},

	netconf.Mainnet: {
		srcChain: evmchain.IDOptimism,
		dstChain: evmchain.IDArbitrumOne,
	},
}
