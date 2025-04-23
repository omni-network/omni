package cctp

import (
	"context"

	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/cast"
	"github.com/omni-network/omni/lib/cctp/db"
	"github.com/omni-network/omni/lib/cctp/types"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/expbackoff"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

type getMessageSentFunc func(logs []ethtypes.Log) (*MessageTransmitterMessageSent, bool, error)
type getDepositForBurnFunc func(logs []ethtypes.Log) (*TokenMessengerDepositForBurn, bool, error)
type isReceivedFunc func(ctx context.Context, msg types.MsgSendUSDC) (bool, error)

// AuditForever streams finalized CCTP SendUSDC messages to `recipient`, and reconiles them db state.
// It does this for all chains in `chains`.
// - messages missed are inserted
// - messages with incorrect fields are corrected
// - messages marked as `minted` are cofirmed minted, else re-marked as `submitted`.
// The audit process progresses db cursors.
func AuditForever(
	ctx context.Context,
	db *db.DB,
	networkID netconf.ID,
	xprov xchain.Provider,
	clients map[uint64]ethclient.Client,
	chains []evmchain.Metadata,
	recipient common.Address,
) error {
	ctx = log.WithCtx(ctx, "process", "cctp.AuditForever")

	if err := maybeInitCursors(ctx, db, chains, clients); err != nil {
		return errors.Wrap(err, "init cursors")
	}

	for _, chain := range chains {
		chainID := chain.ChainID

		client, ok := clients[chainID]
		if !ok {
			return errors.New("chain client not found", "chain", chain.Name)
		}

		contracts, err := newContracts(chainID, client)
		if err != nil {
			return errors.Wrap(err, "new contracts")
		}

		go auditChainForever(
			ctx,
			db,
			networkID,
			xprov,
			clients,
			contracts,
			chain,
			recipient)
	}

	return nil
}

// auditChainForever audits messages on a single chain forever.
func auditChainForever(
	ctx context.Context,
	db *db.DB,
	networkID netconf.ID,
	xprov xchain.Provider,
	clients map[uint64]ethclient.Client,
	contracts msgContracts,
	chain evmchain.Metadata,
	recipient common.Address,
) {
	chainID := chain.ChainID

	ctx = log.WithCtx(ctx, "chain", chain.Name, "recipient", recipient)

	proc := newEventProc(db, networkID, chainID,
		newIsReceived(clients),
		newDepositForBurnGetter(contracts.TokenMessenger, contracts.TokenMessengerAddress, recipient),
		newMessageSentGetter(contracts.MessageTransmitter, contracts.MessageTransmitterAddress),
	)

	log.Info(ctx, "Starting event processor")

	backoff := expbackoff.New(ctx)
	for {
		from, ok, err := db.GetCursor(ctx, chainID)
		if !ok || err != nil {
			log.Warn(ctx, "Failed reading cursor (will retry)", err)
			backoff()

			continue
		}

		req := xchain.EventLogsReq{
			ChainID:         chainID,
			Height:          from,
			ConfLevel:       xchain.ConfFinalized,
			FilterAddresses: []common.Address{contracts.TokenMessengerAddress, contracts.MessageTransmitterAddress},
			FilterTopics:    []common.Hash{depositForBurnEvent.ID, messageSentEvent.ID},
		}

		err = xprov.StreamEventLogs(ctx, req, proc)

		if ctx.Err() != nil {
			return
		}

		log.Warn(ctx, "Failure processing cctp events (will retry)", err)
		backoff()
	}
}

// newEventProc returns an xchain.EventLogsCallback that processes CCTP DepositForBurn & MessageSent events.
func newEventProc(
	db *db.DB,
	networkID netconf.ID,
	chainID uint64,
	isReceived isReceivedFunc,
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

			msg, err := eventPairToMsg(networkID, chainID, burn, send)
			if err != nil {
				return errors.Wrap(err, "event pair to msg")
			}

			msgs = append(msgs, msg)
		}

		if err := upsertMsgs(ctx, db, msgs, isReceived); err != nil {
			return errors.Wrap(err, "upsert msgs")
		}

		return db.SetCursor(ctx, chainID, header.Number.Uint64())
	}
}

