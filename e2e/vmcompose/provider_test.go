package vmcompose_test

import (
	"context"
	"os"
	"path/filepath"
	"regexp"
	"testing"

	"github.com/omni-network/omni/e2e/app"
	"github.com/omni-network/omni/e2e/vmcompose"
	"github.com/omni-network/omni/lib/tutil"

	"github.com/stretchr/testify/require"
)

//go:generate go test . -golden -clean

func TestSetup(t *testing.T) {
	t.Parallel()
	manifestFile, dataFile := vmcompose.SetupDataFixtures(t)

	def, err := app.MakeDefinition(context.Background(), app.DefinitionConfig{
		ManifestFile:  manifestFile,
		InfraProvider: vmcompose.ProviderName,
		InfraDataFile: dataFile,
		OmniImgTag:    "7d1ae53",
	}, "")
	require.NoError(t, err)

	err = os.MkdirAll(def.Testnet.Dir, 0o755)
	require.NoError(t, err)

	err = def.Infra.Setup()
	require.NoError(t, err)

	files, err := filepath.Glob(filepath.Join(def.Testnet.Dir, "*compose.yaml"))
	require.NoError(t, err)

	for _, file := range files {
		t.Run(filepath.Base(file), func(t *testing.T) {
			t.Parallel()
			bz, err := os.ReadFile(file)
			require.NoError(t, err)

			// Replace non-deterministic fields with placeholders

			re1 := regexp.MustCompile(`--nodekeyhex=([0-9a-fA-F]+)`)
			bz = re1.ReplaceAll(bz, []byte("--nodekeyhex=<nodekeyhex>"))

			re2 := regexp.MustCompile(`enode://([0-9a-fA-F]+)`)
			bz = re2.ReplaceAll(bz, []byte("enode://<enode_pubkey>"))

			tutil.RequireGoldenBytes(t, bz)
		})
	}
}
