package cmd

import (
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/scripts/trade/config"

	"github.com/spf13/pflag"
)

func bindRunFlags(flags *pflag.FlagSet, cfg *config.Config) {
	netconf.BindFlag(flags, &cfg.Network)
	flags.StringVar(&cfg.RPCListen, "rpc-listen", cfg.RPCListen, "The address to listen on for RPC requests")
	flags.StringVar(&cfg.DBConn, "db-conn", cfg.DBConn, "Postgres connection string for the database")
}
