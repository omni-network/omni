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

func (c flowConfig) IsSwap() bool {
	src, _ := evmchain.MetadataByID(c.srcChain)
	dst, _ := evmchain.MetadataByID(c.dstChain)

	return src.NativeToken != dst.NativeToken
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
		{
			// Native ETH <> Native Omni
			srcChain: evmchain.IDBase,
			dstChain: evmchain.IDOmniMainnet,
		},
	},
}
