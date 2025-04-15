package cctp

import (
	"context"
	"math/big"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/cast"
	"github.com/omni-network/omni/lib/cctp/db"
	"github.com/omni-network/omni/lib/cctp/types"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/evmchain"
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
	db *db.DB,
	srcChainID, destChainID uint64,
	user common.Address,
	amount *big.Int,
) error {
	srcChain := evmchain.Name(srcChainID)
	dstChain := evmchain.Name(destChainID)

	usdc, ok := tokens.ByAsset(srcChainID, tokens.USDC)
	if !ok {
		return errors.New("no usdc", "src_chain", srcChain)
	}

	tknMessenger, tknMessengerAddr, err := newTokenMessenger(srcChainID, backend)
	if err != nil {
		return errors.Wrap(err, "new token messenger")
	}

	msgTransmitter, _, err := newMessageTransmitter(srcChainID, backend)
	if err != nil {
		return errors.Wrap(err, "new message transmitter")
	}

	if err := maybeApproveMessenger(ctx, backend, usdc, user, amount, tknMessengerAddr); err != nil {
		return errors.Wrap(err, "approve")
	}

	// CCTP uses bytes32 addresses
	recipient := cast.EthAddress32(user)

	txOpts, err := backend.BindOpts(ctx, user)
	if err != nil {
		return errors.Wrap(err, "bind opts")
	}

	tx, err := tknMessenger.DepositForBurn(txOpts, amount, umath.MustToUint32(destChainID), recipient, usdc.Address)
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

	msg := types.MsgSendUSDC{
		TxHash:       tx.Hash(),
		SrcChainID:   srcChainID,
		DestChainID:  destChainID,
		Amount:       amount,
		MessageBytes: messageBz,
		MessageHash:  messageHash,
		Recipient:    user,
		Status:       types.MsgStatusSubmitted,
	}

	attrs := []any{
		"tx", tx.Hash(),
		"src_chain", srcChain,
		"dest_chain", dstChain,
		"amount", usdc.FormatAmt(amount),
		"message_bytes", hexutil.Encode(messageBz),
		"message_hash", messageHash,
		"recipient", user,
	}

	log.Info(ctx, "Sent USDC", attrs...)

	if err := db.InsertMsg(ctx, msg); err != nil {
		return errors.Wrap(err, "insert message", "tx", tx.Hash())
	}

	return nil
}

// newTokenMessenger returns a new TokenMessenger instance for chainID.
func newTokenMessenger(chainID uint64, client ethclient.Client) (*TokenMessenger, common.Address, error) {
	addr, ok := tokenMessengers[chainID]
	if !ok {
		return nil, common.Address{}, errors.New("no messenger", "chain", evmchain.Name(chainID))
	}

	msgr, err := NewTokenMessenger(addr, client)
	if err != nil {
		return nil, common.Address{}, err
	}

	return msgr, addr, nil
}

// newMessageTransmitter returns a new MessageTransmitter instance for chainID.
func newMessageTransmitter(chainID uint64, client ethclient.Client) (*MessageTransmitter, common.Address, error) {
	addr, ok := messageTransmitters[chainID]
	if !ok {
		return nil, common.Address{}, errors.New("no transmitter", "chain", evmchain.Name(chainID))
	}

	transmitter, err := NewMessageTransmitter(addr, client)
	if err != nil {
		return nil, common.Address{}, err
	}

	return transmitter, addr, nil
}

// maybeApproveMessenger approves the TokenMessenger to spend USDC, if needed.
func maybeApproveMessenger(
	ctx context.Context,
	backend *ethbackend.Backend,
	usdc tokens.Token,
	user common.Address,
	amount *big.Int,
	messenger common.Address,
) error {
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

	log.Info(ctx, "Approved USDC spend", "chain", backend.Name(), "usdc", usdc.Address, "messenger", messenger, "tx", tx.Hash())

	return nil
}

// parseMessageSent finds and returns the message bytes from the MessageSent event in a transaction receipt.
func parseMessageSent(receipt *ethclient.Receipt, msgTransmitter *MessageTransmitter) ([]byte, error) {
	for _, log := range receipt.Logs {
		if len(log.Topics) > 0 && log.Topics[0] == messageSentEvent.ID {
			ev, err := msgTransmitter.ParseMessageSent(*log)
			if err != nil {
				return nil, errors.Wrap(err, "parse message sent")
			}

			return ev.Message, nil
		}
	}

	return nil, errors.New("event not found")
}
