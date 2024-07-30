package netconf_test

import (
	"context"
	"encoding/hex"
	"flag"
	"fmt"
	"sort"
	"strings"
	"testing"

	"github.com/omni-network/omni/e2e/app/key"
	"github.com/omni-network/omni/e2e/manifests"
	"github.com/omni-network/omni/e2e/types"
	"github.com/omni-network/omni/lib/contracts"
	"github.com/omni-network/omni/lib/k1util"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/tutil"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/ethereum/go-ethereum/p2p/enode"

	"github.com/BurntSushi/toml"
	"github.com/stretchr/testify/require"
)

//go:generate go test -golden -run=TestGenConsSeeds

// TestGenConsSeeds generates <network>/consensus-seeds.txt by loading e2e manifests and parsing seed* p2p_consensus keys.
func TestGenConsSeeds(t *testing.T) {
	t.Parallel()
	tests := []struct {
		network      netconf.ID
		manifestFunc func() []byte
	}{
		{
			network:      netconf.Omega,
			manifestFunc: manifests.Omega,
		},
		{
			network:      netconf.Staging,
			manifestFunc: manifests.Staging,
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
				if !isSeedNode(test.network, node) {
					continue
				}

				for typ, addr := range manifest.Keys[node] {
					if typ != key.P2PConsensus {
						continue
					}
					addr = strings.ToLower(addr) // CometBFT P2P IDs are lowercase.

					peers = append(peers, fmt.Sprintf("%s@%s.%s.omni.network:26656", addr, node, test.network)) // abcde123@seed01.staging.omni.network:26656
				}
			}

			seeds := strings.Join(peers, "\n")
			seedsFile := fmt.Sprintf("../%s/consensus-seeds.txt", test.network)
			tutil.RequireGoldenBytes(t, []byte(seeds), tutil.WithFilename(seedsFile))
		})
	}
}

var genExecutionSeeds = flag.Bool("gen-execution-seeds", false, "Enable to generate execution-seeds.txt. Note this requires GCP secret manager read-access")

//go:generate go test -golden -gen-execution-seeds -run=TestGenExecutionSeeds

func TestGenExecutionSeeds(t *testing.T) {
	t.Parallel()
	if !*genExecutionSeeds {
		t.Skip("Skipping since --gen-execution-seeds=false")
		return
	}

	tests := []struct {
		network      netconf.ID
		manifestFunc func() []byte
	}{
		{
			network:      netconf.Omega,
			manifestFunc: manifests.Omega,
		},
		{
			network:      netconf.Staging,
			manifestFunc: manifests.Staging,
		},
	}
	for _, test := range tests {
		t.Run(test.network.String(), func(t *testing.T) {
			t.Parallel()
			ctx := context.Background()
			var manifest types.Manifest
			_, err := toml.Decode(string(test.manifestFunc()), &manifest)
			require.NoError(t, err)

			var peers []string
			for _, node := range sortedKeys(manifest.Keys) {
				if !isSeedNode(test.network, node) {
					continue
				}

				for typ, addr := range manifest.Keys[node] {
					if typ != key.P2PExecution {
						continue
					}

					key, err := key.Download(ctx, test.network, node, typ, addr)
					require.NoError(t, err)

					stdPrivKey, err := key.ECDSA()
					require.NoError(t, err)

					pubkey64, err := k1util.PubKeyToBytes64(&stdPrivKey.PublicKey)
					require.NoError(t, err)
					pubkeyHex := hex.EncodeToString(pubkey64)
					nodeName := strings.TrimSuffix(node, "_evm")

					peers = append(peers, fmt.Sprintf("enode://%s@%s.%s.omni.network:30303", pubkeyHex, nodeName, test.network))
				}
			}

			seeds := strings.Join(peers, "\n")
			seedsFile := fmt.Sprintf("../%s/execution-seeds.txt", test.network)
			tutil.RequireGoldenBytes(t, []byte(seeds), tutil.WithFilename(seedsFile))
		})
	}
}

func TestConsensusSeeds(t *testing.T) {
	t.Parallel()

	seeds := netconf.Omega.Static().ConsensusSeeds()
	require.Len(t, seeds, 2)
	for _, seed := range seeds {
		require.NotEmpty(t, seed)
		parts := strings.Split(seed, "@")
		require.Len(t, parts, 2)
		require.NotEmpty(t, parts[0])
		require.NotEmpty(t, parts[1])
		t.Logf("Consensus Seed: %s", seed)
	}
}

func TestExecutionSeeds(t *testing.T) {
	t.Skip("testnet shutdown at the moment")
	t.Parallel()

	seeds := netconf.Omega.Static().ExecutionSeeds()
	require.Len(t, seeds, 2)
	for _, seed := range seeds {
		node, err := enode.ParseV4(seed)
		require.NoError(t, err)

		require.EqualValues(t, 30303, node.TCP())
		require.EqualValues(t, 30303, node.UDP())
		t.Logf("Seed IP: %s: %s", node.IP(), seed)
		require.NotEmpty(t, node.IP())
	}
}

func TestConfLevels(t *testing.T) {
	t.Parallel()

	chain := netconf.Chain{
		Shards: []xchain.ShardID{xchain.ShardBroadcast0, xchain.ShardLatest0},
	}
	require.Len(t, chain.ConfLevels(), 2)
	require.EqualValues(t, []xchain.ConfLevel{xchain.ConfLatest, xchain.ConfFinalized}, chain.ConfLevels())
}

func TestAddrs(t *testing.T) {
	t.Parallel()

	// test that hardcoded address in netconf match lib/contract addresses
	for _, deployment := range netconf.Omega.Static().Portals {
		require.Equal(t, contracts.Portal(netconf.Omega), deployment.Address)
	}

	require.Equal(t, contracts.AVS(netconf.Omega), netconf.Omega.Static().AVSContractAddress)
	require.Equal(t, contracts.AVS(netconf.Mainnet), netconf.Mainnet.Static().AVSContractAddress)
}

// isSeedNode returns true if the node should be added to seed node static config.
func isSeedNode(network netconf.ID, node string) bool {
	// All "seed*" nodes in the manifest are seed noes
	if strings.HasPrefix(node, "seed") {
		return true
	}

	// In staging, we only have 1 seednode, so add the fullnode as well
	if network == netconf.Staging && strings.HasPrefix(node, "full") {
		return true
	}

	return false
}

func sortedKeys[T any](m map[string]T) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	return keys
}
