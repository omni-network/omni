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

	"github.com/ethereum/go-ethereum/common"
)

func SetSolverNetRoutes(ctx context.Context, network netconf.Network, backends ethbackend.Backends) error {
	addrs, err := contracts.GetAddresses(ctx, network.ID)
	if err != nil {
		return errors.Wrap(err, "get addresses", "network", network.ID)
	}

	chainIDs := network.EVMChains()
	manager := eoa.MustAddress(network.ID, eoa.RoleManager)

	for _, chain := range chainIDs {
		var remoteChainIDs []uint64
		var remoteOutboxes []common.Address
		var remoteInboxConfigs []bindings.ISolverNetOutboxInboxConfig
		for _, c := range chainIDs {
			if c.ID != chain.ID {
				remoteChainIDs = append(remoteChainIDs, c.ID)
				remoteOutboxes = append(remoteOutboxes, addrs.SolverNetOutbox)

				provider, err := solvernet.Provider(c.ID)
				if err != nil {
					return errors.Wrap(err, "get provider", "chain", c.Name)
				}

				remoteInboxConfigs = append(remoteInboxConfigs, bindings.ISolverNetOutboxInboxConfig{
					Inbox:    addrs.SolverNetInbox,
					Provider: provider,
				})
			}
		}

		backend, err := backends.Backend(chain.ID)
		if err != nil {
			return errors.Wrap(err, "get backend", "chain", chain.Name)
		}

		isDeployed, err := contracts.IsDeployed(ctx, backend, addrs.SolverNetInbox)
		if !isDeployed || err != nil {
			return errors.Wrap(err, "solvernetinbox not deployed", "chain", chain.Name)
		}
		isDeployed, err = contracts.IsDeployed(ctx, backend, addrs.SolverNetOutbox)
		if !isDeployed || err != nil {
			return errors.Wrap(err, "solvernetoutbox not deployed", "chain", chain.Name)
		}

		txOpts, err := backend.BindOpts(ctx, manager)
		if err != nil {
			return errors.Wrap(err, "bind opts", "chain", chain.Name)
		}

		inbox, err := bindings.NewSolverNetInbox(addrs.SolverNetInbox, backend)
		if err != nil {
			return errors.Wrap(err, "bind inbox", "chain", chain.Name)
		}

		tx, err := inbox.SetOutboxes(txOpts, remoteChainIDs, remoteOutboxes)
		if err != nil {
			return errors.Wrap(err, "set outboxes on inbox", "chain", chain.Name)
		}

		_, err = backend.WaitMined(ctx, tx)
		if err != nil {
			return errors.Wrap(err, "wait mined set outboxes on inbox", "chain", chain.Name)
		}

		outbox, err := bindings.NewSolverNetOutbox(addrs.SolverNetOutbox, backend)
		if err != nil {
			return errors.Wrap(err, "bind outbox", "chain", chain.Name)
		}

		tx, err = outbox.SetInboxes(txOpts, remoteChainIDs, remoteInboxConfigs)
		if err != nil {
			return errors.Wrap(err, "set inboxes on outbox", "chain", chain.Name)
		}

		_, err = backend.WaitMined(ctx, tx)
		if err != nil {
			return errors.Wrap(err, "wait mined set inboxes on outbox", "chain", chain.Name)
		}
	}

	return nil
}
