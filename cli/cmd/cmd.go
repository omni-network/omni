package cmd

import (
	libcmd "github.com/omni-network/omni/lib/cmd"

	"github.com/spf13/cobra"
)

// New returns a new root cobra command that handles our command line tool.
func New() *cobra.Command {
	return libcmd.NewRootCmd(
		"omni",
		"CLI providing tools for omni operators",
		newOperatorCmds(),
		newDevnetCmds(),
	)
}

func newOperatorCmds() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "operator",
		Short: "Operator commands",
		Args:  cobra.NoArgs,
	}

	cmd.AddCommand(newRegisterCmd())

	return cmd
}

func newRegisterCmd() *cobra.Command {
	var omniAVSAddress string

	cmd := &cobra.Command{
		Use:   "register",
		Short: "Register an operator with the Omni AVS contract",
		Long: `Register command expects a Eigen-Layer yaml config file as an argument
to successfully register an operator address with the Omni AVS contract.

Note the operator must already be registered with Eigen-Layer.`,
		Example: "  omni operator register <eigen-configuration-file>",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return Register(cmd.Context(), args[0], omniAVSAddress)
		},
	}

	cmd.Flags().StringVar(&omniAVSAddress, "omni-avs-address", omniAVSAddress, "Optional address of the Omni AVS contract.")

	return cmd
}
