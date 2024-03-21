package chainids

const (
	// mainnets.
	Ethereum = uint64(1)

	// testnets.
	Goerli     = uint64(5)
	Holesky    = uint64(17000)
	ArbSepolia = uint64(421614)
	OpSepolia  = uint64(11155420)
)

//nolint:gochecknoglobals // constant values
var (
	mainnets = []uint64{
		Ethereum,
	}

	testnets = []uint64{
		Goerli,
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
