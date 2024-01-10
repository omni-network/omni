package provider

import (
	"context"
	"time"

	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/ethereum/go-ethereum/common"
)

const (
	SyncCheckInterval = 5 * time.Second // wait interval between sync checks
	BlockChannelSize  = 10              // not of xblocks can be in the channel at any given time
)

type Streamer struct {
	chainConfig     *ChainConfig       // the chain config which also has the subscription information
	rpcClient       *RPCClient         // the rpc client which manages the block production
	lastBlockHash   common.Hash        // the hash of the last block delivered (used for parent hash)
	lastBlockHeight uint64             // the height of the last block delivered
	blockC          chan *xchain.Block // the channel through which the block is transmitted from rpc to streamer
}

// NewStreamer manages the rpc client to collect xblocks and it delivers the blocks to the
// subscriber callback.
func NewStreamer(ctx context.Context, config *ChainConfig) (*Streamer, error) {
	blockChannel := make(chan *xchain.Block, BlockChannelSize)

	// create the rpc client and do few validations
	client, err := NewRPCClient(ctx, config, blockChannel)
	if err != nil {
		return nil, err
	}

	// initialize the streamer structure with the received configuration
	stream := &Streamer{
		chainConfig: config,
		rpcClient:   client,
		blockC:      blockChannel,
	}

	return stream, nil
}

// start triggers the rpc to produce blocks and streamer to consume and deliver blocks.
func (s *Streamer) start(ctx context.Context) {
	// wait for the chain to sync
	s.waitForChainToSync(ctx)

	s.consumeXBlock(ctx)
	go s.rpcClient.produceBlocks(ctx, s.chainConfig.fromHeight)
}

// stop kills the streaming of blocks and quits.
func (s *Streamer) stop(ctx context.Context) {
	ctx.Done()
	close(s.blockC)
}

// consumeXBlock is a forever loop (until interrupted by context) which collects the xblocks from
// the rpc client and delivers it to the.
func (s *Streamer) consumeXBlock(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case block := <-s.blockC:
			s.deliverBlock(ctx, block)
			s.lastBlockHeight = block.BlockHeight
			s.lastBlockHash = block.BlockHash
			noOfMsgs := len(block.Msgs)
			noOfReceipts := len(block.Receipts)
			log.Info(ctx, "Delivered xblock",
				"chain name", s.chainConfig.name,
				"chain id", s.chainConfig.id,
				"block height", block.BlockHeight,
				"block hash", block.BlockHash,
				"XMsgs count", noOfMsgs,
				"XReceipts count", noOfReceipts)
		}
	}
}

// deliverBlock does the actual delivery of xblocks. if there is any error in
// delivering the block, it re-tries until it succeeds.
func (s *Streamer) deliverBlock(ctx context.Context, block *xchain.Block) {
	err := s.chainConfig.callback(ctx, block)
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
				err = s.chainConfig.callback(ctx, block)
				if err != nil {
					log.Error(ctx, "Error while delivering xblock", err,
						"chain name", s.chainConfig.name,
						"chain id", s.chainConfig.id,
						"block height", block.BlockHeight,
						"block hash", block.BlockHash)
				} else {
					return // block successfully delivered
				}
			}
		}
	}
}

// waitForChainToSync is called during startup of the provider to check if the chain is
// in sync condition. if not, it waits until the chain catches up with the canonical tip
// of the chain.
func (s *Streamer) waitForChainToSync(ctx context.Context) {
	for {
		syncProgress, err := s.rpcClient.getSyncStatus(ctx)
		if syncProgress == nil && err == nil {
			log.Info(ctx, "Chain is in sync status")

			return
		}
		if syncProgress != nil {
			log.Info(ctx, "Syncing in progress", ""+
				"chainName", s.chainConfig.name,
				"chainId", s.chainConfig.id,
				"highest block", syncProgress.HighestBlock,
				"current block", syncProgress.CurrentBlock)
		}
		time.Sleep(SyncCheckInterval)
	}
}
