package emitcache

import (
	"context"
	"sync"
	"time"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/umath"
	"github.com/omni-network/omni/lib/xchain"

	ormv1alpha1 "cosmossdk.io/api/cosmos/orm/v1alpha1"
	"cosmossdk.io/core/store"
	"cosmossdk.io/orm/model/ormdb"
	"cosmossdk.io/orm/model/ormlist"
	"cosmossdk.io/orm/types/ormerrors"
	db "github.com/cosmos/cosmos-db"
	"golang.org/x/sync/errgroup"
)

const (
	// cacheRetain is the number of block heights after which cursors are evicted from the cache.
	cacheRetain = 10_000
	// cacheStartLag is the number of blocks behind latest to start streaming and populating the cache.
	// 128 is the default number of historical block state that geth stores in non-archive mode.
	cacheStartLag = 128
)

type Cache interface {
	Get(ctx context.Context, height uint64, stream xchain.StreamID) (xchain.EmitCursor, bool, error)
	AtOrBefore(ctx context.Context, height uint64, stream xchain.StreamID) (xchain.EmitCursor, bool, error)
}

// Start subscribes the xprovider iot populate the emit cursor cache.
// It returns a cache that will be populated and trimmed asynchronously.
func Start(
	ctx context.Context,
	network netconf.Network,
	xprov xchain.Provider,
	db db.DB,
) (Cache, error) {
	cache, err := newEmitCursorCache(db)
	if err != nil {
		return nil, err
	}

	for _, chain := range network.Chains {
		callback := func(ctx context.Context, block xchain.Block) error {
			// Ignore blocks that are not attested.
			if !block.ShouldAttest(chain.AttestInterval) {
				return nil
			}

			// Update the emit cursor cache for each stream for this height (in parallel).
			var eg errgroup.Group
			for _, stream := range network.StreamsFrom(chain.ID) {
				eg.Go(func() error {
					ref := xchain.EmitRef{Height: &block.BlockHeight}
					emit, _, err := xprov.GetEmittedCursor(ctx, ref, stream)
					if err != nil {
						latest, _ := xprov.ChainVersionHeight(ctx, xchain.ChainVersion{ID: chain.ID, ConfLevel: xchain.ConfLatest})
						return errors.Wrap(err, "get emit cursor", err,
							"stream", network.StreamName(stream),
							"lagging", umath.SubtractOrZero(latest, block.BlockHeight),
						)
					}
					// Populate a zero cursor if not found.

					return cache.set(ctx, block.BlockHeight, emit)
				})
			}

			if err := eg.Wait(); err != nil {
				// Log warn and continue, don't block or retry.
				log.Warn(ctx, "Failed (partially) to populate emit cursor cache (skipping)", err, "height", block.BlockHeight, "chain", chain.Name)
			} else {
				log.Debug(ctx, "Populated emit cursor cache", "height", block.BlockHeight, "chain", chain.Name)
			}

			return nil
		}

		// Figure out where to start streaming from.
		latest, err := xprov.ChainVersionHeight(ctx, xchain.ChainVersion{ID: chain.ID, ConfLevel: xchain.ConfLatest})
		if err != nil {
			return nil, errors.Wrap(err, "latest height", "chain", chain.Name)
		}

		fromHeight := uintSub(latest, cacheStartLag) // Start as far back as cacheStartLag blocks.
		if fromHeight < chain.DeployHeight {
			fromHeight = chain.DeployHeight // But not before chain deploy height.
		}

		req := xchain.ProviderRequest{
			ChainID:   chain.ID,
			Height:    fromHeight,
			ConfLevel: xchain.ConfLatest, // Stream latest height to ensure state is available for querying.
		}

		log.Info(ctx, "Subscribing to xblocks to populate emit cursor cache", "chain", chain.Name, "from_height", fromHeight)

		if err := xprov.StreamAsync(ctx, req, callback); err != nil {
			return nil, err
		}

		// Start a goroutine to trim this chain.
		go cache.trimForever(ctx, network.ID, chain.ID)
	}

	return cache, nil
}

