package solve

import (
	"context"

	"github.com/omni-network/omni/lib/contracts/solvernet/inbox"
	"github.com/omni-network/omni/lib/contracts/solvernet/outbox"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"

	"golang.org/x/sync/errgroup"
)

// Deploy deploys solve inbox / outbox contracts, and devnet app (if devnet).
func Deploy(ctx context.Context, network netconf.Network, backends ethbackend.Backends) error {
	if !network.ID.IsEphemeral() {
		log.Warn(ctx, "Skipping solvernet deploy", nil)
		return nil
	}

	var eg errgroup.Group
	eg.Go(func() error { return deployBoxes(ctx, network, backends) })
	eg.Go(func() error { return maybeDeployMockTokens(ctx, network, backends) })
	eg.Go(func() error { return maybeFundSolver(ctx, network.ID, backends) })
	if err := eg.Wait(); err != nil {
		return errors.Wrap(err, "deploy")
	}

	return nil
}

func deployBoxes(ctx context.Context, network netconf.Network, backends ethbackend.Backends) error {
	var eg errgroup.Group
	for _, chain := range network.EVMChains() {
		eg.Go(func() error {
			backend, err := backends.Backend(chain.ID)
			if err != nil {
				return errors.Wrap(err, "get backend", "chain", chain.Name)
			}

			addr, _, err := inbox.DeployIfNeeded(ctx, network, backend)
			if err != nil {
				return errors.Wrap(err, "deploy solve inbox", "chain", chain.Name)
			}

			log.Debug(ctx, "SolverNetInbox deployed", "addr", addr.Hex(), "chain", chain.Name)

			addr, _, err = outbox.DeployIfNeeded(ctx, network, backend)
			if err != nil {
				return errors.Wrap(err, "deploy solve outbox", "chain", chain.Name)
			}

			log.Debug(ctx, "SolverNetOutbox deployed", "addr", addr.Hex(), "chain", chain.Name)

			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		return errors.Wrap(err, "deploy solvernet boxes")
	}

	return nil
}
