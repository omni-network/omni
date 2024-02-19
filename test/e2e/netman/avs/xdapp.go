package avs

import (
	"context"
	"math/big"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/test/e2e/netman"
	"github.com/omni-network/omni/test/e2e/types"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type XDapp struct {
	// Immutable state
	cfg        AVSConfig
	eigen      EigenDeployments
	portalAddr common.Address
	chain      types.EVMChain
	ethCl      *ethclient.Client
	txOpts     *bind.TransactOpts // TODO(corver): Replace this with a txmgr.

	// Mutable state
	contract     *bindings.OmniAVS
	contractAddr common.Address
	height       uint64
}

func New(cfg AVSConfig, eigen EigenDeployments, portalAddr common.Address,
	chain types.EVMChain, ethCl *ethclient.Client, txOpts *bind.TransactOpts) *XDapp {
	return &XDapp{
		cfg:        cfg,
		eigen:      eigen,
		portalAddr: portalAddr,
		chain:      chain,
		ethCl:      ethCl,
		txOpts:     txOpts,
	}
}

func (m *XDapp) Deploy(ctx context.Context) error {
	if m.contract != nil {
		return errors.New("avs already deployed")
	}

	log.Info(ctx, "Deploying AVS contracts", "chain", m.chain.Name)

	if m.ethCl == nil {
		return errors.New("avs client not set")
	}

	height, err := m.ethCl.BlockNumber(ctx)
	if err != nil {
		return errors.Wrap(err, "get block number")
	}

	// TODO: use same proxy admin for portal & avs on same chain
	proxyAdmin, err := netman.DeployProxyAdmin(ctx, m.txOpts, m.ethCl)
	if err != nil {
		return errors.Wrap(err, "deploy proxy admin")
	}

	stratParms := make([]bindings.IOmniAVSStrategyParams, len(m.cfg.StrategyParams))
	for i, sp := range m.cfg.StrategyParams {
		stratParms[i] = bindings.IOmniAVSStrategyParams{
			Strategy:   sp.Strategy,
			Multiplier: sp.Multiplier,
		}
	}

	addr, err := deployOmniAVS(ctx, m.ethCl, m.txOpts, proxyAdmin, m.txOpts.From,
		m.portalAddr, m.chain.ID, m.eigen.DelegationManager, m.eigen.AVSDirectory,
		m.cfg.MinimumOperatorStake, m.cfg.MaximumOperatorCount, stratParms)
	if err != nil {
		return errors.Wrap(err, "deploy avs")
	}

	contract, err := bindings.NewOmniAVS(addr, m.ethCl)
	if err != nil {
		return errors.Wrap(err, "instantiate avs")
	}

	m.contract = contract
	m.contractAddr = addr
	m.height = height

	log.Info(ctx, "Deployed AVS contract", "address", addr.Hex(), "chain", m.chain.Name)

	return nil
}

// ExportDeployInfo sets the contract addresses in the given DeployInfos.
func (m *XDapp) ExportDeployInfo(i types.DeployInfos) {
	i.Set(m.chain.ID, types.ContractOmniAVS, m.contractAddr, m.height)

	const elHeight uint64 = 0 // TODO(corver): Maybe figure this out?
	i.Set(m.chain.ID, types.ContractELAVSDirectory, m.eigen.AVSDirectory, elHeight)
	i.Set(m.chain.ID, types.ContractELDelegationManager, m.eigen.DelegationManager, elHeight)
	i.Set(m.chain.ID, types.ContractELStrategyManager, m.eigen.StrategyManager, elHeight)
	i.Set(m.chain.ID, types.ContractELPodManager, m.eigen.EigenPodManager, elHeight)
}

func deployOmniAVS(ctx context.Context, client *ethclient.Client, txOpts *bind.TransactOpts,
	proxyAdmin common.Address, owner common.Address, portal common.Address, omniChainID uint64,
	delegationManager common.Address, avsDirectory common.Address, minOperatorStake *big.Int,
	maxOperators uint32, strategyParams []bindings.IOmniAVSStrategyParams,
) (common.Address, error) {
	impl, tx, _, err := bindings.DeployOmniAVS(txOpts, client, delegationManager, avsDirectory)
	if err != nil {
		return common.Address{}, errors.Wrap(err, "deploy avs impl")
	}

	receipt, err := bind.WaitMined(ctx, client, tx)
	if err != nil || receipt.Status != ethtypes.ReceiptStatusSuccessful {
		return common.Address{}, errors.Wrap(err, "wait mined avs impl")
	}

	abi, err := bindings.OmniAVSMetaData.GetAbi()
	if err != nil {
		return common.Address{}, errors.Wrap(err, "get avs abi")
	}

	enc, err := abi.Pack("initialize", owner, portal, omniChainID,
		minOperatorStake, maxOperators, strategyParams)
	if err != nil {
		return common.Address{}, errors.Wrap(err, "encode avs initializer")
	}

	proxy, tx, _, err := bindings.DeployTransparentUpgradeableProxy(txOpts, client, impl, proxyAdmin, enc)
	if err != nil {
		return common.Address{}, errors.Wrap(err, "deploy avs proxy")
	}

	receipt, err = bind.WaitMined(ctx, client, tx)
	if err != nil || receipt.Status != ethtypes.ReceiptStatusSuccessful {
		return common.Address{}, errors.Wrap(err, "wait mined avs proxy")
	}

	return proxy, nil
}
