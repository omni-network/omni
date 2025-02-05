package app

import (
	"context"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/log"

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

type RejectOrErr struct {
	Reason rejectReason
	Err    error
}

// ShouldReject returns true if reject reason is not none.
func (r RejectOrErr) ShouldReject() bool {
	return r.Reason != rejectNone
}

// ShouldError returns true if there was an error without a rejection.
func (r RejectOrErr) ShouldError() bool {
	return r.Reason == rejectNone && r.Err != nil
}

// ShouldReturn returns true if there was a rejection or an error.
func (r RejectOrErr) ShouldReturn() bool {
	return r.Reason != rejectNone || r.Err != nil
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
		// rejectOrErr returns either an error, or a rejectReason with shouldReject == true.
		// For rejections, it swallows the error, only sampling / logging it.
		rejectOrErr := func(r RejectOrErr) (rejectReason, bool, error) {
			if !r.ShouldReturn() {
				return rejectNone, false, errors.New("[BUG] unexpected rejectOrErr call")
			}

			if r.ShouldError() {
				return rejectNone, false, r.Err
			}

			err := errors.Wrap(r.Err, "reject",
				"order_id", order.ID.String(),
				"dest_chain_id", order.DestinationChainID,
				"src_chain_id", order.SourceChainID,
				"target", targetName(order))

			rejectedOrders.WithLabelValues(
				chainName(order.SourceChainID),
				chainName(order.DestinationChainID),
				targetName(order),
				r.Reason.String(),
			).Inc()

			log.InfoErr(ctx, "Rejecting request", err, "reason", r.Reason)

			return r.Reason, true, nil
		}

		reject := func(reason rejectReason, err error) (rejectReason, bool, error) {
			return rejectOrErr(RejectOrErr{Reason: reason, Err: err})
		}

		returnErr := func(err error) (rejectReason, bool, error) {
			return rejectOrErr(RejectOrErr{Err: err})
		}

		if srcChainID != order.SourceChainID {
			return returnErr(errors.New("source chain id mismatch [BUG]", "got", order.SourceChainID, "expected", srcChainID))
		}

		backend, err := backends.Backend(order.DestinationChainID)
		if err != nil {
			return reject(rejectUnsupportedDestChain, err)
		}

		deposits, r := parseDeposits(order)
		if r.ShouldReturn() {
			return rejectOrErr(r)
		}

		expenses, r := parseExpenses(order)
		if r.ShouldReturn() {
			return rejectOrErr(r)
		}

		r = checkQuote(deposits, expenses)
		if r.ShouldReturn() {
			return rejectOrErr(r)
		}

		r = checkLiquidity(ctx, expenses, backend, solverAddr)
		if r.ShouldReturn() {
			return rejectOrErr(r)
		}

		return rejectNone, false, nil
	}
}

// parseDeposits parses order.MinReceived, checks all tokens are supported, returns the list of deposits.
func parseDeposits(order Order) ([]Payment, RejectOrErr) {
	var deposits []Payment
	for _, output := range order.MinReceived {
		chainID := output.ChainId.Uint64()

		// inbox contract order resolution should ensure minReceived[].output.chainId matches order.SourceChainID
		if chainID != order.SourceChainID {
			return nil, RejectOrErr{
				Err: errors.New("min received chain id mismatch [BUG]", "got", chainID, "expected", order.SourceChainID),
			}
		}

		addr := toEthAddr(output.Token)
		if !cmpAddrs(addr, output.Token) {
			return nil, RejectOrErr{
				Reason: rejectUnsupportedDeposit,
				Err:    errors.New("non-eth addressed token", "addr", hexutil.Encode(output.Token[:])),
			}
		}

		tkn, ok := tokens.find(chainID, addr)
		if !ok {
			return nil, RejectOrErr{
				Reason: rejectUnsupportedDeposit,
				Err:    errors.New("unsupported token", "addr", addr),
			}
		}

		deposits = append(deposits, Payment{
			Token:  tkn,
			Amount: output.Amount,
		})
	}

	return deposits, RejectOrErr{}
}

// parseExpenses parses order.MaxSpent, checks all tokens are supported, returns the list of expenses.
func parseExpenses(order Order) ([]Payment, RejectOrErr) {
	var expenses []Payment
	for _, output := range order.MaxSpent {
		chainID := output.ChainId.Uint64()

		// inbox contract order resolution should ensure maxSpent[].output.chainId matches order.DestinationChainID
		if chainID != order.DestinationChainID {
			return nil, RejectOrErr{
				Err: errors.New("max spent chain id mismatch [BUG]", "got", chainID, "expected", order.DestinationChainID),
			}
		}

		addr := toEthAddr(output.Token)
		if !cmpAddrs(addr, output.Token) {
			return nil, RejectOrErr{
				Reason: rejectUnsupportedExpense,
				Err:    errors.New("non-eth addressed token", "addr", hexutil.Encode(output.Token[:])),
			}
		}

		tkn, ok := tokens.find(chainID, addr)
		if !ok {
			return nil, RejectOrErr{
				Reason: rejectUnsupportedExpense,
				Err:    errors.New("unsupported token", "addr", addr),
			}
		}

		expenses = append(expenses, Payment{
			Token:  tkn,
			Amount: output.Amount,
		})
	}

	return expenses, RejectOrErr{}
}

// checkQuote checks if deposits match or exceed quote for expenses.
// only single expense supported with matching deposit is supported.
func checkQuote(deposits, expenses []Payment) RejectOrErr {
	quote, r := getQuote(tkns(deposits), expenses)
	if r.ShouldReturn() {
		return r
	}

	r = coversQuote(deposits, quote)
	if r.ShouldReturn() {
		return r
	}

	return RejectOrErr{}
}

// checkLiquidity checks that the solver has enough liquidity to pay for the expenses.
func checkLiquidity(ctx context.Context, expenses []Payment, backend *ethbackend.Backend, solverAddr common.Address) RejectOrErr {
	for _, expense := range expenses {
		bal, err := balanceOf(ctx, expense.Token, backend, solverAddr)
		if err != nil {
			return RejectOrErr{
				Err: errors.Wrap(err, "get balance", "token", expense.Token.Symbol),
			}
		}

		// TODO: for native tokens, even if we have enough, we don't want to
		// spend out whole balance. we'll need to keep some for gas
		if bal.Cmp(expense.Amount) < 0 {
			return RejectOrErr{
				Reason: rejectInsufficientInventory,
				Err:    errors.New("insufficient balance", "token", expense.Token.Symbol),
			}
		}
	}

	return RejectOrErr{}
}

func tkns(payments []Payment) []Token {
	tkns := make([]Token, len(payments))
	for i, p := range payments {
		tkns[i] = p.Token
	}

	return tkns
}
