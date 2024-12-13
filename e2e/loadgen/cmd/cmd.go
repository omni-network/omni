package cmd

import (
	"github.com/omni-network/omni/e2e/loadgen/app"
	libcmd "github.com/omni-network/omni/lib/cmd"
	"github.com/omni-network/omni/lib/log"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// New returns a new root cobra command that runs the xcall load generator.
func New() *cobra.Command {
	cmd := libcmd.NewRootCmd(
		"loadgen",
		"Generate xcall load on a specified network.",
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
	flags.IntVar(&cfg.Count, "count", cfg.Count, "number of xcalls to make per chain id pair")
	flags.StringVar(&cfg.Network, "network", cfg.Network, "network to generate load on")
	flags.StringVar(&cfg.Role, "role", cfg.Role, "role to use to pay for xcalls")
	flags.StringVar(&cfg.ChainIDPairs, "chainid-pairs", cfg.ChainIDPairs, "comma separated list of pairs of chain ids to make xcalls from:to e.g: ethereum:arbitrum_one,arbitrum_one:optimism,base:arbitrum_one")
}
