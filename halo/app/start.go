package app

import (
	"context"
	"encoding/hex"
	"os"
	"time"

	"github.com/omni-network/omni/halo/comet"
	halocfg "github.com/omni-network/omni/halo/config"
	"github.com/omni-network/omni/halo/genutil/genserve"
	"github.com/omni-network/omni/lib/buildinfo"
	"github.com/omni-network/omni/lib/cchain"
	cprovider "github.com/omni-network/omni/lib/cchain/provider"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/tracer"
	etypes "github.com/omni-network/omni/octane/evmengine/types"

	cmtcfg "github.com/cometbft/cometbft/config"
	"github.com/cometbft/cometbft/crypto"
	"github.com/cometbft/cometbft/node"
	"github.com/cometbft/cometbft/p2p"
	"github.com/cometbft/cometbft/privval"
	"github.com/cometbft/cometbft/proxy"
	rpclocal "github.com/cometbft/cometbft/rpc/client/local"
	cmttypes "github.com/cometbft/cometbft/types"

	"github.com/ethereum/go-ethereum/common"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"

	sdklog "cosmossdk.io/log"
	"cosmossdk.io/store"
	pruningtypes "cosmossdk.io/store/pruning/types"
	"cosmossdk.io/store/snapshots"
	snapshottypes "cosmossdk.io/store/snapshots/types"
	storetypes "cosmossdk.io/store/types"
	dbm "github.com/cosmos/cosmos-db"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	sdkflags "github.com/cosmos/cosmos-sdk/client/flags"
	sdkserver "github.com/cosmos/cosmos-sdk/server"
	"github.com/cosmos/cosmos-sdk/server/api"
	"github.com/cosmos/cosmos-sdk/server/grpc"
	sdkservertypes "github.com/cosmos/cosmos-sdk/server/types"
	sdktelemetry "github.com/cosmos/cosmos-sdk/telemetry"
	"github.com/cosmos/cosmos-sdk/types/mempool"
	grpc1 "github.com/cosmos/gogoproto/grpc"
)

// Config wraps the halo (app) and comet (client) configurations.
type Config struct {
	halocfg.Config
	Comet cmtcfg.Config
}

// BackendType returns the halo config backend type
// or the comet backend type otherwise.
func (c Config) BackendType() dbm.BackendType {
	if c.Config.BackendType == "" {
		return dbm.BackendType(c.Comet.DBBackend)
	}

	return dbm.BackendType(c.Config.BackendType)
}

// Run runs the halo client until the context is canceled.
//
//nolint:contextcheck // Explicit new stop context.
func Run(ctx context.Context, cfg Config) error {
	async, stopFunc, err := Start(ctx, cfg)
	if err != nil {
		return err
	}

	select {
	case <-ctx.Done():
		log.Info(ctx, "Shutdown detected, stopping...")
	case err := <-async:
		return err
	}

	// Use a fresh context for stopping (only allow 5 seconds).
	stopCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return stopFunc(stopCtx)
}

