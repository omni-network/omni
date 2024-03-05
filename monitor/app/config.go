package monitor

type Config struct {
	NetworkFile    string
	MonitoringAddr string
	AVSAddress     string
}

func DefaultConfig() Config {
	return Config{
		NetworkFile:    "network.json",
		MonitoringAddr: ":26660",
		AVSAddress:     "0x0000000000000000000000000000000000000000",
	}
}
