package commands

import (
	_ "embed"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"path/filepath"
	"strings"
)

// nolint:gochecknoinits
func init() {
	cobra.EnableCommandSorting = false
}

type command struct {
	root    *cobra.Command
	config  *viper.Viper
	cfgFile string
	homeDir string
}

type option func(*command)

//go:embed omni-welcome-message.txt
var omniWelcomeMessage string

func newCommand(opts ...option) (c *command, err error) {
	c = &command{
		root: &cobra.Command{
			Use:           "omni",
			Short:         "Omni Network Node",
			SilenceErrors: true,
			SilenceUsage:  true,
			PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
				fmt.Print(omniWelcomeMessage)
				return c.initConfig()
			},
		},
	}
	c.initVersionCmd()
	return c, nil
}

func (c *command) initConfig() (err error) {
	config := viper.New()
	configName := ".omni"
	if c.cfgFile != "" {
		// Use config file from the flag.
		config.SetConfigFile(c.cfgFile)
	} else {
		// Search config in home directory with name ".omni" (without extension).
		config.AddConfigPath(c.homeDir)
		config.SetConfigName(configName)
	}

	// Environment
	config.SetEnvPrefix("omni")
	config.AutomaticEnv() // read in environment variables that match
	config.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))

	if c.homeDir != "" && c.cfgFile == "" {
		c.cfgFile = filepath.Join(c.homeDir, configName+".yaml")
	}

	// If a config file is found, read it in.
	if err := config.ReadInConfig(); err != nil {
		var e viper.ConfigFileNotFoundError
		if !errors.As(err, &e) {
			return err
		}
	}
	c.config = config
	return nil
}

// Execute parses command line arguments and runs appropriate functions.
func Execute() (err error) {
	c, err := newCommand()
	if err != nil {
		return err
	}
	return c.Execute()
}

func (c *command) Execute() (err error) {
	return c.root.Execute()
}
