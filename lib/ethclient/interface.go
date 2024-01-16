package ethclient

import (
	"context"

	"github.com/omni-network/omni/lib/xchain"

	"github.com/ethereum/go-ethereum/core/types"
)

// Client is the interface to the rpc connection module that connects to the
// chain and provides information.
type Client interface {
	// GetBlock fetches the XBlock from the event logs given a block height
	GetBlock(ctx context.Context, height uint64, header *types.Header) (xchain.Block, bool, error)

	// GetCurrentFinalisedBlockHeader delivers the recently finalized block header
	GetCurrentFinalisedBlockHeader(ctx context.Context) (*types.Header, error)
}
