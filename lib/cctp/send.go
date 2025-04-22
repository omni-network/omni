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
	"github.com/ethereum/go-ethereum/crypto"
)

type SendUSDCArgs struct {
	Sender      common.Address
	Recipient   common.Address
	SrcChainID  uint64
	DestChainID uint64
	Amount      *big.Int
}

// SendUSDC sends USDC from one chain to another using CCTP, storing the message in DB.
// It does not receive USDC on the destination chain.
func SendUSDC(
	ctx context.Context,
	db *db.DB,
	backend *ethbackend.Backend,
	args SendUSDCArgs,
) error {
	srcChain := evmchain.Name(args.SrcChainID)
	dstChain := evmchain.Name(args.DestChainID)

	usdc, ok := tokens.ByAsset(args.SrcChainID, tokens.USDC)
	if !ok {
		return errors.New("no usdc", "src_chain", srcChain)
	}

	c, err := newContracts(args.SrcChainID, backend)
	if err != nil {
		return errors.Wrap(err, "new contracts")
	}

	if err := maybeApproveMessenger(ctx, backend, usdc, args.Sender, args.Amount, c.TokenMessageAddress); err != nil {
		return errors.Wrap(err, "approve")
	}

	// CCTP uses bytes32 addresses
	recipient := cast.EthAddress32(args.Recipient)

	txOpts, err := backend.BindOpts(ctx, args.Sender)
	if err != nil {
		return errors.Wrap(err, "bind opts")
	}

	domainID, ok := domains[args.DestChainID]
	if !ok {
		return errors.New("unknown domain ID", "dest_chain", dstChain)
	}

	tx, err := c.TokenMessenger.DepositForBurn(txOpts, args.Amount, domainID, recipient, usdc.Address)
	if err != nil {
		return errors.Wrap(err, "deposit for burn")
	}

	receipt, err := backend.WaitMined(ctx, tx)
	if err != nil {
		return errors.Wrap(err, "wait mined")
	}

	messageBz, err := parseMessageSent(receipt, c.MessageTransmitter)
	if err != nil {
		return errors.Wrap(err, "parse message sent")
	}

	messageHash := crypto.Keccak256Hash(messageBz)

	msg := types.MsgSendUSDC{
		TxHash:       receipt.TxHash,
		BlockHeight:  receipt.BlockNumber.Uint64(),
		SrcChainID:   args.SrcChainID,
		DestChainID:  args.DestChainID,
		Amount:       args.Amount,
		MessageBytes: messageBz,
		MessageHash:  messageHash,
		Recipient:    args.Recipient,
		Status:       types.MsgStatusSubmitted,
	}

	log.Info(ctx, "Sent USDC",
		"tx", receipt.TxHash,
		"block_height", receipt.BlockNumber.Uint64(),
		"src_chain", srcChain,
		"dest_chain", dstChain,
		"amount", usdc.FormatAmt(args.Amount),
		"message_hash", messageHash,
		"recipient", args.Recipient)

	if err := db.InsertMsg(ctx, msg); err != nil {
		return errors.Wrap(err, "insert message", "tx", tx.Hash())
	}

	return nil
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
