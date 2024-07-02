package vmcompose

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

// SetupDataFixtures returns the test fixture filenames of manifest and infrastructure data files.
// This accesses private types, so it's in the same package as the types.
func SetupDataFixtures(t *testing.T) (string, string) {
	t.Helper()

	// Write manifest to disk
	manifest := `
network = "devnet"
anvil_chains = ["mock_l1"]
multi_omni_evms = true

[node.validator01]
[node.validator02]

[node.seed01]
mode = "seed"

[node.fullnode01]
mode = "full"
`
	manifestFile := filepath.Join(t.TempDir(), "test.toml")
	err := os.WriteFile(manifestFile, []byte(manifest), 0o644)
	require.NoError(t, err)

	const vm1, vm2, vm3, vm4, vm5, vm6 = "vm1", "vm2", "vm3", "vm4", "vm5", "vm6"

	dataJSON := dataJSON{
		NetworkCIDR: "127.0.0.1/24",
		VMs: []vmJSON{
			{Name: vm1, IP: "127.0.0.1"},
			{Name: vm2, IP: "127.0.0.2"},
			{Name: vm3, IP: "127.0.0.3"},
			{Name: vm4, IP: "127.0.0.4"},
			{Name: vm5, IP: "127.0.0.5"},
			{Name: vm6, IP: "127.0.0.6"},
		},
		ServicesByVM: map[string]string{
			"validator01":     vm1,
			"validator01_evm": vm1,

			"validator02":     vm2,
			"validator02_evm": vm2,

			"mock_l1": vm3,
			"relayer": vm3,

			"seed01":     vm4,
			"seed01_evm": vm4,

			"fullnode01":     vm5,
			"fullnode01_evm": vm5,
		},
	}

	// Write raw data json to disk
	bz, err := json.Marshal(dataJSON)
	require.NoError(t, err)
	dataFile := filepath.Join(t.TempDir(), "data.json")
	err = os.WriteFile(dataFile, bz, 0o644)
	require.NoError(t, err)

	return manifestFile, dataFile
}
