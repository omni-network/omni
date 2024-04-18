package portal

import (
	"context"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/e2e/app/eoa"
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
	OmniEChainID   uint64
	OmniCChainID   uint64
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
	if cfg.OmniEChainID == 0 {
		return errors.New("omni EVM chain ID is zero")
	}
	if cfg.OmniCChainID == 0 {
		return errors.New("omni cons chain ID is zero")
	}
	if (cfg.ExpectedAddr == common.Address{}) {
		return errors.New("expected address is zero")
	}

	return nil
}

func getDeployCfg(chainID uint64, network netconf.ID, valSetID uint64, vals []bindings.Validator) (DeploymentConfig, error) {
	if !chainids.IsMainnetOrTestnet(chainID) && network == netconf.Devnet {
		return devnetCfg(valSetID, vals), nil
	}

	if chainids.IsMainnet(chainID) && network == netconf.Mainnet {
		return mainnetCfg(), nil
	}

	if chainids.IsTestnet(chainID) && network == netconf.Testnet {
		return testnetCfg(valSetID, vals), nil
	}

	if !chainids.IsMainnet(chainID) && network == netconf.Staging {
		return stagingCfg(valSetID, vals), nil
	}

	return DeploymentConfig{}, errors.New("unsupported chain for network", "chain_id", chainID, "network", network)
}

func mainnetCfg() DeploymentConfig {
	return DeploymentConfig{
		Create3Factory: contracts.MainnetCreate3Factory(),
		Create3Salt:    contracts.PortalSalt(netconf.Mainnet),
		Owner:          eoa.MustAddress(netconf.Mainnet, eoa.RolePortalAdmin),
		Deployer:       eoa.MustAddress(netconf.Mainnet, eoa.RoleDeployer),
		// TODO: fill in the rest
	}
}

func testnetCfg(valSetID uint64, vals []bindings.Validator) DeploymentConfig {
	return DeploymentConfig{
		Create3Factory: contracts.TestnetCreate3Factory(),
		Create3Salt:    contracts.PortalSalt(netconf.Testnet),
		Owner:          eoa.MustAddress(netconf.Testnet, eoa.RolePortalAdmin),
		Deployer:       eoa.MustAddress(netconf.Testnet, eoa.RoleDeployer),
		ProxyAdmin:     contracts.TestnetProxyAdmin(),
		OmniEChainID:   netconf.Testnet.Static().OmniExecutionChainID,
		OmniCChainID:   netconf.Testnet.Static().OmniConsensusChainIDUint64(),
		ValSetID:       valSetID,
		Validators:     vals,
		ExpectedAddr:   contracts.TestnetPortal(),
	}
}

func stagingCfg(valSetID uint64, vals []bindings.Validator) DeploymentConfig {
	return DeploymentConfig{
		Create3Factory: contracts.StagingCreate3Factory(),
		Create3Salt:    contracts.PortalSalt(netconf.Staging),
		Owner:          eoa.MustAddress(netconf.Staging, eoa.RolePortalAdmin),
		Deployer:       eoa.MustAddress(netconf.Staging, eoa.RoleDeployer),
		ProxyAdmin:     contracts.StagingProxyAdmin(),
		OmniEChainID:   netconf.Staging.Static().OmniExecutionChainID,
		OmniCChainID:   netconf.Staging.Static().OmniConsensusChainIDUint64(),
		ValSetID:       valSetID,
		Validators:     vals,
		ExpectedAddr:   contracts.StagingPortal(),
	}
}

func devnetCfg(valSetID uint64, vals []bindings.Validator) DeploymentConfig {
	return DeploymentConfig{
		Create3Factory: contracts.DevnetCreate3Factory(),
		Create3Salt:    contracts.PortalSalt(netconf.Devnet),
		Owner:          eoa.MustAddress(netconf.Devnet, eoa.RolePortalAdmin),
		Deployer:       eoa.MustAddress(netconf.Devnet, eoa.RoleDeployer),
		ProxyAdmin:     contracts.DevnetProxyAdmin(),
		OmniEChainID:   netconf.Devnet.Static().OmniExecutionChainID,
		OmniCChainID:   netconf.Devnet.Static().OmniConsensusChainIDUint64(),
		ValSetID:       valSetID,
		Validators:     vals,
		ExpectedAddr:   contracts.DevnetPortal(),
	}
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
func IsDeployed(ctx context.Context, network netconf.ID, backend *ethbackend.Backend) (bool, common.Address, error) {
	addr, ok := AddrForNetwork(network)
	if !ok {
		return false, addr, errors.New("unsupported network", "network", network)
	}

	code, err := backend.CodeAt(ctx, addr, nil)
	if err != nil {
		return false, addr, errors.Wrap(err, "code at")
	}

	if len(code) == 0 {
		return false, addr, nil
	}

	return true, addr, nil
}

// Deploy deploys a new Portal contract and returns the address and receipt.
// It only allows deployments to explicitly supported chains.
func Deploy(ctx context.Context, network netconf.ID, backend *ethbackend.Backend, valSetID uint64, validators []bindings.Validator,
) (common.Address, *ethtypes.Receipt, error) {
	chainID, err := backend.ChainID(ctx)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "chain id")
	}

	cfg, err := getDeployCfg(chainID.Uint64(), network, valSetID, validators)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "get deployment config")
	}

	return deploy(ctx, cfg, backend)
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

	feeOracle, tx, _, err := bindings.DeployFeeOracleV1(txOpts, backend)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "deploy fee oracle")
	}

	receipt, err := bind.WaitMined(ctx, backend, tx)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "wait mined fee oracle")
	} else if receipt.Status != ethtypes.ReceiptStatusSuccessful {
		return common.Address{}, nil, errors.New("deploy fee oracle failed")
	}

	impl, tx, _, err := bindings.DeployOmniPortal(txOpts, backend)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "deploy impl")
	}

	receipt, err = bind.WaitMined(ctx, backend, tx)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "wait mined portal")
	} else if receipt.Status != ethtypes.ReceiptStatusSuccessful {
		return common.Address{}, nil, errors.New("deploy impl failed")
	}

	initCode, err := packInitCode(cfg, feeOracle, impl)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "pack init code")
	}

	tx, err = factory.Deploy(txOpts, salt, initCode)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "deploy proxy")
	}

	receipt, err = bind.WaitMined(ctx, backend, tx)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "wait mined proxy")
	} else if receipt.Status != ethtypes.ReceiptStatusSuccessful {
		return common.Address{}, nil, errors.New("deploy proxy failed")
	}

	return addr, receipt, nil
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

	initializer, err := portalAbi.Pack("initialize", cfg.Owner, feeOracle,
		cfg.OmniEChainID, cfg.OmniCChainID, cfg.ValSetID, cfg.Validators)
	if err != nil {
		return nil, errors.Wrap(err, "encode portal initializer")
	}

	return contracts.PackInitCode(proxyAbi, bindings.TransparentUpgradeableProxyBin, impl, cfg.ProxyAdmin, initializer)
}
