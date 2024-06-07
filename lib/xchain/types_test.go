package xchain_test

import (
	"testing"

	"github.com/omni-network/omni/lib/xchain"

	"github.com/stretchr/testify/require"
)

func TestShards(t *testing.T) {
	t.Parallel()

	s := xchain.ShardFinalized0
	require.Equal(t, xchain.ConfFinalized, s.ConfLevel())
	require.False(t, s.Broadcast())

	s = xchain.ShardLatest0
	require.Equal(t, xchain.ConfLatest, s.ConfLevel())
	require.False(t, s.Broadcast())

	s = xchain.ShardBroadcast0
	require.Equal(t, xchain.ConfFinalized, s.ConfLevel())
	require.True(t, s.Broadcast())
}
