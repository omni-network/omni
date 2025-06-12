package app

import (
	"context"
	"crypto/ecdsa"
	"net/http"
	"net/http/pprof"
	"time"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/halo/genutil/evm/predeploys"
	"github.com/omni-network/omni/lib/buildinfo"
	"github.com/omni-network/omni/lib/cctp"
	"github.com/omni-network/omni/lib/contracts"
	"github.com/omni-network/omni/lib/contracts/solvernet"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/expbackoff"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/tokenpricer"
	"github.com/omni-network/omni/lib/tokens"
	"github.com/omni-network/omni/lib/tracer"
	"github.com/omni-network/omni/lib/unibackend"
	"github.com/omni-network/omni/lib/xchain"
	xprovider "github.com/omni-network/omni/lib/xchain/provider"
	"github.com/omni-network/omni/solver/job"
	"github.com/omni-network/omni/solver/rebalance"
	"github.com/omni-network/omni/solver/targets"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"

	"github.com/gagliardetto/solana-go/rpc"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// chainVerFromID returns the chain version to stream/process per chain ID.
func chainVerFromID(network netconf.ID, chainID uint64) xchain.ChainVersion {
	// For ethereum L1, we stream min2 to reduce reorg risk
	if chainID == netconf.EthereumChainID(network) {
		return xchain.NewChainVersion(chainID, xchain.ConfMin2)
	}

	// For solana, we use confirmed as per docs
	if chainID == netconf.SolanaChainID(network) {
		return xchain.NewChainVersion(chainID, xchain.ConfConfirmed)
	}

	// On other chains, we only stream latest for now,
	return xchain.NewChainVersion(chainID, xchain.ConfLatest)
}

// Run starts the solver service.
func Run(ctx context.Context, cfg Config) error {
	log.Info(ctx, "Starting solver service")

	buildinfo.Instrument(ctx)

	tracerID := tracer.Identifiers{Network: cfg.Network.String(), Service: "solver"}
	stopTracer, err := tracer.Init(ctx, tracerID, cfg.Tracer)
	if err != nil {
		return err
	}
	defer stopTracer(ctx) //nolint:errcheck // Tracing shutdown errors not critical

	go targets.RefreshForever(ctx, cfg.Network)

	// Start monitoring first, so app is "up"
	monitorChan := serveMonitoring(cfg.MonitoringAddr)

	portalReg, err := makePortalRegistry(ctx, cfg.Network, cfg.RPCEndpoints)
	if err != nil {
		return err
	}

	network, err := netconf.AwaitOnExecutionChain(ctx, cfg.Network, portalReg, solvernet.OnlyCoreEndpoints(cfg.RPCEndpoints).Keys())
	if err != nil {
		return err
	}

	network = solvernet.AddNetwork(ctx, network, solvernet.FilterByContracts(ctx, cfg.RPCEndpoints))

	if cfg.SolverPrivKey == "" {
		return errors.New("private key not set")
	}
	privKey, err := ethcrypto.LoadECDSA(cfg.SolverPrivKey)
	if err != nil {
		return errors.Wrap(err, "load private key")
	}
	solverAddr := ethcrypto.PubkeyToAddress(privKey.PublicKey)
	log.Debug(ctx, "Using solver address", "address", solverAddr.Hex())

	backends, err := ethbackend.BackendsFromNetwork(ctx, network, cfg.RPCEndpoints, privKey)
	if err != nil {
		return err
	}
	backends.StartIdleConnectionClosing(ctx)
	uniBackends := unibackend.EVMBackends(backends)

	// Add SVM backends if any
	for _, svmChain := range network.SVMChains() {
		endpoint, err := cfg.RPCEndpoints.ByNameOrID(svmChain.Name, svmChain.ID)
		if err != nil {
			return err
		}
		uniBackends[svmChain.ID] = unibackend.SVMBackend(rpc.New(endpoint), svmChain.ID)
	}

	xprov := xprovider.New(network, backends.Clients(), nil)

	db, err := newSolverDB(cfg.DBDir)
	if err != nil {
		return err
	}

	jobDB, err := job.New(db)
	if err != nil {
		return err
	}
	// TODO(corver): Remove once all networks migrated
	count, err := job.Migrate(ctx, jobDB)
	if err != nil {
		return errors.Wrap(err, "migrate job db")
	}
	log.Info(ctx, "Migrated job db", "count", count)

	cursors, err := newCursors(db)
	if err != nil {
		return errors.Wrap(err, "create cursor store")
	}

	addrs, err := contracts.GetAddresses(ctx, network.ID)
	if err != nil {
		return errors.Wrap(err, "get contract addresses")
	}

	// Create inbox contracts early so they can be used by both event processing and relay
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

		// Initialize cursors for each chain
		chainVer := chainVerFromID(network.ID, chain)
		loopCtx := log.WithCtx(ctx, "chain_version", network.ChainVersionName(chainVer))
		if err := maybeBootstrapCursor(loopCtx, inbox, cursors, chainVer); err != nil {
			return err
		}
	}

	if err := approveOutboxes(ctx, network, backends, solverAddr); err != nil {
		return errors.Wrap(err, "approve outboxes")
	}

	pricer := newPricer(ctx, network.ID, cfg.CoinGeckoAPIKey)
	priceFunc := newPriceFunc(pricer)

	go monitorPricesForever(ctx, priceFunc)

	err = startProcessingEvents(ctx, network, xprov, jobDB, uniBackends, privKey, addrs, cursors, pricer, priceFunc, inboxContracts)
	if err != nil {
		return errors.Wrap(err, "start event streams")
	}

	if err := rebalance.Start(ctx, network, newCCTPClient(network.ID), pricer, backends, solverAddr, cfg.DBDir); err != nil {
		log.Warn(ctx, "Failed to start rebalancing [BUG]", err)
	}

	// TODO: move to solver/rebalance
	err = startRebalancingOMNI(ctx, network, backends, newSimpleGasPnLFunc(pricer, network.ChainName))
	if err != nil {
		return errors.Wrap(err, "start rebalancing omni")
	}

	callAllower := newCallAllower(network.ID, addrs.SolverNetExecutor)

	log.Info(ctx, "Serving API", "address", cfg.APIAddr)

	// Build base handlers that are always available
	checkFunc := newChecker(uniBackends, callAllower, priceFunc, solverAddr, addrs.SolverNetOutbox)
	handlers := []Handler{
		newCheckHandler(
			checkFunc,
			newTracer(backends, solverAddr, addrs.SolverNetOutbox),
		),
		newContractsHandler(addrs),
		newQuoteHandler(newQuoter(priceFunc)),
		newPriceHandler(wrapPriceHandlerFunc(priceFunc)),
		newTokensHandler(uniBackends, solverAddr),
	}

	// Only add relay handler for ephemeral networks
	if network.ID.IsEphemeral() {
		log.Debug(ctx, "Adding relay handler for ephemeral network", "network", network.ID)
		handlers = append(handlers, newRelayHandler(newRelayer(inboxContracts, uniBackends, solverAddr, addrs.SolverNetInbox, checkFunc)))
	}

	//nolint:contextcheck // False positive, inner context is used for shutdown
	apiChan, apiCancel := serveAPI(cfg.APIAddr, handlers...)
	defer apiCancel()

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

func makePortalRegistry(ctx context.Context, network netconf.ID, endpoints xchain.RPCEndpoints) (*bindings.PortalRegistry, error) {
	meta := netconf.MetadataByID(network, network.Static().OmniExecutionChainID)
	rpc, err := endpoints.ByNameOrID(meta.Name, meta.ChainID)
	if err != nil {
		return nil, err
	}

	ethCl, err := ethclient.DialContext(ctx, meta.Name, rpc)
	if err != nil {
		return nil, err
	}

	resp, err := bindings.NewPortalRegistry(common.HexToAddress(predeploys.PortalRegistry), ethCl)
	if err != nil {
		return nil, errors.Wrap(err, "create portal registry")
	}

	return resp, nil
}

// startProcessingEvents starts the event processing for the solver.
// It starts processing all existing jobs in the DB, as well as
// streaming new events and inserting into job DB and processing them.
// TODO(corver): Make this robust against chains not be available on startup.
func startProcessingEvents(
	ctx context.Context,
	network netconf.Network,
	xprov xchain.Provider,
	jobDB *job.DB,
	backends unibackend.Backends,
	solverKey *ecdsa.PrivateKey,
	addrs contracts.Addresses,
	cursors *cursors,
	pricer tokenpricer.Pricer,
	priceFunc priceFunc,
	inboxContracts map[uint64]*bindings.SolverNetInbox,
) error {
	solverAddr := ethcrypto.PubkeyToAddress(solverKey.PublicKey)
	ethBackends := backends.EVMBackends()

	outboxChains, err := detectContractChains(ctx, network, ethBackends, addrs.SolverNetOutbox)
	if err != nil {
		return errors.Wrap(err, "detect outbox chains")
	}

	outboxContracts := make(map[uint64]*bindings.SolverNetOutbox)
	for _, chain := range outboxChains {
		name := network.ChainName(chain)
		log.Debug(ctx, "Using outbox contract", "chain", name, "address", addrs.SolverNetOutbox.Hex())

		backend, err := ethBackends.Backend(chain)
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
			return "invalid"
		}

		// use last call target for target name
		call := fill.Calls[len(fill.Calls)-1]

		if tkn, ok := tokens.ByAddress(pendingData.DestinationChainID, call.Target); ok {
			return "ERC20:" + tkn.Symbol
		}

		// Native bridging has zero call data and positive value
		isNative := call.Selector == [4]byte{} && len(call.Params) == 0 && call.Value.Sign() > 0
		if nativeTkn, ok := tokens.Native(pendingData.DestinationChainID); ok && isNative {
			return "Native:" + nativeTkn.Symbol
		}

		if target, ok := targets.Get(pendingData.DestinationChainID, call.Target); ok {
			return target.Name
		}

		return "unknown"
	}

	procName := func(chainID uint64) string {
		return network.ChainVersionName(chainVerFromID(network.ID, chainID))
	}

	debugFunc := func(ctx context.Context, order Order, event Event) {
		debugPendingData(ctx, targetName, order, event)
		debugOrderPrice(ctx, priceFunc, order)
	}

	callAllower := newCallAllower(network.ID, addrs.SolverNetExecutor)

	ageCache := newAgeCache(backends)
	go monitorAgeCacheForever(ctx, network, ageCache)

	filledPnL := newFilledPnlFunc(pricer, targetName, network.ChainName, ageCache.InstrumentDestFilled)
	updatePnL := newUpdatePnLFunc(pricer, network.ChainName)

	deps := procDeps{
		GetOrder:          newOrderGetter(inboxContracts),
		ShouldReject:      newShouldRejector(backends, callAllower, priceFunc, solverAddr, addrs.SolverNetOutbox),
		DidFill:           newDidFiller(outboxContracts),
		Reject:            newRejector(inboxContracts, backends, solverAddr, updatePnL),
		Fill:              newFiller(outboxContracts, backends, solverAddr, addrs.SolverNetOutbox, filledPnL),
		Claim:             newClaimer(inboxContracts, backends, solverAddr, updatePnL),
		ChainName:         network.ChainName,
		ProcessorName:     procName,
		TargetName:        targetName,
		InstrumentAge:     ageCache.InstrumentAge,
		DebugPendingOrder: debugFunc,
	}

	// Create all event processing functions per EVM chain
	procs := make(map[uint64]eventProcFunc)
	for chainID := range inboxContracts {
		procs[chainID] = newEventProcFunc(deps, chainID)
	}

	// Create event processing functions for SVM chains
	for _, svmChain := range network.SVMChains() {
		backend, err := backends.Backend(svmChain.ID)
		if err != nil {
			return err
		}
		svmDeps := svmProcDeps(backend.SVMClient(), addrs.SolverNetOutbox, solverKey, deps)
		procs[svmChain.ID] = newEventProcFunc(svmDeps, svmChain.ID)
	}

	// Create the async worker function
	asyncWork := newAsyncWorkerFunc(jobDB, procs, network.ChainName)

	// Start all processing all existing jobs
	jobs, err := jobDB.All(ctx)
	if err != nil {
		return err
	}
	log.Debug(ctx, "Restarting existing jobs", "count", len(jobs))
	for _, j := range jobs {
		if err := asyncWork(ctx, j); err != nil {
			return err
		}
	}

	// Start streaming events for all chains
	for chainID := range inboxContracts {
		chainVer := chainVerFromID(network.ID, chainID)
		loopCtx := log.WithCtx(ctx, "chain_version", network.ChainVersionName(chainVer))
		go streamEventsForever(loopCtx, chainVer, xprov, cursors, addrs.SolverNetInbox, jobDB, asyncWork)
	}

	// Start streaming events for all SVM chains
	for _, svmChain := range network.SVMChains() {
		backend, err := backends.Backend(svmChain.ID)
		if err != nil {
			return err
		}
		chainVer := chainVerFromID(network.ID, svmChain.ID)
		loopCtx := log.WithCtx(ctx, "chain_version", network.ChainVersionName(chainVer))
		if err := maybeBootstrapSVMCursor(loopCtx, backend.SVMClient(), cursors, chainVer); err != nil {
			return errors.Wrap(err, "bootstrap SVM cursor")
		}
		go streamSVMForever(loopCtx, chainVer, backend.SVMClient(), cursors, jobDB, asyncWork)
	}

	return nil
}

