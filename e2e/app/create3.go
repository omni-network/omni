package app

import (
	"context"

	"github.com/omni-network/omni/lib/contracts/create3"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/common"
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

	addr, receipt, err := deployCreate3(ctx, def, cfg.ChainID)
	if err != nil {
		return errors.Wrap(err, "deploy create3")
	}

	log.Info(ctx, "Create3 factory deployed", "chain", cfg.ChainID, "addr", addr.Hex(), "block", receipt.BlockNumber)

	return nil
}

func deployAllCreate3(ctx context.Context, def Definition) error {
	for _, chain := range def.Testnet.EVMChains() {
		_, _, err := deployCreate3(ctx, def, chain.ChainID)
		if err != nil {
			return errors.Wrap(err, "deploy create3", "chain", chain.Name)
		}
	}

	return nil
}

func DeployAllCreate3(ctx context.Context, network netconf.Network, backends ethbackend.Backends) error {
	log.Debug(ctx, "Deploying create3 factory", "network", network.ID)

	for _, chain := range network.EVMChains() {
		backend, err := backends.Backend(chain.ID)
		if err != nil {
			return errors.Wrap(err, "get backend", "chain", chain.Name)
		}

		_, _, err = create3.DeployIfNeeded(ctx, network.ID, backend)
		if err != nil {
			return errors.Wrap(err, "deploy create3", "chain", chain.Name)
		}
	}

	return nil
}

func deployCreate3(ctx context.Context, def Definition, chainID uint64) (common.Address, *ethclient.Receipt, error) {
	backend, err := def.Backends().Backend(chainID)
	if err != nil {
		return common.Address{}, nil, err
	}

	return create3.DeployIfNeeded(ctx, def.Testnet.Network, backend)
}
