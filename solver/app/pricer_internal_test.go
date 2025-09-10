package app

import (
	"context"
	"fmt"
	"math/big"
	"net/http/httptest"
	"testing"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/tokenpricer"
	"github.com/omni-network/omni/lib/tokens"
	"github.com/omni-network/omni/lib/tutil"
	"github.com/omni-network/omni/solver/client"
	"github.com/omni-network/omni/solver/types"

	"github.com/stretchr/testify/require"
)

func must(token tokens.Token, ok bool) tokens.Token {
	if !ok {
		panic("token not found")
	}

	return token
}

//go:generate go test . -run=TestPrice -golden

func TestPrice(t *testing.T) {
	t.Parallel()

	omniNative := must(tokens.Native(evmchain.IDOmniDevnet))
	l1Native := must(tokens.Native(evmchain.IDMockL1))
	l2Native := must(tokens.Native(evmchain.IDMockL2))
	l1NOM := must(tokens.ByAsset(evmchain.IDMockL1, tokens.NOM))
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
			Deposit: l1NOM,
			Expense: omniNative,
			Price:   big.NewRat(1, 1),
			NoFees:  true,
		},
		{
			Deposit: l1Native,
			Expense: omniNative,
			Price:   big.NewRat(3000.0*75.0, 5.0), // 1 unit ETH in NOM =  3000 * 75 (conversion rate) / 5 = 45000 NOM/ETH
		},
		{
			Deposit: l2wstETH,
			Expense: l1wstETH,
			Price:   big.NewRat(1, 1), // 1 unit WSTETH in NOM = 1 WSTETH
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

			for _, includeFees := range []bool{true, false} {
				req := types.PriceRequest{
					SourceChainID:      test.Deposit.ChainID,
					DestinationChainID: test.Expense.ChainID,
					DepositToken:       test.Deposit.UniAddress(),
					ExpenseToken:       test.Expense.UniAddress(),
					IncludeFees:        includeFees,
				}
				actual, err := handlerFunc(t.Context(), req)
				require.NoError(t, err)

				expected := types.Price{
					Price:   test.Price,
					Deposit: test.Deposit.Asset,
					Expense: test.Expense.Asset,
				}
				if includeFees && !test.NoFees {
					expected = expected.WithFeeBips(feeBips(test.Deposit.Asset, test.Expense.Asset))
				}

				require.Equalf(t, expected, actual, "expected=%v, actual=%v, includeFees=%v", expected, actual, includeFees)

				if includeFees {
					tutil.RequireGoldenJSON(t, req, tutil.WithFilename(t.Name()+"/req_body.json"))
					tutil.RequireGoldenJSON(t, actual, tutil.WithFilename(t.Name()+"/resp_body.json"))
				}
			}
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

func TestPriceEndpoint(t *testing.T) {
	t.Parallel()

	handler := handlerAdapter(newPriceHandler(wrapPriceHandlerFunc(unaryPrice)))

	srv := httptest.NewServer(handler)
	defer srv.Close()

	cl := client.New(srv.URL)
	realPrice, err := cl.Price(t.Context(), types.PriceRequest{
		SourceChainID:      evmchain.IDBaseSepolia,
		DestinationChainID: evmchain.IDArbSepolia,
		IncludeFees:        false,
	})
	require.NoError(t, err)

	priceWithFees, err := cl.Price(t.Context(), types.PriceRequest{
		SourceChainID:      evmchain.IDBaseSepolia,
		DestinationChainID: evmchain.IDArbSepolia,
		IncludeFees:        true,
	})
	require.NoError(t, err)
	require.Greater(t, realPrice.F64(), priceWithFees.F64()) // Price with fees is less, since you get less expense tokens for same deposit.
	require.Equal(t, priceWithFees, realPrice.WithFeeBips(feeBips(realPrice.Deposit, realPrice.Expense)))
}
