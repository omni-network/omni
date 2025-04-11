package app

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/contracts/solvernet"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/tokens"
	"github.com/omni-network/omni/lib/tokens/tokenutil"
	"github.com/omni-network/omni/solver/types"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// RejectionError implement error, but represents a logical (expected) rejection, not an unexpected system error.
// We combine rejections with errors for detailed internal structured errors.
type RejectionError struct {
	Reason types.RejectReason // Succinct human-readable reason for rejection.
	Err    error              // Internal detailed reject condition
}

// Error implements error.
func (r *RejectionError) Error() string {
	// TODO(corver): Improve how to add errors attributes, instead of using hacky empty wraps.
	errMsg := errors.Format(r.Err)
	for strings.HasPrefix(errMsg, ": ") {
		errMsg = errMsg[2:]
	}

	return fmt.Sprintf("%s: %v", r.Reason.String(), errMsg)
}

// newRejection is a convenience function to create a new RejectionError error.
func newRejection(reason types.RejectReason, err error) *RejectionError {
	return &RejectionError{Reason: reason, Err: err}
}

// newShouldRejector returns as ShouldReject function for the given network.
//
// ShouldReject returns true and a reason if the request should be rejected.
// It returns false if the request should be accepted.
// Errors are unexpected and refer to internal problems.
//
// It will return true if the order has rleady been filled. DidFill check
// should be made before calling ShouldReject.
func newShouldRejector(
	backends ethbackend.Backends,
	isAllowedCall callAllowFunc,
	priceFunc priceFunc,
	solverAddr, outboxAddr common.Address,
) func(ctx context.Context, order Order) (types.RejectReason, bool, error) {
	return func(ctx context.Context, order Order) (types.RejectReason, bool, error) {
		pendingData, err := order.PendingData()
		if err != nil {
			return types.RejectNone, false, err
		}

		// Internal logic just return errors (convert them to rejections below)
		err = func(ctx context.Context, order Order) error {
			backend, err := backends.Backend(pendingData.DestinationChainID)
			if err != nil {
				return newRejection(types.RejectUnsupportedDestChain, err)
			}

			if err := checkOrderCalls(pendingData, isAllowedCall); err != nil {
				return err
			}

			deposits, err := parseMinReceived(order)
			if err != nil {
				return err
			}

			expenses, err := parseMaxSpent(pendingData, outboxAddr)
			if err != nil {
				return err
			}

			if err := checkQuote(ctx, priceFunc, deposits, expenses); err != nil {
				return err
			}

			if err := checkLiquidity(ctx, expenses, backend, solverAddr); err != nil {
				return err
			}

			if err := checkApprovals(ctx, expenses, backend, solverAddr, outboxAddr); err != nil {
				return err
			}

			return checkFill(ctx, backend, order.ID, pendingData.FillOriginData, nativeAmt(expenses), solverAddr, outboxAddr)
		}(ctx, order)

		if err == nil { // No error, no rejection
			return types.RejectNone, false, nil
		}

		r := new(RejectionError)
		if !errors.As(err, &r) { // Error, but no rejection
			return types.RejectNone, false, err
		}

		return r.Reason, true, nil
	}
}

// parseMinReceived parses order.MinReceived, checks all tokens are supported, returns the list of deposits.
func parseMinReceived(order Order) ([]TokenAmt, error) {
	minReceived, err := order.MinReceived()
	if err != nil {
		return nil, err
	}

	var deposits []TokenAmt
	for _, output := range minReceived {
		chainID := output.ChainId.Uint64()

		// inbox contract order resolution should ensure minReceived[].output.chainId matches order.SourceChainID
		if chainID != order.SourceChainID {
			return nil, errors.New("min received chain id mismatch [BUG]", "got", chainID, "expected", order.SourceChainID)
		}

		addr := toEthAddr(output.Token)
		if !cmpAddrs(addr, output.Token) {
			return nil, newRejection(types.RejectUnsupportedDeposit, errors.New("non-eth addressed token", "addr", hexutil.Encode(output.Token[:])))
		}

		tkn, ok := tokens.ByAddress(chainID, addr)
		if !ok || !IsSupportedToken(tkn) {
			return nil, newRejection(types.RejectUnsupportedDeposit, errors.New("unsupported token", "addr", addr))
		}

		deposits = append(deposits, TokenAmt{
			Token:  tkn,
			Amount: output.Amount,
		})
	}

	return deposits, nil
}

