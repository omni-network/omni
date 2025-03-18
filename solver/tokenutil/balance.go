// Package tokenutil provides some useful token related logic.
// TODO(corver): extract app.Tokens to a separate package.
package tokenutil

import (
	"context"
	"math/big"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/solver/app"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

// Balance returns the balance of the given token and address.
func Balance(ctx context.Context, backend *ethbackend.Backend, tkn app.Token, address common.Address) (*big.Int, error) {
	if tkn.IsNative() {
		return nativeBalance(ctx, backend, address)
	}

	return erc20Balance(ctx, backend, tkn, address)
}

func nativeBalance(ctx context.Context, backend *ethbackend.Backend, address common.Address) (*big.Int, error) {
	balance, err := backend.BalanceAt(ctx, address, nil)
	if err != nil {
		return nil, errors.Wrap(err, "balanceAt")
	}

	return balance, nil
}

func erc20Balance(ctx context.Context, backend *ethbackend.Backend, token app.Token, address common.Address) (*big.Int, error) {
	contract, err := bindings.NewIERC20(token.Address, backend)
	if err != nil {
		return nil, errors.Wrap(err, "new IERC20")
	}

	balance, err := contract.BalanceOf(&bind.CallOpts{Context: ctx}, address)
	if err != nil {
		return nil, errors.Wrap(err, "balanceOf")
	}

	return balance, nil
}
