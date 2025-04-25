package solve

import (
	"context"

	"github.com/omni-network/omni/lib/contracts/solvernet"
	"github.com/omni-network/omni/lib/contracts/solvernet/executor"
	"github.com/omni-network/omni/lib/contracts/solvernet/inbox"
	"github.com/omni-network/omni/lib/contracts/solvernet/middleman"
	"github.com/omni-network/omni/lib/contracts/solvernet/outbox"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"

	"golang.org/x/sync/errgroup"
)

// Deploy deploys solve inbox / outbox / middleman contracts, and devnet app (if devnet).
func Deploy(ctx context.Context, network netconf.Network, backends ethbackend.Backends) error {
	network = solvernet.AddHLNetwork(network)

	var eg1 errgroup.Group
	eg1.Go(func() error { return deployBoxes(ctx, network, backends) })
	eg1.Go(func() error { return maybeDeployMockTokens(ctx, network, backends) })

	if err := eg1.Wait(); err != nil {
		return errors.Wrap(err, "deploy prerequisites")
	}

	// These routines need to wait for `maybeDeployMockTokens` because they might use mock tokens.
	var eg2 errgroup.Group
	eg2.Go(func() error { return maybeFundERC20Solver(ctx, network.ID, backends) })
	eg2.Go(func() error { return maybeFundERC20Flowgen(ctx, network.ID, backends) })
	eg2.Go(func() error { return maybeDeployMockVault(ctx, network, backends) })

	if err := eg2.Wait(); err != nil {
		return errors.Wrap(err, "deploy dependent tasks")
	}

	return nil
}

func deployBoxes(ctx context.Context, network netconf.Network, backends ethbackend.Backends) error {
	var eg errgroup.Group
	for _, chain := range network.EVMChains() {
		backend, err := backends.Backend(chain.ID)
		if err != nil {
			return errors.Wrap(err, "get backend", "chain", chain.Name)
		}

		eg.Go(func() error {
			addr, receipt, err := inbox.Deploy(ctx, network, backend)
			if err != nil {
				return errors.Wrap(err, "deploy inbox", "chain", chain.Name)
			}

			log.Info(ctx, "SolverNetInbox deployed", "addr", addr.Hex(), "chain", chain.Name, "tx", maybeTxHash(receipt))

			return nil
		})

		eg.Go(func() error {
			addr, receipt, err := outbox.Deploy(ctx, network, backend)
			if err != nil {
				return errors.Wrap(err, "deploy outbox", "chain", chain.Name)
			}

			log.Info(ctx, "SolverNetOutbox deployed", "addr", addr.Hex(), "chain", chain.Name, "tx", maybeTxHash(receipt))

			return nil
		})

		eg.Go(func() error {
			addr, receipt, err := middleman.Deploy(ctx, network, backend)
			if err != nil {
				return errors.Wrap(err, "deploy middleman", "chain", chain.Name)
			}

			log.Info(ctx, "SolverNetMiddleman deployed", "addr", addr.Hex(), "chain", chain.Name, "tx", maybeTxHash(receipt))

			return nil
		})

		eg.Go(func() error {
			addr, receipt, err := executor.Deploy(ctx, network, backend)
			if err != nil {
				return errors.Wrap(err, "deploy executor", "chain", chain.Name)
			}

			log.Info(ctx, "SolverNetExecutor deployed", "addr", addr.Hex(), "chain", chain.Name, "tx", maybeTxHash(receipt))

			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		return errors.Wrap(err, "deploy solvernet boxes")
	}

	return nil
}

func maybeTxHash(receipt *ethclient.Receipt) string {
	if receipt != nil {
		return receipt.TxHash.Hex()
	}

	return "nil"
}
