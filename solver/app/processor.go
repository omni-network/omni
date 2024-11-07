package solver

import (
	"context"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/ethereum/go-ethereum/core/types"
)

// newEventProcessor returns a callback provided to xchain.Provider::StreamEventLogs processing
// all inbox contract events and driving request lifecycle.
func newEventProcessor(deps procDeps, chainID uint64) xchain.EventLogsCallback {
	return func(ctx context.Context, _ uint64, elogs []types.Log) error {
		for _, elog := range elogs {
			event, ok := eventsByTopic[elog.Topics[0]]
			if !ok {
				return errors.New("unknown event [BUG]")
			}

			reqID, err := deps.ParseID(chainID, elog)
			if err != nil {
				return errors.Wrap(err, "parse id")
			}

			ctx := log.WithCtx(ctx, log.Hex7("req_id", reqID[:]))

			req, _, err := deps.GetRequest(ctx, chainID, reqID)
			if err != nil {
				return errors.Wrap(err, "current status")
			} else if event.Status != req.Status {
				// TODO(corver): Detect unexpected on-chain status.
				log.Info(ctx, "Ignoring mismatching old event", "actual", statusString(req.Status), "event", statusString(event.Status))
				continue
			}

			switch event.Status {
			case statusPending:
				reason, reject, err := deps.ShouldReject(ctx, chainID, req)
				if err != nil {
					return errors.Wrap(err, "should reject")
				} else if reject {
					return deps.Reject(ctx, chainID, req, reason)
				}

				return deps.Accept(ctx, chainID, req)
			case statusAccepted:
				return deps.Fulfill(ctx, chainID, req)
			case statusFulfilled:
				return deps.Claim(ctx, chainID, req)
			case statusRejected, statusReverted, statusClaimed:
			// Ignore for now
			default:
				return errors.New("unknown status [BUG]")
			}
		}

		return nil
	}
}
