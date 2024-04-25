package resolvers

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

type Query struct {
	BlocksResolver
}

type BlocksProvider interface { //nolint: interfacebloat // We need this interface to define all of our methods for the library we are using
	XBlock(ctx context.Context, sourceChainID uint64, height uint64) (*XBlock, bool, error)
	XBlockRange(ctx context.Context, amount uint64, offset uint64) ([]*XBlock, bool, error)
	XBlockCount(ctx context.Context) (*hexutil.Big, bool, error)
	XMsgCount(ctx context.Context) (*hexutil.Big, bool, error)
	XMsgRange(ctx context.Context, amount uint64, offset uint64) ([]*XMsg, bool, error)
	XMsg(ctx context.Context, sourceChainID uint64, destChainID uint64, streamOffset uint64) (*XMsg, bool, error)
	XMsgs(ctx context.Context, limit uint64, cursor, sourceChainID, destChainID *uint64, address *common.Address) (*XMsgResult, bool, error)
	XReceiptCount(ctx context.Context) (*hexutil.Big, bool, error)
	XReceipt(ctx context.Context, sourceChainID, destChainID, streamOffset uint64) (*XReceipt, bool, error)
	SupportedChains(ctx context.Context) ([]*Chain, bool, error)
	Search(ctx context.Context, query string) (*SearchResult, bool, error)
}

type BlocksResolver struct {
	BlocksProvider BlocksProvider
}
