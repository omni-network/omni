package resolvers

import (
	"context"
	"testing"

	"github.com/omni-network/omni/explorer/db/ent"
	"github.com/omni-network/omni/explorer/db/ent/enttest"
	"github.com/omni-network/omni/explorer/db/ent/migrate"

	"github.com/ethereum/go-ethereum/common"
)

func CreateTestBlock(ctx context.Context, t *testing.T, client *ent.Client) {
	t.Helper()

	sourceChainID := uint64(1234)
	blockHeight := uint64(0)
	blockHashBytes := []byte{1, 3, 23, 111, 27, 45, 98, 103, 94, 55, 1, 3, 23, 111, 27, 45, 98, 103, 94, 55}
	blockHashValue := common.Hash{}
	blockHashValue.SetBytes(blockHashBytes)

	client.Block.Create().
		SetSourceChainID(sourceChainID).
		SetBlockHeight(blockHeight).
		SetBlockHash(blockHashValue.Bytes()).
		SaveX(ctx)
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