// streamEventsForever streams events from the inbox contract on the given chain version.
func streamEventsForever(
	ctx context.Context,
	chainVer xchain.ChainVersion,
	xprov xchain.Provider,
	cursors *cursors,
	inboxAddr common.Address,
	jobDB *job.DB,
	asyncWork asyncWorkFunc,
) {
	backoff := expbackoff.New(ctx)
	for {
		from, ok, err := cursors.Get(ctx, chainVer)
		if !ok || err != nil {
			log.Warn(ctx, "Failed reading cursor (will retry)", err)
			backoff()

			continue
		}

		req := xchain.EventLogsReq{
			ChainID:         chainVer.ID,
			Height:          from, // Note the previous height is re-processed (idempotency FTW)
			ConfLevel:       chainVer.ConfLevel,
			FilterAddresses: []common.Address{inboxAddr},
			FilterTopics:    solvernet.AllEventTopics(),
		}
		err = xprov.StreamEventLogs(ctx, req, func(ctx context.Context, header *types.Header, elogs []types.Log) error {
			for _, elog := range elogs {
				// Insert each event/job into the jobDB, and start work async
				j, err := jobDB.InsertLog(ctx, chainVer.ID, elog)
				if err != nil {
					return err
				}

				err = asyncWork(ctx, j)
				if err != nil {
					return err
				}
			}

			return cursors.Set(ctx, chainVer, header.Number.Uint64())
		})
		if ctx.Err() != nil {
			return
		}

		log.Warn(ctx, "Failure processing inbox events (will retry)", err)
		backoff()
	}
}

// newCCTPClient creates a new CCTP client based on the network ID.
func newCCTPClient(networkID netconf.ID) cctp.Client {
	api := cctp.TestnetAPI
	if networkID == netconf.Mainnet {
		api = cctp.MainnetAPI
	}

	return cctp.NewClient(api)
}
