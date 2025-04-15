package cctp

import (
	"context"

	"github.com/omni-network/omni/lib/cast"
	"github.com/omni-network/omni/lib/cctp/db"
	"github.com/omni-network/omni/lib/cctp/types"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/expbackoff"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

type getMessageSentFunc func(logs []ethtypes.Log) (*MessageTransmitterMessageSent, bool, error)
type getDepositForBurnFunc func(logs []ethtypes.Log) (*TokenMessengerDepositForBurn, bool, error)

// StoreMessagesForever streams CCTP SendUSDC messages, and saves them to the database.
func StoreMessagesForever(
	ctx context.Context,
	db *db.DB,
	chainVer xchain.ChainVersion,
	client ethclient.Client,
	xprov xchain.Provider,
	recipient common.Address,
) error {
	msgTransmitter, msgTransmitterAddr, err := newMessageTransmitter(chainVer.ID, client)
	if err != nil {
		return errors.Wrap(err, "message transmitter")
	}

	tknMessenger, tknMessengerAddr, err := newTokenMessenger(chainVer.ID, client)
	if err != nil {
		return errors.Wrap(err, "token messenger")
	}

	proc := newEventProc(db, chainVer,
		newDepositForBurnGetter(tknMessenger, tknMessengerAddr, recipient),
		newMessageSentGetter(msgTransmitter, msgTransmitterAddr),
	)

	backoff := expbackoff.New(ctx)
	for {
		from, ok, err := db.GetCursor(ctx, chainVer.ID)
		if !ok || err != nil {
			log.Warn(ctx, "Failed reading cursor (will retry)", err)
			backoff()

			continue
		}

		req := xchain.EventLogsReq{
			ChainID:         chainVer.ID,
			Height:          from,
			ConfLevel:       chainVer.ConfLevel,
			FilterAddresses: []common.Address{tknMessengerAddr, msgTransmitterAddr},
			FilterTopics:    []common.Hash{depositForBurnEvent.ID, messageSentEvent.ID},
		}

		err = xprov.StreamEventLogs(ctx, req, proc)

		if ctx.Err() != nil {
			//nolint:nilerr // Allow context timeout.
			return nil
		}

		log.Warn(ctx, "Failure processing inbox events (will retry)", err)
		backoff()
	}
}

// newEventProc returns an xchain.EventLogsCallback that processes CCTP DepositForBurn & MessageSent events.
func newEventProc(
	db *db.DB,
	chainVer xchain.ChainVersion,
	getDepositForBurn getDepositForBurnFunc,
	getMessageSent getMessageSentFunc,
) xchain.EventLogsCallback {
	return func(ctx context.Context, header *ethtypes.Header, elogs []ethtypes.Log) error {
		// Group logs by transaction hash
		byTxHash := make(map[common.Hash][]ethtypes.Log)
		for _, log := range elogs {
			byTxHash[log.TxHash] = append(byTxHash[log.TxHash], log)
		}

		var msgs []types.MsgSendUSDC
		for _, logs := range byTxHash {
			// Get DepositForBurn event
			burn, ok, err := getDepositForBurn(logs)
			if err != nil {
				return err
			} else if !ok {
				continue
			}

			// Get MessageSent event
			send, ok, err := getMessageSent(logs)
			if err != nil {
				return err
			} else if !ok {
				continue
			}

			msg := eventPairToMsg(chainVer.ID, burn, send)
			msgs = append(msgs, msg)
		}

		if err := upsertMsgs(ctx, db, msgs); err != nil {
			return errors.Wrap(err, "upsert msgs")
		}

		return db.SetCursor(ctx, chainVer.ID, header.Number.Uint64())
	}
}

// upsertMsgs upserts a list of MsgSendUSDC by tx hash, if necessary.
func upsertMsgs(ctx context.Context, db *db.DB, msgs []types.MsgSendUSDC) error {
	for _, msg := range msgs {
		curr, ok, err := db.GetMsg(ctx, msg.TxHash)
		if err != nil {
			return errors.Wrap(err, "has msg")
		}

		// Message missed, insert.
		if !ok {
			if err := db.InsertMsg(ctx, withStatus(msg, types.MsgStatusSubmitted)); err != nil {
				return errors.Wrap(err, "insert msg")
			}

			continue
		}

		sanityCheck := func() error {
			// Minted, but message hash changed
			if curr.Status == types.MsgStatusMinted && curr.MessageHash != msg.MessageHash {
				return errors.New("message hash changed post mint", "tx_hash", msg.TxHash, "old", curr.MessageHash, "new", msg.MessageHash)
			}

			// Same message hash, but different content
			if curr.MessageHash == msg.MessageHash && !curr.Equals(msg) {
				return errors.New("message same for different content", "tx_hash", msg.TxHash, "msg_hash", msg.MessageHash)
			}

			return nil
		}

		// Maybe warn.
		if err := sanityCheck(); err != nil {
			log.Warn(ctx, "Failed sanity check [BUG]", err)
		}

		// Already minted, skip.
		if curr.Status == types.MsgStatusMinted {
			continue
		}

		// Already saved, skip (expected).
		if curr.MessageHash == msg.MessageHash {
			continue
		}

		// Message hash changed, update.
		if err := db.SetMsg(ctx, withStatus(msg, types.MsgStatusSubmitted)); err != nil {
			return errors.Wrap(err, "set msg")
		}

		continue
	}

	return nil
}

// eventPairToMsg converts a (TokenMessenger.DepositForBurn, MessageTransmitter.MessageSent) pair to a MsgSendUSDC.
// It assumes the events are from the same transaction, and recipient is a valid ETH address.
func eventPairToMsg(
	srcChainID uint64,
	burn *TokenMessengerDepositForBurn,
	send *MessageTransmitterMessageSent,
) types.MsgSendUSDC {
	messageBz := send.Message
	messageHash := crypto.Keccak256Hash(messageBz)

	return types.MsgSendUSDC{
		TxHash:       burn.Raw.TxHash,
		Recipient:    cast.MustEthAddress(burn.MintRecipient[12:]),
		MessageBytes: messageBz,
		MessageHash:  messageHash,
		Amount:       burn.Amount,
		SrcChainID:   srcChainID,
		DestChainID:  uint64(burn.DestinationDomain),
		Status:       types.MsgStatusUnknown, // unknown, resolve in upsert
	}
}

// newDepositForBurnGetter returns a func that finds an TokenMessenger.DepositForBurn event in a list of logs.
func newDepositForBurnGetter(contract *TokenMessenger, addr common.Address, recipient common.Address) getDepositForBurnFunc {
	return func(logs []ethtypes.Log) (*TokenMessengerDepositForBurn, bool, error) {
		for _, log := range logs {
			// Skip if not TokenMessage.DepositForBurn event
			if log.Address != addr || log.Topics[0] != depositForBurnEvent.ID {
				continue
			}

			ev, err := contract.ParseDepositForBurn(log)
			if err != nil {
				return nil, false, errors.Wrap(err, "parse deposit for burn")
			}

			if ev.MintRecipient != cast.EthAddress32(recipient) {
				continue
			}

			return ev, true, nil
		}

		return nil, false, nil
	}
}

// newMessageSentGetter returns a func that finds an MessageTransmitter.MessageSent event in a list of logs.
func newMessageSentGetter(contract *MessageTransmitter, addr common.Address) getMessageSentFunc {
	return func(logs []ethtypes.Log) (*MessageTransmitterMessageSent, bool, error) {
		for _, log := range logs {
			// Skip if not MessageTransmitter.MessageSent event
			if log.Address != addr || log.Topics[0] != messageSentEvent.ID {
				continue
			}

			ev, err := contract.ParseMessageSent(log)
			if err != nil {
				return nil, false, errors.Wrap(err, "parse message sent")
			}

			return ev, true, nil
		}

		return nil, false, nil
	}
}

// withStatus sets the status of a MsgSendUSDC message.
func withStatus(msg types.MsgSendUSDC, status types.MsgStatus) types.MsgSendUSDC {
	msg.Status = status
	return msg
}
