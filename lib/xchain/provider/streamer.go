package provider

import (
	"context"
	"sync/atomic"
	"time"

	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/xchain"
)

const (
	BlockFetchInterval = 1 * time.Second // time interval between each block fetch
)

// Streamer maintains the config and the destination for each chain.
type Streamer struct {
	chainConfig *ChainConfig            // the chain config which also has the subscription information
	minHeight   uint64                  // the minimum height to start the getting the xchain messages
	callback    xchain.ProviderCallback // the callback to call on receiving a xblock
	quitC       chan struct{}           // the quit channel to stop all the streamer operations
}

// NewStreamer manages the rpc client to collect xblocks and delivers it to the
// subscriber through callback.
func NewStreamer(config *ChainConfig,
	minHeight uint64,
	callback xchain.ProviderCallback,
	quitC chan struct{},
) (*Streamer, error) {
	// initialize the streamer structure with the received configuration
	stream := &Streamer{
		chainConfig: config,
		minHeight:   minHeight,
		callback:    callback,
		quitC:       quitC,
	}

	return stream, nil
}

func (s *Streamer) streamBlocks(ctx context.Context, currentHeight uint64) {
	// produce blocks on every BlockFetchInterval
	ticker := time.NewTicker(BlockFetchInterval)
	defer ticker.Stop()

	var locker uint32 // variable to take channel backpressure
	for {
		select {
		case <-ctx.Done():
			log.Info(ctx, "Stopping to produce blocks",
				"chainName", s.chainConfig.name,
				"chainID", s.chainConfig.id,
				"height", currentHeight)

			return

		case <-s.quitC:
			log.Info(ctx, "Stopping to produce blocks",
				"chainName", s.chainConfig.name,
				"chainID", s.chainConfig.id,
				"height", currentHeight)

			return

		case <-ticker.C:
			// if the previous xblocks are not consumed yet, then skip this interval
			if !atomic.CompareAndSwapUint32(&locker, 0, 1) {
				continue
			}

			// get the message and receipts from the chain for this block if any
			xBlock, err := s.chainConfig.rpcClient.GetBlock(ctx, currentHeight)
			if err != nil {
				log.Error(ctx, "Could not get cross chain block from rpc client", err,
					"chainName", s.chainConfig.name,
					"chainID", s.chainConfig.id,
					"height", currentHeight)

				continue
			}

			// ignore of there is no messages in this block
			if xBlock == nil {
				log.Info(ctx, "No cross chain block in this height",
					"chainName", s.chainConfig.name,
					"chainID", s.chainConfig.id,
					"height", currentHeight)

				continue
			}

			// deliver the block
			callbackErr := s.callback(ctx, xBlock)
			if callbackErr != nil {
				log.Error(ctx, "Error while delivering xblock", callbackErr,
					"chainName", s.chainConfig.name,
					"chainID", s.chainConfig.id,
					"blockHeight", currentHeight)

				continue
			}
			log.Info(ctx, "Delivered xBlock",
				"sourceChainID", xBlock.SourceChainID,
				"blockHeight", xBlock.BlockHeight,
				"blockHash", xBlock.BlockHash,
				"noOfMsgs", len(xBlock.Msgs),
				"noOfReceipts", len(xBlock.Receipts))

			// move to the next block
			currentHeight++

			// release the lock to accept new xblocks
			atomic.StoreUint32(&locker, 0)
		}
	}
}
