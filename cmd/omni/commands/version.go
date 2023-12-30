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
	"github.com/cometbft/cometbft/version"
	"github.com/omni-network/omni"
	"github.com/omni-network/omni/pkg/halo"
	"github.com/spf13/cobra"
	"strconv"
)

func PrintVersionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Print version number",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return config.viperConfig.BindPFlags(cmd.Flags())
		},
		RunE: printVersion,
	}
	setVersionFlags(cmd)
	return cmd
}

func setVersionFlags(cmd *cobra.Command) {
	cmd.Flags().Bool(optionNameVerbose, false, "display in verbose format")
}

func printVersion(cmd *cobra.Command, args []string) error {
	// accept only one argument, verbosity
	if len(args) > 1 {
		return cmd.Help()
	}
	verbose := config.viperConfig.GetBool(optionNameVerbose)

	logger := config.ologger
	logger.Info("---- Version Info ------- ")
	logger.Info("Omni-Node: " + omni.Version)
	logger.Info("Omni-Halo: " + halo.HaloProtocolVersion)
	if verbose {
		logger.Info("CometBFT-Core: " + version.TMCoreSemVer)
		logger.Info("CometBFT-ABCI: " + version.ABCIVersion)
		logger.Info("CometBFT-P2P: " + strconv.Itoa(int(version.P2PProtocol)))
		logger.Info("CometBFT-BlockProtocol: " + strconv.Itoa(int(version.BlockProtocol)))
	}
	return nil
}
