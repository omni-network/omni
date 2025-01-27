package xbridge

import (
	"context"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/create3"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

// deploymentConfig contains all of the shared parameters for a deployment.
type deploymentConfig struct {
	Create3Salt    string
	Create3Factory common.Address
	ExpectedAddr   common.Address
	Deployer       common.Address
}

// deploymentParams contains all of the parameters for a generic deployment.
type deploymentParams struct {
	Config         deploymentConfig
	ValidateConfig func() error
	DeployImpl     func(*bind.TransactOpts, *ethbackend.Backend) (common.Address, *ethtypes.Transaction, error)
	PackInitCode   func(common.Address) ([]byte, error)
}

func (cfg deploymentConfig) validateDeploymentConfig() error {
	if cfg.Create3Salt == "" {
		return errors.New("create3 salt is empty")
	}
	if isEmpty(cfg.Create3Factory) {
		return errors.New("create3 factory is zero")
	}
	if isEmpty(cfg.ExpectedAddr) {
		return errors.New("expected address is zero")
	}
	if isEmpty(cfg.Deployer) {
		return errors.New("deployer is zero")
	}

	return nil
}

func isEmpty(addr common.Address) bool {
	return addr == common.Address{}
}

// isDeployed returns true if the rlusd contract is already deployed to its expected address.
func isDeployed(ctx context.Context, backend *ethbackend.Backend, expectedAddr common.Address) (bool, common.Address, error) {
	code, err := backend.CodeAt(ctx, expectedAddr, nil)
	if err != nil {
		return false, expectedAddr, errors.Wrap(err, "code at", "address", expectedAddr)
	}

	if len(code) == 0 {
		return false, expectedAddr, nil
	}

	return true, expectedAddr, nil
}

// deployPrep handles the common deployment preparation flow for all contracts.
func deployPrep(ctx context.Context, network netconf.ID, backend *ethbackend.Backend, cfg deploymentConfig) (*bindings.Create3, common.Address, [32]byte, *bind.TransactOpts, error) {
	chainID, err := backend.ChainID(ctx)
	if err != nil {
		return nil, common.Address{}, [32]byte{}, nil, errors.Wrap(err, "chain id")
	}

	if chainID.Uint64() == network.Static().OmniExecutionChainID {
		return nil, common.Address{}, [32]byte{}, nil, errors.New("cannot deploy on omni evm")
	}

	txOpts, err := backend.BindOpts(ctx, cfg.Deployer)
	if err != nil {
		return nil, common.Address{}, [32]byte{}, nil, errors.Wrap(err, "bind opts")
	}

	factory, err := bindings.NewCreate3(cfg.Create3Factory, backend)
	if err != nil {
		return nil, common.Address{}, [32]byte{}, nil, errors.Wrap(err, "new create3")
	}

	salt := create3.HashSalt(cfg.Create3Salt)

	addr, err := factory.GetDeployed(nil, txOpts.From, salt)
	if err != nil {
		return nil, common.Address{}, [32]byte{}, nil, errors.Wrap(err, "get deployed")
	} else if (cfg.ExpectedAddr != common.Address{}) && addr != cfg.ExpectedAddr {
		return nil, common.Address{}, [32]byte{}, nil, errors.New("unexpected address", "expected", cfg.ExpectedAddr, "actual", addr)
	}

	return factory, addr, salt, txOpts, nil
}

// performDeployment handles the common deployment flow for all contracts.
func performDeployment(ctx context.Context, network netconf.ID, backend *ethbackend.Backend,
	params deploymentParams) (common.Address, *ethtypes.Receipt, error) {
	if err := params.ValidateConfig(); err != nil {
		return common.Address{}, nil, errors.Wrap(err, "validate config")
	}

	factory, addr, salt, txOpts, err := deployPrep(ctx, network, backend, params.Config)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "deploy prep")
	}

	impl, tx, err := params.DeployImpl(txOpts, backend)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "deploy impl")
	}

	_, err = backend.WaitMined(ctx, tx)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "wait mined impl")
	}

	initCode, err := params.PackInitCode(impl)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "pack init code")
	}

	tx, err = factory.DeployWithRetry(txOpts, salt, initCode) //nolint:contextcheck // Context is txOpts
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "deploy proxy")
	}

	receipt, err := backend.WaitMined(ctx, tx)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "wait mined proxy")
	}

	return addr, receipt, nil
}
