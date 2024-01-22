package app

import (
	"context"
	"time"

	"github.com/omni-network/omni/halo/attest"
	"github.com/omni-network/omni/halo/comet"
	"github.com/omni-network/omni/lib/engine"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/gitinfo"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"
	"github.com/omni-network/omni/lib/xchain/provider"

	abci "github.com/cometbft/cometbft/abci/types"
	cmtconfig "github.com/cometbft/cometbft/config"
	"github.com/cometbft/cometbft/node"
	"github.com/cometbft/cometbft/p2p"
	"github.com/cometbft/cometbft/privval"
	"github.com/cometbft/cometbft/proxy"
	cmttypes "github.com/cometbft/cometbft/types"
)

type Config struct {
	HaloConfig
	Comet cmtconfig.Config
}

// Run runs the halo client.
func Run(ctx context.Context, cfg Config) error {
	log.Info(ctx, "Starting halo consensus client")

	commit, timestamp := gitinfo.Get()
	log.Info(ctx, "Version info", "git_commit", commit, "git_timestamp", timestamp)

	// Load private validator key and state from disk (this hard exits on any error).
	privVal := privval.LoadFilePV(cfg.Comet.PrivValidatorKeyFile(), cfg.Comet.PrivValidatorStateFile())

	network, err := netconf.Load(cfg.NetworkFile())
	if err != nil {
		return errors.Wrap(err, "load network")
	} else if err := network.Validate(); err != nil {
		return errors.Wrap(err, "validate network configuration")
	}

	ethCl, err := newEthCl(ctx, cfg.HaloConfig, network)
	if err != nil {
		return err
	}

	xprovider, err := newXProvider(network)
	if err != nil {
		return errors.Wrap(err, "create xchain provider")
	}

	attSvc, err := attest.LoadAttester(ctx, privVal.Key.PrivKey, cfg.AttestStateFile(), xprovider, network.ChainIDs())
	if err != nil {
		return errors.Wrap(err, "create attester")
	}

	appState, err := comet.LoadOrGenState(cfg.AppStateDir(), cfg.AppStatePersistInterval)
	if err != nil {
		return errors.Wrap(err, "load or gen app state")
	}

	snapshotStore, err := comet.NewSnapshotStore(cfg.SnapshotDir())
	if err != nil {
		return errors.Wrap(err, "create snapshot store")
	}

	app := comet.NewApp(ethCl, attSvc, appState, snapshotStore, cfg.SnapshotInterval)

	cmtNode, err := newCometNode(ctx, &cfg.Comet, app, privVal)
	if err != nil {
		return errors.Wrap(err, "create comet node")
	}

	log.Info(ctx, "Starting CometBFT", "listeners", cmtNode.Listeners())

	if err := cmtNode.Start(); err != nil {
		return errors.Wrap(err, "start comet node")
	}

	if err := maybeSetupSimnetRelayer(ctx, network, app, xprovider); err != nil {
		return errors.Wrap(err, "setup simnet relayer")
	}

	<-ctx.Done()
	log.Info(ctx, "Shutdown detected, stopping...")

	if err := cmtNode.Stop(); err != nil {
		return errors.Wrap(err, "stop comet node")
	}

	return nil
}

// newXProvider returns a new xchain provider.
func newXProvider(network netconf.Network) (xchain.Provider, error) {
	if network.Name == netconf.Simnet {
		return provider.NewMock(time.Millisecond * 750), nil // Slightly faster than our chain.
	}

	return nil, errors.New("only simnet is supported at this point")
}

// newEthCl returns a new engine API client.
func newEthCl(ctx context.Context, cfg HaloConfig, network netconf.Network) (engine.API, error) {
	if network.Name == netconf.Simnet {
		return engine.NewMock()
	}

	jwtBytes, err := engine.LoadJWTHexFile(cfg.EngineJWTFile)
	if err != nil {
		return nil, errors.Wrap(err, "load engine JWT file")
	}

	omniChain, ok := network.OmniChain()
	if !ok {
		return nil, errors.New("omni chain not found in network")
	}

	ethCl, err := engine.NewClient(ctx, omniChain.RPCURL, jwtBytes)
	if err != nil {
		return nil, errors.Wrap(err, "create engine client")
	}

	return ethCl, nil
}

func newCometNode(ctx context.Context, config *cmtconfig.Config, app abci.Application, privVal cmttypes.PrivValidator,
) (*node.Node, error) {
	nodeKey, err := p2p.LoadOrGenNodeKey(config.NodeKeyFile())
	if err != nil {
		return nil, errors.Wrap(err, "load or gen node key", "key_file", config.NodeKeyFile())
	}

	// TxIndex config always disabled
	config.TxIndex = &cmtconfig.TxIndexConfig{
		Indexer: "null",
	}

	cmtLog, err := newCmtLogger(ctx, config.LogLevel)
	if err != nil {
		return nil, err
	}

	cmtNode, err := node.NewNode(config,
		privVal,
		nodeKey,
		proxy.NewLocalClientCreator(app),
		node.DefaultGenesisDocProviderFunc(config),
		cmtconfig.DefaultDBProvider,
		node.DefaultMetricsProvider(config.Instrumentation),
		cmtLog,
	)
	if err != nil {
		return nil, errors.Wrap(err, "create node")
	}

	return cmtNode, nil
}
