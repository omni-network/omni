package monitor

import (
	"context"
	"net/http"
	"net/http/pprof"
	"time"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/halo/genutil/evm/predeploys"
	"github.com/omni-network/omni/lib/buildinfo"
	cprovider "github.com/omni-network/omni/lib/cchain/provider"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"
	xprovider "github.com/omni-network/omni/lib/xchain/provider"
	"github.com/omni-network/omni/monitor/account"
	"github.com/omni-network/omni/monitor/avs"
	"github.com/omni-network/omni/monitor/contract"
	"github.com/omni-network/omni/monitor/loadgen"
	"github.com/omni-network/omni/monitor/routerecon"
	"github.com/omni-network/omni/monitor/validator"
	"github.com/omni-network/omni/monitor/xfeemngr"
	"github.com/omni-network/omni/monitor/xmonitor"
	"github.com/omni-network/omni/monitor/xmonitor/indexer"

	"github.com/cometbft/cometbft/rpc/client"
	comethttp "github.com/cometbft/cometbft/rpc/client/http"

	"github.com/ethereum/go-ethereum/common"

	dbm "github.com/cosmos/cosmos-db"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Run starts the monitor service.
func Run(ctx context.Context, cfg Config) error {
	log.Info(ctx, "Starting monitor service")

	buildinfo.Instrument(ctx)

	// Start monitoring first, so app is "up"
	monitorChan := serveMonitoring(cfg.MonitoringAddr)

	portalReg, err := makePortalRegistry(cfg.Network, cfg.RPCEndpoints)
	if err != nil {
		return err
	}

	network, err := netconf.AwaitOnExecutionChain(ctx, cfg.Network, portalReg, cfg.RPCEndpoints.Keys())
	if err != nil {
		return err
	}

	ethClients, err := initializeEthClients(network.EVMChains(), cfg.RPCEndpoints)
	if err != nil {
		return err
	}

	if cfg.HaloURL == "" {
		return errors.New("empty --halo-url flag")
	}

	tmClient, err := newClient(cfg.HaloURL)
	if err != nil {
		return err
	}

	cprov := cprovider.NewABCI(tmClient, network.ID)

	xprov := xprovider.New(network, ethClients, cprov)

	if err := avs.StartMonitor(ctx, network, ethClients); err != nil {
		return errors.Wrap(err, "monitor AVS")
	}

	account.StartMonitoring(ctx, network, ethClients)

	if err := contract.StartMonitoring(ctx, network, cfg.RPCEndpoints, ethClients); err != nil {
		return errors.Wrap(err, "monitor contracts")
	}

	if err := startLoadGen(ctx, cfg, network, ethClients); err != nil {
		return errors.Wrap(err, "start load generator")
	}

	if err := xmonitor.Start(ctx, network, xprov, cprov, ethClients); err != nil {
		return errors.Wrap(err, "start xchain monitor")
	}

	if err := startIndexer(ctx, cfg, network, xprov); err != nil {
		return errors.Wrap(err, "start xchain indexer")
	}

	if err := xfeemngr.Start(ctx, network, cfg.XFeeMngr, cfg.PrivateKey); err != nil {
		log.Error(ctx, "Failed to start xfee manager [BUG]", err)
	}

	startMonitoringSyncDiff(ctx, network, ethClients)
	go runHistoricalBaselineForever(ctx, network, cprov)
	go monitorUpgradesForever(ctx, cprov)
	go routerecon.ReconForever(ctx, network, xprov, ethClients)
	go validator.MonitorForever(ctx, cprov)
	go monitorPublicRPCForever(ctx, network, ethClients)

	select {
	case <-ctx.Done():
		log.Info(ctx, "Shutdown detected, stopping...")
		return nil
	case err := <-monitorChan:
		return err
	}
}

// startIndexer starts the xchain indexer.
func startIndexer(
	ctx context.Context,
	cfg Config,
	network netconf.Network,
	xprov xchain.Provider,
) error {
	var db dbm.DB
	if cfg.DBDir == "" {
		log.Warn(ctx, "No --db-dir provided, using in-memory DB", nil)
		db = dbm.NewMemDB()
	} else {
		var err error
		db, err = dbm.NewGoLevelDB("indexer", cfg.DBDir, nil)
		if err != nil {
			return errors.Wrap(err, "new golevel db")
		}
	}

	return indexer.Start(ctx, network, xprov, db)
}

// serveMonitoring starts a goroutine that serves the monitoring API. It
// returns a channel that will receive an error if the server fails to start.
func serveMonitoring(address string) <-chan error {
	errChan := make(chan error)
	go func() {
		mux := http.NewServeMux()
		mux.Handle("/metrics", promhttp.Handler())

		// Copied from net/http/pprof/pprof.go
		mux.HandleFunc("/debug/pprof/", pprof.Index)
		mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
		mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
		mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
		mux.HandleFunc("/debug/pprof/trace", pprof.Trace)

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

func startLoadGen(ctx context.Context, cfg Config, network netconf.Network, ethClients map[uint64]ethclient.Client) error {
	if err := loadgen.Start(ctx, network, ethClients, cfg.LoadGen); err != nil {
		return errors.Wrap(err, "start load generator")
	}

	return nil
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

// newClient returns a new tendermint HTTP RPC client.
func newClient(tmNodeAddr string) (client.Client, error) {
	c, err := comethttp.New("tcp://"+tmNodeAddr, "/websocket")
	if err != nil {
		return nil, errors.Wrap(err, "new tendermint client")
	}

	return c, nil
}

// initializeEthClients initializes the RPC clients for the given chains.
func initializeEthClients(chains []netconf.Chain, endpoints xchain.RPCEndpoints) (map[uint64]ethclient.Client, error) {
	rpcClientPerChain := make(map[uint64]ethclient.Client)
	for _, chain := range chains {
		rpc, err := endpoints.ByNameOrID(chain.Name, chain.ID)
		if err != nil {
			return nil, err
		}
		c, err := ethclient.Dial(chain.Name, rpc)
		if err != nil {
			return nil, errors.Wrap(err, "dial rpc", "chain_name", chain.Name, "chain_id", chain.ID, "rpc_url", rpc)
		}
		rpcClientPerChain[chain.ID] = c
	}

	return rpcClientPerChain, nil
}
