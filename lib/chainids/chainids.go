package chainids

const (
	// Mainnets.
	Ethereum = uint64(1)
	Omni     = uint64(166)

	// Testnets.
	OmniTestnet = uint64(165)
	Holesky     = uint64(17000)
	ArbSepolia  = uint64(421614)
	OpSepolia   = uint64(11155420)

	// Localnets.
	OmniDevnet = uint64(16561)
)

//nolint:gochecknoglobals // constant values
var (
	mainnets = []uint64{
		Ethereum,
		Omni,
	}

	testnets = []uint64{
		OmniTestnet,
		Holesky,
		ArbSepolia,
		OpSepolia,
	}
)

func IsMainnet(chainID uint64) bool {
	for _, mainnet := range mainnets {
		if chainID == mainnet {
			return true
		}
	}

	return false
}

func IsTestnet(chainID uint64) bool {
	for _, testnet := range testnets {
		if chainID == testnet {
			return true
		}
	}

	return false
}

func IsMainnetOrTestnet(chainID uint64) bool {
	return IsMainnet(chainID) || IsTestnet(chainID)
}
