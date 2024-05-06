package resolvers

import (
	"context"

	"github.com/omni-network/omni/lib/errors"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

type XReceiptArgs struct {
	SourceChainID hexutil.Big
	DestChainID   hexutil.Big
	Offset        hexutil.Big
}

func (b *BlocksResolver) XReceipt(ctx context.Context, args XReceiptArgs) (*XReceipt, error) {
	res, found, err := b.Provider.XReceipt(ctx, args.SourceChainID.ToInt().Uint64(), args.DestChainID.ToInt().Uint64(), args.Offset.ToInt().Uint64())
	if err != nil {
		return nil, errors.New("failed to fetch receipt")
	}
	if !found {
		return nil, errors.New("receipt not found")
	}

	return res, nil
}
