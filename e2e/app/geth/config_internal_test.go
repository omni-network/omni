package geth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/omni-network/omni/lib/tutil"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/p2p/enode"

	"github.com/stretchr/testify/require"
)

//go:generate go test . -golden -clean

func TestWriteConfigTOML(t *testing.T) {
	t.Parallel()

	testKey, _ := crypto.HexToECDSA("45a915e4d060149eb4365960e6a7a45f334393093061116b197e3240065ff2d8")
	node1 := enode.NewV4(&testKey.PublicKey, net.IP{127, 0, 0, 1}, 1, 1)
	node2 := enode.NewV4(&testKey.PublicKey, net.IP{127, 0, 0, 2}, 2, 2)

	tests := map[string]bool{
		"archive": true,
		"full":    false,
	}
	for name, isArchive := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			data := Config{
				Moniker:      name,
				BootNodes:    []*enode.Node{node1},
				TrustedNodes: []*enode.Node{node1, node2},
				ChainID:      15651,
				IsArchive:    isArchive,
			}

			tempFile := filepath.Join(t.TempDir(), name+".toml")

			err := writeConfigTOML(data, tempFile)
			require.NoError(t, err)

			bz, err := os.ReadFile(tempFile)
			require.NoError(t, err)

			tutil.RequireGoldenBytes(t, bz)

			// Compare our generated config against the output of `geth dumpconfig` with this as the base config.
			// Geth does some custom config parsing/sanitizing/updating of the config, so we ensure our config doesn't
			// get silently updated by geth.
			// See https://github.com/ethereum/go-ethereum/blob/master/cmd/utils/flags.go#L1640
			result := gethDumpConfigToml(t, MakeGethConfig(data))
			require.Equal(t, string(bz), string(result))
		})
	}
}

// TestGethVersion checks if the geth version is up to date.
func TestGethVersion(t *testing.T) {
	t.Parallel()

	resp, err := http.Get("https://api.github.com/repos/ethereum/go-ethereum/releases/latest") //nolint:noctx // Test is ok
	require.NoError(t, err)
	defer resp.Body.Close()

	bz, err := io.ReadAll(resp.Body)
	require.NoError(t, err)

	var v version
	require.NoError(t, json.Unmarshal(bz, &v))

	require.Equal(t, Version, v.TagName, "A new version of geth has been released, update `geth.Version`")
}

type version struct {
	TagName string `json:"tag_name"`
}

// gethDumpConfigToml executes `geth dumpconfig` using the provided base config and
// returns the resulting toml config file content.
func gethDumpConfigToml(t *testing.T, baseCfg FullConfig) []byte {
	t.Helper()

	bz, err := tomlSettings.Marshal(baseCfg)
	require.NoError(t, err)

	baseFile := filepath.Join(t.TempDir(), "base.toml")
	err = os.WriteFile(baseFile, bz, 0o644)
	require.NoError(t, err)

	var stdout, stderr bytes.Buffer
	cmd := exec.Command("docker", "run",
		fmt.Sprintf("--volume=%s:/tmp/config.toml", baseFile),
		fmt.Sprintf("ethereum/client-go:%s", Version),
		"dumpconfig",
		"--config=/tmp/config.toml")
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err = cmd.Run()
	require.NoError(t, err, stderr.String())

	t.Logf("geth dumpconfig logs:\n%s", stderr.String())

	return stdout.Bytes()
}
