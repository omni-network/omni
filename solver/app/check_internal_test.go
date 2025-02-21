package app

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/lib/contracts"
	"github.com/omni-network/omni/lib/netconf"

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

			handler := newCheckHandler(newChecker(backends, solver, inbox, outbox))

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
			req, err := http.NewRequestWithContext(ctx, http.MethodPost, "api/v1/check", bytes.NewBuffer(body))
			require.NoError(t, err)

			rr := httptest.NewRecorder()
			handler.ServeHTTP(rr, req)

			require.Equal(t, http.StatusOK, rr.Code)

			var res CheckResponse
			err = json.NewDecoder(rr.Body).Decode(&res)
			require.NoError(t, err)

			t.Logf("res: %+v", res)

			require.Equal(t, tt.res.Rejected, res.Rejected)
			require.Equal(t, tt.res.RejectReason, res.RejectReason)
			require.Equal(t, tt.res.Accepted, res.Accepted)

			clients.Finish(t)
		})
	}
}
