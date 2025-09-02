package feeoraclev2

import (
	"github.com/omni-network/omni/lib/tokens"
)

var (
	gasTokenIDs = map[tokens.Asset]uint16{
		tokens.NOM: 1,
		tokens.ETH: 2,
	}
)

func gasTokenID(t tokens.Asset) (uint16, bool) {
	id, ok := gasTokenIDs[t]
	return id, ok
}

func allGasTokens() map[tokens.Asset]uint16 {
	result := make(map[tokens.Asset]uint16, len(gasTokenIDs))
	for k, v := range gasTokenIDs {
		result[k] = v
	}

	return result
}