// checkApprovals checks if the outbox is approved to spend all expenses.
func checkApprovals(ctx context.Context, expenses []TokenAmt, client ethclient.Client, solverAddr, outboxAddr common.Address) error {
	for _, expense := range expenses {
		tkn := expense.Token

		if tkn.IsNative() {
			continue
		}

		isAppproved, err := isAppproved(ctx, tkn.Address, client, solverAddr, outboxAddr, expense.Amount)
		if err != nil {
			return errors.Wrap(err, "is approved")
		}

		if !isAppproved {
			return errors.New("outbox not approved to spend token",
				"token", tkn.Symbol,
				"chain_id", tkn.ChainID,
				"addr", tkn.Address.Hex(),
				"amount", expense.Amount,
			)
		}
	}

	return nil
}

// checkFill checks if a destination call reverts. Does not check if order was already filled.
func checkFill(
	ctx context.Context,
	client ethclient.Client,
	orderID OrderID,
	fillOriginData []byte,
	nativeValue *big.Int,
	solverAddr, outboxAddr common.Address,
) error {
	outbox, err := bindings.NewSolverNetOutbox(outboxAddr, client)
	if err != nil {
		return errors.Wrap(err, "new outbox")
	}

	// xcall fee
	fee, err := outbox.FillFee(&bind.CallOpts{Context: ctx}, fillOriginData)
	if err != nil {
		return errors.Wrap(err, "get fulfill fee")
	}

	fillCallData, err := solvernet.PackFillCalldata(orderID, fillOriginData)
	if err != nil {
		return errors.Wrap(err, "pack fill inputs")
	}

	msg := ethereum.CallMsg{
		To:    &outboxAddr,
		From:  solverAddr,
		Value: bi.Add(nativeValue, fee),
		Data:  fillCallData,
	}

	returnData, err := client.CallContract(ctx, msg, nil)
	if err != nil {
		return &RejectionError{
			Reason: types.RejectDestCallReverts,
			Err:    errors.Wrap(err, "call contract", "return_data", hexutil.Encode(returnData), "solidity_err", solvernet.DetectCustomError(err)),
		}
	}

	return nil
}

// parseMaxSpent parses order.MaxSpent, checks all tokens are supported, returns the list of expenses.
func parseMaxSpent(pendingData PendingData, outboxAddr common.Address) ([]TokenAmt, error) {
	var expenses []TokenAmt
	var hasNative bool
	for _, output := range pendingData.MaxSpent {
		chainID := output.ChainId.Uint64()

		// order resolution ensures maxSpent[].output.chainId matches order.DestinationChainID
		if chainID != pendingData.DestinationChainID {
			return nil, errors.New("max spent chain id mismatch [BUG]", "got", chainID, "expected", pendingData.DestinationChainID)
		}

		// order resolve ensures maxSpent[].output.recipient is outboxAddr
		if toEthAddr(output.Recipient) != outboxAddr {
			return nil, errors.New("unexpected max spent recipient [BUG]", "got", output.Recipient, "expected", outboxAddr)
		}

		addr := toEthAddr(output.Token)
		if !cmpAddrs(addr, output.Token) {
			return nil, newRejection(types.RejectUnsupportedExpense, errors.New("non-eth addressed token", "addr", hexutil.Encode(output.Token[:])))
		}

		tkn, ok := tokens.ByAddress(chainID, addr)
		if !ok || !IsSupportedToken(tkn) {
			return nil, newRejection(types.RejectUnsupportedExpense, errors.New("unsupported token", "addr", addr))
		}

		if output.Token == [32]byte{} {
			if hasNative {
				// inbox contract enforces max 1 native expense
				return nil, errors.New("multiple native expenses [BUG]")
			}

			hasNative = true
		}

		bounds := GetSpendBounds(tkn)
		if bounds.MaxSpend != nil && bi.GT(output.Amount, bounds.MaxSpend) {
			return nil, newRejection(types.RejectExpenseOverMax, errors.New("expense over max", "token", tkn.Symbol, "max", bounds.MaxSpend, "amount", output.Amount))
		}

		if bounds.MinSpend != nil && bi.LT(output.Amount, bounds.MinSpend) {
			return nil, newRejection(types.RejectExpenseUnderMin, errors.New("expense under min", "token", tkn.Symbol, "min", bounds.MinSpend, "amount", output.Amount))
		}

		expenses = append(expenses, TokenAmt{
			Token:  tkn,
			Amount: output.Amount,
		})
	}

	return expenses, nil
}

