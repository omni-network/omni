package app

type Config struct {
	ListenAddr    string
	ChainID       uint64
	LoadState     string
	BlockTimeSecs uint64
	Silent        bool
	SlotsInEpoch  uint64
}

func DefaultConfig() Config {
	return Config{
		ListenAddr:    "0.0.0.0:8545",
		ChainID:       1337,
		LoadState:     "",
		Silent:        true,
		BlockTimeSecs: 1,
		SlotsInEpoch:  32,
	}
}
