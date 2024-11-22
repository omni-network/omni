package geth_test

import (
	"testing"

	"github.com/omni-network/omni/e2e/app/geth"

	"github.com/stretchr/testify/require"
)

func TestVersions(t *testing.T) {
	t.Parallel()
	require.Len(t, geth.SupportedVersions, 4) // Only support (max) 4 versions
	require.NotContains(t, geth.SupportedVersions, geth.Version)
}
