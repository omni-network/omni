package app

import (
	"context"
	"net/http"
	"net/http/pprof"
	"time"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/halo/genutil/evm/predeploys"
	"github.com/omni-network/omni/lib/buildinfo"
	"github.com/omni-network/omni/lib/contracts"
	"github.com/omni-network/omni/lib/contracts/solvernet"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/expbackoff"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	tokenslib "github.com/omni-network/omni/lib/tokens"
	"github.com/omni-network/omni/lib/tokens/coingecko"
	"github.com/omni-network/omni/lib/tracer"
	"github.com/omni-network/omni/lib/umath"
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

	// if mainnet, just run monitoring and api (/live only)
	if cfg.Network == netconf.Mainnet {
		log.Info(ctx, "Serving API", "address", cfg.APIAddr)
		apiChan := serveAPI(cfg.APIAddr)

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

	portalReg, err := makePortalRegistry(cfg.Network, cfg.RPCEndpoints)
	if err != nil {
		return err
	}

	network, err := netconf.AwaitOnExecutionChain(ctx, cfg.Network, portalReg, cfg.RPCEndpoints.Keys())
	if err != nil {
		return err
	}

	// add back holesky manually on holesky on omega
	// temporary fix to enable holesky solving before it is re-enabled in core
	// It was disabled here https://github.com/omni-network/omni/pull/3259/files
	network = maybeAddHolesky(network)

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

	pricer := newPricer(ctx, cfg.CoinGeckoAPIKey)

	err = startEventStreams(ctx, network, xprov, backends, solverAddr, addrs, cursors, pricer)
	if err != nil {
		return errors.Wrap(err, "start event streams")
	}

	log.Info(ctx, "Serving API", "address", cfg.APIAddr)
	apiChan := serveAPI(cfg.APIAddr,
		newCheckHandler(newChecker(backends, solverAddr, addrs.SolverNetInbox, addrs.SolverNetOutbox)),
		newContractsHandler(addrs),
		newQuoteHandler(quoter),
	)

	if err := approveOutboxes(ctx, network, backends, solverAddr); err != nil {
		return errors.Wrap(err, "approve outboxes")
	}

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
	inboxTimestamps := make(map[uint64]func(uint64) time.Time)
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

		inboxTimestamps[chain] = func(height uint64) time.Time {
			header, err := backend.HeaderByNumber(ctx, umath.NewBigInt(height))
			if err != nil {
				return time.Time{} // Best effort, ignore for now.
			}
			timeI64, err := umath.ToInt64(header.Time)
			if err != nil {
				return time.Time{} // Best effort, ignore for now.
			}

			return time.Unix(timeI64, 0)
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

		// use last call target for target name
		call := fill.Calls[len(fill.Calls)-1]

		if tkn, ok := tokens.Find(o.DestinationChainID, call.Target); ok {
			return "ERC20:" + tkn.Symbol
		}

		// Native bridging has zero call data and positive value
		isNative := call.Selector == [4]byte{} && len(call.Params) == 0 && call.Value.Sign() > 0
		if nativeTkn, ok := tokens.Find(o.DestinationChainID, common.Address{}); ok && isNative {
			return "Native:" + nativeTkn.Symbol
		}

		return call.Target.Hex()[:7] // Short hex.
	}

	blockTimestamps := func(chainID uint64, height uint64) time.Time {
		f, ok := inboxTimestamps[chainID]
		if !ok {
			return time.Time{}
		}

		return f(height)
	}

	deps := procDeps{
		ParseID:        newIDParser(inboxContracts),
		GetOrder:       newOrderGetter(inboxContracts),
		ShouldReject:   newShouldRejector(backends, solverAddr, addrs.SolverNetOutbox),
		DidFill:        newDidFiller(outboxContracts),
		Reject:         newRejector(inboxContracts, backends, solverAddr),
		Fill:           newFiller(outboxContracts, backends, solverAddr, addrs.SolverNetOutbox, pricer),
		Claim:          newClaimer(inboxContracts, backends, solverAddr, pricer),
		SetCursor:      cursorSetter,
		ChainName:      network.ChainName,
		TargetName:     targetName,
		BlockTimestamp: blockTimestamps,
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

		var delay uint64
		// For L1 we delay the processing of events by two heights to minimize the risks
		// of the intent trasaction getting excluded during a reorg.
		if chainID == evmchain.IDEthereum {
			delay = 2
		}

		req := xchain.EventLogsReq{
			ChainID:       chainID,
			Height:        from, // Note the previous height is re-processed (idempotency FTW)
			ConfLevel:     confLevel,
			FilterAddress: inboxAddr,
			FilterTopics:  solvernet.AllEventTopics(),
			Delay:         delay,
		}
		err = xprov.StreamEventLogs(ctx, req, newEventProcessor(deps, chainID))
		if ctx.Err() != nil {
			return
		}

		log.Warn(ctx, "Failure streaming inbox events (will retry)", err)
		backoff()
	}
}

func maybeAddHolesky(network netconf.Network) netconf.Network {
	if network.ID != netconf.Omega {
		return network
	}

	// if holesky already exists, return
	for _, chain := range network.Chains {
		if chain.ID == evmchain.IDHolesky {
			return network
		}
	}

	// from omega netconf static
	deployHeight := 2130892
	portalAddr := common.HexToAddress("0xcB60A0451831E4865bC49f41F9C67665Fc9b75C3")

	// from e2e/types
	shards := []xchain.ShardID{xchain.ShardFinalized0, xchain.ShardLatest0}

	meta, ok := evmchain.MetadataByID(evmchain.IDHolesky)
	if !ok {
		// will not happen
		return network
	}

	network.Chains = append(network.Chains, netconf.Chain{
		ID:             evmchain.IDHolesky,
		Name:           meta.Name,
		PortalAddress:  portalAddr,
		DeployHeight:   uint64(deployHeight),
		BlockPeriod:    meta.BlockPeriod,
		Shards:         shards,
		AttestInterval: intervalFromPeriod(network.ID, meta.BlockPeriod),
	})

	return network
}

// from e2e/types testnet.go (temporary).
func intervalFromPeriod(network netconf.ID, period time.Duration) uint64 {
	target := time.Hour
	if network == netconf.Staging {
		target = time.Minute * 10
	} else if network == netconf.Devnet {
		target = time.Second * 10
	}

	if period == 0 {
		return 0
	}

	return uint64(target / period)
}
