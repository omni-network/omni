package solve

import (
	"context"
	"math/big"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/contracts"
	"github.com/omni-network/omni/lib/contracts/solvernet"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/umath"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"

	"golang.org/x/sync/errgroup"
)

const invalidChainID = 1234

var (
	zeroAddr            common.Address
	addrs               = mustAddrs(netconf.Devnet)
	invalidTokenAddress = common.HexToAddress("0x1234")
)

func mustAddrs(network netconf.ID) contracts.Addresses {
	addrs, err := contracts.GetAddresses(context.Background(), network)
	if err != nil {
		panic(err)
	}

	return addrs
}

func erc20Deposit(amt *big.Int, addr common.Address) solvernet.Deposit {
	return solvernet.Deposit{Token: addr, Amount: amt}
}

func nativeTransferCall(amt *big.Int, to common.Address) []solvernet.Call {
	return []solvernet.Call{{
		Value:  amt,
		Target: to,
		Data:   nil,
	}}
}

func nativeExpense(amt *big.Int) solvernet.Expenses {
	return []solvernet.Expense{{Amount: amt}}
}

func unsupportedExpense(amt *big.Int) solvernet.Expenses {
	return []solvernet.Expense{{Amount: amt, Token: invalidTokenAddress}}
}

func invalidExpense() solvernet.Expenses {
	return nativeExpense(big.NewInt(params.Ether))
}

func unsupportedDeposit(amt *big.Int) solvernet.Deposit {
	return solvernet.Deposit{Amount: amt, Token: invalidTokenAddress}
}

func nativeDeposit(amt *big.Int) solvernet.Deposit {
	return solvernet.Deposit{Amount: amt}
}

func mintAndApproveAll(ctx context.Context, backends ethbackend.Backends, orders []TestOrder) error {
	var eg errgroup.Group
	for _, order := range orders {
		eg.Go(func() error { return mintAndApprove(ctx, backends, order) })
	}

	if err := eg.Wait(); err != nil {
		return errors.Wrap(err, "wait group")
	}

	return nil
}

func mintAndApprove(ctx context.Context, backends ethbackend.Backends, order TestOrder) error {
	if order.IsDepositTokenEmpty() || order.IsDepositTokenInvalid() {
		// native, nothing to do
		return nil
	}

	backend, err := backends.Backend(order.SourceChainID)
	if err != nil {
		return errors.Wrap(err, "get backend")
	}

	addrs, err := contracts.GetAddresses(ctx, netconf.Devnet)
	if err != nil {
		return errors.Wrap(err, "get addrs")
	}

	txOpts, err := backend.BindOpts(ctx, order.Owner)
	if err != nil {
		return errors.Wrap(err, "bind opts")
	}

	contract, err := bindings.NewMockERC20(order.Deposit.Token, backend)
	if err != nil {
		return errors.Wrap(err, "bind contract")
	}

	tx, err := contract.Mint(txOpts, order.Owner, order.Deposit.Amount)
	if err != nil {
		return errors.Wrap(err, "mint tx")
	}

	_, err = backend.WaitMined(ctx, tx)
	if err != nil {
		return errors.Wrap(err, "wait mined")
	}

	tx, err = contract.Approve(txOpts, addrs.SolverNetInbox, umath.MaxUint256)
	if err != nil {
		return errors.Wrap(err, "mint tx")
	}

	_, err = backend.WaitMined(ctx, tx)
	if err != nil {
		return errors.Wrap(err, "wait mined")
	}

	return nil
}
