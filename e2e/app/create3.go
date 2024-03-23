package app

import (
	"context"

	"github.com/omni-network/omni/lib/contracts/create3"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
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
	backend, err := def.Backends.Backend(cfg.ChainID)
	if err != nil {
		return err
	}

	addr, receipt, err := create3.Deploy(ctx, def.Testnet.Network, backend)
	if err != nil {
		return err
	}

	log.Info(ctx, "Create3 factory deployed", "chain", cfg.ChainID, "addr", addr.Hex(), "block", receipt.BlockNumber)

	return nil
}

func deployCreate3Factories(ctx context.Context, def Definition) error {
	// TODO: support all networks
	if def.Testnet.Network != netconf.Devnet {
		return nil
	}

	for _, c := range def.Testnet.AnvilChains {
		cfg := Create3DeployConfig{ChainID: c.Chain.ID}

		if err := Create3Deploy(ctx, def, cfg); err != nil {
			return errors.Wrap(err, "deploy create3")
		}
	}

	// only deploy to omni evm once
	if len(def.Testnet.OmniEVMs) > 0 {
		c := def.Testnet.OmniEVMs[0]
		cfg := Create3DeployConfig{ChainID: c.Chain.ID}

		if err := Create3Deploy(ctx, def, cfg); err != nil {
			return errors.Wrap(err, "deploy create3")
		}
	}

	return nil
}
