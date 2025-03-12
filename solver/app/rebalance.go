package app

import (
	"context"
	"math/big"
	"time"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/lib/contracts"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/umath"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"
)

// maxL1OMNI is the max amount of L1 OMNI the solver should hold by network.
var maxL1OMNI = map[netconf.ID]*big.Int{
	// 1 OMNI for ephemeral networks (tests rebalancing more frequently)
	netconf.Devnet:  new(big.Int).Mul(big.NewInt(1), big.NewInt(params.Ether)),
	netconf.Staging: new(big.Int).Mul(big.NewInt(1), big.NewInt(params.Ether)),

	// 1000 OMNI for protected networks (reduces gas spend)
	netconf.Omega:   new(big.Int).Mul(big.NewInt(1000), big.NewInt(params.Ether)),
	netconf.Mainnet: new(big.Int).Mul(big.NewInt(1000), big.NewInt(params.Ether)),
}

// startRebalancing starts rebalancing of tokens that the solver is able to rebalance.
func startRebalancing(ctx context.Context, network netconf.Network, backends ethbackend.Backends) error {
	if err := startRebalancingOMNI(ctx, network, backends); err != nil {
		return errors.Wrap(err, "rebalance OMNI")
	}

	return nil
}

// startRebalancingOMNI starts the rebalancing of OMNI tokens.
func startRebalancingOMNI(ctx context.Context, network netconf.Network, backends ethbackend.Backends) error {
	l1, ok := network.EthereumChain()

	// if no l1, nothing to rebalance
	if !ok {
		return nil
	}

	backend, err := backends.Backend(l1.ID)
	if err != nil {
		return errors.Wrap(err, "get backend")
	}

	go rebalanceOMNIForever(ctx, network, l1, backend)

	return nil
}

// rebalanceOMNIForever periodically bridges L1 ERC20 OMNI back to Omni's EVM.
func rebalanceOMNIForever(ctx context.Context, network netconf.Network, l1 netconf.Chain, backend *ethbackend.Backend) {
	ctx = log.WithCtx(ctx, "rebalancer", "OMNI", "chain", l1.Name, "network", network.ID)
	log.Info(ctx, "Rebalancing OMNI tokens")

	interval := time.Second * 30
	if network.ID.IsEphemeral() {
		interval = time.Second * 5
	}

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			err := rebalanceOMNIOnce(ctx, network, backend)
			if ctx.Err() != nil {
				return
			} else if err != nil {
				log.Error(ctx, "Rebalancing OMNI failed (will retry)", err)

				continue
			}
		}
	}
}

// rebalanceOMNIOnce moves any OMNI balance (above 1000) on L1 back to Omni's EVM.
func rebalanceOMNIOnce(ctx context.Context, network netconf.Network, backend *ethbackend.Backend) error {
	solverAddr := eoa.MustAddress(network.ID, eoa.RoleSolver)
	tokenAddr := contracts.TokenAddr(network.ID)

	token, err := bindings.NewIERC20(tokenAddr, backend)
	if err != nil {
		return errors.Wrap(err, "new erc20")
	}

	balance, err := token.BalanceOf(&bind.CallOpts{Context: ctx}, solverAddr)
	if err != nil {
		return errors.Wrap(err, "get balance")
	}

	// if balance not above max, do nothing
	if balance.Cmp(maxL1OMNI[network.ID]) < 0 {
		return nil
	}

	addrs, err := contracts.GetAddresses(ctx, network.ID)
	if err != nil {
		return errors.Wrap(err, "get addrs")
	}

	txOpts, err := backend.BindOpts(ctx, solverAddr)
	if err != nil {
		return errors.Wrap(err, "bind opts")
	}

	if err := maybeApprove(ctx, backend, token, solverAddr, addrs.L1Bridge, balance); err != nil {
		return errors.Wrap(err, "maybe approve")
	}

	bridge, err := bindings.NewOmniBridgeL1(addrs.L1Bridge, backend)
	if err != nil {
		return errors.Wrap(err, "l1 bridge")
	}

	fee, err := bridge.BridgeFee(&bind.CallOpts{Context: ctx}, txOpts.From, solverAddr, balance)
	if err != nil {
		return errors.Wrap(err, "bridge fee")
	}

	txOpts.Value = fee

	tx, err := bridge.Bridge(txOpts, solverAddr, balance)
	if err != nil {
		return errors.Wrap(err, "bridge")
	}

	_, err = backend.WaitMined(ctx, tx)
	if err != nil {
		return errors.Wrap(err, "wait mined")
	}

	log.Info(ctx, "Bridged L1-to-native OMNI", "to", solverAddr, "amount", balance, "fee", fee, "tx", tx.Hash())

	return nil
}

func maybeApprove(
	ctx context.Context,
	backend *ethbackend.Backend,
	token *bindings.IERC20,
	owner, spender common.Address,
	amount *big.Int,
) error {
	allowance, err := token.Allowance(&bind.CallOpts{Context: ctx}, owner, spender)
	if err != nil {
		return errors.Wrap(err, "allowance")
	}

	if allowance.Cmp(amount) >= 0 {
		return nil
	}

	txOpts, err := backend.BindOpts(ctx, owner)
	if err != nil {
		return errors.Wrap(err, "bind opts")
	}

	tx, err := token.Approve(txOpts, spender, umath.MaxUint256)
	if err != nil {
		return errors.Wrap(err, "approve")
	}

	_, err = backend.WaitMined(ctx, tx)
	if err != nil {
		return errors.Wrap(err, "wait mined")
	}

	return nil
}
