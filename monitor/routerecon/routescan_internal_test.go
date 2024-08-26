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

	ctx := context.Background()
	resp, err := paginateLatestCrossTx(ctx, queryFilter{})
	require.NoError(t, err)
	require.NotEmpty(t, resp.ID)

	bz, err := json.MarshalIndent(resp, "", "  ")
	require.NoError(t, err)
	t.Log(string(bz))
}

// TestIntegrationFalse ensures the integration flag defaults to false.
func TestIntegrationFalse(t *testing.T) {
	t.Parallel()
	require.False(t, *integration)
}
