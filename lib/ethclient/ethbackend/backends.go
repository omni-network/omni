package ethbackend

import (
	"context"
	"crypto/ecdsa"
	"strings"
	"time"

	"github.com/omni-network/omni/e2e/types"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/txmgr"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

const (
	interval = 3

	// keys of pre-funded anvil account 0.
	privKeyHex0 = "0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
)

//nolint:gochecknoglobals // Static mapping.
var (
	privateDeployKey = mustHexToKey(privKeyHex0)
)

// Backends is a wrapper around a set of Backends, one for each chain.
// At this point, it only supports "a single account for all Backends".
//
// See Backends godoc for more information.
type Backends struct {
	backends map[uint64]*Backend
}

func NewBackends(testnet types.Testnet, deployKeyFile string) (Backends, error) {
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
		// todo(lazar): remove this when we figure out why txs are stuck in geth mempool upon initial run
		// task https://app.asana.com/0/1206208509925075/1206887969751598/f
		chain, err := testnet.FirstOmniValidatorEVM() // Connect to a geth node connected to a validator
		if err != nil {
			return Backends{}, errors.Wrap(err, "omni evm validator")
		}
		ethCl, err := ethclient.Dial(chain.Chain.Name, chain.ExternalRPC)
		if err != nil {
			return Backends{}, errors.Wrap(err, "dial")
		}

		inner[chain.Chain.ID], err = NewBackend(chain.Chain.Name, chain.Chain.ID, chain.Chain.BlockPeriod, ethCl, privateDeployKey)
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

		inner[chain.Chain.ID], err = NewBackend(chain.Chain.Name, chain.Chain.ID, chain.Chain.BlockPeriod, ethCl, privateDeployKey)
		if err != nil {
			return Backends{}, errors.Wrap(err, "new anvil Backend")
		}
	}

	// Configure public EVM Backends
	for _, chain := range testnet.PublicChains {
		if publicDeployKey == nil {
			return Backends{}, errors.New("public deploy key required")
		}
		ethCl, err := ethclient.Dial(chain.Chain.Name, chain.RPCAddress)
		if err != nil {
			return Backends{}, errors.Wrap(err, "dial")
		}

		inner[chain.Chain.ID], err = NewBackend(chain.Chain.Name, chain.Chain.ID, chain.Chain.BlockPeriod, ethCl, publicDeployKey)
		if err != nil {
			return Backends{}, errors.Wrap(err, "new public Backend")
		}
	}

	return Backends{
		backends: inner,
	}, nil
}

// BindOpts is a convenience function that returns the single account and bind.TransactOpts and Backend for a given chain.
func (b Backends) BindOpts(ctx context.Context, sourceChainID uint64) (common.Address, *bind.TransactOpts, *Backend, error) {
	backend, ok := b.backends[sourceChainID]
	if !ok {
		return common.Address{}, nil, nil, errors.New("unknown chain", "chain", sourceChainID)
	}

	if len(backend.accounts) != 1 {
		return common.Address{}, nil, nil, errors.New("only single account Backends supported")
	}

	// Get the first account
	var addr common.Address
	for a := range backend.accounts {
		addr = a
		break
	}

	opts, err := backend.BindOpts(ctx, addr)
	if err != nil {
		return common.Address{}, nil, nil, errors.Wrap(err, "bind opts")
	}

	return addr, opts, backend, nil
}

func (b Backends) RPCClients() map[uint64]ethclient.Client {
	clients := make(map[uint64]ethclient.Client)
	for chainID, backend := range b.backends {
		clients[chainID] = backend.Client
	}

	return clients
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

func mustHexToKey(privKeyHex string) *ecdsa.PrivateKey {
	privKey, err := crypto.HexToECDSA(strings.TrimPrefix(privKeyHex, "0x"))
	if err != nil {
		panic(err)
	}

	return privKey
}
