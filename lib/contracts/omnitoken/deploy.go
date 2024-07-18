package omnitoken

import (
	"context"
	"math/big"

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

// NOTE: mainnet not supported, as mainnet token is already deployed.

var configs = map[netconf.ID]DeploymentConfig{
	netconf.Omega: {
		Create3Factory: contracts.OmegaCreate3Factory(),
		Create3Salt:    contracts.TokenSalt(netconf.Omega),
		Deployer:       eoa.MustAddress(netconf.Omega, eoa.RoleDeployer),
		Recipient:      eoa.MustAddress(netconf.Omega, eoa.RoleTester),
		ExpectedAddr:   contracts.OmegaToken(),
	},
	netconf.Staging: {
		Create3Factory: contracts.StagingCreate3Factory(),
		Create3Salt:    contracts.TokenSalt(netconf.Staging),
		Deployer:       eoa.MustAddress(netconf.Staging, eoa.RoleDeployer),
		Recipient:      eoa.MustAddress(netconf.Staging, eoa.RoleTester),
		ExpectedAddr:   contracts.StagingToken(),
	},
	netconf.Devnet: {
		Create3Factory: contracts.DevnetCreate3Factory(),
		Create3Salt:    contracts.TokenSalt(netconf.Devnet),
		Deployer:       eoa.MustAddress(netconf.Devnet, eoa.RoleDeployer),
		Recipient:      eoa.MustAddress(netconf.Devnet, eoa.RoleTester),
		ExpectedAddr:   contracts.DevnetToken(),
	},
}

func InitialSupplyRecipient(network netconf.ID) (common.Address, bool) {
	if cfg, ok := configs[network]; ok {
		return cfg.Recipient, true
	}

	return common.Address{}, false
}

func AddrForNetwork(network netconf.ID) (common.Address, bool) {
	switch network {
	case netconf.Mainnet:
		return contracts.MainnetToken(), true
	case netconf.Omega:
		return contracts.OmegaToken(), true
	case netconf.Staging:
		return contracts.StagingToken(), true
	case netconf.Devnet:
		return contracts.DevnetToken(), true
	default:
		return common.Address{}, false
	}
}

// isDeployed returns true if the token contract is already deployed to its expected address.
func isDeployed(ctx context.Context, network netconf.ID, backend *ethbackend.Backend) (bool, common.Address, error) {
	addr, ok := AddrForNetwork(network)
	if !ok {
		return false, addr, errors.New("unsupported network", "network", network)
	}

	code, err := backend.CodeAt(ctx, addr, nil)
	if err != nil {
		return false, addr, errors.Wrap(err, "code at", "address", addr)
	}

	if len(code) == 0 {
		return false, addr, nil
	}

	return true, addr, nil
}

// DeployIfNeeded deploys a new token contract if it is not already deployed.
// If the contract is already deployed, the receipt is nil.
func DeployIfNeeded(ctx context.Context, network netconf.ID, backend *ethbackend.Backend) (common.Address, *ethtypes.Receipt, error) {
	deployed, addr, err := isDeployed(ctx, network, backend)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "is deployed")
	}
	if deployed {
		return addr, nil, nil
	}

	return Deploy(ctx, network, backend)
}

// Deploy deploys a new ERC20 OMNI token contract and returns the address and receipt.
//
// NOTE: the mainnet ERC20 OMNI token is already deployed to ETH mainnet. We use
// this code for test / ephemeral networks.
func Deploy(ctx context.Context, network netconf.ID, backend *ethbackend.Backend) (common.Address, *ethtypes.Receipt, error) {
	cfg, ok := configs[network]
	if !ok {
		return common.Address{}, nil, errors.New("unsupported network", "network", network)
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
