package app

import (
	"context"

	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/lib/anvil"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/umath"
)

// fundAnvil funds EOAs on anvil chains.
func fundAnvil(ctx context.Context, def Definition) error {
	if def.Testnet.Network.IsProtected() {
		return nil
	}

	toFund := eoa.MustAddresses(def.Testnet.Network, eoa.AllRoles()...)
	amt := umath.EtherToWei(1_000_000) // 1M Ether

	for _, chain := range def.Testnet.AnvilChains {
		backend, err := def.Backends().Backend(chain.Chain.ChainID)
		if err != nil {
			return errors.Wrap(err, "get backend")
		}

		if err := anvil.FundAccounts(ctx, backend.Client, amt, toFund...); err != nil {
			return errors.Wrap(err, "fund anvil account")
		}
	}

	return nil
}
