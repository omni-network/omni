package resolvers

import (
	"context"
	"testing"

	"github.com/omni-network/omni/explorer/db/ent"
	"github.com/omni-network/omni/explorer/db/ent/enttest"
	"github.com/omni-network/omni/explorer/db/ent/migrate"

	"github.com/ethereum/go-ethereum/common"
)

func CreateTestBlock(ctx context.Context, t *testing.T, client *ent.Client) ent.Block {
	t.Helper()

	sourceChainID := uint64(1)
	blockHeight := uint64(0)
	blockHashBytes := []byte{1, 3, 23, 111, 27, 45, 98, 103, 94, 55, 1, 3, 23, 111, 27, 45, 98, 103, 94, 55}
	blockHashValue := common.Hash{}
	blockHashValue.SetBytes(blockHashBytes)

	block := client.Block.Create().
		SetSourceChainID(sourceChainID).
		SetBlockHeight(blockHeight).
		SetBlockHash(blockHashValue.Bytes()).
		SaveX(ctx)

	return *block
}

// CreateTestBlocks creates n test blocks with n messages and n-1 receipts.
func CreateTestBlocks(ctx context.Context, t *testing.T, client *ent.Client, count int) {
	t.Helper()
	var msg *ent.Msg
	for i := 0; i < count; i++ {
		block := CreateTestBlock(ctx, t, client)
		if msg != nil {
			createReceipt(ctx, t, client, *msg)
		}
		msg = createXMsg(ctx, t, client, block)
	}
}

func createXMsg(ctx context.Context, t *testing.T, client *ent.Client, block ent.Block) *ent.Msg {
	t.Helper()

	destChain := uint64(2)
	sourceMessageSender := [20]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	destAddress := [20]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 21}
	data := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	txHash := []byte{1, 2, 3, 4}
	msg := client.Msg.Create().
		SetSourceMsgSender(sourceMessageSender[:]).
		SetDestAddress(destAddress[:]).
		SetData(data).
		SetDestGasLimit(100).
		SetSourceChainID(block.SourceChainID).
		SetDestChainID(destChain).
		SetStreamOffset(block.BlockHeight).
		SetTxHash(txHash).
		SetBlockID(block.ID).
		SetBlock(&block).
		SaveX(ctx)

	return msg
}

func createReceipt(ctx context.Context, t *testing.T, client *ent.Client, msg ent.Msg) ent.Receipt {
	t.Helper()

	relayerAddress := [20]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 22}
	receipt := client.Receipt.Create().
		SetGasUsed(100).
		SetSuccess(true).
		SetRelayerAddress(relayerAddress[:]).
		SetSourceChainID(msg.SourceChainID).
		SetDestChainID(msg.DestChainID).
		SetStreamOffset(msg.StreamOffset).
		SetTxHash(msg.TxHash).
		SaveX(ctx)

	return *receipt
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
