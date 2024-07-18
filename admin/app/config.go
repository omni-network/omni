package app

import (
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/netconf"
)

const chainAll = "all"

type Config struct {
	Chain       string // Name of chain to run admin command on, use "all" to run on all chains
	FireAPIKey  string
	FireKeyPath string
	RPCs        map[string]string // map chain name to rpc url
	Network     netconf.ID
}

func DefaultConfig() Config {
	return Config{
		Chain: "",
	}
}

func (cfg Config) Validate() error {
	if cfg.Chain == "" {
		return errors.New("chain must be set")
	}

	if cfg.Network == "" {
		return errors.New("network must be set")
	}

	if err := cfg.Network.Verify(); err != nil {
		return errors.New("invalid network", "network", cfg.Network)
	}

	if cfg.Network != netconf.Devnet && (cfg.FireAPIKey == "" || cfg.FireKeyPath == "") {
		return errors.New("fireblocks api key and key path required for non-devnet networks")
	}

	return nil
}
