package types_test

import (
	"strings"
	"testing"

	"github.com/omni-network/omni/e2e/manifests"
	"github.com/omni-network/omni/e2e/types"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/tokens"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/stretchr/testify/require"
)

var allShards = []xchain.ShardID{xchain.ShardFinalized0, xchain.ShardLatest0}

func TestPublicChains(t *testing.T) {
	t.Parallel()

	for _, network := range netconf.All() {
		if network == netconf.Devnet || network == netconf.Simnet {
			continue
		}

		chains, err := manifests.EVMChains(network)
		require.NoError(t, err)

		for _, meta := range chains {
			if meta.NativeToken == tokens.NOM {
				continue
			}
			if strings.Contains(meta.Name, "mock") {
				continue
			}
			if meta.ChainID == evmchain.IDSepolia {
				continue
			}

			require.NotEmpty(t, types.PublicRPCByName(meta.Name), "name=%s", meta.Name)

			chain, err := types.PublicChainByName(meta.Name)
			require.NoError(t, err, "name=%s", meta.Name)
			require.Equal(t, meta, chain.Metadata, "name=%s", meta.Name)
			require.True(t, chain.IsPublic, "name=%s", meta.Name)
			require.Equal(t, allShards, chain.Shards, "name=%s", meta.Name)
		}
	}
}
