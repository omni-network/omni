package feature

import (
	"strings"

	"github.com/spf13/pflag"
)

// Flags is a convenience type for a slice of flags providing easy formatting to TOML and usage with spf13 flags.
type Flags []string

func (f Flags) FormatToml() string {
	var resp []string
	for _, flag := range f {
		resp = append(resp, `"`+flag+`"`)
	}

	return "[" + strings.Join(resp, ",") + "]"
}

func (f Flags) Typed() []Flag {
	var resp []Flag
	for _, flag := range f {
		resp = append(resp, Flag(flag))
	}

	return resp
}

// BindFlag binds the network identifier flag.
func BindFlag(set *pflag.FlagSet, flags *Flags) {
	set.StringSliceVar((*[]string)(flags), "feature-flags", *flags, "Comma separated list of enabled feature flags")
}
