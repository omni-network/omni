package app

import (
	"context"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
)

// newEventProcessor returns a callback provided to xchain.Provider::StreamEventLogs processing
// all inbox contract events and driving order lifecycle.
func newEventProcessor(deps procDeps, chainID uint64) xchain.EventLogsCallback {
	return func(ctx context.Context, height uint64, elogs []types.Log) error {
		for _, elog := range elogs {
			event, ok := eventsByTopic[elog.Topics[0]]
			if !ok {
				return errors.New("unknown event [BUG]")
			}

			orderID, err := deps.ParseID(chainID, elog)
			if err != nil {
				return errors.Wrap(err, "parse id")
			}

			order, found, err := deps.GetOrder(ctx, chainID, orderID)
			if err != nil {
				return errors.Wrap(err, "get order")
			} else if !found {
				return errors.New("order not found [BUG]")
			}

			target := deps.TargetName(order)
			statusOffset.WithLabelValues(deps.ChainName(chainID), target, statusString(event.Status)).Set(float64(orderID.Uint64()))

			attrs := []any{
				"order_id", order.ID.String(),
				"status", order.Status,
				"src_chain_id", order.SourceChainID,
				"dst_chain_id", order.DestinationChainID,
			}

			fill, err := order.ParsedFillOriginData()
			if err != nil {
				log.Warn(ctx, "Failed to parse fill origin data", err, attrs...)
				attrs = append(attrs, "calls", unknown)
			} else {
				// use last call target for logs
				lastCall := fill.Calls[len(fill.Calls)-1]

				attrs = append(attrs,
					"calls", len(fill.Calls),
					"call_target", lastCall.Target.Hex(),
					"call_selector", hexutil.Encode(lastCall.Selector[:]),
					"call_params", hexutil.Encode(lastCall.Params),
					"call_value", lastCall.Value.String(),
				)
			}

			ctx := log.WithCtx(ctx, attrs...)

			log.Debug(ctx, "Processing order event")

			if event.Status != order.Status {
				// TODO(corver): Detect unexpected on-chain status.
				log.Info(ctx, "Ignoring mismatching old event", "actual", statusString(order.Status))
				continue
			}

			// maybeReject rejects orders if necessary, logging and counting them.
			maybeReject := func() (ShouldRejectResult, error) {
				r, err := deps.ShouldReject(ctx, chainID, order)

				if err != nil {
					return r, errors.Wrap(err, "should reject")
				}

				if r.AlreadyFilled || !r.Reject() {
					return r, nil
				}

				log.InfoErr(ctx, "Rejecting order", err, "reason", r.Reason.String())

				// reject, log and count, swallow err
				if err := deps.Reject(ctx, chainID, order, r.Reason); err != nil {
					return r, errors.Wrap(err, "reject order")
				}

				rejectedOrders.WithLabelValues(
					deps.ChainName(order.SourceChainID),
					deps.ChainName(order.DestinationChainID),
					target,
					r.Reason.String(),
				).Inc()

				return r, nil
			}

			switch event.Status {
			case statusPending:
				if r, err := maybeReject(); err != nil {
					return err
				} else if r.Reject() || r.AlreadyFilled {
					continue
				}

				log.Info(ctx, "Accepting order")
				if err := deps.Accept(ctx, chainID, order); err != nil {
					return errors.Wrap(err, "accept order")
				}
			case statusAccepted:
				// check reject again, as liquidity (or other conditions) may have changed
				if r, err := maybeReject(); err != nil {
					return err
				} else if r.Reject() || r.AlreadyFilled {
					continue
				}

				log.Info(ctx, "Filling order")
				if err := deps.Fill(ctx, chainID, order); err != nil {
					return errors.Wrap(err, "fill order")
				}
			case statusFilled:
				log.Info(ctx, "Claiming order")
				if err := deps.Claim(ctx, chainID, order); err != nil {
					return errors.Wrap(err, "claim order")
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
