package relayer

type Config struct {
	PrivateKey  string
	HaloURL     string
	NetworkFile string
}

func DefaultRelayerConfig() Config {
	return Config{
		PrivateKey:  "relayer.key",
		HaloURL:     "localhost:26657",
		NetworkFile: "network.json",
	}
}
