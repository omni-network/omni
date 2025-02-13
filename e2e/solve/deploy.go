package solve

import (
	"context"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/halo/genutil/evm/predeploys"
	"github.com/omni-network/omni/lib/contracts/solvernet/inbox"
	"github.com/omni-network/omni/lib/contracts/solvernet/middleman"
	"github.com/omni-network/omni/lib/contracts/solvernet/outbox"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"

	"golang.org/x/sync/errgroup"
)

// Deploy deploys solve inbox / outbox / middleman contracts, and devnet app (if devnet).
func Deploy(ctx context.Context, networkID netconf.ID, backends ethbackend.Backends) error {
	var eg errgroup.Group

	portalReg, err := makePortalRegistry(networkID, backends)
	if err != nil {
		return err
	}

	network, err := netconf.AwaitOnExecutionChain(ctx, networkID, portalReg, chainNames(networkID, backends))
	if err != nil {
		return err
	}

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

			addr, receipt, err := inbox.Deploy(ctx, network, backend)
			if err != nil {
				return errors.Wrap(err, "deploy solve inbox", "chain", chain.Name)
			}

			log.Info(ctx, "SolverNetInbox deployed", "addr", addr.Hex(), "chain", chain.Name, "tx", maybeTxHash(receipt))

			addr, receipt, err = outbox.Deploy(ctx, network, backend)
			if err != nil {
				return errors.Wrap(err, "deploy solve outbox", "chain", chain.Name)
			}

			log.Info(ctx, "SolverNetOutbox deployed", "addr", addr.Hex(), "chain", chain.Name, "tx", maybeTxHash(receipt))

			addr, receipt, err = middleman.Deploy(ctx, network, backend)
			if err != nil {
				return errors.Wrap(err, "deploy solve middleman", "chain", chain.Name)
			}

			log.Info(ctx, "SolverNetMiddleman deployed", "addr", addr.Hex(), "chain", chain.Name, "tx", maybeTxHash(receipt))

			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		return errors.Wrap(err, "deploy solvernet boxes")
	}

	return nil
}

func maybeTxHash(receipt *ethtypes.Receipt) string {
	if receipt != nil {
		return receipt.TxHash.Hex()
	}

	return "nil"
}

func chainNames(networkID netconf.ID, backends ethbackend.Backends) []string {
	var names []string
	namer := netconf.ChainNamer(networkID)

	for chainID := range backends.All() {
		names = append(names, namer(chainID))
	}

	return names
}

func makePortalRegistry(network netconf.ID, backends ethbackend.Backends) (*bindings.PortalRegistry, error) {
	meta := netconf.MetadataByID(network, network.Static().OmniExecutionChainID)
	backend, err := backends.Backend(meta.ChainID)
	if err != nil {
		return nil, errors.Wrap(err, "backend", "chain", meta.Name)
	}

	resp, err := bindings.NewPortalRegistry(common.HexToAddress(predeploys.PortalRegistry), backend)
	if err != nil {
		return nil, errors.Wrap(err, "create portal registry")
	}

	return resp, nil
}
