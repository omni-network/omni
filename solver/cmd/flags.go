package cmd

import (
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/tracer"
	"github.com/omni-network/omni/lib/xchain"
	solver "github.com/omni-network/omni/solver/app"

	"github.com/spf13/pflag"
)

func bindRunFlags(flags *pflag.FlagSet, cfg *solver.Config) {
	netconf.BindFlag(flags, &cfg.Network)
	xchain.BindFlags(flags, &cfg.RPCEndpoints)
	tracer.BindFlags(flags, &cfg.Tracer)
	flags.StringVar(&cfg.SolverPrivKey, "private-key", cfg.SolverPrivKey, "The path to the solver private key e.g path/private.key")
	flags.StringVar(&cfg.MonitoringAddr, "monitoring-addr", cfg.MonitoringAddr, "The address to bind the monitoring server")
	flags.StringVar(&cfg.APIAddr, "api-addr", cfg.APIAddr, "The address to bind the API server")
	flags.StringVar(&cfg.DBDir, "db-dir", cfg.DBDir, "The path to the database directory")
	flags.StringVar(&cfg.CoinGeckoAPIKey, "coingecko-apikey", cfg.CoinGeckoAPIKey, "The CoinGecko API key to use for fetching token prices")
	flags.StringToStringVar(&cfg.RPCOverrides, "rpc-overrides", cfg.RPCOverrides, "Public chain rpc overrides: '<chain1>=<url1>,<url2>'")
}
