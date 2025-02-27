package cmd

import (
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"
	monitor "github.com/omni-network/omni/monitor/app"
	"github.com/omni-network/omni/monitor/loadgen"
	"github.com/omni-network/omni/monitor/xfeemngr"

	"github.com/spf13/pflag"
)

func bindRunFlags(flags *pflag.FlagSet, cfg *monitor.Config) {
	netconf.BindFlag(flags, &cfg.Network)
	xchain.BindFlags(flags, &cfg.RPCEndpoints)
	flags.StringVar(&cfg.PrivateKey, "private-key", cfg.PrivateKey, "The path to the private key e.g path/private.key")
	flags.StringVar(&cfg.FlowGenKey, "flowgen-key", cfg.FlowGenKey, "The path to the flowgen private key e.g path/flowgen.key")
	flags.StringVar(&cfg.MonitoringAddr, "monitoring-addr", cfg.MonitoringAddr, "The address to bind the monitoring server")
	flags.StringVar(&cfg.HaloCometURL, "halo-url", cfg.HaloCometURL, "The URL of the halo node e.g localhost:26657")
	flags.StringVar(&cfg.HaloGRPCURL, "halo-grpc-url", cfg.HaloGRPCURL, "The gRPC URL of the halo node e.g localhost:9999")
	flags.StringVar(&cfg.DBDir, "db-dir", cfg.DBDir, "The path to the database directory")
	flags.StringVar(&cfg.RouteScanAPIKey, "routescan-apikey", cfg.RouteScanAPIKey, "The RouteScan API key to use their APIs with higher rate limits")
}

func bindLoadGenFlags(flags *pflag.FlagSet, cfg *loadgen.Config) {
	flags.StringVar(&cfg.ValidatorKeysGlob, "loadgen-validator-keys-glob", cfg.ValidatorKeysGlob, "Glob path to the validator keys used for self-delegation load generation. Only applicable to devnet and staging")
	flags.StringVar(&cfg.XCallerKey, "loadgen-xcaller-key", cfg.XCallerKey, "Path to the xcaller key used for xcall loadgen")
}

func bindXFeeMngrFlags(flags *pflag.FlagSet, cfg *xfeemngr.Config) {
	flags.StringToStringVar((*map[string]string)(&cfg.RPCEndpoints), "xfeemngr-rpc-endpoints", cfg.RPCEndpoints, "Cross-chain EVM RPC endpoints. e.g. \"ethereum=http://geth:8545,optimism=https://optimism.io\"")
	flags.StringVar(&cfg.CoinGeckoAPIKey, "xfeemngr-coingecko-apikey", cfg.CoinGeckoAPIKey, "The CoinGecko API key to use for fetching token prices")
}
