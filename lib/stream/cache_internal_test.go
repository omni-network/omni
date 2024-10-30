package stream

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func EnsureFloor[E any](t *testing.T, cache Cache[E]) {
	t.Helper()

	mc, ok := cache.(*mapCache[E])
	require.True(t, ok)

	mc.mu.RLock()
	defer mc.mu.RUnlock()

	for k := range mc.elems {
		require.GreaterOrEqual(t, k, mc.floor)
	}
}
