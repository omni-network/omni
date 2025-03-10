package app

import (
	"context"
	"sync"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	ormv1alpha1 "cosmossdk.io/api/cosmos/orm/v1alpha1"
	"cosmossdk.io/core/store"
	"cosmossdk.io/orm/model/ormdb"
	"cosmossdk.io/orm/types/ormerrors"
	db "github.com/cosmos/cosmos-db"
)

// newSolverDB returns a new DB backend based on the given directory
// or in-memory-based if the directory is empty.
func newSolverDB(dbDir string) (db.DB, error) {
	resp := db.DB(db.NewMemDB())
	if dbDir != "" {
		var err error
		resp, err = db.NewGoLevelDB("solver", dbDir, nil)
		if err != nil {
			return nil, errors.Wrap(err, "new golevel db")
		}
	}

	return resp, nil
}

func newCursors(db db.DB) (*cursors, error) {
	schema := &ormv1alpha1.ModuleSchemaDescriptor{SchemaFile: []*ormv1alpha1.ModuleSchemaDescriptor_FileEntry{
		{Id: 1, ProtoFileName: File_solver_app_solver_proto.Path()},
	}}

	storeSvc := dbStoreService{DB: db}

	modDB, err := ormdb.NewModuleDB(schema, ormdb.ModuleDBOptions{KVStoreService: storeSvc})
	if err != nil {
		return nil, errors.Wrap(err, "create ormdb module db")
	}

	dbStore, err := NewSolverStore(modDB)
	if err != nil {
		return nil, errors.Wrap(err, "create store")
	}

	return &cursors{
		table: dbStore.CursorTable(),
	}, nil
}

// cursors provides a thread-safe persisted cursor store.
type cursors struct {
	mu    sync.Mutex
	table CursorTable
}

// Get returns the block height of the cursor for the given chain version, or false, or an error.
func (c *cursors) Get(ctx context.Context, chainVer xchain.ChainVersion) (uint64, bool, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	cursor, err := c.table.Get(ctx, chainVer.ID, uint32(chainVer.ConfLevel))
	if ormerrors.IsNotFound(err) {
		return 0, false, nil
	} else if err != nil {
		return 0, false, errors.Wrap(err, "get cursor")
	}

	return cursor.GetBlockHeight(), true, nil
}

// Set sets the block height of the cursor for the given chain version.
func (c *cursors) Set(ctx context.Context, chainVer xchain.ChainVersion, height uint64) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	err := c.table.Save(ctx, &Cursor{
		ChainId:     chainVer.ID,
		ConfLevel:   uint32(chainVer.ConfLevel),
		BlockHeight: height,
	})
	if err != nil {
		return errors.Wrap(err, "save cursor")
	}

	return nil
}

// dbStoreService wraps a cosmos-db instance and provides it via OpenKVStore.
type dbStoreService struct {
	db.DB
}

func (db dbStoreService) OpenKVStore(context.Context) store.KVStore {
	return db.DB
}

// maybeBootstrapCursor bootstraps a cursor if not present.
// It either uses the height of an existing cursor (of same chain), or the deploy height of the inbox.
func maybeBootstrapCursor(
	ctx context.Context,
	inbox *bindings.SolverNetInbox,
	cursors *cursors,
	chainVer xchain.ChainVersion,
) error {
	// Check if cursor already bootstrapped
	if _, ok, err := cursors.Get(ctx, chainVer); err != nil {
		return errors.Wrap(err, "get cursor")
	} else if ok {
		return nil
	}

	// Try to use existing cursor for same chain
	for _, confLevel := range xchain.AllConfLevels() {
		height, ok, err := cursors.Get(ctx, xchain.NewChainVersion(chainVer.ID, confLevel))
		if err != nil {
			return err
		} else if !ok {
			continue
		}

		log.Info(ctx, "Bootstrap cursor from existing", "existing", confLevel, "height", height)

		if err := cursors.Set(ctx, chainVer, height); err != nil {
			return errors.Wrap(err, "set cursor")
		}

		return nil
	}

	// No existing cursor found, bootstrap from contract deploy height
	deployHeight, err := inbox.DeployedAt(&bind.CallOpts{Context: ctx})
	if err != nil {
		return errors.Wrap(err, "deployed at")
	}

	log.Info(ctx, "Bootstrap cursor from inbox deploy height", "height", deployHeight)

	return cursors.Set(ctx, chainVer, deployHeight.Uint64())
}
