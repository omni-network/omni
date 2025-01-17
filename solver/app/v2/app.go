package appv2

import (
	"context"
	"net/http"
	"net/http/pprof"
	"time"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/halo/genutil/evm/predeploys"
	"github.com/omni-network/omni/lib/buildinfo"
	"github.com/omni-network/omni/lib/contracts"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/expbackoff"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"
	xprovider "github.com/omni-network/omni/lib/xchain/provider"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// confLevel of solver streamers.
const (
	confLevel = xchain.ConfLatest
	unknown   = "unknown"
)

func chainVerFromID(id uint64) xchain.ChainVersion {
	return xchain.ChainVersion{ID: id, ConfLevel: confLevel}
}

// Run starts the solver service.
func Run(ctx context.Context, cfg Config) error {
	log.Info(ctx, "Starting solver v2 service")

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

	// TODO: log supported tokens / balances

	if cfg.SolverPrivKey == "" {
		return errors.New("private key not set")
	}
	privKey, err := ethcrypto.LoadECDSA(cfg.SolverPrivKey)
	if err != nil {
		return errors.Wrap(err, "load private key")
	}
	solverAddr := ethcrypto.PubkeyToAddress(privKey.PublicKey)
	log.Debug(ctx, "Using solver address", "address", solverAddr.Hex())

	backends, err := ethbackend.BackendsFromNetwork(network, cfg.RPCEndpoints, privKey)
	if err != nil {
		return err
	}

	// This starts devapp loadgen, currently written for v1 inbox / outbox.
	// if err := maybeStartLoadGen(ctx, cfg, network.ID, backends); err != nil {
	// 	return err
	// }

	xprov := xprovider.New(network, backends.Clients(), nil)

	db, err := newSolverDB(cfg.DBDir)
	if err != nil {
		return err
	}

	cursors, err := newCursors(db)
	if err != nil {
		return errors.Wrap(err, "create cursor store")
	}

	err = startEventStreams(ctx, network, xprov, backends, solverAddr, cursors)
	if err != nil {
		return errors.Wrap(err, "start event streams")
	}

	select {
	case <-ctx.Done():
		log.Info(ctx, "Shutdown detected, stopping...")
		return nil
	case err := <-monitorChan:
		return err
	}
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

// startEventStreams starts the event streams for the solver.
// TODO(corver): Make this robust against chains not be available on startup.
func startEventStreams(
	ctx context.Context,
	network netconf.Network,
	xprov xchain.Provider,
	backends ethbackend.Backends,
	solverAddr common.Address,
	cursors *cursors,
) error {
	addrs, err := contracts.GetAddresses(ctx, network.ID)
	if err != nil {
		return errors.Wrap(err, "get contract addresses")
	}

	inboxChains, err := detectContractChains(ctx, network, backends, addrs.SolverNetInbox)
	if err != nil {
		return errors.Wrap(err, "detect inbox chains")
	}

	inboxContracts := make(map[uint64]*bindings.SolverNetInbox)
	for _, chain := range inboxChains {
		name := network.ChainName(chain)
		chainVer := chainVerFromID(chain)
		log.Debug(ctx, "Using inbox contract", "chain", name, "address", addrs.SolverNetInbox.Hex())

		backend, err := backends.Backend(chain)
		if err != nil {
			return err
		}

		inbox, err := bindings.NewSolverNetInbox(addrs.SolverNetInbox, backend)
		if err != nil {
			return errors.Wrap(err, "create inbox contract", "chain", name)
		}
		inboxContracts[chain] = inbox

		// Check if cursor store should be initialized with deploy height
		if _, ok, err := cursors.Get(ctx, chainVer); err != nil {
			return errors.Wrap(err, "get cursor", "chain", name)
		} else if ok { // Cursor already set, skip
			continue
		}

		height, err := inbox.DeployedAt(&bind.CallOpts{Context: ctx})
		if err != nil {
			return errors.New("get inbox deploy height", "chain", name)
		}

		log.Info(ctx, "Initializing inbox cursor", "chain", name, "deployed_at", height)

		if err := cursors.Set(ctx, chainVer, height.Uint64()); err != nil {
			return err
		}
	}

	outboxChains, err := detectContractChains(ctx, network, backends, addrs.SolverNetOutbox)
	if err != nil {
		return errors.Wrap(err, "detect outbox chains")
	}

	outboxContracts := make(map[uint64]*bindings.SolverNetOutbox)
	for _, chain := range outboxChains {
		name := network.ChainName(chain)
		log.Debug(ctx, "Using outbox contract", "chain", name, "address", addrs.SolverNetOutbox.Hex())

		backend, err := backends.Backend(chain)
		if err != nil {
			return err
		}

		outbox, err := bindings.NewSolverNetOutbox(addrs.SolverNetOutbox, backend)
		if err != nil {
			return errors.Wrap(err, "create outbox contract", "chain", name)
		}
		outboxContracts[chain] = outbox
	}

	cursorSetter := func(ctx context.Context, chainID uint64, height uint64) error {
		return cursors.Set(ctx, chainVerFromID(chainID), height)
	}

	targetName := func(o Order) string {
		fill, err := o.ParsedFillOriginData()
		if err != nil {
			return unknown
		}

		// TODO: Give known targets friendly names
		return toEthAddr(fill.Call.Target).Hex()
	}

	deps := procDeps{
		ParseID:      newIDParser(inboxContracts),
		GetOrder:     newOrderGetter(inboxContracts),
		ShouldReject: newShouldRejector(backends, solverAddr, targetName, network.ChainName),
		Accept:       newAcceptor(inboxContracts, backends, solverAddr),
		Reject:       newRejector(inboxContracts, backends, solverAddr),
		Fill:         newFiller(outboxContracts, backends, solverAddr, addrs.SolverNetOutbox),
		Claim:        newClaimer(inboxContracts, backends, solverAddr),
		SetCursor:    cursorSetter,
		ChainName:    network.ChainName,
		TargetName:   targetName,
	}

	for _, chain := range inboxChains {
		log.Info(ctx, "Starting inbox event stream", "chain", network.ChainName(chain))
		go streamEventsForever(ctx, chain, xprov, deps, cursors, addrs.SolverNetInbox)
	}

	return nil
}

// streamEventsForever streams events from the inbox contract on the given chain.
func streamEventsForever(
	ctx context.Context,
	chainID uint64,
	xprov xchain.Provider,
	deps procDeps,
	cursors *cursors,
	inboxAddr common.Address,
) {
	backoff := expbackoff.New(ctx, expbackoff.WithPeriodicConfig(time.Second*5))
	for {
		from, ok, err := cursors.Get(ctx, xchain.ChainVersion{ID: chainID, ConfLevel: confLevel})
		if !ok || err != nil {
			log.Warn(ctx, "Failed reading cursor (will retry)", err)
			backoff()

			continue
		}

		req := xchain.EventLogsReq{
			ChainID:       chainID,
			Height:        from, // Note the previous height is re-processed (idempotency FTW)
			ConfLevel:     confLevel,
			FilterAddress: inboxAddr,
			FilterTopics:  allEventTopics,
		}
		err = xprov.StreamEventLogs(ctx, req, newEventProcessor(deps, chainID))
		if ctx.Err() != nil {
			return
		}

		log.Warn(ctx, "Failure streaming inbox events (will retry)", err)
		backoff()
	}
}
