// Package cmd provides the cli for running the halo consensus client.
package cmd

import (
	"context"

	halo1 "github.com/omni-network/omni/halo/app"
	halo2 "github.com/omni-network/omni/halo2/app"
	libcmd "github.com/omni-network/omni/lib/cmd"
	"github.com/omni-network/omni/lib/log"

	"github.com/spf13/cobra"
)

// New returns a new root cobra command that handles our command line tool.
func New() *cobra.Command {
	return libcmd.NewRootCmd(
		"halo",
		"Halo is a consensus client implementation for the Omni Protocol",
		newRunCmd("run1", halo1.Run),
		newInitCmd(),
		newRunCmd("run", halo2.Run),
	)
}

// newRunCmd returns a new cobra command that runs the halo consensus client.
func newRunCmd(name string, runFunc func(context.Context, halo1.Config) error) *cobra.Command {
	haloCfg := halo1.DefaultHaloConfig()
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

			return runFunc(ctx, halo1.Config{
				HaloConfig: haloCfg,
				Comet:      cometCfg,
			})
		},
	}

	bindRunFlags(cmd.Flags(), &haloCfg)
	log.BindFlags(cmd.Flags(), &logCfg)

	return cmd
}
