package solve

import (
	"context"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/lib/contracts"
	"github.com/omni-network/omni/lib/contracts/solvernet"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"golang.org/x/sync/errgroup"
)

// Route represents a SolverNet route from a source chain to a destination chain.
type Route struct {
	ChainID      uint64
	InboxConfig  common.Address
	OutboxConfig bindings.ISolverNetOutboxInboxConfig
}

// SetSolverNetRoutes sets the SolverNet routes for all chains in a given network.
func SetSolverNetRoutes(ctx context.Context, network netconf.Network, backends ethbackend.Backends) error {
	addrs, err := contracts.GetAddresses(ctx, network.ID)
	if err != nil {
		return errors.Wrap(err, "get addresses", "network", network.ID)
	}

	eg, childCtx := errgroup.WithContext(ctx)

	for _, chain := range network.EVMChains() {
		backend, err := backends.Backend(chain.ID)
		if err != nil {
			return errors.Wrap(err, "get backend", "chain", chain.Name)
		}

		routes := getRoutes(chain, network, addrs.SolverNetInbox, addrs.SolverNetOutbox)
		routes, err = filterRoutes(ctx, chain, network, backend, addrs.SolverNetInbox, addrs.SolverNetOutbox, routes)
		if err != nil {
			return errors.Wrap(err, "filter routes", "chain", chain.Name)
		}

		isDeployed, err := checkDeployed(ctx, backend, addrs.SolverNetInbox, addrs.SolverNetOutbox)
		if !isDeployed || err != nil {
			return errors.Wrap(err, "isDeployed", "chain", chain.Name)
		}

		txOpts, err := backend.BindOpts(ctx, eoa.MustAddress(network.ID, eoa.RoleManager))
		if err != nil {
			return errors.Wrap(err, "bind opts", "chain", chain.Name)
		}

		// Capture loop variables for the goroutine closure to avoid race conditions
		_chain := chain
		eg.Go(func() error {
			err := configureInbox(childCtx, backend, txOpts, addrs.SolverNetInbox, routes)
			if err != nil {
				return errors.Wrap(err, "configure inbox", "chain", _chain.Name)
			}

			err = configureOutbox(childCtx, backend, txOpts, addrs.SolverNetOutbox, routes)
			if err != nil {
				return errors.Wrap(err, "configure outbox", "chain", _chain.Name)
			}

			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		return errors.Wrap(err, "set routes")
	}

	return nil
}

// getRoutes returns the remote chain IDs, outboxes, and inbox configs for a given chain.
func getRoutes(src netconf.Chain, network netconf.Network, inboxAddr common.Address, outboxAddr common.Address) []Route {
	var routes []Route
	for _, dest := range network.EVMChains() {
		if solvernet.IsDisabled(src.ID) || solvernet.IsDisabled(dest.ID) {
			// If disabled, configure zero values for routes to/from disabled chains.
			routes = append(routes, Route{ChainID: dest.ID})
			continue
		}

		provider, ok := solvernet.Provider(src.ID, dest.ID)
		if !ok {
			continue
		}

		routes = append(routes, Route{
			ChainID:      dest.ID,
			InboxConfig:  outboxAddr,
			OutboxConfig: bindings.ISolverNetOutboxInboxConfig{Inbox: inboxAddr, Provider: provider},
		})
	}

	return routes
}

// filterRoutes filters out routes that are already configured on a given chain.
func filterRoutes(ctx context.Context, src netconf.Chain, network netconf.Network, backend *ethbackend.Backend, inboxAddr common.Address, outboxAddr common.Address, routes []Route) ([]Route, error) {
	var currentRoutes []Route
	for _, dest := range network.EVMChains() {
		callOpts := &bind.CallOpts{Context: ctx}

		inbox, err := bindings.NewSolverNetInbox(inboxAddr, backend)
		if err != nil {
			return nil, errors.Wrap(err, "bind inbox", "chain", backend.Name())
		}

		inboxConfig, err := inbox.GetOutbox(callOpts, dest.ID)
		if err != nil {
			return nil, errors.Wrap(err, "get inbox outbox", "chain", backend.Name())
		}

		outbox, err := bindings.NewSolverNetOutbox(outboxAddr, backend)
		if err != nil {
			return nil, errors.Wrap(err, "bind outbox", "chain", backend.Name())
		}

		outboxConfig, err := outbox.GetInboxConfig(callOpts, src.ID)
		if err != nil {
			return nil, errors.Wrap(err, "get outbox inbox config", "chain", backend.Name())
		}

		currentRoutes = append(currentRoutes, Route{
			ChainID:      dest.ID,
			InboxConfig:  inboxConfig,
			OutboxConfig: outboxConfig,
		})
	}

	// Filter out routes that are already configured.
	var filteredRoutes []Route
	for _, route := range routes {
		for _, currentRoute := range currentRoutes {
			if route.ChainID == currentRoute.ChainID && route != currentRoute {
				filteredRoutes = append(filteredRoutes, route)
			}
		}
	}

	return filteredRoutes, nil
}

// checkDeployed returns true if the SolverNet inbox and outbox are deployed on a given chain.
func checkDeployed(ctx context.Context, backend *ethbackend.Backend, inbox common.Address, outbox common.Address) (bool, error) {
	isDeployed, err := contracts.IsDeployed(ctx, backend, inbox)
	if !isDeployed {
		return false, errors.New("inbox not deployed", "chain", backend.Name())
	} else if err != nil {
		return false, errors.Wrap(err, "is deployed inbox", "chain", backend.Name())
	}

	isDeployed, err = contracts.IsDeployed(ctx, backend, outbox)
	if !isDeployed {
		return false, errors.New("outbox not deployed", "chain", backend.Name())
	} else if err != nil {
		return false, errors.Wrap(err, "is deployed outbox", "chain", backend.Name())
	}

	return true, nil
}

// configureInbox configures the routes on a SolverNetInbox contract for a given chain.
func configureInbox(ctx context.Context, backend *ethbackend.Backend, txOpts *bind.TransactOpts, inbox common.Address, routes []Route) error {
	solverNetInbox, err := bindings.NewSolverNetInbox(inbox, backend)
	if err != nil {
		return errors.Wrap(err, "bind inbox", "chain", backend.Name())
	}

	tx, err := solverNetInbox.SetOutboxes(txOpts, chainIDs(routes), outboxes(routes))
	if err != nil {
		return errors.Wrap(err, "set outboxes", "chain", backend.Name())
	}

	receipt, err := backend.WaitMined(ctx, tx)
	if err != nil {
		return errors.Wrap(err, "wait mined", "chain", backend.Name())
	}

	log.Info(ctx, "SolverNetInbox configured", "addr", inbox.Hex(), "chain", backend.Name(), "block", receipt.BlockNumber, "tx", maybeTxHash(receipt))

	return nil
}

// configureOutbox configures the routes on a SolverNetOutbox contract for a given chain.
func configureOutbox(ctx context.Context, backend *ethbackend.Backend, txOpts *bind.TransactOpts, outbox common.Address, routes []Route) error {
	solverNetOutbox, err := bindings.NewSolverNetOutbox(outbox, backend)
	if err != nil {
		return errors.Wrap(err, "bind outbox", "chain", backend.Name())
	}

	tx, err := solverNetOutbox.SetInboxes(txOpts, chainIDs(routes), inboxConfigs(routes))
	if err != nil {
		return errors.Wrap(err, "set inboxes", "chain", backend.Name())
	}

	receipt, err := backend.WaitMined(ctx, tx)
	if err != nil {
		return errors.Wrap(err, "wait mined", "chain", backend.Name())
	}

	log.Info(ctx, "SolverNetOutbox configured", "addr", outbox.Hex(), "chain", backend.Name(), "block", receipt.BlockNumber, "tx", maybeTxHash(receipt))

	return nil
}

// chainIDs returns the chain IDs of the routes.
func chainIDs(routes []Route) []uint64 {
	chainIDs := make([]uint64, 0, len(routes))
	for _, route := range routes {
		chainIDs = append(chainIDs, route.ChainID)
	}

	return chainIDs
}

// outboxes returns the outboxes of the routes.
func outboxes(routes []Route) []common.Address {
	outboxes := make([]common.Address, 0, len(routes))
	for _, route := range routes {
		outboxes = append(outboxes, route.InboxConfig)
	}

	return outboxes
}

// inboxConfigs returns the inbox configs of the routes.
func inboxConfigs(routes []Route) []bindings.ISolverNetOutboxInboxConfig {
	inboxConfigs := make([]bindings.ISolverNetOutboxInboxConfig, 0, len(routes))
	for _, route := range routes {
		inboxConfigs = append(inboxConfigs, route.OutboxConfig)
	}

	return inboxConfigs
}
