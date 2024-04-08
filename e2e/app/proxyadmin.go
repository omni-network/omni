package app

import (
	"context"

	"github.com/omni-network/omni/lib/contracts/proxyadmin"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
)

func deployPrivateProxyAdmin(ctx context.Context, def Definition) error {
	for _, c := range def.Testnet.AnvilChains {
		if err := deployProxyAdmin(ctx, def, c.Chain.ID); err != nil {
			return errors.Wrap(err, "deploy proxy admin")
		}
	}

	// only deploy to omni evm once
	if len(def.Testnet.OmniEVMs) > 0 {
		c := def.Testnet.OmniEVMs[0]
		if err := deployProxyAdmin(ctx, def, c.Chain.ID); err != nil {
			return errors.Wrap(err, "deploy proxy admin")
		}
	}

	return nil
}

func deployPublicProxyAdmin(ctx context.Context, def Definition) error {
	for _, c := range def.Testnet.PublicChains {
		if err := deployProxyAdmin(ctx, def, c.Chain().ID); err != nil {
			return errors.Wrap(err, "deploy proxy admin")
		}
	}

	return nil
}

func deployProxyAdmin(ctx context.Context, def Definition, chainID uint64) error {
	backend, err := def.Backends().Backend(chainID)
	if err != nil {
		return errors.Wrap(err, "backend")
	}

	addr, receipt, err := proxyadmin.DeployIfNeeded(ctx, def.Testnet.Network, backend)
	if err != nil {
		return errors.Wrap(err, "deploy")
	}

	if receipt != nil {
		log.Info(ctx, "Deployed ProxyAdmin", "chain", chainID, "address", addr.Hex(), "block", receipt.BlockNumber)
	} else {
		log.Info(ctx, "ProxyAdmin already deployed", "chain", chainID, "address", addr.Hex())
	}

	return nil
}
