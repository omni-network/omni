package create3

import (
	"context"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/lib/contracts"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/common"
)

// IsDeployed checks if the Create3 factory contract is deployed to the provided backend
// to its expected network address.
func IsDeployed(ctx context.Context, network netconf.ID, backend *ethbackend.Backend) (bool, common.Address, error) {
	addrs, err := contracts.GetAddresses(ctx, network)
	if err != nil {
		return false, common.Address{}, errors.Wrap(err, "get addrs")
	}

	code, err := backend.CodeAt(ctx, addrs.Create3Factory, nil)
	if err != nil {
		return false, addrs.Create3Factory, errors.Wrap(err, "code at", "address", addrs.Create3Factory)
	}

	if len(code) == 0 {
		return false, addrs.Create3Factory, nil
	}

	return true, addrs.Create3Factory, nil
}

// DeployIfNeeded deploys a new Create3 factory contract if it is not already deployed.
func DeployIfNeeded(ctx context.Context, network netconf.ID, backend *ethbackend.Backend) (common.Address, *ethclient.Receipt, error) {
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
func Deploy(ctx context.Context, network netconf.ID, backend *ethbackend.Backend) (common.Address, *ethclient.Receipt, error) {
	cfg, ok := eoa.Address(network, eoa.RoleCreate3Deployer)
	if !ok {
		return common.Address{}, nil, errors.New("unsupported network", "network", network)
	}

	return deploy(ctx, cfg, backend)
}

func deploy(ctx context.Context, deployer common.Address, backend *ethbackend.Backend) (common.Address, *ethclient.Receipt, error) {
	txOpts, err := backend.BindOpts(ctx, deployer)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "bind opts")
	}

	nonce, err := backend.PendingNonceAt(ctx, deployer)
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
