package solve

import (
	"context"
	"math/big"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/contracts"
	"github.com/omni-network/omni/lib/contracts/solvernet"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/tokens"
	"github.com/omni-network/omni/lib/umath"
	"github.com/omni-network/omni/lib/uni"
	solver "github.com/omni-network/omni/solver/types"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"

	"golang.org/x/sync/errgroup"
)

const invalidChainID = 1234

var (
	zeroAddr            common.Address
	addrs               = mustAddrs(netconf.Devnet)
	invalidTokenAddress = common.HexToAddress("0x1234")
	invalidCallData     = hexutil.MustDecode("0x00000000")
	minETHSpend         = bi.Wei(1)
	maxETHSpend         = bi.Ether(1)
	validETHSpend       = bi.DivRaw( // mid = (min + max) / 2
		bi.Add(minETHSpend, maxETHSpend),
		2,
	)
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

func expenseAndCall(amt *big.Int, dstToken tokens.Token, user common.Address) ([]solvernet.Expense, []solvernet.Call) {
	if dstToken.IsNative() {
		return nativeExpense(amt), nativeTransferCall(amt, user)
	}

	expense := solvernet.Expense{
		Token:  dstToken.Address,
		Amount: amt,
	}

	call, err := erc20Call(dstToken, amt, user)
	if err != nil {
		panic(err)
	}

	return []solvernet.Expense{expense}, call
}

func erc20Call(dstToken tokens.Token, amt *big.Int, user common.Address) ([]solvernet.Call, error) {
	abi, err := bindings.IERC20MetaData.GetAbi()
	if err != nil {
		return nil, errors.Wrap(err, "get abi")
	}
	data, err := abi.Pack("transfer", user, amt)
	if err != nil {
		return nil, errors.Wrap(err, "pack transfer")
	}

	return []solvernet.Call{{
		Target: dstToken.Address,
		Data:   data,
	}}, nil
}

func nativeTransferCall(amt *big.Int, to common.Address) []solvernet.Call {
	return []solvernet.Call{{
		Value:  amt,
		Target: to,
		Data:   nil,
	}}
}

// contractCallWithInvalidCallData calls SolverNetInbox.sol contract with invalid calldata causing the tx to be reverted.
func contractCallWithInvalidCallData() []solvernet.Call {
	return []solvernet.Call{{
		Value:  validETHSpend,
		Target: addrs.SolverNetInbox,
		Data:   invalidCallData, // will revert
	}}
}

func nativeExpense(amt *big.Int) []solvernet.Expense {
	return []solvernet.Expense{{Amount: amt}}
}

func multipleNativeExpenses(amt *big.Int) []solvernet.Expense {
	return []solvernet.Expense{{Amount: amt}, {Amount: amt}}
}

func multipleERC20Expenses(amt *big.Int) []solvernet.Expense {
	return []solvernet.Expense{{Amount: amt, Token: addrs.NomToken}, {Amount: amt, Token: addrs.NomToken}}
}

func unsupportedExpense(amt *big.Int) []solvernet.Expense {
	return []solvernet.Expense{{Amount: amt, Token: invalidTokenAddress}}
}

func invalidExpenseOutOfBounds() []solvernet.Expense {
	return nativeExpense(bi.Ether(1))
}

func unsupportedERC20Deposit(amt *big.Int) solvernet.Deposit {
	return solvernet.Deposit{Amount: amt, Token: invalidTokenAddress}
}

func nativeDeposit(amt *big.Int) solvernet.Deposit {
	return solvernet.Deposit{Amount: amt}
}

func mintAndApproveAll(ctx context.Context, backends ethbackend.Backends, orders []TestOrder) error {
	var eg errgroup.Group
	for _, order := range orders {
		eg.Go(func() error {
			token, _ := tokens.ByAddress(order.SourceChainID, order.Deposit.Token)

			if err := mintAndApprove(ctx, backends, order); err != nil {
				return errors.Wrap(err, "mint and approve",
					"chain", evmchain.Name(order.SourceChainID),
					"token", token,
				)
			}

			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		return errors.Wrap(err, "wait group")
	}

	return nil
}

func mintAndApprove(ctx context.Context, backends ethbackend.Backends, order TestOrder) error {
	if isDepositTokenEmpty(order) || isDepositTokenInvalid(order) {
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

func addrAmtFromDeposit(d solvernet.Deposit) solver.AddrAmt {
	return solver.AddrAmt{Token: uni.EVMAddress(d.Token), Amount: d.Amount}
}

func callsFromBindings(calls []solvernet.Call) []solver.Call {
	var resp []solver.Call
	for _, c := range calls {
		resp = append(resp, solver.Call(c))
	}

	return resp
}

func expensesFromBindings(expenses []solvernet.Expense) []solver.Expense {
	var resp []solver.Expense
	for _, e := range expenses {
		resp = append(resp, solver.Expense(e))
	}

	return resp
}
