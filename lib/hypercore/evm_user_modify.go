package hypercore

import (
	"context"

	"github.com/omni-network/omni/lib/errors"
)

func (c client) UseBigBlocks(ctx context.Context) error {
	action := ActionEVMUserModify{
		Type:           "evmUserModify",
		UsingBigBlocks: true,
	}

	if err := c.do(ctx, action); err != nil {
		return errors.Wrap(err, "do action", "type", action.Type)
	}

	return nil
}
