// Copyright 2023 The Omni Authors. All rights reserved.
// This file is part of the omni library.
//
// The omni library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The omni library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.

// You should have received a copy of the GNU Lesser General Public License
// along with the omni library. If not, see <http://www.gnu.org/licenses/>.

package commands

import (
	"encoding/json"
	"github.com/spf13/cobra"
)

// PrintConfigCmd prints the configuration of omni and cometBFT to
// console in json format.
func PrintConfigCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "printconfig",
		Short: "resolve the configuration from env., yaml file or from command line and print",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return config.viperConfig.BindPFlags(cmd.Flags())
		},
		RunE: printConfiguration,
	}
	setPrintConfigFlags(cmd)
	config.cmd.AddCommand(cmd)
	return cmd
}

func setPrintConfigFlags(cmd *cobra.Command) {
	cmd.Flags().Bool(optionNameVerbose, false, "display in verbose format")
}

func printConfiguration(cmd *cobra.Command, args []string) error {
	if len(args) > 1 {
		return cmd.Help()
	}
	verbose := config.viperConfig.GetBool(optionNameVerbose)

	ologger.Info("---- Omni Configuration ------- ")
	str, _ := json.MarshalIndent(config.omniConfig, "", "  ")
	ologger.Info(string(str))
	if verbose {
		ologger.Info("---- cometBFT Configuration ------- ")
		str, _ = json.MarshalIndent(config.cometBFTConfig, "", "  ")
		ologger.Info(string(str))
	}
	return nil
}
