package db_test

import (
	"context"
	"testing"

	"github.com/omni-network/omni/explorer/db"

	"github.com/stretchr/testify/assert"
)

func TestMsgHooks(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	client := db.CreateTestEntClient(t)
	blocks := db.CreateTestBlocks(ctx, t, client, 2)

	b := blocks[0]
	// TODO: fix this
	msg := b.QueryMsgs().FirstX(ctx)
	msgCount := msg.QueryReceipts().CountX(ctx)
	// receiptIds := msg.QueryReceipts().IDsX(ctx)
	assert.Equal(t, 0, msgCount)
}
