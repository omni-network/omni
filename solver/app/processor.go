package app

import (
	"context"

	"github.com/omni-network/omni/lib/contracts/solvernet"
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
			event, ok := solvernet.EventByTopic(elog.Topics[0])
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

			statusOffset.WithLabelValues(deps.ProcessorName, event.Status.String()).Set(float64(order.Offset))

			ctx := log.WithCtx(ctx,
				"order_id", order.ID.String(),
				"offset", order.Offset,
				"status", order.Status,
			)

			if event.Status != order.Status {
				log.Debug(ctx, "Ignoring old order event (status already changed)", "event_status", event.Status.String())
				continue
			}

			age := deps.InstrumentAge(ctx, chainID, height, order)

			log.Debug(ctx, "Processing order event", age)

			alreadyFilled := func() bool {
				// ignore err. maybeReject will handle unsupported dest chain
				filled, _ := deps.DidFill(ctx, order)
				return filled
			}

			// maybeReject rejects orders if necessary, logging and counting them, returning true if rejected.
			maybeReject := func() (bool, error) {
				reason, reject, err := deps.ShouldReject(ctx, order)
				if err != nil {
					return false, errors.Wrap(err, "should reject")
				} else if !reject {
					return false, nil
				}

				log.InfoErr(ctx, "Rejecting order", err, "reason", reason.String())

				// reject, log and count, swallow err
				if err := deps.Reject(ctx, order, reason); err != nil {
					return false, errors.Wrap(err, "reject order")
				}

				rejectedOrders.WithLabelValues(
					deps.ChainName(order.SourceChainID),
					reason.String(),
				).Inc()

				return true, nil
			}

			switch event.Status {
			case solvernet.StatusPending:
				if alreadyFilled() {
					return nil
				}

				debugPendingData(ctx, deps, order, elog)

				if didReject, err := maybeReject(); err != nil {
					return err
				} else if didReject {
					continue
				}

				log.Info(ctx, "Filling order")
				if err := deps.Fill(ctx, order); err != nil {
					return errors.Wrap(err, "fill order")
				}
			case solvernet.StatusFilled:
				log.Info(ctx, "Claiming order")
				if err := deps.Claim(ctx, order); err != nil {
					return errors.Wrap(err, "claim order")
				}
			case solvernet.StatusRejected, solvernet.StatusClosed, solvernet.StatusClaimed:
				// Noop for now
			default:
				return errors.New("unknown status [BUG]")
			}

			processedEvents.WithLabelValues(deps.ProcessorName, event.Status.String()).Inc()
		}

		return deps.SetCursor(ctx, chainID, height)
	}
}

func debugPendingData(ctx context.Context, deps procDeps, order Order, elog types.Log) {
	pendingData, err := order.PendingData()
	if err != nil {
		log.Warn(ctx, "Order not pending [BUG]", err)
		return
	}

	fill, err := pendingData.ParsedFillOriginData()
	if err != nil {
		log.Warn(ctx, "Failed to parse fill origin data", err)
		return
	}

	// use last call target for logs
	lastCall := fill.Calls[len(fill.Calls)-1]

	log.Debug(ctx, "Pending order data",
		"calls", len(fill.Calls),
		"call_target", lastCall.Target.Hex(),
		"call_selector", hexutil.Encode(lastCall.Selector[:]),
		"call_params", hexutil.Encode(lastCall.Params),
		"call_value", lastCall.Value.String(),
		"dst_chain", deps.ChainName(pendingData.DestinationChainID),
		"full_order_id", order.ID.Hex(),
		"target", deps.TargetName(pendingData),
		"tx", elog.TxHash.Hex(),
	)
}
