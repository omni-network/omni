package app

import (
	"context"
	"encoding/json"
	"math/big"
	"net/http"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/contracts/solvernet"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/solver/types"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

type (
	CheckRequest  = types.CheckRequest
	CheckResponse = types.CheckResponse
	Expense       = solvernet.Expense
	Call          = solvernet.Call
	Deposit       = solvernet.Deposit

	checkFunc func(context.Context, CheckRequest) error
)

// newChecker returns a checkFunc that can be used to see if an order would be accepted or rejected.
// It is the logic behind the /check endpoint.
func newChecker(backends ethbackend.Backends, solverAddr, inboxAddr, outboxAddr common.Address) checkFunc {
	return func(ctx context.Context, req CheckRequest) error {
		if req.SourceChainID == req.DestinationChainID {
			return newRejection(rejectSameChain, errors.New("source and destination chain are the same"))
		}

		srcBackend, err := backends.Backend(req.SourceChainID)
		if err != nil {
			return newRejection(rejectUnsupportedSrcChain, errors.New("unsupported source chain", "chain_id", req.SourceChainID))
		}

		dstBackend, err := backends.Backend(req.DestinationChainID)
		if err != nil {
			return newRejection(rejectUnsupportedDestChain, errors.New("unsupported destination chain", "chain_id", req.DestinationChainID))
		}

		deposit, err := parseDeposit(req.SourceChainID, req.Deposit.Parse())
		if err != nil {
			return err
		}

		expenses, err := parseExpenses(req.DestinationChainID, req.Expenses.Parse(), req.Calls.Parse())
		if err != nil {
			return err
		}

		quote, err := getQuote([]Token{deposit.Token}, expenses)
		if err != nil {
			return err
		}

		err = coversQuote([]Payment{deposit}, quote)
		if err != nil {
			return err
		}

		if err := checkLiquidity(ctx, expenses, dstBackend, solverAddr); err != nil {
			return err
		}

		if err := checkApprovals(ctx, expenses, dstBackend, solverAddr, outboxAddr); err != nil {
			return err
		}

		orderID, err := getNextOrderID(ctx, srcBackend, inboxAddr)
		if err != nil {
			return err
		}

		fillOriginData, err := getFillOriginData(req)
		if err != nil {
			return err
		}

		return checkFill(ctx, dstBackend, orderID, fillOriginData, nativeAmt(expenses), solverAddr, outboxAddr)
	}
}

// getNextOrderID returns the next order ID for the given inbox.
func getNextOrderID(ctx context.Context, client ethclient.Client, inboxAddr common.Address) (OrderID, error) {
	inbox, err := bindings.NewSolverNetInbox(inboxAddr, client)
	if err != nil {
		return OrderID{}, errors.Wrap(err, "new inbox")
	}

	orderID, err := inbox.GetNextId(&bind.CallOpts{Context: ctx})
	if err != nil {
		return OrderID{}, errors.Wrap(err, "get next order id")
	}

	return orderID, nil
}

// getFillOriginData returns packed fill origin data for a check request.
func getFillOriginData(req CheckRequest) ([]byte, error) {
	fillOriginData := bindings.SolverNetFillOriginData{
		FillDeadline: req.FillDeadline,
		SrcChainId:   req.SourceChainID,
		DestChainId:  req.DestinationChainID,
		Expenses:     req.Expenses.Parse().NoNative(),
		Calls:        req.Calls.Parse().ToBindings(),
	}

	fillOriginDataBz, err := solvernet.PackFillOriginData(fillOriginData)
	if err != nil {
		return nil, errors.Wrap(err, "pack fill origin data")
	}

	return fillOriginDataBz, nil
}

// newCheckHandler returns a handler for the /check endpoint.
// It is responsible for http request / response handling, and delegates
// logic to a checkFunc.
func newCheckHandler(checkFunc checkFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, rr *http.Request) {
		ctx := rr.Context()

		w.Header().Set("Content-Type", "application/json")

		writeError := func(statusCode int, err error) {
			log.DebugErr(ctx, "Error handling /check request", err)

			writeJSON(ctx, w, CheckResponse{
				Error: &JSONErrorResponse{
					Code:    statusCode,
					Status:  http.StatusText(statusCode),
					Message: removeBUG(err.Error()),
				},
			})
		}

		var req CheckRequest
		if err := json.NewDecoder(rr.Body).Decode(&req); err != nil {
			writeError(http.StatusBadRequest, errors.Wrap(err, "decode request"))
			return
		}

		err := checkFunc(ctx, req)
		if r := new(RejectionError); errors.As(err, &r) { // RejectionError
			writeJSON(ctx, w, CheckResponse{
				Rejected:          true,
				RejectReason:      r.Reason.String(),
				RejectDescription: r.Err.Error(),
			})
		} else if err != nil { // Error
			writeError(http.StatusInternalServerError, err)
		} else {
			writeJSON(ctx, w, CheckResponse{Accepted: true}) // Success
		}
	})
}

