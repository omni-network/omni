package cmd

import (
	"github.com/omni-network/omni/lib/feature"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"
	solver "github.com/omni-network/omni/solver/app"

	"github.com/spf13/pflag"
)

func bindRunFlags(flags *pflag.FlagSet, cfg *solver.Config) {
	netconf.BindFlag(flags, &cfg.Network)
	xchain.BindFlags(flags, &cfg.RPCEndpoints)
	feature.BindFlag(flags, &cfg.FeatureFlags)
	flags.StringVar(&cfg.SolverPrivKey, "private-key", cfg.SolverPrivKey, "The path to the solver private key e.g path/private.key")
	flags.StringVar(&cfg.LoadGenPrivKey, "loadgen-key", cfg.LoadGenPrivKey, "The path to the loadgen private key e.g path/loadgen.key (not applicable to protected networks)")
	flags.StringVar(&cfg.MonitoringAddr, "monitoring-addr", cfg.MonitoringAddr, "The address to bind the monitoring server")
	flags.StringVar(&cfg.DBDir, "db-dir", cfg.DBDir, "The path to the database directory")
}
