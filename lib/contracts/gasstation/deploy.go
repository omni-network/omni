package gasstation

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
	Deployer        common.Address
	Portal          common.Address
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
	if isDeadOrEmpty(cfg.ProxyAdminOwner) {
		return errors.New("proxy admin is zero")
	}
	if isDeadOrEmpty(cfg.Deployer) {
		return errors.New("deployer is not set")
	}
	if isDeadOrEmpty(cfg.Owner) {
		return errors.New("owner is not set")
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
	addrs, err := contracts.GetAddresses(ctx, network)
	if err != nil {
		return false, common.Address{}, errors.Wrap(err, "get addresses")
	}

	code, err := backend.CodeAt(ctx, addrs.GasStation, nil)
	if err != nil {
		return false, addrs.GasStation, errors.Wrap(err, "code at", "address", addrs.GasStation)
	}

	if len(code) == 0 {
		return false, addrs.GasStation, nil
	}

	return true, addrs.GasStation, nil
}

// DeployIfNeeded deploys a new token contract if it is not already deployed.
// If the contract is already deployed, the receipt is nil.
func DeployIfNeeded(
	ctx context.Context,
	network netconf.ID,
	backend *ethbackend.Backend,
	gasPumps []bindings.OmniGasStationGasPump,
) (common.Address, *ethtypes.Receipt, error) {
	deployed, addr, err := isDeployed(ctx, network, backend)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "is deployed")
	}
	if deployed {
		return addr, nil, nil
	}

	return Deploy(ctx, network, backend, gasPumps)
}

// Deploy deploys a new L1Bridge contract and returns the address and receipt.
func Deploy(
	ctx context.Context,
	network netconf.ID,
	backend *ethbackend.Backend,
	gasPumps []bindings.OmniGasStationGasPump,
) (common.Address, *ethtypes.Receipt, error) {
	addrs, err := contracts.GetAddresses(ctx, network)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "get addresses")
	}

	salts, err := contracts.GetSalts(ctx, network)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "get salts")
	}

	cfg := DeploymentConfig{
		Create3Factory:  addrs.Create3Factory,
		Create3Salt:     salts.GasStation,
		Owner:           eoa.MustAddress(network, eoa.RoleManager),
		Deployer:        eoa.MustAddress(network, eoa.RoleDeployer),
		ProxyAdminOwner: eoa.MustAddress(network, eoa.RoleUpgrader),
		Portal:          addrs.Portal,
		ExpectedAddr:    addrs.GasStation,
	}

	return deploy(ctx, network, cfg, backend, gasPumps)
}

func deploy(
	ctx context.Context,
	network netconf.ID,
	cfg DeploymentConfig,
	backend *ethbackend.Backend,
	gasPumps []bindings.OmniGasStationGasPump,
) (common.Address, *ethtypes.Receipt, error) {
	if err := cfg.Validate(); err != nil {
		return common.Address{}, nil, errors.Wrap(err, "validate config")
	}

	chainID, err := backend.ChainID(ctx)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "chain id")
	}

	if chainID.Uint64() != network.Static().OmniExecutionChainID {
		return common.Address{}, nil, errors.New("must deploy on omni evm")
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

	impl, tx, _, err := bindings.DeployOmniGasStation(txOpts, backend)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "deploy impl")
	}

	_, err = backend.WaitMined(ctx, tx)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "wait mined impl")
	}

	initCode, err := packInitCode(cfg, impl, gasPumps)
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

func packInitCode(cfg DeploymentConfig, impl common.Address, gasPumps []bindings.OmniGasStationGasPump) ([]byte, error) {
	gasStationAbi, err := bindings.OmniGasStationMetaData.GetAbi()
	if err != nil {
		return nil, errors.Wrap(err, "get abi")
	}

	proxyAbi, err := bindings.TransparentUpgradeableProxyMetaData.GetAbi()
	if err != nil {
		return nil, errors.Wrap(err, "get proxy abi")
	}

	initializer, err := gasStationAbi.Pack("initialize", cfg.Portal, cfg.Owner, gasPumps)
	if err != nil {
		return nil, errors.Wrap(err, "encode initializer")
	}

	return contracts.PackInitCode(proxyAbi, bindings.TransparentUpgradeableProxyBin, impl, cfg.ProxyAdminOwner, initializer)
}
