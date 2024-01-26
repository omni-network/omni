// Package cmd provides the cli for running the indexer.
package cmd

import (
	"github.com/omni-network/omni/explorer/indexer/app"
	libcmd "github.com/omni-network/omni/lib/cmd"
	"github.com/omni-network/omni/lib/log"

	"github.com/spf13/cobra"
)

// New returns a new root cobra command that handles our command line tool.
func New() *cobra.Command {
	return libcmd.NewRootCmd(
		"indexer",
		"Indexer is a service that will initialize our streams to listen to our portals and index "+
			"data and put it in our Omni Blocks DB",
		newRunCmd(),
	)
}

// newRunCmd returns a new cobra command that runs the Indexer.
func newRunCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "run",
		Short: "Runs the indexer",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			log.Info(ctx, "Indexer started")
			conf := app.DefaultExplorerAPIConfig()

			err := app.Run(ctx, conf)
			if err != nil {
				log.Error(ctx, "Failed to start Indexer", err)
				<-ctx.Done()

				return err
			}

			log.Info(ctx, "Press Ctrl+C to stop")
			<-ctx.Done()
			log.Info(ctx, "Indexer stopped")

			return nil
		},
	}
}
