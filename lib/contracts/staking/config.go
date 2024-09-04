package staking

import (
	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/common"
)

type Config struct {
	Allowlist        []common.Address
	AllowlistEnabled bool
}

var statics = map[netconf.ID]Config{
	netconf.Omega: {
		Allowlist:        []common.Address{},
		AllowlistEnabled: true,
	},
}

func ConfigByNetwork(id netconf.ID) (Config, bool) {
	cfg, ok := statics[id]
	return cfg, ok
}
