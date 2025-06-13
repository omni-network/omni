package config_test

import (
	"os"
	"path/filepath"
	"testing"

	halocfg "github.com/omni-network/omni/halo/config"
	"github.com/omni-network/omni/lib/feature"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/tutil"

	"github.com/stretchr/testify/require"
)

//go:generate go test . -golden -clean

func TestDefaultConfigReference(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		cfgFunc func() halocfg.Config
	}{
		{
			name:    "default",
			cfgFunc: halocfg.DefaultConfig,
		},
		{
			name: "test",
			cfgFunc: func() halocfg.Config {
				cfg := halocfg.DefaultConfig()
				cfg.Network = netconf.Omega
				cfg.EngineEndpoint = "http://omni_evm:8551"
				cfg.EngineJWTFile = "/geth/jwtsecret"
				cfg.RPCEndpoints = map[string]string{
					"mock": "http://mock_rpc:8545",
				}
				cfg.UnsafeSkipUpgrades = []int{1, 2, 3}
				cfg.FeatureFlags = feature.Flags{"a", "b"}
				cfg.EVMProxyListen = "0.0.0.0:8545"
				cfg.EVMProxyTarget = "http://omni_evm:8545"

				return cfg
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			cfg := tt.cfgFunc()
			tempDir := t.TempDir()
			cfg.HomeDir = tempDir

			require.NoError(t, os.MkdirAll(filepath.Join(tempDir, "config"), 0o755))
			require.NoError(t, halocfg.WriteConfigTOML(cfg, log.DefaultConfig()))

			b, err := os.ReadFile(filepath.Join(tempDir, "config", "halo.toml"))
			require.NoError(t, err)

			tutil.RequireGoldenBytes(t, b, tutil.WithFilename(tt.name+"_halo.toml"))
		})
	}
}
