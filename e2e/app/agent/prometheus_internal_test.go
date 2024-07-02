package agent

import (
	"context"
	"testing"

	"github.com/omni-network/omni/e2e/types"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/tutil"

	e2e "github.com/cometbft/cometbft/test/e2e/pkg"

	"github.com/stretchr/testify/require"
)

//go:generate go test . -golden -clean

func TestPromGen(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name         string
		network      netconf.ID
		nodes        []string
		newNodes     []string
		geths        []string
		newGeths     []string
		newServices  []string
		hostname     string
		agentSecrets bool
	}{
		{
			name:         "manifest1",
			network:      netconf.Devnet,
			nodes:        []string{"validator01", "validator02"},
			hostname:     "localhost",
			newNodes:     []string{"validator01"},
			geths:        []string{"omni_evm"},
			newGeths:     []string{"omni_evm"},
			newServices:  nil,
			agentSecrets: false,
		},
		{
			name:         "manifest2",
			network:      netconf.Staging,
			nodes:        []string{"validator01", "validator02", "fullnode03"},
			hostname:     "vm",
			newNodes:     []string{"fullnode04"},
			geths:        []string{"validator01_evm", "validator02_evm", "validator03_evm"},
			newGeths:     []string{"fullnode04_evm"},
			newServices:  []string{"relayer"},
			agentSecrets: true,
		},
		{
			name:         "manifest3",
			network:      netconf.Devnet,
			nodes:        []string{"validator01", "validator02"},
			hostname:     "localhost",
			newNodes:     []string{"validator01"},
			newServices:  []string{"monitor"},
			agentSecrets: false,
		},
		{
			name:         "manifest4",
			network:      netconf.Staging,
			nodes:        []string{"validator01", "validator02", "fullnode03"},
			hostname:     "vm",
			newNodes:     []string{"fullnode04"},
			newServices:  []string{"relayer", "monitor"},
			agentSecrets: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			ctx := context.Background()

			var nodes []*e2e.Node
			for _, name := range test.nodes {
				nodes = append(nodes, &e2e.Node{Name: name})
			}

			var geths []types.OmniEVM
			for _, name := range test.geths {
				geths = append(geths, types.OmniEVM{InstanceName: name})
			}

			testnet := types.Testnet{
				Network: test.network,
				Testnet: &e2e.Testnet{
					Name:  test.name,
					Nodes: nodes,
				},
				OmniEVMs: geths,
			}

			var agentSecrets Secrets
			if test.agentSecrets {
				agentSecrets = Secrets{
					URL:  "https://grafana.com",
					User: "admin",
					Pass: "password",
				}
			}

			cfg1, err := genPromConfig(ctx, testnet, agentSecrets, test.hostname)
			require.NoError(t, err)

			services := make(map[string]bool)
			for _, newService := range test.newServices {
				services[newService] = true
			}

			cfg2 := ConfigForHost(cfg1, test.hostname+"-2", test.newNodes, test.newGeths, services)

			t.Run("gen", func(t *testing.T) {
				t.Parallel()
				tutil.RequireGoldenBytes(t, cfg1)
			})

			t.Run("update", func(t *testing.T) {
				t.Parallel()
				tutil.RequireGoldenBytes(t, cfg2)
			})
		})
	}
}
