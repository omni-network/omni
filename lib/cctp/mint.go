package cctp

import (
	"context"
	"math/big"
	"time"

	"github.com/omni-network/omni/lib/bi"
	cctpdb "github.com/omni-network/omni/lib/cctp/db"
	"github.com/omni-network/omni/lib/cctp/types"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/tokens"
	"github.com/omni-network/omni/lib/tokens/tokenutil"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// MintForeverOption is a functional option for configuring MintForever.
type MintForeverOption func(*mintForeverOpts)

type mintForeverOpts struct {
	mintInterval  time.Duration
	purgeInterval time.Duration
}

// WithMintInterval sets the cadence for the mint loop.
func WithMintInterval(interval time.Duration) MintForeverOption {
	return func(c *mintForeverOpts) {
		c.mintInterval = interval
	}
}

// WithPurgeInterval sets the cadence for the purge loop.
func WithPurgeInterval(interval time.Duration) MintForeverOption {
	return func(c *mintForeverOpts) {
		c.purgeInterval = interval
	}
}

func defaultMintOpts() *mintForeverOpts {
	return &mintForeverOpts{
		mintInterval:  30 * time.Second,
		purgeInterval: 20 * time.Minute,
	}
}

// MintForever mints submitted CCTP MsgSendUSDC messages (from the db) on all chains.
func MintForever(
	ctx context.Context,
	db *cctpdb.DB,
	client Client,
	backends ethbackend.Backends,
	chains []evmchain.Metadata,
	minter common.Address,
	opts ...MintForeverOption,
) error {
	o := defaultMintOpts()
	for _, opt := range opts {
		opt(o)
	}

	for _, chain := range chains {
		backend, err := backends.Backend(chain.ChainID)
		if err != nil {
			return errors.Wrap(err, "get backend")
		}

		go mintChainForever(ctx, db, client, backend, chain, minter, o.mintInterval)
		go purgeChainForever(ctx, db, backend, chain, minter, o.purgeInterval)
	}

	return nil
}

// mintChainForever mints submitted CCTP MsgSendUSDC on a chain forever.
func mintChainForever(
	ctx context.Context,
	db *cctpdb.DB,
	client Client,
	backend *ethbackend.Backend,
	chain evmchain.Metadata,
	minter common.Address,
	interval time.Duration,
) {
	ctx = log.WithCtx(ctx,
		"subprocess", "cctp.MintForever",
		"chain", chain.Name,
		"minter", minter)

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	log.Info(ctx, "Starting mint loop", "interval", interval)

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			err := tryMintSubmitted(ctx, db, minter, chain.ChainID, client, backend)
			if ctx.Err() != nil {
				return
			} else if err != nil {
				log.Warn(ctx, "Mint failed (will retry)", err)
			}
		}
	}
}

func tryMintSubmitted(
	ctx context.Context,
	db *cctpdb.DB,
	minter common.Address,
	chainID uint64,
	client Client,
	backend *ethbackend.Backend,
) error {
	// Get submitted (not minted) messages for this chain
	msgs, err := db.GetMsgsBy(ctx, cctpdb.MsgFilter{
		Status:      types.MsgStatusSubmitted,
		DestChainID: chainID,
	})
	if err != nil {
		return errors.Wrap(err, "get submitted msgs")
	}

	msgTransmitter, _, err := newMessageTransmitter(chainID, backend)
	if err != nil {
		return errors.Wrap(err, "new message transmitter")
	}

	usdc, ok := tokens.ByAsset(chainID, tokens.USDC)
	if !ok {
		return errors.New("no usdc")
	}

	for _, msg := range msgs {
		if err := tryMint(ctx, db, usdc, minter, client, backend, msgTransmitter, msg); err != nil {
			return errors.Wrap(err, "try mint once")
		}
	}

	return nil
}

