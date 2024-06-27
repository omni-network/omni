package portal

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
	Create3Factory        common.Address
	Create3Salt           string
	ProxyAdmin            common.Address
	Deployer              common.Address
	Owner                 common.Address
	OmniChainID           uint64
	OmniCChainID          uint64
	XMsgMinGasLimit       uint64
	XMsgMaxGasLimit       uint64
	XMsgMaxDataSize       uint16
	XReceiptMaxErrorBytes uint16
	CChainXMsgOffset      uint64
	CChainXBlockOffset    uint64
	ExpectedAddr          common.Address
}

const (
	XMsgMinGasLimit       = 21_000
	XMsgMaxGasLimit       = 5_000_000
	XMsgMaxDataSize       = 20_000
	XReceiptMaxErrorBytes = 256

	// We use default cchain xmsg and xblock offsets of 1. This xmsg / block
	// contains the genesis validator set. This is set in the portal contract
	// on initialization, and therefore does not need to be relayed.
	//
	// Setting xmsg and xblock offsets to 1 works well for portals that are
	// deployed at "network genesis" - the start of an omni network. For portals
	// added to an existing network, we should use higher offsets on initialization,
	// so that the entire history of cchain xmsgs do not need to be relayed.

	GenesisCChainXMsgOffset   = 1
	GenesisCChainXBlockOffset = 1
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
	if cfg.XMsgMaxDataSize == 0 {
		return errors.New("xmsg max data size is zero")
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
	if cfg.CChainXMsgOffset == 0 {
		return errors.New("cchain xmsg offset is zero")
	}
	if cfg.CChainXBlockOffset == 0 {
		return errors.New("cchain xblock offset is zero")
	}
	if (cfg.ExpectedAddr == common.Address{}) {
		return errors.New("expected address is zero")
	}

	return nil
}

func getDeployCfg(network netconf.ID) (DeploymentConfig, error) {
	if network == netconf.Devnet {
		return devnetCfg(), nil
	}

	if network == netconf.Mainnet {
		return mainnetCfg(), nil
	}

	if network == netconf.Omega {
		return omegaCfg(), nil
	}

	if network == netconf.Staging {
		return stagingCfg(), nil
	}

	return DeploymentConfig{}, errors.New("unsupported network", "network", network)
}

func mainnetCfg() DeploymentConfig {
	return DeploymentConfig{
		Create3Factory: contracts.MainnetCreate3Factory(),
		Create3Salt:    contracts.PortalSalt(netconf.Mainnet),
		Owner:          eoa.MustAddress(netconf.Mainnet, eoa.RoleAdmin),
		Deployer:       eoa.MustAddress(netconf.Mainnet, eoa.RoleDeployer),
		// TODO: fill in the rest
	}
}

func omegaCfg() DeploymentConfig {
	return DeploymentConfig{
		Create3Factory:        contracts.OmegaCreate3Factory(),
		Create3Salt:           contracts.PortalSalt(netconf.Omega),
		Owner:                 eoa.MustAddress(netconf.Omega, eoa.RoleAdmin),
		Deployer:              eoa.MustAddress(netconf.Omega, eoa.RoleDeployer),
		ProxyAdmin:            contracts.OmegaProxyAdmin(),
		OmniChainID:           netconf.Omega.Static().OmniExecutionChainID,
		OmniCChainID:          netconf.Omega.Static().OmniConsensusChainIDUint64(),
		XMsgMinGasLimit:       XMsgMinGasLimit,
		XMsgMaxGasLimit:       XMsgMaxGasLimit,
		XMsgMaxDataSize:       XMsgMaxDataSize,
		CChainXMsgOffset:      GenesisCChainXMsgOffset,
		CChainXBlockOffset:    GenesisCChainXBlockOffset,
		XReceiptMaxErrorBytes: XReceiptMaxErrorBytes,
		ExpectedAddr:          contracts.OmegaPortal(),
	}
}

func stagingCfg() DeploymentConfig {
	return DeploymentConfig{
		Create3Factory:        contracts.StagingCreate3Factory(),
		Create3Salt:           contracts.PortalSalt(netconf.Staging),
		Owner:                 eoa.MustAddress(netconf.Staging, eoa.RoleAdmin),
		Deployer:              eoa.MustAddress(netconf.Staging, eoa.RoleDeployer),
		ProxyAdmin:            contracts.StagingProxyAdmin(),
		OmniChainID:           netconf.Staging.Static().OmniExecutionChainID,
		OmniCChainID:          netconf.Staging.Static().OmniConsensusChainIDUint64(),
		XMsgMinGasLimit:       XMsgMinGasLimit,
		XMsgMaxGasLimit:       XMsgMaxGasLimit,
		XMsgMaxDataSize:       XMsgMaxDataSize,
		CChainXMsgOffset:      GenesisCChainXMsgOffset,
		CChainXBlockOffset:    GenesisCChainXBlockOffset,
		XReceiptMaxErrorBytes: XReceiptMaxErrorBytes,
		ExpectedAddr:          contracts.StagingPortal(),
	}
}

func devnetCfg() DeploymentConfig {
	return DeploymentConfig{
		Create3Factory:        contracts.DevnetCreate3Factory(),
		Create3Salt:           contracts.PortalSalt(netconf.Devnet),
		Owner:                 eoa.MustAddress(netconf.Devnet, eoa.RoleAdmin),
		Deployer:              eoa.MustAddress(netconf.Devnet, eoa.RoleDeployer),
		ProxyAdmin:            contracts.DevnetProxyAdmin(),
		OmniChainID:           netconf.Devnet.Static().OmniExecutionChainID,
		OmniCChainID:          netconf.Devnet.Static().OmniConsensusChainIDUint64(),
		XMsgMinGasLimit:       XMsgMinGasLimit,
		XMsgMaxGasLimit:       XMsgMaxGasLimit,
		XMsgMaxDataSize:       XMsgMaxDataSize,
		CChainXMsgOffset:      GenesisCChainXMsgOffset,
		CChainXBlockOffset:    GenesisCChainXBlockOffset,
		XReceiptMaxErrorBytes: XReceiptMaxErrorBytes,
		ExpectedAddr:          contracts.DevnetPortal(),
	}
}

func AddrForNetwork(network netconf.ID) (common.Address, bool) {
	switch network {
	case netconf.Mainnet:
		return contracts.MainnetPortal(), true
	case netconf.Omega:
		return contracts.OmegaPortal(), true
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
	cfg, err := getDeployCfg(network)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "get deployment config")
	}

	return deploy(ctx, cfg, backend, feeOracle, valSetID, validators)
}

