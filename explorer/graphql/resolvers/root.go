package resolvers

import (
	"context"
)

type Query struct {
	BlocksResolver
}

type BlocksProvider interface {
	XBlock(ctx context.Context, sourceChainID uint64, height uint64) (*XBlock, bool, error)
	XMsg(ctx context.Context, sourceChainID uint64, destChainID uint64, offset uint64) (*XMsg, bool, error)
	XMsgs(ctx context.Context, Limit uint64, cursor *uint64) (*XMsgResult, bool, error)
	XReceipt(ctx context.Context, sourceChainID, destChainID, offset uint64) (*XReceipt, bool, error)
	SupportedChains(ctx context.Context) ([]*Chain, bool, error)
	Search(ctx context.Context, query string) (*SearchResult, bool, error)
}

type BlocksResolver struct {
	Provider BlocksProvider
}
