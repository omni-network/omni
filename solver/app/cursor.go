package app

import (
	"context"
	"sync"

	"github.com/omni-network/omni/anchor/anchorinbox"
	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/cast"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	ormv1alpha1 "cosmossdk.io/api/cosmos/orm/v1alpha1"
	"cosmossdk.io/core/store"
	"cosmossdk.io/orm/model/ormdb"
	"cosmossdk.io/orm/types/ormerrors"
	db "github.com/cosmos/cosmos-db"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
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

// GetTxSig returns the tx sig of the cursor for the given chain version, or false, or an error.
func (c *cursors) GetTxSig(ctx context.Context, chainVer xchain.ChainVersion) (solana.Signature, bool, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	cursor, err := c.table.Get(ctx, chainVer.ID, uint32(chainVer.ConfLevel))
	if ormerrors.IsNotFound(err) {
		return solana.Signature{}, false, nil
	} else if err != nil {
		return solana.Signature{}, false, errors.Wrap(err, "get cursor")
	}

	txSig, err := cast.Array64(cursor.GetTxSig())
	if err != nil {
		return solana.Signature{}, false, errors.Wrap(err, "cast tx sig")
	}

	return txSig, true, nil
}

// SetTxSig sets the tx sig of the cursor for the given chain version.
func (c *cursors) SetTxSig(ctx context.Context, chainVer xchain.ChainVersion, sig solana.Signature) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	err := c.table.Save(ctx, &Cursor{
		ChainId:   chainVer.ID,
		ConfLevel: uint32(chainVer.ConfLevel),
		TxSig:     sig[:],
	})
	if err != nil {
		return errors.Wrap(err, "save cursor")
	}

	return nil
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

func maybeBootstrapSVMCursor(ctx context.Context, cl *rpc.Client, cursors *cursors, chainVer xchain.ChainVersion) error {
	// Check if cursor already bootstrapped
	if _, ok, err := cursors.GetTxSig(ctx, chainVer); err != nil {
		return errors.Wrap(err, "get cursor")
	} else if ok {
		return nil
	}

	initSig, err := anchorinbox.GetInitSig(ctx, cl)
	if err != nil {
		return errors.Wrap(err, "get init sig")
	}

	log.Debug(ctx, "Bootstrapping SVM cursor", "sig", initSig)

	return cursors.SetTxSig(ctx, chainVer, initSig)
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
	if height, ok, err := cursors.Get(ctx, chainVer); err != nil {
		return errors.Wrap(err, "get cursor")
	} else if ok {
		// HyperEVM had semi-frequent outages that makes some blocks unavailable.
		// We use this to skip over blocks when necessary.
		const hyperSkipTo uint64 = 9014000
		if chainVer.ID == evmchain.IDHyperEVM && height < hyperSkipTo {
			log.Info(ctx, "Updating HyperEVM cursor to skip height", "height", height, "skip_to", hyperSkipTo)
			err := cursors.Set(ctx, chainVer, hyperSkipTo)
			if err != nil {
				return err
			}
		}

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
