package app

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/tokenpricer"
	"github.com/omni-network/omni/lib/tokens"
	"github.com/omni-network/omni/solver/types"

	"github.com/stretchr/testify/require"
)

func must(token tokens.Token, ok bool) tokens.Token {
	if !ok {
		panic("token not found")
	}

	return token
}

func TestPriceHandler(t *testing.T) {
	t.Parallel()

	omniNative := must(tokens.Native(evmchain.IDOmniDevnet))
	l1Native := must(tokens.Native(evmchain.IDMockL1))
	l2Native := must(tokens.Native(evmchain.IDMockL2))
	l1OMNI := must(tokens.ByAsset(evmchain.IDMockL1, tokens.OMNI))
	l1USDC := must(tokens.ByAsset(evmchain.IDMockL1, tokens.USDC))
	l2wstETH := must(tokens.ByAsset(evmchain.IDMockL2, tokens.WSTETH))
	l1wstETH := must(tokens.ByAsset(evmchain.IDMockL1, tokens.WSTETH))

	handlerFunc := wrapPriceHandlerFunc(newPriceFunc(tokenpricer.NewDevnetMock()))

	tests := []struct {
		Deposit tokens.Token
		Expense tokens.Token
		Price   *big.Rat // See tokenpricer.NewDevnetMock for prices
		NoFees  bool
	}{
		{
			Deposit: l1OMNI,
			Expense: omniNative,
			Price:   big.NewRat(1, 1),
			NoFees:  true,
		},
		{
			Deposit: l1Native,
			Expense: omniNative,
			Price:   big.NewRat(3000.0, 5.0), // 1 unit ETH in OMNI =  3000 / 5 = 600 OMNI/ETH
		},
		{
			Deposit: l2wstETH,
			Expense: l1wstETH,
			Price:   big.NewRat(1, 1), // 1 unit WSTETH in OMNI = 1 WSTETH
		},
		{
			Deposit: l1wstETH,
			Expense: l2Native,
			Price:   big.NewRat(4000.0, 3000.0), //  1 unit WSTETH in ETH = 4000 / 3000 = 4/3 WSTETH/ETH
		},
		{
			Deposit: l1USDC,
			Expense: l2Native,
			Price:   big.NewRat(1.0, 3000.0), // 1 unit USDC in ETH = 1 / 3000  USDC/ETH
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("%d-%s-%s", i, test.Deposit, test.Expense), func(t *testing.T) {
			t.Parallel()
			actual, err := handlerFunc(t.Context(), types.PriceRequest{
				SourceChainID:      test.Deposit.ChainID,
				DestinationChainID: test.Expense.ChainID,
				DepositToken:       test.Deposit.Address,
				ExpenseToken:       test.Expense.Address,
			})
			require.NoError(t, err)

			// Add the expected fees to the test price above
			expected := types.Price{
				Price:   test.Price,
				Deposit: test.Deposit.Asset,
				Expense: test.Expense.Asset,
			}
			if !test.NoFees {
				expected = expected.WithFeeBips(feeBips(test.Deposit.Asset, test.Expense.Asset))
			}

			require.Equalf(t, expected, actual, "expected=%v, actual=%v", expected, actual)
		})
	}
}
