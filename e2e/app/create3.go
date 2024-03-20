package app

import (
	"context"

	"github.com/omni-network/omni/lib/contracts/create3"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
)

type Create3DeployConfig struct {
	ChainID uint64 // chain id of the chain to deploy to
}

// Validate validates the Create3DeployConfig.
func (cfg Create3DeployConfig) Validate() error {
	if cfg.ChainID == 0 {
		return errors.New("chain id is zero")
	}

	return nil
}

// Create3Deploy deploys the Omni Create3 contracts.
func Create3Deploy(ctx context.Context, def Definition, cfg Create3DeployConfig) error {
	_, _, backend, err := def.Backends.BindOpts(ctx, cfg.ChainID)
	if err != nil {
		return errors.Wrap(err, "bind opts")
	}

	addr, receipt, err := create3.Deploy(ctx, backend)
	if err != nil {
		return errors.Wrap(err, "deploy")
	}

	log.Info(ctx, "Create3 factory deployed", "chain", cfg.ChainID, "addr", addr.Hex(), "block", receipt.BlockNumber)

	return nil
}
