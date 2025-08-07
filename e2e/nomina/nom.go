package nomina

import (
	"context"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/lib/cast"
	"github.com/omni-network/omni/lib/contracts"
	"github.com/omni-network/omni/lib/createx"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type nomDeploymentConfig struct {
	CreateXSalt    string
	CreateXFactory common.Address
	ExpectedAddr   common.Address
	Deployer       common.Address
	MintAuthority  common.Address
}

func isNomTokenDeployed(ctx context.Context, network netconf.ID, backend *ethbackend.Backend) (bool, error) {
	nomToken := contracts.NomAddr(network)

	code, err := backend.CodeAt(ctx, nomToken, nil)
	if err != nil {
		return false, errors.Wrap(err, "code at", "address", nomToken)
	}

	if len(code) == 0 {
		return false, nil
	}

	return true, nil
}

func validateNomTokenDeploymentConfig(cfg nomDeploymentConfig) error {
	if cfg.CreateXSalt == "" {
		return errors.New("createx salt is empty")
	}
	if isEmpty(cfg.CreateXFactory) {
		return errors.New("createx factory is zero")
	}
	if isEmpty(cfg.ExpectedAddr) {
		return errors.New("expected address is zero")
	}
	if isEmpty(cfg.Deployer) {
		return errors.New("deployer is zero")
	}
	if isEmpty(cfg.MintAuthority) {
		return errors.New("mint authority is zero")
	}

	return nil
}

func deployNomTokenIfNeeded(ctx context.Context, network netconf.Network, backends ethbackend.Backends) error {
	ethereum, ok := network.EthereumChain()
	if !ok {
		return errors.New("no ethereum chain")
	}

	backend, err := backends.Backend(ethereum.ID)
	if err != nil {
		return errors.Wrap(err, "get backend")
	}

	deployed, err := isNomTokenDeployed(ctx, network.ID, backend)
	if err != nil {
		return errors.Wrap(err, "is nom token deployed")
	}

	if deployed {
		return nil
	}

	return deployNomToken(ctx, network.ID, backend)
}

func deployNomToken(ctx context.Context, network netconf.ID, backend *ethbackend.Backend) error {
	addrs, err := contracts.GetAddresses(ctx, network)
	if err != nil {
		return errors.Wrap(err, "get addresses")
	}

	salts, err := contracts.GetSalts(ctx, network)
	if err != nil {
		return errors.Wrap(err, "get salts")
	}

	cfg := nomDeploymentConfig{
		CreateXSalt:    salts.NomToken,
		CreateXFactory: addrs.CreateXFactory,
		ExpectedAddr:   addrs.NomToken,
		Deployer:       eoa.MustAddress(network, eoa.RoleDeployer),
		MintAuthority:  eoa.MustAddress(network, eoa.RoleCold), // TODO(zodomo): Replace with proper mint authority
	}

	if err := validateNomTokenDeploymentConfig(cfg); err != nil {
		return errors.Wrap(err, "validate nom token deployment config")
	}

	txOpts, err := backend.BindOpts(ctx, cfg.Deployer)
	if err != nil {
		return errors.Wrap(err, "bind opts")
	}

	factory, err := bindings.NewICreateX(cfg.CreateXFactory, backend)
	if err != nil {
		return errors.Wrap(err, "new createx factory")
	}

	abi, err := bindings.NominaMetaData.GetAbi()
	if err != nil {
		return errors.Wrap(err, "get nomina abi")
	}

	initCode, err := contracts.PackInitCode(abi, bindings.NominaMetaData.Bin, addrs.Token, cfg.MintAuthority)
	if err != nil {
		return errors.Wrap(err, "pack nomina init code")
	}

	initCodeHash := crypto.Keccak256Hash(initCode)

	saltBytes := []byte(cfg.CreateXSalt)
	var salt [32]byte
	var guardedSalt [32]byte

	// If salt is exactly 32 bytes and starts with deployer address, it's pre-formatted
	if len(saltBytes) == 32 && cast.MustEthAddress(saltBytes[:20]) == cfg.Deployer {
		// Use the original salt bytes directly (already deployer-formatted)
		copy(salt[:], saltBytes)
	} else {
		// For string salts, hash them to get a 32-byte salt
		hashedSalt := crypto.Keccak256Hash(saltBytes)
		copy(salt[:], hashedSalt[:])
	}

	guardedSalt = createx.GuardSalt(salt, cfg.Deployer)

	// Confirm our expected address matches the factory's derivation
	addr, err := factory.ComputeCreate2Address(nil, guardedSalt, initCodeHash)
	if err != nil {
		return errors.Wrap(err, "compute create2 address")
	}

	if addr == (common.Address{}) || (cfg.ExpectedAddr != common.Address{} && addr != cfg.ExpectedAddr) {
		return errors.New("unexpected address", "expected", cfg.ExpectedAddr, "actual", addr)
	}

	tx, err := factory.DeployCreate2(txOpts, salt, initCode)
	if err != nil {
		return errors.Wrap(err, "deploy create2")
	}

	_, err = backend.WaitMined(ctx, tx)
	if err != nil {
		return errors.Wrap(err, "wait mined")
	}

	log.Info(ctx, "Nomina token deployed", "address", addr, "network", network)

	return nil
}
