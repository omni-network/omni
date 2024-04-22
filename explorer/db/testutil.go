package db

import (
	"context"
	"strconv"
	"testing"

	"github.com/omni-network/omni/explorer/db/ent"
	"github.com/omni-network/omni/explorer/db/ent/enttest"
	"github.com/omni-network/omni/explorer/db/ent/migrate"

	"github.com/ethereum/go-ethereum/common"
)

func CreateTestChain(t *testing.T, ctx context.Context, client *ent.Client, chainID uint64) ent.Chain {
	t.Helper()

	name := "test-chain" + strconv.FormatUint(chainID, 10)
	chain := client.Chain.Create().
		SetChainID(chainID).
		SetName(name).
		SaveX(ctx)

	return *chain
}

func CreateTestBlock(t *testing.T, ctx context.Context, client *ent.Client, height int) ent.Block {
	t.Helper()

	sourceChainID := uint64(1)
	blockHashBytes := []byte{1, 3, 23, 111, 27, 45, 98, 103, 94, 55, 1, 3, 23, 111, 27, 45, 98, 103, 94, 55}
	blockHashValue := common.Hash{}
	blockHashValue.SetBytes(blockHashBytes)

	b := client.Block.Create().
		SetSourceChainID(sourceChainID).
		SetBlockHeight(uint64(height)).
		SetBlockHash(blockHashValue.Bytes()).
		SaveX(ctx)

	return *b
}

// CreateTestBlocks creates n test blocks with n messages and n-1 receipts.
func CreateTestBlocks(t *testing.T, ctx context.Context, client *ent.Client, count int) []ent.Block {
	t.Helper()
	destChainID := uint64(2)
	var msg *ent.Msg
	var blocks []ent.Block
	for i := 0; i < count; i++ {
		b := CreateTestBlock(t, ctx, client, i)
		if msg != nil {
			CreateReceipt(t, ctx, client, b, msg.DestChainID, msg.StreamOffset)
		}
		msg = CreateXMsg(t, ctx, client, b, destChainID, uint64(i))
		blocks = append(blocks, b)
	}

	return blocks
}

func CreateXMsg(t *testing.T, ctx context.Context, client *ent.Client, b ent.Block, destChainID uint64, streamOffset uint64) *ent.Msg {
	t.Helper()

	sourceMessageSender := [20]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	destAddress := [20]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 21}
	data := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	txHash := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32}

	msg := client.Msg.Create().
		SetSourceMsgSender(sourceMessageSender[:]).
		SetDestAddress(destAddress[:]).
		SetDestChainID(destChainID).
		SetStreamOffset(streamOffset).
		SetData(data).
		SetDestGasLimit(100).
		SetSourceChainID(b.SourceChainID).
		SetTxHash(txHash).
		SetBlockHash(b.BlockHash).
		SetBlockHeight(b.BlockHeight).
		SaveX(ctx)

	client.Block.UpdateOne(&b).AddMsgs(msg).SaveX(ctx)

	return msg
}

func CreateReceipt(t *testing.T, ctx context.Context, client *ent.Client, b ent.Block, destChainID uint64, streamOffset uint64) *ent.Receipt {
	t.Helper()
	relayerAddress := [20]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 22}
	txHash := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 33}

	receipt := client.Receipt.Create().
		SetGasUsed(100).
		SetSuccess(true).
		SetRelayerAddress(relayerAddress[:]).
		SetSourceChainID(b.SourceChainID).
		SetDestChainID(destChainID).
		SetStreamOffset(streamOffset).
		SetTxHash(txHash).
		SaveX(ctx)

	client.Block.UpdateOne(&b).AddReceipts(receipt).SaveX(ctx)

	return receipt
}

func CreateTestEntClient(t *testing.T) *ent.Client {
	t.Helper()

	entOpts := []enttest.Option{
		enttest.WithOptions(ent.Log(t.Log)),
		enttest.WithMigrateOptions(migrate.WithGlobalUniqueID(true)),
	}

	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1", entOpts...)

	return client
}
