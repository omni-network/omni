package xbridge

import (
	"context"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/lib/contracts"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

const (
	rlusdTokenSymbol   = "RLUSD"
	rlusdWrapperSymbol = "RLUSDe"
)

type tokenDeploymentConfig struct {
	Config     deploymentConfig
	Name       string
	Symbol     string
	Minter     common.Address
	Admin      common.Address
	Upgrader   common.Address
	Pauser     common.Address
	Clawbacker common.Address
}

func (cfg tokenDeploymentConfig) validateTokenConfig() error {
	if err := cfg.Config.validateDeploymentConfig(); err != nil {
		return errors.Wrap(err, "validate config")
	}
	if cfg.Name == "" {
		return errors.New("name is empty")
	}
	if cfg.Symbol == "" {
		return errors.New("symbol is empty")
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

// tokenAddress returns the token contract address for the given network.
func tokenAddress(ctx context.Context, network netconf.ID, wrapper bool) (common.Address, error) {
	saltPrefix := rlusdTokenSymbol
	if wrapper {
		saltPrefix = rlusdWrapperSymbol
	}

	return contracts.Create3Address(ctx, network, saltPrefix)
}

// deployTokenIfNeeded deploys a new rlusd contract if it is not already deployed.
// If the contract is already deployed, the receipt is nil.
func DeployTokenIfNeeded(ctx context.Context, network netconf.ID, backend *ethbackend.Backend, wrapper bool) (common.Address, *ethtypes.Receipt, error) {
	tokenAddr, err := tokenAddress(ctx, network, wrapper)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "token address")
	}

	deployed, addr, err := isDeployed(ctx, backend, tokenAddr)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "is deployed")
	}
	if deployed {
		return addr, nil, nil
	}

	return deployToken(ctx, network, backend, wrapper)
}

// deployRLUSD deploys a new rlusd contract and returns the address and receipt.
func deployToken(ctx context.Context, network netconf.ID, backend *ethbackend.Backend, wrapper bool) (common.Address, *ethtypes.Receipt, error) {
	addrs, err := contracts.GetAddresses(ctx, network)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "get addrs")
	}

	bridgeAddr, err := bridgeAddress(ctx, network)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "bridge address")
	}

	tokenAddr, err := tokenAddress(ctx, network, wrapper)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "token address")
	}

	saltPrefix := rlusdTokenSymbol
	if wrapper {
		saltPrefix = rlusdWrapperSymbol
	}
	tokenSalt, err := contracts.Create3Salt(ctx, network, saltPrefix)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "token salt")
	}

	// as we reuse the same token template for both RLUSD and RLUSDe, we need to set values accordingly
	name := "Ripple USD"
	symbol := rlusdTokenSymbol
	minter := eoa.MustAddress(network, eoa.RoleManager)
	if wrapper {
		name = "Bridged RLUSD (Omni)"
		symbol = rlusdWrapperSymbol
		minter = bridgeAddr
	}

	deployCfg := deploymentConfig{
		Create3Salt:    tokenSalt,
		Create3Factory: addrs.Create3Factory,
		ExpectedAddr:   tokenAddr,
		Deployer:       eoa.MustAddress(network, eoa.RoleDeployer),
	}

	cfg := tokenDeploymentConfig{
		Config:     deployCfg,
		Name:       name,
		Symbol:     symbol,
		Minter:     minter,
		Admin:      eoa.MustAddress(network, eoa.RoleManager),
		Upgrader:   eoa.MustAddress(network, eoa.RoleUpgrader),
		Pauser:     eoa.MustAddress(network, eoa.RoleManager),
		Clawbacker: eoa.MustAddress(network, eoa.RoleManager),
	}

	return performTokenDeployment(ctx, network, backend, cfg)
}

// performTokenDeployment handles the common deployment flow for the token contract.
func performTokenDeployment(ctx context.Context, network netconf.ID, backend *ethbackend.Backend, cfg tokenDeploymentConfig) (common.Address, *ethtypes.Receipt, error) {
	params := deploymentParams{
		Config:         cfg.Config,
		ValidateConfig: cfg.validateTokenConfig,
		DeployImpl: func(txOpts *bind.TransactOpts, backend *ethbackend.Backend) (common.Address, *ethtypes.Transaction, error) {
			addr, tx, _, err := bindings.DeployStablecoinUpgradeable(txOpts, backend)
			return addr, tx, err
		},
		PackInitCode: func(impl common.Address) ([]byte, error) {
			return packTokenInitCode(cfg, impl)
		},
	}

	return performDeployment(ctx, network, backend, params)
}

// packTokenInitCode packs the initialization code for the token contract proxy.
func packTokenInitCode(cfg tokenDeploymentConfig, impl common.Address) ([]byte, error) {
	stablecoinAbi, err := bindings.StablecoinUpgradeableMetaData.GetAbi()
	if err != nil {
		return nil, errors.Wrap(err, "get abi")
	}

	proxyAbi, err := bindings.StablecoinProxyMetaData.GetAbi()
	if err != nil {
		return nil, errors.Wrap(err, "get proxy abi")
	}

	initializer, err := stablecoinAbi.Pack("initialize", cfg.Name, cfg.Symbol, cfg.Minter, cfg.Admin, cfg.Upgrader, cfg.Pauser, cfg.Clawbacker)
	if err != nil {
		return nil, errors.Wrap(err, "encode initializer")
	}

	return contracts.PackInitCode(proxyAbi, bindings.StablecoinProxyBin, impl, initializer)
}
