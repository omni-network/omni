package cmd

import (
	libcmd "github.com/omni-network/omni/lib/cmd"

	"github.com/spf13/cobra"
)

func DeRegisterOperatorFromOmniAVS() *cobra.Command {
	deregisterFromAVSCmd := &cobra.Command{
		Use:   "deregister",
		Short: "deregister validator from omni avs",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			return libcmd.LogFlags(ctx, cmd.Flags())
		},
	}

	return deregisterFromAVSCmd
}
