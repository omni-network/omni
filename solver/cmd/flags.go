package cmd

import (
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"
	solver "github.com/omni-network/omni/solver/app"

	"github.com/spf13/pflag"
)

func bindRunFlags(flags *pflag.FlagSet, cfg *solver.Config) {
	netconf.BindFlag(flags, &cfg.Network)
	xchain.BindFlags(flags, &cfg.RPCEndpoints)
	flags.StringVar(&cfg.PrivateKey, "private-key", cfg.PrivateKey, "The path to the private key e.g path/private.key")
	flags.StringVar(&cfg.MonitoringAddr, "monitoring-addr", cfg.MonitoringAddr, "The address to bind the monitoring server")
	flags.StringVar(&cfg.DBDir, "db-dir", cfg.DBDir, "The path to the database directory")
}
