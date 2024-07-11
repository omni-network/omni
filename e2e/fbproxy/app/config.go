package app

import "github.com/omni-network/omni/lib/netconf"

type Config struct {
	ListenAddr  string
	Network     netconf.ID
	BaseRPC     string
	FireAPIKey  string
	FireKeyPath string
}

func DefaultConfig() Config {
	return Config{
		Network:     netconf.Devnet,
		ListenAddr:  "0.0.0.0:8545",
		BaseRPC:     "",
		FireAPIKey:  "",
		FireKeyPath: "",
	}
}
