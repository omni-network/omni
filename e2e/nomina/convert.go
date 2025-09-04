package nomina

import (
	"context"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/halo/evmredenom"
	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/contracts"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

func ConvertOmni(ctx context.Context, network netconf.Network, backends ethbackend.Backends) error {
	ethereum, ok := network.EthereumChain()
	if !ok {
		return errors.New("no ethereum chain")
	}

	backend, err := backends.Backend(ethereum.ID)
	if err != nil {
		return errors.Wrap(err, "get backend")
	}

	addrs, err := contracts.GetAddresses(ctx, network.ID)
	if err != nil {
		return errors.Wrap(err, "get addresses")
	}

	code, err := backend.CodeAt(ctx, addrs.NomToken, nil)
	if err != nil {
		return errors.Wrap(err, "code at", "address", addrs.NomToken)
	}

	if len(code) == 0 {
		return errors.New("nom token not deployed")
	}

	omni, err := bindings.NewOmni(addrs.Token, backend)
	if err != nil {
		return errors.Wrap(err, "new omni")
	}

	nomina, err := bindings.NewNomina(addrs.NomToken, backend)
	if err != nil {
		return errors.Wrap(err, "new nomina")
	}

	accounts := eoa.AllAccounts(network.ID)
	for _, account := range accounts {
		address := account.Address

		callOpts := &bind.CallOpts{
			Context: ctx,
			From:    address,
		}

		omniBalance, err := omni.BalanceOf(callOpts, address)
		if err != nil {
			return errors.Wrap(err, "balance of", "address", address, "role", account.Role)
		}

		if bi.IsZero(omniBalance) {
			log.Info(ctx, "Skipping account with no OMNI balance",
				"network", network.ID,
				"address", address,
				"role", account.Role,
			)

			continue
		}

		ethBalance, err := backend.BalanceAt(ctx, address, nil)
		if err != nil {
			return errors.Wrap(err, "balance at", "address", address, "role", account.Role)
		}

		if bi.IsZero(ethBalance) {
			log.Info(ctx, "Skipping account with no ETH balance",
				"network", network.ID,
				"address", address,
				"role", account.Role,
			)
		}

		txOpts, err := backend.BindOpts(ctx, address)
		if err != nil {
			return errors.Wrap(err, "bind opts approve", "address", address, "role", account.Role)
		}

		tx, err := omni.Approve(txOpts, addrs.NomToken, omniBalance)
		if err != nil {
			return errors.Wrap(err, "approve", "address", address, "role", account.Role)
		}

		_, err = backend.WaitMined(ctx, tx)
		if err != nil {
			return errors.Wrap(err, "wait approve mined", "address", address, "role", account.Role)
		}

		txOpts, err = backend.BindOpts(ctx, address)
		if err != nil {
			return errors.Wrap(err, "bind opts convert", "address", address, "role", account.Role)
		}

		tx, err = nomina.Convert(txOpts, address, omniBalance)
		if err != nil {
			return errors.Wrap(err, "convert", "address", address, "role", account.Role)
		}

		_, err = backend.WaitMined(ctx, tx)
		if err != nil {
			return errors.Wrap(err, "wait convert mined", "address", address, "role", account.Role)
		}

		log.Info(ctx, "Converted OMNI to NOM",
			"network", network.ID,
			"address", address,
			"role", account.Role,
			"omni", bi.ToEtherF64(omniBalance),
			"nom", bi.ToEtherF64(bi.MulRaw(omniBalance, evmredenom.Factor)),
		)
	}

	return nil
}
