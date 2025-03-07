package cmd

import (
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"
	relayer "github.com/omni-network/omni/relayer/app"

	"github.com/spf13/pflag"
)

func bindRunFlags(flags *pflag.FlagSet, cfg *relayer.Config) {
	netconf.BindFlag(flags, &cfg.Network)
	xchain.BindFlags(flags, &cfg.RPCEndpoints)
	flags.StringVar(&cfg.PrivateKey, "private-key", cfg.PrivateKey, "The path to the private key e.g path/private.key")
	flags.StringVar(&cfg.HaloCometURL, "halo-url", cfg.HaloCometURL, "The URL of the halo node e.g localhost:26657")
	flags.StringVar(&cfg.HaloGRPCURL, "halo-grpc-url", cfg.HaloGRPCURL, "The gRPC URL of the halo node e.g localhost:9999")
	flags.StringVar(&cfg.MonitoringAddr, "monitoring-addr", cfg.MonitoringAddr, "The address to bind the monitoring server")
	flags.StringVar(&cfg.DBDir, "db-dir", cfg.DBDir, "The path to the database directory")
	flags.StringVar(&cfg.CoinGeckoAPIKey, "coingecko-apikey", cfg.CoinGeckoAPIKey, "The CoinGecko API key to use for fetching token prices")
}
