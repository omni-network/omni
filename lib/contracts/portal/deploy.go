package portal

import (
	"context"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/halo/genutil/evm/predeploys"
	"github.com/omni-network/omni/lib/contracts"
	"github.com/omni-network/omni/lib/create3"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

type DeploymentConfig struct {
	Create3Factory        common.Address
	Create3Salt           string
	ProxyAdmin            common.Address
	Deployer              common.Address
	Owner                 common.Address
	OmniChainID           uint64
	OmniCChainID          uint64
	XMsgMinGasLimit       uint64
	XMsgMaxGasLimit       uint64
	XReceiptMaxErrorBytes uint16
	ExpectedAddr          common.Address
}

const (
	XMsgMinGasLimit       = 21_000
	XMsgMaxGasLimit       = 5_000_000
	XReceiptMaxErrorBytes = uint16(256)
)

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
	if (cfg.ProxyAdmin == common.Address{}) {
		return errors.New("proxy admin is zero")
	}
	if isDeadOrEmpty(cfg.Deployer) {
		return errors.New("deployer is not set")
	}
	if isDeadOrEmpty(cfg.Owner) {
		return errors.New("owner is not set")
	}
	if cfg.XMsgMinGasLimit == 0 {
		return errors.New("xmsg min gas limit is zero")
	}
	if cfg.XMsgMaxGasLimit == 0 {
		return errors.New("xmsg max gas limit is zero")
	}
	if cfg.XReceiptMaxErrorBytes == 0 {
		return errors.New("xreceipt max error bytes is zero")
	}
	if cfg.OmniChainID == 0 {
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

func getDeployCfg(chainID uint64, network netconf.ID) (DeploymentConfig, error) {
	if network == netconf.Devnet {
		return devnetCfg(), nil
	}

	if network == netconf.Mainnet {
		return mainnetCfg(), nil
	}

	if network == netconf.Omega {
		return testnetCfg(), nil
	}

	if network == netconf.Staging {
		return stagingCfg(), nil
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

func testnetCfg() DeploymentConfig {
	return DeploymentConfig{
		Create3Factory:        contracts.TestnetCreate3Factory(),
		Create3Salt:           contracts.PortalSalt(netconf.Omega),
		Owner:                 eoa.MustAddress(netconf.Omega, eoa.RolePortalAdmin),
		Deployer:              eoa.MustAddress(netconf.Omega, eoa.RoleDeployer),
		ProxyAdmin:            contracts.TestnetProxyAdmin(),
		OmniChainID:           netconf.Omega.Static().OmniExecutionChainID,
		OmniCChainID:          netconf.Omega.Static().OmniConsensusChainIDUint64(),
		XMsgMinGasLimit:       XMsgMinGasLimit,
		XMsgMaxGasLimit:       XMsgMaxGasLimit,
		XReceiptMaxErrorBytes: XReceiptMaxErrorBytes,
		ExpectedAddr:          contracts.TestnetPortal(),
	}
}

func stagingCfg() DeploymentConfig {
	return DeploymentConfig{
		Create3Factory:        contracts.StagingCreate3Factory(),
		Create3Salt:           contracts.PortalSalt(netconf.Staging),
		Owner:                 eoa.MustAddress(netconf.Staging, eoa.RolePortalAdmin),
		Deployer:              eoa.MustAddress(netconf.Staging, eoa.RoleDeployer),
		ProxyAdmin:            contracts.StagingProxyAdmin(),
		OmniChainID:           netconf.Staging.Static().OmniExecutionChainID,
		OmniCChainID:          netconf.Staging.Static().OmniConsensusChainIDUint64(),
		XMsgMinGasLimit:       XMsgMinGasLimit,
		XMsgMaxGasLimit:       XMsgMaxGasLimit,
		XReceiptMaxErrorBytes: XReceiptMaxErrorBytes,
		ExpectedAddr:          contracts.StagingPortal(),
	}
}

func devnetCfg() DeploymentConfig {
	return DeploymentConfig{
		Create3Factory:        contracts.DevnetCreate3Factory(),
		Create3Salt:           contracts.PortalSalt(netconf.Devnet),
		Owner:                 eoa.MustAddress(netconf.Devnet, eoa.RolePortalAdmin),
		Deployer:              eoa.MustAddress(netconf.Devnet, eoa.RoleDeployer),
		ProxyAdmin:            contracts.DevnetProxyAdmin(),
		OmniChainID:           netconf.Devnet.Static().OmniExecutionChainID,
		OmniCChainID:          netconf.Devnet.Static().OmniConsensusChainIDUint64(),
		XMsgMinGasLimit:       XMsgMinGasLimit,
		XMsgMaxGasLimit:       XMsgMaxGasLimit,
		XReceiptMaxErrorBytes: XReceiptMaxErrorBytes,
		ExpectedAddr:          contracts.DevnetPortal(),
	}
}

func AddrForNetwork(network netconf.ID) (common.Address, bool) {
	switch network {
	case netconf.Mainnet:
		return contracts.MainnetPortal(), true
	case netconf.Omega:
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
func Deploy(ctx context.Context, network netconf.ID, backend *ethbackend.Backend, feeOracle common.Address, valSetID uint64, validators []bindings.Validator,
) (common.Address, *ethtypes.Receipt, error) {
	chainID, err := backend.ChainID(ctx)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "chain id")
	}

	cfg, err := getDeployCfg(chainID.Uint64(), network)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "get deployment config")
	}

	return deploy(ctx, cfg, chainID.Uint64(), backend, feeOracle, valSetID, validators)
}

func deploy(ctx context.Context, cfg DeploymentConfig, chainID uint64, backend *ethbackend.Backend, feeOracle common.Address, valSetID uint64, validators []bindings.Validator,
) (common.Address, *ethtypes.Receipt, error) {
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

	var xregistry common.Address
	if chainID == cfg.OmniChainID {
		// On Omni, the main XRegistry is the "replica"
		xregistry = common.HexToAddress(predeploys.XRegistry)
	} else {
		// On other chains, deploy a new XRegistry replica
		addr, tx, _, err := bindings.DeployXRegistryReplica(txOpts, backend, addr)
		if err != nil {
			return common.Address{}, nil, errors.Wrap(err, "new xregistry replica")
		}
		xregistry = addr

		_, err = backend.WaitMined(ctx, tx)
		if err != nil {
			return common.Address{}, nil, errors.Wrap(err, "wait mined xregistry replica")
		}
	}

	impl, tx, _, err := bindings.DeployOmniPortal(txOpts, backend)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "deploy impl")
	}

	_, err = backend.WaitMined(ctx, tx)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "wait mined portal")
	}

	initCode, err := packInitCode(cfg, feeOracle, xregistry, impl, valSetID, validators)
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

func packInitCode(cfg DeploymentConfig, feeOracle common.Address, xregistry common.Address, impl common.Address, valSetID uint64, validators []bindings.Validator,
) ([]byte, error) {
	portalAbi, err := bindings.OmniPortalMetaData.GetAbi()
	if err != nil {
		return nil, errors.Wrap(err, "get portal abi")
	}

	proxyAbi, err := bindings.TransparentUpgradeableProxyMetaData.GetAbi()
	if err != nil {
		return nil, errors.Wrap(err, "get proxy abi")
	}

	initializer, err := portalAbi.Pack("initialize", cfg.Owner, feeOracle, xregistry, cfg.OmniChainID, cfg.OmniCChainID,
		cfg.XMsgMaxGasLimit, cfg.XMsgMinGasLimit, cfg.XReceiptMaxErrorBytes, valSetID, validators)
	if err != nil {
		return nil, errors.Wrap(err, "encode portal initializer")
	}

	return contracts.PackInitCode(proxyAbi, bindings.TransparentUpgradeableProxyBin, impl, cfg.ProxyAdmin, initializer)
}
