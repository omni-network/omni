// Package headerdb provides a header db/cache implementation.
package headerdb

import (
	"context"
	"encoding/json"
	"sync"

	"github.com/omni-network/omni/lib/cast"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/umath"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	ormv1alpha1 "cosmossdk.io/api/cosmos/orm/v1alpha1"
	"cosmossdk.io/core/store"
	"cosmossdk.io/orm/model/ormdb"
	"cosmossdk.io/orm/model/ormlist"
	"cosmossdk.io/orm/types/ormerrors"
	db "github.com/cosmos/cosmos-db"
)

// New returns a new headerdb backed by the given db.
func New(db db.DB) (*DB, error) {
	schema := &ormv1alpha1.ModuleSchemaDescriptor{SchemaFile: []*ormv1alpha1.ModuleSchemaDescriptor_FileEntry{
		{Id: 1, ProtoFileName: File_lib_ethclient_headerdb_headerdb_proto.Path()},
	}}

	storeSvc := dbStoreService{DB: db}

	modDB, err := ormdb.NewModuleDB(schema, ormdb.ModuleDBOptions{KVStoreService: storeSvc})
	if err != nil {
		return nil, errors.Wrap(err, "create ormdb module db")
	}

	dbStore, err := NewHeaderdbStore(modDB)
	if err != nil {
		return nil, errors.Wrap(err, "create store")
	}

	return &DB{
		table: dbStore.HeaderTable(),
	}, nil
}

// DB implements a header db/cache.
// It is reorg aware, via AddAndReorg, so safe to use for both HeaderByHash and HeaderByHeight.
// Without reorg detection, HeaderByHeight could return invalid headers.
// It doesn't preemptively fill gaps, this allows multiple workers to populate the DB.
// When a reorg is however detected, it replaces/fixes any existing invalid headers.
// Fetching valid parent headers (via fetchFunc) is required to identify all reorged-out-headers.
type DB struct {
	mu    sync.Mutex // Lock public write methods, even though cosmos-db supports concurrent access.
	table HeaderTable
}

// set saves the header to the db.
// It is a noop if the header (by hash) already exists.
// It returns ormerrors.UniqueKeyViolation if a header with the same height already exists (reorg).
// Rather use AddAndReorg to handle reorgs.
func (db *DB) set(ctx context.Context, h *types.Header) error {
	hpb, err := toProto(h)
	if err != nil {
		return err
	}

	if _, ok, err := db.ByHash(ctx, h.Hash()); err != nil {
		return err
	} else if ok {
		return nil
	}

	if err := db.table.Insert(ctx, hpb); err != nil {
		return errors.Wrap(err, "save header")
	}

	return nil
}

// ByHash returns the header with the given hash or false if not found.
func (db *DB) ByHash(ctx context.Context, hash common.Hash) (*types.Header, bool, error) {
	hpb, err := db.table.GetByHash(ctx, hash[:])
	if ormerrors.IsNotFound(err) {
		return nil, false, nil
	} else if err != nil {
		return nil, false, errors.Wrap(err, "get header by hash")
	}

	h, err := fromProto(hpb)
	if err != nil {
		return nil, false, err
	}

	return h, true, nil
}

// ByHeight returns the header at the given height or false if not found.
func (db *DB) ByHeight(ctx context.Context, height uint64) (*types.Header, bool, error) {
	hpb, err := db.table.GetByHeight(ctx, height)
	if ormerrors.IsNotFound(err) {
		return nil, false, nil
	} else if err != nil {
		return nil, false, errors.Wrap(err, "get header by height")
	}

	h, err := fromProto(hpb)
	if err != nil {
		return nil, false, err
	}

	return h, true, nil
}

type fetchFunc func(ctx context.Context, hash common.Hash) (*types.Header, error)

// AddAndReorg adds the known header to the chain, and deletes any existing headers not part of known header's chain.
// It replaces all invalid parents of the known head (using fetchfunc).
// It returns the number of headers deleted (reorg depth).
func (db *DB) AddAndReorg(ctx context.Context, known *types.Header, fetch fetchFunc) (int, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	var deleted int

	reorgHeight, fetched, err := db.getReorgHeight(ctx, known, fetch)
	if err != nil {
		return 0, err
	} else if reorgHeight > 0 {
		deleted, err = db.deleteFrom(ctx, reorgHeight)
		if err != nil {
			return 0, err
		}
	}

	// Add new known header
	if err := db.set(ctx, known); err != nil {
		return 0, err
	}

	// Add fetched headers
	for _, h := range fetched {
		if err := db.set(ctx, h); err != nil {
			return 0, err
		}
	}

	return deleted, nil
}

