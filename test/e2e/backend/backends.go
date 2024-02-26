package backend

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"strings"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/txmgr"
	"github.com/omni-network/omni/test/e2e/types"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
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

// Backends is a wrapper around a set of adapted ethclients (backends)
// that delegate transaction sending to a txmgr.TxManager for reliable sending.
//
// It should be used with bindings based contracts, as it provides bind.TransactOpts.
type Backends struct {
	backends map[uint64]backend
}

func New(testnet types.Testnet, deployKeyFile string) (Backends, error) {
	var err error

	var publicDeployKey *ecdsa.PrivateKey
	if testnet.Network == netconf.Devnet {
		if deployKeyFile != "" {
			return Backends{}, errors.New("deploy key not supported in devnet")
		}
	} else if testnet.Network == netconf.Staging {
		publicDeployKey, err = crypto.LoadECDSA(deployKeyFile)
	} else {
		return Backends{}, errors.New("unknown extNetwork")
	}
	if err != nil {
		return Backends{}, errors.Wrap(err, "load deploy key")
	}

	backends := make(map[uint64]backend)

	// Configure omni EVM backend
	{
		chain := testnet.OmniEVMs[0] // Connect to the first omni evm instance for now.
		backends[chain.Chain.ID], err = newBackend(chain.Chain, chain.ExternalRPC, privateDeployKey)
		if err != nil {
			return Backends{}, errors.Wrap(err, "new omni backend")
		}
	}

	// Configure anvil EVM backends
	for _, chain := range testnet.AnvilChains {
		backends[chain.Chain.ID], err = newBackend(chain.Chain, chain.ExternalRPC, privateDeployKey)
		if err != nil {
			return Backends{}, errors.Wrap(err, "new anvil backend")
		}
	}

	// Configure public EVM backends
	for _, chain := range testnet.PublicChains {
		if publicDeployKey == nil {
			return Backends{}, errors.New("public deploy key required")
		}
		backends[chain.Chain.ID], err = newBackend(chain.Chain, chain.RPCAddress, publicDeployKey)
		if err != nil {
			return Backends{}, errors.Wrap(err, "new public backend")
		}
	}

	return Backends{
		backends: backends,
	}, nil
}

// BindOpts returns a new TransactOpts and Backend for interacting with bindings based contracts.
// The TransactOpts are partially stubbed, since txmgr handles nonces and signing.
// The returned backend is a normal ethclient, except that it delegates SendTransaction to txmgr.
//
// Do not cache or store the TransactOpts, as they are not safe for concurrent use (pointer).
// Rather create a new TransactOpts for each transaction.
func (b Backends) BindOpts(ctx context.Context, sourceChainID uint64) (*bind.TransactOpts, Backend, error) {
	backend, ok := b.backends[sourceChainID]
	if !ok {
		return nil, nil, errors.New("unknown backend")
	}

	if header, err := backend.HeaderByNumber(ctx, nil); err != nil {
		return nil, nil, errors.Wrap(err, "header by number")
	} else if header.BaseFee == nil {
		return nil, nil, errors.New("only dynamic transaction backends supported")
	}

	// Stub nonce and signer since txmgr will handle this.
	// Bindings will estimate gas.
	return &bind.TransactOpts{
		From:  backend.from,
		Nonce: big.NewInt(1),
		Signer: func(_ common.Address, tx *ethtypes.Transaction) (*ethtypes.Transaction, error) {
			return tx, nil
		},
		Context: ctx,
	}, backend, nil
}

func newTxMgr(ethCl ethclient.Client, chain types.EVMChain, privateKey *ecdsa.PrivateKey) (txmgr.TxManager, error) {
	// creates our new CLI config for our tx manager
	cliConfig := txmgr.NewCLIConfig(
		chain.ID,
		chain.BlockPeriod/interval,
		txmgr.DefaultSenderFlagValues,
	)

	// get the config for our tx manager
	cfg, err := txmgr.NewConfig(cliConfig, privateKey, ethCl)
	if err != nil {
		return nil, errors.Wrap(err, "new config")
	}

	// create a simple tx manager from our config
	txMgr, err := txmgr.NewSimpleTxManagerFromConfig(chain.Name, cfg)
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
