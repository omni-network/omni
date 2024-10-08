package app

import (
	"context"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"

	cmtcfg "github.com/cometbft/cometbft/config"
	rpcclient "github.com/cometbft/cometbft/rpc/client"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// monitorCometForever blocks until the context is canceled.
// It periodically calls monitorCometOnce.
func monitorCometForever(
	ctx context.Context,
	network netconf.ID,
	rpcClient rpcclient.Client,
	isSyncing func() bool,
	dbDir string,
	status *readinessStatus,
) {
	if network == netconf.Simnet {
		return // Simnet doesn't need to monitor cometBFT, since no p2p.
	}

	ticker := time.NewTicker(time.Second * 10)
	defer ticker.Stop()

	var lastHeight int64

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			height, err := monitorCometOnce(ctx, rpcClient, isSyncing, lastHeight, status)
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
func monitorCometOnce(ctx context.Context, rpcClient rpcclient.Client, isSyncing func() bool, lastHeight int64, status *readinessStatus) (int64, error) {
	netInfo, err := rpcClient.NetInfo(ctx)
	status.setConsensusP2PPeers(netInfo.NPeers)
	if err != nil {
		return 0, errors.Wrap(err, "net info")
	} else if netInfo.NPeers == 0 {
		log.Error(ctx, "Halo has 0 consensus p2p peers", nil)
	}

	synced := !isSyncing()
	setConstantGauge(cometSynced, synced)
	status.setConsensusSynced(synced)

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
func monitorEVMForever(ctx context.Context, cfg Config, ethCl ethclient.Client, status *readinessStatus) {
	if cfg.Network == netconf.Simnet {
		return // Simnet doesn't have an EVM tp monitor.
	}

	ticker := time.NewTicker(time.Second * 10)
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

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			err := monitorEVMOnce(ctx, ethCl, status)
			if err != nil {
				log.Warn(ctx, "Failed monitoring attached omni evm (will retry)", err, "addr", ethCl.Address())
				status.setExecutionConnected(false)
			}
		}
	}
}

// monitorEVMOnce monitors the attached omni_evm height, peers, and sync status.
func monitorEVMOnce(ctx context.Context, ethCl ethclient.Client, status *readinessStatus) error {
	// Best effort monitoring of peer count, since method not available in auth API.
	peers, err := ethCl.PeerCount(ctx)
	if ethclient.IsErrMethodNotAvailable(err) { //nolint:revive // Empty block skips error handling below.
		// Do not set the metric if the method is not available.
	} else if err != nil {
		return errors.Wrap(err, "peer count")
	} else {
		if peers == 0 {
			log.Error(ctx, "Attached omni evm has 0 peers", nil)
		}
		evmPeers.Set(float64(peers))
		status.setExecutionP2PPeers(peers)
	}

	if syncing, err := ethCl.SyncProgress(ctx); err != nil {
		return errors.Wrap(err, "sync progress")
	} else if syncing != nil && !syncing.Done() {
		// SyncProgress returns nil of not syncing.
		evmSynced.Set(0)
		log.Warn(ctx, "Attached omni evm is syncing", nil, "highest_block", syncing.HighestBlock, "current_block", syncing.CurrentBlock, "tx_indexing", syncing.TxIndexRemainingBlocks)
	} else {
		evmSynced.Set(1)
		status.setExecutionSynced(true)
	}

	status.setExecutionConnected(true)

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

// startMonitoringAPI registers metrics and health endpoints.
// Return the function shutting down the HTTP server.
func startMonitoringAPI(
	cfg *cmtcfg.Config,
	asyncAbort chan<- error,
	status *readinessStatus) func(context.Context) error {
	mux := http.NewServeMux()

	mux.Handle("/metrics", promhttp.Handler())
	// We also export the metrics on the root path as this is how the default prometheus server works.
	mux.Handle("/", promhttp.Handler())

	// On the `/ready` endpoint, we serve the readiness of the halo node.
	// Additionally, in case when the node is not ready, the response status code is set to 503.
	mux.HandleFunc("/ready", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if !status.ready() {
			w.WriteHeader(http.StatusServiceUnavailable)
		}
		if err := status.serialize(w); err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	})

	server := &http.Server{
		Addr:              cfg.Instrumentation.PrometheusListenAddr,
		ReadHeaderTimeout: 3 * time.Second,
		Handler:           mux,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			asyncAbort <- errors.Wrap(err, "http server failed to start")
		}
	}()

	return server.Shutdown
}
