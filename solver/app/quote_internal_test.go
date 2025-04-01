package app

import (
	"bytes"
	"encoding/json"
	"io"
	"math/big"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/tutil"
	"github.com/omni-network/omni/solver/types"

	"github.com/ethereum/go-ethereum/common"

	fuzz "github.com/google/gofuzz"
	"github.com/stretchr/testify/require"
)

func TestQuoteExpense(t *testing.T) {
	t.Parallel()
	tests := []struct {
		Name         string
		DepositGwei  uint64
		DepositEther float64
		Expense      string
	}{
		{
			Name:        "1 gwei",
			DepositGwei: 1,
			Expense:     "997_008_973",
		},
		{
			Name:        "1000 gwei",
			DepositGwei: 1_000,
			Expense:     "997_008_973_080",
		},
		{
			Name:         "0.01 eth",
			DepositEther: 0.01,
			Expense:      "9_970_089_730_807_577",
		},
		{
			Name:         "1 eth",
			DepositEther: 1,
			Expense:      "997_008_973_080_757_726",
		},
		{
			Name:         "1000 eth",
			DepositEther: 1000,
			Expense:      "997_008_973_080_757_726_819",
		},
		{
			Name:        "1003 gwei",
			DepositGwei: 1003,
			Expense:     "1_000_000_000_000",
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()

			dep := bi.Gwei(test.DepositGwei)
			if bi.IsZero(dep) {
				dep = bi.Ether(test.DepositEther)
			}

			expense := expenseFor(dep, standardFeeBips)
			expenseFormatted := addThousandSeparators(expense.String())
			require.Equal(t, test.Expense, expenseFormatted)

			deposit2 := depositFor(expense, standardFeeBips)
			require.Equal(t, dep, deposit2)
		})
	}
}

func addThousandSeparators(num string) string {
	n := len(num)
	if n <= 3 {
		return num
	}

	var result []string
	for i, digit := range num {
		if (n-i)%3 == 0 && i != 0 {
			result = append(result, "_")
		}
		result = append(result, string(digit))
	}

	return strings.Join(result, "")
}

//go:generate go test . -run=TestQuote -golden

