package solvernet

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStatus(t *testing.T) {
	t.Parallel()

	// All statuses can be reached from pending
	require.True(t, StatusPending.ValidTarget(StatusPending))
	require.True(t, StatusPending.ValidTarget(StatusRejected))
	require.True(t, StatusPending.ValidTarget(StatusClosed))
	require.True(t, StatusPending.ValidTarget(StatusFilled))
	require.True(t, StatusPending.ValidTarget(StatusClaimed))

	// Filled is followed by Claimed only
	require.True(t, StatusFilled.ValidTarget(StatusFilled))
	require.True(t, StatusFilled.ValidTarget(StatusClaimed))
	require.False(t, StatusFilled.ValidTarget(StatusPending))
	require.False(t, StatusFilled.ValidTarget(StatusRejected))
	require.False(t, StatusFilled.ValidTarget(StatusClosed))

	// Following states are all end states, they can only reach themselves
	for _, s := range []OrderStatus{
		StatusRejected,
		StatusClosed,
		StatusClaimed,
	} {
		require.False(t, s.ValidTarget(StatusPending))
		require.False(t, s.ValidTarget(StatusFilled))
		require.Equal(t, s == StatusRejected, s.ValidTarget(StatusRejected))
		require.Equal(t, s == StatusClosed, s.ValidTarget(StatusClosed))
		require.Equal(t, s == StatusClaimed, s.ValidTarget(StatusClaimed))
	}
}
