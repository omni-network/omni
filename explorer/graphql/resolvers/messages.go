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
	Offset        hexutil.Big
}

type XMsgsArgs struct {
	Limit  *hexutil.Big
	Cursor *hexutil.Big
}

const MsgsLimit = 25

func (b *BlocksResolver) XMsg(ctx context.Context, args XMsgArgs) (*XMsg, error) {
	res, found, err := b.Provider.XMsg(ctx, args.SourceChainID.ToInt().Uint64(), args.DestChainID.ToInt().Uint64(), args.Offset.ToInt().Uint64())
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

	if limit > MsgsLimit || args.Limit == nil {
		limit = MsgsLimit
	}

	if args.Cursor != nil {
		c := args.Cursor.ToInt().Uint64()
		cursor = &c
	}

	res, found, err := b.Provider.XMsgs(ctx, limit, cursor)
	if err != nil {
		return nil, errors.New("failed to fetch messages")
	}
	if !found {
		return nil, errors.New("messages not found")
	}

	return res, nil
}
