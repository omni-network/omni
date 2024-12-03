package devapp

import (
	"context"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/lib/cast"
	"github.com/omni-network/omni/lib/contracts"
	"github.com/omni-network/omni/lib/create3"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"

	"cosmossdk.io/math"
)

const (
	l1VaultSalt = "l1-vault"
	l1TokenSalt = "l1-token"
	l2TokenSalt = "l2-token"
)

// Deploy deploys the mock tokens and vaults to devnet.
func Deploy(ctx context.Context, network netconf.Network, backends ethbackend.Backends) error {
	if network.ID != netconf.Devnet {
		return errors.New("onl devnet")
	}

	l1Backend, err := backends.Backend(static.L1.ChainID)
	if err != nil {
		return errors.Wrap(err, "backend mock l1")
	}

	l2Backend, err := backends.Backend(static.L2.ChainID)
	if err != nil {
		return errors.Wrap(err, "backend mock l2")
	}

	if err := deployToken(ctx, l1Backend, l1TokenSalt); err != nil {
		return errors.Wrap(err, "deploy l1 token")
	}

	if err := deployVault(ctx, l1Backend, l1VaultSalt, static.L1Token); err != nil {
		return errors.Wrap(err, "deploy vault")
	}

	if err := deployToken(ctx, l2Backend, l2TokenSalt); err != nil {
		return errors.Wrap(err, "deploy l2 token")
	}

	if err := fundSolver(ctx, l1Backend, static.L1Token); err != nil {
		return errors.Wrap(err, "fund solver")
	}

	return nil
}

func fundSolver(ctx context.Context, backend *ethbackend.Backend, tokenAddr common.Address) error {
	mngr := eoa.MustAddress(netconf.Devnet, eoa.RoleManager) // we use mngr to mint, but this doesn't matter. could be any dev addr
	slvr := eoa.MustAddress(netconf.Devnet, eoa.RoleSolver)

	token, err := bindings.NewMockToken(tokenAddr, backend)
	if err != nil {
		return errors.Wrap(err, "new mock token")
	}

	txOpts, err := backend.BindOpts(ctx, mngr)
	if err != nil {
		return errors.Wrap(err, "bind opts")
	}

	eth1m := math.NewInt(1_000_000).MulRaw(params.Ether).BigInt()
	tx, err := token.Mint(txOpts, slvr, eth1m)
	if err != nil {
		return errors.Wrap(err, "mint")
	}

	_, err = backend.WaitMined(ctx, tx)
	if err != nil {
		return errors.Wrap(err, "wait mined")
	}

	return nil
}

// AllowOutboxCalls allows the outbox to call the L1 vault deposit method.
func AllowOutboxCalls(ctx context.Context, network netconf.Network, backends ethbackend.Backends) error {
	if network.ID != netconf.Devnet {
		return errors.New("onl devnet")
	}

	addrs, err := contracts.GetAddresses(ctx, network.ID)
	if err != nil {
		return errors.Wrap(err, "get addresses")
	}

	mockl1, ok := network.Chain(evmchain.IDMockL1)
	if !ok {
		return errors.New("no mock l1")
	}

	mockl1Backend, err := backends.Backend(mockl1.ID)
	if err != nil {
		return errors.Wrap(err, "backend mock l1")
	}

	if err := allowCalls(ctx, mockl1Backend, addrs.SolveOutbox); err != nil {
		return errors.Wrap(err, "allow calls")
	}

	return nil
}

// allowCalls allows the outbox to call the L1 vault deposit method.
func allowCalls(ctx context.Context, backend *ethbackend.Backend, outboxAddr common.Address) error {
	outbox, err := bindings.NewSolveOutbox(outboxAddr, backend)
	if err != nil {
		return errors.Wrap(err, "new solve outbox")
	}

	txOpts, err := backend.BindOpts(ctx, manager)
	if err != nil {
		return errors.Wrap(err, "bind opts")
	}

	vaultDepositID, err := cast.Array4(vaultDeposit.ID[:4])
	if err != nil {
		return err
	}

	tx, err := outbox.SetAllowedCall(txOpts, static.L1Vault, vaultDepositID, true)
	if err != nil {
		return errors.Wrap(err, "set allowed call")
	} else if _, err := backend.WaitMined(ctx, tx); err != nil {
		return errors.Wrap(err, "wait mined")
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
		return errors.Wrap(err, "deploy proxy")
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
