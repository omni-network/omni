package app

import (
	"context"

	"github.com/omni-network/omni/halo/attest"
	"github.com/omni-network/omni/halo/consensus"
	"github.com/omni-network/omni/lib/engine"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"

	abci "github.com/cometbft/cometbft/abci/types"
	cmtconfig "github.com/cometbft/cometbft/config"
	"github.com/cometbft/cometbft/node"
	"github.com/cometbft/cometbft/p2p"
	"github.com/cometbft/cometbft/privval"
	"github.com/cometbft/cometbft/proxy"
	cmttypes "github.com/cometbft/cometbft/types"
)

// Config defines al the halo run config.
type Config struct {
	EngineJWTFile           string
	AttestStateFile         string
	AppStateDir             string
	AppStatePersistInterval uint64
	Network                 netconf.Network
	Comet                   cmtconfig.Config
}

// Run runs the halo client.
func Run(ctx context.Context, config Config) error {
	// Load private validator key and state from disk (this hard exits on any error).
	privVal := privval.LoadFilePV(config.Comet.PrivValidatorKeyFile(), config.Comet.PrivValidatorStateFile())

	omniChain, ok := config.Network.OmniChain()
	if !ok {
		return errors.New("omni chain not found in network")
	}

	jwtBytes, err := engine.LoadJWTHexFile(config.EngineJWTFile)
	if err != nil {
		return errors.Wrap(err, "load engine JWT file")
	}

	ethCl, err := engine.NewClient(ctx, omniChain.RPCURL, jwtBytes)
	if err != nil {
		return errors.Wrap(err, "create engine client")
	}

	attState, err := attest.LoadState(config.AttestStateFile)
	if err != nil {
		return errors.Wrap(err, "load attest state")
	}

	var xprovider xchain.Provider
	// TODO(corver): Instantiate xprovider

	attSvc, err := attest.NewAttester(ctx, attState, privVal.Key.PrivKey, xprovider, config.Network.ChainIDs())
	if err != nil {
		return errors.Wrap(err, "create attester")
	}

	appState, err := consensus.LoadOrGenState(config.AppStateDir, config.AppStatePersistInterval)
	if err != nil {
		return errors.Wrap(err, "load or gen app state")
	}

	core := consensus.NewCore(ethCl, attSvc, appState)

	cmtNode, err := newCometNode(ctx, &config.Comet, core, privVal)
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
