package avs

import (
	"context"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/test/e2e/backend"
	"github.com/omni-network/omni/test/e2e/netman"
	"github.com/omni-network/omni/test/e2e/types"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

type Contracts struct {
	OmniAVS           *bindings.OmniAVS
	DelegationManager *bindings.DelegationManager
	StrategyManager   *bindings.StrategyManager
	WETHStrategy      *bindings.StrategyBase
	WETHToken         *bindings.MockERC20
	AVSDirectory      *bindings.AVSDirectory

	OmniAVSAddr           common.Address
	DelegationManagerAddr common.Address
	StrategyManagerAddr   common.Address
	WETHStrategyAddr      common.Address
	WETHTokenAddr         common.Address
	AVSDirectoryAddr      common.Address
}

type Deployer struct {
	// Immutable state
	cfg         Config
	eigen       EigenDeployments
	portalAddr  common.Address
	omniChainID uint64
	chainID     uint64

	// Mutable state
	contract     *bindings.OmniAVS
	contractAddr common.Address
	height       uint64
}

func NewDeployer(
	cfg Config,
	eigen EigenDeployments,
	portalAddr common.Address,
	omniChainID uint64,
) *Deployer {
	return &Deployer{
		cfg:         cfg,
		eigen:       eigen,
		portalAddr:  portalAddr,
		omniChainID: omniChainID,
	}
}

// Contracts returns the deployed contracts.
func (d *Deployer) Contracts(backend backend.Backend) (Contracts, error) {
	if d.contract == nil {
		return Contracts{}, errors.New("avs not deployed")
	}

	delMan, err := bindings.NewDelegationManager(d.eigen.DelegationManager, backend)
	if err != nil {
		return Contracts{}, errors.Wrap(err, "delegation manager")
	}

	stratMan, err := bindings.NewStrategyManager(d.eigen.StrategyManager, backend)
	if err != nil {
		return Contracts{}, errors.Wrap(err, "strategy manager")
	}

	wethStratAddr, ok := d.eigen.Strategies["WETH"]
	if !ok || (wethStratAddr == common.Address{}) {
		return Contracts{}, errors.New("missing WETH strategy address")
	}

	wethStrategy, err := bindings.NewStrategyBase(wethStratAddr, backend)
	if err != nil {
		return Contracts{}, errors.Wrap(err, "weth strategy")
	}

	wethTokenAddr, err := wethStrategy.UnderlyingToken(&bind.CallOpts{})
	if err != nil {
		return Contracts{}, errors.Wrap(err, "underlying token")
	}

	wethToken, err := bindings.NewMockERC20(wethTokenAddr, backend)
	if err != nil {
		return Contracts{}, errors.Wrap(err, "weth token")
	}

	avsDir, err := bindings.NewAVSDirectory(d.eigen.AVSDirectory, backend)
	if err != nil {
		return Contracts{}, errors.Wrap(err, "avs directory")
	}

	return Contracts{
		OmniAVS:           d.contract,
		DelegationManager: delMan,
		StrategyManager:   stratMan,
		WETHStrategy:      wethStrategy,
		WETHToken:         wethToken,
		AVSDirectory:      avsDir,

		OmniAVSAddr:           d.contractAddr,
		DelegationManagerAddr: d.eigen.DelegationManager,
		StrategyManagerAddr:   d.eigen.StrategyManager,
		WETHStrategyAddr:      wethStratAddr,
		WETHTokenAddr:         wethTokenAddr,
		AVSDirectoryAddr:      d.eigen.AVSDirectory,
	}, nil
}

func (d *Deployer) Deploy(ctx context.Context, backend backend.Backend, owner common.Address) error {
	if d.contract != nil {
		return errors.New("avs already deployed")
	}

	chainName, chainID := backend.Chain()
	d.chainID = chainID

	log.Info(ctx, "Deploying AVS contracts", "chain", chainName)

	height, err := backend.BlockNumber(ctx)
	if err != nil {
		return errors.Wrap(err, "get block number")
	}

	txOpts, err := backend.BindOpts(ctx, owner)
	if err != nil {
		return err
	}

	// TODO: use same proxy admin for portal & avs on same chain
	proxyAdmin, err := netman.DeployProxyAdmin(ctx, txOpts, backend)
	if err != nil {
		return errors.Wrap(err, "deploy proxy admin")
	}

	addr, err := d.deployContracts(ctx, txOpts, backend, proxyAdmin, owner)
	if err != nil {
		return errors.Wrap(err, "deploy avs")
	}

	contract, err := bindings.NewOmniAVS(addr, backend)
	if err != nil {
		return errors.Wrap(err, "instantiate avs")
	}

	d.contract = contract
	d.contractAddr = addr
	d.height = height

	log.Debug(ctx, "Deployed AVS contract", "address", addr.Hex(), "chain", chainName)

	return nil
}

// ExportDeployInfo sets the contract addresses in the given DeployInfos.
func (d *Deployer) ExportDeployInfo(i types.DeployInfos) {
	i.Set(d.chainID, types.ContractOmniAVS, d.contractAddr, d.height)

	const elHeight uint64 = 0 // TODO(corver): Maybe figure this out?

	i.Set(d.chainID, types.ContractELAVSDirectory, d.eigen.AVSDirectory, elHeight)
	i.Set(d.chainID, types.ContractELDelegationManager, d.eigen.DelegationManager, elHeight)
	i.Set(d.chainID, types.ContractELStrategyManager, d.eigen.StrategyManager, elHeight)
	i.Set(d.chainID, types.ContractELPodManager, d.eigen.EigenPodManager, elHeight)
	i.Set(d.chainID, types.ContractELWETHStrategy, d.eigen.Strategies["WETH"], elHeight)
}

func (d *Deployer) deployContracts(ctx context.Context, txOpts *bind.TransactOpts, backend backend.Backend,
	proxyAdmin common.Address, owner common.Address,
) (common.Address, error) {
	if txOpts.From != owner {
		return common.Address{}, errors.New("txOpts not from owner")
	}

	impl, tx, _, err := bindings.DeployOmniAVS(txOpts, backend, d.eigen.DelegationManager, d.eigen.AVSDirectory)
	if err != nil {
		return common.Address{}, errors.Wrap(err, "deploy avs impl")
	}
	_, err = backend.WaitMined(ctx, tx)
	if err != nil {
		return common.Address{}, errors.Wrap(err, "wait mined avs proxy")
	}

	abi, err := bindings.OmniAVSMetaData.GetAbi()
	if err != nil {
		return common.Address{}, errors.Wrap(err, "get avs abi")
	}

	stratParms := make([]bindings.IOmniAVSStrategyParam, len(d.cfg.StrategyParams))
	for i, sp := range d.cfg.StrategyParams {
		stratParms[i] = bindings.IOmniAVSStrategyParam{
			Strategy:   sp.Strategy,
			Multiplier: sp.Multiplier,
		}
	}

	enc, err := abi.Pack("initialize", owner, d.portalAddr, d.omniChainID, d.cfg.EthStakeInbox, stratParms)
	if err != nil {
		return common.Address{}, errors.Wrap(err, "encode avs initializer")
	}

	proxy, tx, _, err := bindings.DeployTransparentUpgradeableProxy(txOpts, backend, impl, proxyAdmin, enc)
	if err != nil {
		return common.Address{}, errors.Wrap(err, "deploy avs proxy")
	}
	_, err = backend.WaitMined(ctx, tx)
	if err != nil {
		return common.Address{}, errors.Wrap(err, "wait mined avs proxy")
	}

	return proxy, nil
}
