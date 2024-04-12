package cmd

import (
	monitor "github.com/omni-network/omni/monitor/app"
	"github.com/omni-network/omni/monitor/loadgen"

	"github.com/spf13/pflag"
)

func bindRunFlags(flags *pflag.FlagSet, cfg *monitor.Config) {
	flags.StringVar(&cfg.PrivateKey, "private-key", cfg.PrivateKey, "The path to the private key e.g path/private.key")
	flags.StringVar(&cfg.NetworkFile, "network-file", cfg.NetworkFile, "The path to the network file e.g path/network.json")
	flags.StringVar(&cfg.MonitoringAddr, "monitoring-addr", cfg.MonitoringAddr, "The address to bind the monitoring server")
}

func bindLoadGenFlags(flags *pflag.FlagSet, cfg *loadgen.Config) {
	flags.StringVar(&cfg.ValidatorKeysGlob, "loadgen-validator-keys-glob", cfg.ValidatorKeysGlob, "Glob path to the validator keys used for self-delegation load generation. Only applicable to devnet and staging")
}
