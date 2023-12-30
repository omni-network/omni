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
	"io"
	"os"
	"path/filepath"
	"strings"

	ccfg "github.com/cometbft/cometbft/config"
	cflags "github.com/cometbft/cometbft/libs/cli/flags"
	"github.com/cometbft/cometbft/libs/log"
	ocfg "github.com/omni-network/omni/pkg/config"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	config *Config //configurations for omni and cometBFT
)

// command options
const (
	envPrefix             = "OMNI"
	optionNameHomeDir     = "home-dir"
	optionNameLogLevel    = "log-level"
	optionNameLogWriter   = "log-writer"
	optionNameOmniRootDir = "root-dir"
	optionNameVerbose     = "verbose"
)

type Config struct {
	rootCmd        *cobra.Command   // the cobra command which has all the commandline
	viperConfig    *viper.Viper     // viper for reading config from files
	omniConfig     *ocfg.OmniConfig // omni related configuration
	cometBFTConfig *ccfg.Config     // cometBFT related configuration
	cmdWriter      io.Writer        // cobra writer to test output
	ologger        log.Logger       // logger used in omni
	clogger        log.Logger       // logger used in cometBFT
}

func setPersistentFlags(cmd *cobra.Command) {
	// set the persistent flags
	cmd.PersistentFlags().String(optionNameHomeDir, os.ExpandEnv("$HOME"), "home directory for creating omni root")
	cmd.PersistentFlags().String(optionNameLogLevel, ocfg.DefaultLogLevel, "log level (info, debug, error, none)")
	cmd.PersistentFlags().String(optionNameLogWriter, ocfg.DefaultLogWriter, "where to write the log output")
}

// NewRootCommand  is the root command for omni node
func NewRootCommand() (*cobra.Command, error) {
	RootCmd := &cobra.Command{
		Use:           "omni",
		Short:         "Omni node hosting halo protocol for Xchain messages",
		SilenceErrors: true,
		SilenceUsage:  true,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			// construct the configuration for omni and cometBFT
			var err error
			config, err = initializeConfig(cmd)
			if err != nil {
				return err
			}

			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {},
	}
	setPersistentFlags(RootCmd)
	return RootCmd, nil
}

func initializeConfig(cmd *cobra.Command) (*Config, error) {
	//
	//  Step-1: load the default values for omni and cometBFT
	//
	defaultConfig, err := loadDefaultConfig(cmd)
	if err != nil {
		return nil, err
	}
	v := defaultConfig.viperConfig
	co := defaultConfig.omniConfig

	//
	// Step-2: override the default by loading the values from the config files
	writeConfigFileFlag, err := loadConfigFromFile(cmd, defaultConfig)
	if err != nil {
		return nil, err
	}

	//
	// step-3: - load from environment variables
	//
	v.SetEnvPrefix(envPrefix)
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

	//
	// step-4: load the resolved viper values to cobra, except the command line ones
	//
	bindFlags(cmd, v)

	// if a config file is not present for omni, write it
	if writeConfigFileFlag {
		configFileName := filepath.Join(co.OmniConfigDir, co.OmniConfigFileName+".yml")
		if err := writeConfigFile(configFileName, co); err != nil {
			return nil, err
		}
	}
	return defaultConfig, nil
}

// Bind each cobra flag to its associated viper configuration (config file and environment variable)
func bindFlags(cmd *cobra.Command, v *viper.Viper) {
	cmd.Flags().VisitAll(func(f *pflag.Flag) {
		// Determine the naming convention of the flags when represented in the config file
		configName := f.Name

		// Apply the viper config value to the flag when the flag is not set and viper has a value
		if !f.Changed && v.IsSet(configName) {
			val := v.Get(configName)
			err := cmd.Flags().Set(f.Name, fmt.Sprintf("%v", val))
			if err != nil {
				return
			}
		}
	})
}

