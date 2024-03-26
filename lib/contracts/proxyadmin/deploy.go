package proxyadmin

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
	Owner          common.Address
	Deployer       common.Address
	ExpectedAddr   common.Address
}

func (cfg DeploymentConfig) Validate() error {
	if (cfg.Create3Factory == common.Address{}) {
		return errors.New("create3 factory is zero")
	}
	if (cfg.Owner == common.Address{}) {
		return errors.New("owner is zero")
	}
	if (cfg.Deployer == common.Address{}) {
		return errors.New("deployer is zero")
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
		Create3Salt:    contracts.ProxyAdminSalt(netconf.Mainnet),
		Owner:          contracts.MainnetProxyAdminOwner(),
		Deployer:       contracts.MainnetDeployer(),
	}
}

func testnetCfg() DeploymentConfig {
	return DeploymentConfig{
		Create3Factory: contracts.TestnetCreate3Factory(),
		Create3Salt:    contracts.ProxyAdminSalt(netconf.Testnet),
		Owner:          contracts.TestnetProxyAdminOwner(),
		Deployer:       contracts.TestnetDeployer(),
	}
}

func stagingCfg() DeploymentConfig {
	return DeploymentConfig{
		Create3Factory: contracts.StagingCreate3Factory(),
		Create3Salt:    contracts.ProxyAdminSalt(netconf.Staging),
		Owner:          contracts.StagingProxyAdminOwner(),
		Deployer:       contracts.StagingDeployer(),
		ExpectedAddr:   contracts.StagingProxyAdmin(),
	}
}

func devnetCfg() DeploymentConfig {
	return DeploymentConfig{
		Create3Factory: contracts.DevnetCreate3Factory(),
		Create3Salt:    contracts.ProxyAdminSalt(netconf.Devnet),
		Owner:          contracts.DevnetProxyAdminOwner(),
		Deployer:       contracts.DevnetDeployer(),
		ExpectedAddr:   contracts.DevnetProxyAdmin(),
	}
}

func AddrForNetwork(network netconf.ID) (common.Address, bool) {
	switch network {
	case netconf.Mainnet:
		return contracts.MainnetProxyAdmin(), true
	case netconf.Testnet:
		return contracts.TestnetProxyAdmin(), true
	case netconf.Staging:
		return contracts.StagingProxyAdmin(), true
	case netconf.Devnet:
		return contracts.DevnetProxyAdmin(), true
	default:
		return common.Address{}, false
	}
}

// IsDeployed checks if the ProxyAdmin contract is deployed to the provided backend
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

// DeployIfNeeded deploys a new ProxyAdmin contract if it is not already deployed.
func DeployIfNeeded(ctx context.Context, network netconf.ID, backend *ethbackend.Backend) (contracts.Deployment, error) {
	deployed, deployment, err := IsDeployed(ctx, network, backend)
	if err != nil {
		return contracts.Deployment{}, errors.Wrap(err, "is deployed")
	}
	if deployed {
		return deployment, nil
	}

	return Deploy(ctx, network, backend)
}

// Deploy deploys a new ProxyAdmin contract and returns the address and receipt.
// It only allows deployments to explicitly supported chains.
func Deploy(ctx context.Context, network netconf.ID, backend *ethbackend.Backend) (contracts.Deployment, error) {
	chainID, err := backend.ChainID(ctx)
	if err != nil {
		return contracts.Deployment{}, errors.Wrap(err, "chain id")
	}

	cfg, err := getDeployCfg(chainID.Uint64(), network)
	if err != nil {
		return contracts.Deployment{}, errors.Wrap(err, "get deployment config")
	}

	return deploy(ctx, cfg, backend)
}

func deploy(ctx context.Context, cfg DeploymentConfig, backend *ethbackend.Backend) (contracts.Deployment, error) {
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

	initCode, err := packInitCode(cfg)
	if err != nil {
		return contracts.Deployment{}, errors.Wrap(err, "pack init code")
	}

	tx, err := factory.Deploy(txOpts, salt, initCode)
	if err != nil {
		return contracts.Deployment{}, errors.Wrap(err, "deploy proxy admin")
	}

	receipt, err := bind.WaitMined(ctx, backend, tx)
	if err != nil {
		return contracts.Deployment{}, errors.Wrap(err, "wait mined proxy admin")
	} else if receipt.Status != ethtypes.ReceiptStatusSuccessful {
		return contracts.Deployment{}, errors.New("deploy proxy failed")
	}

	deployment := contracts.Deployment{
		Address:     addr,
		BlockHeight: receipt.BlockNumber.Uint64(),
	}

	return deployment, nil
}

func packInitCode(cfg DeploymentConfig) ([]byte, error) {
	abi, err := bindings.ProxyAdminMetaData.GetAbi()
	if err != nil {
		return nil, errors.Wrap(err, "get abi")
	}

	return contracts.PackInitCode(abi, bindings.ProxyAdminBin, cfg.Owner)
}
