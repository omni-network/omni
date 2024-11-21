package devapp

import (
	"context"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/lib/contracts"
	"github.com/omni-network/omni/lib/create3"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/common"
)

const (
	l1VaultSalt = "l1-vault"
	l1TokenSalt = "l1-token"
	l2TokenSalt = "l2-token"
)

var (
	create3Factory = contracts.Create3Factory(netconf.Devnet)
	deployer       = eoa.MustAddress(netconf.Devnet, eoa.RoleDeployer)
)

// Deploy deploys the mock tokens and vaults to devnet.
func Deploy(ctx context.Context, network netconf.Network, backends ethbackend.Backends) error {
	if network.ID != netconf.Devnet {
		return errors.New("onl devnet")
	}

	mockl1, ok := network.Chain(evmchain.IDMockL1)
	if !ok {
		return errors.New("no mock l1")
	}

	mockl2, ok := network.Chain(evmchain.IDMockL2)
	if !ok {
		return errors.New("no mock l2")
	}

	mockl1Backend, err := backends.Backend(mockl1.ID)
	if err != nil {
		return errors.Wrap(err, "backend mock l1")
	}

	mockl2Backend, err := backends.Backend(mockl2.ID)
	if err != nil {
		return errors.Wrap(err, "backend mock l2")
	}

	if err := deployToken(ctx, mockl1Backend, l1TokenSalt); err != nil {
		return errors.Wrap(err, "deploy l1 token")
	}

	if err := deployVault(ctx, mockl1Backend, l1VaultSalt, static.L1Token); err != nil {
		return errors.Wrap(err, "deploy vault")
	}

	if err := deployToken(ctx, mockl2Backend, l2TokenSalt); err != nil {
		return errors.Wrap(err, "deploy l2 token")
	}

	return nil
}

func deployVault(ctx context.Context, backend *ethbackend.Backend, salt string, collaterlTkn common.Address) error {
	txOpts, err := backend.BindOpts(ctx, deployer)
	if err != nil {
		return errors.Wrap(err, "bind opts")
	}

	factory, err := bindings.NewCreate3(create3Factory, backend)
	if err != nil {
		return errors.Wrap(err, "new create3")
	}

	initCode, err := contracts.PackInitCode(vaultABI, bindings.MockVaultBin, collaterlTkn)
	if err != nil {
		return errors.Wrap(err, "pack init code")
	}

	tx, err := factory.DeployWithRetry(txOpts, create3.HashSalt(salt), initCode) //nolint:contextcheck // Context is txOpts
	if err != nil {
		return errors.Wrap(err, "deploy")
	}

	_, err = backend.WaitMined(ctx, tx)
	if err != nil {
		return errors.Wrap(err, "wait mined")
	}

	return nil
}

func deployToken(ctx context.Context, backend *ethbackend.Backend, salt string) error {
	txOpts, err := backend.BindOpts(ctx, deployer)
	if err != nil {
		return errors.Wrap(err, "bind opts")
	}

	factory, err := bindings.NewCreate3(create3Factory, backend)
	if err != nil {
		return errors.Wrap(err, "new create3")
	}

	initCode, err := contracts.PackInitCode(tokenABI, bindings.MockTokenBin)
	if err != nil {
		return errors.Wrap(err, "pack init code")
	}

	tx, err := factory.DeployWithRetry(txOpts, create3.HashSalt(salt), initCode) //nolint:contextcheck // Context is txOpts
	if err != nil {
		return errors.Wrap(err, "deploy")
	}

	_, err = backend.WaitMined(ctx, tx)
	if err != nil {
		return errors.Wrap(err, "wait mined")
	}

	return nil
}
