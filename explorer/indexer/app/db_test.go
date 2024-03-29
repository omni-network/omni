package app_test

import (
	"context"
	"testing"
	"time"

	"github.com/omni-network/omni/explorer/db"
	"github.com/omni-network/omni/explorer/db/ent"
	"github.com/omni-network/omni/explorer/db/ent/block"
	"github.com/omni-network/omni/explorer/db/ent/msg"
	"github.com/omni-network/omni/explorer/db/ent/receipt"
	"github.com/omni-network/omni/explorer/indexer/app"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/ethereum/go-ethereum/common"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/require"
)

type results struct {
	blocks   []ent.Block
	messages []*ent.Msg
	receipts []*ent.Receipt
}

type prerequisite func(t *testing.T, ctx context.Context, client *ent.Client) results

func TestDbTransaction(t *testing.T) {
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
			name: "insert_block_with_msgs_and_receipts",
			prerequisites: func(t *testing.T, ctx context.Context, client *ent.Client) results {
				t.Helper()
				tx, err := client.BeginTx(ctx, nil)
				require.NoError(t, err)

				sourceChainID := uint64(1)
				destChainID := uint64(2)

				//client.XProviderCursor.Create().SetChainID(sourceChainID).SetHeight(0).SaveX(ctx)
				//client.XProviderCursor.Create().SetChainID(destChainID).SetHeight(0).SaveX(ctx)

				blockHash := common.Hash([32]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 30})
				blockHeight := uint64(1)

				sourceMessageSender := [20]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
				destAddress := [20]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 21}
				msgData := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
				gasLimit := uint64(1000)
				msgTxHash := common.Hash([32]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32})
				streamOffset := uint64(0)

				gasUsed := uint64(100)
				relayerAddress := [20]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 22}
				receiptTxHash := common.Hash([32]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 33})

				err = app.InsertBlockTX(ctx, tx, xchain.Block{
					BlockHeader: xchain.BlockHeader{
						SourceChainID: sourceChainID,
						BlockHeight:   blockHeight,
						BlockHash:     blockHash,
					},
					Msgs: []xchain.Msg{
						{
							MsgID: xchain.MsgID{
								StreamID: xchain.StreamID{
									SourceChainID: sourceChainID,
									DestChainID:   destChainID,
								},
								StreamOffset: streamOffset,
							},
							SourceMsgSender: sourceMessageSender,
							DestAddress:     destAddress,
							Data:            msgData,
							DestGasLimit:    gasLimit,
							TxHash:          msgTxHash,
						},
					},
					Receipts: []xchain.Receipt{
						{
							MsgID: xchain.MsgID{
								StreamID: xchain.StreamID{
									SourceChainID: sourceChainID,
									DestChainID:   destChainID,
								},
								StreamOffset: streamOffset,
							},
							GasUsed:        gasUsed,
							Success:        true,
							RelayerAddress: common.Address(relayerAddress[:]),
							TxHash:         receiptTxHash,
						},
					},
					Timestamp: time.Now(),
				})
				require.NoError(t, err)

				b := client.Block.Query().Where(block.BlockHeight(blockHeight)).OnlyX(ctx)
				m := client.Msg.Query().Where(msg.SourceChainID(sourceChainID), msg.DestChainID(destChainID), msg.StreamOffset(streamOffset)).OnlyX(ctx)
				r := client.Receipt.Query().Where(receipt.SourceChainID(sourceChainID), receipt.DestChainID(destChainID), receipt.StreamOffset(streamOffset)).OnlyX(ctx)

				require.NotNil(t, b)
				require.NotNil(t, m)
				require.NotNil(t, r)

				return results{
					blocks:   []ent.Block{*b},
					messages: []*ent.Msg{m},
					receipts: []*ent.Receipt{r},
				}
			},
			want: want{
				BlockCount:        1,
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

// Calculate the number of receipts that are not associated with a block.
func calcStrayReceipts(ctx context.Context, receipts []*ent.Receipt) int {
	var count int
	for _, b := range receipts {
		cnt := b.QueryBlock().CountX(ctx)
		count += cnt
	}

	return count
}

// Calculate the number of receipts that are not associated with a block.
func calcStrayMessages(ctx context.Context, msgs []*ent.Msg) int {
	var count int
	for _, b := range msgs {
		cnt := b.QueryBlock().CountX(ctx)
		count += cnt
	}

	return count
}
