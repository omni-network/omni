package app

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
	return func(ctx context.Context, height uint64, elogs []types.Log) error {
		for _, elog := range elogs {
			event, ok := eventsByTopic[elog.Topics[0]]
			if !ok {
				return errors.New("unknown event [BUG]")
			}

			reqID, err := deps.ParseID(chainID, elog)
			if err != nil {
				return errors.Wrap(err, "parse id")
			}

			offset := reqIDOffset(reqID)
			statusOffset.WithLabelValues(deps.ChainName(chainID), statusString(event.Status)).Set(float64(offset))
			ctx := log.WithCtx(ctx, "status", statusString(event.Status), "req_id", offset)
			log.Debug(ctx, "Processing event")

			req, _, err := deps.GetRequest(ctx, chainID, reqID)
			if err != nil {
				return errors.Wrap(err, "current status")
			} else if event.Status != req.Status {
				// TODO(corver): Detect unexpected on-chain status.
				log.Info(ctx, "Ignoring mismatching old event", "actual", statusString(req.Status))
				continue
			}

			switch event.Status {
			case statusPending:
				if reason, reject, err := deps.ShouldReject(ctx, chainID, req); err != nil {
					return errors.Wrap(err, "should reject")
				} else if reject {
					if err := deps.Reject(ctx, chainID, req, reason); err != nil {
						return errors.Wrap(err, "reject request")
					}
				} else {
					if err := deps.Accept(ctx, chainID, req); err != nil {
						return errors.Wrap(err, "accept request")
					}
				}
			case statusAccepted:
				if err := deps.Fulfill(ctx, chainID, req); err != nil {
					return errors.Wrap(err, "fulfill request")
				}
			case statusFulfilled:
				if err := deps.Claim(ctx, chainID, req); err != nil {
					return errors.Wrap(err, "claim request")
				}
			case statusRejected, statusReverted, statusClaimed:
			// Ignore for now
			default:
				return errors.New("unknown status [BUG]")
			}
		}

		return deps.SetCursor(ctx, chainID, height)
	}
}
