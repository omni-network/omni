package mantle

import (
	"context"
	"math/big"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/tokens"
	"github.com/omni-network/omni/lib/umath"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

var (
	l1USDC = mustToken(evmchain.IDEthereum, tokens.USDC)
	l2USDC = mustToken(evmchain.IDMantle, tokens.USDC)
)

// DepositUSDC bridges USDC from L1 to Mantle L2 via the native bridge.
func DepositUSDC(
	ctx context.Context,
	backend *ethbackend.Backend,
	user common.Address,
	amount *big.Int,
) (*ethclient.Receipt, error) {
	return depositERC20(ctx, backend, user, l1USDC, l2USDC, amount)
}

// depositERC20 bridges an ERC20 token from L1 to Mantle L2 via the native bridge.
func depositERC20(
	ctx context.Context,
	backend *ethbackend.Backend,
	user common.Address,
	l1Token, l2Token tokens.Token,
	amount *big.Int,
) (*ethclient.Receipt, error) {
	chainID := l1Token.ChainID

	bridgeAddr, ok := l1Bridges[chainID]
	if !ok {
		return nil, errors.New("no L1 bridge", "chain_id", chainID)
	}

	bridge, err := NewL1Bridge(bridgeAddr, backend)
	if err != nil {
		return nil, errors.Wrap(err, "new l1 bridge")
	}

	if err := maybeApprove(ctx, backend, user, bridgeAddr, l1Token, amount); err != nil {
		return nil, errors.Wrap(err, "approve l1 bridge")
	}

	txOpts, err := backend.BindOpts(ctx, user)
	if err != nil {
		return nil, errors.Wrap(err, "bind opts")
	}

	const minGasLimit = 200_000
	extraData := []byte{}
	tx, err := bridge.DepositERC20(txOpts, l1Token.Address, l2Token.Address, amount, minGasLimit, extraData)
	if err != nil {
		return nil, errors.Wrap(err, "deposit erc20")
	}

	receipt, err := backend.WaitMined(ctx, tx)
	if err != nil {
		return nil, errors.Wrap(err, "wait mined")
	}

	log.Info(ctx, "Deposited to Mantle", "amount", l1Token.FormatAmt(amount), "tx", receipt.TxHash)

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

func mustToken(chainID uint64, asset tokens.Asset) tokens.Token {
	tkn, ok := tokens.ByAsset(chainID, asset)
	if !ok {
		panic("token not found")
	}

	return tkn
}
