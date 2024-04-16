package resolvers

import (
	"context"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

type Query struct {
	BlocksResolver
}

type BlocksProvider interface {
	XBlock(ctx context.Context, SourceChainID uint64, Height uint64) (*XBlock, bool, error)
	XBlockRange(ctx context.Context, Amount uint64, Offset uint64) ([]*XBlock, bool, error)
	XBlockCount(ctx context.Context) (*hexutil.Big, bool, error)
	XMsgCount(ctx context.Context) (*hexutil.Big, bool, error)
	XReceiptCount(ctx context.Context) (*hexutil.Big, bool, error)
	XMsgRange(ctx context.Context, Amount uint64, Offset uint64) ([]*XMsg, bool, error)
	XMsg(ctx context.Context, SourceChainID uint64, DestChainID uint64, StreamOffset uint64) (*XMsg, bool, error)
	XReceipt(ctx context.Context, SourceChainID, DestChainID, StreamOffset uint64) (*XReceipt, bool, error)
	SupportedChains(ctx context.Context) ([]*Chain, bool, error)
	Search(ctx context.Context, query string) (*SearchResult, bool, error)
}

type BlocksResolver struct {
	BlocksProvider BlocksProvider
}
