package resolvers

import (
	"context"

	"github.com/omni-network/omni/lib/errors"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

type XMsgRangeArgs struct {
	From hexutil.Big
	To   hexutil.Big
}

type XMsgArgs struct {
	SourceChainID hexutil.Big
	DestChainID   hexutil.Big
	StreamOffset  hexutil.Big
}

type XMsgsArgs struct {
	Limit  *hexutil.Big
	Cursor *hexutil.Big
}

func (b *BlocksResolver) XMsgCount(ctx context.Context) (*hexutil.Big, error) {
	res, found, err := b.BlocksProvider.XMsgCount(ctx)
	if err != nil {
		return nil, errors.New("failed to fetch message count")
	}
	if !found {
		return nil, errors.New("message count not found")
	}

	return res, nil
}

func (b *BlocksResolver) XMsgRange(ctx context.Context, args XMsgRangeArgs) ([]*XMsg, error) {
	if args.From.ToInt().Cmp(args.To.ToInt()) >= 0 {
		return nil, errors.New("invalid range")
	}

	res, found, err := b.BlocksProvider.XMsgRange(ctx, args.From.ToInt().Uint64(), args.To.ToInt().Uint64())
	if err != nil {
		return nil, errors.New("failed to fetch messages")
	}
	if !found {
		return nil, errors.New("messages not found")
	}

	return res, nil
}

func (b *BlocksResolver) XMsg(ctx context.Context, args XMsgArgs) (*XMsg, error) {
	res, found, err := b.BlocksProvider.XMsg(ctx, args.SourceChainID.ToInt().Uint64(), args.DestChainID.ToInt().Uint64(), args.StreamOffset.ToInt().Uint64())
	if err != nil {
		return nil, errors.New("failed to fetch message")
	}
	if !found {
		return nil, errors.New("message not found")
	}

	return res, nil
}

func (b *BlocksResolver) XMsgs(ctx context.Context, args XMsgsArgs) (*XMsgResult, error) {
	limit := uint64(1)
	var cursor *uint64

	if args.Limit != nil {
		limit = args.Limit.ToInt().Uint64()
	}

	if args.Cursor != nil {
		c := args.Cursor.ToInt().Uint64()
		cursor = &c
	}

	res, found, err := b.BlocksProvider.XMsgs(ctx, limit, cursor)
	if err != nil {
		return nil, errors.New("failed to fetch messages")
	}
	if !found {
		return nil, errors.New("messages not found")
	}

	return res, nil
}
