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
