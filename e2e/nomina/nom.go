package nomina

import (
	"context"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/lib/contracts"
	"github.com/omni-network/omni/lib/createx"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
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
		MintAuthority:  eoa.MustAddress(network, eoa.RoleNomAuthority),
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

	if network.IsEphemeral() {
		abi, err = bindings.MockNominaMetaData.GetAbi()
		if err != nil {
			return errors.Wrap(err, "get mock nomina abi")
		}

		initCode, err = contracts.PackInitCode(abi, bindings.MockNominaMetaData.Bin, addrs.Token)
		if err != nil {
			return errors.Wrap(err, "pack mock nomina init code")
		}
	}

	initCodeHash := crypto.Keccak256Hash(initCode)

	var salt [32]byte
	var guardedSalt [32]byte

	// Normalise salt: prefer 0x-hex that decodes to 32 bytes, else raw 32-byte string, else hash.
	if b, err := hexutil.Decode(cfg.CreateXSalt); err == nil && len(b) == 32 {
		copy(salt[:], b)
	} else if len([]byte(cfg.CreateXSalt)) == 32 {
		copy(salt[:], []byte(cfg.CreateXSalt))
	} else {
		hashed := crypto.Keccak256Hash([]byte(cfg.CreateXSalt))
		copy(salt[:], hashed[:])
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
