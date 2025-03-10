package app

import (
	"context"
	"time"

	"github.com/omni-network/omni/lib/contracts/solvernet"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
)

const (
	statusPending    = solvernet.StatusPending
	statusFilled     = solvernet.StatusFilled
	statusRejected   = solvernet.StatusRejected
	statusClosed     = solvernet.StatusClosed
	statusClaimed    = solvernet.StatusClaimed
	statusDestFilled = "dest_filled" // Unofficial state after destination fill

	maxAgeCache = 10000 // Max orders to track in age cache
)

// newEventProcessor returns a callback provided to xchain.Provider::StreamEventLogs processing
// all inbox contract events and driving order lifecycle.
func newEventProcessor(deps procDeps, chainID uint64) xchain.EventLogsCallback {
	ageCache := newAgeCache()

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

			target := deps.TargetName(order)
			timestamp := deps.BlockTimestamp(chainID, elog.BlockNumber)
			age := ageCache.InstrumentAge(order.ID, target, order.Status.String(), timestamp)
			statusOffset.WithLabelValues(deps.ChainName(chainID), target, event.Status.String()).Set(float64(orderID.Uint64()))

			ctx := log.WithCtx(ctx,
				"order_id", order.ID.String(),
				"status", order.Status,
				"src_chain", deps.ChainName(order.SourceChainID),
				"dst_chain", deps.ChainName(order.DestinationChainID),
				"age", age,
				"target", target,
			)

			log.Debug(ctx, "Processing order event")

			if event.Status != order.Status {
				// TODO(corver): Detect unexpected on-chain status.
				log.Info(ctx, "Ignoring mismatching old event", "actual", order.Status.String())
				continue
			}

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
				}

				if !reject {
					return false, nil
				}

				log.InfoErr(ctx, "Rejecting order", err, "reason", reason.String())

				// reject, log and count, swallow err
				if err := deps.Reject(ctx, order, reason); err != nil {
					return false, errors.Wrap(err, "reject order")
				}

				rejectedOrders.WithLabelValues(
					deps.ChainName(order.SourceChainID),
					deps.ChainName(order.DestinationChainID),
					target,
					reason.String(),
				).Inc()

				return true, nil
			}

			switch event.Status {
			case statusPending:
				if alreadyFilled() {
					return nil
				}

				// Track all orders for now, since we reject explicitly.
				ageCache.Add(order.ID, timestamp)
				debugOriginData(ctx, order)

				if didReject, err := maybeReject(); err != nil {
					return err
				} else if didReject {
					continue
				}

				log.Info(ctx, "Filling order")
				if err := deps.Fill(ctx, order); err != nil {
					return errors.Wrap(err, "fill order")
				}
				ageCache.InstrumentAge(order.ID, target, statusDestFilled, time.Now())
			case statusFilled:
				log.Info(ctx, "Claiming order")
				if err := deps.Claim(ctx, order); err != nil {
					return errors.Wrap(err, "claim order")
				}
				ageCache.Remove(order.ID) // Delete from cache on final state
			case statusRejected, statusClosed, statusClaimed:
				// Noop for now
				ageCache.Remove(order.ID) // Delete from cache on final state
			default:
				return errors.New("unknown status [BUG]")
			}

			processedEvents.WithLabelValues(deps.ChainName(chainID), target, event.Status.String()).Inc()
		}

		if ageCache.MaybePurge() {
			log.Warn(ctx, "Purged overflowing age cache [BUG]", nil)
		}

		return deps.SetCursor(ctx, chainID, height)
	}
}

func newAgeCache() *ageCache {
	return &ageCache{
		createdAts: make(map[solvernet.OrderID]time.Time),
	}
}

// ageCache enables best-effort tracking of order ages.
// Since on-chain state doesn't contain "created_height", a cache is used.
type ageCache struct {
	createdAts map[solvernet.OrderID]time.Time
}

// InstrumentAge records the age of an order, returning the age.
func (a *ageCache) InstrumentAge(order OrderID, target, status string, timestamp time.Time) time.Duration {
	if timestamp.IsZero() {
		return 0 // Best effort ignore for now
	}
	t0, ok := a.createdAts[order]
	if !ok {
		return 0 // Best effort ignore for now
	}
	age := timestamp.Sub(t0)
	orderAge.WithLabelValues("", target, status).Observe(age.Seconds())

	return age
}

// Remove removes an order from the cache.
func (a *ageCache) Remove(order OrderID) {
	delete(a.createdAts, order)
}

// Add adds a new order to the cache.
func (a *ageCache) Add(order OrderID, timestamp time.Time) {
	a.createdAts[order] = timestamp
}

// MaybePurge returns true if the cache was purged.
// This is required to prevent memory leaks.
func (a *ageCache) MaybePurge() bool {
	if len(a.createdAts) < maxAgeCache {
		return false
	}

	a.createdAts = make(map[solvernet.OrderID]time.Time)

	return true
}

func debugOriginData(ctx context.Context, order Order) {
	fill, err := order.ParsedFillOriginData()
	if err != nil {
		log.Warn(ctx, "Failed to parse fill origin data", err)
		return
	}

	// use last call target for logs
	lastCall := fill.Calls[len(fill.Calls)-1]

	log.Debug(ctx, "Fill origin data",
		"calls", len(fill.Calls),
		"call_target", lastCall.Target.Hex(),
		"call_selector", hexutil.Encode(lastCall.Selector[:]),
		"call_params", hexutil.Encode(lastCall.Params),
		"call_value", lastCall.Value.String(),
	)
}
