package create3

import (
	"context"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/chainids"
	"github.com/omni-network/omni/lib/contracts"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

type DeploymentConfig struct {
	Deployer          common.Address
	AllowNonZeroNonce bool
}

func (cfg DeploymentConfig) Validate() error {
	if (cfg.Deployer == common.Address{}) {
		return errors.New("deployer is zero")
	}

	return nil
}

func getDeployCfg(chainID uint64) (DeploymentConfig, error) {
	if chainids.IsMainnet(chainID) {
		return mainnetDeployCfg(), nil
	}

	if chainids.IsTestnet(chainID) {
		return testnetDeployCfg(), nil
	}

	return DeploymentConfig{}, errors.New("unsupported chain")
}

func testnetDeployCfg() DeploymentConfig {
	return DeploymentConfig{
		Deployer:          contracts.TestnetCreate3Deployer,
		AllowNonZeroNonce: false,
	}
}

func mainnetDeployCfg() DeploymentConfig {
	return DeploymentConfig{
		Deployer:          contracts.MainnetCreate3Deployer,
		AllowNonZeroNonce: false,
	}
}

func devnetDeployCfg() DeploymentConfig {
	return DeploymentConfig{
		Deployer:          contracts.DevnetCreate3Deployer,
		AllowNonZeroNonce: true,
	}
}

// Deploy deploys a new Create3 factory contract and returns the address and receipt.
// It only allows deployments to explicitly supported chains.
func Deploy(ctx context.Context, backend *ethbackend.Backend) (common.Address, *ethtypes.Receipt, error) {
	chainID, err := backend.ChainID(ctx)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "chain id")
	}

	cfg, err := getDeployCfg(chainID.Uint64())
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "get deployment config")
	}

	return deploy(ctx, cfg, backend)
}

// DeployDevnet deploys the devnet AVS contract and returns the address receipt.
func DeployDevnet(ctx context.Context, backend *ethbackend.Backend) (common.Address, *ethtypes.Receipt, error) {
	chainID, err := backend.ChainID(ctx)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "chain id")
	}

	if chainids.IsMainnetOrTestnet(chainID.Uint64()) {
		return common.Address{}, nil, errors.New("not a devnet")
	}

	return deploy(ctx, devnetDeployCfg(), backend)
}

func deploy(ctx context.Context, cfg DeploymentConfig, backend *ethbackend.Backend) (common.Address, *ethtypes.Receipt, error) {
	if err := cfg.Validate(); err != nil {
		return common.Address{}, nil, errors.Wrap(err, "validate config")
	}

	txOpts, err := backend.BindOpts(ctx, cfg.Deployer)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "bind opts")
	}

	if !cfg.AllowNonZeroNonce {
		nonce, err := backend.PendingNonceAt(ctx, cfg.Deployer)
		if err != nil {
			return common.Address{}, nil, errors.Wrap(err, "pending nonce")
		} else if nonce != 0 {
			return common.Address{}, nil, errors.New("nonce not zero")
		}
	}

	addr, tx, _, err := bindings.DeployCreate3(txOpts, backend)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "deploy create3")
	}

	receipt, err := bind.WaitMined(ctx, backend, tx)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "wait mined")
	} else if receipt.Status != ethtypes.ReceiptStatusSuccessful {
		return common.Address{}, nil, errors.New("receipt status failed")
	}

	return addr, receipt, nil
}
