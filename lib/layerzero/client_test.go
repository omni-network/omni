package layerzero_test

import (
	"flag"
	"testing"

	"github.com/omni-network/omni/lib/layerzero"

	"github.com/stretchr/testify/require"
)

var integration = flag.Bool("integration", false, "run integration tests")

func TestGetMessagesByTx(t *testing.T) {
	t.Parallel()

	if !*integration {
		t.Skip("Skipping integration test. Use -integration flag to run")
	}

	client := layerzero.NewClient(layerzero.MainnetAPI)
	ctx := t.Context()

	// Known transaction hashes for testing
	deliveredTxHash := "0xed381b137e2c7f7bd9630a7163274c08f0ddc47ac223c76418e1c1629c5bc83e"
	failedTxHash := "0x3f1b368e7be03a426a145feae08796d52819d55bb3a2daab0a63ba103903a694"

	t.Run("delivered message", func(t *testing.T) {
		t.Parallel()

		messages, err := client.GetMessagesByTx(ctx, deliveredTxHash)
		require.NoError(t, err)
		require.Len(t, messages, 1, "Expected exactly one message")
		require.True(t, messages[0].IsDelivered(), "Expected message to be delivered")
	})

	t.Run("failed message", func(t *testing.T) {
		t.Parallel()

		messages, err := client.GetMessagesByTx(ctx, failedTxHash)
		require.NoError(t, err)
		require.Len(t, messages, 1, "Expected exactly one message")
		require.True(t, messages[0].IsFailed(), "Expected message to be failed")
	})
}
