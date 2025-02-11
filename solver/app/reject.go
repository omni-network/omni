package app

import (
	"context"
	"fmt"
	"math/big"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/contracts/solvernet"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/log"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

//go:generate stringer -type=rejectReason -trimprefix=reject
type rejectReason uint8

const (
	rejectNone                  rejectReason = 0
	rejectDestCallReverts       rejectReason = 1
	rejectInvalidDeposit        rejectReason = 2
	rejectInvalidExpense        rejectReason = 3
	rejectInsufficientDeposit   rejectReason = 4
	rejectInsufficientInventory rejectReason = 5
	rejectUnsupportedDeposit    rejectReason = 6
	rejectUnsupportedExpense    rejectReason = 7
	rejectUnsupportedDestChain  rejectReason = 8
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

type AlreadyFilledError struct {
	OrderID OrderID
}

// Error implements error.
func (e *AlreadyFilledError) Error() string {
	return "already filled: " + e.OrderID.String()
}

// newShouldRejector returns as ShouldReject function for the given network.
//
// ShouldReject returns true and a reason if the request should be rejected.
// It returns false if the request should be accepted.
// Errors are unexpected and refer to internal problems.
func newShouldRejector(
	backends ethbackend.Backends,
	solverAddr common.Address,
	targetName func(Order) string,
	chainName func(uint64) string,
) func(ctx context.Context, srcChainID uint64, order Order) (rejectReason, bool, error) {
	return func(ctx context.Context, srcChainID uint64, order Order) (rejectReason, bool, error) {
		// Internal logic just return errors (convert them to rejections below)
		err := func(ctx context.Context, srcChainID uint64, order Order) error {
			if srcChainID != order.SourceChainID {
				return errors.New("source chain id mismatch [BUG]", "got", order.SourceChainID, "expected", srcChainID)
			}

			backend, err := backends.Backend(order.DestinationChainID)
			if err != nil {
				return newRejection(rejectUnsupportedDestChain, err)
			}

			deposits, err := parseDeposits(order)
			if err != nil {
				return err
			}

			expenses, err := parseExpenses(order)
			if err != nil {
				return err
			}

			if err := checkQuote(deposits, expenses); err != nil {
				return err
			}

			if err := checkLiquidity(ctx, expenses, backend, solverAddr); err != nil {
				return err
			}

			return checkFill(ctx, backend, order, solverAddr)
		}(ctx, srcChainID, order)

		if err == nil { // No error, no rejection
			return rejectNone, false, nil
		}

		r := new(RejectionError)
		if !errors.As(err, &r) { // Error, but no rejection
			return rejectNone, false, err
		}

		// TODO:  handle upstream, do not accept if already filled
		e := new(AlreadyFilledError)
		if errors.As(err, &e) { // Already filled, no rejection
			log.InfoErr(ctx, "Already filled", err, "order_id", order.ID.String())
			return rejectNone, false, nil
		}

		// Handle rejection
		rejectedOrders.WithLabelValues(
			chainName(order.SourceChainID),
			chainName(order.DestinationChainID),
			targetName(order),
			r.Reason.String(),
		).Inc()

		err = errors.Wrap(r.Err, "reject",
			"reason", r.Reason.String(),
			"order_id", order.ID.String(),
			"dest_chain_id", order.DestinationChainID,
			"src_chain_id", order.SourceChainID,
			"target", targetName(order))

		log.InfoErr(ctx, "Rejecting request", err)

		return r.Reason, true, nil
	}
}

// parseDeposits parses order.MinReceived, checks all tokens are supported, returns the list of deposits.
func parseDeposits(order Order) ([]Payment, error) {
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

// checkFill checks if a destination call reverts.
// TODO: approve outbox spend (will revert without approvals).
func checkFill(ctx context.Context, backend *ethbackend.Backend, order Order, solverAddr common.Address) error {
	outbox, err := bindings.NewSolverNetOutbox(order.DestinationSettler, backend)
	if err != nil {
		return errors.Wrap(err, "new outbox")
	}

	callOpts := &bind.CallOpts{Context: ctx}
	if ok, err := outbox.DidFill(callOpts, order.ID, order.FillOriginData); err != nil {
		return errors.Wrap(err, "did fill")
	} else if ok {
		return &AlreadyFilledError{OrderID: order.ID}
	}

	nativeValue := big.NewInt(0)
	for _, output := range order.MaxSpent {
		if output.ChainId.Uint64() != order.DestinationChainID {
			// We error on this case for now, as our contracts only allow single dest chain orders
			// ERC7683 allows for orders with multiple destination chains, so continue-ing here
			// would also be appropriate.
			return errors.New("[BUG] destination chain mismatch")
		}

		// zero token address means native token
		if output.Token == [32]byte{} {
			nativeValue.Add(nativeValue, output.Amount)
			continue
		}

		// approve outbox spend before checking for reverts, because lack of allowance will cause revert
		if err := approveOutboxSpend(ctx, output, backend, solverAddr, order.DestinationSettler); err != nil {
			return errors.Wrap(err, "approve outbox spend")
		}
	}

	// xcall fee
	fee, err := outbox.FillFee(callOpts, order.FillOriginData)
	if err != nil {
		return errors.Wrap(err, "get fulfill fee")
	}

	fillerData := []byte{} // fillerData is optional ERC7683 custom filler specific data, unused in our contracts
	fillCallData, err := outboxABI.Pack("fill", order.ID, order.FillOriginData, fillerData)
	if err != nil {
		return errors.Wrap(err, "pack fill inputs")
	}

	msg := ethereum.CallMsg{
		To:    &order.DestinationSettler,
		From:  solverAddr,
		Value: new(big.Int).Add(nativeValue, fee),
		Data:  fillCallData,
	}

	returnData, err := backend.CallContract(ctx, msg, nil)
	if err != nil {
		return &RejectionError{
			Reason: rejectDestCallReverts,
			Err:    errors.Wrap(err, "return_data", hexutil.Encode(returnData), "custom", solvernet.DetectCustomError(err)),
		}
	}

	return nil
}

// parseExpenses parses order.MaxSpent, checks all tokens are supported, returns the list of expenses.
func parseExpenses(order Order) ([]Payment, error) {
	var expenses []Payment
	for _, output := range order.MaxSpent {
		chainID := output.ChainId.Uint64()

		// inbox contract order resolution should ensure maxSpent[].output.chainId matches order.DestinationChainID
		if chainID != order.DestinationChainID {
			return nil, errors.New("max spent chain id mismatch [BUG]", "got", chainID, "expected", order.DestinationChainID)
		}

		addr := toEthAddr(output.Token)
		if !cmpAddrs(addr, output.Token) {
			return nil, newRejection(rejectUnsupportedExpense, errors.New("non-eth addressed token", "addr", hexutil.Encode(output.Token[:])))
		}

		tkn, ok := tokens.Find(chainID, addr)
		if !ok {
			return nil, newRejection(rejectUnsupportedExpense, errors.New("unsupported token", "addr", addr))
		}

		expenses = append(expenses, Payment{
			Token:  tkn,
			Amount: output.Amount,
		})
	}

	return expenses, nil
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
