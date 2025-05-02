package app

import (
	"context"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/lib/contracts"
	"github.com/omni-network/omni/lib/contracts/solvernet"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

// SetSolverNetRoutes sets the SolverNet routes for all chains in a given network.
func SetSolverNetRoutes(ctx context.Context, network netconf.Network, backends ethbackend.Backends) error {
	addrs, err := contracts.GetAddresses(ctx, network.ID)
	if err != nil {
		return errors.Wrap(err, "get addresses", "network", network.ID)
	}

	chainIDs := network.EVMChains()
	for _, chain := range chainIDs {
		remoteChainIDs, remoteOutboxes, remoteInboxConfigs, err := getRoutes(chain, chainIDs, addrs.SolverNetInbox, addrs.SolverNetOutbox)
		if err != nil {
			return errors.Wrap(err, "get routes", "chain", chain.Name)
		}

		backend, err := backends.Backend(chain.ID)
		if err != nil {
			return errors.Wrap(err, "get backend", "chain", chain.Name)
		}

		isDeployed, err := isDeployed(ctx, backend, addrs.SolverNetInbox, addrs.SolverNetOutbox)
		if !isDeployed || err != nil {
			return errors.Wrap(err, "isDeployed", "chain", chain.Name)
		}

		txOpts, err := backend.BindOpts(ctx, eoa.MustAddress(network.ID, eoa.RoleManager))
		if err != nil {
			return errors.Wrap(err, "bind opts", "chain", chain.Name)
		}

		err = configureInbox(ctx, backend, txOpts, addrs.SolverNetInbox, remoteChainIDs, remoteOutboxes)
		if err != nil {
			return errors.Wrap(err, "configure inbox", "chain", chain.Name)
		}

		err = configureOutbox(ctx, backend, txOpts, addrs.SolverNetOutbox, remoteChainIDs, remoteInboxConfigs)
		if err != nil {
			return errors.Wrap(err, "configure outbox", "chain", chain.Name)
		}
	}

	return nil
}

// getRoutes returns the remote chain IDs, outboxes, and inbox configs for a given chain.
func getRoutes(chain netconf.Chain, allChains []netconf.Chain, inbox common.Address, outbox common.Address) ([]uint64, []common.Address, []bindings.ISolverNetOutboxInboxConfig, error) {
	var remoteChainIDs []uint64
	var remoteOutboxes []common.Address
	var remoteInboxConfigs []bindings.ISolverNetOutboxInboxConfig

	for _, c := range allChains {
		if c.ID != chain.ID {
			var solverNetInbox common.Address
			var solverNetOutbox common.Address
			provider := solvernet.None
			var err error

			if !solvernet.IsDisabled(chain.ID) && !solvernet.IsDisabled(c.ID) {
				solverNetInbox = inbox
				solverNetOutbox = outbox
				provider, err = solvernet.Provider(c.ID)
				if err != nil {
					return nil, nil, nil, errors.Wrap(err, "get provider", "chain", c.Name)
				}
			}

			remoteChainIDs = append(remoteChainIDs, c.ID)
			remoteOutboxes = append(remoteOutboxes, solverNetOutbox)
			remoteInboxConfigs = append(remoteInboxConfigs, bindings.ISolverNetOutboxInboxConfig{
				Inbox:    solverNetInbox,
				Provider: provider,
			})
		}
	}

	return remoteChainIDs, remoteOutboxes, remoteInboxConfigs, nil
}

// isDeployed returns true if the SolverNet inbox and outbox are deployed on a given chain.
func isDeployed(ctx context.Context, backend *ethbackend.Backend, inbox common.Address, outbox common.Address) (bool, error) {
	isDeployed, err := contracts.IsDeployed(ctx, backend, inbox)
	if !isDeployed || err != nil {
		return false, errors.Wrap(err, "solvernetinbox not deployed", "chain", backend.Name())
	}

	isDeployed, err = contracts.IsDeployed(ctx, backend, outbox)
	if !isDeployed || err != nil {
		return false, errors.Wrap(err, "solvernetoutbox not deployed", "chain", backend.Name())
	}

	return true, nil
}

// configureInbox configures the routes on a SolverNetInbox contract for a given chain.
func configureInbox(ctx context.Context, backend *ethbackend.Backend, txOpts *bind.TransactOpts, inbox common.Address, remoteChainIDs []uint64, remoteOutboxes []common.Address) error {
	solverNetInbox, err := bindings.NewSolverNetInbox(inbox, backend)
	if err != nil {
		return errors.Wrap(err, "bind inbox", "chain", backend.Name())
	}

	tx, err := solverNetInbox.SetOutboxes(txOpts, remoteChainIDs, remoteOutboxes)
	if err != nil {
		return errors.Wrap(err, "set outboxes on inbox", "chain", backend.Name())
	}

	_, err = backend.WaitMined(ctx, tx)
	if err != nil {
		return errors.Wrap(err, "wait mined set outboxes on inbox", "chain", backend.Name())
	}

	return nil
}

// configureOutbox configures the routes on a SolverNetOutbox contract for a given chain.
func configureOutbox(ctx context.Context, backend *ethbackend.Backend, txOpts *bind.TransactOpts, outbox common.Address, remoteChainIDs []uint64, remoteInboxConfigs []bindings.ISolverNetOutboxInboxConfig) error {
	solverNetOutbox, err := bindings.NewSolverNetOutbox(outbox, backend)
	if err != nil {
		return errors.Wrap(err, "bind outbox", "chain", backend.Name())
	}

	tx, err := solverNetOutbox.SetInboxes(txOpts, remoteChainIDs, remoteInboxConfigs)
	if err != nil {
		return errors.Wrap(err, "set inboxes on outbox", "chain", backend.Name())
	}

	_, err = backend.WaitMined(ctx, tx)
	if err != nil {
		return errors.Wrap(err, "wait mined set inboxes on outbox", "chain", backend.Name())
	}

	return nil
}
