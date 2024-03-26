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
	if cfg.ValSetID == 0 {
		return errors.New("validator set ID is zero")
	}
	if len(cfg.Validators) == 0 {
		return errors.New("validators is empty")
	}

	return nil
}

func getDeployCfg(chainID uint64, network netconf.ID) (DeploymentConfig, error) {
	if !chainids.IsMainnetOrTestnet(chainID) && network == netconf.Devnet {
		return devnetCfg(), nil
	}

	if chainids.IsMainnet(chainID) && network == netconf.Mainnet {
		return mainnetCfg(), nil
	}

	if chainids.IsTestnet(chainID) && network == netconf.Testnet {
		return testnetCfg(), nil
	}

	if !chainids.IsMainnet(chainID) && network == netconf.Staging {
		return stagingCfg(), nil
	}

	return DeploymentConfig{}, errors.New("unsupported chain for network", "chain_id", chainID, "network", network)
}

func mainnetCfg() DeploymentConfig {
	return DeploymentConfig{
		Create3Factory: contracts.MainnetCreate3Factory(),
		Create3Salt:    contracts.PortalSalt(netconf.Mainnet),
		Owner:          contracts.MainnetPortalAdmin(),
		Deployer:       contracts.MainnetDeployer(),
		// TODO: fill in the rest
	}
}

func testnetCfg() DeploymentConfig {
	return DeploymentConfig{
		Create3Factory: contracts.TestnetCreate3Factory(),
		Create3Salt:    contracts.PortalSalt(netconf.Testnet),
		Owner:          contracts.TestnetPortalAdmin(),
		Deployer:       contracts.TestnetDeployer(),
		// TODO: fill in the rest
	}
}

func stagingCfg() DeploymentConfig {
	return DeploymentConfig{
		Create3Factory: contracts.StagingCreate3Factory(),
		Create3Salt:    contracts.PortalSalt(netconf.Staging),
		Owner:          contracts.StagingPortalAdmin(),
		Deployer:       contracts.StagingDeployer(),
		ProxyAdmin:     contracts.StagingProxyAdmin(),
		ExpectedAddr:   contracts.StagingPortalAdmin(),
	}
}

func devnetCfg() DeploymentConfig {
	return DeploymentConfig{
		Create3Factory: contracts.DevnetCreate3Factory(),
		Create3Salt:    contracts.PortalSalt(netconf.Devnet),
		Owner:          contracts.DevnetPortalAdmin(),
		Deployer:       contracts.DevnetDeployer(),
		ProxyAdmin:     contracts.DevnetProxyAdmin(),
		ExpectedAddr:   contracts.DevnetPortal(),
	}
}

func (cfg *DeploymentConfig) addValidators(valSetID uint64, validators []bindings.Validator) {
	cfg.ValSetID = valSetID
	cfg.Validators = validators
}

func AddrForNetwork(network netconf.ID) (common.Address, bool) {
	switch network {
	case netconf.Mainnet:
		return contracts.MainnetPortal(), true
	case netconf.Testnet:
		return contracts.TestnetPortal(), true
	case netconf.Staging:
		return contracts.StagingPortal(), true
	case netconf.Devnet:
		return contracts.DevnetPortal(), true
	default:
		return common.Address{}, false
	}
}

// IsDeployed checks if the Portal contract is deployed to the provided backend
// to its expected network address.
func IsDeployed(ctx context.Context, network netconf.ID, backend *ethbackend.Backend) (bool, contracts.Deployment, error) {
	chainID, err := backend.ChainID(ctx)
	if err != nil {
		return false, contracts.Deployment{}, errors.Wrap(err, "chain id")
	}

	cfg, err := getDeployCfg(chainID.Uint64(), network)
	if err != nil {
		return false, contracts.Deployment{}, errors.Wrap(err, "get deployment config")
	}

	factory, err := bindings.NewCreate3(cfg.Create3Factory, backend)
	if err != nil {
		return false, contracts.Deployment{}, errors.Wrap(err, "new create3")
	}

	salt := create3.HashSalt(cfg.Create3Salt)
	height, err := factory.GetDeployedHeight(nil, cfg.Deployer, salt)
	if err != nil {
		return false, contracts.Deployment{}, errors.Wrap(err, "get deployed height")
	}

	if height.Uint64() == 0 {
		return false, contracts.Deployment{}, nil
	}

	deployment := contracts.Deployment{
		Address:     create3.Address(cfg.Create3Factory, cfg.Create3Salt, cfg.Deployer),
		BlockHeight: height.Uint64(),
	}

	return true, deployment, nil
}

// DeployIfNeeded deploys a new Portal contract if it is not already deployed.
func DeployIfNeeded(ctx context.Context, network netconf.ID, backend *ethbackend.Backend, valSetID uint64, validators []bindings.Validator,
) (contracts.Deployment, error) {
	deployed, deployment, err := IsDeployed(ctx, network, backend)
	if err != nil {
		return contracts.Deployment{}, errors.Wrap(err, "is deployed")
	}
	if deployed {
		return deployment, nil
	}

	return Deploy(ctx, network, backend, valSetID, validators)
}

