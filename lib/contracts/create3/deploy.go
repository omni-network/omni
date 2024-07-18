package create3

import (
	"context"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/lib/contracts"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

type DeploymentConfig struct {
	Deployer common.Address
}

func (cfg DeploymentConfig) Validate() error {
	if (cfg.Deployer == common.Address{}) {
		return errors.New("deployer is zero")
	}

	return nil
}

var configs = map[netconf.ID]DeploymentConfig{
	netconf.Mainnet: {
		Deployer: eoa.MustAddress(netconf.Mainnet, eoa.RoleCreate3Deployer),
	},
	netconf.Omega: {
		Deployer: eoa.MustAddress(netconf.Omega, eoa.RoleCreate3Deployer),
	},
	netconf.Staging: {
		Deployer: eoa.MustAddress(netconf.Staging, eoa.RoleCreate3Deployer),
	},
	netconf.Devnet: {
		Deployer: eoa.MustAddress(netconf.Devnet, eoa.RoleCreate3Deployer),
	},
}

func AddrForNetwork(network netconf.ID) (common.Address, bool) {
	switch network {
	case netconf.Mainnet:
		return contracts.MainnetCreate3Factory(), true
	case netconf.Omega:
		return contracts.OmegaCreate3Factory(), true
	case netconf.Staging:
		return contracts.StagingCreate3Factory(), true
	case netconf.Devnet:
		return contracts.DevnetCreate3Factory(), true
	default:
		return common.Address{}, false
	}
}

// IsDeployed checks if the Create3 factory contract is deployed to the provided backend
// to its expected network address.
func IsDeployed(ctx context.Context, network netconf.ID, backend *ethbackend.Backend) (bool, common.Address, error) {
	addr, ok := AddrForNetwork(network)
	if !ok {
		return false, addr, errors.New("unsupported network", "network", network)
	}

	code, err := backend.CodeAt(ctx, addr, nil)
	if err != nil {
		return false, addr, errors.Wrap(err, "code at", "address", addr)
	}

	if len(code) == 0 {
		return false, addr, nil
	}

	return true, addr, nil
}

// DeployIfNeeded deploys a new Create3 factory contract if it is not already deployed.
func DeployIfNeeded(ctx context.Context, network netconf.ID, backend *ethbackend.Backend) (common.Address, *ethtypes.Receipt, error) {
	deployed, addr, err := IsDeployed(ctx, network, backend)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "is deployed")
	}
	if deployed {
		return addr, nil, nil
	}

	return Deploy(ctx, network, backend)
}

// Deploy deploys a new Create3 factory contract and returns the address and receipt.
// It only allows deployments to explicitly supported chains.
func Deploy(ctx context.Context, network netconf.ID, backend *ethbackend.Backend) (common.Address, *ethtypes.Receipt, error) {
	cfg, ok := configs[network]
	if !ok {
		return common.Address{}, nil, errors.New("unsupported network", "network", network)
	}

	return deploy(ctx, cfg, backend)
}

func deploy(ctx context.Context, cfg DeploymentConfig, backend *ethbackend.Backend) (common.Address, *ethtypes.Receipt, error) {
	if err := cfg.Validate(); err != nil {
		return common.Address{}, nil, errors.Wrap(err, "validate config")
	}

	txOpts, err := backend.BindOpts(ctx, cfg.Deployer)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "bind opts")
	}

	nonce, err := backend.PendingNonceAt(ctx, cfg.Deployer)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "pending nonce")
	} else if nonce != 0 {
		return common.Address{}, nil, errors.New("nonce not zero")
	}

	addr, tx, _, err := bindings.DeployCreate3(txOpts, backend)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "deploy create3")
	}

	receipt, err := backend.WaitMined(ctx, tx)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "wait mined")
	}

	return addr, receipt, nil
}
