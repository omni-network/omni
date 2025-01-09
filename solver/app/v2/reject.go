package appv2

import (
	"context"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/log"
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

// newShouldRejector returns as ShouldReject function for the given network.
//
// ShouldReject returns true and a reason if the request should be rejected.
// It returns false if the request should be accepted.
// Errors are unexpected and refer to internal server problems.
func newShouldRejector(network netconf.ID) func(ctx context.Context, chainID uint64, req bindings.SolveRequest) (rejectReason, bool, error) {
	return func(ctx context.Context, srcChainID uint64, req bindings.SolveRequest) (rejectReason, bool, error) {
		// reject swallows the error (only logging it) and returns true and the reject reason.
		reject := func(reason rejectReason, err error) (rejectReason, bool, error) {
			log.InfoErr(ctx, "Rejecting request", err, "reason", reason)
			return reason, true, nil
		}

		target, err := getTarget(network, req.Call)
		if err != nil {
			return reject(rejectInvalidTarget, err)
		}

		if err := target.Verify(srcChainID, req.Call, req.Deposits); err != nil {
			return reject(rejectInsufficientInventory, err) // TODO(corver): Fix reason
		}

		return rejectNone, false, nil
	}
}