// getReorgHeight returns the height of the reorg, or zero if no reorg detected.
// It also returns any fetched parents if reorg height is lower than known height.
func (db *DB) getReorgHeight(ctx context.Context, known *types.Header, fetch fetchFunc) (uint64, []*types.Header, error) {
	// First check if parent height is in new known chain
	if parentHeight, ok := umath.Subtract(known.Number.Uint64(), 1); ok { //nolint:nestif // Not too bad
		if parent, ok, err := db.ByHeight(ctx, parentHeight); err != nil {
			return 0, nil, err
		} else if ok && parent.Hash() != known.ParentHash {
			// Parent exists, and is not in new known chain
			// Walk up the chain to find reorg height
			parent, err := fetch(ctx, known.ParentHash)
			if err != nil {
				return 0, nil, errors.Wrap(err, "fetch parent")
			}

			reorgHeight, fetched, err := db.getReorgHeight(ctx, parent, fetch)
			if err != nil {
				return 0, nil, err
			} else if reorgHeight == 0 {
				return 0, nil, errors.New("missing reorg height [BUG]")
			}

			return reorgHeight, append(fetched, parent), nil
		}
		// else parent is either in known chain, or it doesn't exist, continue below
	}

	// Check current height
	if current, ok, err := db.ByHeight(ctx, known.Number.Uint64()); err != nil {
		return 0, nil, err
	} else if ok && current.Hash() != known.Hash() {
		// Reorg occurred at this height
		return known.Number.Uint64(), nil, nil
	} else if ok && current.Hash() == known.Hash() {
		// Known header is part of the chain
		return 0, nil, nil
	} // else current height doesn't exist, continue below

	// Check child height
	if child, ok, err := db.ByHeight(ctx, known.Number.Uint64()+1); err != nil {
		return 0, nil, err
	} else if ok && child.ParentHash == known.Hash() {
		// Chain continues on known header
		return 0, nil, nil
	} else if ok && child.ParentHash != known.Hash() {
		// Reorg occurred at this height
		return known.Number.Uint64(), nil, nil
	} // else no child

	// No child, no current height, so reorg unknown, assume no reorg
	return 0, nil, nil
}

// MaybePrune deletes headers ensuring max limit header with highest height.
func (db *DB) MaybePrune(ctx context.Context, limit int) (int, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	// Reverse iterations take a read lock, so we can't delete while iterating.
	iter, err := db.table.List(ctx, HeaderHeightIndexKey{}, ormlist.Reverse())
	if err != nil {
		return 0, errors.Wrap(err, "list header")
	}
	defer iter.Close()

	var toDelete []uint64
	for i := 0; iter.Next(); i++ {
		if i < limit {
			continue
		}

		val, err := iter.Value()
		if err != nil {
			return 0, errors.Wrap(err, "value")
		}

		toDelete = append(toDelete, val.GetId())
	}

	// Close the iterator, which releases the read lock.
	iter.Close()

	for _, id := range toDelete {
		if err := db.table.Delete(ctx, &Header{Id: id}); err != nil {
			return 0, errors.Wrap(err, "delete header")
		}
	}

	return len(toDelete), nil
}

// deleteFrom deletes all headers from the given height and higher.
func (db *DB) deleteFrom(ctx context.Context, height uint64) (int, error) {
	// Not using DeleteRange since it doesn't return the number deleted (which we need).
	iter, err := db.table.ListRange(ctx, HeaderHeightIndexKey{}.WithHeight(height), HeaderHeightIndexKey{})
	if err != nil {
		return 0, errors.Wrap(err, "list from height")
	}
	defer iter.Close()

	// Collect to values to delete after iterating, since writes while iterating isn't supported.
	var toDelete []uint64
	for iter.Next() {
		val, err := iter.Value()
		if err != nil {
			return 0, errors.Wrap(err, "value")
		}

		toDelete = append(toDelete, val.GetId())
	}

	// Close the iterator, which releases the read lock.
	iter.Close()

	for _, id := range toDelete {
		if err := db.table.Delete(ctx, &Header{Id: id}); err != nil {
			return 0, errors.Wrap(err, "delete header")
		}
	}

	return len(toDelete), nil
}

func toProto(h *types.Header) (*Header, error) {
	if h == nil {
		return nil, errors.New("nil header")
	}

	bz, err := json.Marshal(h)
	if err != nil {
		return nil, errors.Wrap(err, "marshal header")
	}

	hash := h.Hash()

	return &Header{
		Height:     h.Number.Uint64(),
		Hash:       hash[:],
		HeaderJson: bz,
	}, nil
}

func fromProto(hpb *Header) (*types.Header, error) {
	if hpb == nil {
		return nil, errors.New("nil header")
	}

	var h types.Header
	if err := json.Unmarshal(hpb.GetHeaderJson(), &h); err != nil {
		return nil, errors.Wrap(err, "unmarshal header")
	}

	if hash, err := cast.EthHash(hpb.GetHash()); err != nil {
		return nil, err
	} else if hash != h.Hash() {
		return nil, errors.New("hash mismatch")
	} else if hpb.GetHeight() != h.Number.Uint64() {
		return nil, errors.New("height mismatch")
	}

	return &h, nil
}

// dbStoreService wraps a cosmos-db instance and provides it via OpenKVStore.
type dbStoreService struct {
	db.DB
}

func (db dbStoreService) OpenKVStore(context.Context) store.KVStore {
	return db.DB
}
