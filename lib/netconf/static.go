package netconf

// Static defines static config and data for a network.
type Static struct {
	OmniExecutionChainID uint64
	OmniConsensusChainID uint64
}

//nolint:gochecknoglobals // Static mappings.
var statics = map[string]Static{
	Simnet: {
		OmniExecutionChainID: 16561,
		OmniConsensusChainID: 2,
	},
	Devnet: {
		OmniExecutionChainID: 16561,
		OmniConsensusChainID: 2,
	},
	Staging: {
		OmniExecutionChainID: 16561,
		OmniConsensusChainID: 2,
	},
}

// GetStatic returns the static config for a network or the zero value if it does not exist.
func GetStatic(network string) Static {
	return statics[network]
}
