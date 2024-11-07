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
// At 2024-11-06 10:45:44.637 UTC.
var bootstrapCursors = map[netconf.ID][]bootstrapCursor{
	netconf.Omega: {
		{"arb_sepolia|F", "base_sepolia", 10119},
		{"arb_sepolia|F", "holesky", 10119},
		{"arb_sepolia|F", "omni_evm", 10119},
		{"arb_sepolia|F", "op_sepolia", 10119},
		{"arb_sepolia|L", "base_sepolia", 10120},
		{"arb_sepolia|L", "holesky", 10120},
		{"arb_sepolia|L", "omni_evm", 10120},
		{"arb_sepolia|L", "op_sepolia", 10120},
		{"base_sepolia|F", "arb_sepolia", 8387},
		{"base_sepolia|F", "holesky", 8387},
		{"base_sepolia|F", "omni_evm", 8387},
		{"base_sepolia|F", "op_sepolia", 8387},
		{"base_sepolia|L", "arb_sepolia", 8388},
		{"base_sepolia|L", "holesky", 8388},
		{"base_sepolia|L", "omni_evm", 8388},
		{"base_sepolia|L", "op_sepolia", 8388},
		{"holesky|F", "arb_sepolia", 15471},
		{"holesky|F", "base_sepolia", 15471},
		{"holesky|F", "omni_evm", 15471},
		{"holesky|F", "op_sepolia", 15471},
		{"holesky|L", "arb_sepolia", 15473},
		{"holesky|L", "base_sepolia", 15473},
		{"holesky|L", "omni_evm", 15473},
		{"holesky|L", "op_sepolia", 15473},
		{"omni_consensus|F", "arb_sepolia", 39},
		{"omni_consensus|F", "base_sepolia", 39},
		{"omni_consensus|F", "holesky", 39},
		{"omni_consensus|F", "omni_evm", 39},
		{"omni_consensus|F", "op_sepolia", 39},
		{"omni_evm|F", "arb_sepolia", 19031},
		{"omni_evm|F", "base_sepolia", 19031},
		{"omni_evm|F", "holesky", 19031},
		{"omni_evm|F", "op_sepolia", 19031},
		{"op_sepolia|F", "arb_sepolia", 8896},
		{"op_sepolia|F", "base_sepolia", 8896},
		{"op_sepolia|F", "holesky", 8896},
		{"op_sepolia|F", "omni_evm", 8896},
		{"op_sepolia|L", "arb_sepolia", 8896},
		{"op_sepolia|L", "base_sepolia", 8896},
		{"op_sepolia|L", "holesky", 8896},
		{"op_sepolia|L", "omni_evm", 8896},
	},
}
