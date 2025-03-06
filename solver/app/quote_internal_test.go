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
		expErr *types.JSONErrorResponse
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
			expErr: &types.JSONErrorResponse{
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
			expErr: &types.JSONErrorResponse{
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
			expErr: &types.JSONErrorResponse{
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
			expErr: &types.JSONErrorResponse{
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
			expErr: &types.JSONErrorResponse{
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
			expErr: &types.JSONErrorResponse{
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

		// Response is either a QuoteResponse or a JSONErrorResponse
		var res struct {
			types.QuoteResponse
			types.JSONErrorResponse
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

// TestQuoteRequestParsing calls the handlerAdapter directly for valid and invalid JSON scenarios.
func TestQuoteRequestParsing(t *testing.T) {
	t.Parallel()

	quoteHandler := newQuoteHandler(func(ctx context.Context, req types.QuoteRequest) (types.QuoteResponse, error) {
		return types.QuoteResponse{}, nil // No-op logic, just testing request parsing.
	})

	tests := []struct {
		name           string
		jsonPayload    string
		expectedStatus int
	}{
		{
			name:           "malformed JSON",
			jsonPayload:    `{"sourceChainId": 1, "destChainId": 2, "deposit": {"token": "0x1234567890123456789012345678901234567890", "amount": "0x5f5e100"}`,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "wrong data type for sourceChainId (string instead of int)",
			jsonPayload:    `{"sourceChainId": "one", "destChainId": 2, "deposit": {"token": "0x1234567890123456789012345678901234567890", "amount": "0x5f5e100"}}`,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "wrong data type for destChainId (string instead of int)",
			jsonPayload:    `{"sourceChainId": 1, "destChainId": "two", "deposit": {"token": "0x1234567890123456789012345678901234567890", "amount": "0x5f5e100"}}`,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "wrong data type for sourceChainId (boolean instead of integer)",
			jsonPayload:    `{"sourceChainId": true, "destChainId": 2, "deposit": {"token": "0x1234567890123456789012345678901234567890", "amount": "0x5f5e100"}}`,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "wrong data type for destChainId (boolean instead of integer)",
			jsonPayload:    `{"sourceChainId": 1, "destChainId": true, "deposit": {"token": "0x1234567890123456789012345678901234567890", "amount": "0x5f5e100"}}`,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "wrong data type for deposit (array instead of object)",
			jsonPayload:    `{"sourceChainId": 1, "destChainId": 2, "deposit": []}`,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "extra unexpected field",
			jsonPayload:    `{"sourceChainId": 1, "destChainId": 2, "randomField": "unexpected", "deposit": {"token": "0x1234567890123456789012345678901234567890", "amount": "0x5f5e100"}}`,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "null value in deposit amount",
			jsonPayload:    `{"sourceChainId": 1, "destChainId": 2, "deposit": {"token": "0x1234567890123456789012345678901234567890", "amount": null}}`,
			expectedStatus: http.StatusOK,
		},
		{
			name:           "empty JSON",
			jsonPayload:    `{}`,
			expectedStatus: http.StatusOK,
		},
		{
			name:           "invalid address format",
			jsonPayload:    `{"sourceChainId": 1, "destChainId": 2, "deposit": {"token": "InvalidAddress", "amount": "0x5f5e100"}}`,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "empty string for address field",
			jsonPayload:    `{"sourceChainId": 1, "destChainId": 2, "deposit": {"token": "", "amount": "0x5f5e100"}}`,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "empty deposit object is allowed",
			jsonPayload:    `{"sourceChainId": 1, "destChainId": 2, "deposit": {}}`,
			expectedStatus: http.StatusOK,
		},
		{
			name:           "negative deposit amount as hex string",
			jsonPayload:    `{"sourceChainId": 1, "destChainId": 2, "deposit": {"token": "0x1234567890123456789012345678901234567890", "amount": "0xfffffffffffffffffffffffffffffffffffffffffffffffffffffa0a1fef00"}}`,
			expectedStatus: http.StatusOK,
		},
		{
			name:           "negative deposit amount as number",
			jsonPayload:    `{"sourceChainId": 1, "destChainId": 2, "deposit": {"token": "0x1234567890123456789012345678901234567890", "amount": -100000000}}`,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "zero deposit amount as number",
			jsonPayload:    `{"sourceChainId": 1, "destChainId": 2, "deposit": {"token": "0x1234567890123456789012345678901234567890", "amount": 0}}`,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "zero deposit amount as string (non hex encoded)",
			jsonPayload:    `{"sourceChainId": 1, "destChainId": 2, "deposit": {"token": "0x1234567890123456789012345678901234567890", "amount": "0"}}`,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "missing token field in deposit",
			jsonPayload:    `{"sourceChainId": 1, "destChainId": 2, "deposit": {"amount": "0x5f5e100"}}`,
			expectedStatus: http.StatusOK,
		},
		{
			name:           "missing amount field in deposit",
			jsonPayload:    `{"sourceChainId": 1, "destChainId": 2, "deposit": {"token": "0x1234567890123456789012345678901234567890"}}`,
			expectedStatus: http.StatusOK,
		},
		// TODO: duplicate field detection
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			req := httptest.NewRequest(http.MethodPost, endpointQuote, bytes.NewBufferString(tt.jsonPayload))
			req.Header.Set("Content-Type", "application/json")
			rec := httptest.NewRecorder()
			handlerAdapter(quoteHandler).ServeHTTP(rec, req)
			require.Equal(t, tt.expectedStatus, rec.Code)
		})
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
