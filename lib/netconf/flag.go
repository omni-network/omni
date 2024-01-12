package netconf

import "github.com/spf13/pflag"

// BindFlag binds the standard flag to provide network config path at runtime.
func BindFlag(flags *pflag.FlagSet, netConf *string) {
	flags.StringVar(netConf, "network-config", "network.json", "The path to the omni network config file")
}
