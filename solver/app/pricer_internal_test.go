package app

import (
	"context"
	"fmt"
	"math/big"
	"testing"

	"github.com/omni-network/omni/lib/errors"
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

// unaryPrice is a priceFunc that returns a price for like-for-like 1-to-1 pairs or an error.
// This is the legacy (pre-swaps) behavior.
func unaryPrice(_ context.Context, deposit, expense tokens.Token) (types.Price, error) {
	if !areEqualBySymbol(deposit, expense) {
		return types.Price{}, errors.New("deposit token must match expense token", "deposit", deposit, "expense", expense)
	}

	if deposit.ChainClass != expense.ChainClass {
		// we should reject with UnsupportedDestChain before quoting tokens of different chain classes.
		return types.Price{}, errors.New("deposit and expense must be of the same chain class (e.g. mainnet, testnet)", "deposit", deposit.ChainClass, "expense", expense.ChainClass)
	}

	return types.Price{
		Price:   big.NewRat(1, 1),
		Deposit: deposit.Asset,
		Expense: expense.Asset,
	}, nil
}
