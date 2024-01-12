// Package cmd provides the cli for running the explorerapi.
package cmd

import (
	"github.com/omni-network/omni/explorerapi/app"
	libcmd "github.com/omni-network/omni/lib/cmd"
	"github.com/omni-network/omni/lib/log"

	"github.com/spf13/cobra"
)

// New returns a new root cobra command that handles our command line tool.
func New() *cobra.Command {
	return libcmd.NewRootCmd(
		"collector",
		"Collector is a service that will initialize our streams to listen to our portals and index "+
			"data and put it in our Omni Blocks DB",
		newRunCmd(),
	)
}

// newRunCmd returns a new cobra command that runs the explorerapi.
func newRunCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "run",
		Short: "Runs the collector",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			log.Info(ctx, "Collector started")
			conf := app.Config{}

			err := app.Run(ctx, conf)
			if err != nil {
				log.Error(ctx, "Failed to start Collector", err)
				<-ctx.Done()

				return err
			}

			log.Info(ctx, "Press Ctrl+C to stop")
			<-ctx.Done()
			log.Info(ctx, "Collector stopped")

			return nil
		},
	}
}
