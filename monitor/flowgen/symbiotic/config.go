package symbiotic

import (
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/tokens"
	"github.com/omni-network/omni/solver/targets"

	"github.com/ethereum/go-ethereum/common"
)

type flowConfig struct {
	srcChain     uint64
	dstChain     uint64
	depositToken tokens.Token
	expenseToken tokens.Token
	vaultAddr    common.Address
}

var config = map[netconf.ID]flowConfig{
	netconf.Omega: {
		srcChain:     evmchain.IDBaseSepolia,
		dstChain:     evmchain.IDHolesky,
		depositToken: tokens.WSTETH,
		expenseToken: tokens.WSTETH,
		vaultAddr:    targets.SymbioticHoleskyWSTETHVault1,
	},

	netconf.Devnet: {
		srcChain:     evmchain.IDMockL1,
		dstChain:     evmchain.IDMockL2,
		depositToken: tokens.WSTETH,
		expenseToken: tokens.WSTETH,
		// Deterministic address on devnet
		vaultAddr: common.HexToAddress("0x81487c7b22a0babadC98D5cA1d7D21240beB14Cc"),
	},
}