func TestQuote(t *testing.T) {
	t.Parallel()

	omegaOMNIAddr := omniERC20(netconf.Omega).Address

	tests := []struct {
		name     string
		req      types.QuoteRequest
		res      types.QuoteResponse
		expErr   types.JSONError
		testdata bool
	}{
		{
			name: "quote deposit 1 eth expense",
			req: types.QuoteRequest{
				SourceChainID:      evmchain.IDEthereum,
				DestinationChainID: evmchain.IDArbitrumOne,
				Deposit:            zeroAddrAmt,
				Expense:            mockAddrAmt("1000000000000000000"),
			},
			res: types.QuoteResponse{
				Deposit: mockAddrAmt("1003000000000000000"),
				Expense: mockAddrAmt("1000000000000000000"),
			},
			testdata: true,
		},
		{
			name: "quote deposit 2 eth expense",
			req: types.QuoteRequest{
				SourceChainID:      evmchain.IDEthereum,
				DestinationChainID: evmchain.IDArbitrumOne,
				Deposit:            zeroAddrAmt,
				Expense:            mockAddrAmt("2000000000000000000"),
			},
			res: types.QuoteResponse{
				Deposit: mockAddrAmt("2006000000000000000"),
				Expense: mockAddrAmt("2000000000000000000"),
			},
		},
		{
			name: "quote expense 1.003 eth deposit",
			req: types.QuoteRequest{
				SourceChainID:      evmchain.IDEthereum,
				DestinationChainID: evmchain.IDArbitrumOne,
				Deposit:            mockAddrAmt("1003000000000000000"),
				Expense:            zeroAddrAmt,
			},
			res: types.QuoteResponse{
				Deposit: mockAddrAmt("1003000000000000000"),
				Expense: mockAddrAmt("1000000000000000000"),
			},
		},
		{
			// no fees for OMNI
			name: "quote deposit 1 OMNI expense",
			req: types.QuoteRequest{
				SourceChainID:      evmchain.IDHolesky,
				DestinationChainID: evmchain.IDOmniOmega,
				Deposit:            types.AddrAmt{Token: omegaOMNIAddr},
				Expense:            mockAddrAmt("10000000000000000000"),
			},
			res: types.QuoteResponse{
				Deposit: types.AddrAmt{Amount: parseInt("10000000000000000000"), Token: omegaOMNIAddr},
				Expense: mockAddrAmt("10000000000000000000"),
			},
		},
		{
			// no fees for OMNI
			name: "quote expense 1 OMNI deposit",
			req: types.QuoteRequest{
				SourceChainID:      evmchain.IDHolesky,
				DestinationChainID: evmchain.IDOmniOmega,
				Deposit:            types.AddrAmt{Amount: parseInt("10000000000000000000"), Token: omegaOMNIAddr},
				Expense:            zeroAddrAmt,
			},
			res: types.QuoteResponse{
				Deposit: types.AddrAmt{Amount: parseInt("10000000000000000000"), Token: omegaOMNIAddr},
				Expense: mockAddrAmt("10000000000000000000"),
			},
			testdata: true,
		},
		{
			name: "no deposit of expense amount specified",
			req: types.QuoteRequest{
				SourceChainID:      evmchain.IDEthereum,
				DestinationChainID: evmchain.IDArbitrumOne,
				Deposit:            zeroAddrAmt,
				Expense:            zeroAddrAmt,
			},
			expErr: types.JSONError{
				Code:    http.StatusBadRequest,
				Status:  http.StatusText(http.StatusBadRequest),
				Message: "deposit and expense amount cannot be both zero or both non-zero",
			},
			testdata: true,
		},
		{
			name: "both deposit and expense amount specified",
			req: types.QuoteRequest{
				SourceChainID:      evmchain.IDEthereum,
				DestinationChainID: evmchain.IDArbitrumOne,
				Deposit:            mockAddrAmt("1000000000000000000"),
				Expense:            mockAddrAmt("1000000000000000000"),
			},
			expErr: types.JSONError{
				Code:    http.StatusBadRequest,
				Status:  http.StatusText(http.StatusBadRequest),
				Message: "deposit and expense amount cannot be both zero or both non-zero",
			},
		},
		{
			name: "unsupported deposit token",
			req: types.QuoteRequest{
				SourceChainID:      evmchain.IDEthereum,
				DestinationChainID: evmchain.IDArbitrumOne,
				Deposit:            types.AddrAmt{Token: common.HexToAddress("0x1234")},
				Expense:            mockAddrAmt("1000000000000000000"),
			},
			expErr: types.JSONError{
				Code:    http.StatusNotFound,
				Status:  http.StatusText(http.StatusNotFound),
				Message: "unsupported deposit token",
			},
		},
		{
			name: "unsupported expense token",
			req: types.QuoteRequest{
				SourceChainID:      evmchain.IDEthereum,
				DestinationChainID: evmchain.IDArbitrumOne,
				Deposit:            mockAddrAmt("1000000000000000000"),
				Expense:            types.AddrAmt{Token: common.HexToAddress("0x1234")},
			},
			expErr: types.JSONError{
				Code:    http.StatusNotFound,
				Status:  http.StatusText(http.StatusNotFound),
				Message: "unsupported expense token",
			},
		},
		{
			name: "invalid deposit (OMNI for ETH)",
			req: types.QuoteRequest{
				SourceChainID:      evmchain.IDOmniMainnet,
				DestinationChainID: evmchain.IDEthereum,
				Deposit:            zeroAddrAmt,
				Expense:            mockAddrAmt("1000000000000000000"),
			},
			expErr: types.JSONError{
				Code:    http.StatusBadRequest,
				Status:  http.StatusText(http.StatusBadRequest),
				Message: "InvalidDeposit: deposit token must match expense token",
			},
		},
		{
			name: "invalid deposit (chain mismatch)",
			req: types.QuoteRequest{
				SourceChainID:      evmchain.IDEthereum,
				DestinationChainID: evmchain.IDHolesky,
				Deposit:            mockAddrAmt("1000000000000000000"),
				Expense:            zeroAddrAmt,
			},
			expErr: types.JSONError{
				Code:    http.StatusBadRequest,
				Status:  http.StatusText(http.StatusBadRequest),
				Message: "InvalidDeposit: deposit and expense must be of the same chain class (e.g. mainnet, testnet)",
			},
			testdata: true,
		},
		{
			name: "expense over max",
			req: types.QuoteRequest{
				SourceChainID:      evmchain.IDEthereum,
				DestinationChainID: evmchain.IDArbitrumOne,
				Deposit:            zeroAddrAmt,
				Expense:            types.AddrAmt{Amount: bi.Ether(10)},
			},
			res: types.QuoteResponse{
				Deposit:           types.AddrAmt{Amount: bi.Ether(10.03)},
				Expense:           types.AddrAmt{Amount: bi.Ether(10)},
				Rejected:          true,
				RejectCode:        types.RejectExpenseOverMax,
				RejectReason:      types.RejectExpenseOverMax.String(),
				RejectDescription: "requested expense exceeds maximum: ask=10 ETH, max=6 ETH",
			},
			testdata: true,
		},
		{
			name: "expense under min",
			req: types.QuoteRequest{
				SourceChainID:      evmchain.IDEthereum,
				DestinationChainID: evmchain.IDArbitrumOne,
				Deposit:            types.AddrAmt{Amount: bi.Ether(0.0001003)},
				Expense:            zeroAddrAmt,
			},
			res: types.QuoteResponse{
				Deposit:           types.AddrAmt{Amount: bi.Ether(0.0001003)},
				Expense:           types.AddrAmt{Amount: bi.Ether(0.0001)},
				Rejected:          true,
				RejectCode:        types.RejectExpenseUnderMin,
				RejectReason:      types.RejectExpenseUnderMin.String(),
				RejectDescription: "requested expense is below minimum: ask=0.0001 ETH, min=0.001 ETH",
			},
			testdata: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			handler := handlerAdapter(newQuoteHandler(quoter))

			body, err := json.Marshal(tt.req)
			require.NoError(t, err)

			ctx := t.Context()
			req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpointQuote, bytes.NewBuffer(body))
			require.NoError(t, err)

			rr := httptest.NewRecorder()
			handler.ServeHTTP(rr, req)

			respBody, err := io.ReadAll(rr.Body)
			require.NoError(t, err)

			// Response is either a QuoteResponse or a JSONErrorResponse
			var res struct {
				types.QuoteResponse
				types.JSONErrorResponse
			}
			require.NoError(t, json.Unmarshal(respBody, &res))
			if rr.Code != http.StatusOK {
				require.Equal(t, res.Error.Code, rr.Code)
			}
			require.Equal(t, tt.res, res.QuoteResponse)
			require.Equal(t, tt.expErr, res.Error)

			if tt.testdata {
				tutil.RequireGoldenBytes(t, indent(body), tutil.WithFilename(t.Name()+"/req_body.json"))
				tutil.RequireGoldenBytes(t, indent(respBody), tutil.WithFilename(t.Name()+"/resp_body.json"))
			}
		})
	}
}

