package feeoraclev2

import (
	"github.com/omni-network/omni/lib/tokenmeta"
)

var (
	gasTokenIDs = map[tokenmeta.Meta]uint16{
		tokenmeta.OMNI: 1,
		tokenmeta.ETH:  2,
	}
)

func gasTokenID(t tokenmeta.Meta) (uint16, bool) {
	id, ok := gasTokenIDs[t]
	return id, ok
}

func allGasTokens() map[tokenmeta.Meta]uint16 {
	result := make(map[tokenmeta.Meta]uint16, len(gasTokenIDs))
	for k, v := range gasTokenIDs {
		result[k] = v
	}

	return result
}
