package app

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/solver/types"

	"github.com/ethereum/go-ethereum/common"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

//nolint:bodyclose,noctx // Not critical for tests
func TestCheckHandlerRequests(t *testing.T) {
	t.Parallel()

	tests := []struct {
		Name    string
		Input   string
		Request types.CheckRequest
	}{
		{
			Name:    "empty",
			Input:   `{}`,
			Request: types.CheckRequest{},
		}, {
			Name:  "missing deposit amount",
			Input: `{"deposit":{"token":"0x0123456789012345678901234567890123456789"}}`,
			Request: types.CheckRequest{
				Deposit: types.AddrAmt{
					Token:  common.HexToAddress("0x0123456789012345678901234567890123456789"),
					Amount: bi.Zero(),
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			handler := newCheckHandler(func(_ context.Context, req types.CheckRequest) error {
				if !assert.Equal(t, test.Request, req) {
					return errors.New("unexpected request")
				}

				return nil
			})

			srv := httptest.NewServer(handlerAdapter(handler))

			req, err := http.NewRequest(http.MethodPost, srv.URL, bytes.NewBufferString(test.Input))
			require.NoError(t, err)

			resp, err := http.DefaultClient.Do(req)
			require.NoError(t, err)
			require.Equal(t, http.StatusOK, resp.StatusCode)
		})
	}
}
