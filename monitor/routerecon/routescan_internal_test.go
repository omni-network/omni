package routerecon

import (
	"context"
	"encoding/json"
	"flag"
	"testing"

	"github.com/stretchr/testify/require"
)

//go:generate go test . -integration -v -run=TestQueryLatestXChain

var integration = flag.Bool("integration", false, "run routescan integration tests")

func TestQueryLatestXChain(t *testing.T) {
	t.Parallel()
	if !*integration {
		t.Skip("skipping integration test")
	}
	t.Parallel()
	ctx := context.Background()
	resp, err := paginateLatestCrossTx(ctx)
	require.NoError(t, err)
	require.NotEmpty(t, resp.ID)
	bz, err := json.MarshalIndent(resp, "", "  ")
	require.NoError(t, err)
	t.Log(string(bz))
}
