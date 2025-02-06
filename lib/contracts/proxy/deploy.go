package proxy

import (
	"context"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/lib/contracts"
	"github.com/omni-network/omni/lib/create3"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

type DeployParams struct {
	Network      netconf.ID
	Create3Salt  string
	DeployImpl   func(*bind.TransactOpts, *ethbackend.Backend) (common.Address, *ethtypes.Transaction, error)
	PackInitCode func(common.Address) ([]byte, error)
}

// Deploy deploys a proxy via a network's create3 factory.
func Deploy(ctx context.Context, backend *ethbackend.Backend, params DeployParams) (common.Address, *ethtypes.Receipt, error) {
	network := params.Network
	deployImpl := params.DeployImpl
	packInitCode := params.PackInitCode

	deployer := eoa.MustAddress(network, eoa.RoleDeployer)
	txOpts, err := backend.BindOpts(ctx, deployer)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "bind opts", "deployer", deployer)
	}

	salt := create3.HashSalt(params.Create3Salt)

	factory, err := bindings.NewCreate3(contracts.Create3Factory(network), backend)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "new create3")
	}

	addr := create3.Address(contracts.Create3Factory(network), params.Create3Salt, eoa.MustAddress(network, eoa.RoleDeployer))

	if ok, err := isDeployed(ctx, backend, addr); err != nil {
		return addr, nil, errors.Wrap(err, "is deployed")
	} else if ok {
		return addr, nil, nil // already deployed
	}

	impl, tx, err := deployImpl(txOpts, backend)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "deploy impl")
	}

	_, err = backend.WaitMined(ctx, tx)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "wait mined impl")
	}

	initCode, err := packInitCode(impl)
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

func isDeployed(ctx context.Context, client ethclient.Client, addr common.Address) (bool, error) {
	code, err := client.CodeAt(ctx, addr, nil)
	if err != nil {
		return false, errors.Wrap(err, "code at", "address", addr)
	}

	if len(code) == 0 {
		return false, nil
	}

	return true, nil
}