// newEmitCursorCache creates a new emit cursor cache using the provided DB.
func newEmitCursorCache(db db.DB) (*emitCursorCache, error) {
	schema := &ormv1alpha1.ModuleSchemaDescriptor{SchemaFile: []*ormv1alpha1.ModuleSchemaDescriptor_FileEntry{
		{Id: 1, ProtoFileName: File_monitor_xmonitor_emitcache_emitcursor_proto.Path()},
	}}

	storeSvc := dbStoreService{DB: db}

	modDB, err := ormdb.NewModuleDB(schema, ormdb.ModuleDBOptions{KVStoreService: storeSvc})
	if err != nil {
		return nil, errors.Wrap(err, "create ormdb module db")
	}

	dbStore, err := NewEmitcursorStore(modDB)
	if err != nil {
		return nil, errors.Wrap(err, "create store")
	}

	return &emitCursorCache{
		table: dbStore.EmitCursorTable(),
	}, nil
}

// emitCursorCache is a cache of the last 10k emit cursors for each stream.
// This is used to avoid querying chain state (emit cursor) for historical blocks
// as this requires archive nodes.
// Instead we cache the emit cursor of latest blocks, and query the cache for historical blocks
// while monitoring attested stream offsets.
type emitCursorCache struct {
	mu    sync.RWMutex
	table EmitCursorTable
}

// set adds a cursor to the cache for the given height and stream.
// It updates the cursor if it already exists.
func (c *emitCursorCache) set(ctx context.Context, height uint64, cursor xchain.EmitCursor) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	err := c.table.Insert(ctx, &EmitCursor{
		SrcChainId: cursor.SourceChainID,
		Height:     height,
		DstChainId: cursor.DestChainID,
		ShardId:    uint64(cursor.ShardID),
		MsgOffset:  cursor.MsgOffset,
	})
	if errors.Is(err, ormerrors.UniqueKeyViolation) {
		// Cursor already exists, update it
		existing, err := c.table.GetBySrcChainIdDstChainIdShardIdHeight(ctx, cursor.SourceChainID, cursor.DestChainID, uint64(cursor.ShardID), height)
		if err != nil {
			return errors.Wrap(err, "get emit cursor")
		}
		existing.MsgOffset = cursor.MsgOffset
		if err := c.table.Update(ctx, existing); err != nil {
			return errors.Wrap(err, "update emit cursor")
		}

		return nil
	} else if err != nil {
		return errors.Wrap(err, "insert emit cursor")
	}

	return nil
}

// Get returns the emit cursor for the given height and stream.
func (c *emitCursorCache) Get(ctx context.Context, height uint64, stream xchain.StreamID) (xchain.EmitCursor, bool, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	cursor, err := c.table.GetBySrcChainIdDstChainIdShardIdHeight(ctx, stream.SourceChainID, stream.DestChainID, uint64(stream.ShardID), height)
	if ormerrors.IsNotFound(err) {
		return xchain.EmitCursor{}, false, nil
	} else if err != nil {
		return xchain.EmitCursor{}, false, errors.Wrap(err, "get emit cursor")
	}

	return xchain.EmitCursor{
		StreamID:  stream,
		MsgOffset: cursor.GetMsgOffset(),
	}, true, nil
}

