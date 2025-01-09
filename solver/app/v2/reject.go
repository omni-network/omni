package appv2

import (
	"context"

	"github.com/omni-network/omni/lib/netconf"
)

//go:generate stringer -type=rejectReason -trimprefix=reject
type rejectReason uint8

const (
	rejectNone                  rejectReason = 0
	rejectDestCallReverts       rejectReason = 1
	rejectInsufficientFee       rejectReason = 2
	rejectInsufficientInventory rejectReason = 3
	rejectInvalidTarget         rejectReason = 4
)

func newShouldRejector(_ netconf.ID) func(ctx context.Context, chainID uint64, order Order) (rejectReason, bool, error) {
	return func(_ context.Context, _ uint64, _ Order) (rejectReason, bool, error) {
		// TODO: check liquidity, reject
		return rejectNone, false, nil
	}
}
