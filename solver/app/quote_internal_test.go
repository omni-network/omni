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
	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/stretchr/testify/require"
)

type JSONQuoteUnit = types.JSONQuoteUnit

func TestQuote(t *testing.T) {
	t.Parallel()

	omegaOMNIAddr := omniERC20(netconf.Omega).Address

	tests := []struct {
		name   string
		req    QuoteRequest
		res    QuoteResponse
		expErr *JSONErrorResponse
	}{
		{
			name: "quote deposit 1 eth expense",
			req: QuoteRequest{
				SourceChainID:      evmchain.IDEthereum,
				DestinationChainID: evmchain.IDArbitrumOne,
				Deposit:            JSONQuoteUnit{},
				Expense:            JSONQuoteUnit{Amount: parseInt("1000000000000000000")},
			},
			res: QuoteResponse{
				Deposit: JSONQuoteUnit{Amount: parseInt("1003000000000000000")},
				Expense: JSONQuoteUnit{Amount: parseInt("1000000000000000000")},
			},
		},
		{
			name: "quote deposit 10 eth expense",
			req: QuoteRequest{
				SourceChainID:      evmchain.IDEthereum,
				DestinationChainID: evmchain.IDArbitrumOne,
				Deposit:            JSONQuoteUnit{},
				Expense:            JSONQuoteUnit{Amount: parseInt("10000000000000000000")},
			},
			res: QuoteResponse{
				Deposit: JSONQuoteUnit{Amount: parseInt("10030000000000000000")},
				Expense: JSONQuoteUnit{Amount: parseInt("10000000000000000000")},
			},
		},
		{
			name: "quote expense 1.003 eth deposit",
			req: QuoteRequest{
				SourceChainID:      evmchain.IDEthereum,
				DestinationChainID: evmchain.IDArbitrumOne,
				Deposit:            JSONQuoteUnit{Amount: parseInt("1003000000000000000")},
				Expense:            JSONQuoteUnit{},
			},
			res: QuoteResponse{
				Deposit: JSONQuoteUnit{Amount: parseInt("1003000000000000000")},
				Expense: JSONQuoteUnit{Amount: parseInt("1000000000000000000")},
			},
		},
		{
			// no fees for OMNI
			name: "quote deposit 1 OMNI expense",
			req: QuoteRequest{
				SourceChainID:      evmchain.IDHolesky,
				DestinationChainID: evmchain.IDOmniOmega,
				Deposit:            JSONQuoteUnit{Token: omegaOMNIAddr},
				Expense:            JSONQuoteUnit{Amount: parseInt("10000000000000000000")},
			},
			res: QuoteResponse{
				Deposit: JSONQuoteUnit{Amount: parseInt("10000000000000000000"), Token: omegaOMNIAddr},
				Expense: JSONQuoteUnit{Amount: parseInt("10000000000000000000")},
			},
		},
		{
			// no fees for OMNI
			name: "quote expense 1 OMNI deposit",
			req: QuoteRequest{
				SourceChainID:      evmchain.IDHolesky,
				DestinationChainID: evmchain.IDOmniOmega,
				Deposit:            JSONQuoteUnit{Amount: parseInt("10000000000000000000"), Token: omegaOMNIAddr},
				Expense:            JSONQuoteUnit{},
			},
			res: QuoteResponse{
				Deposit: JSONQuoteUnit{Amount: parseInt("10000000000000000000"), Token: omegaOMNIAddr},
				Expense: JSONQuoteUnit{Amount: parseInt("10000000000000000000")},
			},
		},
		{
			name: "no deposit of expense amount specified",
			req: QuoteRequest{
				SourceChainID:      evmchain.IDEthereum,
				DestinationChainID: evmchain.IDArbitrumOne,
				Deposit:            JSONQuoteUnit{},
				Expense:            JSONQuoteUnit{},
			},
			expErr: &JSONErrorResponse{
				Code:    http.StatusBadRequest,
				Status:  http.StatusText(http.StatusBadRequest),
				Message: "deposit and expense amount cannot be both zero or both non-zero",
			},
		},
		{
			name: "both deposit and expense amount specified",
			req: QuoteRequest{
				SourceChainID:      evmchain.IDEthereum,
				DestinationChainID: evmchain.IDArbitrumOne,
				Deposit:            JSONQuoteUnit{Amount: parseInt("1000000000000000000")},
				Expense:            JSONQuoteUnit{Amount: parseInt("1000000000000000000")},
			},
			expErr: &JSONErrorResponse{
				Code:    http.StatusBadRequest,
				Status:  http.StatusText(http.StatusBadRequest),
				Message: "deposit and expense amount cannot be both zero or both non-zero",
			},
		},
		{
			name: "unsupported deposit token",
			req: QuoteRequest{
				SourceChainID:      evmchain.IDEthereum,
				DestinationChainID: evmchain.IDArbitrumOne,
				Deposit:            JSONQuoteUnit{Token: common.HexToAddress("0x1234")},
				Expense:            JSONQuoteUnit{Amount: parseInt("1000000000000000000")},
			},
			expErr: &JSONErrorResponse{
				Code:    http.StatusNotFound,
				Status:  http.StatusText(http.StatusNotFound),
				Message: "unsupported deposit token",
			},
		},
		{
			name: "unsupported expense token",
			req: QuoteRequest{
				SourceChainID:      evmchain.IDEthereum,
				DestinationChainID: evmchain.IDArbitrumOne,
				Deposit:            JSONQuoteUnit{Amount: parseInt("1000000000000000000")},
				Expense:            JSONQuoteUnit{Token: common.HexToAddress("0x1234")},
			},
			expErr: &JSONErrorResponse{
				Code:    http.StatusNotFound,
				Status:  http.StatusText(http.StatusNotFound),
				Message: "unsupported expense token",
			},
		},
		{
			name: "invalid deposit (OMNI for ETH)",
			req: QuoteRequest{
				SourceChainID:      evmchain.IDOmniMainnet,
				DestinationChainID: evmchain.IDEthereum,
				Deposit:            JSONQuoteUnit{},
				Expense:            JSONQuoteUnit{Amount: parseInt("1000000000000000000")},
			},
			expErr: &JSONErrorResponse{
				Code:    http.StatusBadRequest,
				Status:  http.StatusText(http.StatusBadRequest),
				Message: "InvalidDeposit: deposit token must match expense token",
			},
		},
		{
			name: "invalid deposit (chain mismatch)",
			req: QuoteRequest{
				SourceChainID:      evmchain.IDEthereum,
				DestinationChainID: evmchain.IDHolesky,
				Deposit:            JSONQuoteUnit{Amount: parseInt("1000000000000000000")},
				Expense:            JSONQuoteUnit{},
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
			QuoteResponse
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

func parseInt(s string) *hexutil.Big {
	b, ok := new(big.Int).SetString(s, 10)
	if !ok {
		panic("invalid big int")
	}

	return (*hexutil.Big)(b)
}
