package solve

import (
	"context"
	"math/big"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/lib/anvil"
	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/tokens"

	"github.com/ethereum/go-ethereum/common"
)

func maybeFundERC20Solver(ctx context.Context, network netconf.ID, backends ethbackend.Backends) error {
	// funding solver with l1 wsETH uses anvil_setStorageAt util, which is only available on devnet
	if network != netconf.Devnet {
		return nil
	}

	wstETHOnMockL1, err := Find(evmchain.IDMockL1, tokens.WSTETH.Symbol)
	if err != nil {
		return err
	}
	wstETHOnMockL2, err := Find(evmchain.IDMockL2, tokens.WSTETH.Symbol)
	if err != nil {
		return err
	}
	usdcOnMockL1, err := Find(evmchain.IDMockL1, tokens.USDC.Symbol)
	if err != nil {
		return err
	}
	usdcOnMockL2, err := Find(evmchain.IDMockL2, tokens.USDC.Symbol)
	if err != nil {
		return err
	}

	// erc20 tokens to fund solver with on devnet. useful for solvernet development when forking public networks
	toFund := []struct {
		chainID uint64
		addr    common.Address
	}{
		// devnet wstETH
		{chainID: evmchain.IDMockL1, addr: wstETHOnMockL1},
		{chainID: evmchain.IDMockL2, addr: wstETHOnMockL2},
		// devnet USDC
		{chainID: evmchain.IDMockL1, addr: usdcOnMockL1},
		{chainID: evmchain.IDMockL2, addr: usdcOnMockL2},
	}

	solver := eoa.MustAddress(netconf.Devnet, eoa.RoleSolver)
	eth1m := bi.Ether(1_000_000)

	for _, tkn := range toFund {
		backend, err := backends.Backend(tkn.chainID)
		if err != nil {
			return err
		}

		merc20, err := bindings.NewMockERC20(tkn.addr, backend)
		if err != nil {
			return errors.Wrap(err, "new mock erc20")
		}

		txOpts, err := backend.BindOpts(ctx, solver)
		if err != nil {
			return err
		}

		tx, err := merc20.Mint(txOpts, solver, eth1m)
		if err != nil {
			return errors.Wrap(err, "mint token")
		} else if _, err := backend.WaitMined(ctx, tx); err != nil {
			return errors.Wrap(err, "wait for mint tx")
		}
	}

	return nil
}

func maybeFundERC20Flowgen(ctx context.Context, network netconf.ID, backends ethbackend.Backends) error {
	if network != netconf.Devnet {
		return nil
	}

	wstETHOnMockL1, err := Find(evmchain.IDMockL1, tokens.WSTETH.Symbol)
	if err != nil {
		return err
	}

	toFund := []struct {
		chainID uint64
		addr    common.Address
	}{
		// devnet wstETH
		{chainID: evmchain.IDMockL1, addr: wstETHOnMockL1},
	}

	flowgen := eoa.MustAddress(netconf.Devnet, eoa.RoleFlowgen)
	eth1m := bi.Ether(1_000_000)

	for _, tkn := range toFund {
		backend, err := backends.Backend(tkn.chainID)
		if err != nil {
			return err
		}

		merc20, err := bindings.NewMockERC20(tkn.addr, backend)
		if err != nil {
			return errors.Wrap(err, "new mock erc20")
		}

		txOpts, err := backend.BindOpts(ctx, flowgen)
		if err != nil {
			return err
		}

		tx, err := merc20.Mint(txOpts, flowgen, eth1m)
		if err != nil {
			return errors.Wrap(err, "mint token")
		} else if _, err := backend.WaitMined(ctx, tx); err != nil {
			return errors.Wrap(err, "wait for mint tx")
		}

		log.Debug(ctx, "Funded flowgen", "chain", tkn.chainID, "address", flowgen, "token_address", tkn.addr)
	}

	return nil
}

// setSolverAccountNativeBalance calls anvil_setBalance to set the solver account native balance to a certain amount.
func setSolverAccountNativeBalance(ctx context.Context, chainID uint64, backends ethbackend.Backends, amt *big.Int) error {
	solver := eoa.MustAddress(netconf.Devnet, eoa.RoleSolver)

	ethCl, ok := backends.Clients()[chainID]
	if !ok {
		return errors.New("eth client not found", "chain_id", chainID)
	}

	if err := anvil.FundAccounts(ctx, ethCl, amt, solver); err != nil {
		return errors.Wrap(err, "set solver account balance failed")
	}

	return nil
}
