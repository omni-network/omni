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

type lockboxDeploymentConfig struct {
	Config          deploymentConfig
	ProxyAdminOwner common.Address
	Admin           common.Address
	Pauser          common.Address
	Token           common.Address
	Wrapped         common.Address
}

func (cfg lockboxDeploymentConfig) validate() error {
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
	if isEmpty(cfg.Token) {
		return errors.New("token is zero")
	}
	if isEmpty(cfg.Wrapped) {
		return errors.New("wrapped is zero")
	}

	return nil
}

// lockboxAddress returns the Lockbox contract address for the given network.
func lockboxAddress(ctx context.Context, network netconf.ID, deployment tokenDescriptors) (common.Address, error) {
	return contracts.Create3Address(ctx, network, deployment.symbol+"lockbox")
}

// lockboxSalt returns the salt for the lockbox contract for the given network.
func lockboxSalt(ctx context.Context, network netconf.ID, deployment tokenDescriptors) (string, error) {
	return contracts.Create3Salt(ctx, network, deployment.symbol+"lockbox")
}

// deployLockboxIfNeeded deploys a new lockbox contract if it is not already deployed.
// If the contract is already deployed, the receipt is nil.
func DeployLockboxIfNeeded(ctx context.Context, network netconf.ID, backend *ethbackend.Backend, deployment xBridgeDeployment) (common.Address, *ethtypes.Receipt, error) {
	lockboxAddr, err := lockboxAddress(ctx, network, deployment.token)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "lockbox address", "deployment", deployment)
	}

	deployed, addr, err := isDeployed(ctx, backend, lockboxAddr)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "is deployed", "deployment", deployment)
	}
	if deployed {
		return addr, nil, nil
	}

	return deployLockbox(ctx, network, backend, deployment)
}

// deployLockbox deploys a new lockbox contract and returns the address and receipt.
func deployLockbox(ctx context.Context, network netconf.ID, backend *ethbackend.Backend, deployment xBridgeDeployment) (common.Address, *ethtypes.Receipt, error) {
	addrs, err := contracts.GetAddresses(ctx, network)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "get addrs", "deployment", deployment)
	}

	tokenAddr, err := tokenAddress(ctx, network, deployment.token)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "token address", "deployment", deployment)
	}

	wrappedAddr, err := tokenAddress(ctx, network, deployment.wrapped)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "wrapped address", "deployment", deployment)
	}

	lockboxAddr, err := lockboxAddress(ctx, network, deployment.token)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "lockbox address", "deployment", deployment)
	}

	lockboxSalt, err := lockboxSalt(ctx, network, deployment.token)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "lockbox salt", "deployment", deployment)
	}

	deployCfg := deploymentConfig{
		Create3Salt:    lockboxSalt,
		Create3Factory: addrs.Create3Factory,
		ExpectedAddr:   lockboxAddr,
		Deployer:       eoa.MustAddress(network, eoa.RoleDeployer),
	}

	cfg := lockboxDeploymentConfig{
		Config:          deployCfg,
		ProxyAdminOwner: eoa.MustAddress(network, eoa.RoleUpgrader),
		Admin:           eoa.MustAddress(network, eoa.RoleManager),
		Pauser:          eoa.MustAddress(network, eoa.RoleManager),
		Token:           tokenAddr,
		Wrapped:         wrappedAddr,
	}

	return performLockboxDeployment(ctx, network, backend, cfg)
}

// performLockboxDeployment handles the common deployment flow for the lockbox contract.
func performLockboxDeployment(ctx context.Context, network netconf.ID, backend *ethbackend.Backend, cfg lockboxDeploymentConfig) (common.Address, *ethtypes.Receipt, error) {
	params := deploymentParams{
		Config:         cfg.Config,
		ValidateConfig: cfg.validate,
		DeployImpl: func(txOpts *bind.TransactOpts, backend *ethbackend.Backend) (common.Address, *ethtypes.Transaction, error) {
			addr, tx, _, err := bindings.DeployLockbox(txOpts, backend)
			return addr, tx, err
		},
		PackInitCode: func(impl common.Address) ([]byte, error) {
			return packLockboxInitCode(cfg, impl)
		},
	}

	return performDeployment(ctx, network, backend, params)
}

// packLockboxInitCode packs the initialization code for the lockbox contract proxy.
func packLockboxInitCode(cfg lockboxDeploymentConfig, impl common.Address) ([]byte, error) {
	lockboxAbi, err := bindings.LockboxMetaData.GetAbi()
	if err != nil {
		return nil, errors.Wrap(err, "get abi")
	}

	proxyAbi, err := bindings.TransparentUpgradeableProxyMetaData.GetAbi()
	if err != nil {
		return nil, errors.Wrap(err, "get proxy abi")
	}

	initializer, err := lockboxAbi.Pack("initialize", cfg.Admin, cfg.Pauser, cfg.Token, cfg.Wrapped)
	if err != nil {
		return nil, errors.Wrap(err, "encode initializer")
	}

	return contracts.PackInitCode(proxyAbi, bindings.TransparentUpgradeableProxyBin, impl, cfg.ProxyAdminOwner, initializer)
}
