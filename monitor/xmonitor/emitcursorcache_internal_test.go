package xmonitor

import (
	"math"
	"testing"

	"github.com/omni-network/omni/lib/xchain"

	"github.com/stretchr/testify/require"
)

func TestEmitCursorCache(t *testing.T) {
	t.Parallel()
	cache := newEmitCursorCache()

	assertContains := func(t *testing.T, height uint64, stream xchain.StreamID, cursor xchain.EmitCursor) {
		t.Helper()
		c, ok := cache.Get(height, stream)
		require.True(t, ok)
		require.Equal(t, cursor, c)

		c, ok = cache.AtOrBefore(height, stream)
		require.True(t, ok)
		require.Equal(t, cursor, c)
	}

	assertHighest := func(t *testing.T, stream xchain.StreamID, cursor xchain.EmitCursor) {
		t.Helper()
		const maxHeight = math.MaxUint64
		c, ok := cache.AtOrBefore(maxHeight, stream)
		require.True(t, ok)
		require.Equal(t, cursor, c)
	}

	assertNotContains := func(t *testing.T, height uint64, stream xchain.StreamID) {
		t.Helper()
		_, ok := cache.Get(height, stream)
		require.False(t, ok)
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
	cache.set(1, stream1, cursor11)
	assertContains(t, 1, stream1, cursor11)
	cache.trim(0, stream1) // Nothing trimmed
	assertContains(t, 1, stream1, cursor11)
	assertHighest(t, stream1, cursor11)

	assertNotContains(t, 2, stream1)
	cache.set(2, stream1, cursor12)
	cache.trim(0, stream1) // Nothing trimmed
	assertContains(t, 2, stream1, cursor12)
	assertContains(t, 1, stream1, cursor11)
	assertHighest(t, stream1, cursor12)

	assertNotContains(t, 1, stream2)
	assertNotContains(t, 2, stream2)
	cache.set(1, stream2, cursor21)
	assertContains(t, 1, stream2, cursor21)
	assertHighest(t, stream2, cursor21)

	assertNotContains(t, 2, stream2)
	cache.set(2, stream2, cursor22)
	assertContains(t, 2, stream2, cursor22)
	assertContains(t, 1, stream2, cursor21)
	assertHighest(t, stream2, cursor22)

	assertNotContains(t, 1, stream99)
	assertNotContains(t, 2, stream99)

	cache.trim(1, stream1)
	assertNotContains(t, 1, stream1)
	assertContains(t, 2, stream1, cursor12)

	cache.trim(2, stream1)
	assertNotContains(t, 2, stream1)

	cache.trim(2, stream2)
	assertNotContains(t, 1, stream2)
	assertNotContains(t, 2, stream2)

	// ensure internal state
	require.Empty(t, cache.cursors)

	require.Len(t, cache.heights, 2)
	require.Empty(t, cache.heights[stream1])
	require.Empty(t, cache.heights[stream2])
}
