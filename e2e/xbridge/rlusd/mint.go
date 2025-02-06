package rlusd

import (
	"context"
	"math/big"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/common"
)

func MintCanonical(ctx context.Context, network netconf.Network, backends ethbackend.Backends, to common.Address, amount *big.Int) error {
	if _, ok := canonicals[network.ID]; ok {
		return errors.New("can only mint mock canonical")
	}

	// admin is minter
	admin := eoa.MustAddress(network.ID, eoa.RoleManager)

	canon, err := XToken().Canonical(ctx, network.ID)
	if err != nil {
		return errors.Wrap(err, "canonical")
	}

	backend, err := backends.Backend(canon.ChainID)
	if err != nil {
		return errors.Wrap(err, "get backend")
	}

	contract, err := bindings.NewStablecoinUpgradeable(canon.Address, backend)
	if err != nil {
		return errors.Wrap(err, "new contract")
	}

	txOpts, err := backend.BindOpts(ctx, admin)
	if err != nil {
		return errors.Wrap(err, "bind opts")
	}

	tx, err := contract.Mint(txOpts, to, amount)
	if err != nil {
		return errors.Wrap(err, "mint")
	}

	_, err = backend.WaitMined(ctx, tx)
	if err != nil {
		return errors.Wrap(err, "wait mined")
	}

	return nil
}
