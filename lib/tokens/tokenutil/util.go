package tokenutil

import (
	"context"
	"math/big"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/tokens"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

// BalanceOf retrieves the balance of a given address for a specific token.
func BalanceOf(
	ctx context.Context,
	client ethclient.Client,
	tkn tokens.Token,
	addr common.Address,
) (*big.Int, error) {
	switch {
	case tkn.IsNative():
		return client.BalanceAt(ctx, addr, nil)
	default:
		contract, err := bindings.NewIERC20(tkn.Address, client)
		if err != nil {
			return nil, err
		}

		return contract.BalanceOf(&bind.CallOpts{Context: ctx}, addr)
	}
}
