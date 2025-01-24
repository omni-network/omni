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
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"

	"cosmossdk.io/math"
)

const (
	l1VaultSalt = "l1-vault"
	l1TokenSalt = "l1-token"
	l2TokenSalt = "l2-token"
)

// MaybeDeploy deploys the mock tokens / vaults, and mints solver with mock tokens, if needed.
func MaybeDeploy(ctx context.Context, network netconf.ID, backends ethbackend.Backends) error {
	if !network.IsEphemeral() {
		return errors.New("only ephemeral networks")
	}

	app := MustGetApp(network)

	l1Backend, err := backends.Backend(app.L1.ChainID)
	if err != nil {
		return errors.Wrap(err, "backend mock l1")
	}

	l2Backend, err := backends.Backend(app.L2.ChainID)
	if err != nil {
		return errors.Wrap(err, "backend mock l2")
	}

	if err := maybeDeployToken(ctx, network, l1Backend, l1TokenSalt); err != nil {
		return errors.Wrap(err, "deploy l1 token")
	}

	if err := maybeDeployVault(ctx, network, l1Backend, l1VaultSalt, app.L1Token); err != nil {
		return errors.Wrap(err, "deploy vault")
	}

	if err := maybeDeployToken(ctx, network, l2Backend, l2TokenSalt); err != nil {
		return errors.Wrap(err, "deploy l2 token")
	}

	if err := maybeFundSolver(ctx, network, l1Backend, app.L1Token); err != nil {
		return errors.Wrap(err, "fund solver")
	}

	return nil
}

func maybeFundSolver(ctx context.Context, network netconf.ID, backend *ethbackend.Backend, tokenAddr common.Address) error {
	funder := eoa.MustAddress(network, eoa.RoleHot)
	solver := eoa.MustAddress(network, eoa.RoleSolver)

	token, err := bindings.NewMockERC20(tokenAddr, backend)
	if err != nil {
		return errors.Wrap(err, "new mock token")
	}

	txOpts, err := backend.BindOpts(ctx, funder)
	if err != nil {
		return errors.Wrap(err, "bind opts")
	}

	eth1m := math.NewInt(1_000_000).MulRaw(params.Ether).BigInt()
	eth1k := math.NewInt(1_000).MulRaw(params.Ether).BigInt()

	balance, err := token.BalanceOf(&bind.CallOpts{Context: ctx}, solver)
	if err != nil {
		return errors.Wrap(err, "balance of")
	}

	// if more than 1k, do nothing. if less, mint 1m
	if balance.Cmp(eth1k) >= 0 {
		return nil
	}

	tx, err := token.Mint(txOpts, solver, eth1m)
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
func AllowOutboxCalls(ctx context.Context, network netconf.ID, backends ethbackend.Backends) error {
	if !network.IsEphemeral() {
		return errors.New("only ephemeral networks")
	}

	app := MustGetApp(network)

	addrs, err := contracts.GetAddresses(ctx, network)
	if err != nil {
		return errors.Wrap(err, "get addresses")
	}

	l1Backend, err := backends.Backend(app.L1.ChainID)
	if err != nil {
		return errors.Wrap(err, "backend mock l1")
	}

	if err := allowCalls(ctx, network, l1Backend, addrs.SolveOutbox); err != nil {
		return errors.Wrap(err, "allow calls")
	}

	return nil
}

// allowCalls allows the outbox to call the L1 vault deposit method.
func allowCalls(ctx context.Context, network netconf.ID, backend *ethbackend.Backend, outboxAddr common.Address) error {
	app := MustGetApp(network)
	manager := eoa.MustAddress(network, eoa.RoleManager)

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

	tx, err := outbox.SetAllowedCall(txOpts, app.L1Vault, vaultDepositID, true)
	if err != nil {
		return errors.Wrap(err, "set allowed call")
	} else if _, err := backend.WaitMined(ctx, tx); err != nil {
		return errors.Wrap(err, "wait mined")
	}

	return nil
}

func maybeDeployVault(ctx context.Context, network netconf.ID, backend *ethbackend.Backend, salt string, collaterlTkn common.Address) error {
	deployer := eoa.MustAddress(network, eoa.RoleDeployer)

	txOpts, err := backend.BindOpts(ctx, deployer)
	if err != nil {
		return errors.Wrap(err, "bind opts")
	}

	factory, err := bindings.NewCreate3(contracts.Create3Factory(network), backend)
	if err != nil {
		return errors.Wrap(err, "new create3")
	}

	addr, deployed, err := isDeployed(ctx, backend, factory, deployer, salt)
	if err != nil {
		return errors.Wrap(err, "is deployed")
	}

	if deployed {
		log.Info(ctx, "MockVault already deployed", "addr", addr, "salt", salt)
		return nil
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

	log.Info(ctx, "MockVault deployed", "addr", addr, "salt", salt)

	return nil
}

func maybeDeployToken(ctx context.Context, network netconf.ID, backend *ethbackend.Backend, salt string) error {
	deployer := eoa.MustAddress(network, eoa.RoleDeployer)

	txOpts, err := backend.BindOpts(ctx, deployer)
	if err != nil {
		return errors.Wrap(err, "bind opts")
	}

	factory, err := bindings.NewCreate3(contracts.Create3Factory(network), backend)
	if err != nil {
		return errors.Wrap(err, "new create3")
	}

	addr, deployed, err := isDeployed(ctx, backend, factory, deployer, salt)
	if err != nil {
		return errors.Wrap(err, "is deployed")
	}

	if deployed {
		log.Info(ctx, "MockToken already deployed", "addr", addr, "salt", salt)
		return nil
	}

	initCode, err := contracts.PackInitCode(tokenABI, bindings.MockERC20Bin, "Mock Token", "MTK")
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

	log.Info(ctx, "MockToken deployed", "addr", addr, "salt", salt)

	return nil
}

func isDeployed(ctx context.Context, backend *ethbackend.Backend,
	factory *bindings.Create3, deployer common.Address, salt string) (common.Address, bool, error) {
	addr, err := factory.GetDeployed(&bind.CallOpts{Context: ctx}, deployer, create3.HashSalt(salt))
	if err != nil {
		return addr, false, errors.Wrap(err, "get deployed")
	}

	code, err := backend.CodeAt(ctx, addr, nil)
	if err != nil {
		return addr, false, errors.Wrap(err, "code at")
	}

	if len(code) > 0 {
		return addr, true, nil
	}

	return addr, false, nil
}
