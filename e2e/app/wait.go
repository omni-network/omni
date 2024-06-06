package app

import (
	"context"
	"time"

	"github.com/omni-network/omni/e2e/netman"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/expbackoff"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"

	e2e "github.com/cometbft/cometbft/test/e2e/pkg"
)

// Wait waits for a number of blocks to be produced, and for all nodes to catch
// up with it.
func Wait(ctx context.Context, testnet *e2e.Testnet, blocks int64) error {
	block, _, err := waitForHeight(ctx, testnet, 0)
	if err != nil {
		return err
	}

	return WaitUntil(ctx, testnet, block.Height+blocks)
}

// WaitUntil waits until a given height has been reached.
func WaitUntil(ctx context.Context, testnet *e2e.Testnet, height int64) error {
	log.Info(ctx, "Waiting for nodes to reach height", "height", height)
	_, err := waitForAllNodes(ctx, testnet, height, waitingTime(len(testnet.Nodes), height))
	if err != nil {
		return err
	}

	return nil
}

// waitingTime estimates how long it should take for a node to reach the height.
// More nodes in a network implies we may expect a slower network and may have to wait longer.
func waitingTime(nodes int, height int64) time.Duration {
	return time.Duration(20+(int64(nodes)*height)) * time.Second
}

func WaitAllSubmissions(ctx context.Context, network netconf.Network, portals map[uint64]netman.Portal, minimum uint64) error {
	log.Info(ctx, "Waiting for submissions on all destination chains")
	ctx, cancel := context.WithTimeout(ctx, 5*time.Minute)
	defer cancel()

	for _, dest := range portals {
		for _, stream := range network.StreamsTo(dest.Chain.ChainID) {
			if stream.SourceChainID == network.ID.Static().OmniConsensusChainIDUint64() {
				continue
			}

			backoff := expbackoff.New(ctx, expbackoff.WithPeriodicConfig(time.Second))
			for i := 0; ; i++ {
				if ctx.Err() != nil {
					return errors.Wrap(ctx.Err(), "timeout waiting for submissions")
				}

				srcOffset, err := portals[stream.SourceChainID].Contract.OutXMsgOffset(nil, dest.Chain.ChainID, uint64(stream.ShardID))
				if err != nil {
					return errors.Wrap(err, "get outXMsgOffset")
				}

				destOffset, err := dest.Contract.InXMsgOffset(nil, stream.SourceChainID, uint64(stream.ShardID))
				if err != nil {
					return errors.Wrap(err, "get inXmsgOffset")
				}

				if srcOffset >= minimum && destOffset == srcOffset {
					break
				}

				if i%5 == 0 { // Only log every 5th iteration (5s)
					log.Debug(ctx, "Waiting for submissions on destination chain",
						"stream", network.StreamName(stream),
						"src_offset", srcOffset,
						"dest_offset", destOffset,
						"minimum", minimum,
					)
				}

				backoff()
			}
		}
	}

	return nil
}
