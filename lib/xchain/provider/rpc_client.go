package provider

import (
	"context"
	"fmt"
	"sync/atomic"
	"time"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	RPCConnectInterval = 10 * time.Second // rpc connection re-try interval
	BlockBatchSize     = 100              // the no of blocks to scan in an iteration
	BlockBatchInterval = 5 * time.Second  // time interval between every block iteration
)

var ErrPortalAddressNotContract = errors.New("portal address is not a contract")

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
	// dial the rpc client until it gets connected
	eClient := connectToURL(ctx, config.rpcURL)

	// check if we are connecting to the proper chain id
	networkID, err := eClient.NetworkID(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "error while getting network id from rpc")
	}
	if config.id != networkID.Uint64() {
		errString := fmt.Sprintf("chain id mismatch. expected %d, got %d ", config.id, networkID.Uint64())
		return nil, errors.New(errString)
	}

	// check if portal address is a valid address and is a smart contract
	bytecode, err := eClient.CodeAt(context.WithoutCancel(ctx), config.portalAddress, nil)
	if err != nil {
		return nil, errors.Wrap(err, "error while getting the portal code")
	}
	if len(bytecode) == 0 {
		return nil, ErrPortalAddressNotContract
	}

	return &RPCClient{
		config: config,
		client: eClient,
		blockC: blockChannel,
	}, nil
}

// produceBlocks inspects the event logs for XMsgs and XRceceipts and constructs the XBlock
// out of it. This runs forever, looking in to logs in batches and waits for soe interval
// in-between each collection operation.
func (r *RPCClient) produceBlocks(ctx context.Context, fromHeight uint64) {
	// produce blocks on every BlockBatchInterval
	ticker := time.NewTicker(BlockBatchInterval)
	defer ticker.Stop()

	startHeight := fromHeight
	var locker uint32 // variable to take channel backpressure
	for {
		select {
		case <-ctx.Done():
			log.Info(ctx, "Stopping to produce blocks",
				"chain name", r.config.name,
				"chain id", r.config.id)

			return
		case <-ticker.C:
			// if the previous xblocks are not consumed yet, then skip this interval
			if atomic.CompareAndSwapUint32(&locker, 0, 1) {
				// blocks := r.checkForBlockInBatch() // ctx, startHeight, BlockBatchSize)
				r.checkForBlockInBatch()
				// for _, block := range blocks {
				// 	r.blockC <- block
				// }
				startHeight += BlockBatchSize
				atomic.StoreUint32(&locker, 0) // release the lock to accept new xblocks
			}
		}
	}
}

// connectToURL dials to the given rpc url. upon encountering any connection errors it
// backs off for sometime and re-try until the connection is successful.
func connectToURL(ctx context.Context, rpcURL string) *ethclient.Client {
	connCount := 0
	for {
		// TODO(jmozah): may be we should quit the subscription after few attempts
		eClient, err := ethclient.Dial(rpcURL)
		if err != nil {
			connCount++
			log.Warn(ctx, "Could not connect to chain. will again try after sometime")
			time.Sleep(RPCConnectInterval)

			continue
		}
		log.Info(ctx, "Connected to chain", "rpcURL", rpcURL)

		return eClient
	}
}

// checkForBlockInBatch this function looks for any XMsgs or XReceipts log event between
// a given block range. If present, it collects them and creates XBlocks with respective
// additional information, like streamId, block header etc.
func (r *RPCClient) checkForBlockInBatch() { // []*xchain.Block {
	// TODO(jmozah): get every log event from the start and end block and construct XBlocks
	//
	r.client.Client()

	// return nil
}

// getSyncStatus retrieves the syncing status of the rollup chain.
func (r *RPCClient) getSyncStatus(ctx context.Context) (*ethereum.SyncProgress, error) {
	syncProgress, err := r.client.SyncProgress(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "error while calling sync progress")
	}

	return syncProgress, nil
}
