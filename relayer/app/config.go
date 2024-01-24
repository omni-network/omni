package relayer

import "path/filepath"

const (
	networkFile = "network.json"
	configDir   = "relayer/config"
)

type Config struct {
	PrivateKeyPath string `toml:"private_key_path"`
	HaloURL        string `toml:"halo_url"`
}

func (Config) NetworkFile() string {
	return filepath.Join(configDir, networkFile)
}

func DefaultRelayerConfig() Config {
	return Config{
		PrivateKeyPath: "",
		HaloURL:        "http://localhost:26657",
	}
}
