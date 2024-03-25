package resolvers

import (
	"context"

	"github.com/omni-network/omni/lib/errors"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

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
