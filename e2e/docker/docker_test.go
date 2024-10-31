package docker_test

import (
	"net"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/omni-network/omni/e2e/docker"
	"github.com/omni-network/omni/e2e/types"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/tutil"

	e2e "github.com/cometbft/cometbft/test/e2e/pkg"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/p2p/enode"

	"github.com/stretchr/testify/require"
)

//go:generate go test . -golden -clean

func TestComposeTemplate(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		tag        string
		isEmpheral bool
		upgrade    string
	}{
		{
			name:       "commit",
			tag:        "7d1ae53",
			isEmpheral: false,
		},
		{
			name:       "empheral_network",
			tag:        "main",
			isEmpheral: true,
			upgrade:    "omniops/halo:v1.0",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			_, ipNet, err := net.ParseCIDR("10.186.73.0/24")
			require.NoError(t, err)

			key, err := crypto.HexToECDSA("59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d")
			require.NoError(t, err)
			en := enode.NewV4(&key.PublicKey, ipNet.IP, 30303, 30303)

			const node0 = "node0"
			dir := t.TempDir()
			testnet := types.Testnet{
				Manifest: types.Manifest{
					PinnedRelayerTag: "v2",
					PinnedMonitorTag: "v3",
				},
				Testnet: &e2e.Testnet{
					Name:       "test",
					IP:         ipNet,
					Dir:        dir,
					Prometheus: true,
					Nodes: []*e2e.Node{{
						Name:       node0,
						Version:    "omniops/halo:" + test.tag,
						InternalIP: ipNet.IP,
						ProxyPort:  8584,
					}},
					UpgradeVersion: "omniops/halo:" + test.tag,
				},
				OmniEVMs: []types.OmniEVM{
					{
						Chain:        types.OmniEVMByNetwork(netconf.Simnet),
						InstanceName: "omni_evm_0",
						AdvertisedIP: ipNet.IP,
						ProxyPort:    8000,
						NodeKey:      key,
						Enode:        en,
						Peers:        []*enode.Node{en},
					},
				},
				AnvilChains: []types.AnvilChain{
					{
						Chain: types.EVMChain{Metadata: evmchain.Metadata{
							ChainID:     99,
							Name:        "mock_rollup",
							BlockPeriod: time.Second,
						}},
						InternalIP: ipNet.IP,
						ProxyPort:  9000,
					},
					{
						Chain: types.EVMChain{Metadata: evmchain.Metadata{
							ChainID:     1,
							Name:        "mock_l1",
							BlockPeriod: time.Hour,
						}},
						InternalIP: ipNet.IP,
						ProxyPort:  9000,
						LoadState:  "path/to/anvil/state.json",
					},
				},
			}

			// If the network is empheral, we use the devnet configuration.
			if test.isEmpheral {
				testnet.Network = netconf.Devnet
			}

			if test.upgrade != "" {
				testnet.UpgradeVersion = test.upgrade
			}

			p := docker.NewProvider(testnet, types.InfrastructureData{}, test.tag)
			require.NoError(t, err)

			require.NoError(t, p.Setup())

			bz, err := os.ReadFile(filepath.Join(dir, "docker-compose.yaml"))
			require.NoError(t, err)

			tutil.RequireGoldenBytes(t, bz)

			t.Run("upgrade", func(t *testing.T) {
				t.Parallel()
				err := docker.ReplaceUpgradeImage(dir, node0)
				if test.upgrade == "" {
					require.Error(t, err)
					return
				}
				require.NoError(t, err)

				bz, err := os.ReadFile(filepath.Join(dir, "docker-compose.yaml"))
				require.NoError(t, err)

				tutil.RequireGoldenBytes(t, bz)
			})
		})
	}
}
