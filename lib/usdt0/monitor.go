package usdt0

import (
	"context"
	"time"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/layerzero"
	"github.com/omni-network/omni/lib/log"
)

// MonitorSendsForever monitors USDT0 sends in db, logging and udpating their status.
func MonitorSendsForever(ctx context.Context, db *DB, client layerzero.Client) {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				if err := monitorSends(ctx, db, client); err != nil {
					log.Error(ctx, "Failed to monitor sends (will retry)", err)
				}
			}
		}
	}()
}

// monitorSends checks the status of all messages in the database and updates them accordingly.
func monitorSends(ctx context.Context, db *DB, client layerzero.Client) error {
	// Get all messages that are not in a final state
	msgs, err := db.GetMsgs(ctx, FilterMsgByStatus(
		layerzero.MsgStatusUnknown,
		layerzero.MsgStatusInFlight,
		layerzero.MsgStatusConfirming,
		layerzero.MsgStatusPayloadStored,
	))
	if err != nil {
		return errors.Wrap(err, "get msgs")
	}

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
