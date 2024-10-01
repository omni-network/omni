package app

import (
	"context"
	"os"
	"path/filepath"
	"time"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"

	rpcclient "github.com/cometbft/cometbft/rpc/client"
)

// monitorCometForever blocks until the context is canceled.
// It periodically calls monitorCometOnce.
func monitorCometForever(
	ctx context.Context,
	network netconf.ID,
	rpcClient rpcclient.Client,
	isSyncing func() bool,
	dbDir string,
	readiness *ReadyResponse,
) {
	if network == netconf.Simnet {
		return // Simnet doesn't need to monitor cometBFT, since no p2p.
	}

	ticker := time.NewTicker(time.Second * 30)
	defer ticker.Stop()

	var lastHeight int64

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			height, err := monitorCometOnce(ctx, rpcClient, isSyncing, lastHeight, readiness)
			if err != nil {
				log.Warn(ctx, "Failed monitoring cometBFT (will retry)", err)
				// Don't reset lastHeight to zero.
			} else {
				lastHeight = height
			}

			// Monitor db size.
			size, err := dirSize(dbDir)
			if err != nil {
				log.Warn(ctx, "Failed monitoring db size (will retry)", err)
			} else {
				dbSize.Set(float64(size))
			}
		}
	}
}

// monitorCometOnce monitors the cometBFT peers, and sync status.
func monitorCometOnce(ctx context.Context, rpcClient rpcclient.Client, isSyncing func() bool, lastHeight int64, readiness *ReadyResponse) (int64, error) {
	if netInfo, err := rpcClient.NetInfo(ctx); err != nil {
		return 0, errors.Wrap(err, "net info")
	} else if netInfo.NPeers == 0 {
		log.Error(ctx, "Halo has 0 consensus p2p peers", nil)
	} else {
		readiness.SetConsensusP2PPeers(netInfo.NPeers)
	}

	synced := !isSyncing()
	setConstantGauge(cometSynced, synced)
	readiness.SetConsensusSynced(synced)

	abciInfo, err := rpcClient.ABCIInfo(ctx)
	if err != nil {
		return 0, errors.Wrap(err, "abci info")
	} else if !isSyncing() && lastHeight > 0 && abciInfo.Response.LastBlockHeight <= lastHeight {
		log.Warn(ctx, "Halo height is not increasing, evm syncing?", nil, "height", abciInfo.Response.LastBlockHeight)
	}

	return abciInfo.Response.LastBlockHeight, nil
}

// monitorEVMForever blocks until the contract is canceled.
// It periodically calls monitorEVMOnce.
func monitorEVMForever(ctx context.Context, cfg Config, ethCl ethclient.Client, readiness *ReadyResponse) {
	if cfg.Network == netconf.Simnet {
		return // Simnet doesn't have an EVM tp monitor.
	}

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
			readiness.SetExecutionConnected(true)
		}
	}

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			err := monitorEVMOnce(ctx, ethCl, readiness)
			if err != nil {
				log.Warn(ctx, "Failed monitoring attached omni evm (will retry)", err, "addr", ethCl.Address())
			}
		}
	}
}

// monitorEVMOnce monitors the attached omni_evm height, peers, and sync status.
func monitorEVMOnce(ctx context.Context, ethCl ethclient.Client, readiness *ReadyResponse) error {
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
		readiness.SetExecutionP2PPeers(peers)
	}

	if syncing, err := ethCl.SyncProgress(ctx); err != nil {
		return errors.Wrap(err, "sync progress")
	} else if syncing != nil && !syncing.Done() {
		// SyncProgress returns nil of not syncing.
		evmSynced.Set(0)
		log.Warn(ctx, "Attached omni evm is syncing", nil, "highest_block", syncing.HighestBlock, "current_block", syncing.CurrentBlock, "tx_indexing", syncing.TxIndexRemainingBlocks)
	} else {
		evmSynced.Set(1)
		readiness.SetConsensusSynced(true)
	}

	latest, err := ethCl.BlockNumber(ctx)
	if err != nil {
		return errors.Wrap(err, "block number")
	}
	evmHeight.Set(float64(latest))

	return nil
}

// dirSize returns the total size of the directory at path.
func dirSize(path string) (int64, error) {
	var size int64
	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if os.IsNotExist(err) {
			return nil // Ignore files deleted during walk
		} else if err != nil {
			return err
		}

		if !info.IsDir() {
			size += info.Size()
		}

		return nil
	})
	if err != nil {
		return 0, errors.Wrap(err, "walk")
	}

	return size, nil
}
