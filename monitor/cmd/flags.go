package cmd

import (
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"
	monitor "github.com/omni-network/omni/monitor/app"
	"github.com/omni-network/omni/monitor/loadgen"

	"github.com/spf13/pflag"
)

func bindRunFlags(flags *pflag.FlagSet, cfg *monitor.Config) {
	netconf.BindFlag(flags, &cfg.Network)
	xchain.BindFlags(flags, &cfg.RPCEndpoints)
	flags.StringVar(&cfg.PrivateKey, "private-key", cfg.PrivateKey, "The path to the private key e.g path/private.key")
	flags.StringVar(&cfg.MonitoringAddr, "monitoring-addr", cfg.MonitoringAddr, "The address to bind the monitoring server")
	flags.StringVar(&cfg.HaloURL, "halo-url", cfg.HaloURL, "The URL of the halo node e.g localhost:26657")
	flags.StringVar(&cfg.DBDir, "db-dir", cfg.DBDir, "The path to the database directory")
}

func bindLoadGenFlags(flags *pflag.FlagSet, cfg *loadgen.Config) {
	flags.StringVar(&cfg.ValidatorKeysGlob, "loadgen-validator-keys-glob", cfg.ValidatorKeysGlob, "Glob path to the validator keys used for self-delegation load generation. Only applicable to devnet and staging")
}
