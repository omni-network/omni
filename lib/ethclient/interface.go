package ethclient

import (
	"context"

	"github.com/omni-network/omni/lib/xchain"
)

// Client is the interface to the rpc connection module that connects to the
// chain and provides information.
type Client interface {
	// GetBlock fetches the XBlock from the event logs given a block height
	GetBlock(ctx context.Context, height uint64) (xchain.Block, bool, error)
}
