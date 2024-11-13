package solve

import (
	"context"

	"github.com/omni-network/omni/e2e/solve/devapp"
	"github.com/omni-network/omni/lib/contracts/solveinbox"
	"github.com/omni-network/omni/lib/contracts/solveoutbox"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
)

// DeployContracts deploys solve inbox / outbox contracts, and devnet app (if devnet).
func DeployContracts(ctx context.Context, network netconf.Network, backends ethbackend.Backends) error {
	if network.ID != netconf.Devnet {
		log.Warn(ctx, "Skipping solve deploy", nil)
		return nil
	}

	log.Info(ctx, "Deploying solve contracts")

	if err := deployBoxes(ctx, network, backends); err != nil {
		return errors.Wrap(err, "deploy boxes")
	}

	if err := devapp.Deploy(ctx, network, backends); err != nil {
		return errors.Wrap(err, "deploy devapp")
	}

	return nil
}

// DeployContracts deploys solve inbox / outbox contracts.
func deployBoxes(ctx context.Context, network netconf.Network, backends ethbackend.Backends) error {
	for _, chain := range network.EVMChains() {
		backend, err := backends.Backend(chain.ID)
		if err != nil {
			return errors.Wrap(err, "get backend", "chain", chain.Name)
		}

		addr, _, err := solveinbox.DeployIfNeeded(ctx, network.ID, backend)
		if err != nil {
			return errors.Wrap(err, "deploy solve inbox")
		}

		log.Info(ctx, "SolveInbox deployed", "addr", addr.Hex(), "chain", chain.Name)

		addr, _, err = solveoutbox.DeployIfNeeded(ctx, network.ID, backend)
		if err != nil {
			return errors.Wrap(err, "deploy solve outbox")
		}

		log.Info(ctx, "SolveOutbox deployed", "addr", addr.Hex(), "chain", chain.Name)
	}

	return nil
}
