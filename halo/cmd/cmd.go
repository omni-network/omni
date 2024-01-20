// Package cmd provides the cli for running the halo consensus client.
package cmd

import (
	"context"

	"github.com/omni-network/omni/halo/app"
	libcmd "github.com/omni-network/omni/lib/cmd"

	"github.com/spf13/cobra"
)

// New returns a new root cobra command that handles our command line tool.
func New() *cobra.Command {
	return libcmd.NewRootCmd(
		"halo",
		"Halo is a consensus client implementation for the Omni Protocol",
		newRunCmd(app.Run),
		newInitCmd(),
	)
}

// newRunCmd returns a new cobra command that runs the halo consensus client.
func newRunCmd(runFunc func(context.Context, app.Config) error) *cobra.Command {
	cfg := app.Config{
		HaloConfig: app.DefaultHaloConfig(),
	}

	cmd := &cobra.Command{
		Use:   "run",
		Short: "Runs the halo consensus client",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()

			if err := logConfig(ctx, cmd.Flags()); err != nil {
				return err
			}

			var err error
			cfg.Comet, err = parseCometConfig(ctx, cfg.HomeDir)
			if err != nil {
				return err
			}

			return runFunc(ctx, cfg)
		},
	}

	bindRunFlags(cmd.Flags(), &cfg.HaloConfig)

	return cmd
}
