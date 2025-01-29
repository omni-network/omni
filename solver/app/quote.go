package app

import (
	"math/big"

	"github.com/omni-network/omni/lib/errors"
)

// getQuote returns payment in `depositTkns` required to pay for `expenses`.
//
// For now, this is a simple quote that requires a single expense, paid
// for by an equal amount of an equivalent deposit token. Token equivalence is
// determined by symbol (ex arbitrum "ETH" is equivalent to optimism "ETH").
func getQuote(depositTkns []Token, expenses []Payment) ([]Payment, RejectOrErr) {
	if len(depositTkns) != 1 {
		return nil, RejectOrErr{
			Reason: rejectInvalidDeposit,
			Err:    errors.New("only single deposit token supported"),
		}
	}

	if len(expenses) != 1 {
		return nil, RejectOrErr{
			Reason: rejectInvalidExpense,
			Err:    errors.New("only single expense supported"),
		}
	}

	expense := expenses[0]
	depositTkn := depositTkns[0]

	if expense.Token.Symbol != depositTkn.Symbol {
		return nil, RejectOrErr{
			Reason: rejectInvalidDeposit,
			Err:    errors.New("deposit token must match expense token"),
		}
	}

	// make sure chain class (e.g. mainnet, testnet) matches
	// we should reject with UnsupportedDestChain before this. the solver is
	// initialized by network, which only includes chains of the same class
	if expense.Token.ChainClass != depositTkn.ChainClass {
		return nil, RejectOrErr{
			Reason: rejectInvalidDeposit,
			Err:    errors.New("deposit and expense must be of the same chain class (e.g. mainnet, testnet)"),
		}
	}

	return []Payment{
		{
			Token:  depositTkn,
			Amount: expense.Amount,
		},
	}, RejectOrErr{}
}

// coversQuote checks if `deposits` match or exceed a `quote` for expenses.
func coversQuote(deposits, quote []Payment) RejectOrErr {
	if len(quote) != len(deposits) {
		return RejectOrErr{}
	}

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
			return RejectOrErr{
				Reason: rejectInsufficientDeposit,
				Err:    errors.New("missing deposit", "token", tkn),
			}
		}

		if d.Cmp(q) < 0 {
			return RejectOrErr{
				Reason: rejectInsufficientDeposit,
				Err:    errors.New("insufficient deposit", "token", tkn, "deposit", d, "quote", q),
			}
		}
	}

	return RejectOrErr{}
}
