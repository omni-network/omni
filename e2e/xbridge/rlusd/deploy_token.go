package rlusd

import (
	"context"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/e2e/xbridge/types"
	"github.com/omni-network/omni/lib/contracts"
	"github.com/omni-network/omni/lib/contracts/proxy"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

type TokenConfig struct {
	Token      types.TokenDescriptors
	Admin      common.Address
	Upgrader   common.Address
	Pauser     common.Address
	Minter     common.Address
	Clawbacker common.Address
}

func isEmpty(addr common.Address) bool {
	return addr == common.Address{}
}

func (cfg TokenConfig) Validate() error {
	if cfg.Token.Name == "" {
		return errors.New("token name is empty")
	}
	if cfg.Token.Symbol == "" {
		return errors.New("token symbol is empty")
	}
	if isEmpty(cfg.Minter) {
		return errors.New("minter is zero")
	}
	if isEmpty(cfg.Admin) {
		return errors.New("admin is zero")
	}
	if isEmpty(cfg.Upgrader) {
		return errors.New("upgrader is zero")
	}
	if isEmpty(cfg.Pauser) {
		return errors.New("pauser is zero")
	}
	if isEmpty(cfg.Clawbacker) {
		return errors.New("clawbacker is zero")
	}

	return nil
}

func saltID(tkn types.TokenDescriptors) string {
	return tkn.Symbol
}

func deployXToken(
	ctx context.Context,
	network netconf.ID,
	backend *ethbackend.Backend,
	bridge common.Address) (common.Address, *ethtypes.Receipt, error) {
	cfg := TokenConfig{
		Token:      xtoken,
		Upgrader:   eoa.MustAddress(network, eoa.RoleUpgrader),
		Admin:      eoa.MustAddress(network, eoa.RoleManager),
		Pauser:     eoa.MustAddress(network, eoa.RoleManager),
		Clawbacker: bridge,
		Minter:     bridge,
	}

	return deployToken(ctx, cfg, network, backend)
}

func deployCanonical(
	ctx context.Context,
	network netconf.ID,
	backend *ethbackend.Backend) (common.Address, *ethtypes.Receipt, error) {
	cfg := TokenConfig{
		Token:      wraps,
		Upgrader:   eoa.MustAddress(network, eoa.RoleUpgrader),
		Admin:      eoa.MustAddress(network, eoa.RoleManager),
		Pauser:     eoa.MustAddress(network, eoa.RoleManager),
		Clawbacker: eoa.MustAddress(network, eoa.RoleManager),
		Minter:     eoa.MustAddress(network, eoa.RoleManager),
	}

	return deployToken(ctx, cfg, network, backend)
}

func deployToken(
	ctx context.Context,
	cfg TokenConfig,
	network netconf.ID,
	backend *ethbackend.Backend) (common.Address, *ethtypes.Receipt, error) {
	salt, err := contracts.Create3Salt(ctx, network, saltID(cfg.Token))
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "salt")
	}

	return proxy.Deploy(ctx, backend, proxy.DeployParams{
		Network:     network,
		Create3Salt: salt,
		DeployImpl: func(txOpts *bind.TransactOpts, backend *ethbackend.Backend) (common.Address, *ethtypes.Transaction, error) {
			addr, tx, _, err := bindings.DeployStablecoinUpgradeable(txOpts, backend)
			return addr, tx, err
		},
		PackInitCode: func(impl common.Address) ([]byte, error) {
			return packInitCode(cfg, impl)
		},
	})
}

func packInitCode(cfg TokenConfig, impl common.Address) ([]byte, error) {
	stablecoinAbi, err := bindings.StablecoinUpgradeableMetaData.GetAbi()
	if err != nil {
		return nil, errors.Wrap(err, "get abi")
	}

	proxyAbi, err := bindings.StablecoinProxyMetaData.GetAbi()
	if err != nil {
		return nil, errors.Wrap(err, "get proxy abi")
	}

	initializer, err := stablecoinAbi.Pack("initialize",
		cfg.Token.Name, cfg.Token.Symbol, cfg.Minter,
		cfg.Admin, cfg.Upgrader, cfg.Pauser, cfg.Clawbacker,
	)
	if err != nil {
		return nil, errors.Wrap(err, "encode initializer")
	}

	return contracts.PackInitCode(proxyAbi, bindings.StablecoinProxyBin, impl, initializer)
}
