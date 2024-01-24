package relayer

import "path/filepath"

const (
	networkFile = "network.json"
	configDir   = "relayer/config"
)

type Config struct {
	PrivateKey string `toml:"private_key"`
	HaloURL    string `toml:"halo_url"`
}

func (Config) NetworkFile() string {
	return filepath.Join(configDir, networkFile)
}

func DefaultRelayerConfig() Config {
	return Config{
		PrivateKey: "relayer.key",
		HaloURL:    "http://localhost:26657",
	}
}
