package portal

import (
	"context"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/chainids"
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
	ProxyAdmin     common.Address
	Deployer       common.Address
	Owner          common.Address
	FeeOracle      common.Address
	ValSetID       uint64
	Validators     []bindings.Validator
	ExpectedAddr   common.Address
}

func (cfg DeploymentConfig) Validate() error {
	if (cfg.Create3Factory == common.Address{}) {
		return errors.New("create3 factory is zero")
	}
	if cfg.Create3Salt == "" {
		return errors.New("create3 salt is empty")
	}
	if (cfg.ProxyAdmin == common.Address{}) {
		return errors.New("proxy admin is zero")
	}
	if (cfg.Deployer == common.Address{}) {
		return errors.New("deployer is zero")
	}
	if (cfg.Owner == common.Address{}) {
		return errors.New("owner is zero")
	}
	if (cfg.FeeOracle == common.Address{}) {
		return errors.New("fee oracle is zero")
	}
	if cfg.ValSetID == 0 {
		return errors.New("validator set ID is zero")
	}
	if len(cfg.Validators) == 0 {
		return errors.New("validators is empty")
	}

	return nil
}

func getDeployCfg(chainID uint64, network string) (DeploymentConfig, error) {
	if chainids.IsMainnet(chainID) && network == netconf.Mainnet {
		return mainnetDeployCfg(), nil
	}

	if chainids.IsTestnet(chainID) && network == netconf.Testnet {
		return testnetDeployCfg(), nil
	}

	return DeploymentConfig{}, errors.New("unsupported chain for network", "chain_id", chainID, "network", network)
}

func mainnetDeployCfg() DeploymentConfig {
	return DeploymentConfig{
		Create3Factory: contracts.MainnetCreate3Factory,
		Create3Salt:    contracts.PortalSalt(netconf.Mainnet),
		Owner:          contracts.MainnetPortalAdmin,
		Deployer:       contracts.MainnetDeployer,
		// TODO: fill in the rest
	}
}

func testnetDeployCfg() DeploymentConfig {
	return DeploymentConfig{
		Create3Factory: contracts.TestnetCreate3Factory,
		Create3Salt:    contracts.PortalSalt(netconf.Testnet),
		Owner:          contracts.TestnetPortalAdmin,
		Deployer:       contracts.TestnetDeployer,
		// TODO: fill in the rest
	}
}

func devnetDeployCfg(vals []bindings.Validator) DeploymentConfig {
	return DeploymentConfig{
		Create3Factory: contracts.DevnetCreate3Factory,
		Create3Salt:    contracts.PortalSalt(netconf.Devnet),
		FeeOracle:      contracts.DevnetFeeOracleV1,
		Owner:          contracts.DevnetPortalAdmin,
		Deployer:       contracts.DevnetDeployer,
		ProxyAdmin:     contracts.DevnetProxyAdmin,
		ValSetID:       1,
		Validators:     vals,
		ExpectedAddr:   contracts.DevnetPortal,
	}
}

// Deploy deploys a new Portal contract and returns the address and receipt.
// It only allows deployments to explicitly supported chains.
func Deploy(ctx context.Context, network string, backend *ethbackend.Backend) (common.Address, *ethtypes.Receipt, error) {
	chainID, err := backend.ChainID(ctx)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "chain id")
	}

	cfg, err := getDeployCfg(chainID.Uint64(), network)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "get deployment config")
	}

	return deploy(ctx, cfg, backend)
}

// DeployDevnet deploys the devnet Portal and returns the address receipt.
func DeployDevnet(ctx context.Context, backend *ethbackend.Backend, vals []bindings.Validator) (common.Address, *ethtypes.Receipt, error) {
	chainID, err := backend.ChainID(ctx)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "chain id")
	}

	if chainids.IsMainnetOrTestnet(chainID.Uint64()) {
		return common.Address{}, nil, errors.New("not a devnet")
	}

	return deploy(ctx, devnetDeployCfg(vals), backend)
}

func deploy(ctx context.Context, cfg DeploymentConfig, backend *ethbackend.Backend) (common.Address, *ethtypes.Receipt, error) {
	if err := cfg.Validate(); err != nil {
		return common.Address{}, nil, errors.Wrap(err, "validate")
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

	impl, tx, _, err := bindings.DeployOmniPortal(txOpts, backend)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "deploy portal impl")
	}

	receipt, err := bind.WaitMined(ctx, backend, tx)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "wait mined portal")
	} else if receipt.Status != ethtypes.ReceiptStatusSuccessful {
		return common.Address{}, nil, errors.New("deploy portal failed")
	}

	initCode, err := packInitCode(cfg, impl)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "pack init code")
	}

	tx, err = factory.Deploy(txOpts, salt, initCode)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "deploy proxy admin")
	}

	receipt, err = bind.WaitMined(ctx, backend, tx)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "wait mined upgradable proxy")
	} else if receipt.Status != ethtypes.ReceiptStatusSuccessful {
		return common.Address{}, nil, errors.New("deploy upgradable proxy failed")
	}

	return addr, receipt, nil
}

func packInitCode(cfg DeploymentConfig, impl common.Address) ([]byte, error) {
	portalAbi, err := bindings.OmniPortalMetaData.GetAbi()
	if err != nil {
		return nil, errors.Wrap(err, "get portal abi")
	}

	proxyAbi, err := bindings.TransparentUpgradeableProxyMetaData.GetAbi()
	if err != nil {
		return nil, errors.Wrap(err, "get proxy abi")
	}

	initializer, err := portalAbi.Pack("initialize", cfg.Owner, cfg.FeeOracle, cfg.ValSetID, cfg.Validators)
	if err != nil {
		return nil, errors.Wrap(err, "encode portal initializer")
	}

	return contracts.PackInitCode(proxyAbi, bindings.TransparentUpgradeableProxyBin, impl, cfg.ProxyAdmin, initializer)
}
