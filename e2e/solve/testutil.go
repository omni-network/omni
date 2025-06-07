package solve

import (
	"context"
	"math/big"
	"time"

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

// OrderLike defines common interface for order types that need minting and approval.
type OrderLike interface {
	GetSourceChainID() uint64
	GetDepositPayer() common.Address // The account that needs tokens minted and approved
	GetDeposit() solvernet.Deposit

	// Methods needed for testCheckAPI
	GetFillDeadline() time.Time
	GetDestChainID() uint64
	GetExpenses() []solvernet.Expense
	GetCalls() []solvernet.Call
	GetShouldReject() bool
	GetRejectReason() string
}

// Implement OrderLike for TestOrder.
func (o TestOrder) GetSourceChainID() uint64 {
	return o.SourceChainID
}

func (o TestOrder) GetDepositPayer() common.Address {
	return o.Owner // For regular orders, owner pays the deposit (owner is the sender in tests)
}

func (o TestOrder) GetDeposit() solvernet.Deposit {
	return o.Deposit
}

func (o TestOrder) GetFillDeadline() time.Time {
	return o.FillDeadline
}

func (o TestOrder) GetDestChainID() uint64 {
	return o.DestChainID
}

func (o TestOrder) GetExpenses() []solvernet.Expense {
	return o.Expenses
}

func (o TestOrder) GetCalls() []solvernet.Call {
	return o.Calls
}

func (o TestOrder) GetShouldReject() bool {
	return o.ShouldReject
}

func (o TestOrder) GetRejectReason() string {
	return o.RejectReason
}

// Implement OrderLike for TestGaslessOrder.
func (o TestGaslessOrder) GetSourceChainID() uint64 {
	return o.SourceChainID
}

func (o TestGaslessOrder) GetDepositPayer() common.Address {
	return o.User // For gasless orders, user pays the deposit via Permit2
}

func (o TestGaslessOrder) GetDeposit() solvernet.Deposit {
	return o.Deposit
}

func (o TestGaslessOrder) GetFillDeadline() time.Time {
	return o.FillDeadline
}

func (o TestGaslessOrder) GetDestChainID() uint64 {
	return o.DestChainID
}

func (o TestGaslessOrder) GetExpenses() []solvernet.Expense {
	return o.Expenses
}

func (o TestGaslessOrder) GetCalls() []solvernet.Call {
	return o.Calls
}

func (o TestGaslessOrder) GetShouldReject() bool {
	return o.ShouldReject
}

func (o TestGaslessOrder) GetRejectReason() string {
	return o.RejectReason
}

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
	return []solvernet.Expense{{Amount: amt, Token: addrs.Token}, {Amount: amt, Token: addrs.Token}}
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

func isDepositTokenEmptyForOrderLike(deposit solvernet.Deposit) bool {
	return deposit.Token == zeroAddr
}

func isDepositTokenInvalidForOrderLike(deposit solvernet.Deposit) bool {
	return deposit.Token == invalidTokenAddress
}

// Helper functions to convert specific order types to OrderLike slices.
func testOrdersToOrderLike(orders []TestOrder) []OrderLike {
	result := make([]OrderLike, len(orders))
	for i, order := range orders {
		result[i] = order
	}

	return result
}

func testGaslessOrdersToOrderLike(orders []TestGaslessOrder) []OrderLike {
	result := make([]OrderLike, len(orders))
	for i, order := range orders {
		result[i] = order
	}

	return result
}

func mintAndApproveAll(ctx context.Context, backends ethbackend.Backends, orders []OrderLike) error {
	var eg errgroup.Group
	for _, order := range orders {
		eg.Go(func() error {
			token, _ := tokens.ByAddress(order.GetSourceChainID(), order.GetDeposit().Token)

			if err := mintAndApprove(ctx, backends, order); err != nil {
				return errors.Wrap(err, "mint and approve",
					"chain", evmchain.Name(order.GetSourceChainID()),
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

func mintAndApprove(ctx context.Context, backends ethbackend.Backends, order OrderLike) error {
	deposit := order.GetDeposit()

	if isDepositTokenEmptyForOrderLike(deposit) || isDepositTokenInvalidForOrderLike(deposit) {
		// native, nothing to do
		return nil
	}

	backend, err := backends.Backend(order.GetSourceChainID())
	if err != nil {
		return errors.Wrap(err, "get backend")
	}

	addrs, err := contracts.GetAddresses(ctx, netconf.Devnet)
	if err != nil {
		return errors.Wrap(err, "get addrs")
	}

	depositPayer := order.GetDepositPayer()
	txOpts, err := backend.BindOpts(ctx, depositPayer)
	if err != nil {
		return errors.Wrap(err, "bind opts")
	}

	contract, err := bindings.NewMockERC20(deposit.Token, backend)
	if err != nil {
		return errors.Wrap(err, "bind contract")
	}

	tx, err := contract.Mint(txOpts, depositPayer, deposit.Amount)
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
	return solver.AddrAmt{Token: d.Token, Amount: d.Amount}
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

// Helper functions for OrderLike interface to support testCheckAPI logic.
func isInsufficientInventoryForOrderLike(order OrderLike) bool {
	return order.GetRejectReason() == solver.RejectInsufficientInventory.String()
}

// Helper function to convert OrderLike to solver.CheckRequest.
func orderLikeToCheckRequest(order OrderLike) solver.CheckRequest {
	return solver.CheckRequest{
		FillDeadline:       uint32(order.GetFillDeadline().Unix()),
		SourceChainID:      order.GetSourceChainID(),
		DestinationChainID: order.GetDestChainID(),
		Expenses:           expensesFromBindings(order.GetExpenses()),
		Calls:              callsFromBindings(order.GetCalls()),
		Deposit:            addrAmtFromDeposit(order.GetDeposit()),
	}
}
