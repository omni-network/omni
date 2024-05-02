package netconf

import "github.com/spf13/pflag"

// BindFlag binds the network identifier flag.
func BindFlag(flags *pflag.FlagSet, network *ID) {
	flags.StringVar((*string)(network), "network", string(*network), "Omni network to participate in: mainnet, testnet, devnet")
}
