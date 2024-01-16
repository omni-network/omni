package provider

import (
	"context"
	"time"

	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/ethereum/go-ethereum/core/types"
)

const (
	BlockFetchInterval = 1 * time.Second // time interval between each block fetch
)

// Streamer maintains the config and the destination for each chain.
type Streamer struct {
	chainConfig *ChainConfig            // the chain config which also has the subscription information
	minHeight   uint64                  // the minimum height to start the getting the xchain messages
	callback    xchain.ProviderCallback // the callback to call on receiving a xblock
	batchSize   uint64
	backoffFunc func(context.Context) (func(), func())
}

// NewStreamer manages the rpc client to collect xblocks and delivers it to the
// subscriber through callback.
func NewStreamer(config *ChainConfig,
	minHeight uint64,
	callback xchain.ProviderCallback,
	batchSize uint64,
	backoffFunc func(context.Context) (func(), func()),
) *Streamer {
	// initialize the streamer structure with the received configuration
	stream := &Streamer{
		chainConfig: config,
		minHeight:   minHeight,
		callback:    callback,
		batchSize:   batchSize,
		backoffFunc: backoffFunc,
	}

	return stream
}

// streamBlocks triggers a continuously running routine that fetches and delivers xBlocks.
func (s *Streamer) streamBlocks(ctx context.Context, height uint64) {
	go func() {
		backoff, reset := s.backoffFunc(ctx)
		// stream blocks until the context is canceled
		for ctx.Err() == nil {
			// get the current finalized block
			finalisedHeader, noOfBlocksToDeliver, err := s.getNumberOfBlockToDeliver(ctx, height)
			if ctx.Err() != nil {
				return
			}
			if err != nil || noOfBlocksToDeliver == 0 {
				backoff()

				continue
			}
			reset() // reset the fetching of finalized blocks
			log.Debug(ctx, "No of blocks to deliver",
				"start_height", height,
				"no_of_blocks", noOfBlocksToDeliver)

			// if height is not in the tip, fetch and deliver blocks in batches
			// this will be useful in cases we have a lot of blocks to catchup
			for currentHeight := height; currentHeight < noOfBlocksToDeliver; currentHeight++ {
				log.Debug(ctx, "Fetching block", "height", currentHeight)

				// fetch xBlock
				xBlock, exists := s.fetchXBlock(ctx, currentHeight, finalisedHeader, backoff, reset)

				// no cross chain logs in this height, so go to the next height
				if !exists {
					log.Debug(ctx, "No cross chain block", "height", currentHeight)

					continue
				}

				// deliver the fetched xBlock
				s.deliverXBlock(ctx, currentHeight, xBlock, backoff, reset)
			}
		}
	}()
}

func (s *Streamer) fetchXBlock(ctx context.Context,
	currentHeight uint64,
	finalisedHeader *types.Header,
	backoff func(),
	reset func(),
) (xchain.Block, bool) {
	var hdr *types.Header
	if finalisedHeader.Number.Uint64() == currentHeight {
		hdr = finalisedHeader
	}

	// fetch xBlock
	var blk xchain.Block
	for ctx.Err() == nil {
		// get the message and receipts from the chain for this block if any

		xBlock, exists, err := s.chainConfig.rpcClient.GetBlock(ctx, currentHeight, hdr)
		if ctx.Err() != nil {
			return xBlock, false
		}
		if err != nil {
			log.Warn(ctx, "Could not get cross chain block", err, "height", currentHeight)
			backoff() // backoff and retry fetching the block

			continue
		}
		reset() // reset the GetBlock backoff

		return xBlock, exists
	}

	return blk, false
}

func (s *Streamer) deliverXBlock(ctx context.Context,
	currentHeight uint64,
	xBlock xchain.Block,
	backoff func(),
	reset func(),
) {
	// deliver the fetched xBlock
	for ctx.Err() == nil {
		callbackErr := s.callback(ctx, &xBlock)
		if ctx.Err() != nil {
			return
		}
		if callbackErr != nil {
			log.Warn(ctx, "Error while delivering xBlock", callbackErr,
				"height", currentHeight,
				"hash", xBlock.BlockHash)
			backoff() // try delivering after sometime

			continue
		}
		reset() // delivery backoff reset

		break // successfully delivered the xBlock
	}

	log.Debug(ctx, "Delivered xBlock", "height", currentHeight)
}

// getNumberOfBlockToDeliver finds the number of blocks to.
func (s *Streamer) getNumberOfBlockToDeliver(ctx context.Context, height uint64) (*types.Header, uint64, error) {
	// get the current finalized block
	finalisedHeader, err := s.chainConfig.rpcClient.GetCurrentFinalisedBlockHeader(ctx)
	if err != nil {
		return nil, 0, err
	}

	// calculate the number of blocks to fetch in this batch
	var blocksToFetch uint64
	if height < finalisedHeader.Number.Uint64() {
		blocksToFetch = finalisedHeader.Number.Uint64() - height
		if blocksToFetch > s.batchSize {
			blocksToFetch = s.batchSize
		}
	} else if height == finalisedHeader.Number.Uint64() {
		blocksToFetch = 1 // we are already in tip
	} else {
		log.Debug(ctx, "Height is greater than the finalized tip",
			"height", height,
			"finalized_height", finalisedHeader.Number.Uint64())

		return nil, 0, nil
	}

	return finalisedHeader, blocksToFetch, nil
}
