package headerdb

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"cosmossdk.io/orm/types/ormerrors"
	dbm "github.com/cosmos/cosmos-db"
	"github.com/stretchr/testify/require"
)

func TestGetSetPrune(t *testing.T) {
	t.Parallel()
	const limit = 2
	db, err := New(dbm.NewMemDB(), limit)
	require.NoError(t, err)

	h1 := fuzzHeader(t, 1, nil)
	f1 := fuzzHeader(t, 1, nil) // Fork at height 1
	h2 := fuzzHeader(t, 22, nil)
	h3 := fuzzHeader(t, 333, nil)
	h4 := fuzzHeader(t, 4444, nil)
	h5 := fuzzHeader(t, 55555, nil)

	ctx := context.Background()

	require.NoError(t, db.Set(ctx, h1))
	require.NoError(t, db.Set(ctx, h1))                               // Noop
	require.ErrorIs(t, db.Set(ctx, f1), ormerrors.UniqueKeyViolation) // Duplicate heights error

	assertExists(ctx, t, db, h1)
	assertNotExists(ctx, t, db, h2, h3, h4, h5)

	require.NoError(t, db.Set(ctx, h2))
	require.NoError(t, db.Set(ctx, h3))
	assertExists(ctx, t, db, h1, h2, h3)
	assertNotExists(ctx, t, db, h4, h5)

	pruned, err := db.maybePrune(ctx)
	require.NoError(t, err)
	require.Equal(t, 1, pruned)
	assertNotExists(ctx, t, db, h1, h4, h5)
	assertExists(ctx, t, db, h2, h3)

	require.NoError(t, db.Set(ctx, h4))
	require.NoError(t, db.Set(ctx, h5))
	assertNotExists(ctx, t, db, h1)
	assertExists(ctx, t, db, h2, h3, h4, h5)

	pruned, err = db.maybePrune(ctx)
	require.NoError(t, err)
	require.Equal(t, 2, pruned)
	assertNotExists(ctx, t, db, h1, h2, h3)
	assertExists(ctx, t, db, h4, h5)
}

func TestAddAndReorg(t *testing.T) {
	t.Parallel()

	db, err := New(dbm.NewMemDB(), -1)
	require.NoError(t, err)

	h1 := fuzzHeader(t, 1, nil)
	h2 := fuzzHeader(t, 2, h1)
	h3 := fuzzHeader(t, 3, h2)
	h4 := fuzzHeader(t, 4, h3)
	h5 := fuzzHeader(t, 5, h4)

	f3 := fuzzHeader(t, 3, h2) // Fork at height 3
	f4 := fuzzHeader(t, 4, f3) // Fork at height 4
	f5 := fuzzHeader(t, 5, f4) // Fork at height 5

	ctx := context.Background()

	require.NoError(t, db.Set(ctx, h1))
	require.NoError(t, db.Set(ctx, h2))
	require.NoError(t, db.Set(ctx, h3))
	require.NoError(t, db.Set(ctx, h4))
	require.NoError(t, db.Set(ctx, h5))
	assertExists(ctx, t, db, h1, h2, h3, h4, h5)
	assertNotExists(ctx, t, db, f3, f4, f5)

	// Reorg to f3, results in h3, h4, h5 being deleted
	deleted, err := db.AddAndReorg(ctx, f3, noFetch)
	require.NoError(t, err)
	require.Equal(t, 3, deleted)
	assertExists(ctx, t, db, h1, h2, f3)
	assertNotExists(ctx, t, db, h3, h4, h5)
	assertNotExists(ctx, t, db, f4, f5)
	// DB contains h1, h2, f3

	// Reorg to h1, results in noop (since h1 is in chain/DB)
	deleted, err = db.AddAndReorg(ctx, h1, noFetch)
	require.NoError(t, err)
	require.Equal(t, 0, deleted)
	assertExists(ctx, t, db, h1, h2, f3)
	assertNotExists(ctx, t, db, h3, h4, h5)
	assertNotExists(ctx, t, db, f4, f5)
	// DB contains h1, h2, f3

	// Reorg to h5, results in noop (since there is gap between f3 and h5)
	deleted, err = db.AddAndReorg(ctx, h5, noFetch)
	require.NoError(t, err)
	require.Equal(t, 0, deleted)
	assertExists(ctx, t, db, h1, h2, f3, h5)
	assertNotExists(ctx, t, db, h3, h4)
	assertNotExists(ctx, t, db, f4, f5)
	// DB contains h1, h2, f3, h5

	// Reorg to h4, results in f3 (and h5) being deleted, and h3 fetched
	deleted, err = db.AddAndReorg(ctx, h4, expectFetch(h3))
	require.NoError(t, err)
	require.Equal(t, 2, deleted)
	assertExists(ctx, t, db, h1, h2, h3, h4)
	assertNotExists(ctx, t, db, h5) // This isn't great, but it's expected for now
	assertNotExists(ctx, t, db, f3, f4, f5)

	require.NoError(t, db.Set(ctx, h5))
	// DB contains h1, h2, h3, h4, h5

	// Reorg to f5, results in h3+h4+f5 deleted, and f3+f4 fetched
	deleted, err = db.AddAndReorg(ctx, f5, expectFetch(f3, f4))
	require.NoError(t, err)
	require.Equal(t, 3, deleted)
	assertExists(ctx, t, db, h1, h2, f3, f4, f5)
	assertNotExists(ctx, t, db, h3, h4, h5)
}

