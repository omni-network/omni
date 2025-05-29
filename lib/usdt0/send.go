package usdt0

import (
	"context"
	"math/big"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/layerzero"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/tokens"
	"github.com/omni-network/omni/lib/umath"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

// Send sends USDT0 from one chain to another.
func Send(
	ctx context.Context,
	backend *ethbackend.Backend,
	user common.Address,
	srcChainID uint64,
	destChainID uint64,
	amount *big.Int,
) (*ethclient.Receipt, error) {
	oftAddr, ok := oftByChain[srcChainID]
	if !ok {
		return nil, errors.New("no oft", "chain_id", srcChainID)
	}

	destEID, ok := layerzero.EIDByChain(destChainID)
	if !ok {
		return nil, errors.New("no eid", "chain_id", destChainID)
	}

	oft, err := NewIOFT(oftAddr, backend)
	if err != nil {
		return nil, errors.Wrap(err, "new oft")
	}

	srcToken, ok := tokenByChain[srcChainID]
	if !ok {
		return nil, errors.New("no token", "chain_id", srcChainID)
	}

	// Used for logging amount received
	dstToken, ok := tokenByChain[destChainID]
	if !ok {
		return nil, errors.New("no token", "chain_id", destChainID)
	}

	// Used to log native fee
	srcNativeTkn, ok := tokens.Native(srcChainID)
	if !ok {
		return nil, errors.New("no native token", "chain_id", srcChainID)
	}

	if err := maybeApprove(ctx, backend, user, oftAddr, srcToken, amount); err != nil {
		return nil, errors.Wrap(err, "approve l1 bridge")
	}

	params := SendParam{
		DstEid:   destEID,
		To:       toBz32(user),
		AmountLD: amount,
	}

	_, _, oftReceipt, err := oft.QuoteOFT(&bind.CallOpts{Context: ctx}, params)
	if err != nil {
		return nil, errors.Wrap(err, "quote oft")
	}

	// TODO: check amount received, require within acceptable range
	params.MinAmountLD = oftReceipt.AmountReceivedLD

	fee, err := oft.QuoteSend(&bind.CallOpts{Context: ctx}, params, false)
	if err != nil {
		return nil, errors.Wrap(err, "quote send")
	}

	txOpts, err := backend.BindOpts(ctx, user)
	if err != nil {
		return nil, errors.Wrap(err, "bind opts")
	}

	txOpts.Value = fee.NativeFee

	tx, err := oft.Send(txOpts, params, fee, user)
	if err != nil {
		return nil, errors.Wrap(err, "send")
	}

	receipt, err := backend.WaitMined(ctx, tx)
	if err != nil {
		return nil, errors.Wrap(err, "wait mined")
	}

	log.Info(ctx, "Sent USDT0",
		"src_chain", evmchain.Name(srcChainID),
		"dest_chain", evmchain.Name(destChainID),
		"sent", srcToken.FormatAmt(oftReceipt.AmountSentLD),
		"received", dstToken.FormatAmt(oftReceipt.AmountReceivedLD),
		"fee", srcNativeTkn.FormatAmt(fee.NativeFee),
		"tx", receipt.TxHash)

	return receipt, nil
}

// maybeApproveRouter approves `spender` to spend `amount` of `token` on behalf of `owner`.
func maybeApprove(
	ctx context.Context,
	backend *ethbackend.Backend,
	owner common.Address,
	spender common.Address,
	token tokens.Token,
	amount *big.Int,
) error {
	erc20, err := bindings.NewIERC20(token.Address, backend)
	if err != nil {
		return errors.Wrap(err, "new token")
	}

	allowance, err := erc20.Allowance(&bind.CallOpts{Context: ctx}, owner, spender)
	if err != nil {
		return errors.Wrap(err, "get allowance")
	}

	if bi.GTE(allowance, amount) {
		return nil
	}

	txOpts, err := backend.BindOpts(ctx, owner)
	if err != nil {
		return errors.Wrap(err, "bind opts")
	}

	tx, err := erc20.Approve(txOpts, spender, umath.MaxUint256)
	if err != nil {
		return errors.Wrap(err, "approve")
	}

	if _, err = backend.WaitMined(ctx, tx); err != nil {
		return errors.Wrap(err, "wait mined")
	}

	log.Info(ctx, "Approved token spend",
		"token", token.Symbol,
		"chain", evmchain.Name(token.ChainID),
		"owner", owner.Hex(),
		"spender", spender.Hex(),
		"tx", tx.Hash().Hex())

	return nil
}

// toBz32 converts an Ethereum address to a 32-byte address.
func toBz32(addr common.Address) [32]byte {
	var bz [32]byte
	copy(bz[12:], addr.Bytes())

	return bz
}
