package provider

import (
	"context"
	"sync/atomic"
	"time"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	BlockNumberOfRetry      = 10              // no fo times to retry when fetching a block
	BlockRetrySleepInterval = 1 * time.Second // time interval between retry
	BlockBatchInterval      = 1 * time.Second // time interval between every block iteration
)

// RPCClient stores the chain config and the rpc client.
type RPCClient struct {
	config *ChainConfig       // the rollup chain configuration
	client *ethclient.Client  // rpc client for the chain
	blockC chan *xchain.Block // channel to which the collected xblocks are sent for processing
}

// NewRPCClient is the client implementation to connect to the rollup chain and collect XBlocks.
func NewRPCClient(
	ctx context.Context,
	config *ChainConfig,
	blockChannel chan *xchain.Block,
) (*RPCClient, error) {
	// dial the rpc client and try to get connected
	eClient, err := ethclient.Dial(config.rpcURL)
	if err != nil {
		log.Warn(ctx, "Could not connect to chain. will again try after sometime", err,
			"chainName", config.name,
			"chainId", config.id,
			"rpcURL", config.rpcURL)
	}
	log.Info(ctx, "Connected to chain. ",
		"chainName", config.name,
		"chainId", config.id,
		"rpcURL", config.rpcURL)

	// TODO(jmozah): validations like checking networkId, portal Address etc
	return &RPCClient{
		config: config,
		client: eClient,
		blockC: blockChannel,
	}, nil
}

// produceBlocks inspects the event logs for XMsgs and XRceceipts and constructs the XBlock
// out of it. This runs forever, looking in to logs in batches and waits for soe interval
// in-between each collection operation.
func (r *RPCClient) produceBlocks(ctx context.Context, minHeight uint64, quitC chan bool) {
	// produce blocks on every BlockBatchInterval
	ticker := time.NewTicker(BlockBatchInterval)
	defer ticker.Stop()

	var locker uint32 // variable to take channel backpressure
	for {
		select {
		case <-ctx.Done():
			log.Info(ctx, "Stopping to produce blocks",
				"chainName", r.config.name,
				"chainID", r.config.id)

			return

		case <-quitC:
			log.Info(ctx, "Stopping to produce blocks",
				"chainName", r.config.name,
				"chainID", r.config.id)

			return

		case <-ticker.C:
			// if the previous xblocks are not consumed yet, then skip this interval
			if !atomic.CompareAndSwapUint32(&locker, 0, 1) {
				continue
			}

			block, err := r.checkForBlockInHeight(ctx) // TODO(jmozah): add minHeight, quitC later
			if err != nil {
				continue // if some error in getting the block
			}

			// send the block for processing if it is present
			if block != nil {
				r.blockC <- block
			}

			// go to the next block
			minHeight++

			// release the lock to accept new xblocks
			atomic.StoreUint32(&locker, 0)
		}
	}
}

// checkForBlockInBatch this function looks for any XMsgs or XReceipts log event between
// a given block range. If present, it collects them and creates XBlocks with respective
// additional information, like streamId, block header etc.
func (r *RPCClient) checkForBlockInHeight(ctx context.Context) (*xchain.Block, error) {
	noOfRetry := 0
	for {
		block, err := r.getBlockFromLog(ctx) // TODO(jmozah): add height, quitC later
		if err != nil {
			if noOfRetry == BlockNumberOfRetry {
				return nil, err
			}
			noOfRetry++
			time.Sleep(BlockRetrySleepInterval)

			continue
		}

		return block, nil
	}
}

// getBlockFromLog connects to the chain and fetches the block.
func (r *RPCClient) getBlockFromLog(ctx context.Context) (*xchain.Block, error) {
	// check for connection
	if r.client == nil {
		eClient, err := ethclient.Dial(r.config.rpcURL)
		if err != nil {
			log.Info(ctx, "Lost connection to chain. Trying to connect again",
				"chainName", r.config.name,
				"chainId", r.config.id,
				"rpcURL", r.config.rpcURL)
			r.client = nil

			return nil, errors.Wrap(err, "lost client connection")
		}
		log.Info(ctx, "Re-connected to chain",
			"chainName", r.config.name,
			"chainId", r.config.id,
			"rpcURL", r.config.rpcURL)
		r.client = eClient
	}

	var block *xchain.Block

	return block, nil
}