func expectFetch(expected ...*types.Header) func(ctx context.Context, hash common.Hash) (*types.Header, error) {
	fetched := make(map[common.Hash]bool)
	return func(ctx context.Context, hash common.Hash) (*types.Header, error) {
		if fetched[hash] {
			return nil, errors.New("duplicate fetch")
		}
		fetched[hash] = true

		for _, e := range expected {
			if e.Hash() == hash {
				return e, nil
			}
		}

		return nil, errors.New("unexpected fetch")
	}
}

func noFetch(ctx context.Context, hash common.Hash) (*types.Header, error) {
	return nil, errors.New("no fetch expected")
}

func fuzzHeader(t *testing.T, height uint64, parent *types.Header) *types.Header {
	t.Helper()
	fuzzer := ethclient.NewFuzzer(0)

	var header types.Header
	fuzzer.Fuzz(&header)
	header.Number = bi.N(height)
	if parent != nil {
		header.ParentHash = parent.Hash()
	}

	// Ensure that the header is valid.
	bz1, err := json.MarshalIndent(header, "", " ")
	require.NoError(t, err)

	var header2 types.Header
	err = json.Unmarshal(bz1, &header2)
	require.NoError(t, err)

	headersEqual(t, &header, &header2)

	return &header
}

func assertNotExists(ctx context.Context, t *testing.T, db DB, headers ...*types.Header) {
	t.Helper()

	for _, header := range headers {
		_, ok, err := db.ByHash(ctx, header.Hash())
		require.NoError(t, err)
		require.False(t, ok)

		r, ok, err := db.ByHeight(ctx, header.Number.Uint64())
		require.NoError(t, err)
		if ok {
			require.NotEqual(t, header.Hash(), r.Hash())
		}
	}
}

func assertExists(ctx context.Context, t *testing.T, db DB, headers ...*types.Header) {
	t.Helper()

	for _, header := range headers {
		r, ok, err := db.ByHash(ctx, header.Hash())
		require.NoError(t, err)
		require.True(t, ok)
		headersEqual(t, header, r)

		r, ok, err = db.ByHeight(ctx, header.Number.Uint64())
		require.NoError(t, err)
		require.True(t, ok)
		headersEqual(t, header, r)
	}
}

func headersEqual(t *testing.T, h1, h2 *types.Header) {
	t.Helper()
	bz1, err := json.MarshalIndent(h1, "", " ")
	require.NoError(t, err)
	bz2, err := json.MarshalIndent(h2, "", " ")
	require.NoError(t, err)
	require.Equal(t, string(bz1), string(bz2))
}
