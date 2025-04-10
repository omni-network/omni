package cctp

import (
	"context"
	"math/big"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/tokens"
	"github.com/omni-network/omni/lib/umath"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

// SendUSDC sends USDC from one chain to another using CCTP, storing the message in DB.
// It does not receive USDC on the destination chain.
func SendUSDC(
	ctx context.Context,
	backend *ethbackend.Backend,
	db DB,
	srcChainID, destChainID uint64,
	user common.Address,
	amount *big.Int,
) error {
	usdc, ok := tokens.ByAsset(srcChainID, tokens.USDC)
	if !ok {
		return errors.New("no usdc", "chain_id", srcChainID)
	}

	tknMessenger, err := newTokenMessenger(srcChainID, backend)
	if err != nil {
		return errors.Wrap(err, "new token messenger")
	}

	msgTransmitter, err := newMessageTransmitter(srcChainID, backend)
	if err != nil {
		return errors.Wrap(err, "new message transmitter")
	}

	if err := maybeApproveMessenger(ctx, backend, usdc, user, amount); err != nil {
		return errors.Wrap(err, "approve")
	}

	// CCTP uses bytes32 addresses
	var recipient [32]byte
	copy(recipient[12:], user.Bytes())

	txOpts, err := backend.BindOpts(ctx, user)
	if err != nil {
		return errors.Wrap(err, "bind opts")
	}

	//nolint:gosec // chain ids are small
	destChainID32 := uint32(destChainID)

	tx, err := tknMessenger.DepositForBurn(txOpts, amount, destChainID32, recipient, usdc.Address)
	if err != nil {
		return errors.Wrap(err, "deposit for burn")
	}

	receipt, err := backend.WaitMined(ctx, tx)
	if err != nil {
		return errors.Wrap(err, "wait mined")
	}

	messageBz, err := parseMessageSent(receipt, msgTransmitter)
	if err != nil {
		return errors.Wrap(err, "parse message sent")
	}

	messageHash := crypto.Keccak256Hash(messageBz)

	msg := MsgSendUSDC{
		TxHash:       tx.Hash(),
		SrcChainID:   srcChainID,
		DestChainID:  destChainID,
		Amount:       amount,
		MessageBytes: messageBz,
		MessageHash:  messageHash,
		Recipient:    user,
	}

	attrs := []any{
		"tx", tx.Hash(),
		"src_chain_id", srcChainID,
		"dest_chain_id", destChainID,
		"amount", amount,
		"message_bytes", hexutil.Encode(messageBz),
		"message_hash", messageHash,
		"recipient", user,
	}

	log.Info(ctx, "Sent USDC", attrs...)

	if err := db.Insert(ctx, msg); err != nil {
		log.Error(ctx, "Failed to insert message into DB", err, attrs...)
		return errors.Wrap(err, "insert message")
	}

	return nil
}

// newTokenMessenger returns a new TokenMessenger instance for chainID.
func newTokenMessenger(chainID uint64, backend *ethbackend.Backend) (*TokenMessenger, error) {
	addr, ok := tokenMessengers[chainID]
	if !ok {
		return nil, errors.New("not found", "chain_id", chainID)
	}

	return NewTokenMessenger(addr, backend)
}

// newMessageTransmitter returns a new MessageTransmitter instance for chainID.
func newMessageTransmitter(chainID uint64, backend *ethbackend.Backend) (*MessageTransmitter, error) {
	addr, ok := messageTransmitters[chainID]
	if !ok {
		return nil, errors.New("not found", "chain_id", chainID)
	}

	return NewMessageTransmitter(addr, backend)
}

// maybeApproveMessenger approves the TokenMessenger to spend USDC, if needed.
func maybeApproveMessenger(
	ctx context.Context,
	backend *ethbackend.Backend,
	usdc tokens.Token,
	user common.Address,
	amount *big.Int,
) error {
	messenger, ok := tokenMessengers[usdc.ChainID]
	if !ok {
		return errors.New("no token messenger", "chain_id", usdc.ChainID)
	}

	erc20, err := bindings.NewIERC20(usdc.Address, backend)
	if err != nil {
		return errors.Wrap(err, "new token")
	}

	allowance, err := erc20.Allowance(&bind.CallOpts{Context: ctx}, user, messenger)
	if err != nil {
		return errors.Wrap(err, "get allowance")
	}

	if bi.GTE(allowance, amount) {
		return nil
	}

	txOpts, err := backend.BindOpts(ctx, user)
	if err != nil {
		return errors.Wrap(err, "bind opts")
	}

	tx, err := erc20.Approve(txOpts, messenger, umath.MaxUint256)
	if err != nil {
		return errors.Wrap(err, "approve")
	}

	if _, err = backend.WaitMined(ctx, tx); err != nil {
		return errors.Wrap(err, "wait mined")
	}

	log.Info(ctx, "Approved USDC spend", "usdc", usdc.Address, "message", messenger, "tx", tx.Hash())

	return nil
}

// parseMessageSent finds and returns the message bytes from the MessageSent event in a transaction receipt.
func parseMessageSent(receipt *ethclient.Receipt, msgTransmitter *MessageTransmitter) ([]byte, error) {
	topic := crypto.Keccak256Hash([]byte("MessageSent(bytes)"))

	for _, log := range receipt.Logs {
		if len(log.Topics) > 0 && log.Topics[0] == topic {
			ev, err := msgTransmitter.ParseMessageSent(*log)
			if err != nil {
				return nil, errors.Wrap(err, "parse message sent")
			}

			return ev.Message, nil
		}
	}

	return nil, errors.New("event not found")
}
