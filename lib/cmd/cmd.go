// Package cmd provides a common utilities and helper function to standarise
// the way omni apps use cobra and viper to produce consistent cli experience
// for both users and devs.
package cmd

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// Main is the main entry point for the omni application binaries.
// Usage:
//
//	   func main() {
//		     libcmd.Main(appcmd.New())
//	   }
func Main(cmd *cobra.Command) {
	wrapRunCmd(cmd)

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	err := cmd.ExecuteContext(ctx)
	cancel()

	if err != nil {
		const errExitCode = 1
		os.Exit(errExitCode) //nolint:revive // Deep exit is exactly the point of this helper function.
	}
}

// NewRootCmd returns a new root cobra command that handles our command line tool.
// It sets up the general viper config and binds the cobra flags to the viper flags.
func NewRootCmd(appName string, appDescription string, subCmds ...*cobra.Command) *cobra.Command {
	root := &cobra.Command{
		Use:   appName,
		Short: appDescription,
		Args:  cobra.NoArgs,
		PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {
			return initializeConfig(appName, cmd)
		},
		RunE: func(*cobra.Command, []string) error {
			// Callers should either add sub-commands or override RunE.
			return errors.New("no sub-command specified, see --help")
		},
	}

	root.AddCommand(subCmds...)

	return root
}

// SilenceErrUsage silences the usage and error printing.
func SilenceErrUsage(cmd *cobra.Command) {
	cmd.SilenceUsage = true
	cmd.SilenceErrors = true
	for _, cmd := range cmd.Commands() {
		SilenceErrUsage(cmd)
	}
}

// initializeConfig sets up the general viper config and binds the cobra flags to the viper flags.
func initializeConfig(appName string, cmd *cobra.Command) error {
	v := viper.New()

	v.SetConfigName(appName)

	// Set config path to <home>/config/ if --home flag is used in this app.
	if home := cmd.Flag(homeFlag); home != nil {
		v.AddConfigPath(filepath.Join(home.Value.String(), "config"))
	} else {
		// Otherwise, set config path to current directory
		v.AddConfigPath(".")
	}

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
			strings.Replace(f.Name, "-", ".", 1), // Support 1 tier of TOML groups using first term before "-".
		}

		for _, name := range viperNames {
			if !v.IsSet(name) {
				continue
			}

			// Special case handling of slice flags.
			if sliceVal, ok := f.Value.(pflag.SliceValue); ok {
				for _, v := range v.GetStringSlice(name) {
					err := sliceVal.Append(v)
					if err != nil {
						lastErr = errors.Wrap(err, "append flag", "name", f.Name, "type", f.Value.Type(), "value", v)
					}
				}

				break
			}

			val := v.Get(name)

			// Special case handling of map[string]string flags.
			if f.Value.Type() == "stringToString" {
				strMap := v.GetStringMapString(name)
				if len(strMap) == 0 {
					// There is no way to set an empty value for Cobra's map[string]string flags.
					// It must either not be set or be non-empty.
					// So skip empty viper maps (as if not set) assuming the default value is empty.
					continue
				}

				var kvs []string
				for k, v := range strMap {
					kvs = append(kvs, fmt.Sprintf("%s=%s", k, v))
				}

				val = strings.Join(kvs, ",")
				// Set flag below
			}

			err := f.Value.Set(fmt.Sprintf("%v", val))
			if err != nil {
				lastErr = errors.Wrap(err, "set flag", "name", f.Name, "type", f.Value.Type(), "value", val)
			}

			break
		}
	})

	return lastErr
}

// wrapRunCmd wraps the "app run" command to custom fatal error log and silence cobra output.
func wrapRunCmd(cmd *cobra.Command) {
	runCmd := getRunCmd(cmd)
	SilenceErrUsage(runCmd)
	runFunc := runCmd.RunE
	runCmd.RunE = func(cmd *cobra.Command, args []string) error {
		if runFunc == nil {
			return errors.New("run command RunE nil [BUG]")
		}

		err := runFunc(cmd, args)
		if err != nil {
			log.Error(cmd.Context(), "!! Fatal error occurred, app died !!", err)
		}

		return err
	}
}

// getRunCmd returns the "run" subcommand of the given command or the command itself.
func getRunCmd(cmd *cobra.Command) *cobra.Command {
	const name = "run"
	for _, sub := range cmd.Commands() {
		if sub.Use == name {
			return sub
		}
	}

	return cmd
}
