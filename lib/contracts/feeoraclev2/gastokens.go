package feeoraclev2

import (
	"github.com/omni-network/omni/lib/tokens"
)

var (
	gasTokenIDs = map[tokens.Token]uint8{
		tokens.OMNI: 1,
		tokens.ETH:  2,
	}
)

func gasTokenID(t tokens.Token) (uint8, bool) {
	id, ok := gasTokenIDs[t]
	return id, ok
}

func allGasTokens() map[tokens.Token]uint8 {
	result := make(map[tokens.Token]uint8, len(gasTokenIDs))
	for k, v := range gasTokenIDs {
		result[k] = v
	}

	return result
}
