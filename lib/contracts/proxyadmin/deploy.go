package proxyadmin

import (
	"context"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/chainids"
	"github.com/omni-network/omni/lib/contracts"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

type DeploymentConfig struct {
	Create3Factory common.Address
	Create3Salt    string
	Owner          common.Address
	Deployer       common.Address
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

func getDeployCfg(chainID uint64, network string) (DeploymentConfig, error) {
	if chainids.IsMainnet(chainID) {
		return mainnetDeployCfg(network), nil
	}

	if chainids.IsTestnet(chainID) {
		return testnetDeployCfg(network), nil
	}

	return DeploymentConfig{}, errors.New("unsupported chain id")
}

func mainnetDeployCfg(network string) DeploymentConfig {
	return DeploymentConfig{
		Create3Factory: contracts.MainnetCreate3Factory,
		Create3Salt:    network + "-proxy-admin",
		Owner:          contracts.MainnetProxyAdminOwner,
		Deployer:       contracts.MainnetDeployer,
	}
}

func testnetDeployCfg(network string) DeploymentConfig {
	return DeploymentConfig{
		Create3Factory: contracts.TestnetCreate3Factory,
		Create3Salt:    network + "-proxy-admin",
		Owner:          contracts.TestnetProxyAdminOwner,
		Deployer:       contracts.TestnetDeployer,
	}
}

func devnetDeployCfg() DeploymentConfig {
	return DeploymentConfig{
		Create3Factory: common.HexToAddress("0x1234"), // TODO: currently unused
		Create3Salt:    "devnet-proxy-admin",
		// anvil account 0
		Owner:    common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"),
		Deployer: common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"),
	}
}

// Deploy deploys a new ProxyAdmin contract and returns the address and receipt.
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

func deploy(ctx context.Context, cfg DeploymentConfig, backend *ethbackend.Backend) (common.Address, *ethtypes.Receipt, error) {
	txOpts, err := backend.BindOpts(ctx, cfg.Deployer)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "bind opts")
	}

	addr, tx, _, err := bindings.DeployProxyAdmin(txOpts, backend, cfg.Owner)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "deploy proxy admin")
	}

	receipt, err := bind.WaitMined(ctx, backend, tx)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "wait mined proxy admin")
	} else if receipt.Status != ethtypes.ReceiptStatusSuccessful {
		return common.Address{}, nil, errors.New("deploy proxy failed")
	}

	return addr, receipt, nil
}