// Start starts the halo client returning a stop function or an error.
//
// Note that the original context used to start the app must be canceled first
// before calling the stop function and a fresh context should be passed into the stop function.
func Start(ctx context.Context, cfg Config) (<-chan error, func(context.Context) error, error) {
	log.Info(ctx, "Starting halo consensus client", "moniker", cfg.Comet.Moniker)

	if err := cfg.Verify(); err != nil {
		return nil, nil, errors.Wrap(err, "verify halo config")
	}

	buildinfo.Instrument(ctx)

	tracerIDs := tracer.Identifiers{Network: cfg.Network, Service: "halo", Instance: cfg.Comet.Moniker}
	stopTracer, err := tracer.Init(ctx, tracerIDs, cfg.Tracer)
	if err != nil {
		return nil, nil, err
	}

	metrics, err := enableSDKTelemetry(cfg.Network)
	if err != nil {
		return nil, nil, errors.Wrap(err, "enable cosmos-sdk telemetry")
	}

	privVal, err := loadPrivVal(cfg)
	if err != nil {
		return nil, nil, errors.Wrap(err, "load validator key")
	}
	logPrivVal(ctx, privVal)

	db, err := dbm.NewDB("application", cfg.BackendType(), cfg.DataDir())
	if err != nil {
		return nil, nil, errors.Wrap(err, "create db")
	}

	baseAppOpts, err := makeBaseAppOpts(cfg)
	if err != nil {
		return nil, nil, errors.Wrap(err, "make base app opts")
	}

	engineCl, err := newEngineClient(ctx, cfg, cfg.Network, privVal.Key.PubKey)
	if err != nil {
		return nil, nil, err
	}

	voter, err := newVoterLoader(privVal.Key.PrivKey) // Construct a lazy voter loader
	if err != nil {
		return nil, nil, err
	}

	sdkLogger := newSDKLogger(ctx)
	asyncAbort := make(chan error, 1) // Allows async processes to abort the app

	//nolint:contextcheck // False positive
	app, err := newApp(
		sdkLogger,
		db,
		engineCl,
		voter,
		netconf.ChainVersionNamer(cfg.Network),
		netconf.ChainNamer(cfg.Network),
		burnEVMFees{},
		serverAppOptsFromCfg(cfg),
		asyncAbort,
		baseAppOpts...,
	)
	if err != nil {
		return nil, nil, errors.Wrap(err, "create app")
	}

	if err := registerGenesisServer(ctx, app.GRPCQueryRouter(), cfg); err != nil {
		return nil, nil, err
	}

	app.EVMEngKeeper.SetBuildDelay(cfg.EVMBuildDelay)
	app.EVMEngKeeper.SetBuildOptimistic(cfg.EVMBuildOptimistic)

	cmtNode, err := newCometNode(ctx, &cfg.Comet, app, privVal)
	if err != nil {
		return nil, nil, errors.Wrap(err, "create comet node")
	}

	rpcClient := rpclocal.New(cmtNode)
	cmtAPI := comet.NewAPI(rpcClient)
	app.SetCometAPI(cmtAPI)

	clientCtx := app.ClientContext(ctx).WithClient(rpcClient).WithHomeDir(cfg.HomeDir)
	if err := startRPCServers(ctx, cfg, app, sdkLogger, metrics, asyncAbort, clientCtx); err != nil {
		return nil, nil, err
	}

	cProvider, err := newCProvider(rpcClient, cfg)
	if err != nil {
		return nil, nil, err
	}

	go func() {
		err := voter.LazyLoad(
			ctx,
			cfg.Network,
			engineCl,
			cfg.RPCEndpoints,
			cProvider,
			privVal.Key.PrivKey,
			cfg.VoterStateFile(),
			cmtAPI,
			asyncAbort,
		)
		if err != nil {
			asyncAbort <- err
		}
	}()

	log.Info(ctx, "Starting CometBFT")

	if err := cmtNode.Start(); err != nil {
		return nil, nil, errors.Wrap(err, "start comet node")
	}

	status := new(readinessStatus)
	go instrumentReadiness(ctx, status)

	stopMonitoringAPI := startMonitoringAPI(&cfg.Comet, asyncAbort, status)

	go monitorCometForever(ctx, cfg.Network, rpcClient, cmtNode.ConsensusReactor().WaitSync, cfg.DataDir(), status)
	go monitorEVMForever(ctx, cfg, engineCl, status)

	// Return asyncAbort and stop functions.
	// Note that the original context used to start the app must be canceled first.
	// And a fresh context should be passed into the stop function.
	return asyncAbort, func(ctx context.Context) error {
		voter.WaitDone()

		if err := cmtNode.Stop(); err != nil {
			return errors.Wrap(err, "stop comet node")
		}
		cmtNode.Wait()

		// Note that cometBFT doesn't shut down cleanly. It leaves a bunch of goroutines running...

		if err := stopMonitoringAPI(ctx); err != nil {
			return errors.Wrap(err, "stop monitoring API")
		}

		if err := stopTracer(ctx); err != nil {
			return errors.Wrap(err, "stop tracer")
		}

		log.Info(ctx, "Halo consensus client stopped")

		return nil
	}, nil
}

// newCProvider returns a new cchain provider. Either GRPC if enabled since it is faster,
// otherwise the ABCI provider.
//

func newCProvider(rpcClient *rpclocal.Local, cfg Config) (cchain.Provider, error) {
	if cfg.SDKGRPC.Enable {
		return cprovider.NewGRPC(cfg.SDKGRPC.Address, cfg.Network)
	}

	return cprovider.NewABCI(rpcClient, cfg.Network), nil
}

