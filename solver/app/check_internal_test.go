package app

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/lib/contracts"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/solver/types"

	"github.com/stretchr/testify/require"
)

//nolint:tparallel // subtests use same mock controller
func TestCheck(t *testing.T) {
	t.Parallel()

	solver := eoa.MustAddress(netconf.Devnet, eoa.RoleSolver)

	// inbox / outbox addr only matters for mocks, using devnet
	addrs, err := contracts.GetAddresses(context.Background(), netconf.Devnet)
	require.NoError(t, err)
	outbox := addrs.SolverNetOutbox
	inbox := addrs.SolverNetInbox

	for _, tt := range checkTestCases(t, solver) {
		t.Run(tt.name, func(t *testing.T) {
			backends, clients := testBackends(t)

			handler := handlerAdapter(newCheckHandler(newChecker(backends, solver, inbox, outbox)))

			if tt.mock != nil {
				tt.mock(clients)

				destClient := clients.Client(t, tt.req.DestinationChainID)
				mockDidFill(t, destClient, outbox, false)
				mockFill(t, destClient, outbox, tt.res.RejectReason == rejectDestCallReverts.String())
				mockFillFee(t, destClient, outbox)

				srcClient := clients.Client(t, tt.req.SourceChainID)
				mockGetNextID(t, srcClient, inbox)
			}

			body, err := json.Marshal(tt.req)
			require.NoError(t, err)

			ctx := context.Background()
			req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpointCheck, bytes.NewBuffer(body))
			require.NoError(t, err)

			rr := httptest.NewRecorder()
			handler.ServeHTTP(rr, req)

			require.Equal(t, http.StatusOK, rr.Code)

			respBody, err := io.ReadAll(rr.Body)
			require.NoError(t, err)

			var res types.CheckResponse
			err = json.Unmarshal(respBody, &res)
			require.NoError(t, err)

			t.Logf("resp_body: %s", respBody)

			require.Equal(t, tt.res.Rejected, res.Rejected)
			require.Equal(t, tt.res.RejectReason, res.RejectReason)
			require.Equal(t, tt.res.Accepted, res.Accepted)

			clients.Finish(t)
		})
	}
}

// TestCheckRequestParsing calls the handlerAdapter directly for valid and invalid JSON scenarios.
func TestCheckRequestParsing(t *testing.T) {
	t.Parallel()

	checkHandler := newCheckHandler(func(ctx context.Context, req types.CheckRequest) error {
		return nil // noop logic, just testing request parsing.
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
			name:           "wrong data type for sourceChainId (boolean instead of integer)",
			jsonPayload:    `{"sourceChainId": true, "destChainId": 2, "deposit": {"token": "0x1234567890123456789012345678901234567890", "amount": "0x5f5e100"}}`,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "wrong data type for destChainId (string instead of int)",
			jsonPayload:    `{"sourceChainId": 1, "destChainId": "two", "deposit": {"token": "0x1234567890123456789012345678901234567890", "amount": "0x5f5e100"}}`,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "wrong data type for destChainId (boolean instead of integer)",
			jsonPayload:    `{"sourceChainId": 1, "destChainId": true, "deposit": {"token": "0x1234567890123456789012345678901234567890", "amount": "0x5f5e100"}}`,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "wrong data type for calls (object instead of array)",
			jsonPayload:    `{"sourceChainId": 1, "destChainId": 2, "calls": {}}`,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "wrong data type for expenses (object instead of array)",
			jsonPayload:    `{"sourceChainId": 1, "destChainId": 2, "expenses": {}}`,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "wrong data type for deposit (array instead of object)",
			jsonPayload:    `{"sourceChainId": 1, "destChainId": 2, "deposit": []}`,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "wrong data type for deposit (number instead of object)",
			jsonPayload:    `{"sourceChainId": 1, "destChainId": 2, "deposit": 123}`,
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
			name:           "empty calls array is allowed",
			jsonPayload:    `{"sourceChainId": 1, "destChainId": 2, "calls": []}`,
			expectedStatus: http.StatusOK,
		},
		{
			name:           "empty expenses array is allowed",
			jsonPayload:    `{"sourceChainId": 1, "destChainId": 2, "expenses": []}`,
			expectedStatus: http.StatusOK,
		},
		{
			name:           "empty deposit object is allowed",
			jsonPayload:    `{"sourceChainId": 1, "destChainId": 2, "deposit": {}}`,
			expectedStatus: http.StatusOK,
		},
		{
			name:           "incorrect hex encoding in calls",
			jsonPayload:    `{"sourceChainId": 1, "destChainId": 2, "calls": [{"target": "0x1234567890123456789012345678901234567890", "data": "XYZ123", "value": 1000}], "deposit": {"token": "0x1234567890123456789012345678901234567890", "amount": "0x5f5e100"}}`,
			expectedStatus: http.StatusBadRequest,
		},
		{
			// 0xfffffffffffffffffffffffffffffffffffffffffffffffffffffa0a1fef00 is the hex representation of -100000000.
			// This check should pass on request parsing, but will fail on validation checks.
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
			name:           "negative deposit amount as number",
			jsonPayload:    `{"sourceChainId": 1, "destChainId": 2, "deposit": {"token": "0x1234567890123456789012345678901234567890", "amount": -100000000}}`,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "missing token field in deposit",
			jsonPayload:    `{"sourceChainId": 1, "destChainId": 2, "deposit": {"amount": "0x5f5e100"}}`,
			expectedStatus: http.StatusOK,
		},
		{
			// default value of "amount" field is 0
			name:           "missing amount field in deposit",
			jsonPayload:    `{"sourceChainId": 1, "destChainId": 2, "deposit": {"token": "0x1234567890123456789012345678901234567890"}}`,
			expectedStatus: http.StatusOK,
		},
		// TODO: duplicate field detection
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			req := httptest.NewRequest(http.MethodPost, endpointCheck, bytes.NewBufferString(tt.jsonPayload))
			req.Header.Set("Content-Type", "application/json")
			rec := httptest.NewRecorder()
			handlerAdapter(checkHandler).ServeHTTP(rec, req)
			require.Equal(t, tt.expectedStatus, rec.Code)
		})
	}
}
