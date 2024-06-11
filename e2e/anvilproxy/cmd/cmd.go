// Package cmd provides the cli for running the api.
package cmd

import (
	"github.com/omni-network/omni/e2e/anvilproxy/app"
	libcmd "github.com/omni-network/omni/lib/cmd"
	"github.com/omni-network/omni/lib/log"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// New returns a new root cobra command that runs the anvilproxy server.
func New() *cobra.Command {
	cmd := libcmd.NewRootCmd(
		"anvilproxy",
		"Anvil proxy server supporting fuzzy head",
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
	flags.StringVar(&cfg.ListenAddr, "listen-addr", cfg.ListenAddr, "Address for proxy to listen on")
	flags.Uint64Var(&cfg.ChainID, "chain-id", cfg.ChainID, "Anvil chain id")
	flags.StringVar(&cfg.LoadState, "load-state", cfg.LoadState, "Initialize the chain from a previously saved state snapshot")
	flags.Uint64Var(&cfg.BlockTimeSecs, "block-time", cfg.BlockTimeSecs, "Block time in seconds for interval mining")
	flags.BoolVar(&cfg.Silent, "silent", cfg.Silent, "Don't print anything on startup and don't print logs")
	flags.Uint64Var(&cfg.SlotsInEpoch, "slots-in-an-epoch", cfg.SlotsInEpoch, "Slots in an epoch")
}
