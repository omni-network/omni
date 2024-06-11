package omnitoken

import (
	"context"
	"math/big"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/lib/anvil"
	"github.com/omni-network/omni/lib/contracts"
	"github.com/omni-network/omni/lib/create3"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

// TotalSupply is the 100M, total supply of the token.
var TotalSupply = new(big.Int).Mul(big.NewInt(100e6), big.NewInt(1e18))

type DeploymentConfig struct {
	Create3Factory common.Address
	Create3Salt    string
	Deployer       common.Address
	Recipient      common.Address
	ExpectedAddr   common.Address
}

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
	if isDeadOrEmpty(cfg.Deployer) {
		return errors.New("deployer is not set")
	}
	if isDeadOrEmpty(cfg.Recipient) {
		return errors.New("recipient is not set")
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

	if network == netconf.Staging {
		return stagingCfg(), nil
	}

	return DeploymentConfig{}, errors.New("unsupported network", "network", network)
}

func stagingCfg() DeploymentConfig {
	return DeploymentConfig{
		Create3Factory: contracts.StagingCreate3Factory(),
		Create3Salt:    contracts.TokenSalt(netconf.Staging),
		Deployer:       eoa.MustAddress(netconf.Staging, eoa.RoleDeployer),
		Recipient:      eoa.MustAddress(netconf.Staging, eoa.RoleFbDev),
		ExpectedAddr:   contracts.StagingToken(),
	}
}

func devnetCfg() DeploymentConfig {
	return DeploymentConfig{
		Create3Factory: contracts.DevnetCreate3Factory(),
		Create3Salt:    contracts.TokenSalt(netconf.Devnet),
		Deployer:       eoa.MustAddress(netconf.Devnet, eoa.RoleDeployer),
		Recipient:      anvil.DevAccount7(),
		ExpectedAddr:   contracts.DevnetToken(),
	}
}

func InitialSupplyRecipient(network netconf.ID) (common.Address, bool) {
	if network == netconf.Devnet {
		return anvil.DevAccount7(), true
	}

	if network == netconf.Staging {
		return eoa.MustAddress(netconf.Staging, eoa.RoleFbDev), true
	}

	return common.Address{}, false
}

// Deploy deploys a new ERC20 OMNI token contract and returns the address and receipt.
func Deploy(ctx context.Context, network netconf.ID, backend *ethbackend.Backend) (common.Address, *ethtypes.Receipt, error) {
	cfg, err := getDeployCfg(network)
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

	initCode, err := packInitCode(cfg)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "pack init code")
	}

	tx, err := factory.Deploy(txOpts, salt, initCode)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "deploy proxy")
	}

	receipt, err := backend.WaitMined(ctx, tx)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "wait mined proxy")
	}

	return addr, receipt, nil
}

func packInitCode(cfg DeploymentConfig) ([]byte, error) {
	abi, err := bindings.OmniMetaData.GetAbi()
	if err != nil {
		return nil, errors.Wrap(err, "get abi")
	}

	return contracts.PackInitCode(abi, bindings.OmniBin, TotalSupply, cfg.Recipient)
}
