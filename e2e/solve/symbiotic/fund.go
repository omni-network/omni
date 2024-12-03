package symbiotic

import (
	"context"

	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/lib/anvil"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/params"

	"cosmossdk.io/math"
)

func FundSolver(ctx context.Context, network netconf.ID, backends ethbackend.Backends) error {
	// funding solver with l1 wsETH uses anvil_setStorageAt util, which is only available on devnet
	if network != netconf.Devnet {
		return errors.New("only devnet")
	}

	app := MustGetApp(network)

	ethCl, ok := backends.Clients()[app.L1.ChainID]
	if !ok {
		return errors.New("missing eth client", "chain", app.L1.Name)
	}

	eth1m := math.NewInt(1_000_000).MulRaw(params.Ether).BigInt()

	return anvil.FundERC20(ctx, ethCl, app.L1wstETH, eth1m, eoa.MustAddress(netconf.Devnet, eoa.RoleSolver))
}
