package app

import (
	"context"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

func DefaultCreate3DeployConfig() Create3DeployConfig {
	return Create3DeployConfig{RequireNonceZero: true}
}

type Create3DeployConfig struct {
	ChainID          uint64 // chain id of the chain to deploy to
	Create3Deployer  string // required Create3 factory deployer address
	RequireNonceZero bool   // require the deployer to have a zero nonce
}

// Validate validates the Create3DeployConfig.
func (cfg Create3DeployConfig) Validate() error {
	if cfg.ChainID == 0 {
		return errors.New("chain id is zero")
	}
	if cfg.Create3Deployer == "" {
		return errors.New("create3 deployer is empty")
	}

	return nil
}

// Create3Deploy deploys the Omni Create3 contracts.
func Create3Deploy(ctx context.Context, def Definition, cfg Create3DeployConfig) error {
	addr, err := deployCreate3(ctx, def, cfg)
	if err != nil {
		return errors.Wrap(err, "deploy create3")
	}

	log.Info(ctx, "Create3 contract deployed", "addr", addr.Hex())

	return nil
}

func deployCreate3(ctx context.Context, def Definition, cfg Create3DeployConfig) (common.Address, error) {
	err := cfg.Validate()
	if err != nil {
		return common.Address{}, errors.Wrap(err, "validate")
	}

	deployer, txOpts, backend, err := def.Backends.BindOpts(ctx, cfg.ChainID)
	if err != nil {
		return common.Address{}, errors.Wrap(err, "bind opts")
	}
	if deployer.Hex() != cfg.Create3Deployer {
		return common.Address{}, errors.New("incorrect deployer address")
	}

	nonce, err := backend.PendingNonceAt(ctx, deployer)
	if err != nil {
		return common.Address{}, errors.Wrap(err, "pending nonce")
	}
	if cfg.RequireNonceZero && nonce != 0 {
		return common.Address{}, errors.New("nonce not zero")
	}

	addr, tx, _, err := bindings.DeployCreate3(txOpts, backend)
	if err != nil {
		return common.Address{}, errors.Wrap(err, "deploy create3")
	}

	receipt, err := bind.WaitMined(ctx, backend, tx)
	if err != nil {
		return common.Address{}, errors.Wrap(err, "wait mined")
	}
	if receipt.Status == ethtypes.ReceiptStatusFailed {
		return common.Address{}, errors.New("receipt status failed")
	}

	return addr, nil
}
