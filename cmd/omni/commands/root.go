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
	"errors"
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
	"strings"

	ccfg "github.com/cometbft/cometbft/config"
	cflags "github.com/cometbft/cometbft/libs/cli/flags"
	"github.com/cometbft/cometbft/libs/log"
	ocfg "github.com/omni-network/omni/pkg/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// load default configurations for omni and cometBFT
	config = loadDefaultConfig()

	// set the default omni logger
	ologger = log.NewTMLogger(log.NewSyncWriter(os.Stdout))

	// set the default cometBFT logger
	clogger = log.NewTMLogger(log.NewSyncWriter(os.Stdout))
)

// command options
const (
	optionNameHomeDir     = "home_dir"
	optionNameOmniRootDir = "root_dir"
	optionNameLogLevel    = "log_level"
	optionNameLogWriter   = "log_writer"
	optionNameVerbose     = "verbose"
)

type Config struct {
	cmd            *cobra.Command   // the cobra command which has all the commandline
	viperConfig    *viper.Viper     // viper for reading config from files
	omniConfig     *ocfg.OmniConfig // omni related configuration
	cometBFTConfig *ccfg.Config     // cometBFT related configuration
}

// GetRootCommand  is the root command for omni node
func GetRootCommand() (*cobra.Command, error) {
	RootCmd := &cobra.Command{
		Use:           "omni",
		Short:         "Omni node hosting halo protocol for Xchain messages",
		SilenceErrors: true,
		SilenceUsage:  true,
	}

	// adds persistent commands to the root so that it is available for all children
	addPersistentCommands(RootCmd)

	// binds the cobra flags to the viper
	if err := viper.BindPFlags(RootCmd.Flags()); err != nil {
		return nil, err
	}

	// parse both the omni and cometbft configuration
	config, err := parseConfig(RootCmd)
	if err != nil {
		return nil, err
	}

	// initialise the omni logger
	ologger, err = getLogger(ologger, config.omniConfig.LogLevel, ocfg.DefaultLogLevel, "omni")
	if err != nil {
		return nil, err
	}

	// initialise the cometBFT logger
	clogger, err = getLogger(clogger, config.cometBFTConfig.LogLevel, ocfg.DefaultLogLevel, "comt")
	if err != nil {
		return nil, err
	}

	return RootCmd, nil
}

func addPersistentCommands(cmd *cobra.Command) {
	cmd.PersistentFlags().StringP(optionNameHomeDir, "", os.ExpandEnv("$HOME"), "home directory for creating omni root")
	cmd.PersistentFlags().String(optionNameLogLevel, config.omniConfig.LogLevel, "log level (info, debug, error, none)")
	cmd.PersistentFlags().String(optionNameLogWriter, config.omniConfig.LogWriter, "where to write the log output")
}

func loadDefaultConfig() *Config {
	// create default Omni config
	oConfig := ocfg.LoadDefaultConfig()

	// create default cometBFT config
	cConfig := ccfg.DefaultConfig()

	// overwrite few defaults in cometbft config
	cConfig.Moniker = oConfig.NodeName  // cometBFT moniker
	cConfig.LogLevel = oConfig.LogLevel // cometBFT loglevel

	return &Config{
		viperConfig:    viper.New(),
		omniConfig:     oConfig,
		cometBFTConfig: cConfig,
	}
}

func parseConfig(cmd *cobra.Command) (*Config, error) {
	// config loading order
	// - defaults values loaded first
	// - env vars are applied on top of that
	// - config file values are applied on top of that
	// - commandline variables overrides everything and applied last

	// set the command line arguments
	config.cmd = cmd

	// read omni config from environment variables
	config.viperConfig.SetEnvPrefix("OMNI")
	config.viperConfig.AutomaticEnv() // read in environment variables that match
	config.viperConfig.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))

	// parse the home dir if it is given in the command line
	hDir, err := cmd.Flags().GetString(optionNameHomeDir)
	if err == nil {
		config.omniConfig.HomeDirectory = hDir
	}

	// parse the omni root dir if it is given in the command line
	rDir, err := cmd.Flags().GetString(optionNameOmniRootDir)
	if err == nil {
		config.omniConfig.OmniRootDir = rDir
	}

	// make sure the directories exists, otherwise create them
	err = ocfg.EnsureDirectories(config.omniConfig)
	if err != nil {
		return nil, err
	}

	// set the file name to read
	config.viperConfig.AddConfigPath(config.omniConfig.OmniConfigDir)
	config.viperConfig.SetConfigName(config.omniConfig.OmniConfigFileName)

	// If a config file is found, read it in.
	if err := config.viperConfig.ReadInConfig(); err != nil {
		var e viper.ConfigFileNotFoundError
		if !errors.As(err, &e) {
			return nil, err
		} else {
			// write a default config file
			configFileName := filepath.Join(config.omniConfig.OmniConfigDir, config.omniConfig.OmniConfigFileName+".yaml")
			if err := writeConfigFile(configFileName, config.omniConfig); err != nil {
				return nil, err
			}
		}
	}

	// read config for cometBFT
	err = parseCometBFTConfig(config.cometBFTConfig, config.omniConfig.HomeDirectory)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func parseCometBFTConfig(cometBFTConfig *ccfg.Config, homeDir string) error {
	// set the cometBFT root directory with the same home as omni
	cometBFTConfig.RootDir = filepath.Join(homeDir, ccfg.DefaultTendermintDir)

	err := viper.Unmarshal(cometBFTConfig)
	if err != nil {
		return err
	}

	// read and validate the cometBFT values
	cometBFTConfig.SetRoot(cometBFTConfig.RootDir)
	ccfg.EnsureRoot(cometBFTConfig.RootDir)
	if err := cometBFTConfig.ValidateBasic(); err != nil {
		return fmt.Errorf("error in config file: %v", err)
	}
	if warnings := cometBFTConfig.CheckDeprecated(); len(warnings) > 0 {
		for _, warning := range warnings {
			clogger.Info("deprecated usage found in configuration file", "usage", warning)
		}
	}
	return nil
}

func writeConfigFile(fname string, config *ocfg.OmniConfig) error {
	out, err := yaml.Marshal(config)
	if err != nil {
		return err
	}
	if err := os.WriteFile(fname, out, ocfg.DefaultDirPerm); err != nil {
		return err
	}
	return nil
}

func getLogger(rcvdLogger log.Logger, logLevel string, defaultVal string, module string) (log.Logger, error) {
	logger, err := cflags.ParseLogLevel(logLevel, rcvdLogger, defaultVal)
	if err != nil {
		return nil, err
	}
	logger = logger.With("module", module)
	return logger, nil
}
