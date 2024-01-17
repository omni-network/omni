package cmd

import "github.com/spf13/pflag"

const homeFlag = "home"

// BindHomeFlag binds the home flag to the given flag set.
// This is generally only required for apps that require multiple config files or persist data to disk.
// It will result in the viper config directory to be updated from default "." to "<home>/config".
func BindHomeFlag(flags *pflag.FlagSet, homeDir *string) {
	flags.StringVar(homeDir, homeFlag, ".", "The application home directory containing config and data")
}
