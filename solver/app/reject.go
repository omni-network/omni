package app

import (
	"context"
	"fmt"
	"math/big"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/contracts/solvernet"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/solver/types"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

type rejectReason = types.RejectReason

const (
	rejectNone                  = types.RejectNone
	rejectDestCallReverts       = types.RejectDestCallReverts
	rejectInvalidDeposit        = types.RejectInvalidDeposit
	rejectInvalidExpense        = types.RejectInvalidExpense
	rejectInsufficientDeposit   = types.RejectInsufficientDeposit
	rejectInsufficientInventory = types.RejectInsufficientInventory
	rejectUnsupportedDeposit    = types.RejectUnsupportedDeposit
	rejectUnsupportedExpense    = types.RejectUnsupportedExpense
	rejectUnsupportedDestChain  = types.RejectUnsupportedDestChain
	rejectUnsupportedSrcChain   = types.RejectUnsupportedSrcChain
	rejectSameChain             = types.RejectSameChain
)

// RejectionError implement error, but represents a logical (expected) rejection, not an unexpected system error.
// We combine rejections with errors for detailed internal structured errors.
type RejectionError struct {
	Reason rejectReason // Succinct human-readable reason for rejection.
	Err    error        // Internal detailed reject condition
}

// Error implements error.
func (r *RejectionError) Error() string {
	return fmt.Sprintf("%s: %v", r.Reason.String(), r.Err)
}

// newRejection is a convenience function to create a new RejectionError error.
func newRejection(reason rejectReason, err error) *RejectionError {
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
	solverAddr, outboxAddr common.Address,
) func(ctx context.Context, order Order) (rejectReason, bool, error) {
	return func(ctx context.Context, order Order) (rejectReason, bool, error) {
		// Internal logic just return errors (convert them to rejections below)
		err := func(ctx context.Context, order Order) error {
			backend, err := backends.Backend(order.DestinationChainID)
			if err != nil {
				return newRejection(rejectUnsupportedDestChain, err)
			}

			deposits, err := parseMinReceived(order)
			if err != nil {
				return err
			}

			expenses, err := parseMaxSpent(order, outboxAddr)
			if err != nil {
				return err
			}

			if err := checkQuote(deposits, expenses); err != nil {
				return err
			}

			if err := checkLiquidity(ctx, expenses, backend, solverAddr); err != nil {
				return err
			}

			if err := checkApprovals(ctx, expenses, backend, solverAddr, outboxAddr); err != nil {
				return err
			}

			return checkFill(ctx, backend, order.ID, order.FillOriginData, nativeAmt(expenses), solverAddr, outboxAddr)
		}(ctx, order)

		if err == nil { // No error, no rejection
			return rejectNone, false, nil
		}

		r := new(RejectionError)
		if !errors.As(err, &r) { // Error, but no rejection
			return rejectNone, false, err
		}

		return r.Reason, true, nil
	}
}

// parseMinReceived parses order.MinReceived, checks all tokens are supported, returns the list of deposits.
func parseMinReceived(order Order) ([]Payment, error) {
	var deposits []Payment
	for _, output := range order.MinReceived {
		chainID := output.ChainId.Uint64()

		// inbox contract order resolution should ensure minReceived[].output.chainId matches order.SourceChainID
		if chainID != order.SourceChainID {
			return nil, errors.New("min received chain id mismatch [BUG]", "got", chainID, "expected", order.SourceChainID)
		}

		addr := toEthAddr(output.Token)
		if !cmpAddrs(addr, output.Token) {
			return nil, newRejection(rejectUnsupportedDeposit, errors.New("non-eth addressed token", "addr", hexutil.Encode(output.Token[:])))
		}

		tkn, ok := tokens.Find(chainID, addr)
		if !ok {
			return nil, newRejection(rejectUnsupportedDeposit, errors.New("unsupported token", "addr", addr))
		}

		deposits = append(deposits, Payment{
			Token:  tkn,
			Amount: output.Amount,
		})
	}

	return deposits, nil
}

// checkApprovals checks if the outbox is approved to spend all expenses.
func checkApprovals(ctx context.Context, expenses []Payment, client ethclient.Client, solverAddr, outboxAddr common.Address) error {
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
		Value: new(big.Int).Add(nativeValue, fee),
		Data:  fillCallData,
	}

	returnData, err := client.CallContract(ctx, msg, nil)
	if err != nil {
		return &RejectionError{
			Reason: rejectDestCallReverts,
			Err:    errors.Wrap(err, "return_data", hexutil.Encode(returnData), "custom", solvernet.DetectCustomError(err)),
		}
	}

	return nil
}

// parseMaxSpent parses order.MaxSpent, checks all tokens are supported, returns the list of expenses.
func parseMaxSpent(order Order, outboxAddr common.Address) ([]Payment, error) {
	var expenses []Payment

	hasNative := false

	for _, output := range order.MaxSpent {
		chainID := output.ChainId.Uint64()

		// order resolution ensures maxSpent[].output.chainId matches order.DestinationChainID
		if chainID != order.DestinationChainID {
			return nil, errors.New("max spent chain id mismatch [BUG]", "got", chainID, "expected", order.DestinationChainID)
		}

		// order resolve ensures maxSpent[].output.recipient is outboxAddr
		if toEthAddr(output.Recipient) != outboxAddr {
			return nil, errors.New("unexpected max spent recipient [BUG]", "got", output.Recipient, "expected", outboxAddr)
		}

		addr := toEthAddr(output.Token)
		if !cmpAddrs(addr, output.Token) {
			return nil, newRejection(rejectUnsupportedExpense, errors.New("non-eth addressed token", "addr", hexutil.Encode(output.Token[:])))
		}

		tkn, ok := tokens.Find(chainID, addr)
		if !ok {
			return nil, newRejection(rejectUnsupportedExpense, errors.New("unsupported token", "addr", addr))
		}

		if output.Token == [32]byte{} {
			if hasNative {
				// inbox contract enforces max 1 native expense
				return nil, errors.New("multiple native expenses [BUG]")
			}

			hasNative = true
		}

		expenses = append(expenses, Payment{
			Token:  tkn,
			Amount: output.Amount,
		})
	}

	return expenses, nil
}

func nativeAmt(ps []Payment) *big.Int {
	for _, p := range ps {
		if p.Token.IsNative() {
			return p.Amount
		}
	}

	return big.NewInt(0)
}

// checkQuote checks if deposits match or exceed quote for expenses.
// only single expense supported with matching deposit is supported.
func checkQuote(deposits, expenses []Payment) error {
	quote, err := getQuote(tkns(deposits), expenses)
	if err != nil {
		return err
	}

	return coversQuote(deposits, quote)
}

// checkLiquidity checks that the solver has enough liquidity to pay for the expenses.
func checkLiquidity(ctx context.Context, expenses []Payment, backend *ethbackend.Backend, solverAddr common.Address) error {
	for _, expense := range expenses {
		bal, err := balanceOf(ctx, expense.Token, backend, solverAddr)
		if err != nil {
			return errors.Wrap(err, "get balance", "token", expense.Token.Symbol)
		}

		// TODO: for native tokens, even if we have enough, we don't want to
		// spend out whole balance. we'll need to keep some for gas
		if bal.Cmp(expense.Amount) < 0 {
			return newRejection(rejectInsufficientInventory, errors.New("insufficient balance", "token", expense.Token.Symbol))
		}
	}

	return nil
}

func tkns(payments []Payment) []Token {
	tkns := make([]Token, len(payments))
	for i, p := range payments {
		tkns[i] = p.Token
	}

	return tkns
}
