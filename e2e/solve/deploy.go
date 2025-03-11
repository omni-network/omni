package solve

import (
	"context"

	"github.com/omni-network/omni/lib/contracts/solvernet/executor"
	"github.com/omni-network/omni/lib/contracts/solvernet/inbox"
	"github.com/omni-network/omni/lib/contracts/solvernet/middleman"
	"github.com/omni-network/omni/lib/contracts/solvernet/outbox"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"

	ethtypes "github.com/ethereum/go-ethereum/core/types"

	"golang.org/x/sync/errgroup"
)

// Deploy deploys solve inbox / outbox / middleman contracts, and devnet app (if devnet).
func Deploy(ctx context.Context, network netconf.Network, backends ethbackend.Backends) error {
	var eg errgroup.Group
	eg.Go(func() error { return deployBoxes(ctx, network, backends) })
	eg.Go(func() error { return maybeDeployMockTokens(ctx, network, backends) })
	eg.Go(func() error { return maybeFundERC20Solver(ctx, network.ID, backends) })

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
				return errors.Wrap(err, "deploy inbox", "chain", chain.Name)
			}

			log.Info(ctx, "SolverNetInbox deployed", "addr", addr.Hex(), "chain", chain.Name, "tx", maybeTxHash(receipt))

			addr, receipt, err = outbox.Deploy(ctx, network, backend)
			if err != nil {
				return errors.Wrap(err, "deploy outbox", "chain", chain.Name)
			}

			log.Info(ctx, "SolverNetOutbox deployed", "addr", addr.Hex(), "chain", chain.Name, "tx", maybeTxHash(receipt))

			addr, receipt, err = middleman.Deploy(ctx, network, backend)
			if err != nil {
				return errors.Wrap(err, "deploy middleman", "chain", chain.Name)
			}

			log.Info(ctx, "SolverNetMiddleman deployed", "addr", addr.Hex(), "chain", chain.Name, "tx", maybeTxHash(receipt))

			addr, receipt, err = executor.Deploy(ctx, network, backend)
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

func maybeTxHash(receipt *ethtypes.Receipt) string {
	if receipt != nil {
		return receipt.TxHash.Hex()
	}

	return "nil"
}
