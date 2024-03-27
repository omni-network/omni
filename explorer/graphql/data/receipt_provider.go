package data

import (
	"context"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

func (p Provider) XReceiptCount(ctx context.Context) (*hexutil.Big, bool, error) {
	query, err := p.EntClient.Receipt.Query().
		Count(ctx)
	if err != nil {
		log.Error(ctx, "Graphql provider err", err)
		return nil, false, err
	}

	hex, err := Uint2Hex(uint64(query))
	if err != nil {
		return nil, false, errors.Wrap(err, "decoding block count")
	}

	return &hex, true, nil
}