// startRPCServers starts the Cosmos REST and gRPC servers.
func startRPCServers(
	ctx context.Context,
	cfg Config,
	app *App,
	logger sdklog.Logger,
	metrics *sdktelemetry.Metrics,
	async chan error,
	clientCtx client.Context,
) error {
	app.RegisterTendermintService(clientCtx)
	app.RegisterNodeService(clientCtx, cfg.SDKRPCConfig())

	rpcCfg := cfg.SDKRPCConfig()

	grpcSrv, err := grpc.NewGRPCServer(clientCtx, app, rpcCfg.GRPC)
	if err != nil {
		return errors.Wrap(err, "new grpc server")
	}

	apiSrv := api.New(clientCtx, logger.With("module", "api-server"), grpcSrv)
	apiSrv.SetTelemetry(metrics)
	app.RegisterAPIRoutes(apiSrv, rpcCfg.API)

	if cfg.SDKAPI.Enable {
		go func() {
			async <- apiSrv.Start(ctx, cfg.SDKRPCConfig())
		}()
	}

	if cfg.SDKGRPC.Enable {
		go func() {
			async <- grpc.StartGRPCServer(ctx, logger.With("module", "grpc-server"), rpcCfg.GRPC, grpcSrv)
		}()
	}

	return nil
}

func newCometNode(ctx context.Context, cfg *cmtcfg.Config, app *App, privVal cmttypes.PrivValidator) (
	*node.Node, error) {
	nodeKey, err := p2p.LoadOrGenNodeKey(cfg.NodeKeyFile())
	if err != nil {
		return nil, errors.Wrap(err, "load or gen node key", "key_file", cfg.NodeKeyFile())
	}

	cmtLog, err := NewCmtLogger(ctx, cfg.LogLevel)
	if err != nil {
		return nil, err
	}

	wrapper := newABCIWrapper(
		sdkserver.NewCometABCIWrapper(app),
		app.EVMEngKeeper.PostFinalize,
		func() storetypes.CacheMultiStore {
			return app.CommitMultiStore().CacheMultiStore()
		},
	)

	// Configure CometBFT prometheus metrics as per provided config
	metrics := node.DefaultMetricsProvider(&cmtcfg.InstrumentationConfig{
		Prometheus: cfg.Instrumentation.Prometheus,
		Namespace:  cfg.Instrumentation.Namespace,
	})
	// But don't instantiate the CometBFT prometheus server, we do it startMonitoringAPI.
	cfg.Instrumentation.Prometheus = false

	cmtNode, err := node.NewNode(cfg,
		privVal,
		nodeKey,
		proxy.NewLocalClientCreator(wrapper),
		node.DefaultGenesisDocProviderFunc(cfg),
		cmtcfg.DefaultDBProvider,
		metrics,
		cmtLog,
	)
	if err != nil {
		return nil, errors.Wrap(err, "create comet node")
	}

	return cmtNode, nil
}

func makeBaseAppOpts(cfg Config) ([]func(*baseapp.BaseApp), error) {
	chainID, err := chainIDFromGenesis(cfg)
	if err != nil {
		return nil, err
	}

	snapshotStore, err := newSnapshotStore(cfg)
	if err != nil {
		return nil, err
	}

	snapshotOptions := snapshottypes.NewSnapshotOptions(cfg.SnapshotInterval, cfg.SnapshotKeepRecent)

	pruneOpts := pruningtypes.NewPruningOptionsFromString(cfg.PruningOption)
	if cfg.PruningOption == pruningtypes.PruningOptionDefault {
		// We interpret "default" to be PruningEverything, since historical state isn't very important.
		pruneOpts = pruningtypes.NewPruningOptions(pruningtypes.PruningEverything)
	}

	return []func(*baseapp.BaseApp){
		// baseapp.SetOptimisticExecution(), // Octane doesn't support this.
		baseapp.SetChainID(chainID),
		baseapp.SetMinRetainBlocks(cfg.MinRetainBlocks),
		baseapp.SetPruning(pruneOpts),
		baseapp.SetInterBlockCache(store.NewCommitKVStoreCacheManager()),
		baseapp.SetSnapshot(snapshotStore, snapshotOptions),
		baseapp.SetMempool(mempool.NoOpMempool{}),
	}, nil
}

func newSnapshotStore(cfg Config) (*snapshots.Store, error) {
	db, err := dbm.NewDB("metadata", cfg.BackendType(), cfg.SnapshotDir())
	if err != nil {
		return nil, errors.Wrap(err, "create snapshot db")
	}

	ss, err := snapshots.NewStore(db, cfg.SnapshotDir())
	if err != nil {
		return nil, errors.Wrap(err, "create snapshot store")
	}

	return ss, nil
}

func chainIDFromGenesis(cfg Config) (string, error) {
	genDoc, err := node.DefaultGenesisDocProviderFunc(&cfg.Comet)()
	if err != nil {
		return "", errors.Wrap(err, "load genesis doc")
	}

	return genDoc.ChainID, nil
}

