//nolint:dupl // similar code for different contracts
package solve

import (
	"context"

	"github.com/omni-network/omni/e2e/solve/devapp"
	"github.com/omni-network/omni/e2e/solve/symbiotic"
	"github.com/omni-network/omni/lib/contracts/solveinbox"
	"github.com/omni-network/omni/lib/contracts/solveoutbox"
	"github.com/omni-network/omni/lib/contracts/solvernet/inbox"
	"github.com/omni-network/omni/lib/contracts/solvernet/outbox"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/feature"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"

	"golang.org/x/sync/errgroup"
)

// DeployContracts deploys solve inbox / outbox contracts, and devnet app (if devnet).
func DeployContracts(ctx context.Context, network netconf.Network, backends ethbackend.Backends) error {
	if !network.ID.IsEphemeral() {
		log.Warn(ctx, "Skipping solve deploy", nil)
		return nil
	}

	if feature.FlagSolverV2.Enabled(ctx) {
		log.Info(ctx, "Deploying solve v2 contracts")
		return deployV2Boxes(ctx, network, backends)

		// TODO: remove below when v2 is fully enabled
		//
		// remove devapp
		// replace symbiotic target with symbiotic tokens in solver/app/v2/tokens.go
		//
		// we do use the devapp.L2Token MockERC20 as "wstETH" on base sepolia
		// for testnet symbiotic solving (because there is no canonical wstETH
		// contract on base sepolia)
		//
		// we can replace that by deploying MockERC20s on all chains here,
		// tracking them in solver/app/v2/tokens.go, and use them for solver
		// tests / demos when needed.
	}

	log.Info(ctx, "Deploying solve contracts")
	if err := deployBoxes(ctx, network, backends); err != nil {
		return errors.Wrap(err, "deploy boxes")
	}

	// TODO(kevin): idempotent outbox allow calls
	var eg errgroup.Group
	eg.Go(func() error { return devapp.MaybeDeploy(ctx, network.ID, backends) })
	eg.Go(func() error { return devapp.AllowOutboxCalls(ctx, network.ID, backends) })
	eg.Go(func() error { return symbiotic.MaybeFundSolver(ctx, network.ID, backends) })
	eg.Go(func() error { return symbiotic.AllowOutboxCalls(ctx, network.ID, backends) })
	if err := eg.Wait(); err != nil {
		return errors.Wrap(err, "setup targets")
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

			addr, _, err := solveinbox.DeployIfNeeded(ctx, network.ID, backend)
			if err != nil {
				return errors.Wrap(err, "deploy solve inbox")
			}

			log.Debug(ctx, "SolveInbox deployed", "addr", addr.Hex(), "chain", chain.Name)

			addr, _, err = solveoutbox.DeployIfNeeded(ctx, network.ID, backend)
			if err != nil {
				return errors.Wrap(err, "deploy solve outbox")
			}

			log.Debug(ctx, "SolveOutbox deployed", "addr", addr.Hex(), "chain", chain.Name)

			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		return errors.Wrap(err, "deploy solver boxes")
	}

	return nil
}

func deployV2Boxes(ctx context.Context, network netconf.Network, backends ethbackend.Backends) error {
	var eg errgroup.Group
	for _, chain := range network.EVMChains() {
		eg.Go(func() error {
			backend, err := backends.Backend(chain.ID)
			if err != nil {
				return errors.Wrap(err, "get backend", "chain", chain.Name)
			}

			addr, _, err := inbox.DeployIfNeeded(ctx, network.ID, backend)
			if err != nil {
				return errors.Wrap(err, "deploy solve inbox")
			}

			log.Debug(ctx, "SolverNetInbox deployed", "addr", addr.Hex(), "chain", chain.Name)

			addr, _, err = outbox.DeployIfNeeded(ctx, network.ID, backend)
			if err != nil {
				return errors.Wrap(err, "deploy solve outbox")
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
