package provider

import (
	"context"
	"time"

	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/ethereum/go-ethereum/common"
)

const (
	BlockChannelSize = 10 // not of xblocks can be in the channel at any given time
)

type Streamer struct {
	chainConfig     *ChainConfig            // the chain config which also has the subscription information
	minHeight       uint64                  // the minimum height to start the getting the xchain messages
	callback        xchain.ProviderCallback // the callback to call on receiving a xblock
	rpcClient       *RPCClient              // the rpc client which manages the block production
	lastBlockHash   common.Hash             // the hash of the last block delivered (used for parent hash)
	lastBlockHeight uint64                  // the height of the last block delivered
	blockC          chan *xchain.Block      // the channel through which the block is transmitted from rpc to streamer
	quitC           chan bool               // the quit channel to stop all the streamer operations
}

// NewStreamer manages the rpc client to collect xblocks and delivers it to the
// subscriber through callback.
func NewStreamer(ctx context.Context,
	config *ChainConfig,
	minHeight uint64,
	callback xchain.ProviderCallback,
	quitC chan bool,
) (*Streamer, error) {
	blockChannel := make(chan *xchain.Block, BlockChannelSize)

	// create the rpc client and do few validations
	client, err := NewRPCClient(ctx, config, blockChannel)
	if err != nil {
		return nil, err
	}

	// initialize the streamer structure with the received configuration
	stream := &Streamer{
		chainConfig: config,
		minHeight:   minHeight,
		callback:    callback,
		rpcClient:   client,
		blockC:      blockChannel,
		quitC:       quitC,
	}

	return stream, nil
}

// start triggers the rpc to produce blocks and streamer to consume and deliver blocks.
func (s *Streamer) start(ctx context.Context) {
	// TODO(jmozah) wait for the chain to sync

	go s.rpcClient.produceBlocks(ctx, s.minHeight, s.quitC)
	s.consumeXBlock(ctx)
}

// consumeXBlock is a forever loop (until interrupted by context) which collects the xblocks from
// the rpc client and delivers it to the.
func (s *Streamer) consumeXBlock(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			log.Info(ctx, "Stopping to consume blocks",
				"chainName", s.chainConfig.name,
				"chainID", s.chainConfig.id)

			return
		case <-s.quitC:
			close(s.blockC)
			log.Info(ctx, "Stopping to consume blocks",
				"chainName", s.chainConfig.name,
				"chainID", s.chainConfig.id)

			return
		case block := <-s.blockC:
			s.deliverBlock(ctx, block)
			s.lastBlockHeight = block.BlockHeight
			s.lastBlockHash = block.BlockHash
			noOfMsgs := len(block.Msgs)
			noOfReceipts := len(block.Receipts)
			log.Info(ctx, "Delivered xblock",
				"chainName", s.chainConfig.name,
				"chainID", s.chainConfig.id,
				"blockHeight", block.BlockHeight,
				"blockHash", block.BlockHash,
				"xMsgsCount", noOfMsgs,
				"xReceiptsCount", noOfReceipts)
		}
	}
}

// deliverBlock does the actual delivery of xblocks. if there is any error in
// delivering the block, it re-tries until it succeeds.
func (s *Streamer) deliverBlock(ctx context.Context, block *xchain.Block) {
	err := s.callback(ctx, block)
	if err != nil {
		// if there is some error in delivering block
		// then try delivering after sometime
		// TODO(jmozah): may be we should quit the subscription after few attempts
		ticker := time.NewTicker(BlockBatchInterval)
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				err = s.callback(ctx, block)
				if err != nil {
					log.Error(ctx, "Error while delivering xblock", err,
						"chainName", s.chainConfig.name,
						"chainID", s.chainConfig.id,
						"blockHeight", block.BlockHeight,
						"blockHash", block.BlockHash)
				} else {
					return // block successfully delivered
				}
			}
		}
	}
}
