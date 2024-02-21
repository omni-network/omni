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
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

type XDapp struct {
	// Immutable state
	cfg         AVSConfig
	eigen       EigenDeployments
	portalAddr  common.Address
	chain       types.EVMChain
	omniChainID uint64
	backends    backend.Backends

	// Mutable state
	contract     *bindings.OmniAVS
	contractAddr common.Address
	height       uint64
}

func New(cfg AVSConfig, eigen EigenDeployments, portalAddr common.Address,
	chain types.EVMChain, omniChainID uint64, backends backend.Backends) *XDapp {
	return &XDapp{
		cfg:         cfg,
		eigen:       eigen,
		portalAddr:  portalAddr,
		chain:       chain,
		omniChainID: omniChainID,
		backends:    backends,
	}
}

func (d *XDapp) Deploy(ctx context.Context) error {
	if d.contract != nil {
		return errors.New("avs already deployed")
	}

	log.Info(ctx, "Deploying AVS contracts", "chain", d.chain.Name)

	txOpts, backend, err := d.backends.BindOpts(ctx, d.chain.ID)
	if err != nil {
		return err
	}

	height, err := backend.BlockNumber(ctx)
	if err != nil {
		return errors.Wrap(err, "get block number")
	}

	// TODO: use same proxy admin for portal & avs on same chain
	proxyAdmin, err := netman.DeployProxyAdmin(ctx, txOpts, backend)
	if err != nil {
		return errors.Wrap(err, "deploy proxy admin")
	}

	addr, err := d.deployOmniAVS(ctx, backend, txOpts, proxyAdmin, txOpts.From)
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

	log.Debug(ctx, "Deployed AVS contract", "address", addr.Hex(), "chain", d.chain.Name)

	return nil
}

// ExportDeployInfo sets the contract addresses in the given DeployInfos.
func (d *XDapp) ExportDeployInfo(i types.DeployInfos) {
	i.Set(d.chain.ID, types.ContractOmniAVS, d.contractAddr, d.height)

	const elHeight uint64 = 0 // TODO(corver): Maybe figure this out?

	i.Set(d.chain.ID, types.ContractELAVSDirectory, d.eigen.AVSDirectory, elHeight)
	i.Set(d.chain.ID, types.ContractELDelegationManager, d.eigen.DelegationManager, elHeight)
	i.Set(d.chain.ID, types.ContractELStrategyManager, d.eigen.StrategyManager, elHeight)
	i.Set(d.chain.ID, types.ContractELPodManager, d.eigen.EigenPodManager, elHeight)
	i.Set(d.chain.ID, types.ContractELWETH, d.eigen.WETH, elHeight)
	i.Set(d.chain.ID, types.ContractELWETHStrategy, d.eigen.WETHStrategy, elHeight)

}

func (d *XDapp) deployOmniAVS(ctx context.Context, client backend.Backend, txOpts *bind.TransactOpts,
	proxyAdmin common.Address, owner common.Address,
) (common.Address, error) {
	impl, tx, _, err := bindings.DeployOmniAVS(txOpts, client, d.eigen.DelegationManager, d.eigen.AVSDirectory)
	if err != nil {
		return common.Address{}, errors.Wrap(err, "deploy avs impl")
	}

	receipt, err := bind.WaitMined(ctx, client, tx)
	if err != nil {
		return common.Address{}, errors.Wrap(err, "wait mined avs impl")
	} else if receipt.Status != ethtypes.ReceiptStatusSuccessful {
		return common.Address{}, errors.New("deploy avs impl failed")
	}

	abi, err := bindings.OmniAVSMetaData.GetAbi()
	if err != nil {
		return common.Address{}, errors.Wrap(err, "get avs abi")
	}

	stratParms := make([]bindings.IOmniAVSStrategyParams, len(d.cfg.StrategyParams))
	for i, sp := range d.cfg.StrategyParams {
		stratParms[i] = bindings.IOmniAVSStrategyParams{
			Strategy:   sp.Strategy,
			Multiplier: sp.Multiplier,
		}
	}

	enc, err := abi.Pack("initialize", owner, d.portalAddr, d.omniChainID,
		d.cfg.MinimumOperatorStake, d.cfg.MaximumOperatorCount, []common.Address{} /* allowlist */, stratParms)
	if err != nil {
		return common.Address{}, errors.Wrap(err, "encode avs initializer")
	}

	proxy, tx, _, err := bindings.DeployTransparentUpgradeableProxy(txOpts, client, impl, proxyAdmin, enc)
	if err != nil {
		return common.Address{}, errors.Wrap(err, "deploy avs proxy")
	}

	receipt, err = bind.WaitMined(ctx, client, tx)
	if err != nil {
		return common.Address{}, errors.Wrap(err, "wait mined avs proxy")
	} else if receipt.Status != ethtypes.ReceiptStatusSuccessful {
		return common.Address{}, errors.New("deploy avs proxy failed")
	}

	return proxy, nil
}
