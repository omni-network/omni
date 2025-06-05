package tokenutil

import (
	"context"
	"math/big"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/tokens"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// Transfer transfers tokens from sender to recipient, native or ERC20.
func Transfer(
	ctx context.Context,
	backend *ethbackend.Backend,
	token tokens.Token,
	sender common.Address,
	recipient common.Address,
	amount *big.Int,
) (*ethclient.Receipt, error) {
	if token.IsNative() {
		return nativeTransfer(ctx, backend, sender, recipient, amount)
	}

	return erc20Transfer(ctx, backend, token, sender, recipient, amount)
}

func nativeTransfer(
	ctx context.Context,
	backend *ethbackend.Backend,
	sender common.Address,
	recipient common.Address,
	amount *big.Int,
) (*ethclient.Receipt, error) {
	tx := types.NewTransaction(
		0, // nonce handled by backend tx mngr
		recipient,
		amount,
		21000, // standard eth transfer gas
		nil,   // gas price handle by backend tx mngr
		nil,
	)

	txOpts, err := backend.BindOpts(ctx, sender)
	if err != nil {
		return nil, errors.Wrap(err, "bind opts")
	}

	tx, err = txOpts.Signer(sender, tx)
	if err != nil {
		return nil, errors.Wrap(err, "sign tx")
	}

	if err := backend.SendTransaction(ctx, tx); err != nil {
		return nil, errors.Wrap(err, "send tx")
	}

	return backend.WaitMined(ctx, tx)
}

func erc20Transfer(
	ctx context.Context,
	backend *ethbackend.Backend,
	token tokens.Token,
	sender common.Address,
	recipient common.Address,
	amount *big.Int,
) (*ethclient.Receipt, error) {
	txOpts, err := backend.BindOpts(ctx, sender)
	if err != nil {
		return nil, errors.Wrap(err, "bind opts")
	}

	contract, err := bindings.NewIERC20(token.Address, backend)
	if err != nil {
		return nil, errors.Wrap(err, "new erc20 contract")
	}

	tx, err := contract.Transfer(txOpts, recipient, amount)
	if err != nil {
		return nil, errors.Wrap(err, "transfer tokens")
	}

	return backend.WaitMined(ctx, tx)
}
