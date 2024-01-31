package netman

import (
	"context"
	"crypto/ecdsa"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	// privKeyHex0 of pre-funded anvil account 0.
	privKeyHex0 = "0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
	// privKeyHex1 of pre-funded anvil account 1.
	privKeyHex1 = "0x59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d"
)

//nolint:gochecknoglobals // Static mapping.
var (
	privKey0 = mustHexToKey(privKeyHex0)
	privKey1 = mustHexToKey(privKeyHex1)
)

// Manager abstract logic to deploy and bootstrap a network.
type Manager interface {
	// DeployPublicPortals deploys portals to public chains, like arb-goerli.
	DeployPublicPortals(ctx context.Context) error

	// DeployPrivatePortals deploys portals to private (docker) chains.
	DeployPrivatePortals(ctx context.Context) error

	// Portals returns the deployed portals from both public and private chains.
	Portals() map[uint64]Portal

	// Network returns the network configuration.
	// Note that RPCURLs for private chains are empty, it is defined by the infra provider.
	Network() netconf.Network

	// RelayerKey returns the relayer private key hex.
	RelayerKey() (*ecdsa.PrivateKey, error)

	// AdditionalService returns additional services to run in docker (as opposed to halo validators).
	AdditionalService() []string
}

func NewManager(network string, deployKeyFile string, relayerKeyFile string) (Manager, error) {
	switch network {
	case "", netconf.Devnet: // Default to devnetManager if not specified.
		if deployKeyFile != "" || relayerKeyFile != "" {
			return nil, errors.New("deploy and relayer keys not supported in devnet")
		}

		return &devnetManager{network: defaultDevnet()}, nil
	case netconf.Staging:
		deployKey, err := crypto.LoadECDSA(deployKeyFile)
		if err != nil {
			return nil, errors.Wrap(err, "read deploy key file", "path", deployKeyFile)
		}
		relayerKey, err := crypto.LoadECDSA(relayerKeyFile)
		if err != nil {
			return nil, errors.Wrap(err, "read relayer key file", "path", relayerKeyFile)
		}

		return &stagingManager{
			network:    defaultStaging(),
			deployKey:  deployKey,
			relayerKey: relayerKey,
		}, nil
	default:
		return nil, errors.New("unknown network")
	}
}

type Portal struct {
	Chain    netconf.Chain
	Client   *ethclient.Client
	Contract *bindings.OmniPortal
	txOpts   *bind.TransactOpts
}

// TxOpts returns transaction options using the deploy key.
func (p Portal) TxOpts(ctx context.Context) *bind.TransactOpts {
	clone := *p.txOpts
	clone.Context = ctx

	return &clone
}
