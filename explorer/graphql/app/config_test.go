package app_test

import (
	"os"
	"path/filepath"
	"testing"

	graphql "github.com/omni-network/omni/explorer/graphql/app"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/tutil"

	"github.com/stretchr/testify/require"
)

//go:generate go test . -golden -clean

func TestDefaultConfigReference(t *testing.T) {
	t.Parallel()
	tempDir := t.TempDir()

	tests := []struct {
		name string
		cfg  graphql.Config
	}{
		{
			"default",
			graphql.DefaultConfig(),
		},
		{
			"omega",
			func() graphql.Config {
				cfg := graphql.DefaultConfig()
				cfg.Network = netconf.Omega
				return cfg
			}(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path := filepath.Join(tempDir, "explorer_graphql.toml")

			require.NoError(t, os.MkdirAll(tempDir, 0o755))
			require.NoError(t, graphql.WriteConfigTOML(tt.cfg, log.DefaultConfig(), path))

			b, err := os.ReadFile(path)
			require.NoError(t, err)

			tutil.RequireGoldenBytes(t, b, tutil.WithFilename(tt.name+"_explorer_graphql.toml"))
		})
	}
}
