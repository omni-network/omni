package symbiotic

import (
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/tokens"

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
	// TODO(christian): enable once tested on devnet
	// netconf.Omega: {
	// 	srcChain:     evmchain.IDBaseSepolia,
	// 	dstChain:     evmchain.IDHolesky,
	// 	depositToken: tokens.WSTETH,
	// 	expenseToken: tokens.WSTETH,
	// 	vaultAddr:    targets.SymbioticHoleskyWSTETHVault1,
	// },
}
