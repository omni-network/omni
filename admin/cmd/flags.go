package cmd

import (
	"github.com/omni-network/omni/admin/app"

	"github.com/spf13/pflag"
)

func bindCommmonFlags(flags *pflag.FlagSet, cfg *app.Config) {
	flags.StringVar(&cfg.Chain, "chain", cfg.Chain, "Run admin command on a specific chain (\"--chain=all\" for all chains)")
	flags.StringVar(&cfg.FireAPIKey, "fireblocks-api-key", cfg.FireAPIKey, "FireBlocks api key")
	flags.StringVar(&cfg.FireKeyPath, "fireblocks-key-path", cfg.FireKeyPath, "FireBlocks RSA private key path")
	flags.StringToStringVar(&cfg.RPCs, "rpcs", cfg.RPCs, "Public chain rpc endpoints: '<chain1>=<url1>,...")
	flags.StringVar((*string)(&cfg.Network), "network", string(cfg.Network), "Network to run admin command on (staging, omega, mainnet)")
}
