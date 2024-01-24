package cmd

import (
	"context"
	"strings"

	"github.com/omni-network/omni/lib/log"

	"github.com/spf13/pflag"
)

const homeFlag = "home"

// BindHomeFlag binds the home flag to the given flag set.
// This is generally only required for apps that require multiple config files or persist data to disk.
// Using this flag will result in the viper config directory to be updated from default "." to "<home>/config".
func BindHomeFlag(flags *pflag.FlagSet, homeDir *string) {
	flags.StringVar(homeDir, homeFlag, *homeDir, "The application home directory containing config and data")
}

// LogFlags logs the configured flags kv pairs.
func LogFlags(ctx context.Context, flags *pflag.FlagSet) error {
	skip := map[string]bool{
		"help": true,
	}
	// Flatten config into key/value pairs for logging.
	var fields []any
	flags.VisitAll(func(f *pflag.Flag) {
		if skip[f.Name] {
			return
		}
		// TODO(corver): Allow dashes for one-to-one mapping with actual flags?
		fields = append(fields, strings.ReplaceAll(f.Name, "-", "_"))
		fields = append(fields, f.Value)
	})

	log.Info(ctx, "Parsed config from flags", fields...)

	return nil
}
