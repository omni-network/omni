package app

import (
	"context"
	"net/http"
	"net/http/pprof"
	"sync"
	"time"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/halo/genutil/evm/predeploys"
	"github.com/omni-network/omni/lib/buildinfo"
	"github.com/omni-network/omni/lib/contracts"
	"github.com/omni-network/omni/lib/contracts/solvernet"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/expbackoff"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	tokenslib "github.com/omni-network/omni/lib/tokens"
	"github.com/omni-network/omni/lib/tokens/coingecko"
	"github.com/omni-network/omni/lib/tracer"
	"github.com/omni-network/omni/lib/xchain"
	xprovider "github.com/omni-network/omni/lib/xchain/provider"
	"github.com/omni-network/omni/solver/targets"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// chainVersFromID returns the chain versions to stream/process per chain ID.
func chainVersFromID(network netconf.ID, chainID uint64) []xchain.ChainVersion {
	// On devnet we stream twice (for idempotency testing)
	if network == netconf.Devnet {
		return []xchain.ChainVersion{
			xchain.NewChainVersion(chainID, xchain.ConfLatest),
			xchain.NewChainVersion(chainID, xchain.ConfMin2),
		}
	}

	// For ethereum, we stream min2 to reduce reorg risk
	if chainID == netconf.EthereumChainID(network) {
		return []xchain.ChainVersion{
			xchain.NewChainVersion(chainID, xchain.ConfMin2),
		}
	}

	// On other chains, we only stream latest for now,
	return []xchain.ChainVersion{
		xchain.NewChainVersion(chainID, xchain.ConfLatest),
	}
}

// Run starts the solver service.
func Run(ctx context.Context, cfg Config) error {
	log.Info(ctx, "Starting solver service")

	buildinfo.Instrument(ctx)

	tracerID := tracer.Identifiers{Network: cfg.Network, Service: "solver"}
	stopTracer, err := tracer.Init(ctx, tracerID, cfg.Tracer)
	if err != nil {
		return err
	}
	defer stopTracer(ctx) //nolint:errcheck // Tracing shutdown errors not critical

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

	xprov := xprovider.New(network, backends.Clients(), nil)

	db, err := newSolverDB(cfg.DBDir)
	if err != nil {
		return err
	}

	cursors, err := newCursors(db)
	if err != nil {
		return errors.Wrap(err, "create cursor store")
	}

	addrs, err := contracts.GetAddresses(ctx, network.ID)
	if err != nil {
		return errors.Wrap(err, "get contract addresses")
	}

	if err := approveOutboxes(ctx, network, backends, solverAddr); err != nil {
		return errors.Wrap(err, "approve outboxes")
	}

	pricer := newPricer(ctx, cfg.CoinGeckoAPIKey)

	err = startEventStreams(ctx, network, xprov, backends, solverAddr, addrs, cursors, pricer)
	if err != nil {
		return errors.Wrap(err, "start event streams")
	}

	err = startRebalancing(ctx, network, backends, newSimpleGasPnLFunc(pricer, network.ChainName))
	if err != nil {
		return errors.Wrap(err, "start rebalancing")
	}

	callAllower := newCallAllower(network.ID, addrs.SolverNetMiddleman)

	log.Info(ctx, "Serving API", "address", cfg.APIAddr)
	apiChan := serveAPI(cfg.APIAddr,
		newCheckHandler(newChecker(backends, callAllower, solverAddr, addrs.SolverNetOutbox)),
		newContractsHandler(addrs),
		newQuoteHandler(quoter),
	)

	select {
	case <-ctx.Done():
		log.Info(ctx, "Shutdown detected, stopping...")
		return nil
	case err := <-monitorChan:
		return err
	case err := <-apiChan:
		return err
	}
}

func newPricer(ctx context.Context, apiKey string) tokenslib.Pricer {
	pricer := tokenslib.NewCachedPricer(coingecko.New(coingecko.WithAPIKey(apiKey)))

	// use cached pricer avoid spamming coingecko public api
	const priceCacheEvictInterval = time.Minute * 10
	go pricer.ClearCacheForever(ctx, priceCacheEvictInterval)

	return pricer
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
	addrs contracts.Addresses,
	cursors *cursors,
	pricer tokenslib.Pricer,
) error {
	inboxChains, err := detectContractChains(ctx, network, backends, addrs.SolverNetInbox)
	if err != nil {
		return errors.Wrap(err, "detect inbox chains")
	}

	inboxContracts := make(map[uint64]*bindings.SolverNetInbox)
	for _, chain := range inboxChains {
		name := network.ChainName(chain)
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

		// Bootstrap all streamer cursors for this chain
		for _, chainVer := range chainVersFromID(network.ID, chain) {
			loopCtx := log.WithCtx(ctx, "chain_version", network.ChainVersionName(chainVer))
			if err := maybeBootstrapCursor(loopCtx, inbox, cursors, chainVer); err != nil {
				return err
			}
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

	targetName := func(pendingData PendingData) string {
		fill, err := pendingData.ParsedFillOriginData()
		if err != nil {
			return "unknown"
		}

		// use last call target for target name
		call := fill.Calls[len(fill.Calls)-1]

		if tkn, ok := tokens.Find(pendingData.DestinationChainID, call.Target); ok {
			return "ERC20:" + tkn.Symbol
		}

		// Native bridging has zero call data and positive value
		isNative := call.Selector == [4]byte{} && len(call.Params) == 0 && call.Value.Sign() > 0
		if nativeTkn, ok := tokens.Find(pendingData.DestinationChainID, NativeAddr); ok && isNative {
			return "Native:" + nativeTkn.Symbol
		}

		if target, ok := targets.Get(pendingData.DestinationChainID, call.Target); ok {
			return target.Name
		}

		return call.Target.Hex()[:7] // Short hex.
	}

	callAllower := newCallAllower(network.ID, addrs.SolverNetMiddleman)

	ageCache := newAgeCache(backends)
	go monitorAgeCacheForever(ctx, ageCache, network.ChainName)

	filledPnL := newFilledPnlFunc(pricer, targetName, network.ChainName, addrs.SolverNetOutbox, ageCache.InstrumentDestFilled)
	orderGasPnL := newOrderGasPnLFunc(pricer, network.ChainName)

	for _, chainID := range inboxChains {
		// Ensure chain version processors don't process same height concurrently.
		callbackWrapper := newHeightMutexer()
		for _, chainVer := range chainVersFromID(network.ID, chainID) {
			cursorSetter := func(ctx context.Context, _ uint64, height uint64) error {
				return cursors.Set(ctx, chainVer, height)
			}

			deps := procDeps{
				ParseID:       newIDParser(inboxContracts),
				GetOrder:      newOrderGetter(inboxContracts),
				ShouldReject:  newShouldRejector(backends, callAllower, solverAddr, addrs.SolverNetOutbox),
				DidFill:       newDidFiller(outboxContracts),
				Reject:        newRejector(inboxContracts, backends, solverAddr, orderGasPnL),
				Fill:          newFiller(outboxContracts, backends, solverAddr, addrs.SolverNetOutbox, filledPnL),
				Claim:         newClaimer(network.ID, inboxContracts, backends, solverAddr, orderGasPnL),
				SetCursor:     cursorSetter,
				ChainName:     network.ChainName,
				ProcessorName: network.ChainVersionName(chainVer),
				TargetName:    targetName,
				InstrumentAge: ageCache.InstrumentAge,
			}

			loopCtx := log.WithCtx(ctx, "proc", network.ChainVersionName(chainVer))
			log.Info(loopCtx, "Starting inbox event processor")
			go streamEventsForever(loopCtx, chainVer, xprov, deps, cursors, addrs.SolverNetInbox, callbackWrapper)
		}
	}

	return nil
}

// newHeightMutexer returns a callback wrapper that ensures a specific chain height is NOT processed concurrently.
// This prevents races when multiple chain versions concurrently processing log events.
// Since we wait for submitted txs to be mined, subsequent processors will read state from chain.
// The risk is wasting gas due to re-submitting the same tx.
func newHeightMutexer() func(callback xchain.EventLogsCallback) xchain.EventLogsCallback {
	var mutexes sync.Map
	return func(callback xchain.EventLogsCallback) xchain.EventLogsCallback {
		return func(ctx context.Context, height uint64, events []types.Log) error {
			anyMutex, _ := mutexes.LoadOrStore(height, new(sync.Mutex))
			mutex := anyMutex.(*sync.Mutex) //nolint:revive,forcetypeassert // Known type
			mutex.Lock()
			defer func() {
				mutex.Unlock()
				mutexes.Delete(height)
			}()

			return callback(ctx, height, events)
		}
	}
}

// streamEventsForever streams events from the inbox contract on the given chain version.
func streamEventsForever(
	ctx context.Context,
	chainVer xchain.ChainVersion,
	xprov xchain.Provider,
	deps procDeps,
	cursors *cursors,
	inboxAddr common.Address,
	callbackWrapper func(xchain.EventLogsCallback) xchain.EventLogsCallback,
) {
	backoff := expbackoff.New(ctx, expbackoff.WithPeriodicConfig(time.Second*5))
	for {
		from, ok, err := cursors.Get(ctx, chainVer)
		if !ok || err != nil {
			log.Warn(ctx, "Failed reading cursor (will retry)", err)
			backoff()

			continue
		}

		req := xchain.EventLogsReq{
			ChainID:       chainVer.ID,
			Height:        from, // Note the previous height is re-processed (idempotency FTW)
			ConfLevel:     chainVer.ConfLevel,
			FilterAddress: inboxAddr,
			FilterTopics:  solvernet.AllEventTopics(),
		}
		err = xprov.StreamEventLogs(ctx, req, callbackWrapper(newEventProcessor(deps, chainVer.ID)))
		if ctx.Err() != nil {
			return
		}

		log.Warn(ctx, "Failure processing inbox events (will retry)", err)
		backoff()
	}
}
