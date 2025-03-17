package rlusd

import (
	"context"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"golang.org/x/sync/errgroup"
)

// Deploy deploys RLUSD.e tokens on all chains in the network.
func Deploy(ctx context.Context, network netconf.Network, backends ethbackend.Backends, bridge, lockbox common.Address) error {
	var eg errgroup.Group

	if _, ok := network.EthereumChain(); !ok {
		return errors.New("no ethereum chain")
	}

	if isEmpty(lockbox) {
		return errors.New("lockbox required")
	}

	for _, chain := range network.EVMChains() {
		eg.Go(func() error {
			backend, err := backends.Backend(chain.ID)
			if err != nil {
				return errors.Wrap(err, "get backend", "chain", chain.Name)
			}

			addr, receipt, err := deployXToken(ctx, network.ID, backend, bridge)
			if err != nil {
				return errors.Wrap(err, "deploy xtoken", "chain", chain.Name)
			}

			log.Info(ctx, "RLUSD.e xtoken deployed", "addr", addr.Hex(), "chain", chain.Name, "tx", maybeTxHash(receipt))

			// if not ethereum, no lockox roles to assign
			if !netconf.IsEthereumChain(network.ID, chain.ID) {
				return nil
			}

			err = assignLockboxRoles(ctx, network.ID, backend, lockbox, addr)
			if err != nil {
				return errors.Wrap(err, "assign minter role", "chain", chain.Name)
			}

			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		return errors.Wrap(err, "deploy xtokens")
	}

	return maybeDeployCanonical(ctx, network, backends)
}

// assignLockboxRoles assigns the minter role to the RLUSD lockbox.
func assignLockboxRoles(ctx context.Context, networkID netconf.ID, backend *ethbackend.Backend, lockbox, token common.Address) error {
	contract, err := bindings.NewStablecoinUpgradeable(token, backend)
	if err != nil {
		return errors.Wrap(err, "new contract")
	}

	admin := eoa.MustAddress(networkID, eoa.RoleManager)

	txOpts, err := backend.BindOpts(ctx, admin)
	if err != nil {
		return errors.Wrap(err, "bind opts")
	}

	assignRole := func(name string, role common.Hash) error {
		hasRole, err := contract.HasRole(&bind.CallOpts{Context: ctx}, role, lockbox)
		if err != nil {
			return errors.Wrap(err, "has role")
		}

		if hasRole {
			return nil
		}

		tx, err := contract.GrantRole(txOpts, role, lockbox)
		if err != nil {
			return errors.Wrap(err, "grant role")
		}

		receipt, err := backend.WaitMined(ctx, tx)
		if err != nil {
			return errors.Wrap(err, "wait mined")
		}

		log.Info(ctx, "Role assigned to RLUSD lockbox", "tx", maybeTxHash(receipt), "lockbox", lockbox.Hex(), "xtoken", token.Hex(), "role", name)

		return nil
	}

	// keccak256("MINTER")
	minterRole := common.HexToHash("0xf0887ba65ee2024ea881d91b74c2450ef19e1557f03bed3ea9f16b037cbe2dc9")

	// keccak256("CLAWBACKER")
	clawbackRole := common.HexToHash("0x715bacafb7a853b9b91b59ae724920a9eb0c006c5b318ac393fa1bc8974edd98")

	if err := assignRole("minter", minterRole); err != nil {
		return errors.Wrap(err, "assign minter role")
	}

	if err := assignRole("clawbacker", clawbackRole); err != nil {
		return errors.Wrap(err, "assign clawbacker role")
	}

	return nil
}

func maybeDeployCanonical(ctx context.Context, network netconf.Network, backends ethbackend.Backends) error {
	// if no canonical deployment for network, deploy mock
	_, ok := canonicals[network.ID]
	if ok {
		return nil
	}

	l1, ok := network.EthereumChain()
	if !ok {
		return errors.New("no ethereum chain")
	}

	backend, err := backends.Backend(l1.ID)
	if err != nil {
		return errors.Wrap(err, "get backend", "chain", l1.Name)
	}

	addr, receipt, err := deployCanonical(ctx, network.ID, backend)
	if err != nil {
		return errors.Wrap(err, "deploy mock wrapped", "chain", l1.Name)
	}

	log.Info(ctx, "Mock canonical RLUSD deployed", "addr", addr.Hex(), "chain", l1.Name, "tx", maybeTxHash(receipt))

	return nil
}

func maybeTxHash(receipt *ethclient.Receipt) string {
	if receipt != nil {
		return receipt.TxHash.Hex()
	}

	return "nil"
}
