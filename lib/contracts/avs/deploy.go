package avs

import (
	"context"
	"math/big"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/lib/chainids"
	"github.com/omni-network/omni/lib/contracts"
	"github.com/omni-network/omni/lib/create3"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

const (
	metadataURI = "https://raw.githubusercontent.com/omni-network/omni/main/static/avs-metadata.json"
)

//nolint:gochecknoglobals // static abi type & addresses
var (
	avsABI            = mustGetABI(bindings.OmniAVSMetaData)
	proxyABI          = mustGetABI(bindings.TransparentUpgradeableProxyMetaData)
	stubEthStakeInbox = common.HexToAddress("0x1234") // TODO: replace with actual address halo/genutil/evm/predeploys.EthStakeInbox
)

type DeploymentConfig struct {
	Create3Factory   common.Address
	Create3Salt      string
	Eigen            EigenDeployments
	Deployer         common.Address
	Owner            common.Address
	ProxyAdmin       common.Address
	Portal           common.Address
	OmniChainID      uint64
	MetadataURI      string
	StrategyParams   []StrategyParam
	EthStakeInbox    common.Address
	MinOperatorStake *big.Int
	MaxOperatorCount uint32
	AllowlistEnabled bool
	ExpectedAddr     common.Address
}

func (cfg DeploymentConfig) Validate() error {
	if (cfg.Create3Factory == common.Address{}) {
		return errors.New("create3 factory not set")
	}
	if cfg.Create3Salt == "" {
		return errors.New("create3 salt not set")
	}
	if err := cfg.Eigen.Validate(); err != nil {
		return errors.Wrap(err, "eigen deployments")
	}
	if (cfg.Deployer == common.Address{}) {
		return errors.New("deployer is zero")
	}
	if (cfg.Owner == common.Address{}) {
		return errors.New("owner is zero")
	}
	if (cfg.ProxyAdmin == common.Address{}) {
		return errors.New("proxy admin is zero")
	}
	if cfg.MetadataURI == "" {
		return errors.New("metadata uri not set")
	}
	if cfg.MinOperatorStake == nil {
		return errors.New("min operator stake not set")
	}
	if cfg.MaxOperatorCount == 0 {
		return errors.New("max operator count not set")
	}
	if (cfg.ExpectedAddr == common.Address{}) {
		return errors.New("expected address is zero")
	}
	if (cfg.Portal == common.Address{}) {
		return errors.New("portal is zero")
	}
	if (cfg.EthStakeInbox == common.Address{}) {
		return errors.New("eth stake inbox is zero")
	}

	return nil
}

func getDeployCfg(chainID uint64, network netconf.ID) (DeploymentConfig, error) {
	if !chainids.IsMainnetOrTestnet(chainID) && network == netconf.Devnet {
		return devnetCfg(), nil
	}

	if chainID == chainids.Holesky && network == netconf.Testnet {
		return testnetCfg(), nil
	}

	if !chainids.IsMainnet(chainID) && network == netconf.Staging {
		return stagingCfg(), nil
	}

	return DeploymentConfig{}, errors.New("unsupported chain for network", "chain_id", chainID, "network", network)
}

func testnetCfg() DeploymentConfig {
	return DeploymentConfig{
		Create3Factory:   contracts.TestnetCreate3Factory(),
		Create3Salt:      contracts.AVSSalt(netconf.Testnet),
		Deployer:         eoa.MustAddress(netconf.Testnet, eoa.RoleDeployer),
		Owner:            eoa.MustAddress(netconf.Testnet, eoa.RoleAVSAdmin),
		ProxyAdmin:       contracts.TestnetProxyAdmin(),
		Eigen:            holeskyEigenDeployments(),
		StrategyParams:   holeskyStrategyParams(),
		MetadataURI:      metadataURI,
		OmniChainID:      netconf.Testnet.Static().OmniExecutionChainID,
		Portal:           contracts.TestnetPortal(),
		EthStakeInbox:    stubEthStakeInbox,
		MinOperatorStake: big.NewInt(1e18), // 1 ETH
		MaxOperatorCount: 200,
		AllowlistEnabled: false,
		ExpectedAddr:     contracts.TestnetAVS(),
	}
}

func stagingCfg() DeploymentConfig {
	return DeploymentConfig{
		Create3Factory:   contracts.StagingCreate3Factory(),
		Create3Salt:      contracts.AVSSalt(netconf.Staging),
		Deployer:         eoa.MustAddress(netconf.Staging, eoa.RoleDeployer),
		Owner:            eoa.MustAddress(netconf.Staging, eoa.RoleAVSAdmin),
		ProxyAdmin:       contracts.StagingProxyAdmin(),
		Eigen:            devnetEigenDeployments,
		StrategyParams:   devnetStrategyParams(),
		MetadataURI:      metadataURI,
		OmniChainID:      netconf.Staging.Static().OmniExecutionChainID,
		Portal:           contracts.StagingPortal(),
		EthStakeInbox:    stubEthStakeInbox,
		MinOperatorStake: big.NewInt(1e18), // 1 ETH
		MaxOperatorCount: 10,
		AllowlistEnabled: true,
		ExpectedAddr:     contracts.StagingAVS(),
	}
}

func devnetCfg() DeploymentConfig {
	return DeploymentConfig{
		Create3Factory:   contracts.DevnetCreate3Factory(),
		Create3Salt:      contracts.AVSSalt(netconf.Devnet),
		Deployer:         eoa.MustAddress(netconf.Devnet, eoa.RoleDeployer),
		Owner:            eoa.MustAddress(netconf.Devnet, eoa.RoleAVSAdmin),
		ProxyAdmin:       contracts.DevnetProxyAdmin(),
		Eigen:            devnetEigenDeployments,
		MetadataURI:      metadataURI,
		OmniChainID:      netconf.Devnet.Static().OmniExecutionChainID,
		StrategyParams:   devnetStrategyParams(),
		Portal:           contracts.DevnetPortal(),
		EthStakeInbox:    stubEthStakeInbox,
		MinOperatorStake: big.NewInt(1e18), // 1 ETH
		MaxOperatorCount: 10,
		AllowlistEnabled: true,
		ExpectedAddr:     contracts.DevnetAVS(),
	}
}

func AddrForNetwork(network netconf.ID) (common.Address, bool) {
	switch network {
	case netconf.Mainnet:
		return contracts.MainnetAVS(), true
	case netconf.Testnet:
		return contracts.TestnetAVS(), true
	case netconf.Staging:
		return contracts.StagingAVS(), true
	case netconf.Devnet:
		return contracts.DevnetAVS(), true
	default:
		return common.Address{}, false
	}
}

// IsDeployed checks if the OmniAVS contract is deployed to the provided backend
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

// DeployIfNeeded deploys a new AVS contract if it is not already deployed.
func DeployIfNeeded(ctx context.Context, network netconf.ID, backend *ethbackend.Backend) (common.Address, *ethtypes.Receipt, error) {
	deployed, addr, err := IsDeployed(ctx, network, backend)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "is deployed")
	}
	if deployed {
		return addr, nil, nil
	}

	return Deploy(ctx, network, backend)
}

