package types_test

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/tokens"
	"github.com/omni-network/omni/solver/types"

	"github.com/stretchr/testify/require"
)

func TestPriceBips(t *testing.T) {
	t.Parallel()

	// Unit price with fees
	price := types.Price{
		Price:   big.NewRat(1, 1),
		Deposit: tokens.ETH,
		Expense: tokens.ETH,
	}.WithFeeBips(30) // 0.3% fee

	deposit := bi.Ether(1)
	expense := price.ToExpense(deposit)
	deposit2 := price.ToDeposit(expense)
	require.Equal(t, deposit, deposit2)
}

func TestPriceConvert(t *testing.T) {
	t.Parallel()
	tests := []struct {
		Price         float64
		Deposit       tokens.Asset
		Expense       tokens.Asset
		DepositAmount float64
		ExpenseAmount float64
	}{
		{
			Price:         60000.0, // NOM/ETH
			Deposit:       tokens.ETH,
			Expense:       tokens.NOM,
			DepositAmount: 2.0,
			ExpenseAmount: 120000.0,
		},
		{
			Price:         10.0, // NOM/USDC
			Deposit:       tokens.USDC,
			Expense:       tokens.NOM,
			DepositAmount: 1.0,
			ExpenseAmount: 10.0,
		},
		{
			Price:         1000.0, // USDC/ETH
			Deposit:       tokens.ETH,
			Expense:       tokens.USDC,
			DepositAmount: 1.0,
			ExpenseAmount: 1000.0,
		},
		{
			Price:         1001.001, // USDC/ETH
			Deposit:       tokens.ETH,
			Expense:       tokens.USDC,
			DepositAmount: 2.2,
			ExpenseAmount: 2202.2022,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("%d-%s-%s", i, test.Deposit.Symbol, test.Expense.Symbol), func(t *testing.T) {
			t.Parallel()
			price := types.Price{
				Price:   new(big.Rat).SetFloat64(test.Price),
				Deposit: test.Deposit,
				Expense: test.Expense,
			}

			actualExpense := price.ToExpense(test.Deposit.F64ToAmt(test.DepositAmount))
			require.InEpsilonf(t,
				test.ExpenseAmount,
				test.Expense.AmtToF64(actualExpense),
				0.001,
				"expected %v, got %v", test.ExpenseAmount, test.Expense.AmtToF64(actualExpense),
			)

			actualDeposit := price.ToDeposit(test.Expense.F64ToAmt(test.ExpenseAmount))
			require.InEpsilonf(t,
				test.DepositAmount,
				test.Deposit.AmtToF64(actualDeposit),
				0.001,
				"expected %v, got %v", test.DepositAmount, test.Deposit.AmtToF64(actualDeposit),
			)

			price2 := price.WithFeeBips(100) // 1% fee
			actualExpense2 := price2.ToExpense(test.Deposit.F64ToAmt(test.DepositAmount))
			require.InEpsilonf(t,
				test.ExpenseAmount*0.99,
				test.Expense.AmtToF64(actualExpense2),
				0.001,
				"expected %v, got %v", test.ExpenseAmount, test.Expense.AmtToF64(actualExpense),
			)

			actualDeposit2 := price2.ToDeposit(test.Expense.F64ToAmt(test.ExpenseAmount))
			require.InEpsilonf(t,
				test.DepositAmount*1.01,
				test.Deposit.AmtToF64(actualDeposit2),
				0.001,
				"expected %v, got %v", test.DepositAmount, test.Deposit.AmtToF64(actualDeposit),
			)
		})
	}
}
