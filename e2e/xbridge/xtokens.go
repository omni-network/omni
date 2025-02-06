package xbridge

import (
	"github.com/omni-network/omni/e2e/xbridge/rlusd"
	"github.com/omni-network/omni/e2e/xbridge/types"
)

// Tokens returns the list of currently managed xtokens.
func Tokens() []types.XToken { return []types.XToken{rlusd.XToken()} }
