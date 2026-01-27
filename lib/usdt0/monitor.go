package usdt0

import (
	"context"
	"math/big"
	"time"

	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/layerzero"
	"github.com/omni-network/omni/lib/log"
)

// MonitorSendsForever monitors USDT0 sends in db, logging and udpating their status.
func MonitorSendsForever(ctx context.Context, db *DB, client layerzero.Client, chainIDs []uint64) {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				if err := monitorSends(ctx, db, client, chainIDs); err != nil {
					log.Error(ctx, "Failed to monitor sends (will retry)", err)
				}
			}
		}
	}()
}

// monitorSends checks the status of all messages in the database and updates them accordingly.
func monitorSends(ctx context.Context, db *DB, client layerzero.Client, chainIDs []uint64) error {
	// Get all messages that are not in a final state
	msgs, err := db.GetMsgs(ctx, FilterMsgByStatus(nonFinalStatuses()...))
	if err != nil {
		return errors.Wrap(err, "get msgs")
	}

	gaugePending(chainIDs, msgs)
	guageOldest(ctx, msgs, db)

	log.Info(ctx, "Monitoring USDT0 sends", "num_msgs", len(msgs))

	for _, msg := range msgs {
		ctx := log.WithCtx(ctx,
			"tx", msg.TxHash.Hex(),
			"src_chain", msg.SrcChainID,
			"dest_chain", msg.DestChainID,
			"amount", msg.Amount.String())

		if err := monitorSend(ctx, db, client, msg); err != nil {
			log.Error(ctx, "Error monitoring send", err)
		}
	}

	return nil
}

func monitorSend(ctx context.Context, db *DB, client layerzero.Client, msg MsgSend) error {
	messages, err := client.GetMessagesByTx(ctx, msg.TxHash.Hex())
	if err != nil {
		return errors.Wrap(err, "get message status")
	}

	if len(messages) == 0 {
		return errors.New("no messages found for tx", "tx", msg.TxHash.Hex())
	}

	if len(messages) > 1 {
		return errors.New("multiple messages found for same tx [BUG]")
	}

	status := layerzero.MsgStatus(messages[0].Status.Name)

	log.Debug(ctx, "Fetched message status", "status", status.String(), "tx", msg.TxHash.Hex())

	if err := status.Verify(); err != nil {
		return errors.Wrap(err, "unexpected message status [BUG]", "status", status.String())
	}

	if status == msg.Status {
		return nil
	}

	// If DELIVERED, log and delete
	if status == layerzero.MsgStatusDelivered {
		log.Info(ctx, "Message delivered")

		if err := db.DeleteMsg(ctx, msg.TxHash); err != nil {
			return errors.Wrap(err, "delete delivered message")
		}

		return nil
	}

	// If FAILED, bug
	if status == layerzero.MsgStatusFailed {
		return errors.New("message failed [BUG]", "message", "message failed")
	}

	// IF PAYLOAD_STORED, bug
	if status == layerzero.MsgStatusPayloadStored {
		return errors.New("message payload stored [BUG]", "payload", msg.TxHash.Hex())
	}

	// All other status, log and update
	log.Info(ctx, "USDT0 message status changed", "prev", msg.Status.String(), "new", status.String())
	if err := db.SetMsgStatus(ctx, msg.TxHash, status); err != nil {
		return errors.Wrap(err, "set msg status")
	}

	return nil
}

func isPending(status layerzero.MsgStatus) bool {
	for _, s := range nonFinalStatuses() {
		if status == s {
			return true
		}
	}

	return false
}

func nonFinalStatuses() []layerzero.MsgStatus {
	return []layerzero.MsgStatus{
		layerzero.MsgStatusConfirming,
		layerzero.MsgStatusInFlight,
		layerzero.MsgStatusPayloadStored,
	}
}

// 'pending' includes any non-final statuses: INFLIGHT, CONFIRMING, PAYLOAD_STORED.
func gaugePending(chainIDs []uint64, msgs []MsgSend) {
	// reset all routes
	for _, src := range chainIDs {
		for _, dst := range chainIDs {
			usdt0Pending.WithLabelValues(evmchain.Name(src), evmchain.Name(dst)).Set(0)
			msgsPending.WithLabelValues(evmchain.Name(src), evmchain.Name(dst)).Set(0)
		}
	}

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
		// only consider pending
		if !isPending(msg.Status) {
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
		usdt0Pending.WithLabelValues(r.src, r.dst).Set(float64(v.amt.Uint64()))
		msgsPending.WithLabelValues(r.src, r.dst).Set(float64(v.msgs))
	}
}

// guageOldest sets the oldest msg metric.
func guageOldest(ctx context.Context, msgs []MsgSend, db *DB) {
	gauge := func() error {
		oldestByStatus := make(map[layerzero.MsgStatus]time.Time)

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

		for _, status := range nonFinalStatuses() {
			createdAt, ok := oldestByStatus[status]
			if !ok {
				oldestMsg.WithLabelValues(status.String()).Set(0)
				continue
			}

			oldestMsg.WithLabelValues(status.String()).Set(float64(time.Since(createdAt).Seconds()))
		}

		return nil
	}

	if err := gauge(); err != nil {
		log.Warn(ctx, "Failed gauging oldest messages", err)
	}
}
