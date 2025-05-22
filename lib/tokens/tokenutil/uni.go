package tokenutil

import (
	"context"
	"math/big"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/solutil"
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
	case tkn.IsSol() && tkn.IsNative():
		return solutil.NativeBalanceAt(ctx, b.SolClient(), addr.Sol())
	case tkn.IsSol() && !tkn.IsNative():
		return solutil.TokenBalanceAt(ctx, b.SolClient(), tkn.SolAddress, addr.Sol())
	case tkn.IsEth() && tkn.IsNative():
		return b.EthClient().BalanceAt(ctx, addr.Eth(), nil)
	case tkn.IsEth() && !tkn.IsNative():
		contract, err := bindings.NewIERC20(tkn.Address, b.EthClient())
		if err != nil {
			return nil, err
		}

		return contract.BalanceOf(&bind.CallOpts{Context: ctx}, addr.Eth())
	default:
		return nil, errors.New("impossible [BUG]")
	}
}
