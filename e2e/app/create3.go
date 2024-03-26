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
	if err := cfg.Validate(); err != nil {
		return errors.Wrap(err, "validate")
	}

	err := deployCreate3(ctx, def, cfg.ChainID)
	if err != nil {
		return errors.Wrap(err, "deploy create3")
	}

	return nil
}

func deployPrivateCreate3(ctx context.Context, def Definition) error {
	for _, c := range def.Testnet.AnvilChains {
		err := deployCreate3(ctx, def, c.Chain.ID)
		if err != nil {
			return errors.Wrap(err, "deploy create3")
		}
	}

	// only deploy to omni evm once
	if len(def.Testnet.OmniEVMs) > 0 {
		c := def.Testnet.OmniEVMs[0]
		err := deployCreate3(ctx, def, c.Chain.ID)
		if err != nil {
			return errors.Wrap(err, "deploy create3")
		}
	}

	return nil
}

func deployPublicCreate3(ctx context.Context, def Definition) error {
	for _, c := range def.Testnet.PublicChains {
		err := deployCreate3(ctx, def, c.Chain.ID)
		if err != nil {
			return errors.Wrap(err, "deploy create3")
		}
	}

	return nil
}

func deployCreate3(ctx context.Context, def Definition, chainID uint64) error {
	backend, err := def.Backends.Backend(chainID)
	if err != nil {
		return errors.Wrap(err, "backend")
	}

	addr, err := create3.DeployIfNeeded(ctx, def.Testnet.Network, backend)
	if err != nil {
		return errors.Wrap(err, "deploy")
	}

	log.Info(ctx, "Create3 factory deployed", "chain", chainID, "addr", addr.Hex())

	return nil
}