// upsertMsgs upserts a list of MsgSendUSDC by tx hash, if necessary.
func upsertMsgs(ctx context.Context, db *db.DB, msgs []types.MsgSendUSDC, isReceived isReceivedFunc) error {
	var toInsert []types.MsgSendUSDC
	var toUpdate []types.MsgSendUSDC

	for _, streamed := range msgs {
		stored, ok, err := db.GetMsg(ctx, streamed.TxHash)
		if err != nil {
			return errors.Wrap(err, "has msg")
		}

		// Message missed, insert.
		if !ok {
			toInsert = append(toInsert, streamed)
			continue
		}

		// Maybe warn.
		if err := sanityCheck(stored, streamed); err != nil {
			log.Error(ctx, "Failed sanity check [BUG]", err)
		}

		status := stored.Status

		// Confirm mint.
		if status == types.MsgStatusMinted {
			minted, err := isReceived(ctx, stored)
			if err != nil {
				return errors.Wrap(err, "is received")
			}

			// Message not received, re-mark as submitted.
			if !minted {
				status = types.MsgStatusSubmitted
				log.Warn(ctx, "Message not received, re-marking as submitted", errors.New("message not received"), "tx_hash", streamed.TxHash)
			}
		}

		// Mint confirmed, but message hash changed -> BUG. Block processing.
		if status == types.MsgStatusMinted && stored.MessageHash != streamed.MessageHash {
			return errors.New("message hash changed post confirmed mint [BUG]", "tx_hash", streamed.TxHash, "stored", stored.MessageHash, "streamed", streamed.MessageHash)
		}

		correction := withStatus(streamed, status)

		// Update if correction required.
		if !stored.Equals(correction) {
			log.Debug(ctx, "Message changed", "tx_hash", streamed.TxHash, "diff", stored.Diff(correction))
			toUpdate = append(toUpdate, correction)
		}
	}

	// Insert
	for _, msg := range toInsert {
		if err := db.InsertMsg(ctx, msg); err != nil {
			return errors.Wrap(err, "insert msg")
		}

		log.Debug(ctx, "Inserted missing message", "tx_hash", msg.TxHash, "msg_hash", msg.MessageHash)
	}

	// Update
	for _, msg := range toUpdate {
		if err := db.SetMsg(ctx, msg); err != nil {
			return errors.Wrap(err, "set msg")
		}

		log.Debug(ctx, "Corrected message", "tx_hash", msg.TxHash, "msg_hash", msg.MessageHash)
	}

	return nil
}

// eventPairToMsg converts a (TokenMessenger.DepositForBurn, MessageTransmitter.MessageSent) pair to a MsgSendUSDC.
// It assumes the events are from the same transaction, and recipient is a valid ETH address.
func eventPairToMsg(
	networkID netconf.ID,
	srcChainID uint64,
	burn *TokenMessengerDepositForBurn,
	send *MessageTransmitterMessageSent,
) (types.MsgSendUSDC, error) {
	messageBz := send.Message
	messageHash := crypto.Keccak256Hash(messageBz)

	destChainID, ok := chainIDForDomain(networkID, burn.DestinationDomain)
	if !ok {
		return types.MsgSendUSDC{}, errors.New("chain ID not found for domain", "domain", burn.DestinationDomain)
	}

	return types.MsgSendUSDC{
		TxHash:       burn.Raw.TxHash,
		BlockHeight:  burn.Raw.BlockNumber,
		Recipient:    cast.MustEthAddress(burn.MintRecipient[12:]),
		MessageBytes: messageBz,
		MessageHash:  messageHash,
		Amount:       burn.Amount,
		SrcChainID:   srcChainID,
		DestChainID:  destChainID,
		Status:       types.MsgStatusSubmitted, // we know it's at least submitted, because we a processing a finalized event
	}, nil
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

// sanityCheck errors on unexpected stored vs. streamed msg state.
func sanityCheck(stored, streamed types.MsgSendUSDC) error {
	// Minted, but message hash changed
	if stored.Status == types.MsgStatusMinted && stored.MessageHash != streamed.MessageHash {
		return errors.New("message hash changed post marked mint", "tx_hash", stored.TxHash, "stored", stored.MessageHash, "streamed", stored.MessageHash)
	}

	// Same message hash, but different content
	if stored.MessageHash == streamed.MessageHash && !stored.Equals(withStatus(streamed, stored.Status)) {
		return errors.New("message same for different content", "tx_hash", stored.TxHash, "diff", stored.Diff(streamed))
	}

	// Source chain ID mismatch
	if stored.SrcChainID != streamed.SrcChainID {
		return errors.New("source chain ID mismatch", "tx_hash", stored.TxHash, "stored", stored.SrcChainID, "streamed", streamed.SrcChainID)
	}

	// Destination chain ID mismatch
	if stored.DestChainID != streamed.DestChainID {
		return errors.New("destination chain ID mismatch", "tx_hash", stored.TxHash, "stored", stored.DestChainID, "streamed", streamed.DestChainID)
	}

	// Amount mismatch
	if bi.NEQ(stored.Amount, streamed.Amount) {
		return errors.New("amount mismatch", "tx_hash", stored.TxHash, "stored", stored.Amount, "streamed", streamed.Amount)
	}

	// Recipient mismatch
	if stored.Recipient != streamed.Recipient {
		return errors.New("recipient mismatch", "tx_hash", stored.TxHash, "stored", stored.Recipient, "streamed", streamed.Recipient)
	}

	return nil
}

// maybeInitCursors initializes cursors for each chain to the latest block (if they don't exist).
func maybeInitCursors(ctx context.Context, db *db.DB, chains []evmchain.Metadata, clients map[uint64]ethclient.Client) error {
	for _, chain := range chains {
		chainID := chain.ChainID

		_, ok, err := db.GetCursor(ctx, chainID)
		if err != nil {
			return errors.Wrap(err, "get cursor", "chain", chain.Name)
		} else if ok {
			// Already initialized, skip.
			continue
		}

		client, ok := clients[chainID]
		if !ok {
			return errors.New("client not found", "chain", chain.Name)
		}

		blockNum, err := client.BlockNumber(ctx)
		if err != nil {
			return errors.Wrap(err, "get latest block", "chain", chain.Name)
		}

		// Set cursor to latest block
		if err := db.SetCursor(ctx, chainID, blockNum); err != nil {
			return errors.Wrap(err, "set cursor", "chain", chain.Name)
		}

		log.Info(ctx, "Initialized cursor", "chain", chain.Name, "block", blockNum)
	}

	return nil
}
