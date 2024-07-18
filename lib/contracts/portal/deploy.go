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

type deploymentConfig struct {
	Create3Factory        common.Address
	Create3Salt           string
	ProxyAdminOwner       common.Address
	Deployer              common.Address
	Owner                 common.Address
	OmniChainID           uint64
	OmniCChainID          uint64
	XMsgMinGasLimit       uint64
	XMsgMaxGasLimit       uint64
	XMsgMaxDataSize       uint16
	XReceiptMaxErrorBytes uint16
	XSubValsetCutoff      uint8
	CChainXMsgOffset      uint64
	CChainXBlockOffset    uint64
	ExpectedAddr          common.Address
}

const (
	XMsgMinGasLimit       = 21_000
	XMsgMaxGasLimit       = 5_000_000
	XMsgMaxDataSize       = 20_000
	XReceiptMaxErrorBytes = 256
	XSubValsetCutoff      = 10

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

func (cfg deploymentConfig) Validate() error {
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

// IsDeployed checks if the Portal contract is deployed to the provided backend
// to its expected network address.
func IsDeployed(ctx context.Context, network netconf.ID, backend *ethbackend.Backend) (bool, common.Address, error) {
	addr := contracts.Portal(network)

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
	cfg := deploymentConfig{
		Create3Factory:        contracts.Create3Factory(network),
		Create3Salt:           contracts.PortalSalt(network),
		Owner:                 eoa.MustAddress(network, eoa.RoleAdmin),
		Deployer:              eoa.MustAddress(network, eoa.RoleDeployer),
		ProxyAdminOwner:       eoa.MustAddress(network, eoa.RoleAdmin),
		OmniChainID:           network.Static().OmniExecutionChainID,
		OmniCChainID:          network.Static().OmniConsensusChainIDUint64(),
		XMsgMinGasLimit:       XMsgMinGasLimit,
		XMsgMaxGasLimit:       XMsgMaxGasLimit,
		XMsgMaxDataSize:       XMsgMaxDataSize,
		XReceiptMaxErrorBytes: XReceiptMaxErrorBytes,
		XSubValsetCutoff:      XSubValsetCutoff,
		CChainXMsgOffset:      GenesisCChainXMsgOffset,
		CChainXBlockOffset:    GenesisCChainXBlockOffset,
		ExpectedAddr:          contracts.Portal(network),
	}

	return deploy(ctx, cfg, backend, feeOracle, valSetID, validators)
}

func deploy(ctx context.Context, cfg deploymentConfig, backend *ethbackend.Backend, feeOracle common.Address, valSetID uint64, validators []bindings.Validator,
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

func packInitCode(cfg deploymentConfig, feeOracle common.Address, impl common.Address, valSetID uint64, validators []bindings.Validator,
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
			XsubValsetCutoff:     cfg.XSubValsetCutoff,
			CChainXMsgOffset:     cfg.CChainXMsgOffset,
			CChainXBlockOffset:   cfg.CChainXBlockOffset,
			ValSetId:             valSetID,
			Validators:           validators,
		},
	)
	if err != nil {
		return nil, errors.Wrap(err, "encode portal initializer")
	}

	return contracts.PackInitCode(proxyAbi, bindings.TransparentUpgradeableProxyBin, impl, cfg.ProxyAdminOwner, initializer)
}
