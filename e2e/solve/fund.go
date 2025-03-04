package solve

import (
	"context"
	"math/big"

	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/lib/anvil"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"

	"cosmossdk.io/math"
)

func maybeFundERC20Solver(ctx context.Context, network netconf.ID, backends ethbackend.Backends) error {
	// funding solver with l1 wsETH uses anvil_setStorageAt util, which is only available on devnet
	if network != netconf.Devnet {
		return nil
	}

	// erc20 tokens to fund solver with on devnet. useful for solvernet development when forking public networks
	toFund := []struct {
		chainID uint64
		addr    common.Address
	}{
		// holesky wstETH
		{chainID: evmchain.IDMockL1, addr: common.HexToAddress("0x8d09a4502cc8cf1547ad300e066060d043f6982d")},
		// holesky stETH
		{chainID: evmchain.IDMockL1, addr: common.HexToAddress("0x3f1c547b21f65e10480de3ad8e19faac46c95034")},
	}

	solver := eoa.MustAddress(netconf.Devnet, eoa.RoleSolver)
	eth1m := math.NewInt(1_000_000).MulRaw(params.Ether).BigInt()

	for _, tkn := range toFund {
		ethCl, ok := backends.Clients()[tkn.chainID]
		if !ok {
			continue
		}

		// if devnet is not forking the public chain tkn is deployed on, this sets
		// storage on an unused address, which is fine
		err := anvil.FundERC20(ctx, ethCl, tkn.addr, eth1m, solver)
		if err != nil {
			return errors.Wrap(err, "fund tkn failed", "chain_id", tkn.chainID, "addr", tkn.addr)
		}
	}

	return nil
}

// setSolverAccountNativeBalance calls anvil_setBalance to set the solver account to a passed amount.
func setSolverAccountNativeBalance(ctx context.Context, backends ethbackend.Backends, amt *big.Int) error {
	solver := eoa.MustAddress(netconf.Devnet, eoa.RoleSolver)
	chainID := evmchain.IDMockL1

	ethCl, ok := backends.Clients()[chainID]
	if !ok {
		return errors.New("eth client not found", "chain_id", chainID)
	}

	if err := anvil.FundAccounts(ctx, ethCl, amt, solver); err != nil {
		return errors.Wrap(err, "set solver account balance failed")
	}

	balance, err := ethCl.BalanceAt(ctx, solver, nil)
	if err != nil {
		return errors.Wrap(err, "get account balance failed", "chain_id", chainID)
	}

	log.Debug(ctx, "Current Solver Balance", "balance", balance.String(), "chain_id", chainID)

	return nil
}
