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

			req, _, err := deps.GetRequest(ctx, chainID, reqID)
			if err != nil {
				return errors.Wrap(err, "current status")
			}

			target := deps.TargetName(req)
			statusOffset.WithLabelValues(deps.ChainName(chainID), target, statusString(event.Status)).Set(float64(reqID.Uint64()))
			ctx := log.WithCtx(ctx, "target", target, "status", statusString(event.Status), "req_id", reqID)
			log.Debug(ctx, "Processing request event")

			if event.Status != req.Status {
				// TODO(corver): Detect unexpected on-chain status.
				log.Info(ctx, "Ignoring mismatching old event", "actual", statusString(req.Status))
				continue
			}

			switch event.Status {
			case statusPending:
				if reason, reject, err := deps.ShouldReject(ctx, chainID, req); err != nil {
					return errors.Wrap(err, "should reject")
				} else if reject {
					// ShouldReject does reject logging since it has more information.
					if err := deps.Reject(ctx, chainID, req, reason); err != nil {
						return errors.Wrap(err, "reject request")
					}
				} else {
					log.Info(ctx, "Accepting request")
					if err := deps.Accept(ctx, chainID, req); err != nil {
						return errors.Wrap(err, "accept request")
					}
				}
			case statusAccepted:
				log.Info(ctx, "Accepting request")
				if err := deps.Fulfill(ctx, chainID, req); err != nil {
					return errors.Wrap(err, "fulfill request")
				}
			case statusFulfilled:
				log.Info(ctx, "Claiming request")
				if err := deps.Claim(ctx, chainID, req); err != nil {
					return errors.Wrap(err, "claim request")
				}
			case statusRejected, statusReverted, statusClaimed:
			// Ignore for now
			default:
				return errors.New("unknown status [BUG]")
			}

			processedEvents.WithLabelValues(deps.ChainName(chainID), target, statusString(event.Status)).Inc()
		}

		return deps.SetCursor(ctx, chainID, height)
	}
}