func TestFees(t *testing.T) {
	t.Parallel()

	mustBigStr := func(s string) *big.Int {
		b, ok := new(big.Int).SetString(s, 10)
		if !ok {
			panic("invalid big int")
		}

		return b
	}

	// couple sanity checks w/ known values
	require.Equal(t,
		mustBigStr("1500000000000000000"),
		expenseFor(
			mustBigStr("1504500000000000000"),
			standardFeeBips),
	)

	require.Equal(t,
		mustBigStr("1504500000000000000"),
		depositFor(
			mustBigStr("1500000000000000000"),
			standardFeeBips),
	)

	f := fuzz.New().NilChance(0)
	f.Funcs(func(bi *big.Int, c fuzz.Continue) {
		var val uint64
		c.Fuzz(&val)
		bi.SetUint64(val)
	})

	// then fuzz

	var big *big.Int
	f.Fuzz(&big)

	require.True(t,
		// withinOne, because 1 wei can be lost in rounding
		withinOne(
			depositFor(expenseFor(big, standardFeeBips), standardFeeBips),
			big,
		),
		"depositFor(expenseFor(x)) == x",
	)
}

func parseInt(s string) *big.Int {
	b, ok := new(big.Int).SetString(s, 10)
	if !ok {
		panic("invalid big int")
	}

	return (b)
}

func mockAddrAmt(amt string) types.AddrAmt {
	return types.AddrAmt{Amount: parseInt(amt)}
}

var zeroAddrAmt types.AddrAmt

func withinOne(a, b *big.Int) bool {
	diff := new(big.Int).Sub(a, b) // Compute a - b

	return diff.Abs(diff).Cmp(big.NewInt(1)) <= 0
}
