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

	getDepositForBurn := newDepositForBurnGetter(tknMessenger, tknMessengerAddr, recipient)
	getMessageSent := newMessageSentGetter(msgTransmitter, msgTransmitterAddr)

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

		err = xprov.StreamEventLogs(ctx, req, func(ctx context.Context, header *ethtypes.Header, elogs []ethtypes.Log) error {
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

			if err := insertMsgs(ctx, db, msgs); err != nil {
				return errors.Wrap(err, "insert msgs")
			}

			return db.SetCursor(ctx, chainVer.ID, header.Number.Uint64())
		})

		if ctx.Err() != nil {
			//nolint:nilerr // Allow context timeout.
			return nil
		}

		log.Warn(ctx, "Failure processing inbox events (will retry)", err)
		backoff()
	}
}

// insertMsgs inserts a list of MsgSendUSDC into the database, if they don't already exist.
func insertMsgs(ctx context.Context, db *db.DB, msgs []types.MsgSendUSDC) error {
	for _, msg := range msgs {
		curr, ok, err := db.GetMsg(ctx, msg.TxHash)
		if err != nil {
			return errors.Wrap(err, "has msg")
		}

		if ok {
			if msg.MessageHash == curr.MessageHash {
				// Already saved (expected)
				// TODO(kevin): check full equality, log if buggy
				continue
			}

			if err := db.SetMsg(ctx, curr); err != nil {
				// Message hash changed, update it
				return errors.Wrap(err, "set msg")
			}
		}

		if err := db.InsertMsg(ctx, curr); err != nil {
			// Message missed, insert it
			return errors.Wrap(err, "insert msg")
		}
	}

	return nil
}

// eventPairToMsg converts a (TokenMessenger.DepositForBurn, MessageTransmitter.MessageSent) pair to a MsgSendUSDC.
// It assumes the events are from the same transaction, and recipient is a valid address.
func eventPairToMsg(
	srcChainID uint64,
	burn *TokenMessengerDepositForBurn,
	send *MessageTransmitterMessageSent,
) types.MsgSendUSDC {
	messageBz := send.Message
	messageHash := crypto.Keccak256Hash(messageBz)

	return types.MsgSendUSDC{
		TxHash:       burn.Raw.TxHash,
		Recipient:    cast.MustEthAddress(burn.MintRecipient[:]),
		MessageBytes: messageBz,
		MessageHash:  messageHash,
		Amount:       burn.Amount,
		SrcChainID:   srcChainID,
		DestChainID:  uint64(burn.DestinationDomain),
	}
}

// newDepositForBurnGetter returns a func that finds an TokenMessenger.DepositForBurn event in a list of logs.
func newDepositForBurnGetter(
	contract *TokenMessenger,
	addr common.Address,
	recipient common.Address,
) func(logs []ethtypes.Log) (*TokenMessengerDepositForBurn, bool, error) {
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
func newMessageSentGetter(
	contract *MessageTransmitter,
	addr common.Address,
) func(logs []ethtypes.Log) (*MessageTransmitterMessageSent, bool, error) {
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
