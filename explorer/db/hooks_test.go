package db_test

import (
	"context"
	"testing"

	"github.com/omni-network/omni/explorer/db"
	"github.com/omni-network/omni/explorer/db/ent"

	"github.com/google/go-cmp/cmp"
)

type results struct {
	blocks   []ent.Block
	messages []*ent.Msg
	receipts []*ent.Receipt
}

type prerequisite func(t *testing.T, ctx context.Context, client *ent.Client) results

func TestMsgAndReceiptHooks(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	type want struct {
		BlockCount        int
		MsgCount          int
		ReceiptCount      int
		StrayReceiptCount int
		StrayMessageCount int
	}

	tests := []struct {
		name          string
		prerequisites prerequisite // These functions create entries on our db before the evaluation
		want          want
	}{
		{
			name: "create_initial_block_with_msg_then_following_contain_receipts",
			prerequisites: func(t *testing.T, ctx context.Context, client *ent.Client) results {
				t.Helper()
				blocks := db.CreateTestBlocks(t, ctx, client, 2)

				var messages []*ent.Msg
				var receipts []*ent.Receipt
				for _, b := range blocks {
					messages = append(messages, b.QueryMsgs().AllX(ctx)...)
					receipts = append(receipts, b.QueryReceipts().AllX(ctx)...)
				}

				return results{
					blocks:   blocks,
					messages: messages,
					receipts: receipts,
				}
			},
			want: want{
				BlockCount:        2,
				MsgCount:          2,
				ReceiptCount:      1,
				StrayReceiptCount: 0,
				StrayMessageCount: 1,
			},
		},
		{
			name: "create_block_then_receipt_then_msg",
			prerequisites: func(t *testing.T, ctx context.Context, client *ent.Client) results {
				t.Helper()
				destChainID := uint64(2)
				streamOffset := uint64(0)
				block1 := db.CreateTestBlock(t, ctx, client, 0)
				block2 := db.CreateTestBlock(t, ctx, client, 1)
				receipt := db.CreateReceipt(t, ctx, client, block1, destChainID, streamOffset)
				msg := db.CreateXMsg(t, ctx, client, block2, destChainID, streamOffset)

				messages := []*ent.Msg{msg}
				receipts := []*ent.Receipt{receipt}
				blocks := []ent.Block{block1, block2}

				return results{
					blocks:   blocks,
					messages: messages,
					receipts: receipts,
				}
			},
			want: want{
				BlockCount:        2,
				MsgCount:          1,
				ReceiptCount:      1,
				StrayReceiptCount: 0,
				StrayMessageCount: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			client := setupDB(t)
			results := tt.prerequisites(t, ctx, client)
			actual := want{
				len(results.blocks),
				len(results.messages),
				len(results.receipts),
				calcStrayReceipts(ctx, results.receipts),
				calcStrayMessages(ctx, results.messages),
			}

			if !cmp.Equal(tt.want, actual) {
				t.Errorf("unexpected results: %s", cmp.Diff(tt.want, actual))
			}
		})
	}
}

// Calculate the number of messages that are not associated with a receipt.
func calcStrayMessages(ctx context.Context, messages []*ent.Msg) int {
	var count int
	for _, m := range messages {
		receiptCount := m.QueryReceipts().CountX(ctx)
		if receiptCount == 0 {
			count++
		}
	}

	return count
}

func setupDB(t *testing.T) *ent.Client {
	t.Helper()
	client := db.CreateTestEntClient(t)
	t.Cleanup(func() {
		if err := client.Close(); err != nil {
			t.Error(err)
		}
	})

	return client
}

// Calculate the number of receipts that are not associated with a message.
func calcStrayReceipts(ctx context.Context, receipts []*ent.Receipt) int {
	var count int
	for _, r := range receipts {
		messageCount := r.QueryMsgs().CountX(ctx)
		if messageCount == 0 {
			count++
		}
	}

	return count
}
