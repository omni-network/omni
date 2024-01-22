package handlers_test

import (
	"context"
	"testing"

	"github.com/omni-network/omni/explorer/api/handlers"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHealthCheck(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	resp, err := handlers.GetHealth(ctx)

	require.NoError(t, err)
	assert.NotNil(t, resp)

	require.NotEmptyf(t, resp.Message, "health check response msg")
}