// Deploy deploys a new Portal contract and returns the address and receipt.
// It only allows deployments to explicitly supported chains.
func Deploy(ctx context.Context, network netconf.ID, backend *ethbackend.Backend, valSetID uint64, validators []bindings.Validator,
) (contracts.Deployment, error) {
	chainID, err := backend.ChainID(ctx)
	if err != nil {
		return contracts.Deployment{}, errors.Wrap(err, "chain id")
	}

	cfg, err := getDeployCfg(chainID.Uint64(), network)
	if err != nil {
		return contracts.Deployment{}, errors.Wrap(err, "get deployment config")
	}

	return deploy(ctx, cfg, backend, valSetID, validators)
}

func deploy(ctx context.Context, cfg DeploymentConfig, backend *ethbackend.Backend, valSetID uint64, validators []bindings.Validator,
) (contracts.Deployment, error) {
	cfg.addValidators(valSetID, validators)

	if err := cfg.Validate(); err != nil {
		return contracts.Deployment{}, errors.Wrap(err, "validate")
	}

	txOpts, err := backend.BindOpts(ctx, cfg.Deployer)
	if err != nil {
		return contracts.Deployment{}, errors.Wrap(err, "bind opts")
	}

	factory, err := bindings.NewCreate3(cfg.Create3Factory, backend)
	if err != nil {
		return contracts.Deployment{}, errors.Wrap(err, "new create3")
	}

	salt := create3.HashSalt(cfg.Create3Salt)

	addr, err := factory.GetDeployedAddr(nil, txOpts.From, salt)
	if err != nil {
		return contracts.Deployment{}, errors.Wrap(err, "get deployed")
	} else if (cfg.ExpectedAddr != common.Address{}) && addr != cfg.ExpectedAddr {
		return contracts.Deployment{}, errors.New("unexpected address", "expected", cfg.ExpectedAddr, "actual", addr)
	}

	feeOracle, tx, _, err := bindings.DeployFeeOracleV1(txOpts, backend)
	if err != nil {
		return contracts.Deployment{}, errors.Wrap(err, "deploy fee oracle")
	}

	receipt, err := bind.WaitMined(ctx, backend, tx)
	if err != nil {
		return contracts.Deployment{}, errors.Wrap(err, "wait mined fee oracle")
	} else if receipt.Status != ethtypes.ReceiptStatusSuccessful {
		return contracts.Deployment{}, errors.New("deploy fee oracle failed")
	}

	impl, tx, _, err := bindings.DeployOmniPortal(txOpts, backend)
	if err != nil {
		return contracts.Deployment{}, errors.Wrap(err, "deploy impl")
	}

	receipt, err = bind.WaitMined(ctx, backend, tx)
	if err != nil {
		return contracts.Deployment{}, errors.Wrap(err, "wait mined portal")
	} else if receipt.Status != ethtypes.ReceiptStatusSuccessful {
		return contracts.Deployment{}, errors.New("deploy impl failed")
	}

	initCode, err := packInitCode(cfg, feeOracle, impl)
	if err != nil {
		return contracts.Deployment{}, errors.Wrap(err, "pack init code")
	}

	tx, err = factory.Deploy(txOpts, salt, initCode)
	if err != nil {
		return contracts.Deployment{}, errors.Wrap(err, "deploy proxy")
	}

	receipt, err = bind.WaitMined(ctx, backend, tx)
	if err != nil {
		return contracts.Deployment{}, errors.Wrap(err, "wait mined proxy")
	} else if receipt.Status != ethtypes.ReceiptStatusSuccessful {
		return contracts.Deployment{}, errors.New("deploy proxy failed")
	}

	deployment := contracts.Deployment{
		Address:     addr,
		BlockHeight: receipt.BlockNumber.Uint64(),
	}

	return deployment, nil
}

func packInitCode(cfg DeploymentConfig, feeOracle common.Address, impl common.Address) ([]byte, error) {
	portalAbi, err := bindings.OmniPortalMetaData.GetAbi()
	if err != nil {
		return nil, errors.Wrap(err, "get portal abi")
	}

	proxyAbi, err := bindings.TransparentUpgradeableProxyMetaData.GetAbi()
	if err != nil {
		return nil, errors.Wrap(err, "get proxy abi")
	}

	initializer, err := portalAbi.Pack("initialize", cfg.Owner, feeOracle, cfg.ValSetID, cfg.Validators)
	if err != nil {
		return nil, errors.Wrap(err, "encode portal initializer")
	}

	return contracts.PackInitCode(proxyAbi, bindings.TransparentUpgradeableProxyBin, impl, cfg.ProxyAdmin, initializer)
}
