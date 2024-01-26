// Package cmd provides the cli for running the api.
package cmd

import (
	"github.com/omni-network/omni/explorer/graphql/app"
	libcmd "github.com/omni-network/omni/lib/cmd"
	"github.com/omni-network/omni/lib/log"

	"github.com/spf13/cobra"
)

// New returns a new root cobra command that handles our command line tool.
func New() *cobra.Command {
	return libcmd.NewRootCmd(
		"api",
		"Explorer GraphQL Server",
		newRunCmd(),
		newDebugCmd(),
	)
}

// newRunCmd returns a new cobra command that runs the api.
func newRunCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "run",
		Short: "Runs the GraphQL server",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			log.Info(ctx, "Explorer GraphQL started")
			conf := app.DefaultExplorerAPIConfig()

			err := app.Run(ctx, conf)
			if err != nil {
				log.Error(ctx, "Failed to start Explorer GraphQL", err)
				<-ctx.Done()

				return err
			}

			log.Info(ctx, "Press Ctrl+C to stop")
			<-ctx.Done()
			log.Info(ctx, "ExplorerApi stopped")

			return nil
		},
	}
}

// newRunCmd returns a new cobra command that runs the api.
func newDebugCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "debug",
		Short: "Runs the GraphQL server without a PostgresClient",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			log.Info(ctx, "ExplorerAPI started")
			conf := app.DefaultExplorerAPIConfig()

			err := app.Debug(ctx, conf)
			if err != nil {
				log.Error(ctx, "Failed to start Explorer API", err)
				<-ctx.Done()

				return err
			}

			log.Info(ctx, "Press Ctrl+C to stop")
			<-ctx.Done()
			log.Info(ctx, "ExplorerApi stopped")

			return nil
		},
	}
}
