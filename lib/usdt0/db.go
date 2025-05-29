//nolint:unused // WIP
package usdt0

import (
	"context"
	"sync"

	"github.com/omni-network/omni/lib/errors"

	ormv1alpha1 "cosmossdk.io/api/cosmos/orm/v1alpha1"
	"cosmossdk.io/core/store"
	"cosmossdk.io/orm/model/ormdb"
	db "github.com/cosmos/cosmos-db"
)

// DB provides access to USDT0 storage.
type DB struct {
	mu       sync.Mutex
	msgTable MsgSendUSDT0Table
}

// NewDB returns a new USDT0 database instance.
func NewDB(db db.DB) (*DB, error) {
	schema := &ormv1alpha1.ModuleSchemaDescriptor{SchemaFile: []*ormv1alpha1.ModuleSchemaDescriptor_FileEntry{
		{Id: 1, ProtoFileName: File_lib_usdt0_db_proto.Path()},
	}}

	storeSvc := dbStoreService{DB: db}

	modDB, err := ormdb.NewModuleDB(schema, ormdb.ModuleDBOptions{KVStoreService: storeSvc})
	if err != nil {
		return nil, errors.Wrap(err, "create ormdb module db")
	}

	dbStore, err := NewDbStore(modDB)
	if err != nil {
		return nil, errors.Wrap(err, "create store")
	}

	return &DB{
		msgTable: dbStore.MsgSendUSDT0Table(),
	}, nil
}

type dbStoreService struct {
	db.DB
}

func (db dbStoreService) OpenKVStore(context.Context) store.KVStore {
	return db
}
