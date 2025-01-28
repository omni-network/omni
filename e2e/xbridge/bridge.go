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

type bridgeDeploymentConfig struct {
	Config          deploymentConfig
	ProxyAdminOwner common.Address
	Admin           common.Address
	Pauser          common.Address
	OmniPortal      common.Address
	Token           common.Address
	Lockbox         common.Address
}

func (cfg bridgeDeploymentConfig) validateBridgeConfig() error {
	if err := cfg.Config.validateDeploymentConfig(); err != nil {
		return errors.Wrap(err, "validate config")
	}
	if isEmpty(cfg.ProxyAdminOwner) {
		return errors.New("proxy admin is zero")
	}
	if isEmpty(cfg.Admin) {
		return errors.New("admin is zero")
	}
	if isEmpty(cfg.Pauser) {
		return errors.New("pauser is zero")
	}
	if isEmpty(cfg.OmniPortal) {
		return errors.New("omni portal is zero")
	}
	if isEmpty(cfg.Token) {
		return errors.New("token is zero")
	}

	return nil
}

// bridgeAddress returns the Bridge contract address for the given network.
func bridgeAddress(ctx context.Context, network netconf.ID) (common.Address, error) {
	return contracts.Create3Address(ctx, network, "bridge")
}

// deployBridgeIfNeeded deploys a new bridge contract if it is not already deployed.
// If the contract is already deployed, the receipt is nil.
func DeployBridgeIfNeeded(ctx context.Context, network netconf.ID, backend *ethbackend.Backend, lockbox bool) (common.Address, *ethtypes.Receipt, error) {
	bridgeAddr, err := bridgeAddress(ctx, network)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "bridge address")
	}

	deployed, addr, err := isDeployed(ctx, backend, bridgeAddr)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "is deployed")
	}
	if deployed {
		return addr, nil, nil
	}

	return deployBridge(ctx, network, backend, lockbox)
}

// deployBridge deploys a new bridge contract and returns the address and receipt.
func deployBridge(ctx context.Context, network netconf.ID, backend *ethbackend.Backend, lockbox bool) (common.Address, *ethtypes.Receipt, error) {
	addrs, err := contracts.GetAddresses(ctx, network)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "get addrs")
	}

	wrapper := true
	tokenAddr, err := tokenAddress(ctx, network, wrapper)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "token address")
	}

	lockboxAddr := common.Address{}
	if lockbox {
		lockboxAddr, err = lockboxAddress(ctx, network)
		if err != nil {
			return common.Address{}, nil, errors.Wrap(err, "lockbox address")
		}
	}

	bridgeAddr, err := bridgeAddress(ctx, network)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "bridge address")
	}

	bridgeSalt, err := contracts.Create3Salt(ctx, network, "bridge")
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "bridge salt")
	}

	deployCfg := deploymentConfig{
		Create3Salt:    bridgeSalt,
		Create3Factory: addrs.Create3Factory,
		ExpectedAddr:   bridgeAddr,
		Deployer:       eoa.MustAddress(network, eoa.RoleDeployer),
	}

	cfg := bridgeDeploymentConfig{
		Config:          deployCfg,
		ProxyAdminOwner: eoa.MustAddress(network, eoa.RoleUpgrader),
		Admin:           eoa.MustAddress(network, eoa.RoleManager),
		Pauser:          eoa.MustAddress(network, eoa.RoleManager),
		OmniPortal:      addrs.Portal,
		Token:           tokenAddr,
		Lockbox:         lockboxAddr,
	}

	return performBridgeDeployment(ctx, network, backend, cfg)
}

// performBridgeDeployment handles the common deployment flow for the bridge contract.
func performBridgeDeployment(ctx context.Context, network netconf.ID, backend *ethbackend.Backend, cfg bridgeDeploymentConfig) (common.Address, *ethtypes.Receipt, error) {
	params := deploymentParams{
		Config:         cfg.Config,
		ValidateConfig: cfg.validateBridgeConfig,
		DeployImpl: func(txOpts *bind.TransactOpts, backend *ethbackend.Backend) (common.Address, *ethtypes.Transaction, error) {
			addr, tx, _, err := bindings.DeployBridge(txOpts, backend)
			return addr, tx, err
		},
		PackInitCode: func(impl common.Address) ([]byte, error) {
			return packBridgeInitCode(cfg, impl)
		},
	}

	return performDeployment(ctx, network, backend, params)
}

// packBridgeInitCode packs the initialization code for the bridge contract proxy.
func packBridgeInitCode(cfg bridgeDeploymentConfig, impl common.Address) ([]byte, error) {
	bridgeAbi, err := bindings.BridgeMetaData.GetAbi()
	if err != nil {
		return nil, errors.Wrap(err, "get abi")
	}

	proxyAbi, err := bindings.TransparentUpgradeableProxyMetaData.GetAbi()
	if err != nil {
		return nil, errors.Wrap(err, "get proxy abi")
	}

	initializer, err := bridgeAbi.Pack("initialize", cfg.Admin, cfg.Pauser, cfg.OmniPortal, cfg.Token, cfg.Lockbox)
	if err != nil {
		return nil, errors.Wrap(err, "encode initializer")
	}

	return contracts.PackInitCode(proxyAbi, bindings.TransparentUpgradeableProxyBin, impl, cfg.ProxyAdminOwner, initializer)
}
