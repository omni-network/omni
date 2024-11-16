package cursor

import (
	"context"

	"github.com/omni-network/omni/lib/errors"

	ormv1alpha1 "cosmossdk.io/api/cosmos/orm/v1alpha1"
	"cosmossdk.io/core/store"
	"cosmossdk.io/orm/model/ormdb"
	db "github.com/cosmos/cosmos-db"
)

// listAll returns all cursors by prefix.
// Results are ordered by primary key ascending: SrcChainId-ConfLevel-DstChainId-AttestOffset.
func listAll(ctx context.Context, db CursorTable) ([]*Cursor, error) {
	iterator, err := db.List(ctx, CursorPrimaryKey{})
	if err != nil {
		return nil, errors.Wrap(err, "listAll cursors")
	}
	defer iterator.Close()

	var cursors []*Cursor
	for iterator.Next() {
		cursor, err := iterator.Value()
		if err != nil {
			return nil, errors.Wrap(err, "cursor value")
		}
		cursors = append(cursors, cursor)
	}

	return cursors, nil
}

func newCursorsTable(db db.DB) (CursorTable, error) {
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
