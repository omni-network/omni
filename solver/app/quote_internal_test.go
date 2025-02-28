package app

import (
	"bytes"
	"context"
	"encoding/json"
	"math/big"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/solver/types"

	"github.com/ethereum/go-ethereum/common"

	"github.com/stretchr/testify/require"
)

func TestQuote(t *testing.T) {
	t.Parallel()

	omegaOMNIAddr := omniERC20(netconf.Omega).Address

	tests := []struct {
		name   string
		req    types.QuoteRequest
		res    types.QuoteResponse
		expErr *JSONErrorResponse
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
		},
		{
			name: "quote deposit 10 eth expense",
			req: types.QuoteRequest{
				SourceChainID:      evmchain.IDEthereum,
				DestinationChainID: evmchain.IDArbitrumOne,
				Deposit:            zeroAddrAmt,
				Expense:            mockAddrAmt("10000000000000000000"),
			},
			res: types.QuoteResponse{
				Deposit: mockAddrAmt("10030000000000000000"),
				Expense: mockAddrAmt("10000000000000000000"),
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
		},
		{
			name: "no deposit of expense amount specified",
			req: types.QuoteRequest{
				SourceChainID:      evmchain.IDEthereum,
				DestinationChainID: evmchain.IDArbitrumOne,
				Deposit:            zeroAddrAmt,
				Expense:            zeroAddrAmt,
			},
			expErr: &JSONErrorResponse{
				Code:    http.StatusBadRequest,
				Status:  http.StatusText(http.StatusBadRequest),
				Message: "deposit and expense amount cannot be both zero or both non-zero",
			},
		},
		{
			name: "both deposit and expense amount specified",
			req: types.QuoteRequest{
				SourceChainID:      evmchain.IDEthereum,
				DestinationChainID: evmchain.IDArbitrumOne,
				Deposit:            mockAddrAmt("1000000000000000000"),
				Expense:            mockAddrAmt("1000000000000000000"),
			},
			expErr: &JSONErrorResponse{
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
			expErr: &JSONErrorResponse{
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
			expErr: &JSONErrorResponse{
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
			expErr: &JSONErrorResponse{
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
			expErr: &JSONErrorResponse{
				Code:    http.StatusBadRequest,
				Status:  http.StatusText(http.StatusBadRequest),
				Message: "InvalidDeposit: deposit and expense must be of the same chain class (e.g. mainnet, testnet)",
			},
		},
	}
	for _, tt := range tests {
		handler := handlerAdapter(newQuoteHandler(quoter))

		body, err := json.Marshal(tt.req)
		require.NoError(t, err)

		ctx := context.Background()
		req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpointQuote, bytes.NewBuffer(body))
		require.NoError(t, err)

		rr := httptest.NewRecorder()
		handler.ServeHTTP(rr, req)

		var res struct {
			types.QuoteResponse
			JSONErrorResponse
		}
		require.NoError(t, json.NewDecoder(rr.Body).Decode(&res))
		if rr.Code != http.StatusOK {
			require.Equal(t, res.Code, rr.Code)
		} else {
			require.Empty(t, res.Code)
		}
		require.Equal(t, tt.res, res.QuoteResponse)
	}
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
