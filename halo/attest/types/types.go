package types

// ChainNameFunc is a function that returns the name of a chain given its ID.
type ChainNameFunc func(chainID uint64) string
