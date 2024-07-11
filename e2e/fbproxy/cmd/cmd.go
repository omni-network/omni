package cmd

import (
	"github.com/omni-network/omni/e2e/fbproxy/app"
	libcmd "github.com/omni-network/omni/lib/cmd"
	"github.com/omni-network/omni/lib/log"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// New returns a new root cobra command that runs the anvilproxy server.
func New() *cobra.Command {
	cmd := libcmd.NewRootCmd(
		"fbproxy",
		"Fireblocks ETH JSON-RPC proxy server that supports raw signing and transaction broadcasting",
	)

	cfg := app.DefaultConfig()
	bindFlags(cmd.Flags(), &cfg)

	logCfg := log.DefaultConfig()
	log.BindFlags(cmd.Flags(), &logCfg)

	cmd.RunE = func(cmd *cobra.Command, _ []string) error {
		ctx, err := log.Init(cmd.Context(), logCfg)
		if err != nil {
			return err
		}

		if err := libcmd.LogFlags(ctx, cmd.Flags()); err != nil {
			return err
		}

		return app.Run(ctx, cfg)
	}

	return cmd
}

func bindFlags(flags *pflag.FlagSet, cfg *app.Config) {
	flags.StringVar((*string)(&cfg.Network), "network", string(cfg.Network), "Network ID")
	flags.StringVar(&cfg.ListenAddr, "listen-addr", cfg.ListenAddr, "Address for proxy to listen on")
	flags.StringVar(&cfg.BaseRPC, "base-rpc", cfg.BaseRPC, "Base RPC URL to forward requests to; e.g. http://localhost:8545")
	flags.StringVar(&cfg.FireAPIKey, "fireblocks-api-key", cfg.FireAPIKey, "FireBlocks api key")
	flags.StringVar(&cfg.FireKeyPath, "fireblocks-key-path", cfg.FireKeyPath, "FireBlocks RSA private key path")
}
