package nomina

import (
	"context"
	"math/big"

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

		balance, err := omni.BalanceOf(callOpts, address)
		if err != nil {
			return errors.Wrap(err, "balance of")
		}

		if balance == big.NewInt(0) {
			log.Info(ctx, "Skipping account with no OMNI balance",
				"network", network.ID,
				"address", address,
				"role", account.Role,
			)

			continue
		}

		txOpts, err := backend.BindOpts(ctx, address)
		if err != nil {
			return errors.Wrap(err, "bind opts")
		}

		tx, err := omni.Approve(txOpts, addrs.NomToken, balance)
		if err != nil {
			return errors.Wrap(err, "approve")
		}

		_, err = backend.WaitMined(ctx, tx)
		if err != nil {
			return errors.Wrap(err, "wait approve mined")
		}

		tx, err = nomina.Convert(txOpts, address, balance)
		if err != nil {
			return errors.Wrap(err, "convert")
		}

		_, err = backend.WaitMined(ctx, tx)
		if err != nil {
			return errors.Wrap(err, "wait convert mined")
		}

		log.Info(ctx, "Converted OMNI to NOM",
			"network", network.ID,
			"address", address,
			"role", account.Role,
			"omni", bi.ToEtherF64(balance),
			"nom", bi.ToEtherF64(bi.MulRaw(balance, evmredenom.Factor)),
		)
	}

	return nil
}