func tryMint(
	ctx context.Context,
	db *cctpdb.DB,
	usdc tokens.Token,
	minter common.Address,
	client Client,
	backend *ethbackend.Backend,
	msgTransmitter *MessageTransmitter,
	msg types.MsgSendUSDC,
) error {
	ctx = log.WithCtx(ctx,
		"msg_hash", msg.MessageHash,
		"msg_tx_hash", msg.TxHash,
		"amount", msg.Amount,
		"recipient", msg.Recipient,
	)

	if crypto.Keccak256Hash(msg.MessageBytes) != msg.MessageHash {
		return errors.New("invalid message hash", "msg_hash", msg.MessageHash)
	}

	attestation, status, err := client.GetAttestation(ctx, msg.MessageHash)
	if err != nil {
		return errors.Wrap(err, "get attestation")
	}

	// Attestations pendings, skip
	if status == AttestationStatusPendingConfirmations {
		return nil
	}

	// Marks message as minted
	setMinted := func() error {
		if err := db.SetMsg(ctx, withStatus(msg, types.MsgStatusMinted)); err != nil {
			return errors.Wrap(err, "set minted")
		}

		return nil
	}

	received, err := DidReceive(ctx, backend, msg, nil)
	if err != nil {
		return errors.Wrap(err, "has been received")
	}

	// Already received, mark as minted (this means setMinted failed previously)
	if received {
		if err := setMinted(); err != nil {
			return err
		}

		log.Debug(ctx, "Message already received, marked as minted")

		return nil
	}

	preMintBalance, err := tokenutil.BalanceOf(ctx, backend, usdc, msg.Recipient)
	if err != nil {
		return errors.Wrap(err, "pre mint balance")
	}

	// Receive mint message
	receipt, err := receiveMint(ctx, minter, backend, msgTransmitter, msg, attestation)
	if err != nil {
		return errors.Wrap(err, "mint")
	}

	// Set minted
	if err := setMinted(); err != nil {
		log.Error(ctx, "Failed to set minted", err, "tx_hash", receipt.TxHash)
		return err
	}

	log.Info(ctx, "Mint received", "tx_hash", receipt.TxHash)

	postMintBalance, err := tokenutil.BalanceOf(ctx, backend, usdc, msg.Recipient)
	if err != nil {
		return errors.Wrap(err, "post mint balance")
	}

	// If balance has not increased, warn
	// Do not mark failure, as this may be due to recipient spending USDC
	if bi.GTE(preMintBalance, postMintBalance) {
		log.Warn(ctx, "USDC balance did not increase after mint",
			errors.New("balance did not increase"),
			"pre_mint_balance", preMintBalance,
			"post_mint_balance", postMintBalance)
	}

	return nil
}

// isReceived checks returns an isReceivedFunc for given chains / clients.
func newIsReceived(clients map[uint64]ethclient.Client) isReceivedFunc {
	return func(ctx context.Context, msg types.MsgSendUSDC) (bool, error) {
		client, ok := clients[msg.DestChainID]
		if !ok {
			return false, errors.New("no client for dest chain", "chain_id", msg.DestChainID)
		}

		return DidReceive(ctx, client, msg, nil)
	}
}

// DidReceive checks if a MsgSendUSDC has been received by dest MessageTransmitter.
// It checks MessageTransmitter.UsedNonces(...) to see message nonce has been used.
func DidReceive(ctx context.Context, ethClient ethclient.Client, msg types.MsgSendUSDC, blockNum *big.Int) (bool, error) {
	if len(msg.MessageBytes) < 84 {
		return false, errors.New("message bytes too short", "len", len(msg.MessageBytes))
	}

	msgTransmitter, _, err := newMessageTransmitter(msg.DestChainID, ethClient)
	if err != nil {
		return false, errors.Wrap(err, "message transmitter")
	}

	// Message format:
	//  Field                 Bytes      Type       Index
	//  version               4          uint32     0
	//  sourceDomain          4          uint32     4
	//  destinationDomain     4          uint32     8
	//  nonce                 8          uint64     12
	//  sender                32         bytes32    20
	//  recipient             32         bytes32    52
	//  messageBody           dynamic    bytes      84
	//
	// Nonce key is keccak256(abi.encodePacked(sourceDomain, nonce))
	var nonceBz []byte
	nonceBz = append(nonceBz, msg.MessageBytes[4:8]...)   // source domain
	nonceBz = append(nonceBz, msg.MessageBytes[12:20]...) // nonce
	nonceKey := crypto.Keccak256Hash(nonceBz)

	used, err := msgTransmitter.UsedNonces(&bind.CallOpts{
		Context:     ctx,
		BlockNumber: blockNum,
	}, nonceKey)
	if err != nil {
		return false, errors.Wrap(err, "used nonce")
	}

	// 0 == unused
	if bi.IsZero(used) {
		return false, nil
	}

	return true, nil
}

