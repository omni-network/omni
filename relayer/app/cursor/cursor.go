package cursor

import (
	"context"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/xchain"

	ormv1alpha1 "cosmossdk.io/api/cosmos/orm/v1alpha1"
	"cosmossdk.io/core/store"
	"cosmossdk.io/orm/model/ormdb"
	db "github.com/cosmos/cosmos-db"
)

func (c *Cursor) StreamID(shardID xchain.ShardID) xchain.StreamID {
	return xchain.StreamID{
		SourceChainID: c.GetSrcChainId(),
		DestChainID:   c.GetDstChainId(),
		ShardID:       shardID,
	}
}

func NewCursorsTable(db db.DB) (CursorTable, error) {
	schema := &ormv1alpha1.ModuleSchemaDescriptor{SchemaFile: []*ormv1alpha1.ModuleSchemaDescriptor_FileEntry{
		{Id: 1, ProtoFileName: File_relayer_app_cursor_cursors_proto.Path()},
	}}

	modDB, err := ormdb.NewModuleDB(schema, ormdb.ModuleDBOptions{KVStoreService: dbStoreService{db}})
	if err != nil {
		return nil, errors.Wrap(err, "create ormdb module db")
	}

	dbStore, err := NewCursorsStore(modDB)
	if err != nil {
		return nil, errors.Wrap(err, "create store")
	}

	return dbStore.CursorTable(), nil
}

type dbStoreService struct {
	db.DB
}

func (db dbStoreService) OpenKVStore(context.Context) store.KVStore {
	return db.DB
}
