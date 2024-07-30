package app

import (
	"context"
	"time"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/log"

	rpcclient "github.com/cometbft/cometbft/rpc/client"
)

// monitorCometForever blocks until the context is canceled.
// It periodically calls monitorCometOnce.
func monitorCometForever(ctx context.Context, rpcClient rpcclient.Client, isSyncing func() bool) {
	ticker := time.NewTicker(time.Second * 30)
	defer ticker.Stop()

	// Run initial monitoring immediately.
	lastHeight, _ := monitorCometOnce(ctx, rpcClient, isSyncing, 0)

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			height, err := monitorCometOnce(ctx, rpcClient, isSyncing, lastHeight)
			if err != nil {
				log.Warn(ctx, "Failed monitoring cometBFT (will retry)", err)
				// Don't reset lastHeight to zero.
			} else {
				lastHeight = height
			}
		}
	}
}

// monitorCometOnce monitors the cometBFT peers, and sync status.
func monitorCometOnce(ctx context.Context, rpcClient rpcclient.Client, isSyncing func() bool, lastHeight int64) (int64, error) {
	if netInfo, err := rpcClient.NetInfo(ctx); err != nil {
		return 0, errors.Wrap(err, "net info")
	} else if netInfo.NPeers == 0 {
		log.Error(ctx, "Halo has 0 consensus p2p peers", nil)
	}

	synced := 1
	if isSyncing() {
		synced = 0
		log.Warn(ctx, "Halo is syncing", nil)
	}
	cometSynced.Set(float64(synced))

	abciInfo, err := rpcClient.ABCIInfo(ctx)
	if err != nil {
		return 0, errors.Wrap(err, "abci info")
	} else if !isSyncing() && abciInfo.Response.LastBlockHeight <= lastHeight {
		log.Warn(ctx, "Halo height is not increasing, evm syncing?", nil)
	}

	return abciInfo.Response.LastBlockHeight, nil
}

// monitorEVMForever blocks until the contract is canceled.
// It periodically calls monitorEVMOnce.
func monitorEVMForever(ctx context.Context, cfg Config, ethCl ethclient.Client) {
	ticker := time.NewTicker(time.Second * 30)
	defer ticker.Stop()

	// Geth Auth API (EngineClient) doesn't enable net module, so we can't monitor peer count with it.
	// If a HTTP API also configured in RPCEndpoints, use it instead.
	omniEVM := cfg.Network.Static().OmniExecutionChainName()
	omniEVMRPC, err := cfg.RPCEndpoints.ByNameOrID(omniEVM, cfg.Network.Static().OmniExecutionChainID)
	if err == nil {
		newEthCl, err := ethclient.Dial(omniEVM, omniEVMRPC)
		if err == nil {
			ethCl = newEthCl
			log.Info(ctx, "Using rpc endpoint to monitor attached omni evm", "rpc", omniEVMRPC)
		}
	}

	// Run initial monitoring immediately.
	_ = monitorEVMOnce(ctx, ethCl)

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
		log.Error(ctx, "Attached omni evm has 0 peers", nil)
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
