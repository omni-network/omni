package netconf_test

import (
	"fmt"
	"sort"
	"strings"
	"testing"

	"github.com/omni-network/omni/e2e/app/key"
	"github.com/omni-network/omni/e2e/manifests"
	"github.com/omni-network/omni/e2e/types"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/tutil"

	"github.com/BurntSushi/toml"
	"github.com/stretchr/testify/require"
)

//go:generate go test -golden -run=TestGenSeeds

// TestGenSeeds generates <network>/seeds.txt by loading e2e manifests and parsing seed* p2p_consensus keys.
func TestGenSeeds(t *testing.T) {
	t.Parallel()
	tests := []struct {
		network      netconf.ID
		manifestFunc func() []byte
	}{
		{
			network:      netconf.Testnet,
			manifestFunc: manifests.Testnet,
		},
	}
	for _, test := range tests {
		t.Run(test.network.String(), func(t *testing.T) {
			t.Parallel()
			var manifest types.Manifest
			_, err := toml.Decode(string(test.manifestFunc()), &manifest)
			require.NoError(t, err)

			var peers []string
			for _, node := range sortedKeys(manifest.Keys) {
				if !strings.HasPrefix(node, "seed") {
					continue
				}

				for typ, addr := range manifest.Keys[node] {
					if typ != key.P2PConsensus {
						continue
					}

					peers = append(peers, fmt.Sprintf("%s@%s.%s.omni.network", addr, node, test.network)) // ABCDE123@seed01.staging.omni.network
				}
			}

			seeds := strings.Join(peers, "\n")
			seedsFile := fmt.Sprintf("../%s/seeds.txt", test.network)
			tutil.RequireGoldenBytes(t, []byte(seeds), tutil.WithFilename(seedsFile))
		})
	}
}

func sortedKeys[T any](m map[string]T) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	return keys
}

// TestStrats ensures the netconf.StratX matches ethclient.HeadX.
// Netconf shouldn't import ethclient, so using this test to keep in-sync.
func TestStrats(t *testing.T) {
	t.Parallel()

	require.EqualValues(t, ethclient.HeadLatest, netconf.StratLatest)
	require.EqualValues(t, ethclient.HeadFinalized, netconf.StratFinalized)
}

func TestSeeds(t *testing.T) {
	t.Parallel()

	require.Len(t, netconf.Testnet.Static().Seeds(), 2)
}
