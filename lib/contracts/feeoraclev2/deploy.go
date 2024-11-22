package feeoraclev2

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
	"github.com/omni-network/omni/lib/tokens/coingecko"

	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

type DeploymentConfig struct {
	Create3Salt     string
	Create3Factory  common.Address
	ExpectedAddr    common.Address
	ProxyAdminOwner common.Address
	Owner           common.Address
	Deployer        common.Address
	Manager         common.Address // manager is the address that can set fee parameters (gas price, conversion rates)
	BaseGasLimit    *big.Int       // must fit in uint24 (max: 16,777,215)
	ProtocolFee     *big.Int       // must fit in uint72 (max: 4,722,366,482,869,645,213,695)
}

func isEmpty(addr common.Address) bool {
	return addr == common.Address{}
}

var maxUint24 = new(big.Int).Sub(new(big.Int).Lsh(big.NewInt(1), 24), big.NewInt(1))
var maxUint72 = new(big.Int).Sub(new(big.Int).Lsh(big.NewInt(1), 72), big.NewInt(1))

func (cfg DeploymentConfig) Validate() error {
	if cfg.Create3Salt == "" {
		return errors.New("create3 salt is empty")
	}
	if isEmpty(cfg.Create3Factory) {
		return errors.New("create3 factory is zero")
	}
	if isEmpty(cfg.ExpectedAddr) {
		return errors.New("expected address is zero")
	}
	if isEmpty(cfg.ProxyAdminOwner) {
		return errors.New("proxy admin is zero")
	}
	if isEmpty(cfg.Owner) {
		return errors.New("owner is zero")
	}
	if isEmpty(cfg.Deployer) {
		return errors.New("deployer is zero")
	}
	if isEmpty(cfg.Manager) {
		return errors.New("manager is zero")
	}
	if cfg.BaseGasLimit.Cmp(maxUint24) > 0 {
		return errors.New("base gas limit too high")
	}
	if cfg.ProtocolFee.Cmp(maxUint72) > 0 {
		return errors.New("protocol fee too high")
	}

	return nil
}

// isDeployed returns true if the oracle contract is already deployed to its expected address.
func isDeployed(ctx context.Context, network netconf.ID, backend *ethbackend.Backend) (bool, common.Address, error) {
	addrs, err := contracts.GetAddresses(ctx, network)
	if err != nil {
		return false, common.Address{}, errors.Wrap(err, "get addrs")
	}

	code, err := backend.CodeAt(ctx, addrs.FeeOracleV2, nil)
	if err != nil {
		return false, addrs.FeeOracleV2, errors.Wrap(err, "code at", "address", addrs.FeeOracleV2)
	}

	if len(code) == 0 {
		return false, addrs.FeeOracleV2, nil
	}

	return true, addrs.FeeOracleV2, nil
}

// DeployIfNeeded deploys a new oracle contract if it is not already deployed.
// If the contract is already deployed, the receipt is nil.
func DeployIfNeeded(ctx context.Context, network netconf.ID, chainID uint64, destChainIDs []uint64, backends ethbackend.Backends) (common.Address, *ethtypes.Receipt, error) {
	backend, err := backends.Backend(chainID)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "get backend")
	}

	deployed, addr, err := isDeployed(ctx, network, backend)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "is deployed")
	}
	if deployed {
		return addr, nil, nil
	}

	return Deploy(ctx, network, chainID, destChainIDs, backend, backends)
}

// Deploy deploys a new FeeOracleV2 contract and returns the address and receipt.
func Deploy(ctx context.Context, network netconf.ID, chainID uint64, destChainIDs []uint64, backend *ethbackend.Backend, backends ethbackend.Backends) (common.Address, *ethtypes.Receipt, error) {
	addrs, err := contracts.GetAddresses(ctx, network)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "get addresses")
	}

	salts, err := contracts.GetSalts(ctx, network)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "get salts")
	}

	cfg := DeploymentConfig{
		Create3Salt:     salts.FeeOracleV2,
		Create3Factory:  addrs.Create3Factory,
		ExpectedAddr:    addrs.FeeOracleV2,
		ProxyAdminOwner: eoa.MustAddress(network, eoa.RoleUpgrader),
		Owner:           eoa.MustAddress(network, eoa.RoleManager),
		Deployer:        eoa.MustAddress(network, eoa.RoleDeployer),
		Manager:         eoa.MustAddress(network, eoa.RoleMonitor), // NOTE: monitor is owner of fee oracle contracts, because monitor manages on chain gas prices / conversion rates
		BaseGasLimit:    big.NewInt(100_000),
		ProtocolFee:     big.NewInt(0),
	}

	return deploy(ctx, chainID, destChainIDs, cfg, backend, backends)
}

func deploy(ctx context.Context, chainID uint64, destChainIDs []uint64, cfg DeploymentConfig, backend *ethbackend.Backend, backends ethbackend.Backends) (common.Address, *ethtypes.Receipt, error) {
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

	impl, tx, _, err := bindings.DeployFeeOracleV2(txOpts, backend)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "deploy implementation")
	}

	_, err = backend.WaitMined(ctx, tx)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "wait mined implementation")
	}

	initCode, err := packInitCode(ctx, chainID, destChainIDs, cfg, backends, impl)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "pack init code")
	}

	tx, err = factory.DeployWithRetry(txOpts, salt, initCode) //nolint:contextcheck // Context is txOpts
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "deploy proxy")
	}

	receipt, err := backend.WaitMined(ctx, tx)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "wait mined proxy")
	}

	return addr, receipt, nil
}

func packInitCode(ctx context.Context, chainID uint64, destChainIDs []uint64, cfg DeploymentConfig, backends ethbackend.Backends, impl common.Address) ([]byte, error) {
	feeOracleAbi, err := bindings.FeeOracleV2MetaData.GetAbi()
	if err != nil {
		return nil, errors.Wrap(err, "get fee oracle abi")
	}

	proxyAbi, err := bindings.TransparentUpgradeableProxyMetaData.GetAbi()
	if err != nil {
		return nil, errors.Wrap(err, "get proxy abi")
	}

	feeparams, err := feeParams(ctx, chainID, destChainIDs, backends, coingecko.New())
	if err != nil {
		return nil, errors.Wrap(err, "fee params")
	}

	initializer, err := feeOracleAbi.Pack("initialize", cfg.Owner, cfg.Manager, cfg.BaseGasLimit, cfg.ProtocolFee, feeparams)
	if err != nil {
		return nil, errors.Wrap(err, "pack initialize")
	}

	return contracts.PackInitCode(proxyAbi, bindings.TransparentUpgradeableProxyBin, impl, cfg.ProxyAdminOwner, initializer)
}