// AtOrBefore returns the stream emit cursor at-or-before the given height.
// Only attested heights are populated, so the first cursor at-or-before any
// height will return the correct cursor.
func (c *emitCursorCache) AtOrBefore(ctx context.Context, height uint64, stream xchain.StreamID) (xchain.EmitCursor, bool, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	start := EmitCursorSrcChainIdDstChainIdShardIdHeightIndexKey{}.WithSrcChainIdDstChainIdShardIdHeight(
		stream.SourceChainID,
		stream.DestChainID,
		uint64(stream.ShardID),
		0,
	)

	end := EmitCursorSrcChainIdDstChainIdShardIdHeightIndexKey{}.WithSrcChainIdDstChainIdShardIdHeight(
		stream.SourceChainID,
		stream.DestChainID,
		uint64(stream.ShardID),
		height,
	)

	iter, err := c.table.ListRange(ctx, start, end, ormlist.Reverse(), ormlist.DefaultLimit(1))
	if err != nil {
		return xchain.EmitCursor{}, false, errors.Wrap(err, "list emit cursor")
	}
	defer iter.Close()

	if !iter.Next() {
		return xchain.EmitCursor{}, false, nil // Nothing found
	}

	cursor, err := iter.Value()
	if err != nil {
		return xchain.EmitCursor{}, false, errors.Wrap(err, "emit cursor value")
	}

	if iter.Next() {
		return xchain.EmitCursor{}, false, errors.New("multiple results [BUG]")
	}

	return xchain.EmitCursor{
		StreamID:  stream,
		MsgOffset: cursor.GetMsgOffset(),
	}, true, err
}

func (c *emitCursorCache) trimForever(ctx context.Context, network netconf.ID, chainID uint64) {
	period := time.Hour // Only trim once an hour, since it does table scans.
	if network.IsEphemeral() {
		period = time.Minute * 10 // Make it faster for ephemeral chains
	}

	ticker := time.NewTicker(period)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			err := c.trimOnce(ctx, chainID, cacheRetain)
			if ctx.Err() != nil {
				return // Don't log error on shutdown
			} else if err != nil {
				log.Warn(ctx, "Trim emit cursor cache failed (will retry)", err, "chain", chainID)
			}
		}
	}
}

func (c *emitCursorCache) trimOnce(ctx context.Context, chainID uint64, retain uint64) error {
	t0 := time.Now()

	ids, err := c.detectTrim(ctx, chainID, retain)
	if err != nil {
		return err
	}

	for _, id := range ids {
		c.mu.Lock()
		err := c.table.DeleteBy(ctx, EmitCursorIdIndexKey{}.WithId(id))
		c.mu.Unlock()
		if err != nil {
			return errors.Wrap(err, "delete emit cursor cache")
		}
	}

	var first, last uint64
	if len(ids) > 0 {
		first, last = ids[0], ids[len(ids)-1]
	}

	log.Debug(ctx, "Trimmed emit cursor cache", "chain", chainID, "count", len(ids), "first", first, "last", last, "duration", time.Since(t0))

	return nil
}

// detectTrim returns a list of emit cursor IDs to delete from the cache.
// It returns the IDs of the oldest chainID cursors, excluding the most recent `retainHeight` cursor heights.
func (c *emitCursorCache) detectTrim(ctx context.Context, chainID uint64, retainHeights uint64) ([]uint64, error) {
	// Not taking a read lock since this is a very slow an expensive table scan.
	prefix := EmitCursorSrcChainIdDstChainIdShardIdHeightIndexKey{}.WithSrcChainId(chainID)
	iter, err := c.table.List(ctx, prefix, ormlist.Reverse())
	if err != nil {
		return nil, errors.Wrap(err, "delete emit cursor cache")
	}
	defer iter.Close()

	var ids []uint64
	var prevHeight uint64
	for iter.Next() {
		cursor, err := iter.Value()
		if err != nil {
			return nil, errors.Wrap(err, "emit cursor value")
		}

		if prevHeight == 0 {
			prevHeight = cursor.GetHeight()
		} else if retainHeights > 0 && cursor.GetHeight() < prevHeight {
			// Skip the most recent `retainHeights` cursors.
			prevHeight = cursor.GetHeight()
			retainHeights--
		}

		if retainHeights > 0 {
			continue
		}

		ids = append(ids, cursor.GetId())
	}

	return ids, nil
}

// dbStoreService wraps a cosmos-db instance and provides it via OpenKVStore.
type dbStoreService struct {
	db.DB
}

func (db dbStoreService) OpenKVStore(context.Context) store.KVStore {
	return db.DB
}

func uintSub(a, b uint64) uint64 {
	if a < b {
		return 0
	}

	return a - b
}
