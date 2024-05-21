package testutil

import (
	"context"
	"strconv"
	"testing"
	"time"

	"github.com/omni-network/omni/explorer/db/ent"
	"github.com/omni-network/omni/explorer/db/ent/enttest"
	"github.com/omni-network/omni/explorer/db/ent/migrate"

	"github.com/ethereum/go-ethereum/common"

	gofuzz "github.com/google/gofuzz"
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

func CreateTestBlock(t *testing.T, ctx context.Context, client *ent.Client, height, offset uint64, ts time.Time) ent.Block {
	t.Helper()

	const chainID = 1655
	var blockHash common.Hash
	gofuzz.New().NilChance(0).Fuzz(&blockHash)

	b := client.Block.Create().
		SetChainID(chainID).
		SetHeight(height).
		SetOffset(offset).
		SetHash(blockHash.Bytes()).
		SetTimestamp(ts).
		SaveX(ctx)

	return *b
}

// CreateTestBlocks creates n test blocks with n messages and n-1 receipts.
func CreateTestBlocks(t *testing.T, ctx context.Context, client *ent.Client, count uint64) []ent.Block {
	t.Helper()
	destChainID := uint64(1656)
	var msg *ent.Msg
	var blocks []ent.Block
	ts := time.Now()
	for i := uint64(0); i < count; i++ {
		ts = ts.Add(10 * time.Millisecond) // add some small "delay" to ensure different timestamps in tests
		b := CreateTestBlock(t, ctx, client, i*2, i, ts)
		if msg != nil {
			CreateReceipt(t, ctx, client, b, msg.DestChainID, msg.Offset)
		}
		msg = CreateXMsg(t, ctx, client, b, destChainID, i)
		blocks = append(blocks, b)
	}

	return blocks
}

func CreateXMsg(t *testing.T, ctx context.Context, client *ent.Client, b ent.Block, destChainID uint64, streamOffset uint64) *ent.Msg {
	t.Helper()

	sender := [20]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	to := [20]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 21}
	data := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	txHash := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32}

	msg := client.Msg.Create().
		SetSender(sender[:]).
		SetTo(to[:]).
		SetDestChainID(destChainID).
		SetOffset(streamOffset).
		SetData(data).
		SetGasLimit(100).
		SetSourceChainID(b.ChainID).
		SetTxHash(txHash).
		SaveX(ctx)

	client.Block.UpdateOne(&b).AddMsgs(msg).SaveX(ctx)

	return msg
}

func CreateReceipt(t *testing.T, ctx context.Context, client *ent.Client, b ent.Block, destChainID uint64, offset uint64) *ent.Receipt {
	t.Helper()
	relayerAddress := [20]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 22}
	txHash := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 33}

	receipt := client.Receipt.Create().
		SetGasUsed(100).
		SetSuccess(true).
		SetBlockHash(b.Hash).
		SetRelayerAddress(relayerAddress[:]).
		SetSourceChainID(b.ChainID).
		SetDestChainID(destChainID).
		SetOffset(offset).
		SetTxHash(txHash).
		SaveX(ctx)

	client.Block.UpdateOne(&b).AddReceipts(receipt).SaveX(ctx)

	return receipt
}

func CreateTestEntClient(t *testing.T) *ent.Client {
	t.Helper()

	opts := []enttest.Option{
		enttest.WithOptions(ent.Log(t.Log)),
		enttest.WithMigrateOptions(migrate.WithGlobalUniqueID(true)),
	}

	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1", opts...)

	return client
}
