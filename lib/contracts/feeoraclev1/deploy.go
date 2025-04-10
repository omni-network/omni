package feeoraclev1

import (
	"context"
	"math/big"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/tokenpricer/coingecko"

	"github.com/ethereum/go-ethereum/common"
)

type deploymentConfig struct {
	Owner           common.Address
	Manager         common.Address // manager is the address that can set fee parameters (gas price, conversion rates)
	Deployer        common.Address
	ProxyAdminOwner common.Address
	BaseGasLimit    uint64
	ProtocolFee     *big.Int
}

func isEmpty(addr common.Address) bool {
	return addr == common.Address{}
}

func (cfg deploymentConfig) Validate() error {
	if isEmpty(cfg.Owner) {
		return errors.New("owner is zero")
	}
	if isEmpty(cfg.Manager) {
		return errors.New("manager is zero")
	}
	if isEmpty(cfg.Deployer) {
		return errors.New("deployer is zero")
	}
	if (cfg.ProxyAdminOwner == common.Address{}) {
		return errors.New("proxy admin is zero")
	}

	return nil
}

func Deploy(ctx context.Context, network netconf.ID, chainID uint64, destChainIDs []uint64, backends ethbackend.Backends) (common.Address, *ethclient.Receipt, error) {
	cfg := deploymentConfig{
		Owner:           eoa.MustAddress(network, eoa.RoleManager),
		Manager:         eoa.MustAddress(network, eoa.RoleMonitor), // NOTE: monitor is owner of fee oracle contracts, because monitor manages on chain gas prices / conversion rates
		Deployer:        eoa.MustAddress(network, eoa.RoleDeployer),
		ProxyAdminOwner: eoa.MustAddress(network, eoa.RoleUpgrader),
		BaseGasLimit:    50_000,
		ProtocolFee:     bi.Zero(),
	}
	if err := cfg.Validate(); err != nil {
		return common.Address{}, nil, errors.Wrap(err, "validate cfg")
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
	log.Debug(ctx, "Fee oracle impl", "addr", impl.Hex(), "tx", tx.Hash().Hex(), "chain_id", chainID)

	_, err = backend.WaitMined(ctx, tx)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "wait mined")
	}

	proxy, tx, _, err := bindings.DeployTransparentUpgradeableProxy(txOpts, backend, impl, cfg.ProxyAdminOwner, initializer)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "deploy proxy")
	}
	log.Debug(ctx, "Fee oracle proxy", "addr", proxy.Hex(), "tx", tx.Hash().Hex(), "chain_id", chainID)

	receipt, err := backend.WaitMined(ctx, tx)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "wait mined")
	}

	return proxy, receipt, nil
}
