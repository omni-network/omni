package app

import (
	"encoding/json"
	"math/big"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/svmutil"
	"github.com/omni-network/omni/lib/tokenpricer"
	"github.com/omni-network/omni/lib/tokens"
	"github.com/omni-network/omni/lib/tutil"
	"github.com/omni-network/omni/lib/uni"
	"github.com/omni-network/omni/solver/client"
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

			expense := expenseFor(t, dep)
			expenseFormatted := addThousandSeparators(expense.String())
			require.Equal(t, test.Expense, expenseFormatted)

			deposit2 := depositFor(t, expense)
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

	// omegaNOMAddr := nomERC20(netconf.Omega).UniAddress()

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
		//{
		//	// no fees for NOM
		//	name: "quote deposit 1 NOM expense",
		//	req: types.QuoteRequest{
		//		SourceChainID:      evmchain.IDHolesky,
		//		DestinationChainID: evmchain.IDOmniOmega,
		//		Deposit:            types.AddrAmt{Token: omegaNOMAddr},
		//		Expense:            mockAddrAmt("10000000000000000000"),
		//	},
		//	res: types.QuoteResponse{
		//		Deposit: types.AddrAmt{Amount: parseInt("10000000000000000000"), Token: omegaNOMAddr},
		//		Expense: mockAddrAmt("10000000000000000000"),
		//	},
		// },
		//{
		//	// no fees for NOM
		//	name: "quote expense 1 NOM deposit",
		//	req: types.QuoteRequest{
		//		SourceChainID:      evmchain.IDHolesky,
		//		DestinationChainID: evmchain.IDOmniOmega,
		//		Deposit:            types.AddrAmt{Amount: parseInt("10000000000000000000"), Token: omegaNOMAddr},
		//		Expense:            zeroAddrAmt,
		//	},
		//	res: types.QuoteResponse{
		//		Deposit: types.AddrAmt{Amount: parseInt("10000000000000000000"), Token: omegaNOMAddr},
		//		Expense: mockAddrAmt("10000000000000000000"),
		//	},
		//	testdata: true,
		// },
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
				Deposit:            types.AddrAmt{Token: uni.EVMAddress(common.HexToAddress("0x1234"))},
				Expense:            mockAddrAmt("1000000000000000000"),
			},
			expErr: types.JSONError{
				Code:    http.StatusNotFound,
				Status:  http.StatusText(http.StatusNotFound),
				Message: "unsupported deposit token [chain=42161, address=0x0000000000000000000000000000000000000000]",
			},
		},
		{
			name: "unsupported expense token",
			req: types.QuoteRequest{
				SourceChainID:      evmchain.IDEthereum,
				DestinationChainID: evmchain.IDArbitrumOne,
				Deposit:            mockAddrAmt("1000000000000000000"),
				Expense:            types.AddrAmt{Token: uni.EVMAddress(common.HexToAddress("0x1234"))},
			},
			expErr: types.JSONError{
				Code:    http.StatusNotFound,
				Status:  http.StatusText(http.StatusNotFound),
				Message: "unsupported expense token [chain=42161, address=0x0000000000000000000000000000000000001234]",
			},
		},
		{
			name: "valid native swap (NOM for ETH)",
			req: types.QuoteRequest{
				SourceChainID:      evmchain.IDOmniMainnet,
				DestinationChainID: evmchain.IDEthereum,
				Deposit:            zeroAddrAmt,
				Expense:            mockAddrAmt("1000000000000000000"),
			},
			res: types.QuoteResponse{
				Deposit: mockAddrAmt("45135000000000000000000"), // Price is (3000/5) * 75 ~= 45000 NOM/ETH
				Expense: mockAddrAmt("1000000000000000000"),
			},
		},
		{
			name: "valid erc20 to native swap (base USDC for NOM)",
			req: types.QuoteRequest{
				SourceChainID:      evmchain.IDBase,
				DestinationChainID: evmchain.IDOmniMainnet,
				Deposit: types.AddrAmt{
					Token: erc20(evmchain.IDBase, tokens.USDC).UniAddress(),
				},
				Expense: mockAddrAmt("75000000000000000000"),
			},
			res: types.QuoteResponse{
				Deposit: types.AddrAmt{
					Token:  erc20(evmchain.IDBase, tokens.USDC).UniAddress(),
					Amount: parseInt("5015000"), // Price is $0.066667/OMNI (USDC has 6 decimals)
				},
				Expense: mockAddrAmt("75000000000000000000"),
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
				Message: "deposit and expense must be of the same chain class (e.g. mainnet, testnet) [deposit=mainnet, expense=testnet]",
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
				RejectDescription: "requested expense exceeds maximum [ask=10 ETH, max=6 ETH]",
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
				RejectDescription: "requested expense is below minimum [ask=0.0001 ETH, min=0.001 ETH]",
			},
			testdata: true,
		},
		{
			name: "svm USDC to mockL1 USDC",
			req: types.QuoteRequest{
				SourceChainID:      evmchain.IDSolanaLocal,
				DestinationChainID: evmchain.IDMockL1,
				Deposit: types.AddrAmt{
					Token: uni.SVMAddress(svmutil.DevnetUSDCMint.PublicKey()),
				},
				Expense: types.AddrAmt{
					Amount: bi.Dec6(10),
					Token:  erc20(evmchain.IDMockL1, tokens.USDC).UniAddress(),
				},
			},
			res: types.QuoteResponse{
				Deposit: types.AddrAmt{
					Amount: bi.Dec6(10.03),
					Token:  uni.SVMAddress(svmutil.DevnetUSDCMint.PublicKey()),
				},
				Expense: types.AddrAmt{
					Amount: bi.Dec6(10),
					Token:  erc20(evmchain.IDMockL1, tokens.USDC).UniAddress(),
				},
			},
			testdata: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			priceFunc := newPriceFunc(tokenpricer.NewDevnetMock())
			srv := httptest.NewServer(handlerAdapter(newQuoteHandler(newQuoter(priceFunc))))

			var reqBody, respBody []byte
			cl := client.New(srv.URL, client.WithDebugBodies(
				func(b []byte) { reqBody = b },
				func(b []byte) { respBody = b },
			))

			resp, err := cl.Quote(t.Context(), tt.req)
			if err == nil {
				require.Equal(t, tt.res, resp)
			} else {
				var errResp types.JSONErrorResponse
				require.NoError(t, json.Unmarshal(respBody, &errResp))
				require.Equal(t, tt.expErr, errResp.Error)
			}

			if tt.testdata {
				tutil.RequireGoldenBytes(t, indent(reqBody), tutil.WithFilename(t.Name()+"/req_body.json"))
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
		expenseFor(t,
			mustBigStr("1504500000000000000"),
		),
	)

	require.Equal(t,
		mustBigStr("1504500000000000000"),
		depositFor(t,
			mustBigStr("1500000000000000000"),
		),
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
			depositFor(t, expenseFor(t, big)),
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
	return types.AddrAmt{
		Amount: parseInt(amt),
	}
}

var zeroAddrAmt types.AddrAmt

func withinOne(a, b *big.Int) bool {
	diff := new(big.Int).Sub(a, b) // Compute a - b

	return diff.Abs(diff).Cmp(big.NewInt(1)) <= 0
}
