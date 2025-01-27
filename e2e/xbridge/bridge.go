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

// deployBridgeIfNeeded deploys a new bridge contract if it is not already deployed.
// If the contract is already deployed, the receipt is nil.
func DeployBridgeIfNeeded(ctx context.Context, network netconf.ID, backend *ethbackend.Backend, lockbox bool) (common.Address, *ethtypes.Receipt, error) {
	addrs, err := contracts.GetAddresses(ctx, network)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "get addrs")
	}

	deployed, addr, err := isDeployed(ctx, backend, addrs.Bridge)
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

	salts, err := contracts.GetSalts(ctx, network)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "get salts")
	}

	lockboxAddr := common.Address{}
	if lockbox {
		lockboxAddr = addrs.Lockbox
	}

	deployCfg := deploymentConfig{
		Create3Salt:    salts.Bridge,
		Create3Factory: addrs.Create3Factory,
		ExpectedAddr:   addrs.Bridge,
		Deployer:       eoa.MustAddress(network, eoa.RoleDeployer),
	}

	cfg := bridgeDeploymentConfig{
		Config:          deployCfg,
		ProxyAdminOwner: eoa.MustAddress(network, eoa.RoleUpgrader),
		Admin:           eoa.MustAddress(network, eoa.RoleManager),
		Pauser:          eoa.MustAddress(network, eoa.RoleManager),
		OmniPortal:      addrs.Portal,
		Token:           addrs.RLUSDe,
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
