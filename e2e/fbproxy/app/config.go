package app

type Config struct {
	ListenAddr  string
	Network     string
	BaseRPC     string
	FireAPIKey  string
	FireKeyPath string
}

func DefaultConfig() Config {
	return Config{
		Network:     "devnet",
		ListenAddr:  "0.0.0.0:8545",
		BaseRPC:     "http://localhost:8545",
		FireAPIKey:  "",
		FireKeyPath: "",
	}
}
