package cmd

import (
	monitor "github.com/omni-network/omni/monitor/app"

	"github.com/spf13/pflag"
)

func bindRunFlags(flags *pflag.FlagSet, cfg *monitor.Config) {
	flags.StringVar(&cfg.NetworkFile, "network-file", cfg.NetworkFile, "The path to the network file e.g path/network.json")
	flags.StringVar(&cfg.MonitoringAddr, "monitoring-addr", cfg.MonitoringAddr, "The address to bind the monitoring server")
}
