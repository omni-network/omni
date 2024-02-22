package app

import (
	"context"

	cprovider "github.com/omni-network/omni/lib/cchain/provider"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/test/e2e/netman"

	e2e "github.com/cometbft/cometbft/test/e2e/pkg"
)

func LogMetrics(ctx context.Context, def Definition) error {
	extNetwork := externalNetwork(def.Testnet, def.Netman.DeployInfo())

	// Pick a random node to monitor.
	if err := MonitorCProvider(ctx, random(def.Testnet.Nodes), extNetwork); err != nil {
		return errors.Wrap(err, "monitoring cchain provider")
	}

	if err := MonitorCursors(ctx, def.Netman.Portals(), extNetwork); err != nil {
		return errors.Wrap(err, "monitoring cursors")
	}

	return nil
}

func MonitorCursors(ctx context.Context, portals map[uint64]netman.Portal, network netconf.Network) error {
	for _, dest := range network.Chains {
		for _, src := range network.Chains {
			if src.ID == dest.ID {
				continue
			}

			srcOffset, err := portals[src.ID].Contract.OutXStreamOffset(nil, dest.ID)
			if err != nil {
				return errors.Wrap(err, "getting inXStreamOffset")
			}

			destOffset, err := portals[dest.ID].Contract.InXStreamOffset(nil, src.ID)
			if err != nil {
				return errors.Wrap(err, "getting inXStreamOffset")
			}

			log.Info(ctx, "Submitted cross chain messages",
				"src", src.Name,
				"dest", dest.Name,
				"total_in", destOffset,
				"total_out", srcOffset,
			)
		}
	}

	return nil
}

func MonitorCProvider(ctx context.Context, node *e2e.Node, network netconf.Network) error {
	client, err := node.Client()
	if err != nil {
		return errors.Wrap(err, "getting client")
	}

	cprov := cprovider.NewABCIProvider(client, network.ChainNamesByIDs())

	for _, chain := range network.Chains {
		atts, err := cprov.AttestationsFrom(ctx, chain.ID, chain.DeployHeight)
		if err != nil {
			return errors.Wrap(err, "getting approved attestations")
		}

		log.Info(ctx, "Halo approved attestations", "chain", chain.Name, "count", len(atts))
	}

	return nil
}