func deploy(ctx context.Context, cfg DeploymentConfig, backend *ethbackend.Backend, feeOracle common.Address, valSetID uint64, validators []bindings.Validator,
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

	impl, tx, _, err := bindings.DeployOmniPortal(txOpts, backend)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "deploy impl")
	}

	_, err = backend.WaitMined(ctx, tx)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "wait mined portal")
	}

	initCode, err := packInitCode(cfg, feeOracle, impl, valSetID, validators)
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

func packInitCode(cfg DeploymentConfig, feeOracle common.Address, impl common.Address, valSetID uint64, validators []bindings.Validator,
) ([]byte, error) {
	portalAbi, err := bindings.OmniPortalMetaData.GetAbi()
	if err != nil {
		return nil, errors.Wrap(err, "get portal abi")
	}

	proxyAbi, err := bindings.TransparentUpgradeableProxyMetaData.GetAbi()
	if err != nil {
		return nil, errors.Wrap(err, "get proxy abi")
	}

	initializer, err := portalAbi.Pack("initialize",
		&bindings.OmniPortalInitParams{
			Owner:                cfg.Owner,
			FeeOracle:            feeOracle,
			OmniChainId:          cfg.OmniChainID,
			OmniCChainId:         cfg.OmniCChainID,
			XmsgMaxGasLimit:      cfg.XMsgMaxGasLimit,
			XmsgMinGasLimit:      cfg.XMsgMinGasLimit,
			XmsgMaxDataSize:      cfg.XMsgMaxDataSize,
			XreceiptMaxErrorSize: cfg.XReceiptMaxErrorBytes,
			CChainXMsgOffset:     cfg.CChainXMsgOffset,
			CChainXBlockOffset:   cfg.CChainXBlockOffset,
			ValSetId:             valSetID,
			Validators:           validators,
		},
	)
	if err != nil {
		return nil, errors.Wrap(err, "encode portal initializer")
	}

	return contracts.PackInitCode(proxyAbi, bindings.TransparentUpgradeableProxyBin, impl, cfg.ProxyAdmin, initializer)
}
