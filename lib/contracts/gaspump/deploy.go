package gaspump

import (
	"context"
	"math/big"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/lib/contracts"
	"github.com/omni-network/omni/lib/create3"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

type DeploymentConfig struct {
	Create3Factory common.Address
	Create3Salt    string
	Upgrader       common.Address
	Manager        common.Address
	Deployer       common.Address
	Portal         common.Address
	GasStation     common.Address
	Oracle         common.Address
	MaxSwap        *big.Int
	Toll           *big.Int
	ExpectedAddr   common.Address
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
	if isDeadOrEmpty(cfg.Upgrader) {
		return errors.New("proxy admin is zero")
	}
	if isDeadOrEmpty(cfg.Deployer) {
		return errors.New("deployer is not set")
	}
	if isDeadOrEmpty(cfg.Manager) {
		return errors.New("owner is not set")
	}
	if (cfg.Portal == common.Address{}) {
		return errors.New("portal is zero")
	}
	if (cfg.GasStation == common.Address{}) {
		return errors.New("gas station is zero")
	}
	if (cfg.Oracle == common.Address{}) {
		return errors.New("oracle is zero")
	}
	if cfg.MaxSwap == nil {
		return errors.New("max swap is nil")
	}
	if cfg.Toll == nil {
		return errors.New("toll is nil")
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
		return false, common.Address{}, errors.Wrap(err, "get addrs")
	}

	code, err := backend.CodeAt(ctx, addrs.GasPump, nil)
	if err != nil {
		return false, addrs.GasPump, errors.Wrap(err, "code at", "address", addrs.GasPump)
	}

	if len(code) == 0 {
		return false, addrs.GasPump, nil
	}

	return true, addrs.GasPump, nil
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
	addrs, err := contracts.GetAddresses(ctx, network)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "get addresses")
	}

	salts, err := contracts.GetSalts(ctx, network)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "get salts")
	}

	portal, err := bindings.NewOmniPortal(addrs.Portal, backend)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "new portal")
	}

	oracle, err := portal.FeeOracle(&bind.CallOpts{Context: ctx})
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "fee oracle")
	}

	cfg := DeploymentConfig{
		Create3Factory: addrs.Create3Factory,
		Create3Salt:    salts.GasPump,
		Manager:        eoa.MustAddress(network, eoa.RoleManager),
		Deployer:       eoa.MustAddress(network, eoa.RoleDeployer),
		Upgrader:       eoa.MustAddress(network, eoa.RoleUpgrader),
		Portal:         addrs.Portal,
		GasStation:     addrs.GasStation,
		Oracle:         oracle,
		MaxSwap:        big.NewInt(20000000000000000), // 0.02 ETH
		Toll:           big.NewInt(100),               // 100 / 1000 = 0.1 = 10% (1000 = GasPump.TOLL_DENOM),
		ExpectedAddr:   addrs.GasPump,
	}

	return deploy(ctx, network, cfg, backend)
}

func deploy(ctx context.Context, network netconf.ID, cfg DeploymentConfig, backend *ethbackend.Backend) (common.Address, *ethtypes.Receipt, error) {
	if err := cfg.Validate(); err != nil {
		return common.Address{}, nil, errors.Wrap(err, "validate config")
	}

	chainID, err := backend.ChainID(ctx)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "chain id")
	}

	if chainID.Uint64() == network.Static().OmniExecutionChainID {
		return common.Address{}, nil, errors.New("cannot deploy on omni evm")
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

	impl, tx, _, err := bindings.DeployOmniGasPump(txOpts, backend)
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
	gasPumpAbi, err := bindings.OmniGasPumpMetaData.GetAbi()
	if err != nil {
		return nil, errors.Wrap(err, "get abi")
	}

	proxyAbi, err := bindings.TransparentUpgradeableProxyMetaData.GetAbi()
	if err != nil {
		return nil, errors.Wrap(err, "get proxy abi")
	}

	initParams := bindings.OmniGasPumpInitParams{
		GasStation: cfg.GasStation,
		Oracle:     cfg.Oracle,
		Portal:     cfg.Portal,
		Manager:    cfg.Manager,
		MaxSwap:    cfg.MaxSwap,
		Toll:       cfg.Toll,
	}

	initializer, err := gasPumpAbi.Pack("initialize", initParams)
	if err != nil {
		return nil, errors.Wrap(err, "encode initializer")
	}

	return contracts.PackInitCode(proxyAbi, bindings.TransparentUpgradeableProxyBin, impl, cfg.Upgrader, initializer)
}
