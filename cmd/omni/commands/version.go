package commands

import (
	"github.com/omni-network/omni"
	"github.com/spf13/cobra"
)

func (c *command) initVersionCmd() {
	v := &cobra.Command{
		Use:   "version",
		Short: "Print version number",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Println(omni.Version)
		},
	}
	v.SetOut(c.root.OutOrStdout())
	c.root.AddCommand(v)
}
