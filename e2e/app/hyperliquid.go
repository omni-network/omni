package app

import (
	"context"

	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/fireblocks"
	"github.com/omni-network/omni/lib/hypercore"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
)

// HyperliquidUseBigBlocks enables big blocks on Hyperliquid for specified roles per network.
func HyperliquidUseBigBlocks(
	ctx context.Context,
	networkID netconf.ID,
	fireCl fireblocks.Client,
) error {
	if networkID == netconf.Devnet {
		return errors.New("devnet not supported")
	}

	// list of roles for which to enable big blocks
	// accounts are required to have received least 1 USDC on HyperCore
	// enabling big blocks when already enabled is a no-op
	roles := []eoa.Role{
		eoa.RoleDeployer,
	}

	newClient := func(signer hypercore.Signer) hypercore.Client {
		if networkID == netconf.Mainnet {
			return hypercore.NewClient(signer)
		}

		return hypercore.NewTestnetClient(signer)
	}

	newSigner := func(role eoa.Role) (hypercore.Signer, error) {
		acc, ok := eoa.AccountForRole(networkID, role)
		if !ok {
			return nil, errors.New("no account for role")
		}

		if acc.Type == eoa.TypeRemote {
			return hypercore.NewFireblocksSigner(fireCl, acc.Address), nil
		}

		pk, err := eoa.PrivateKey(ctx, networkID, role)
		if err != nil {
			return nil, errors.Wrap(err, "private key")
		}

		return hypercore.NewPrivateKeySigner(pk), nil
	}

	for _, role := range roles {
		signer, err := newSigner(role)
		if err != nil {
			return errors.Wrap(err, "new signer", "role", role, "network", networkID)
		}

		client := newClient(signer)

		if err := client.UseBigBlocks(ctx); err != nil {
			return errors.Wrap(err, "use big blocks", "role", role, "network", networkID)
		}

		log.Info(ctx, "Enabled Hyperliquid big blocks", "role", role, "network", networkID)
	}

	return nil
}
