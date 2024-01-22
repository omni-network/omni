package server_test

import (
	"context"
	"testing"

	"github.com/omni-network/omni/explorer/api/server"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestServerCreate(t *testing.T) {
	t.Parallel()

	port := 8080
	client := server.NewClient(port)
	ctx := context.Background()

	handler, err := client.CreateServer(ctx)

	require.NoError(t, err)
	assert.NotNil(t, handler)
}
