package l1bridge

import (
	"context"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/lib/contracts"
	"github.com/omni-network/omni/lib/create3"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

type DeploymentConfig struct {
	Create3Factory  common.Address
	Create3Salt     string
	ProxyAdminOwner common.Address
	Owner           common.Address
	Portal          common.Address
	Token           common.Address
	Deployer        common.Address
	ExpectedAddr    common.Address
}

func isDeadOrEmpty(addr common.Address) bool {
	return addr == common.Address{} || addr == common.HexToAddress(eoa.ZeroXDead)
}

func (cfg DeploymentConfig) Validate() error {
	if (cfg.Create3Factory == common.Address{}) {
		return errors.New("create3 factory is zero")
	}
	if cfg.Create3Salt == "" {
		return errors.New("create3 salt is empty")
	}
	if (cfg.ProxyAdminOwner == common.Address{}) {
		return errors.New("proxy admin is zero")
	}
	if isDeadOrEmpty(cfg.Deployer) {
		return errors.New("deployer is not set")
	}
	if isDeadOrEmpty(cfg.Owner) {
		return errors.New("owner is not set")
	}
	if (cfg.Token == common.Address{}) {
		return errors.New("token is zero")
	}
	if (cfg.Portal == common.Address{}) {
		return errors.New("portal is zero")
	}
	if (cfg.ExpectedAddr == common.Address{}) {
		return errors.New("expected address is zero")
	}

	return nil
}

// isDeployed returns true if the token contract is already deployed to its expected address.
func isDeployed(ctx context.Context, network netconf.ID, backend *ethbackend.Backend) (bool, common.Address, error) {
	addr := contracts.L1Bridge(network)

	code, err := backend.CodeAt(ctx, addr, nil)
	if err != nil {
		return false, addr, errors.Wrap(err, "code at", "address", addr)
	}

	if len(code) == 0 {
		return false, addr, nil
	}

	return true, addr, nil
}

// DeployIfNeeded deploys a new token contract if it is not already deployed.
// If the contract is already deployed, the receipt is nil.
func DeployIfNeeded(ctx context.Context, network netconf.ID, backend *ethbackend.Backend) (common.Address, *ethtypes.Receipt, error) {
	deployed, addr, err := isDeployed(ctx, network, backend)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "is deployed")
	}
	if deployed {
		return addr, nil, nil
	}

	return Deploy(ctx, network, backend)
}

// Deploy deploys a new L1Bridge contract and returns the address and receipt.
func Deploy(ctx context.Context, network netconf.ID, backend *ethbackend.Backend) (common.Address, *ethtypes.Receipt, error) {
	cfg := DeploymentConfig{
		Create3Factory:  contracts.Create3Factory(network),
		Create3Salt:     contracts.L1BridgeSalt(network),
		Owner:           eoa.MustAddress(network, eoa.RoleAdmin),
		Deployer:        eoa.MustAddress(network, eoa.RoleDeployer),
		ProxyAdminOwner: eoa.MustAddress(network, eoa.RoleAdmin),
		Portal:          contracts.Portal(network),
		Token:           contracts.Token(network),
		ExpectedAddr:    contracts.L1Bridge(network),
	}

	return deploy(ctx, cfg, backend)
}

func deploy(ctx context.Context, cfg DeploymentConfig, backend *ethbackend.Backend) (common.Address, *ethtypes.Receipt, error) {
	if err := cfg.Validate(); err != nil {
		return common.Address{}, nil, errors.Wrap(err, "validate config")
	}

	txOpts, err := backend.BindOpts(ctx, cfg.Deployer)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "bind opts")
	}

	factory, err := bindings.NewCreate3(cfg.Create3Factory, backend)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "new create3")
	}

	salt := create3.HashSalt(cfg.Create3Salt)

	addr, err := factory.GetDeployed(nil, txOpts.From, salt)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "get deployed")
	} else if (cfg.ExpectedAddr != common.Address{}) && addr != cfg.ExpectedAddr {
		return common.Address{}, nil, errors.New("unexpected address", "expected", cfg.ExpectedAddr, "actual", addr)
	}

	impl, tx, _, err := bindings.DeployOmniBridgeL1(txOpts, backend, cfg.Token)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "deploy impl")
	}

	_, err = backend.WaitMined(ctx, tx)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "wait mined impl")
	}

	initCode, err := packInitCode(cfg, impl)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "pack init code")
	}

	tx, err = factory.Deploy(txOpts, salt, initCode)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "deploy proxy")
	}

	receipt, err := backend.WaitMined(ctx, tx)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "wait mined proxy")
	}

	return addr, receipt, nil
}

func packInitCode(cfg DeploymentConfig, impl common.Address) ([]byte, error) {
	bridgeAbi, err := bindings.OmniBridgeL1MetaData.GetAbi()
	if err != nil {
		return nil, errors.Wrap(err, "get abi")
	}

	proxyAbi, err := bindings.TransparentUpgradeableProxyMetaData.GetAbi()
	if err != nil {
		return nil, errors.Wrap(err, "get proxy abi")
	}

	initializer, err := bridgeAbi.Pack("initialize", cfg.Owner, cfg.Portal)
	if err != nil {
		return nil, errors.Wrap(err, "encode initializer")
	}

	return contracts.PackInitCode(proxyAbi, bindings.TransparentUpgradeableProxyBin, impl, cfg.ProxyAdminOwner, initializer)
}
