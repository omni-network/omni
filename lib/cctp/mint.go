package cctp

import (
	"context"
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
	interval time.Duration
}

// WithInterval sets the cadence for the mint loop.
func WithInterval(interval time.Duration) MintForeverOption {
	return func(c *mintForeverOpts) {
		c.interval = interval
	}
}

func defaultMintOpts() *mintForeverOpts {
	return &mintForeverOpts{
		interval: 30 * time.Second,
	}
}

// MintForever mints submitted CCTP MsgSendUSDC messages forever. Messages are read from the db.
func MintForever(
	ctx context.Context,
	db *cctpdb.DB,
	minter common.Address,
	chain evmchain.Metadata,
	client Client,
	backend *ethbackend.Backend,
	opts ...MintForeverOption,
) {
	ctx = log.WithCtx(ctx,
		"process", "cctp.MintForever",
		"chain", chain.Name,
		"minter", minter)

	o := defaultMintOpts()
	for _, opt := range opts {
		opt(o)
	}

	ticker := time.NewTicker(o.interval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			err := tryMintSubmitted(ctx, db, minter, chain.ChainID, client, backend)
			if ctx.Err() != nil {
				return
			} else if err != nil {
				log.Error(ctx, "Mint failed (will retry)", err)

				continue
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
		return errors.Wrap(err, "list msgs")
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

	attestation, status, err := client.GetAttestation(ctx, msg.MessageHash)
	if err != nil {
		return errors.Wrap(err, "get attestation")
	}

	// Attestations pendings, skip
	if status == AttestationStatusPendingConfirmations {
		log.Debug(ctx, "Attestation confirmations pending")
		return nil
	}

	// Marks message as minted
	setMinted := func() error {
		if err := db.SetMsg(ctx, withStatus(msg, types.MsgStatusMinted)); err != nil {
			return errors.Wrap(err, "set minted")
		}

		return nil
	}

	received, err := didReceive(ctx, msgTransmitter, msg)
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

	log.Info(ctx, "Mint received", "tx_hash", receipt.TxHash)

	// Set minted
	if err := setMinted(); err != nil {
		return err
	}

	log.Info(ctx, "Message marked as minted", "tx_hash", receipt.TxHash)

	postMintBalance, err := tokenutil.BalanceOf(ctx, backend, usdc, msg.Recipient)
	if err != nil {
		return errors.Wrap(err, "post mint balance")
	}

	// If balance has not increased, warn
	// Do not mark failure, as this may be due to recipient spending USDC
	if preMintBalance.Cmp(postMintBalance) >= 0 {
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

		msgTransmitter, _, err := newMessageTransmitter(msg.DestChainID, client)
		if err != nil {
			return false, errors.Wrap(err, "message transmitter")
		}

		return didReceive(ctx, msgTransmitter, msg)
	}
}

// didReceive checks if a MsgSendUSDC has been received by dest MessageTransmitter.
// It checks MessageTransmitter.UsedNonces(...) to see message nonce has been used.
func didReceive(ctx context.Context, msgTransmitter *MessageTransmitter, msg types.MsgSendUSDC) (bool, error) {
	if len(msg.MessageBytes) < 84 {
		return false, errors.New("message bytes too short", "len", len(msg.MessageBytes))
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
	// Nonce key is keccak256(abi.encodePacked(nonce, sourceDomain))
	nonceKey := crypto.Keccak256Hash(append(
		msg.MessageBytes[12:20],  // nonce
		msg.MessageBytes[4:8]..., // source domain
	))

	used, err := msgTransmitter.UsedNonces(&bind.CallOpts{Context: ctx}, nonceKey)
	if err != nil {
		return false, errors.Wrap(err, "used nonce")
	}

	// 0 == unused
	if bi.IsZero(used) {
		return false, nil
	}

	return true, nil
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
