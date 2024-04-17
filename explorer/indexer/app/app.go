package app

import (
	"context"
	"net/http"
	"time"

	"github.com/omni-network/omni/explorer/db"
	"github.com/omni-network/omni/explorer/db/ent"
	"github.com/omni-network/omni/lib/buildinfo"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain/provider"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Run(ctx context.Context, cfg Config) error {
	log.Info(ctx, "Starting Explorer Indexer")

	buildinfo.Instrument(ctx)

	network, err := netconf.Load(cfg.NetworkFile)
	if err != nil {
		return errors.Wrap(err, "load network config")
	}

	entCl, err := db.NewPostgressClient(cfg.ExplorerDBConn)
	if err != nil {
		return errors.Wrap(err, "new db client")
	}

	defer func(entCl *ent.Client) {
		err := entCl.Close()
		if err != nil {
			log.Error(ctx, "Failed to close ent client", err)
		}
	}(entCl)

	if err := db.CreateSchema(ctx, entCl); err != nil {
		return errors.Wrap(err, "create schema")
	}

	err = startXProvider(ctx, network, entCl)
	if err != nil {
		return errors.Wrap(err, "provider")
	}

	select {
	case <-ctx.Done():
		log.Info(ctx, "Shutdown detected, stopping...")
		return nil
	case err := <-serveMonitoring(cfg.MonitoringAddr):
		return err
	}
}

// startXProvider all of our providers and subscribes to the chains in the network config.
func startXProvider(ctx context.Context, network netconf.Network, entCl *ent.Client) error {
	rpcClientPerChain, err := initializeRPCClients(network.EVMChains())
	if err != nil {
		return err
	}

	xprovider := provider.New(network, rpcClientPerChain, nil)
	if xprovider == nil {
		return errors.New("failed to create xchain provider")
	}
	callback := newCallback(entCl)

	for _, chain := range network.EVMChains() {
		fromHeight, err := InitChainCursor(ctx, entCl, chain)
		if err != nil {
			return errors.Wrap(err, "initialize chain cursor", "chain_id", chain.ID)
		}
		log.Info(ctx, "Subscribing to chain", "chain_id", chain.ID, "from_height", fromHeight)

		err = xprovider.StreamAsync(ctx, chain.ID, fromHeight, callback)
		if err != nil {
			return errors.Wrap(err, "subscribe", "chain_id", chain.ID)
		}
	}

	return nil
}

// initializeRPCClients initializes the rpc clients for all evm chains in the network.
func initializeRPCClients(chains []netconf.Chain) (map[uint64]ethclient.Client, error) {
	rpcClientPerChain := make(map[uint64]ethclient.Client)
	for _, chain := range chains {
		client, err := ethclient.Dial(chain.Name, chain.RPCURL)
		if err != nil {
			return nil, errors.Wrap(err, "dial rpc", "chain_id", chain.ID, "rpc_url", chain.RPCURL)
		}
		rpcClientPerChain[chain.ID] = client
	}

	return rpcClientPerChain, nil
}

// InitChainCursor return the initial cursor height to start streaming from (inclusive).
// If a cursor exists, it returns the cursor height + 1.
// Else a new cursor is created with chain deploy height.
func InitChainCursor(ctx context.Context, entCl *ent.Client, chain netconf.Chain) (uint64, error) {
	cursor, ok, err := getCursor(ctx, entCl.XProviderCursor, chain.ID)
	if err != nil {
		return 0, errors.Wrap(err, "get cursor")
	} else if ok {
		return cursor.Height + 1, nil
	}

	// Store the cursor at deploy height - 1, so first cursor update will be at deploy height.
	deployHeight := chain.DeployHeight - 1
	if chain.DeployHeight == 0 { // Except for 0, we handle this explicitly.
		deployHeight = 0
	}

	// cursor doesn't exist yet, create it
	_, err = entCl.XProviderCursor.Create().
		SetChainID(chain.ID).
		SetHeight(deployHeight).
		Save(ctx)
	if err != nil {
		return 0, errors.Wrap(err, "create cursor")
	}

	// if the cursor doesn't exist that means the chain doesn't exist so we have to create it as well
	_, err = entCl.Chain.
		Create().
		SetChainID(chain.ID).
		SetName(chain.Name).
		Save(ctx)
	if err != nil {
		return 0, errors.Wrap(err, "create chain")
	}

	log.Info(ctx, "Created cursor", "chain_id", chain.ID, "height", deployHeight)

	return chain.DeployHeight, nil
}

// serveMonitoring starts a goroutine that serves the monitoring API. It
// returns a channel that will receive an error if the server fails to start.
func serveMonitoring(address string) <-chan error {
	errChan := make(chan error)
	go func() {
		mux := http.NewServeMux()
		mux.Handle("/metrics", promhttp.Handler())

		srv := &http.Server{
			Addr:              address,
			ReadHeaderTimeout: 5 * time.Second,
			IdleTimeout:       5 * time.Second,
			WriteTimeout:      5 * time.Second,
			Handler:           mux,
		}
		errChan <- errors.Wrap(srv.ListenAndServe(), "serve monitoring")
	}()

	return errChan
}