// isConfirmed checks if the MsgSendUSDC has been audited, received and finalized on the destination chain.
// After a confirmed messsage has been purged, isConfirmed will error.
func isConfirmed(ctx context.Context, db *cctpdb.DB, destClient ethclient.Client, msg types.MsgSendUSDC) (bool, error) {
	// Confirm message is tracked in DB
	stored, ok, err := db.GetMsg(ctx, msg.TxHash)
	if err != nil {
		return false, errors.Wrap(err, "get msg")
	} else if !ok {
		return false, errors.New("msg not found", "msg_hash", msg.MessageHash)
	} else if !stored.Equals(withStatus(msg, stored.Status)) {
		return false, errors.New("msg conflict", "msg_hash", msg.MessageHash, "diff", stored.Diff(msg))
	}

	// Confirm audit cursor is past message block height
	cursor, ok, err := db.GetCursor(ctx, msg.SrcChainID)
	if err != nil {
		return false, errors.Wrap(err, "get cursor", "chain_id", msg.SrcChainID)
	} else if !ok {
		return false, errors.New("cursor not found", "chain_id", msg.SrcChainID)
	} else if cursor < msg.BlockHeight {
		// Message tracked but not yet audited, so not confirmed
		return false, nil
	}

	// Confirm dest chain message is received and finalized
	header, err := destClient.HeaderByType(ctx, ethclient.HeadFinalized)
	if err != nil {
		return false, errors.Wrap(err, "get finalized header")
	}

	confirmed, err := DidReceive(ctx, destClient, msg, header.Number)
	if err != nil {
		return false, errors.Wrap(err, "did receive")
	}

	return confirmed, nil
}

// receiveMint submits the MsgSendUSDC and corresponding attestation to MessageTransmitter.receiveMessage.
func receiveMint(
	ctx context.Context,
	minter common.Address,
	backend *ethbackend.Backend,
	msgTransmitter *MessageTransmitter,
	msg types.MsgSendUSDC,
	attestation []byte,
) (*ethclient.Receipt, error) {
	txOpts, err := backend.BindOpts(ctx, minter)
	if err != nil {
		return nil, errors.Wrap(err, "bind opts")
	}

	tx, err := msgTransmitter.ReceiveMessage(txOpts, msg.MessageBytes, attestation)
	if err != nil {
		return nil, errors.Wrap(err, "receive message tx")
	}

	receipt, err := backend.WaitMined(ctx, tx)
	if err != nil {
		return nil, errors.Wrap(err, "wait mined")
	}

	return receipt, nil
}

// purgeChainForever purges all messages confirmed messages on a chain forever.
func purgeChainForever(
	ctx context.Context,
	db *cctpdb.DB,
	backend *ethbackend.Backend,
	chain evmchain.Metadata,
	minter common.Address,
	interval time.Duration,
) {
	ctx = log.WithCtx(ctx,
		"subprocess", "cctp.PurgeForever",
		"chain", chain.Name,
		"minter", minter)

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	log.Info(ctx, "Starting purge loop", "interval", interval)

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			err := tryPurge(ctx, db, backend, chain.ChainID)
			if ctx.Err() != nil {
				return
			} else if err != nil {
				log.Warn(ctx, "Purge failed (will retry)", err)
			}
		}
	}
}

func tryPurge(
	ctx context.Context,
	db *cctpdb.DB,
	backend *ethbackend.Backend,
	chainID uint64,
) error {
	ctx = log.WithCtx(ctx, "chain_id", chainID)

	msgs, err := db.GetMsgsBy(ctx, cctpdb.MsgFilter{
		Status:      types.MsgStatusMinted,
		DestChainID: chainID,
	})
	if err != nil {
		return errors.Wrap(err, "get minted msgs")
	}

	var toDelete []types.MsgSendUSDC

	// Mark confirmed messages for deletion
	for _, msg := range msgs {
		confirmed, err := isConfirmed(ctx, db, backend, msg)
		if err != nil {
			return errors.Wrap(err, "is confirmed")
		}

		if !confirmed {
			continue
		}

		toDelete = append(toDelete, msg)
	}

	// Delete confirmed messages
	for _, msg := range toDelete {
		if err := db.DeleteMsg(ctx, msg.TxHash); err != nil {
			return errors.Wrap(err, "delete msg")
		}

		log.Info(ctx, "Purged confirmed message", "msg_hash", msg.MessageHash, "tx_hash", msg.TxHash)
	}

	return nil
}
