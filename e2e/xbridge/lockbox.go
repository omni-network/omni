package xbridge

import (
	"context"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/lib/contracts"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/netconf"

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

func (cfg lockboxDeploymentConfig) validateLockboxConfig() error {
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

// deployLockboxIfNeeded deploys a new lockbox contract if it is not already deployed.
// If the contract is already deployed, the receipt is nil.
func DeployLockboxIfNeeded(ctx context.Context, network netconf.ID, backend *ethbackend.Backend) (common.Address, *ethtypes.Receipt, error) {
	addrs, err := contracts.GetAddresses(ctx, network)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "get addrs")
	}

	deployed, addr, err := isDeployed(ctx, backend, addrs.Lockbox)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "is deployed")
	}
	if deployed {
		return addr, nil, nil
	}

	return deployLockbox(ctx, network, backend)
}

// deployLockbox deploys a new lockbox contract and returns the address and receipt.
func deployLockbox(ctx context.Context, network netconf.ID, backend *ethbackend.Backend) (common.Address, *ethtypes.Receipt, error) {
	addrs, err := contracts.GetAddresses(ctx, network)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "get addrs")
	}

	salts, err := contracts.GetSalts(ctx, network)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "get salts")
	}

	deployCfg := deploymentConfig{
		Create3Salt:    salts.Lockbox,
		Create3Factory: addrs.Create3Factory,
		ExpectedAddr:   addrs.Lockbox,
		Deployer:       eoa.MustAddress(network, eoa.RoleDeployer),
	}

	cfg := lockboxDeploymentConfig{
		Config:          deployCfg,
		ProxyAdminOwner: eoa.MustAddress(network, eoa.RoleUpgrader),
		Admin:           eoa.MustAddress(network, eoa.RoleManager),
		Pauser:          eoa.MustAddress(network, eoa.RoleManager),
		Token:           addrs.RLUSD,
		Wrapped:         addrs.RLUSDe,
	}

	return performLockboxDeployment(ctx, network, backend, cfg)
}

// performLockboxDeployment handles the common deployment flow for the lockbox contract.
func performLockboxDeployment(ctx context.Context, network netconf.ID, backend *ethbackend.Backend, cfg lockboxDeploymentConfig) (common.Address, *ethtypes.Receipt, error) {
	if err := cfg.validateLockboxConfig(); err != nil {
		return common.Address{}, nil, errors.Wrap(err, "validate config")
	}

	factory, addr, salt, txOpts, err := deployPrep(ctx, network, backend, cfg.Config)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "deploy prep")
	}

	impl, tx, _, err := bindings.DeployLockbox(txOpts, backend)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "deploy impl")
	}

	_, err = backend.WaitMined(ctx, tx)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "wait mined impl")
	}

	initCode, err := packLockboxInitCode(cfg, impl)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "pack init code")
	}

	tx, err = factory.DeployWithRetry(txOpts, salt, initCode) //nolint:contextcheck // Context is txOpts
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "deploy proxy")
	}

	receipt, err := backend.WaitMined(ctx, tx)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "wait mined proxy")
	}

	return addr, receipt, nil
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