// coversQuote checks if `deposits` match or exceed a `quote` for expenses.
func coversQuote(deposits, quote []Payment) error {
	byTkn := func(ps []Payment) map[Token]*big.Int {
		res := make(map[Token]*big.Int)
		for _, p := range ps {
			res[p.Token] = p.Amount
		}

		return res
	}

	quoteByTkn := byTkn(quote)
	depositsByTkn := byTkn(deposits)

	for tkn, q := range quoteByTkn {
		d, ok := depositsByTkn[tkn]
		if !ok {
			return newRejection(rejectInsufficientDeposit, errors.New("missing deposit", "token", tkn))
		}

		if d.Cmp(q) < 0 {
			return newRejection(rejectInsufficientDeposit, errors.New("insufficient deposit", "token", tkn, "deposit", d, "quote", q))
		}
	}

	return nil
}

func parseExpenses(destChainID uint64, expenses []Expense, calls []Call) ([]Payment, error) {
	var ps []Payment

	// sum of call value must be represented in expenses
	callValues := big.NewInt(0)
	for _, c := range calls {
		if c.Value == nil {
			continue
		}

		callValues.Add(callValues, c.Value)
	}

	nativeExpense := big.NewInt(0)
	for _, e := range expenses {
		if e.Amount.Sign() <= 0 {
			return nil, newRejection(rejectInvalidExpense, errors.New("expense amount positive"))
		}

		if isNative(e) {
			// sum of call values must be represented by a single expense
			if nativeExpense.Sign() > 0 {
				return nil, newRejection(rejectInvalidExpense, errors.New("only one native expense supported"))
			}

			nativeExpense = nativeExpense.Set(e.Amount)
		}

		tkn, ok := tokens.Find(destChainID, e.Token)
		if !ok {
			return nil, newRejection(rejectUnsupportedExpense, errors.New("unsupported expense token", "addr", e.Token))
		}

		if tkn.MaxSpend != nil && e.Amount.Cmp(tkn.MaxSpend) > 0 {
			return nil, newRejection(rejectExpenseOverMax, errors.New("expense over max", "token", tkn.Symbol, "max", tkn.MaxSpend, "amount", e.Amount))
		}

		if tkn.MinSpend != nil && e.Amount.Cmp(tkn.MinSpend) < 0 {
			return nil, newRejection(rejectExpenseUnderMin, errors.New("expense under min", "token", tkn.Symbol, "min", tkn.MinSpend, "amount", e.Amount))
		}

		ps = append(ps, Payment{
			Token:  tkn,
			Amount: e.Amount,
		})
	}

	// native expense must match sum of call values
	if nativeExpense.Cmp(callValues) != 0 {
		return nil, newRejection(rejectInvalidExpense, errors.New("native expense must match native value"))
	}

	return ps, nil
}

func parseDeposit(srcChainID uint64, dep Deposit) (Payment, error) {
	tkn, ok := tokens.Find(srcChainID, dep.Token)
	if !ok {
		return Payment{}, newRejection(rejectUnsupportedDeposit, errors.New("unsupported deposit token", "addr", dep.Token))
	}

	return Payment{
		Token:  tkn,
		Amount: dep.Amount,
	}, nil
}

func isNative(e Expense) bool { return e.Token == (common.Address{}) }
