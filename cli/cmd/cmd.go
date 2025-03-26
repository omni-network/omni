package cmd

import (
	"github.com/omni-network/omni/lib/buildinfo"
	libcmd "github.com/omni-network/omni/lib/cmd"

	"github.com/spf13/cobra"
)

// New returns a new root cobra command that handles our command line tool.
func New() *cobra.Command {
	return libcmd.NewRootCmd(
		"omni",
		"CLI providing tools for interacting with omni",
		newOperatorCmds(),
		newDeveloperCmds(),
		newDevnetCmds(),
		buildinfo.NewVersionCmd(),
	)
}

func newOperatorCmds() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "operator",
		Short: "Operator commands",
		Args:  cobra.NoArgs,
	}

	cmd.AddCommand(
		newInitCmd(),
		newCreateValCmd(),
		newCreateOperatorKeyCmd(),
		newCreateConsensusKeyCmd(),
		newUnjailCmd(),
		newDelegateCmd(),
		newEditValCmd(),
	)

	return cmd
}
