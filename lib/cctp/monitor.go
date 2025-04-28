package cctp

import (
	"context"
	"math/big"
	"time"

	"github.com/omni-network/omni/lib/bi"
	cctpdb "github.com/omni-network/omni/lib/cctp/db"
	"github.com/omni-network/omni/lib/cctp/types"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/log"
)

func monitorForever(ctx context.Context, db *cctpdb.DB) {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			if err := monitorOnce(ctx, db); err != nil {
				log.Warn(ctx, "Monitor failed, will retry", err)
			}
		}
	}
}

func monitorOnce(ctx context.Context, db *cctpdb.DB) error {
	msgs, err := db.GetMsgs(ctx)
	if err != nil {
		return errors.Wrap(err, "get msgs")
	}

	gaugeInflight(msgs)

	if err := guageOldestMsg(ctx, msgs, db); err != nil {
		return errors.Wrap(err, "gauge oldest msg")
	}

	return nil
}

// gaugeInflight sets inflight metrics for the given messages.
func gaugeInflight(msgs []types.MsgSendUSDC) {
	type route struct {
		src string
		dst string
	}

	type inflight struct {
		msgs int
		amt  *big.Int
	}

	values := make(map[route]inflight)

	for _, msg := range msgs {
		// only consider non-minted messages
		if msg.Status != types.MsgStatusSubmitted {
			continue
		}

		src := evmchain.Name(msg.SrcChainID)
		dst := evmchain.Name(msg.DestChainID)

		r := route{src, dst}

		v, ok := values[r]
		if !ok {
			values[r] = inflight{
				msgs: 1,
				amt:  msg.Amount,
			}

			continue
		}

		v.msgs++
		v.amt = bi.Add(v.amt, msg.Amount)
		values[r] = v
	}

	for r, v := range values {
		usdcInFlight.WithLabelValues(r.src, r.dst).Set(float64(v.amt.Uint64()))
		msgsInFlight.WithLabelValues(r.src, r.dst).Set(float64(v.msgs))
	}
}

// guageOldestMsg sets the oldest msg metric.
func guageOldestMsg(ctx context.Context, msgs []types.MsgSendUSDC, db *cctpdb.DB) error {
	oldestByStatus := make(map[types.MsgStatus]time.Time)

	for _, msg := range msgs {
		createdAt, err := db.GetMsgCreatedAt(ctx, msg.TxHash)
		if err != nil {
			return errors.Wrap(err, "get msg created at", "tx_hash", msg.TxHash)
		}

		oldest, ok := oldestByStatus[msg.Status]
		if !ok || createdAt.Before(oldest) {
			oldestByStatus[msg.Status] = createdAt
		}
	}

	for _, status := range []types.MsgStatus{types.MsgStatusSubmitted, types.MsgStatusMinted} {
		createdAt, ok := oldestByStatus[status]
		if !ok {
			oldestMsg.WithLabelValues(status.String()).Set(0)
			continue
		}

		oldestMsg.WithLabelValues(status.String()).Set(float64(time.Since(createdAt).Seconds()))
	}

	return nil
}