func nativeAmt(ps []TokenAmt) *big.Int {
	for _, p := range ps {
		if p.Token.IsNative() {
			return p.Amount
		}
	}

	return bi.Zero()
}

// checkQuote checks if deposits match or exceed quote for expenses.
// only single expense supported with matching deposit is supported.
func checkQuote(ctx context.Context, priceFunc priceFunc, deposits, expenses []TokenAmt) error {
	quote, err := getQuote(ctx, priceFunc, tkns(deposits), expenses)
	if err != nil {
		return err
	}

	return coversQuote(deposits, quote)
}

// checkLiquidity checks that the solver has enough liquidity to pay for the expenses.
func checkLiquidity(ctx context.Context, expenses []TokenAmt, backend *ethbackend.Backend, solverAddr common.Address) error {
	for _, expense := range expenses {
		bal, err := tokenutil.BalanceOf(ctx, backend, expense.Token, solverAddr)
		if err != nil {
			return errors.Wrap(err, "get balance", "token", expense.Token.Symbol)
		}

		// TODO: for native tokens, even if we have enough, we don't want to
		// spend out whole balance. we'll need to keep some for gas
		if bi.LT(bal, expense.Amount) {
			if expense.Token.IsOMNI() && expense.Token.IsNative() {
				// for native OMNI, instead of rejecting, we error. the solver
				// will retry this order. this is a special case "fix" to handle
				// potentially high load / order size for genesis stake upgrades.
				// solver rebalances OMNI manually, so should recover
				return errors.New("insufficient balance, will retry",
					"balance", expense.Token.FormatAmt(bal),
					"expense", expense,
				)
			}

			return newRejection(types.RejectInsufficientInventory, errors.New("insufficient balance",
				"balance", expense.Token.FormatAmt(bal),
				"expense", expense,
			))
		}
	}

	return nil
}

// checkOrderCalls checks if all calls in an order are allowed.
func checkOrderCalls(pendingData PendingData, isAllowed callAllowFunc) error {
	fill, err := pendingData.ParsedFillOriginData()
	if err != nil {
		return errors.Wrap(err, "parse fill origin data")
	}

	var calls []types.Call
	for _, call := range fill.Calls {
		calls = append(calls, types.Call{
			Target: call.Target,
			Value:  call.Value,
			Data:   append(call.Selector[:], call.Params...),
		})
	}

	return checkCalls(pendingData.DestinationChainID, calls, isAllowed)
}

// checkCalls checks if all calls to destChainID are allowed.
func checkCalls(destChainID uint64, calls []types.Call, isAllowed callAllowFunc) error {
	for _, call := range calls {
		if !isAllowed(destChainID, call.Target, call.Data) {
			return newRejection(types.RejectCallNotAllowed, errors.New("call not allowed", "target", call.Target.Hex(), "data", hexutil.Encode(call.Data)))
		}
	}

	return nil
}

func tkns(payments []TokenAmt) []tokens.Token {
	tkns := make([]tokens.Token, len(payments))
	for i, p := range payments {
		tkns[i] = p.Token
	}

	return tkns
}
