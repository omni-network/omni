package cmd

import (
	"context"
	"net/url"
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

		if mapVal, err := flags.GetStringToString(f.Name); err == nil { // First check if it is a map flag
			// Redact each map value
			for k, v := range mapVal {
				mapVal[k] = redact(f.Name, v)
			}
			fields = append(fields, f.Name)
			fields = append(fields, mapVal)
		} else if arrayVal, err := flags.GetStringSlice(f.Name); err == nil { // Then check if it is a slice flag
			// Redact each slice element
			var vals []string
			for _, v := range arrayVal {
				vals = append(vals, redact(f.Name, v))
			}
			fields = append(fields, f.Name)
			fields = append(fields, vals)
		} else {
			fields = append(fields, f.Name)
			fields = append(fields, redact(f.Name, f.Value.String()))
		}
	})

	log.Info(ctx, "Parsed config from flags", fields...)

	return nil
}

// redact returns a redacted version of the given flag value. It currently supports redacting
// passwords in valid URLs as well as flags that contains words like "token", "password", "secret", "db" or "key".
func redact(flag, val string) string {
	// Don't redact empty flags ; i.e. show that they are empty.
	if val == "" {
		return ""
	}

	u, err := url.Parse(val)
	if err == nil {
		return u.Redacted()
	}

	// Don't redact --.*path or --.*file flags.
	if strings.Contains(flag, "file") ||
		strings.Contains(flag, "path") {
		return val
	}

	if strings.Contains(flag, "token") ||
		strings.Contains(flag, "password") ||
		strings.Contains(flag, "secret") ||
		strings.Contains(flag, "db") ||
		strings.Contains(flag, "header") ||
		strings.Contains(flag, "key") {
		return "xxxxx"
	}

	return val
}
