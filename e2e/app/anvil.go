package app

import (
	"context"

	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/lib/anvil"
	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/errors"
)

// fundAnvil funds EOAs on anvil chains.
func fundAnvil(ctx context.Context, def Definition) error {
	if def.Testnet.Network.IsProtected() {
		return nil
	}

	amt := bi.Ether(1_000_000) // 1M Ether
	toFund := dedup(append(
		eoa.MustAddresses(def.Testnet.Network, eoa.AllRoles()...),
		append(eoa.DevAccounts(), eoa.CreateXDeployer())...,
	))

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

func dedup[T comparable](s []T) []T {
	seen := make(map[T]bool)
	var out []T

	for _, v := range s {
		if seen[v] {
			continue
		}

		seen[v] = true
		out = append(out, v)
	}

	return out
}
