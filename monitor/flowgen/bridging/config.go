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

var config = map[netconf.ID][]flowConfig{
	netconf.Devnet: {
		{
			// Native ETH <> Native ETH
			srcChain: evmchain.IDMockL1,
			dstChain: evmchain.IDMockL2,
		},
		{
			// Native ETH <> Native Omni
			srcChain: evmchain.IDMockL1,
			dstChain: evmchain.IDOmniDevnet,
		},
	},

	netconf.Staging: {
		{
			// Native ETH <> Native ETH
			srcChain: evmchain.IDBaseSepolia,
			dstChain: evmchain.IDOpSepolia,
		},
		{
			// Native ETH <> Native Omni
			srcChain: evmchain.IDArbSepolia,
			dstChain: evmchain.IDOmniStaging,
		},
	},

	netconf.Omega: {
		{
			// Native ETH <> Native ETH
			srcChain: evmchain.IDOpSepolia,
			dstChain: evmchain.IDArbSepolia,
		},
		{
			// Native ETH <> Native Omni
			srcChain: evmchain.IDBaseSepolia,
			dstChain: evmchain.IDOmniOmega,
		},
	},

	netconf.Mainnet: {
		{
			// Native ETH <> Native ETH
			srcChain: evmchain.IDOptimism,
			dstChain: evmchain.IDArbitrumOne,
		},
		// TODO(corver): Enable once mainnet supports swaps.
		// {
		//	// Native ETH <> Native Omni
		//	srcChain: evmchain.IDBase,
		//	dstChain: evmchain.IDOmniMainnet,
		// },
	},
}
