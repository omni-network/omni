package app

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/uni"
	"github.com/omni-network/omni/solver/types"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRequestCancel(t *testing.T) {
	t.Parallel()

	serving := make(chan struct{})
	served := make(chan struct{})

	h := Handler{
		Endpoint: "test",
		ZeroReq:  func() any { return nil },
		HandleFunc: func(ctx context.Context, _ any) (any, error) {
			<-ctx.Done()
			return nil, ctx.Err()
		},
	}

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Wrap handlerAdapter to detect status code (since client cancels)
		close(serving)
		w2 := instrumentWriter{ResponseWriter: w}
		handlerAdapter(h).ServeHTTP(&w2, r)
		require.Equal(t, http.StatusRequestTimeout, w2.status)
		close(served)
	}))

	ctx, cancel := context.WithCancel(t.Context())
	// Async: when the handler is serving, cancel the context.
	go func() {
		<-serving
		cancel()
	}()

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, srv.URL, nil)
	require.NoError(t, err)

	_, err = new(http.Client).Do(req)
	require.ErrorIs(t, err, context.Canceled)
	<-served
}

//nolint:paralleltest // Global gateway timeout is modified
func TestGatewayTimeout(t *testing.T) {
	// Replace 10s gateway timeout with faster test value
	cached := gatewayTimeout
	gatewayTimeout = time.Millisecond * 50
	t.Cleanup(func() {
		gatewayTimeout = cached
	})

	h := Handler{
		Endpoint: "test",
		ZeroReq:  func() any { return nil },
		HandleFunc: func(ctx context.Context, _ any) (any, error) {
			<-ctx.Done()
			return nil, ctx.Err()
		},
	}

	srv := httptest.NewServer(handlerAdapter(h))

	req, err := http.NewRequestWithContext(t.Context(), http.MethodPost, srv.URL, nil)
	require.NoError(t, err)

	resp, err := new(http.Client).Do(req)
	require.NoError(t, err)
	require.Equal(t, http.StatusGatewayTimeout, resp.StatusCode)
}

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
					Token:  uni.MustHexToAddress("0x0123456789012345678901234567890123456789"),
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
			}, noopTracer)

			srv := httptest.NewServer(handlerAdapter(handler))

			req, err := http.NewRequest(http.MethodPost, srv.URL, bytes.NewBufferString(test.Input))
			require.NoError(t, err)

			resp, err := http.DefaultClient.Do(req)
			require.NoError(t, err)
			require.Equal(t, http.StatusOK, resp.StatusCode)
		})
	}
}
