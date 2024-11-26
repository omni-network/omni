package feeoraclev2

import (
	"sort"

	"github.com/omni-network/omni/lib/evmchain"
)

const defaultBaseGasLimit = 100_000
const defaultBaseDataBuffer = 100
const defaultGasPerByte = 16

const (
	// Mainnets.
	IDEthereum uint64 = 1

	// Testnets.
	IDSepolia uint64 = 11155111
)

type Config struct {
	ID             uint64
	BaseGasLimit   uint32
	BaseDataBuffer uint32
	GasPerByte     uint64
}

func initStatic() map[uint64]Config {
	configs := make(map[uint64]Config, len(evmchain.All()))
	for _, metadata := range evmchain.All() {
		config := Config{
			ID:           metadata.ChainID,
			BaseGasLimit: defaultBaseGasLimit,
		}

		// If the chain doesn't post data to another chain, add data cost configs
		if metadata.PostsTo == 0 {
			config.BaseDataBuffer = defaultBaseDataBuffer
			config.GasPerByte = defaultGasPerByte
		}

		configs[metadata.ChainID] = config
	}

	return configs
}

func GetConfig(chainID uint64) (Config, bool) {
	cfg, ok := static[chainID]
	return cfg, ok
}

func AllConfigs() []Config {
	var resp []Config
	for _, cfg := range static {
		resp = append(resp, cfg)
	}

	sort.Slice(resp, func(i, j int) bool {
		return resp[i].ID < resp[j].ID
	})

	return resp
}

var static = initStatic()
