package db_test

import (
	"context"
	"testing"

	"github.com/omni-network/omni/explorer/db"

	"github.com/stretchr/testify/require"
)

func TestMsgHooks(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	client := db.CreateTestEntClient(t)
	t.Cleanup(func() {
		if err := client.Close(); err != nil {
			t.Error(err)
		}
	})

	// This should provide us with a 2 blocks
	// Block 1: 1 message, 0 receipt
	// Block 2: 1 message, 1 receipt (for the message in block 1)
	blocks := db.CreateTestBlocks(ctx, t, client, 2)

	// grab the first block
	b := blocks[0]

	// get the first message
	msg := b.QueryMsgs().FirstX(ctx)

	// count the number of receipts from the message
	msgCount := msg.QueryReceipts().CountX(ctx)

	// get the first receipt
	receipt := msg.QueryReceipts().FirstX(ctx)

	// count the number of messages from the receipt
	receiptIds := receipt.QueryMsgs().CountX(ctx)

	require.Equal(t, 1, msgCount)
	require.Equal(t, 1, receiptIds)
}
