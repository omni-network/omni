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
	"fmt"
	"strconv"

	cmtos "github.com/cometbft/cometbft/libs/os"
	"github.com/cometbft/cometbft/p2p"
	"github.com/cometbft/cometbft/privval"
	"github.com/cometbft/cometbft/types"
	cmttime "github.com/cometbft/cometbft/types/time"
	config2 "github.com/omni-network/omni/pkg/config"
	"github.com/spf13/cobra"
)

func InitKeyFileConfigCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init",
		Short: "Initialize CometBFT",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return config.viperConfig.BindPFlags(cmd.Flags())
		},
		RunE: initFilesWithConfig,
	}
	setInitFlags(cmd)
	return cmd
}

func setInitFlags(cmd *cobra.Command) {
	testnetChainId := strconv.Itoa(int(config2.Testnet.NetworkID))
	cmd.Flags().String(optionNameChainId, testnetChainId, "use this chain id in genesis file")
	cmd.Flags().String(optionNamePortalAddress, config2.Testnet.OmniPortalAddress, "use this portal address in the config")
}

func initFilesWithConfig(cmd *cobra.Command, args []string) error {
	// accept only one argument, verbosity
	if len(args) > 2 {
		return cmd.Help()
	}
	chainIdString := config.viperConfig.GetString(optionNameChainId)
	portalAddress := config.viperConfig.GetString(optionNamePortalAddress)

	ccfg := config.cometBFTConfig
	ocfg := config.omniConfig
	oLogger := config.ologger

	// parse and load the chain id and portal address
	if chainIdString == "mainnet" {
		ocfg.ChainConfig = config2.Mainnet
	} else if chainIdString == "testnet" {
		ocfg.ChainConfig = config2.Testnet
	} else {
		chainId, err := strconv.Atoi(chainIdString)
		if err != nil {
			return err
		}
		ocfg.ChainConfig.NetworkID = uint64(chainId)
		ocfg.ChainConfig.OmniPortalAddress = portalAddress
	}

	// private validator
	privValKeyFile := ccfg.PrivValidatorKeyFile()
	privValStateFile := ccfg.PrivValidatorStateFile()
	var pv *privval.FilePV
	if cmtos.FileExists(privValKeyFile) {
		pv = privval.LoadFilePV(privValKeyFile, privValStateFile)
		oLogger.Info("Found private validator", "keyFile", privValKeyFile,
			"stateFile", privValStateFile)
	} else {
		pv = privval.GenFilePV(privValKeyFile, privValStateFile)
		pv.Save()
		oLogger.Info("Generated private validator", "keyFile", privValKeyFile,
			"stateFile", privValStateFile)
	}

	nodeKeyFile := ccfg.NodeKeyFile()
	if cmtos.FileExists(nodeKeyFile) {
		oLogger.Info("Found node key", "path", nodeKeyFile)
	} else {
		if _, err := p2p.LoadOrGenNodeKey(nodeKeyFile); err != nil {
			return err
		}
		oLogger.Info("Generated node key", "path", nodeKeyFile)
	}

	// genesis file
	genFile := ccfg.GenesisFile()
	if cmtos.FileExists(genFile) {
		oLogger.Info("Found genesis file", "path", genFile)
	} else {
		genDoc := types.GenesisDoc{
			ChainID:         strconv.Itoa(int(ocfg.ChainConfig.NetworkID)),
			GenesisTime:     cmttime.Now(),
			ConsensusParams: types.DefaultConsensusParams(),
		}
		pubKey, err := pv.GetPubKey()
		if err != nil {
			return fmt.Errorf("can't get pubkey: %w", err)
		}
		genDoc.Validators = []types.GenesisValidator{{
			Address: pubKey.Address(),
			PubKey:  pubKey,
			Power:   10,
		}}

		if err := genDoc.SaveAs(genFile); err != nil {
			return err
		}
		oLogger.Info("Generated genesis file", "path", genFile)
	}

	return nil
}
