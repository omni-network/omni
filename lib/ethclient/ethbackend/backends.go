package ethbackend

import (
	"context"
	"crypto/ecdsa"
	"time"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/fireblocks"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/txmgr"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
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

func BackendsFromNetwork(network netconf.Network, endpoints xchain.RPCEndpoints, privKeys ...*ecdsa.PrivateKey) (Backends, error) {
	inner := make(map[uint64]*Backend)
	for _, chain := range network.EVMChains() {
		endpoint, err := endpoints.ByNameOrID(chain.Name, chain.ID)
		if err != nil {
			return Backends{}, err
		}

		ethCl, err := ethclient.Dial(chain.Name, endpoint)
		if err != nil {
			return Backends{}, errors.Wrap(err, "dial")
		}

		inner[chain.ID], err = NewBackend(chain.Name, chain.ID, chain.BlockPeriod, ethCl, privKeys...)
		if err != nil {
			return Backends{}, errors.Wrap(err, "new backend")
		}
	}

	return Backends{
		backends: inner,
	}, nil
}

func BackendsFrom(backends map[uint64]*Backend) Backends {
	return Backends{backends: backends}
}

func (b Backends) All() map[uint64]*Backend {
	return b.backends
}

// AddAccount adds a in-memory private key account to all backends.
// Note this can be called even if other accounts are fireblocks based.
func (b Backends) AddAccount(privkey *ecdsa.PrivateKey) (common.Address, error) {
	var addr common.Address
	for _, backend := range b.backends {
		var err error
		addr, err = backend.AddAccount(privkey)
		if err != nil {
			return common.Address{}, err
		}
	}

	return addr, nil
}

func (b Backends) Clients() map[uint64]ethclient.Client {
	clients := make(map[uint64]ethclient.Client)
	for chainID, backend := range b.backends {
		clients[chainID] = backend.Client
	}

	return clients
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
