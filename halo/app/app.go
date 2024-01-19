package app

import (
	"context"

	"github.com/omni-network/omni/halo/attest"
	"github.com/omni-network/omni/halo/comet"
	"github.com/omni-network/omni/lib/engine"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/xchain"

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
	// Load private validator key and state from disk (this hard exits on any error).
	privVal := privval.LoadFilePV(cfg.Comet.PrivValidatorKeyFile(), cfg.Comet.PrivValidatorStateFile())

	network, err := cfg.Network()
	if err != nil {
		return errors.Wrap(err, "load network")
	}

	omniChain, ok := network.OmniChain()
	if !ok {
		return errors.New("omni chain not found in network")
	}

	jwtBytes, err := engine.LoadJWTHexFile(cfg.EngineJWTFile)
	if err != nil {
		return errors.Wrap(err, "load engine JWT file")
	}

	ethCl, err := engine.NewClient(ctx, omniChain.RPCURL, jwtBytes)
	if err != nil {
		return errors.Wrap(err, "create engine client")
	}

	attState, err := attest.LoadState(cfg.AttestStateFile())
	if err != nil {
		return errors.Wrap(err, "load attest state")
	}

	var xprovider xchain.Provider
	// TODO(corver): Instantiate xprovider

	attSvc, err := attest.NewAttester(ctx, attState, privVal.Key.PrivKey, xprovider, network.ChainIDs())
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

	if err := cmtNode.Start(); err != nil {
		return errors.Wrap(err, "start comet node")
	}

	<-ctx.Done()

	if err := cmtNode.Stop(); err != nil {
		return errors.Wrap(err, "stop comet node")
	}

	return nil
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

	cmtNode, err := node.NewNode(config,
		privVal,
		nodeKey,
		proxy.NewLocalClientCreator(app),
		node.DefaultGenesisDocProviderFunc(config),
		cmtconfig.DefaultDBProvider,
		node.DefaultMetricsProvider(config.Instrumentation),
		newCmtLogger(ctx),
	)
	if err != nil {
		return nil, errors.Wrap(err, "create node")
	}

	return cmtNode, nil
}
