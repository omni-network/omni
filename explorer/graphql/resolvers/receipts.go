package resolvers

import (
	"context"

	"github.com/omni-network/omni/lib/errors"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

func (b *BlocksResolver) XReceiptCount(ctx context.Context) (*hexutil.Big, error) {
	res, found, err := b.BlocksProvider.XReceiptCount(ctx)
	if err != nil {
		return nil, errors.New("failed to fetch receipt count")
	}
	if !found {
		return nil, errors.New("receipt count not found")
	}

	return res, nil
}
