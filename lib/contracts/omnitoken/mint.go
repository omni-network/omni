package omnitoken

import (
	"context"
	"math/big"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/contracts"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/common"
)

// Mint mints OMNI for user on ephemeral networks. The user adddress pays for the minting, and must be in given backend.
func Mint(ctx context.Context, network netconf.ID, backend *ethbackend.Backend, user common.Address, amount *big.Int) error {
	if !network.IsEphemeral() {
		return errors.New("can only mint on ephemeral networks")
	}

	addrs, err := contracts.GetAddresses(ctx, network)
	if err != nil {
		return errors.Wrap(err, "get addrs")
	}

	txOpts, err := backend.BindOpts(ctx, user)
	if err != nil {
		return errors.Wrap(err, "bind opts")
	}

	contract, err := bindings.NewMockERC20(addrs.Token, backend)
	if err != nil {
		return errors.Wrap(err, "bind contract")
	}

	tx, err := contract.Mint(txOpts, user, amount)
	if err != nil {
		return errors.Wrap(err, "mint tx")
	}

	_, err = backend.WaitMined(ctx, tx)
	if err != nil {
		return errors.Wrap(err, "wait mined")
	}

	return nil
}
