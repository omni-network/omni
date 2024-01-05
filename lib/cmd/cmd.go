// Package cmd provides a common utilities and helper function to standarise
// the way omni apps use cobra and viper to produce consistent cli experience
// for both users and devs.
package cmd

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// Main is the main entry point for the command line tool.
// Usage:
//
//	   func main() {
//		     libcmd.Main(appcmd.New())
//	   }
func Main(cmd *cobra.Command) {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

	err := cmd.ExecuteContext(ctx)

	cancel()

	if err != nil {
		log.Error(ctx, "Fatal error", err)

		const errExitCode = 1
		os.Exit(errExitCode) //nolint:revive // Deep exit is exactly the point of this helper function.
	}
}

// NewRootCmd returns a new root cobra command that handles our command line tool.
// It sets up the general viper config and binds the cobra flags to the viper flags.
// It also silences the usage printing when commands error during "running".
func NewRootCmd(appName string, appDescription string, subCmds ...*cobra.Command) *cobra.Command {
	root := &cobra.Command{
		Use:   appName,
		Short: appDescription,
		Args:  cobra.NoArgs,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return initializeConfig(appName, cmd)
		},
	}

	root.AddCommand(subCmds...)
	root.SilenceErrors = true // Disable default error printing.

	silenceUsage(root)

	return root
}

// silenceUsage silences the usage printing when commands error during "running",
// so only show usage if error occurs before that, e.g., when parsing flags.
func silenceUsage(cmd *cobra.Command) {
	if runFunc := cmd.RunE; runFunc != nil {
		cmd.RunE = func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			return runFunc(cmd, args)
		}
	}

	for _, cmd := range cmd.Commands() {
		silenceUsage(cmd)
	}
}

// initializeConfig sets up the general viper config and binds the cobra flags to the viper flags.
func initializeConfig(appName string, cmd *cobra.Command) error {
	v := viper.New()

	v.SetConfigName(appName)
	v.AddConfigPath(".")

	// Attempt to read the config file, gracefully ignoring errors
	// caused by a config file not being found. Return an error
	// if we cannot parse the config file.
	if err := v.ReadInConfig(); err != nil {
		// It's okay if there isn't a config file
		var cfgError viper.ConfigFileNotFoundError
		if ok := errors.As(err, &cfgError); !ok {
			return errors.Wrap(err, "read config")
		}
	}

	v.SetEnvPrefix(appName)
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))

	// Bind the current command's flags to viper
	return bindFlags(cmd, v)
}

// bindFlags binds each cobra flag to its associated viper configuration (config file and environment variable).
func bindFlags(cmd *cobra.Command, v *viper.Viper) error {
	var lastErr error

	cmd.Flags().VisitAll(func(f *pflag.Flag) {
		// Cobra provided flags take priority
		if f.Changed {
			return
		}

		// Define all the viper flag names to check
		viperNames := []string{
			f.Name,
			strings.ReplaceAll(f.Name, "_", "."), // TOML uses "." to indicate hierarchy, while we use "_".
		}

		for _, name := range viperNames {
			if !v.IsSet(name) {
				continue
			}

			val := v.Get(name)
			err := cmd.Flags().Set(f.Name, fmt.Sprintf("%v", val))
			if err != nil {
				lastErr = err
			}

			break
		}
	})

	return lastErr
}
