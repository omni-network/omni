package solve

import (
	"context"
	"time"

	"github.com/omni-network/omni/lib/contracts/solvernet/inbox"
	"github.com/omni-network/omni/lib/contracts/solvernet/middleman"
	"github.com/omni-network/omni/lib/contracts/solvernet/outbox"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/ethereum/go-ethereum/common"
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

	network = maybeAddHolesky(network)

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

// maybeAddHolesky adds holesky to omega network. temporary fix until we re-enable holesky in netconf.
func maybeAddHolesky(network netconf.Network) netconf.Network {
	if network.ID != netconf.Omega {
		return network
	}

	// if holesky already exists, return
	for _, chain := range network.Chains {
		if chain.ID == evmchain.IDHolesky {
			return network
		}
	}

	// from omega netconf static
	deployHeight := 2130892
	portalAddr := common.HexToAddress("0xcB60A0451831E4865bC49f41F9C67665Fc9b75C3")

	// from e2e/types
	shards := []xchain.ShardID{xchain.ShardFinalized0, xchain.ShardLatest0}

	meta, ok := evmchain.MetadataByID(evmchain.IDHolesky)
	if !ok {
		// will not happen
		return network
	}

	network.Chains = append(network.Chains, netconf.Chain{
		ID:             evmchain.IDHolesky,
		Name:           meta.Name,
		PortalAddress:  portalAddr,
		DeployHeight:   uint64(deployHeight),
		BlockPeriod:    meta.BlockPeriod,
		Shards:         shards,
		AttestInterval: intervalFromPeriod(network.ID, meta.BlockPeriod),
	})

	return network
}

// from e2e/types testnet.go (temporary).
func intervalFromPeriod(network netconf.ID, period time.Duration) uint64 {
	target := time.Hour
	if network == netconf.Staging {
		target = time.Minute * 10
	} else if network == netconf.Devnet {
		target = time.Second * 10
	}

	if period == 0 {
		return 0
	}

	return uint64(target / period)
}
