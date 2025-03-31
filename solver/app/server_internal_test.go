package app

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/omni-network/omni/lib/tutil"

	"github.com/stretchr/testify/require"
)

func TestServerCORS(t *testing.T) {
	t.Parallel()

	addr := fmt.Sprintf("127.0.0.1:%d", tutil.RandomAvailablePort(t))

	_, apiCancel := serveAPI(addr, Handler{
		Endpoint: "/api/v1/test",
		ZeroReq:  func() any { return nil },
		HandleFunc: func(ctx context.Context, req any) (any, error) {
			return struct{}{}, nil
		},
	})
	defer apiCancel()

	req, err := http.NewRequestWithContext(t.Context(), http.MethodOptions, fmt.Sprintf("http://%s/api/v1/test", addr), nil)
	require.NoError(t, err)
	req.Header.Set("Origin", "http://example.com")
	req.Header.Add("Access-Control-Request-Headers", "accept")
	req.Header.Add("Access-Control-Request-Headers", "content-type")
	req.Header.Add("Access-Control-Request-Headers", "origin")
	req.Header.Add("Access-Control-Request-Headers", "user-agent")
	req.Header.Set("Access-Control-Request-Method", "POST")

	var resp *http.Response
	// Allow async server to start
	require.Eventually(t, func() bool {
		resp, err = http.DefaultClient.Do(req)
		return err == nil
	}, time.Second, time.Millisecond*100)

	require.Equal(t, http.StatusNoContent, resp.StatusCode)

	require.Equal(t, "*",
		resp.Header.Get("Access-Control-Allow-Origin"))

	require.Equal(t, []string{"accept", "content-type", "origin", "user-agent"},
		resp.Header.Values("Access-Control-Allow-Headers"))

	require.Equal(t, "POST",
		resp.Header.Get("Access-Control-Allow-Methods"))
}
