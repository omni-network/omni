package app

import (
	"context"

	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/lib/anvil"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/netconf"

	"cosmossdk.io/math"
)

// fundAnvil funds EOAs on anvil chains.
func fundAnvil(ctx context.Context, def Definition) error {
	if def.Testnet.Network.IsProtected() {
		return nil
	}

	toFund := eoa.MustAddresses(netconf.Devnet, eoa.AllRoles()...)
	amt := math.NewInt(1000000).MulRaw(1e18).BigInt() // 1M ETH

	for _, chain := range def.Testnet.AnvilChains {
		if err := anvil.FundAccounts(ctx, chain.ExternalRPC, amt, toFund...); err != nil {
			return errors.Wrap(err, "fund anvil account")
		}
	}

	return nil
}
