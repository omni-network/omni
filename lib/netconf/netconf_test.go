package netconf_test

import (
	"bytes"
	"context"
	"encoding/hex"
	"flag"
	"fmt"
	"net"
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

	"github.com/stretchr/testify/require"
)

func TestStaticNetwork(t *testing.T) {
	t.Parallel()
	for _, chain := range netconf.All() {
		static := chain.Static()
		require.Equal(t, chain, static.Network)
	}
}

//go:generate go test -golden -run=TestGenConsSeeds

// TestGenConsSeeds generates <network>/consensus-seeds.txt by loading e2e manifests and parsing seed* p2p_consensus keys.
func TestGenConsSeeds(t *testing.T) {
	t.Parallel()
	tests := []struct {
		network      netconf.ID
		manifestFunc func() (types.Manifest, error)
	}{
		{
			network:      netconf.Omega,
			manifestFunc: manifests.Omega,
		},
		{
			network:      netconf.Staging,
			manifestFunc: manifests.Staging,
		},
		{
			network:      netconf.Mainnet,
			manifestFunc: manifests.Mainnet,
		},
	}
	for _, test := range tests {
		t.Run(test.network.String(), func(t *testing.T) {
			t.Parallel()
			manifest, err := test.manifestFunc()
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

//go:generate go test -golden -run=TestGenConsArchives

// TestGenConsArchives generates <network>/consensus-archives.txt by loading e2e manifests and parsing archive* p2p_consensus keys.
func TestGenConsArchives(t *testing.T) {
	t.Parallel()
	tests := []struct {
		network      netconf.ID
		manifestFunc func() (types.Manifest, error)
	}{
		{
			network:      netconf.Omega,
			manifestFunc: manifests.Omega,
		},
		{
			network:      netconf.Staging,
			manifestFunc: manifests.Staging,
		},
		{
			network:      netconf.Mainnet,
			manifestFunc: manifests.Mainnet,
		},
	}
	for _, test := range tests {
		t.Run(test.network.String(), func(t *testing.T) {
			t.Parallel()
			manifest, err := test.manifestFunc()
			require.NoError(t, err)

			var peers []string
			for _, node := range sortedKeys(manifest.Keys) {
				if !isArchiveNode(test.network, node) {
					continue
				}

				for typ, addr := range manifest.Keys[node] {
					if typ != key.P2PConsensus {
						continue
					}
					addr = strings.ToLower(addr) // CometBFT P2P IDs are lowercase.

					peers = append(peers, fmt.Sprintf("%s@%s.%s.omni.network:26656", addr, node, test.network)) // abcde123@archive01.staging.omni.network:26656
				}
			}

			seeds := strings.Join(peers, "\n")
			seedsFile := fmt.Sprintf("../%s/consensus-archives.txt", test.network)
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
		manifestFunc func() (types.Manifest, error)
	}{
		{
			network:      netconf.Omega,
			manifestFunc: manifests.Omega,
		},
		{
			network:      netconf.Staging,
			manifestFunc: manifests.Staging,
		},
		{
			network:      netconf.Mainnet,
			manifestFunc: manifests.Mainnet,
		},
	}
	for _, test := range tests {
		t.Run(test.network.String(), func(t *testing.T) {
			t.Parallel()
			ctx := context.Background()
			manifest, err := test.manifestFunc()
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
					tutil.RequireNoError(t, err)

					stdPrivKey, err := key.ECDSA()
					require.NoError(t, err)

					pubkey64, err := k1util.PubKeyToBytes64(&stdPrivKey.PublicKey)
					require.NoError(t, err)
					pubkeyHex := hex.EncodeToString(pubkey64)
					nodeName := strings.TrimSuffix(node, "_evm")

					// Convert hostname to IP since geth doesn't support DNS bootstrap nodes anymore.
					// TODO(corver): Revert once fixed https://github.com/ethereum/go-ethereum/issues/31208
					hostname := fmt.Sprintf("%s.%s.omni.network", nodeName, test.network)
					ips, err := net.LookupIP(hostname)
					require.NoError(t, err)
					require.NotEmpty(t, ips)
					sort.Slice(ips, func(i, j int) bool {
						return bytes.Compare(ips[i], ips[j]) < 0
					})

					//nolint:nosprintfhostport // Not important
					peers = append(peers, fmt.Sprintf("enode://%s@%s:30303", pubkeyHex, ips[0]))
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

	omegaAddrs, err := contracts.GetAddresses(context.Background(), netconf.Omega)
	require.NoError(t, err)

	mainnetAddrs, err := contracts.GetAddresses(context.Background(), netconf.Mainnet)
	require.NoError(t, err)

	// test that hardcoded address in netconf match lib/contract addresses
	for _, deployment := range netconf.Omega.Static().Portals {
		require.Equal(t, omegaAddrs.Portal, deployment.Address)
	}

	require.Equal(t, omegaAddrs.AVS, netconf.Omega.Static().AVSContractAddress)
	require.Equal(t, mainnetAddrs.AVS, netconf.Mainnet.Static().AVSContractAddress)
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

// isArchiveNode returns true if the node should be added to archive node static config.
func isArchiveNode(network netconf.ID, node string) bool {
	// All "seed*" nodes in the manifest are seed noes
	if strings.HasPrefix(node, "archive") {
		return true
	}

	// In staging, we only have 1 archive, it the fullnode as well
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
