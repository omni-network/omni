package feature_test

import (
	"testing"

	"github.com/omni-network/omni/lib/feature"

	"github.com/stretchr/testify/require"
)

func TestFlags(t *testing.T) {
	t.Parallel()
	ctx := t.Context()
	// No flags enabled
	require.False(t, feature.FlagFuzzOctane.Enabled(ctx))

	// Single flag enabled
	ctx = feature.WithFlag(ctx, feature.FlagFuzzOctane)
	require.True(t, feature.FlagFuzzOctane.Enabled(ctx))

	// Unknown flags are ignored (and don't overwrite existing)
	ctx = feature.WithFlags(ctx, feature.Flags{"ignore", "us"})
	require.True(t, feature.FlagFuzzOctane.Enabled(ctx))
}
