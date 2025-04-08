package types

import (
	"strings"
	"testing"

	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/tokenmeta"

	"github.com/stretchr/testify/require"
)

func TestPublicChains(t *testing.T) {
	t.Parallel()
	for _, meta := range evmchain.All() {
		if meta.NativeToken == tokenmeta.OMNI {
			continue
		}
		if strings.Contains(meta.Name, "mock") {
			continue
		}
		if meta.ChainID == evmchain.IDSepolia {
			continue
		}

		require.NotEmpty(t, PublicRPCByName(meta.Name), "name=%s", meta.Name)

		chain, err := PublicChainByName(meta.Name)
		require.NoError(t, err, "name=%s", meta.Name)
		require.Equal(t, meta, chain.Metadata, "name=%s", meta.Name)
		require.True(t, chain.IsPublic, "name=%s", meta.Name)
		require.Equal(t, allShards, chain.Shards, "name=%s", meta.Name)
	}
}
