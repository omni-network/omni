package docker_test

import (
	"net"
	"os"
	"path/filepath"
	"testing"

	"github.com/omni-network/omni/test/e2e/docker"
	"github.com/omni-network/omni/test/e2e/types"
	"github.com/omni-network/omni/test/tutil"

	e2e "github.com/cometbft/cometbft/test/e2e/pkg"

	"github.com/stretchr/testify/require"
)

//go:generate go test . -clean -update

func TestComposeTemplate(t *testing.T) {
	t.Parallel()

	_, ipNet, err := net.ParseCIDR("10.186.73.0/24")
	require.NoError(t, err)

	dir := t.TempDir()
	testnet := types.Testnet{
		Testnet: &e2e.Testnet{
			Name: "test",
			IP:   ipNet,
			Dir:  dir,
			Nodes: []*e2e.Node{{
				Name:       "node0",
				Version:    "v0.0.0",
				InternalIP: ipNet.IP,
				ProxyPort:  8584,
			}},
		},
		OmniEVMs: []types.OmniEVM{
			{
				Chain:        types.ChainOmniEVM,
				InstanceName: "omni_evm_0",
				InternalIP:   ipNet.IP,
				ProxyPort:    8000,
			},
		},
		AnvilChains: []types.AnvilChain{
			{
				Chain:      types.EVMChain{Name: "chain_a", ID: 99},
				InternalIP: ipNet.IP,
				ProxyPort:  9000,
			},
		},
	}

	p := docker.NewProvider(testnet, types.InfrastructureData{})
	require.NoError(t, err)

	require.NoError(t, p.Setup())

	bz, err := os.ReadFile(filepath.Join(dir, "docker-compose.yml"))
	require.NoError(t, err)

	tutil.RequireGoldenBytes(t, bz)
}
