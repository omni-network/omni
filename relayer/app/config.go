package relayer

import "path/filepath"

const (
	networkFile = "network.json"
	configDir   = "config"
)

type Config struct {
	PrivateKeyPath string `toml:"private_key_path"`
}

func (Config) NetworkFile() string {
	return filepath.Join(configDir, networkFile)
}

func DefaultRelayerConfig() Config {
	return Config{
		PrivateKeyPath: "",
	}
}
