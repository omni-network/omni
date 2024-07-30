package app

import (
	"context"
	"time"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/log"
)

// monitorEVMForever blocks until the contract is canceled.
// It periodically calls monitorEVMOnce.
func monitorEVMForever(ctx context.Context, cfg Config, ethCl ethclient.Client) {
	ticker := time.NewTicker(time.Second * 30)
	defer ticker.Stop()

	// Geth Auth API (EngineClient) doesn't enable net module, so we can't monitor peer count with it.
	// If a HTTP API also configured in RPCEndpoints, use it instead.
	const omniEVM = "omni_evm"
	omniEVMRPC, err := cfg.RPCEndpoints.ByNameOrID(omniEVM, cfg.Network.Static().OmniExecutionChainID)
	if err == nil {
		newEthCl, err := ethclient.Dial(omniEVM, omniEVMRPC)
		if err == nil {
			ethCl = newEthCl
			log.Info(ctx, "Using rpc endpoint to monitor attached omni evm", "rpc", omniEVMRPC)
		}
	}

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			err := monitorEVMOnce(ctx, ethCl)
			if err != nil {
				log.Warn(ctx, "Failed monitoring attached omni evm (will retry)", err, "addr", ethCl.Address())
			}
		}
	}
}

// monitorEVMOnce monitors the attached omni_evm height, peers, and sync status.
func monitorEVMOnce(ctx context.Context, ethCl ethclient.Client) error {
	// Best effort monitoring of peer count, since method not available in auth API.
	peers, err := ethCl.PeerCount(ctx)
	if ethclient.IsErrMethodNotAvailable(err) { //nolint:revive // Empty block skips error handling below.
		// Do not set the metric if the method is not available.x
	} else if err != nil {
		return errors.Wrap(err, "peer count")
	} else if peers == 0 {
		log.Warn(ctx, "Attached omni evm has 0 peers", nil)
		evmPeers.Set(0)
	} else {
		evmPeers.Set(float64(peers))
	}

	synced := 1
	if syncing, err := ethCl.SyncProgress(ctx); err != nil {
		return errors.Wrap(err, "sync progress")
	} else if syncing != nil && !syncing.Done() {
		// SyncProgress returns nil of not syncing.
		synced = 0
		log.Warn(ctx, "Attached omni evm is syncing", nil, "highest_block", syncing.HighestBlock, "current_block", syncing.CurrentBlock, "tx_indexing", syncing.TxIndexRemainingBlocks)
	}
	evmSynced.Set(float64(synced))

	latest, err := ethCl.BlockNumber(ctx)
	if err != nil {
		return errors.Wrap(err, "block number")
	}
	evmHeight.Set(float64(latest))

	return nil
}
