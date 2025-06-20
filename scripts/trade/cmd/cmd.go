// Package cmd provides the cli for running the monitor service
package cmd

import (
	libcmd "github.com/omni-network/omni/lib/cmd"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/scripts/trade/app"
	"github.com/omni-network/omni/scripts/trade/config"

	"github.com/spf13/cobra"
)

// New returns a new root cobra command that handles our command line tool.
func New() *cobra.Command {
	cmd := libcmd.NewRootCmd(
		"trade",
		"Trading terminal backend.",
	)

	cfg := config.DefaultConfig()
	bindRunFlags(cmd.Flags(), &cfg)

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
