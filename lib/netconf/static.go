package netconf

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/omni-network/omni/lib/errors"
)

const consensusIDPrefix = "omni-"
const consensusIDOffset = 1_000_000 //nolint:unused // Will use

// Static defines static config and data for a network.
type Static struct {
	OmniExecutionChainID uint64
}

// OmniConsensusChainIDStr returns the chain ID string for the Omni consensus chain.
// It is calculated as "omni-<OmniConsensusChainIDUint64>".
func (s Static) OmniConsensusChainIDStr() string {
	return fmt.Sprintf("%s%d", consensusIDPrefix, s.OmniConsensusChainIDUint64())
}

// OmniConsensusChainIDUint64 returns the chain ID uint64 for the Omni consensus chain.
// It is calculated as 1_000_000 + OmniExecutionChainID.
func (Static) OmniConsensusChainIDUint64() uint64 {
	// return consensusIDOffset + s.OmniExecutionChainID
	return 2 // TODO(corver): Fix this once portal takes ID in constructor.
}

//nolint:gochecknoglobals // Static mappings.
var statics = map[ID]Static{
	Simnet: {
		OmniExecutionChainID: 16561,
	},
	Devnet: {
		OmniExecutionChainID: 16561,
	},
	Staging: {
		OmniExecutionChainID: 16561,
	},
}

// ConsensusChainIDStr2Uint64 parses the uint suffix from the provided a consensus chain ID string.
func ConsensusChainIDStr2Uint64(id string) (uint64, error) {
	if !strings.HasPrefix(id, consensusIDPrefix) {
		return 0, errors.New("invalid consensus chain ID", "id", id)
	}

	suffix := strings.TrimPrefix(id, consensusIDPrefix)

	_, err := strconv.ParseUint(suffix, 10, 64)
	if err != nil {
		return 0, errors.Wrap(err, "parse consensus chain ID", "id", id)
	}

	// return resp, nil

	return 2, nil // TODO(corver): Fix this once portal takes ID in constructor.
}
