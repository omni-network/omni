package fireblocks_test

import (
	"context"
	"testing"

	"github.com/omni-network/omni/lib/fireblocks"

	"github.com/stretchr/testify/require"
)

const host = "https://api.fireblocks.io"
const apiKey = ""
const privateKeyPath = ""

func TestIntegration(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	if apiKey == "" || privateKeyPath == "" {
		t.Skip("apiKey and privateKeyPath are required for integration tests")
	}
	client, err := fireblocks.NewDefaultClient(apiKey, privateKeyPath, host)
	require.NoError(t, err)

	resp, err := client.CreateAndWait(ctx, fireblocks.TransactionRequestOptions{
		Message: fireblocks.UnsignedRawMessage{
			Content: "test",
		},
	})
	require.NoError(t, err)
	require.NotNil(t, resp)
}

func TestCreateTransaction(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	if apiKey == "" || privateKeyPath == "" {
		t.Skip("apiKey and privateKeyPath are required for integration tests")
	}
	client, err := fireblocks.NewDefaultClient(apiKey, privateKeyPath, host)
	require.NoError(t, err)

	opts := fireblocks.TransactionRequestOptions{
		Message: fireblocks.UnsignedRawMessage{
			Content: "test",
		},
	}
	require.NoError(t, err)
	resp, err := client.CreateTransaction(ctx, opts)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, resp.ID)
}
