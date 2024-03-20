package app

import (
	"context"
	"sync"

	"github.com/omni-network/omni/e2e/netman"
	cprovider "github.com/omni-network/omni/lib/cchain/provider"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/forkjoin"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"
	"github.com/omni-network/omni/lib/xchain/provider"

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

// StartMonitoringReceipts starts goroutines that streams all xblock receipts ensuring all are successful.
// It returns a stopfunc that returns an error if any failed receipt was detected before the stopfunc was called.
func StartMonitoringReceipts(ctx context.Context, def Definition) func() error {
	network := externalNetwork(def.Testnet, def.Netman.DeployInfo())
	xprovider := provider.New(network, def.Backends.RPCClients(), nil)

	type void any

	var msgCache sync.Map

	streamReceipts := func(ctx context.Context, chain netconf.Chain) (void, error) {
		return nil, xprovider.StreamBlocks(ctx, chain.ID, chain.DeployHeight,
			func(ctx context.Context, block xchain.Block) error {
				var failed int
				for _, receipt := range block.Receipts {
					if !receipt.Success {
						failed++
					}
				}
				for _, msg := range block.Msgs {
					msgCache.Store(msg.MsgID, msg)
				}
				for _, receipt := range block.Receipts {
					if receipt.Success {
						continue
					}

					// Log AND return so it is clearer in the logs when this happens.
					attrs := []any{
						"stream_offset", receipt.StreamOffset,
						"dest_chain", network.ChainName(receipt.DestChainID),
						"src_chain", network.ChainName(receipt.SourceChainID),
						"gas_used", receipt.GasUsed,
					}

					m, ok := msgCache.Load(receipt.MsgID)
					if !ok {
						log.Error(ctx, "Invalid receipt, missing msg", nil, attrs...)

						return errors.New("invalid receipt, missing msg", attrs...)
					}

					msg := m.(xchain.Msg) //nolint:forcetypeassert // We know it's a msg.
					attrs = append(attrs,
						"msg_address", msg.DestAddress.String(),
						"gas_limit", msg.DestGasLimit,
					)

					log.Error(ctx, "Detected failed receipt", nil, attrs...)

					return errors.New("receipt status failed", attrs...)
				}

				return nil
			})
	}

	results, cancel := forkjoin.NewWithInputs(ctx, streamReceipts, network.EVMChains())

	return func() error {
		log.Debug(ctx, "Checking receipts")
		cancel()
		for res := range results {
			if res.Err != nil {
				return errors.Wrap(res.Err, "streaming receipts", "chain", res.Input.Name)
			}
		}

		return nil
	}
}

func MonitorCursors(ctx context.Context, portals map[uint64]netman.Portal, network netconf.Network) error {
	for _, dest := range network.EVMChains() {
		for _, src := range network.EVMChains() {
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

			log.Debug(ctx, "Submitted cross chain messages",
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

		log.Debug(ctx, "Halo approved attestations", "chain", chain.Name, "count", len(atts))
	}

	return nil
}
