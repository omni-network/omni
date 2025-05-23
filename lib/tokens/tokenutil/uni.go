package tokenutil

import (
	"context"
	"math/big"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/svmutil"
	"github.com/omni-network/omni/lib/tokens"
	"github.com/omni-network/omni/lib/uni"
	"github.com/omni-network/omni/lib/unibackend"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

// UniBalanceOf retrieves the balance of a given address for a specific token.
func UniBalanceOf(
	ctx context.Context,
	b unibackend.Backend,
	tkn tokens.Token,
	addr uni.Address,
) (*big.Int, error) {
	switch {
	case tkn.IsSVM() && tkn.IsNative():
		return svmutil.NativeBalanceAt(ctx, b.SVMClient(), addr.SVM())
	case tkn.IsSVM() && !tkn.IsNative():
		return svmutil.TokenBalanceAt(ctx, b.SVMClient(), tkn.SVMAddress, addr.SVM())
	case tkn.IsEVM() && tkn.IsNative():
		return b.EVMClient().BalanceAt(ctx, addr.EVM(), nil)
	case tkn.IsEVM() && !tkn.IsNative():
		contract, err := bindings.NewIERC20(tkn.Address, b.EVMClient())
		if err != nil {
			return nil, err
		}

		return contract.BalanceOf(&bind.CallOpts{Context: ctx}, addr.EVM())
	default:
		return nil, errors.New("impossible [BUG]")
	}
}
