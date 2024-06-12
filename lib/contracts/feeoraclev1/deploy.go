package feeoraclev1

import (
	"context"
	"math/big"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/lib/contracts"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/tokens/coingecko"

	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

type DeploymentConfig struct {
	Owner        common.Address
	Manager      common.Address // manager is the address that can set fee parameters (gas price, conversion rates)
	Deployer     common.Address
	ProxyAdmin   common.Address
	BaseGasLimit uint64
	ProtocolFee  *big.Int
}

func isDeadOrEmpty(addr common.Address) bool {
	return addr == common.Address{} || addr == common.HexToAddress(eoa.ZeroXDead)
}

func (cfg DeploymentConfig) Validate() error {
	if isDeadOrEmpty(cfg.Owner) {
		return errors.New("owner is zero")
	}
	if isDeadOrEmpty(cfg.Manager) {
		return errors.New("manager is zero")
	}
	if isDeadOrEmpty(cfg.Deployer) {
		return errors.New("deployer is zero")
	}
	if (cfg.ProxyAdmin == common.Address{}) {
		return errors.New("proxy admin is zero")
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

// NOTE: monitor is owner of fee oracle contracts, because monitor manages on chain gas prices / conversion rates

func mainnetCfg() DeploymentConfig {
	return DeploymentConfig{
		Owner:        eoa.MustAddress(netconf.Mainnet, eoa.RolePortalAdmin),
		Manager:      eoa.MustAddress(netconf.Mainnet, eoa.RoleMonitor),
		Deployer:     eoa.MustAddress(netconf.Mainnet, eoa.RoleDeployer),
		ProxyAdmin:   contracts.MainnetProxyAdmin(),
		BaseGasLimit: 50_000,
		ProtocolFee:  big.NewInt(0),
	}
}

func testnetCfg() DeploymentConfig {
	return DeploymentConfig{
		Owner:        eoa.MustAddress(netconf.Omega, eoa.RolePortalAdmin),
		Manager:      eoa.MustAddress(netconf.Omega, eoa.RoleMonitor),
		Deployer:     eoa.MustAddress(netconf.Omega, eoa.RoleDeployer),
		ProxyAdmin:   contracts.TestnetProxyAdmin(),
		BaseGasLimit: 50_000,
		ProtocolFee:  big.NewInt(0),
	}
}

func devnetCfg() DeploymentConfig {
	return DeploymentConfig{
		Owner:        eoa.MustAddress(netconf.Devnet, eoa.RolePortalAdmin),
		Manager:      eoa.MustAddress(netconf.Devnet, eoa.RoleMonitor),
		Deployer:     eoa.MustAddress(netconf.Devnet, eoa.RoleDeployer),
		ProxyAdmin:   contracts.DevnetProxyAdmin(),
		BaseGasLimit: 50_000,
		ProtocolFee:  big.NewInt(0),
	}
}

func stagingCfg() DeploymentConfig {
	return DeploymentConfig{
		Owner:        eoa.MustAddress(netconf.Staging, eoa.RolePortalAdmin),
		Manager:      eoa.MustAddress(netconf.Staging, eoa.RoleMonitor),
		Deployer:     eoa.MustAddress(netconf.Staging, eoa.RoleDeployer),
		ProxyAdmin:   contracts.StagingProxyAdmin(),
		BaseGasLimit: 50_000,
		ProtocolFee:  big.NewInt(0),
	}
}

func Deploy(ctx context.Context, network netconf.ID, chainID uint64, destChainIDs []uint64, backends ethbackend.Backends) (common.Address, *ethtypes.Receipt, error) {
	cfg, err := getDeployCfg(chainID, network)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "get deployment config")
	}

	backend, err := backends.Backend(chainID)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "get backend")
	}

	txOpts, err := backend.BindOpts(ctx, cfg.Deployer)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "bind opts")
	}

	feeparams, err := feeParams(ctx, chainID, destChainIDs, backends, coingecko.New())
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "fee params")
	}

	feeOracleAbi, err := bindings.FeeOracleV1MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "get fee oracle abi")
	}

	initializer, err := feeOracleAbi.Pack("initialize", cfg.Owner, cfg.Manager, cfg.BaseGasLimit, cfg.ProtocolFee, feeparams)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "pack initialize")
	}

	impl, tx, _, err := bindings.DeployFeeOracleV1(txOpts, backend)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "deploy fee oracle")
	}

	_, err = backend.WaitMined(ctx, tx)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "wait mined")
	}

	proxy, tx, _, err := bindings.DeployTransparentUpgradeableProxy(txOpts, backend, impl, cfg.ProxyAdmin, initializer)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "deploy proxy")
	}

	receipt, err := backend.WaitMined(ctx, tx)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "wait mined")
	}

	return proxy, receipt, nil
}
