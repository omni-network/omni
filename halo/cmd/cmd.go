// Package cmd provides the cli for running the halo consensus client.
package cmd

import (
	"context"

	"github.com/omni-network/omni/halo/app"
	halocfg "github.com/omni-network/omni/halo/config"
	"github.com/omni-network/omni/lib/buildinfo"
	libcmd "github.com/omni-network/omni/lib/cmd"
	"github.com/omni-network/omni/lib/log"

	"github.com/spf13/cobra"
)

// New returns a new root cobra command that handles our command line tool.
func New() *cobra.Command {
	return libcmd.NewRootCmd(
		"halo",
		"Halo is a consensus client implementation for the Omni Protocol",
		newRunCmd("run", app.Run),
		newInitCmd(),
		buildinfo.NewVersionCmd(),
	)
}

// newRunCmd returns a new cobra command that runs the halo consensus client.
func newRunCmd(name string, runFunc func(context.Context, app.Config) error) *cobra.Command {
	haloCfg := halocfg.DefaultConfig()
	logCfg := log.DefaultConfig()

	cmd := &cobra.Command{
		Use:   name,
		Short: "Runs the halo consensus client",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := log.Init(cmd.Context(), logCfg)
			if err != nil {
				return err
			}
			if err := libcmd.LogFlags(ctx, cmd.Flags()); err != nil {
				return err
			}

			cometCfg, err := parseCometConfig(ctx, haloCfg.HomeDir)
			if err != nil {
				return err
			}

			return runFunc(ctx, app.Config{
				Config: haloCfg,
				Comet:  cometCfg,
			})
		},
	}

	bindRunFlags(cmd.Flags(), &haloCfg)
	log.BindFlags(cmd.Flags(), &logCfg)

	return cmd
}
