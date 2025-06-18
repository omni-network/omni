package ts_test

import (
	"bytes"
	"encoding/json"
	"flag"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

type Req struct {
	From   string `json:"from"`
	Amount uint64 `json:"amount"`
}

type Resp struct {
	To     string `json:"to"`
	Amount uint64 `json:"amount"`
}

var integration = flag.Bool("integration", false, "Run integration tests")

func TestTS(t *testing.T) {
	t.Parallel()
	if !*integration {
		t.Skip("Skipping integration test")
	}

	req := Req{
		From:   "0x1234567890abcdef1234567890abcdef12345678",
		Amount: 1000000, // 1 million wei
	}
	bz, err := json.Marshal(req)
	require.NoError(t, err)

	resp, err := http.Post("http://localhost:8000/", "application/json", bytes.NewReader(bz))
	require.NoError(t, err)
	defer resp.Body.Close()
	require.Equal(t, http.StatusOK, resp.StatusCode)

	bz, err = io.ReadAll(resp.Body)
	require.NoError(t, err)

	var res Resp
	err = json.Unmarshal(bz, &res)
	require.NoError(t, err)
	require.Equal(t, req.From+"Z", res.To)
	require.Equal(t, req.Amount+1, res.Amount)
}
