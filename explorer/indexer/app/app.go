package app

import (
	"context"
	"net/http"
	"time"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/explorer/db"
	"github.com/omni-network/omni/explorer/db/ent"
	"github.com/omni-network/omni/halo/genutil/evm/predeploys"
	"github.com/omni-network/omni/lib/buildinfo"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"
	"github.com/omni-network/omni/lib/xchain/provider"

	"github.com/ethereum/go-ethereum/common"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Run(ctx context.Context, cfg Config) error {
	log.Info(ctx, "Starting Explorer Indexer")

	buildinfo.Instrument(ctx)

	// Start monitoring first, so app is "up".s
	monitorChan := serveMonitoring(cfg.MonitoringAddr)

	portalReg, err := makePortalRegistry(cfg.Network, cfg.RPCEndpoints)
	if err != nil {
		return err
	}

	network, err := netconf.AwaitOnChain(ctx, cfg.Network, portalReg, cfg.RPCEndpoints.Keys())
	if err != nil {
		return err
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

	err = startXProvider(ctx, network, entCl, cfg.RPCEndpoints)
	if err != nil {
		return errors.Wrap(err, "provider")
	}

	select {
	case <-ctx.Done():
		log.Info(ctx, "Shutdown detected, stopping...")
		return nil
	case err := <-monitorChan:
		return err
	}
}

// startXProvider all of our providers and subscribes to the chains in the network config.
func startXProvider(ctx context.Context, network netconf.Network, entCl *ent.Client, endpoints xchain.RPCEndpoints) error {
	rpcClientPerChain, err := initializeRPCClients(network.EVMChains(), endpoints)
	if err != nil {
		return err
	}

	xprovider := provider.New(network, rpcClientPerChain, nil)
	if xprovider == nil {
		return errors.New("failed to create xchain provider")
	}
	callback := newCallback(entCl)

	for _, chain := range network.EVMChains() {
		height, offset, err := InitChainCursor(ctx, entCl, chain)
		if err != nil {
			return errors.Wrap(err, "initialize chain cursor", "chain_id", chain.ID)
		}
		log.Info(ctx, "Subscribing to chain", "chain_id", chain.ID, "from_height", height)

		// TODO(corver): Store MsgOffset along with heights in cursor table
		//  Currently XBlockOffset isn't supported in x-explorer.
		err = xprovider.StreamAsync(ctx, chain.ID, height, offset+1, callback)
		if err != nil {
			return errors.Wrap(err, "subscribe", "chain_id", chain.ID)
		}
	}

	return nil
}

// initializeRPCClients initializes the rpc clients for all evm chains in the network.
func initializeRPCClients(chains []netconf.Chain, endpoints xchain.RPCEndpoints) (map[uint64]ethclient.Client, error) {
	rpcClientPerChain := make(map[uint64]ethclient.Client)
	for _, chain := range chains {
		rpc, err := endpoints.ByNameOrID(chain.Name, chain.ID)
		if err != nil {
			return nil, err
		}

		client, err := ethclient.Dial(chain.Name, rpc)
		if err != nil {
			return nil, errors.Wrap(err, "dial rpc", "chain_name", chain.Name, "chain_id", chain.ID, "rpc_url", rpc)
		}
		rpcClientPerChain[chain.ID] = client
	}

	return rpcClientPerChain, nil
}

// InitChainCursor return the initial cursor height to start streaming from (inclusive).
// If a cursor exists, it returns the cursor height + 1.
// Else a new cursor is created with chain deploy height.
func InitChainCursor(ctx context.Context, entCl *ent.Client, chain netconf.Chain) (uint64, uint64, error) { //nolint:revive // false positive
	c, ok, err := cursor(ctx, entCl.XProviderCursor, chain.ID)
	if err != nil {
		return 0, 0, errors.Wrap(err, "cursor")
	}
	if ok {
		return c.Height + 1, c.Offset + 1, nil
	}

	// Store the cursor at deploy height - 1, so first cursor update will be at deploy height 0.
	deployHeight := chain.DeployHeight - 1
	if chain.DeployHeight == 0 { // Except for 0, we handle this explicitly.
		deployHeight = 0
	}

	offset := uint64(1)

	// cursor doesn't exist yet, create it
	_, err = entCl.XProviderCursor.Create().SetChainID(chain.ID).SetHeight(deployHeight).SetOffset(offset).Save(ctx)
	if err != nil {
		return 0, 0, errors.Wrap(err, "create cursor")
	}

	// if the cursor doesn't exist that means the chain doesn't exist so we have to create it as well
	_, err = entCl.Chain.Create().SetChainID(chain.ID).SetName(chain.Name).Save(ctx)
	if err != nil {
		return 0, 0, errors.Wrap(err, "create chain")
	}

	log.Debug(ctx, "Created cursor", "chain_id", chain.ID, "height", deployHeight, "offset", offset)

	return chain.DeployHeight, offset, nil
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

func makePortalRegistry(network netconf.ID, endpoints xchain.RPCEndpoints) (*bindings.PortalRegistry, error) {
	meta := netconf.MetadataByID(network, network.Static().OmniExecutionChainID)
	rpc, err := endpoints.ByNameOrID(meta.Name, meta.ChainID)
	if err != nil {
		return nil, err
	}

	ethCl, err := ethclient.Dial(meta.Name, rpc)
	if err != nil {
		return nil, err
	}

	resp, err := bindings.NewPortalRegistry(common.HexToAddress(predeploys.PortalRegistry), ethCl)
	if err != nil {
		return nil, errors.Wrap(err, "create portal registry")
	}

	return resp, nil
}
