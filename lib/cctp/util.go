package cctp

import (
	"context"
	"math/big"

	"github.com/omni-network/omni/lib/bi"
	cctpdb "github.com/omni-network/omni/lib/cctp/db"
	"github.com/omni-network/omni/lib/cctp/types"
	"github.com/omni-network/omni/lib/errors"
)

// GetInflightUSDC returns the amount of USDC inflight to `chainID`.
func GetInflightUSDC(ctx context.Context, db *cctpdb.DB, chainID uint64) (*big.Int, error) {
	msgs, err := db.GetMsgsBy(ctx, cctpdb.MsgFilter{
		DestChainID: chainID,
		Status:      types.MsgStatusSubmitted,
	})
	if err != nil {
		return nil, errors.Wrap(err, "get msgs")
	}

	total := bi.Zero()
	for _, msg := range msgs {
		total = bi.Add(total, msg.Amount)
	}

	return total, nil
}
