package app

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/lib/contracts"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/tutil"
	"github.com/omni-network/omni/solver/types"

	"github.com/ethereum/go-ethereum/common"

	"github.com/stretchr/testify/require"
)

//go:generate go test . -run=TestCheck -golden

//nolint:tparallel,paralleltest // subtests use same mock controller
func TestCheck(t *testing.T) {
	t.Parallel()

	solver := eoa.MustAddress(netconf.Devnet, eoa.RoleSolver)

	// inbox / outbox addr only matters for mocks, using devnet
	addrs, err := contracts.GetAddresses(t.Context(), netconf.Devnet)
	require.NoError(t, err)
	outbox := addrs.SolverNetOutbox

	for _, tt := range checkTestCases(t, solver) {
		t.Run(tt.name, func(t *testing.T) {
			backends, clients := testBackends(t)

			callAllower := func(_ uint64, _ common.Address, _ []byte) bool { return !tt.disallowCall }
			handler := handlerAdapter(newCheckHandler(newChecker(backends, callAllower, solver, outbox)))

			if tt.mock != nil {
				tt.mock(clients)

				destClient := clients.Client(t, tt.req.DestinationChainID)
				mockDidFill(t, destClient, outbox, false)
				mockFill(t, destClient, outbox, tt.res.RejectReason == types.RejectDestCallReverts.String())
				mockFillFee(t, destClient, outbox)
			}

			body, err := json.Marshal(tt.req)
			require.NoError(t, err)

			ctx := t.Context()
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

			require.Equal(t, tt.res.Rejected, res.Rejected)
			require.Equal(t, tt.res.RejectReason, res.RejectReason)
			require.Equal(t, tt.res.Accepted, res.Accepted)

			clients.Finish(t)

			if tt.testdata {
				tutil.RequireGoldenBytes(t, indent(body), tutil.WithFilename(t.Name()+"/req_body.json"))
				tutil.RequireGoldenBytes(t, indent(respBody), tutil.WithFilename(t.Name()+"/resp_body.json"))
			}
		})
	}
}

// indent returns the json bytes indented.
func indent(bz []byte) []byte {
	var buf bytes.Buffer
	err := json.Indent(&buf, bz, "", "  ")
	if err != nil {
		panic(err)
	}

	return buf.Bytes()
}
