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

		deposit, err := parseDeposit(req.SourceChainID, req.Deposit)
		if err != nil {
			return err
		}

		expenses, err := parseExpenses(req.DestinationChainID, req.Expenses, req.Calls)
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

// getFillOriginData returns packed fill origin data for a quote request.
func getFillOriginData(req CheckRequest) ([]byte, error) {
	fillOriginData := bindings.SolverNetFillOriginData{
		FillDeadline: req.FillDeadline,
		SrcChainId:   req.SourceChainID,
		DestChainId:  req.DestinationChainID,
		Expenses:     req.Expenses.ToBindings(),
		Calls:        req.Calls.ToBindings(),
	}

	fillOriginDataBz, err := solvernet.PackFillOriginData(fillOriginData)
	if err != nil {
		return nil, errors.Wrap(err, "pack fill origin data")
	}

	return fillOriginDataBz, nil
}

// newCheckHandler returns a handler for the /quote endpoint.
// It is responsible to http request / response handling, and delegates
// logic to a quoteFunc.
func newCheckHandler(checkFunc checkFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, rr *http.Request) {
		ctx := rr.Context()

		w.Header().Set("Content-Type", "application/json")

		writeError := func(statusCode int, err error) {
			log.DebugErr(ctx, "Error handling /quote request", err)

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

// getQuote returns payment in `depositTkns` required to pay for `expenses`.
//
// For now, this is a simple quote that requires a single expense, paid
// for by an equal amount of an equivalent deposit token. Token equivalence is
// determined by symbol (ex arbitrum "ETH" is equivalent to optimism "ETH").
func getQuote(depositTkns []Token, expenses []Payment) ([]Payment, error) {
	if len(depositTkns) != 1 {
		return nil, newRejection(rejectInvalidDeposit, errors.New("only single deposit token supported"))
	}

	if len(expenses) != 1 {
		return nil, newRejection(rejectInvalidExpense, errors.New("only single expense supported"))
	}

	expense := expenses[0]
	depositTkn := depositTkns[0]

	if expense.Token.Symbol != depositTkn.Symbol {
		return nil, newRejection(rejectInvalidDeposit, errors.New("deposit token must match expense token"))
	}

	// make sure chain class (e.g. mainnet, testnet) matches
	// we should reject with UnsupportedDestChain before this. the solver is
	// initialized by network, which only includes chains of the same class
	if expense.Token.ChainClass != depositTkn.ChainClass {
		return nil, newRejection(rejectInvalidDeposit, errors.New("deposit and expense must be of the same chain class (e.g. mainnet, testnet)"))
	}

	return []Payment{
		{
			Token:  depositTkn,
			Amount: expense.Amount,
		},
	}, nil
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
