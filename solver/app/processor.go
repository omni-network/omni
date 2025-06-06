package app

import (
	"context"

	"github.com/omni-network/omni/lib/contracts/solvernet"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
)

// Event represents an order status change event emitted from on-chain tx.
type Event struct {
	OrderID OrderID
	Status  solvernet.OrderStatus
	Height  uint64
	Tx      string
}

// eventProcFunc abstracts the event processing function.
type eventProcFunc func(ctx context.Context, event Event) error

// newEventProcFunc returns a new event processing function for the provided chain.
// It handles all inbox contract events and driving order lifecycle.
func newEventProcFunc(deps procDeps, chainID uint64) eventProcFunc {
	return func(ctx context.Context, e Event) error {
		order, found, err := deps.GetOrder(ctx, chainID, e.OrderID)
		if err != nil {
			return errors.Wrap(err, "get order")
		} else if !found {
			return errors.New("order not found [BUG]")
		}

		statusOffset.WithLabelValues(deps.ProcessorName(chainID), e.Status.String()).Set(float64(order.Offset))

		ctx = log.WithCtx(ctx,
			"order_id", order.ID.String(),
			"offset", order.Offset,
			"status", order.Status,
		)

		if !e.Status.ValidTarget(order.Status) {
			// Invalid order transition can occur when RPCs return stale data, so just retry for now.
			return errors.New("invalid order transition", "event_status", e.Status)
		} else if e.Status != order.Status {
			log.Debug(ctx, "Ignoring old order event (status already changed)", "event_status", e.Status)
			return nil
		}

		age := deps.InstrumentAge(ctx, chainID, e.Height, order)

		log.Debug(ctx, "Processing order event", age)

		// maybeReject rejects orders if necessary, logging and counting them, returning true if rejected.
		maybeReject := func() (bool, error) {
			reason, reject, err := deps.ShouldReject(ctx, order)
			if !reject && err != nil {
				return false, errors.Wrap(err, "should reject")
			} else if !reject /* && err == nil */ {
				return false, nil
			} /* else reject && err != nil */

			log.InfoErr(ctx, "Rejecting order", err, "reason", reason)

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

		switch e.Status {
		case solvernet.StatusPending:
			if filled, err := deps.DidFill(ctx, order); err != nil {
				return errors.Wrap(err, "already filled")
			} else if filled {
				// TODO(corver): We don't wait for confirmation in this case, so this could still reorg out :(
				log.Info(ctx, "Skipping already filled order")
				break
			}

			deps.DebugPendingOrder(ctx, order, e)

			if didReject, err := maybeReject(); err != nil {
				return err
			} else if didReject {
				break
			}

			log.Debug(ctx, "Filling order")
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

		processedEvents.WithLabelValues(deps.ProcessorName(chainID), e.Status.String()).Inc()

		return nil
	}
}
