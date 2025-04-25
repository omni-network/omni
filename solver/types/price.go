package types

import (
	"encoding/json"
	"fmt"
	"math/big"
	"strings"

	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/tokens"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

// Price represents the canonical solver price.
// It is 1 unit of deposit token denominated in expense tokens. A valid price is always > 0.
//
// Given a price of 800 OMNI/ETH, when depositing 1 ETH the expense amount is 800 OMNI.
//
//	expense amount = deposit amount * price
//	deposit amount = expense amount / price
type Price struct {
	Price   *big.Rat
	Deposit tokens.Asset
	Expense tokens.Asset
}

// WithFeeBips returns a new price with the fee in basis points subtracted from the price.
//
// Fees decrease the expense tokens received
// So it decreases the price (since price denominated in expense tokens).
// E.g. given a real price (pre-fees) of 800 OMNI/ETH,
// when depositing 1 ETH the expense amount (post-fees) is less than 800 OMNI.
func (p Price) WithFeeBips(fee int64) Price {
	clone := p
	if fee > 0 {
		feeRat := big.NewRat(10_000, 10_000+fee)
		clone.Price = new(big.Rat).Mul(clone.Price, feeRat)
	}

	return clone
}

func (p Price) FormatF64() string {
	s := new(big.Float).SetPrec(256).SetRat(p.Price).Text('f', 6)

	// Trim trailing zeroes and dot if needed
	s = strings.TrimRight(s, "0")
	s = strings.TrimRight(s, ".")

	return s
}

func (p Price) String() string {
	return fmt.Sprintf("%s %s/%s",
		p.FormatF64(),
		p.Expense.Symbol,
		p.Deposit.Symbol,
	)
}

func (p Price) Inverse() Price {
	return Price{
		Price:   new(big.Rat).Inv(p.Price),
		Deposit: p.Expense,
		Expense: p.Deposit,
	}
}

type priceJSON struct {
	Numerator   *hexutil.Big `json:"numerator"`
	Denominator *hexutil.Big `json:"denominator"`
	Rate        float64      `json:"imprecise_rate"`
	Deposit     string       `json:"deposit"`
	Expense     string       `json:"expense"`
}

// MarshalJSON implements the json.Marshaler interface.
func (p Price) MarshalJSON() ([]byte, error) {
	rate, _ := p.Price.Float64()
	priceData := priceJSON{
		Numerator:   (*hexutil.Big)(p.Price.Num()),
		Denominator: (*hexutil.Big)(p.Price.Denom()),
		Rate:        rate,
		Deposit:     p.Deposit.Symbol,
		Expense:     p.Expense.Symbol,
	}

	resp, err := json.Marshal(priceData)
	if err != nil {
		return nil, errors.Wrap(err, "marshal price")
	}

	return resp, nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (p *Price) UnmarshalJSON(data []byte) error {
	var priceData priceJSON
	if err := json.Unmarshal(data, &priceData); err != nil {
		return errors.Wrap(err, "unmarshal price")
	}

	if priceData.Deposit == "" {
		return errors.New("missing price deposit")
	} else if priceData.Expense == "" {
		return errors.New("missing price expense")
	} else if priceData.Numerator == nil {
		return errors.New("missing price numerator")
	} else if priceData.Denominator == nil {
		return errors.New("missing price denominator")
	} else if priceData.Rate == 0 {
		return errors.New("missing price rate")
	}

	rat := new(big.Rat)
	rat.SetFrac(priceData.Numerator.ToInt(), priceData.Denominator.ToInt())
	rate, _ := rat.Float64()

	if rat.Sign() <= 0 {
		return errors.New("price must be greater than zero")
	} else if rate != priceData.Rate {
		return errors.New("price rate does not match numerator/denominator")
	}

	deposit, err := tokens.AssetBySymbol(priceData.Deposit)
	if err != nil {
		return err
	}

	expense, err := tokens.AssetBySymbol(priceData.Expense)
	if err != nil {
		return err
	}

	*p = Price{
		Price:   rat,
		Deposit: deposit,
		Expense: expense,
	}

	return nil
}

// ToExpense converts a deposit amount to an expense amount using the price.
// It also rebases to the correct token decimals.
func (p Price) ToExpense(depositAmount *big.Int) *big.Int {
	rebasedDeposit := bi.Rebase(depositAmount, p.Deposit.Decimals, p.Expense.Decimals)

	return bi.MulRatFloor(rebasedDeposit, p.Price)
}

// ToDeposit converts an expense amount to a deposit amount using the price.
// It also rebases to the correct token decimals.
func (p Price) ToDeposit(expenseAmount *big.Int) *big.Int {
	rebasedExpense := bi.Rebase(expenseAmount, p.Expense.Decimals, p.Deposit.Decimals)

	return bi.MulRatCeil(rebasedExpense, p.Inverse().Price)
}
