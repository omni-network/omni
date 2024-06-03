package app_test

import (
	"context"
	"testing"
	"time"

	"github.com/omni-network/omni/explorer/db/ent"
	"github.com/omni-network/omni/explorer/db/ent/block"
	"github.com/omni-network/omni/explorer/db/ent/msg"
	"github.com/omni-network/omni/explorer/db/ent/receipt"
	"github.com/omni-network/omni/explorer/db/testutil"
	"github.com/omni-network/omni/explorer/indexer/app"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/ethereum/go-ethereum/common"

	"github.com/google/go-cmp/cmp"
	fuzz "github.com/google/gofuzz"
	"github.com/stretchr/testify/require"
)

type results struct {
	blocks        []ent.Block
	messages      []*ent.Msg
	receipts      []*ent.Receipt
	blockReceipts map[uint64][]*ent.Receipt
	blockMessages map[uint64][]*ent.Msg
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

				var msgTxHash, receiptTxHash common.Hash
				fuzz.New().NilChance(0).Fuzz(&msgTxHash)
				fuzz.New().NilChance(0).Fuzz(&receiptTxHash)

				var msgTxHash2, receiptTxHash2 common.Hash
				fuzz.New().NilChance(0).Fuzz(&msgTxHash2)
				fuzz.New().NilChance(0).Fuzz(&receiptTxHash2)

				var blockHash common.Hash
				fuzz.New().NilChance(0).Fuzz(&blockHash)

				var sender, to, relayer common.Address
				fuzz.New().NilChance(0).Fuzz(&sender)
				fuzz.New().NilChance(0).Fuzz(&to)
				fuzz.New().NilChance(0).Fuzz(&relayer)

				var msgData []byte
				fuzz.New().NilChance(0).Fuzz(&msgData)

				sourceChainID := uint64(1)
				destChainID := uint64(2)

				_, err := client.XProviderCursor.Create().SetChainID(sourceChainID).SetHeight(0).SetOffset(0).Save(ctx)
				require.NoError(t, err, "creating source chain cursor")

				_, err = client.XProviderCursor.Create().SetChainID(destChainID).SetHeight(0).SetOffset(0).Save(ctx)
				require.NoError(t, err, "creating dest chain cursor")

				blockHeight := uint64(1)
				blockOffset := uint64(1)
				gasLimit := uint64(1000)
				streamOffset := uint64(0)
				gasUsed := uint64(100)

				tx, err := client.BeginTx(ctx, nil)
				require.NoError(t, err)

				err = app.InsertBlockTX(ctx, tx, xchain.Block{
					BlockHeader: xchain.BlockHeader{
						SourceChainID: sourceChainID,
						BlockHeight:   blockHeight,
						BlockOffset:   blockOffset,
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
							SourceMsgSender: sender,
							DestAddress:     to,
							Data:            msgData,
							DestGasLimit:    gasLimit,
							TxHash:          msgTxHash,
						},
						{
							MsgID: xchain.MsgID{
								StreamID: xchain.StreamID{
									SourceChainID: sourceChainID,
									DestChainID:   destChainID,
								},
								StreamOffset: streamOffset + 1,
							},
							SourceMsgSender: sender,
							DestAddress:     to,
							Data:            msgData,
							DestGasLimit:    gasLimit,
							TxHash:          msgTxHash2,
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
							RelayerAddress: common.Address(relayer[:]),
							TxHash:         receiptTxHash,
						},
						{
							MsgID: xchain.MsgID{
								StreamID: xchain.StreamID{
									SourceChainID: sourceChainID,
									DestChainID:   destChainID,
								},
								StreamOffset: streamOffset + 1,
							},
							GasUsed:        gasUsed,
							Success:        false,
							RelayerAddress: common.Address(relayer[:]),
							TxHash:         receiptTxHash2,
							Error:          []byte("something bad happened"),
						},
					},
					Timestamp: time.Now(),
				})
				require.NoError(t, err)

				b := client.Block.Query().Where(block.Height(blockHeight), block.Offset(blockOffset)).OnlyX(ctx)
				msgs, err := client.Msg.Query().Where(msg.SourceChainID(sourceChainID), msg.DestChainID(destChainID)).All(ctx)
				require.NoError(t, err)
				receipts, err := client.Receipt.Query().Where(receipt.SourceChainID(sourceChainID), receipt.DestChainID(destChainID)).All(ctx)
				require.NoError(t, err)

				require.NotNil(t, b)
				require.Len(t, msgs, 2)
				require.Len(t, receipts, 2)

				blockReceipts := make(map[uint64][]*ent.Receipt)
				blockReceipts[b.Height] = receipts

				blockMessages := make(map[uint64][]*ent.Msg)
				blockMessages[b.Height] = msgs

				return results{
					blocks:        []ent.Block{*b},
					messages:      msgs,
					receipts:      receipts,
					blockMessages: blockMessages,
					blockReceipts: blockReceipts,
				}
			},
			want: want{
				BlockCount:        1,
				MsgCount:          2,
				ReceiptCount:      2,
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
			eval(t, results)
			require.Len(t, results.blocks, tt.want.BlockCount)
			require.Len(t, results.messages, tt.want.MsgCount)
			require.Len(t, results.receipts, tt.want.ReceiptCount)
		})
	}
}

func setupDB(t *testing.T) *ent.Client {
	t.Helper()
	client := testutil.CreateTestEntClient(t)
	t.Cleanup(func() {
		if err := client.Close(); err != nil {
			t.Error(err)
		}
	})

	return client
}

func eval(t *testing.T, r results) {
	t.Helper()

	for _, b := range r.blocks {
		expectedMessages := r.blockMessages[b.Height]
		var expectedMessageIDs []int
		for _, m := range expectedMessages {
			expectedMessageIDs = append(expectedMessageIDs, m.ID)
		}
		actualMessageIDs := b.QueryMsgs().IDsX(context.Background())

		if !cmp.Equal(expectedMessageIDs, actualMessageIDs) {
			t.Errorf("got %v want %v", actualMessageIDs, expectedMessageIDs)
		}

		expectedReceipts := r.blockReceipts[b.Height]
		var expectedReceiptIDs []int
		for _, r := range expectedReceipts {
			expectedReceiptIDs = append(expectedReceiptIDs, r.ID)
		}
		actualReceiptIDs := b.QueryReceipts().IDsX(context.Background())

		if !cmp.Equal(expectedReceiptIDs, actualReceiptIDs) {
			t.Errorf("got %v want %v", actualReceiptIDs, expectedReceiptIDs)
		}
	}
}
