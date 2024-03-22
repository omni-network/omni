package docker_test

import (
	"net"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/omni-network/omni/e2e/docker"
	"github.com/omni-network/omni/e2e/types"
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

	tags := []string{"main", "7d1ae53"}

	for _, tag := range tags {
		t.Run("image_tag_"+tag, func(t *testing.T) {
			t.Parallel()
			_, ipNet, err := net.ParseCIDR("10.186.73.0/24")
			require.NoError(t, err)

			key, err := crypto.HexToECDSA("59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d")
			require.NoError(t, err)
			en := enode.NewV4(&key.PublicKey, ipNet.IP, 30303, 30303)

			dir := t.TempDir()
			testnet := types.Testnet{
				Testnet: &e2e.Testnet{
					Name:       "test",
					IP:         ipNet,
					Dir:        dir,
					Prometheus: true,
					Nodes: []*e2e.Node{{
						Name:       "node0",
						Version:    "omniops/halo:" + tag,
						InternalIP: ipNet.IP,
						ProxyPort:  8584,
					}},
				},
				OmniEVMs: []types.OmniEVM{
					{
						Chain:        types.OmniEVMByNetwork(netconf.Simnet),
						InstanceName: "omni_evm_0",
						InternalIP:   ipNet.IP,
						ProxyPort:    8000,
						NodeKey:      key,
						Enode:        en,
						BootNodes:    []*enode.Node{en},
					},
				},
				AnvilChains: []types.AnvilChain{
					{
						Chain:      types.EVMChain{Name: "mock_rollup", ID: 99, BlockPeriod: time.Second},
						InternalIP: ipNet.IP,
						ProxyPort:  9000,
					},
					{
						Chain:      types.EVMChain{Name: "mock_l1", ID: 1, BlockPeriod: time.Hour},
						InternalIP: ipNet.IP,
						ProxyPort:  9000,
						LoadState:  "path/to/anvil/state.json",
					},
				},
			}

			p := docker.NewProvider(testnet, types.InfrastructureData{}, tag)
			require.NoError(t, err)

			require.NoError(t, p.Setup())

			bz, err := os.ReadFile(filepath.Join(dir, "docker-compose.yml"))
			require.NoError(t, err)

			tutil.RequireGoldenBytes(t, bz)
		})
	}
}
