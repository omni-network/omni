package ethbackend

import (
	"context"
	"crypto/ecdsa"
	"time"

	"github.com/omni-network/omni/e2e/types"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/fireblocks"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/txmgr"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

const (
	interval = 3
)

// Backends is a wrapper around a set of Backends, one for each chain.
// At this point, it only supports "a single account for all Backends".
//
// See Backends godoc for more information.
type Backends struct {
	backends map[uint64]*Backend
}

// NewFireBackends returns a multi-backends backed by fireblocks keys that supports configured all chains.
func NewFireBackends(ctx context.Context, testnet types.Testnet, fireCl fireblocks.Client) (Backends, error) {
	inner := make(map[uint64]*Backend)

	// Configure omni EVM Backend
	if testnet.HasOmniEVM() {
		chain := testnet.BroadcastOmniEVM()
		ethCl, err := ethclient.Dial(chain.Chain.Name, chain.ExternalRPC)
		if err != nil {
			return Backends{}, errors.Wrap(err, "dial")
		}

		inner[chain.Chain.ChainID], err = NewFireBackend(ctx, chain.Chain.Name, chain.Chain.ChainID, chain.Chain.BlockPeriod, ethCl, fireCl)
		if err != nil {
			return Backends{}, errors.Wrap(err, "new omni Backend")
		}
	}

	// Configure anvil EVM Backends
	for _, chain := range testnet.AnvilChains {
		ethCl, err := ethclient.Dial(chain.Chain.Name, chain.ExternalRPC)
		if err != nil {
			return Backends{}, errors.Wrap(err, "dial")
		}

		inner[chain.Chain.ChainID], err = NewFireBackend(ctx, chain.Chain.Name, chain.Chain.ChainID, chain.Chain.BlockPeriod, ethCl, fireCl)
		if err != nil {
			return Backends{}, errors.Wrap(err, "new anvil Backend")
		}
	}

	// Configure public EVM Backends
	for _, chain := range testnet.PublicChains {
		ethCl, err := ethclient.Dial(chain.Chain().Name, chain.NextRPCAddress())
		if err != nil {
			return Backends{}, errors.Wrap(err, "dial")
		}

		backend, err := NewFireBackend(ctx, chain.Chain().Name, chain.Chain().ChainID, chain.Chain().BlockPeriod, ethCl, fireCl)
		if err != nil {
			return Backends{}, errors.Wrap(err, "new public Backend")
		} else if err := backend.EnsureSynced(ctx); err != nil {
			return Backends{}, errors.Wrap(err, "ensure public chain synced", "chain", chain.Chain().Name)
		}

		inner[chain.Chain().ChainID] = backend
	}

	return Backends{
		backends: inner,
	}, nil
}

// NewBackends returns a multi-backends backed by in-memory keys that supports configured all chains.
func NewBackends(ctx context.Context, testnet types.Testnet, deployKeyFile string) (Backends, error) {
	var err error

	var publicDeployKey *ecdsa.PrivateKey
	if testnet.Network == netconf.Devnet {
		if deployKeyFile != "" {
			return Backends{}, errors.New("deploy key not supported in devnet")
		}
	} else if testnet.Network == netconf.Staging {
		publicDeployKey, err = crypto.LoadECDSA(deployKeyFile)
	} else {
		return Backends{}, errors.New("unknown network")
	}
	if err != nil {
		return Backends{}, errors.Wrap(err, "load deploy key")
	}

	inner := make(map[uint64]*Backend)

	// Configure omni EVM Backend
	{
		chain := testnet.BroadcastOmniEVM()
		ethCl, err := ethclient.Dial(chain.Chain.Name, chain.ExternalRPC)
		if err != nil {
			return Backends{}, errors.Wrap(err, "dial")
		}

		// dev omni evm uses same dev accounts as anvil
		// TODO: do not use dev anvil backend for prod omni evms
		backend, err := NewAnvilBackend(chain.Chain.Name, chain.Chain.ChainID, chain.Chain.BlockPeriod, ethCl)
		if err != nil {
			return Backends{}, errors.Wrap(err, "new omni Backend")
		}

		inner[chain.Chain.ChainID] = backend
	}

	// Configure anvil EVM Backends
	for _, chain := range testnet.AnvilChains {
		ethCl, err := ethclient.Dial(chain.Chain.Name, chain.ExternalRPC)
		if err != nil {
			return Backends{}, errors.Wrap(err, "dial")
		}

		backend, err := NewAnvilBackend(chain.Chain.Name, chain.Chain.ChainID, chain.Chain.BlockPeriod, ethCl)
		if err != nil {
			return Backends{}, errors.Wrap(err, "new anvil Backend")
		}

		inner[chain.Chain.ChainID] = backend
	}

	// Configure public EVM Backends
	for _, chain := range testnet.PublicChains {
		if publicDeployKey == nil {
			return Backends{}, errors.New("public deploy key required")
		}
		ethCl, err := ethclient.Dial(chain.Chain().Name, chain.NextRPCAddress())
		if err != nil {
			return Backends{}, errors.Wrap(err, "dial")
		}

		backend, err := NewBackend(chain.Chain().Name, chain.Chain().ChainID, chain.Chain().BlockPeriod, ethCl, publicDeployKey)
		if err != nil {
			return Backends{}, errors.Wrap(err, "new public Backend")
		} else if err := backend.EnsureSynced(ctx); err != nil {
			return Backends{}, errors.Wrap(err, "ensure public chain synced", "chain", chain.Chain().Name)
		}

		inner[chain.Chain().ChainID] = backend
	}

	return Backends{
		backends: inner,
	}, nil
}

func (b Backends) All() map[uint64]*Backend {
	return b.backends
}

func (b Backends) Backend(sourceChainID uint64) (*Backend, error) {
	backend, ok := b.backends[sourceChainID]
	if !ok {
		return nil, errors.New("unknown chain", "chain", sourceChainID)
	}

	return backend, nil
}

// BindOpts is a convenience function that an accounts' bind.TransactOpts and Backend for a given chain.
func (b Backends) BindOpts(ctx context.Context, sourceChainID uint64, addr common.Address) (*bind.TransactOpts, *Backend, error) {
	backend, ok := b.backends[sourceChainID]
	if !ok {
		return nil, nil, errors.New("unknown chain", "chain", sourceChainID)
	}

	opts, err := backend.BindOpts(ctx, addr)
	if err != nil {
		return nil, nil, errors.Wrap(err, "bind opts")
	}

	return opts, backend, nil
}

func (b Backends) RPCClients() map[uint64]ethclient.Client {
	clients := make(map[uint64]ethclient.Client)
	for chainID, backend := range b.backends {
		clients[chainID] = backend.Client
	}

	return clients
}

func newFireblocksTxMgr(ethCl ethclient.Client, chainName string, chainID uint64, blockPeriod time.Duration, from common.Address, fireCl fireblocks.Client) (txmgr.TxManager, error) {
	// creates our new CLI config for our tx manager
	defaults := txmgr.DefaultSenderFlagValues
	defaults.NetworkTimeout = time.Minute * 5
	cliConfig := txmgr.NewCLIConfig(
		chainID,
		blockPeriod/interval,
		defaults,
	)

	// get the config for our tx manager
	cfg, err := txmgr.NewConfigWithSigner(cliConfig, fireCl.Sign, from, ethCl)
	if err != nil {
		return nil, errors.Wrap(err, "new config")
	}

	// create a simple tx manager from our config
	txMgr, err := txmgr.NewSimple(chainName, cfg)
	if err != nil {
		return nil, errors.Wrap(err, "new simple")
	}

	return txMgr, nil
}

func newTxMgr(ethCl ethclient.Client, chainName string, chainID uint64, blockPeriod time.Duration, privateKey *ecdsa.PrivateKey) (txmgr.TxManager, error) {
	// creates our new CLI config for our tx manager
	cliConfig := txmgr.NewCLIConfig(
		chainID,
		blockPeriod/interval,
		txmgr.DefaultSenderFlagValues,
	)

	// get the config for our tx manager
	cfg, err := txmgr.NewConfig(cliConfig, privateKey, ethCl)
	if err != nil {
		return nil, errors.Wrap(err, "new config")
	}

	// create a simple tx manager from our config
	txMgr, err := txmgr.NewSimple(chainName, cfg)
	if err != nil {
		return nil, errors.Wrap(err, "new simple")
	}

	return txMgr, nil
}