// newEngineClient returns a new engine API client.
func newEngineClient(ctx context.Context, cfg Config, network netconf.ID, pubkey crypto.PubKey) (ethclient.EngineClient, error) {
	if network == netconf.Simnet {
		return ethclient.NewEngineMock(
			ethclient.WithPortalRegister(netconf.SimnetNetwork()),
			ethclient.WithFarFutureUpgradePlan(),
			ethclient.WithMockSelfDelegation(pubkey, 1),
		)
	}

	jwtBytes, err := ethclient.LoadJWTHexFile(cfg.EngineJWTFile)
	if err != nil {
		return nil, errors.Wrap(err, "load engine JWT file")
	}

	engineCl, err := ethclient.NewAuthClient(ctx, cfg.EngineEndpoint, jwtBytes)
	if err != nil {
		return nil, errors.Wrap(err, "create engine client")
	}

	return engineCl, nil
}

// enableSDKTelemetry enables prometheus based cosmos-sdk telemetry.
func enableSDKTelemetry(id netconf.ID) (*sdktelemetry.Metrics, error) {
	// Skip telemetry for simnet, because it uses globals which conflict when running tests in parallel.
	if id == netconf.Simnet {
		return nil, nil //nolint:nilnil // Not a problem
	}

	const farFuture = time.Hour * 24 * 365 * 10 // 10 years ~= infinity.

	m, err := sdktelemetry.New(sdktelemetry.Config{
		ServiceName:             "cosmos",
		Enabled:                 true,
		PrometheusRetentionTime: int64(farFuture.Seconds()), // Prometheus metrics never expire once created in-app.
	})
	if err != nil {
		return nil, errors.Wrap(err, "enable cosmos-sdk telemetry")
	}

	return m, nil
}

var (
	_ etypes.FeeRecipientProvider = &burnEVMFees{}

	// burnAddress is used as the EVM fee recipient resulting in burned execution fees.
	burnAddress = common.HexToAddress("0x000000000000000000000000000000000000dEaD")
)

// burnEVMFees is a fee recipient provider that burns all execution fees.
type burnEVMFees struct{}

func (burnEVMFees) LocalFeeRecipient() common.Address {
	return burnAddress
}

func (burnEVMFees) VerifyFeeRecipient(address common.Address) error {
	if address != burnAddress {
		return errors.New("fee recipient not the burn address", "addr", address.Hex())
	}

	return nil
}

// registerGenesisServer registers a custom non-cosmos-sdk-module grpc query server that serves the consensus and execution layer genesis files.
// This enables a trusted way to join an ephemeral network without well-known long-lived genesis files.
func registerGenesisServer(ctx context.Context, s grpc1.Server, cfg Config) error {
	if !cfg.Network.IsEphemeral() {
		// Don't do this for protected networks since execution genesis is very big.
		return nil
	}

	consensus, err := os.ReadFile(cfg.Comet.GenesisFile())
	if err != nil {
		return errors.Wrap(err, "read consensus genesis file") // This is expected to succeed
	}

	execution, err := os.ReadFile(cfg.ExecutionGenesisFile())
	if os.IsNotExist(err) {
		// This is optional feature, so not an error.
		log.Info(ctx, "Not serving execution_genesis.json; file not in config folder")
	} else if err != nil {
		log.Warn(ctx, "Not serving execution_genesis.json file", err)
	}

	genserve.Register(s, execution, consensus)

	return nil
}

var _ sdkservertypes.AppOptions = serverAppOpts{}

// serverAppOpts implements the cosmos-sdk server app options interface.
type serverAppOpts map[string]any

func (o serverAppOpts) Get(key string) any {
	return o[key]
}

// serverAppOptsFromCfg returns the cosmos-sdk server app options from the given config.
// This is required by the upgrade module.
func serverAppOptsFromCfg(cfg Config) serverAppOpts {
	return serverAppOpts{
		sdkflags.FlagHome:                cfg.HomeDir,
		sdkserver.FlagUnsafeSkipUpgrades: cfg.UnsafeSkipUpgrades,
	}
}

// logPrivVal logs the private validator key details.
func logPrivVal(ctx context.Context, privVal *privval.FilePV) {
	pk := privVal.Key.PubKey
	ethPK, err := ethcrypto.DecompressPubkey(pk.Bytes())
	if err != nil {
		return
	}

	log.Info(ctx, "Loaded consensus private validator key from disk",
		"pubkey", hex.EncodeToString(pk.Bytes()),
		"comet_addr", pk.Address(),
		"eth_addr", ethcrypto.PubkeyToAddress(*ethPK))
}
