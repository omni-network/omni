package app

import (
	"context"
	"time"

	"github.com/omni-network/omni/halo/comet"
	halocfg "github.com/omni-network/omni/halo/config"
	"github.com/omni-network/omni/lib/buildinfo"
	cprovider "github.com/omni-network/omni/lib/cchain/provider"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/tracer"

	cmtcfg "github.com/cometbft/cometbft/config"
	"github.com/cometbft/cometbft/crypto"
	"github.com/cometbft/cometbft/node"
	"github.com/cometbft/cometbft/p2p"
	"github.com/cometbft/cometbft/proxy"
	rpclocal "github.com/cometbft/cometbft/rpc/client/local"
	cmttypes "github.com/cometbft/cometbft/types"

	"cosmossdk.io/store"
	pruningtypes "cosmossdk.io/store/pruning/types"
	"cosmossdk.io/store/snapshots"
	snapshottypes "cosmossdk.io/store/snapshots/types"
	dbm "github.com/cosmos/cosmos-db"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/server"
	sdktelemetry "github.com/cosmos/cosmos-sdk/telemetry"
	"github.com/cosmos/cosmos-sdk/types/mempool"
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
	log.Info(ctx, "Starting halo consensus client")

	if err := cfg.Verify(); err != nil {
		return nil, nil, errors.Wrap(err, "verify halo config")
	}

	buildinfo.Instrument(ctx)

	tracerIDs := tracer.Identifiers{Network: cfg.Network, Service: "halo", Instance: cfg.Comet.Moniker}
	stopTracer, err := tracer.Init(ctx, tracerIDs, cfg.Tracer)
	if err != nil {
		return nil, nil, err
	}

	if err := enableSDKTelemetry(); err != nil {
		return nil, nil, errors.Wrap(err, "enable cosmos-sdk telemetry")
	}

	privVal, err := loadPrivVal(cfg)
	if err != nil {
		return nil, nil, errors.Wrap(err, "load validator key")
	}

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

	//nolint:contextcheck // False positive
	app, err := newApp(
		newSDKLogger(ctx),
		db,
		engineCl,
		voter,
		netconf.ChainNamer(cfg.Network),
		baseAppOpts...,
	)
	if err != nil {
		return nil, nil, errors.Wrap(err, "create app")
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

	cProvider := cprovider.NewABCIProvider(rpcClient, cfg.Network, netconf.ChainNamer(cfg.Network))

	async := make(chan error, 1)
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
		)
		if err != nil {
			async <- err
		}
	}()

	log.Info(ctx, "Starting CometBFT", "listeners", cmtNode.Listeners())

	if err := cmtNode.Start(); err != nil {
		return nil, nil, errors.Wrap(err, "start comet node")
	}

	// Return async and stop functions.
	// Note that the original context used to start the app must be canceled first.
	// And a fresh context should be passed into the stop function.
	return async, func(ctx context.Context) error {
		voter.WaitDone()

		if err := cmtNode.Stop(); err != nil {
			return errors.Wrap(err, "stop comet node")
		}
		cmtNode.Wait()

		// Note that cometBFT doesn't shut down cleanly. It leaves a bunch of goroutines running...

		if err := stopTracer(ctx); err != nil {
			return errors.Wrap(err, "stop tracer")
		}

		return nil
	}, nil
}

func newCometNode(ctx context.Context, cfg *cmtcfg.Config, app *App, privVal cmttypes.PrivValidator,
) (*node.Node, error) {
	nodeKey, err := p2p.LoadOrGenNodeKey(cfg.NodeKeyFile())
	if err != nil {
		return nil, errors.Wrap(err, "load or gen node key", "key_file", cfg.NodeKeyFile())
	}

	cmtLog, err := newCmtLogger(ctx, cfg.LogLevel)
	if err != nil {
		return nil, err
	}

	wrapper := newABCIWrapper(
		server.NewCometABCIWrapper(app),
		app.EVMEngKeeper.Finalize,
	)

	cmtNode, err := node.NewNode(cfg,
		privVal,
		nodeKey,
		proxy.NewLocalClientCreator(wrapper),
		node.DefaultGenesisDocProviderFunc(cfg),
		cmtcfg.DefaultDBProvider,
		node.DefaultMetricsProvider(cfg.Instrumentation),
		cmtLog,
	)
	if err != nil {
		return nil, errors.Wrap(err, "create node")
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

	snapshotOptions := snapshottypes.NewSnapshotOptions(cfg.SnapshotInterval, uint32(cfg.SnapshotKeepRecent))

	pruneOpts := pruningtypes.NewPruningOptionsFromString(cfg.PruningOption)
	if cfg.PruningOption == pruningtypes.PruningOptionDefault {
		// Override the default cosmosSDK pruning values with much more aggressive defaults
		// since historical state isn't very important for most use-cases.
		pruneOpts = pruningtypes.NewCustomPruningOptions(defaultPruningKeep, defaultPruningInterval)
	}

	return []func(*baseapp.BaseApp){
		// baseapp.SetOptimisticExecution(), // TODO(corver): Enable this.
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
		return ethclient.NewEngineMock(ethclient.WithMockDeposit(pubkey, 1))
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
func enableSDKTelemetry() error {
	const farFuture = time.Hour * 24 * 365 * 10 // 10 years ~= infinity.

	_, err := sdktelemetry.New(sdktelemetry.Config{
		ServiceName:             "cosmos",
		Enabled:                 true,
		PrometheusRetentionTime: int64(farFuture.Seconds()), // Prometheus metrics never expire once created in-app.
	})
	if err != nil {
		return errors.Wrap(err, "enable cosmos-sdk telemetry")
	}

	return nil
}