func loadDefaultConfig(cmd *cobra.Command) (*Config, error) {
	// check if the home-dir is passed through the command line
	commandLineHome, err2 := cmd.Flags().GetString(optionNameHomeDir)
	if err2 != nil {
		return nil, err2
	}

	// create default Omni config
	oConfig, err := ocfg.LoadDefaultConfig(commandLineHome)
	if err != nil {
		return nil, err
	}

	// create default cometBFT config
	cConfig := ccfg.DefaultConfig()

	// overwrite few defaults in cometbft config
	cConfig.Moniker = oConfig.NodeName                                                // cometBFT moniker
	cConfig.LogLevel = oConfig.LogLevel                                               // cometBFT loglevel
	cConfig.RootDir = filepath.Join(oConfig.HomeDirectory, ccfg.DefaultTendermintDir) // cometBFT rootDir

	// initialise the omni logger
	defaultLogger := log.NewTMLogger(log.NewSyncWriter(os.Stdout))
	omniLogger, err := getLogger(defaultLogger, oConfig.LogLevel, ocfg.DefaultLogLevel, "omni")
	if err != nil {
		return nil, err
	}

	// initialise the cometBFT logger
	defaultLogger = log.NewTMLogger(log.NewSyncWriter(os.Stdout))
	cometBFTLogger, err := getLogger(defaultLogger, cConfig.LogLevel, ocfg.DefaultLogLevel, "comt")
	if err != nil {
		return nil, err
	}

	// create the base default config instance
	return &Config{
		rootCmd:        cmd,
		viperConfig:    viper.New(),       // set a new viper instance
		omniConfig:     oConfig,           // set a default omni instance
		cometBFTConfig: cConfig,           // set a default configBFT instance
		cmdWriter:      cmd.OutOrStdout(), // used in unit test cases to capture the stdout
		ologger:        omniLogger,
		clogger:        cometBFTLogger,
	}, nil
}

func setHomeAndRootDirForOmniFromCommandLine(cmd *cobra.Command, defaultConfig *Config) {
	// parse the omni home dir if it is given in the command line
	hDir, err := cmd.Flags().GetString(optionNameHomeDir)
	if err == nil {
		defaultConfig.omniConfig.HomeDirectory = hDir
	}
	// parse the omni root dir if it is given in the command line
	rDir, err := cmd.Flags().GetString(optionNameOmniRootDir)
	if err == nil {
		defaultConfig.omniConfig.OmniRootDir = filepath.Join(defaultConfig.omniConfig.HomeDirectory, rDir)
	}
}

func loadConfigFromFile(cmd *cobra.Command, defaultConfig *Config) (bool, error) {
	writeOmniConfigToFileFlag := false

	v := defaultConfig.viperConfig
	co := defaultConfig.omniConfig
	cc := defaultConfig.cometBFTConfig

	// set the proper home and root directories before reading the config file
	setHomeAndRootDirForOmniFromCommandLine(cmd, defaultConfig)

	// make sure the directories exists
	err := ocfg.EnsureDirectories(defaultConfig.omniConfig)
	if err != nil {
		return false, err
	}

	// read the omni config file
	v.SetConfigName(co.OmniConfigFileName)
	v.AddConfigPath(co.OmniConfigDir)
	if err := v.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if !errors.As(err, &configFileNotFoundError) {
			return false, err
		} else {
			// omni config file not found, write it later after resolving all the values
			writeOmniConfigToFileFlag = true
		}
	}

	// read and validate the cometBFT values
	cc.SetRoot(cc.RootDir)
	ccfg.EnsureRoot(cc.RootDir)
	if err := cc.ValidateBasic(); err != nil {
		return false, fmt.Errorf("error in config file: %v", err)
	}
	if warnings := cc.CheckDeprecated(); len(warnings) > 0 {
		for _, warning := range warnings {
			config.clogger.Info("deprecated usage found in configuration file", "usage", warning)
		}
	}

	return writeOmniConfigToFileFlag, nil
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