// Deploy deploys the AVS contract and returns the address and receipt.
// It only allows deployments to explicitly supported chains.
func Deploy(ctx context.Context, network netconf.ID, backend *ethbackend.Backend) (common.Address, *ethtypes.Receipt, error) {
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

func deploy(ctx context.Context, cfg DeploymentConfig, backend *ethbackend.Backend) (common.Address, *ethtypes.Receipt, error) {
	if err := cfg.Validate(); err != nil {
		return common.Address{}, nil, errors.Wrap(err, "validate config")
	}

	txOpts, err := backend.BindOpts(ctx, cfg.Deployer)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "bind deployer opts")
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

	impl, tx, _, err := bindings.DeployOmniAVS(txOpts, backend, cfg.Eigen.DelegationManager, cfg.Eigen.AVSDirectory)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "deploy impl")
	}

	receipt, err := backend.WaitMined(ctx, tx)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "wait mined portal")
	} else if receipt.Status != ethtypes.ReceiptStatusSuccessful {
		return common.Address{}, nil, errors.New("deploy impl failed")
	}

	initCode, err := packInitCode(cfg, impl)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "pack init code")
	}

	tx, err = factory.Deploy(txOpts, salt, initCode)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "deploy proxy")
	}

	receipt, err = backend.WaitMined(ctx, tx)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "wait mined proxy")
	} else if receipt.Status != ethtypes.ReceiptStatusSuccessful {
		return common.Address{}, nil, errors.New("deploy proxy failed")
	}

	return addr, receipt, nil
}

func packInitCode(cfg DeploymentConfig, impl common.Address) ([]byte, error) {
	initializer, err := packInitialzer(cfg)
	if err != nil {
		return nil, err
	}

	return contracts.PackInitCode(proxyABI, bindings.TransparentUpgradeableProxyBin, impl, cfg.ProxyAdmin, initializer)
}

// packInitializer encodes the initializer parameters for the AVS contract.
func packInitialzer(cfg DeploymentConfig) ([]byte, error) {
	enc, err := avsABI.Pack("initialize",
		cfg.Owner, cfg.Portal, cfg.OmniChainID, cfg.EthStakeInbox,
		cfg.MinOperatorStake, cfg.MaxOperatorCount, strategyParams(cfg),
		cfg.MetadataURI, cfg.AllowlistEnabled)

	if err != nil {
		return nil, errors.Wrap(err, "pack initializer")
	}

	return enc, nil
}

// strategyParams converts the deployment config's strategy params to the.
func strategyParams(cfg DeploymentConfig) []bindings.IOmniAVSStrategyParam {
	params := make([]bindings.IOmniAVSStrategyParam, len(cfg.StrategyParams))
	for i, sp := range cfg.StrategyParams {
		params[i] = bindings.IOmniAVSStrategyParam{
			Strategy:   sp.Strategy,
			Multiplier: sp.Multiplier,
		}
	}

	return params
}

// mustGetABI returns the metadata's ABI as an abi.ABI type.
// It panics on error.
func mustGetABI(metadata *bind.MetaData) *abi.ABI {
	abi, err := metadata.GetAbi()
	if err != nil {
		panic(err)
	}

	return abi
}
