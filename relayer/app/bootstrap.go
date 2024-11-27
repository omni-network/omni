package relayer

import (
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"
)

func getBootstrapOffset(network netconf.Network, destChain netconf.Chain, chainVer xchain.ChainVersion) (uint64, bool) {
	chainVerName := network.ChainVersionName(chainVer)
	for _, b := range bootstrapCursors[network.ID] {
		if b.Worker == destChain.Name && b.ChainVersion == chainVerName {
			return b.AttestOffset, true
		}
	}

	return 0, false
}

type bootstrapCursor struct {
	ChainVersion string
	Worker       string
	AttestOffset uint64
}

// bootstrapCursors obtained from grafana
// Instant value of `max(lib_cprovider_stream_offset{network="omega",job="relayer"}) by (chain_version,worker)`
// At 2024-11-26 19:42:54.227 UTC.
var bootstrapCursors = map[netconf.ID][]bootstrapCursor{
	netconf.Omega: {
		{"arb_sepolia|F", "base_sepolia", 12415},
		{"arb_sepolia|F", "holesky", 12415},
		{"arb_sepolia|F", "omni_evm", 12415},
		{"arb_sepolia|F", "op_sepolia", 12415},
		{"arb_sepolia|L", "base_sepolia", 12416},
		{"arb_sepolia|L", "holesky", 12416},
		{"arb_sepolia|L", "omni_evm", 12416},
		{"arb_sepolia|L", "op_sepolia", 12416},
		{"base_sepolia|F", "arb_sepolia", 10139},
		{"base_sepolia|F", "holesky", 10139},
		{"base_sepolia|F", "omni_evm", 10139},
		{"base_sepolia|F", "op_sepolia", 10139},
		{"base_sepolia|L", "arb_sepolia", 10139},
		{"base_sepolia|L", "holesky", 10139},
		{"base_sepolia|L", "omni_evm", 10139},
		{"base_sepolia|L", "op_sepolia", 10139},
		{"holesky|F", "arb_sepolia", 18388},
		{"holesky|F", "base_sepolia", 18388},
		{"holesky|F", "omni_evm", 18388},
		{"holesky|F", "op_sepolia", 18388},
		{"holesky|L", "arb_sepolia", 18389},
		{"holesky|L", "base_sepolia", 18389},
		{"holesky|L", "omni_evm", 18389},
		{"holesky|L", "op_sepolia", 18389},
		{"omni_consensus|F", "arb_sepolia", 40},
		{"omni_consensus|F", "base_sepolia", 40},
		{"omni_consensus|F", "holesky", 40},
		{"omni_consensus|F", "omni_evm", 40},
		{"omni_consensus|F", "op_sepolia", 40},
		{"omni_evm|F", "arb_sepolia", 23015},
		{"omni_evm|F", "base_sepolia", 23015},
		{"omni_evm|F", "holesky", 23015},
		{"omni_evm|F", "op_sepolia", 23015},
		{"op_sepolia|F", "arb_sepolia", 10764},
		{"op_sepolia|F", "base_sepolia", 10764},
		{"op_sepolia|F", "holesky", 10764},
		{"op_sepolia|F", "omni_evm", 10758},
		{"op_sepolia|L", "arb_sepolia", 10764},
		{"op_sepolia|L", "base_sepolia", 10764},
		{"op_sepolia|L", "holesky", 10764},
		{"op_sepolia|L", "omni_evm", 10754},
	},
}
