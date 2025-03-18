package symbiotic

import (
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/common"
)

type flowConfig struct {
	srcChain  uint64
	dstChain  uint64
	vaultAddr common.Address
}

var config = map[netconf.ID]flowConfig{
	netconf.Devnet: {
		srcChain: evmchain.IDMockL1,
		dstChain: evmchain.IDMockL2,
		// Deterministic address on devnet
		vaultAddr: common.HexToAddress("0x81487c7b22a0babadC98D5cA1d7D21240beB14Cc"),
	},

	// TODO(christian): enable once this is needed.
	// netconf.Omega: {
	// 	srcChain:     evmchain.IDBaseSepolia,
	// 	dstChain:     evmchain.IDHolesky,
	// 	vaultAddr:    targets.SymbioticHoleskyWSTETHVault1,
	// },
}
