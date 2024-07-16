package emitcache

import (
	"context"
	"math"
	"testing"

	"github.com/omni-network/omni/lib/xchain"

	dbm "github.com/cosmos/cosmos-db"
	"github.com/stretchr/testify/require"
)

func TestEmitCursorCache(t *testing.T) {
	t.Parallel()
	db := dbm.NewMemDB()
	cache, err := newEmitCursorCache(db)
	require.NoError(t, err)
	ctx := context.Background()

	assertContains := func(t *testing.T, height uint64, stream xchain.StreamID, cursor xchain.EmitCursor) {
		t.Helper()
		c, ok, err := cache.Get(ctx, height, stream)
		require.NoError(t, err)
		require.True(t, ok)
		require.Equal(t, cursor, c)

		c, ok, err = cache.AtOrBefore(ctx, height, stream)
		require.NoError(t, err)
		require.True(t, ok)
		require.Equal(t, cursor, c)
	}

	assertHighest := func(t *testing.T, stream xchain.StreamID, cursor xchain.EmitCursor) {
		t.Helper()
		const maxHeight = math.MaxUint64
		c, ok, err := cache.AtOrBefore(ctx, maxHeight, stream)
		require.NoError(t, err)
		require.True(t, ok)
		require.Equal(t, cursor, c)
	}

	assertNotContains := func(t *testing.T, height uint64, stream xchain.StreamID) {
		t.Helper()
		_, ok, err := cache.Get(ctx, height, stream)
		require.NoError(t, err)
		require.False(t, ok)
	}

	set := func(t *testing.T, height uint64, cursor xchain.EmitCursor) {
		t.Helper()
		require.NoError(t, cache.set(ctx, height, cursor))
	}

	trim := func(t *testing.T, chainID uint64, retain uint64) {
		t.Helper()
		require.NoError(t, cache.trimOnce(ctx, chainID, retain))
	}

	stream1 := xchain.StreamID{SourceChainID: 1}
	stream2 := xchain.StreamID{SourceChainID: 2}

	cursor11 := xchain.EmitCursor{StreamID: stream1, MsgOffset: 1}
	cursor12 := xchain.EmitCursor{StreamID: stream1, MsgOffset: 2}
	cursor21 := xchain.EmitCursor{StreamID: stream2, MsgOffset: 1}
	cursor22 := xchain.EmitCursor{StreamID: stream2, MsgOffset: 2}

	stream99 := xchain.StreamID{SourceChainID: 99}

	assertNotContains(t, 1, stream1)
	assertNotContains(t, 2, stream1)
	set(t, 1, cursor12)
	assertContains(t, 1, stream1, cursor12)
	set(t, 1, cursor11) // Update it to 11
	assertContains(t, 1, stream1, cursor11)
	trim(t, stream1.SourceChainID, 99) // Nothing trimmed
	assertContains(t, 1, stream1, cursor11)
	assertHighest(t, stream1, cursor11)

	assertNotContains(t, 2, stream1)
	set(t, 2, cursor12)
	trim(t, stream1.SourceChainID, 99) // Nothing trimmed
	assertContains(t, 2, stream1, cursor12)
	assertContains(t, 1, stream1, cursor11)
	assertHighest(t, stream1, cursor12)

	assertNotContains(t, 1, stream2)
	assertNotContains(t, 2, stream2)
	set(t, 1, cursor21)
	assertContains(t, 1, stream2, cursor21)
	assertHighest(t, stream2, cursor21)

	assertNotContains(t, 2, stream2)
	set(t, 2, cursor22)
	assertContains(t, 2, stream2, cursor22)
	assertContains(t, 1, stream2, cursor21)
	assertHighest(t, stream2, cursor22)

	assertNotContains(t, 1, stream99)
	assertNotContains(t, 2, stream99)

	trim(t, stream1.SourceChainID, 1)
	assertNotContains(t, 1, stream1)
	assertContains(t, 2, stream1, cursor12)
	assertContains(t, 1, stream2, cursor21)
	assertContains(t, 2, stream2, cursor22)

	trim(t, stream1.SourceChainID, 0)
	assertNotContains(t, 2, stream1)
	assertContains(t, 1, stream2, cursor21)
	assertContains(t, 2, stream2, cursor22)

	trim(t, stream2.SourceChainID, 0)
	assertNotContains(t, 2, stream2)
	assertNotContains(t, 1, stream2)
}
