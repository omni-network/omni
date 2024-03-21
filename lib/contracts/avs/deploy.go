package avs

import (
	"context"
	"math/big"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/anvil"
	"github.com/omni-network/omni/lib/chainids"
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

//nolint:gochecknoglobals // static abi type
var avsABI = mustGetABI(bindings.OmniAVSMetaData)

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

	return nil
}

func getDeployCfg(chainID uint64) (DeploymentConfig, error) {
	if chainID == chainids.Holesky {
		return holeskeyDeployCfg(), nil
	}

	return DeploymentConfig{}, errors.New("unsupported chain")
}

func holeskeyDeployCfg() DeploymentConfig {
	return DeploymentConfig{
		Create3Salt:      "holesky-avs",
		Eigen:            holeskyEigenDeployments(),
		StrategyParams:   holeskeyStrategyParams(),
		MetadataURI:      metadataURI,
		MinOperatorStake: big.NewInt(1e18), // 1 ETH
		MaxOperatorCount: 200,
		AllowlistEnabled: false,
		// TODO: fill in the rest
	}
}

func devnetDeployCfg() DeploymentConfig {
	return DeploymentConfig{
		Create3Factory:   common.HexToAddress("0x1234"), // TODO: currently unused
		Create3Salt:      "devnet-avs",                  // TODO: currently unused
		Deployer:         anvil.Account0,
		Owner:            anvil.Account0,
		Eigen:            devnetEigenDeployments,
		ProxyAdmin:       anvil.Account1, // should not be an eoa, but does not matter for devnet (yet)
		MetadataURI:      "https://test.com",
		OmniChainID:      netconf.GetStatic("Devnet").OmniExecutionChainID,
		AllowlistEnabled: true,
		StrategyParams:   devnetStrategyParams(),
		EthStakeInbox:    common.HexToAddress("0x1234"), // TODO: replace with actual address
		MinOperatorStake: big.NewInt(1e18),              // 1 ETH
		MaxOperatorCount: 10,
	}
}

// Deploy deploys the AVS contract and returns the address and receipt.
// It only allows deployments to explicitly supported chains.
func Deploy(ctx context.Context, backend *ethbackend.Backend) (common.Address, *ethtypes.Receipt, error) {
	chainID, err := backend.ChainID(ctx)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "chain id")
	}

	cfg, err := getDeployCfg(chainID.Uint64())
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "get deployment config")
	}

	return deploy(ctx, cfg, backend)
}

// DeployDevnet deploys the devnet AVS contract and returns the address receipt.
func DeployDevnet(ctx context.Context, backend *ethbackend.Backend) (common.Address, *ethtypes.Receipt, error) {
	chainID, err := backend.ChainID(ctx)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "chain id")
	}

	if chainids.IsMainnetOrTestnet(chainID.Uint64()) {
		return common.Address{}, nil, errors.New("not a devnet")
	}

	return deploy(ctx, devnetDeployCfg(), backend)
}

// Deploy deploys the AVS contract and returns the deployed contract's address and the transaction receipt.
func deploy(ctx context.Context, cfg DeploymentConfig, backend *ethbackend.Backend) (common.Address, *ethtypes.Receipt, error) {
	if err := cfg.Validate(); err != nil {
		return common.Address{}, nil, errors.Wrap(err, "validate config")
	}

	deployerTxOpts, err := backend.BindOpts(ctx, cfg.Deployer)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "bind deployer opts")
	}

	ownerTxOpts, err := backend.BindOpts(ctx, cfg.Owner)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "bind owner opts")
	}

	enc, err := packInitialzer(cfg)
	if err != nil {
		return common.Address{}, nil, err
	}

	impl, tx, _, err := bindings.DeployOmniAVS(deployerTxOpts, backend, cfg.Eigen.DelegationManager, cfg.Eigen.AVSDirectory)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "deploy avs impl")
	}
	_, err = backend.WaitMined(ctx, tx)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "wait mined avs proxy")
	}

	proxy, tx, _, err := bindings.DeployTransparentUpgradeableProxy(deployerTxOpts, backend, impl, cfg.ProxyAdmin, enc)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "deploy avs proxy")
	}
	receipt, err := backend.WaitMined(ctx, tx)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "wait mined avs proxy")
	}

	avs, err := bindings.NewOmniAVS(proxy, backend)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "bind avs")
	}

	if !cfg.AllowlistEnabled {
		tx, err = avs.DisableAllowlist(ownerTxOpts)
		if err != nil {
			return common.Address{}, nil, errors.Wrap(err, "disable allowlist")
		}

		_, err = backend.WaitMined(ctx, tx)
		if err != nil {
			return common.Address{}, nil, errors.Wrap(err, "wait mined disable allowlist")
		}
	}

	tx, err = avs.SetMetadataURI(ownerTxOpts, cfg.MetadataURI)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "set metadata uri")
	}

	_, err = backend.WaitMined(ctx, tx)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "wait mined set metadata uri")
	}

	return proxy, receipt, nil
}

// packInitializer encodes the initializer parameters for the AVS contract.
func packInitialzer(cfg DeploymentConfig) ([]byte, error) {
	enc, err := avsABI.Pack("initialize",
		cfg.Owner, cfg.Portal, cfg.OmniChainID, cfg.EthStakeInbox,
		cfg.MinOperatorStake, cfg.MaxOperatorCount, strategyParams(cfg))

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
