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
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/tutil"
	"github.com/omni-network/omni/solver/client"
	"github.com/omni-network/omni/solver/types"

	"github.com/ethereum/go-ethereum/common"

	"github.com/stretchr/testify/require"
)

//go:generate go test . -run=TestCheck -golden

//nolint:tparallel,paralleltest // subtests use same mock controller
func TestCheck(t *testing.T) {
	t.Parallel()

	solver := eoa.MustAddress(netconf.Devnet, eoa.RoleSolver)

	priceFunc := unaryPrice

	// inbox / outbox addr only matters for mocks, using devnet
	addrs, err := contracts.GetAddresses(t.Context(), netconf.Devnet)
	require.NoError(t, err)
	outbox := addrs.SolverNetOutbox

	for _, tt := range checkTestCases(t, solver, outbox) {
		t.Run(tt.name, func(t *testing.T) {
			backends, clients := testBackends(t)

			callAllower := func(_ uint64, _ common.Address, _ []byte) bool { return !tt.disallowCall }
			handler := handlerAdapter(newCheckHandler(
				newChecker(backends, callAllower, priceFunc, solver, outbox),
				func(ctx context.Context, req types.CheckRequest) (types.CallTrace, error) {
					require.True(t, tt.req.Debug)
					require.True(t, tt.trace == nil || tt.traceErr == nil)

					if tt.trace == nil {
						return types.CallTrace{}, tt.traceErr
					}

					return *tt.trace, tt.traceErr
				},
			))

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

			if tt.req.Debug && tt.trace != nil {
				require.Equal(t, tt.trace.Map(), res.Trace)
			}

			if tt.req.Debug && tt.traceErr != nil {
				require.Equal(t, map[string]any{"error": errors.Format(tt.traceErr)}, res.Trace)
			}

			if !tt.req.Debug {
				require.Nil(t, res.Trace)
				require.Nil(t, tt.trace)
				require.NoError(t, tt.traceErr)
			}

			res2 := fetchResponseViaClient(t, handler, tt.req)
			require.Equal(t, res, res2)

			clients.Finish(t)

			if tt.testdata {
				tutil.RequireGoldenBytes(t, indent(body), tutil.WithFilename(t.Name()+"/req_body.json"))
				tutil.RequireGoldenBytes(t, indent(respBody), tutil.WithFilename(t.Name()+"/resp_body.json"))
			}
		})
	}
}

func TestCheckError(t *testing.T) {
	t.Parallel()

	msg := "example error message"
	handler := handlerAdapter(newCheckHandler(func(context.Context, types.CheckRequest) error {
		return APIError{
			Err:        errors.New(msg),
			StatusCode: http.StatusServiceUnavailable,
		}
	}, noopTracer))

	srv := httptest.NewServer(handler)
	defer srv.Close()

	cl := client.New(srv.URL)
	_, err := cl.Check(t.Context(), types.CheckRequest{})
	require.ErrorContains(t, err, msg)
}

func fetchResponseViaClient(t *testing.T, h http.Handler, req types.CheckRequest) types.CheckResponse {
	t.Helper()

	srv := httptest.NewServer(h)
	defer srv.Close()

	apiClient := client.New(srv.URL)
	res, err := apiClient.Check(t.Context(), req)
	require.NoError(t, err)

	return res
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
