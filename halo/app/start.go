package app

import (
	"context"
	"time"

	atypes "github.com/omni-network/omni/halo/attest/types"
	"github.com/omni-network/omni/halo/attest/voter"
	"github.com/omni-network/omni/halo/comet"
	halocfg "github.com/omni-network/omni/halo/config"
	"github.com/omni-network/omni/lib/buildinfo"
	"github.com/omni-network/omni/lib/cchain"
	cprovider "github.com/omni-network/omni/lib/cchain/provider"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/tracer"
	"github.com/omni-network/omni/lib/xchain"
	xprovider "github.com/omni-network/omni/lib/xchain/provider"

	cmtcfg "github.com/cometbft/cometbft/config"
	"github.com/cometbft/cometbft/crypto"
	"github.com/cometbft/cometbft/node"
	"github.com/cometbft/cometbft/p2p"
	"github.com/cometbft/cometbft/privval"
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
	stopFunc, err := Start(ctx, cfg)
	if err != nil {
		return err
	}

	<-ctx.Done()
	log.Info(ctx, "Shutdown detected, stopping...")

	// Use a fresh context for stopping (only allow 5 seconds).
	stopCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return stopFunc(stopCtx)
}

// Start starts the halo client returning a stop function or an error.
//
// Note that the original context used to start the app must be canceled first
// before calling the stop function and a fresh context should be passed into the stop function.
func Start(ctx context.Context, cfg Config) (func(context.Context) error, error) {
	log.Info(ctx, "Starting halo consensus client")

	buildinfo.Instrument(ctx)

	network, err := netconf.Load(cfg.NetworkFile())
	if err != nil {
		return nil, errors.Wrap(err, "load network")
	} else if err := network.Validate(); err != nil {
		return nil, errors.Wrap(err, "validate network configuration")
	}

	tracerIDs := tracer.Identifiers{Network: network.ID, Service: "halo", Instance: cfg.Comet.Moniker}
	stopTracer, err := tracer.Init(ctx, tracerIDs, cfg.Tracer)
	if err != nil {
		return nil, err
	}

	if err := enableSDKTelemetry(); err != nil {
		return nil, errors.Wrap(err, "enable cosmos-sdk telemetry")
	}

	privVal, err := loadPrivVal(cfg)
	if err != nil {
		return nil, errors.Wrap(err, "load validator key")
	}

	db, err := dbm.NewDB("application", cfg.BackendType(), cfg.DataDir())
	if err != nil {
		return nil, errors.Wrap(err, "create db")
	}

	baseAppOpts, err := makeBaseAppOpts(cfg)
	if err != nil {
		return nil, errors.Wrap(err, "make base app opts")
	}

	engineCl, err := newEngineClient(ctx, cfg, network, privVal.Key.PubKey)
	if err != nil {
		return nil, err
	}

	//nolint:contextcheck // False positive
	app, err := newApp(
		newSDKLogger(ctx),
		db,
		engineCl,
		network.ChainName,
		baseAppOpts...,
	)
	if err != nil {
		return nil, errors.Wrap(err, "create app")
	}

	app.EVMEngKeeper.SetBuildDelay(cfg.EVMBuildDelay)
	app.EVMEngKeeper.SetBuildOptimistic(cfg.EVMBuildOptimistic)

	cmtNode, err := newCometNode(ctx, &cfg.Comet, app, privVal)
	if err != nil {
		return nil, errors.Wrap(err, "create comet node")
	}

	rpcClient := rpclocal.New(cmtNode)
	cmtAPI := comet.NewAPI(rpcClient)
	app.SetCometAPI(cmtAPI)

	cProvider := cprovider.NewABCIProvider(rpcClient, network.ID, network.ChainNamesByIDs())

	xProvider, err := newXProvider(network, cProvider)
	if err != nil {
		return nil, errors.Wrap(err, "create xchain provider")
	}

	voter, err := newVoter(cmtAPI, network, privVal, xProvider, cProvider, cfg.VoterStateFile())
	if err != nil {
		return nil, errors.Wrap(err, "create voter")
	}
	app.SetVoter(voter)

	voter.Start(ctx)

	log.Info(ctx, "Starting CometBFT", "listeners", cmtNode.Listeners())

	if err := cmtNode.Start(); err != nil {
		return nil, errors.Wrap(err, "start comet node")
	}

	// Return stop function.
	// Note that the original context uesd to start the app must be canceled first.
	// And a fresh context should be passed into the stop function.
	return func(ctx context.Context) error {
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

// newXProvider returns a new xchain provider.
func newXProvider(network netconf.Network, cProvider cchain.Provider) (xchain.Provider, error) {
	if network.ID == netconf.Simnet {
		omni, ok := network.OmniConsensusChain()
		if !ok {
			return nil, errors.New("omni chain not found in network")
		}

		return xprovider.NewMock(omni.BlockPeriod*8/10, omni.ID, cProvider), nil // Slightly faster than our chain.
	}

	clients := make(map[uint64]ethclient.Client)
	for _, chain := range network.EVMChains() {
		ethCl, err := ethclient.Dial(chain.Name, chain.RPCURL)
		if err != nil {
			return nil, errors.Wrap(err, "dial chain",
				"name", chain.Name,
				"id", chain.ID,
				"rpc_url", chain.RPCURL,
			)
		}
		clients[chain.ID] = ethCl
	}

	return xprovider.New(network, clients, cProvider), nil
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

	cmtNode, err := node.NewNode(cfg,
		privVal,
		nodeKey,
		proxy.NewLocalClientCreator(loggingABCIApp{server.NewCometABCIWrapper(app)}),
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

	return []func(*baseapp.BaseApp){
		baseapp.SetChainID(chainID),
		baseapp.SetMinRetainBlocks(cfg.MinRetainBlocks),
		baseapp.SetPruning(pruningtypes.NewPruningOptionsFromString(cfg.PruningOption)),
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
func newEngineClient(ctx context.Context, cfg Config, network netconf.Network, pubkey crypto.PubKey) (ethclient.EngineClient, error) {
	if network.ID == netconf.Simnet {
		return ethclient.NewEngineMock(ethclient.WithMockDeposit(pubkey, 1))
	}

	jwtBytes, err := ethclient.LoadJWTHexFile(cfg.EngineJWTFile)
	if err != nil {
		return nil, errors.Wrap(err, "load engine JWT file")
	}

	omniChain, ok := network.OmniEVMChain()
	if !ok {
		return nil, errors.New("omni chain not found in network")
	}

	engineCl, err := ethclient.NewAuthClient(ctx, omniChain.AuthRPCURL, jwtBytes)
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

func newVoter(
	cmtAPI comet.API,
	network netconf.Network,
	privVal *privval.FilePV,
	xprovider xchain.Provider,
	cProvider cchain.Provider,
	stateFile string,
) (*voter.Voter, error) {
	deps := voteDeps{
		API:      cmtAPI,
		Provider: cProvider,
	}
	voterI, err := voter.LoadVoter(privVal.Key.PrivKey, stateFile, xprovider, deps, network.ChainNamesByIDs())
	if err != nil {
		return nil, errors.Wrap(err, "create voter")
	}

	return voterI, nil
}

var _ atypes.VoterDeps = voteDeps{}

type voteDeps struct {
	comet.API
	cchain.Provider
}

func (v voteDeps) LatestAttestationHeight(ctx context.Context, chainID uint64) (uint64, bool, error) {
	att, ok, err := v.LatestAttestation(ctx, chainID)
	if err != nil {
		return 0, false, err
	} else if !ok {
		return 0, false, nil
	}

	return att.BlockHeight, true, nil
}
