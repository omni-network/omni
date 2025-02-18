package xbridge

import (
	"context"
	"fmt"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/e2e/xbridge/types"
	"github.com/omni-network/omni/lib/contracts"
	"github.com/omni-network/omni/lib/contracts/proxy"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

type LockboxConfig struct {
	ProxyAdminOwner common.Address
	Admin           common.Address
	Pauser          common.Address
	Unpauser        common.Address
	XToken          common.Address
	Token           common.Address
}

func (cfg LockboxConfig) Validate() error {
	if isEmpty(cfg.ProxyAdminOwner) {
		return errors.New("proxy admin is zero")
	}
	if isEmpty(cfg.Admin) {
		return errors.New("admin is zero")
	}
	if isEmpty(cfg.Pauser) {
		return errors.New("pauser is zero")
	}
	if isEmpty(cfg.Unpauser) {
		return errors.New("unpauser is zero")
	}
	if isEmpty(cfg.XToken) {
		return errors.New("xtoken is zero")
	}
	if isEmpty(cfg.Token) {
		return errors.New("token is zero")
	}

	return nil
}

func LockboxAddr(ctx context.Context, network netconf.ID, xtoken types.XToken) (common.Address, error) {
	salt, err := LockboxSalt(ctx, network, xtoken)
	if err != nil {
		return common.Address{}, errors.Wrap(err, "salt")
	}

	return contracts.Create3Address(network, salt), nil
}

func LockboxSalt(ctx context.Context, network netconf.ID, xtoken types.XToken) (string, error) {
	net := network.String()

	if network == netconf.Staging {
		v, err := contracts.StagingID(ctx)
		if err != nil {
			return "", errors.Wrap(err, "staging id")
		}
		net = v
	}

	return fmt.Sprintf("%s-%s-lockbox", net, xtoken.Symbol()), nil
}

// deployLockbox deploys a new lockbox contract and returns the address and receipt.
func deployLockbox(ctx context.Context, network netconf.ID, backends ethbackend.Backends, xtoken types.XToken) error {
	xtkn, err := xtoken.Address(ctx, network)
	if err != nil {
		return errors.Wrap(err, "token address")
	}

	canon, err := xtoken.Canonical(ctx, network)
	if err != nil {
		return errors.Wrap(err, "canonical")
	}

	chain, ok := evmchain.MetadataByID(canon.ChainID)
	if !ok {
		return errors.New("chain meata")
	}

	backend, err := backends.Backend(canon.ChainID)
	if err != nil {
		return errors.Wrap(err, "get backend")
	}

	salt, err := LockboxSalt(ctx, network, xtoken)
	if err != nil {
		return errors.Wrap(err, "lockbox salt")
	}

	cfg := LockboxConfig{
		ProxyAdminOwner: eoa.MustAddress(network, eoa.RoleUpgrader),
		Admin:           eoa.MustAddress(network, eoa.RoleManager),
		Pauser:          eoa.MustAddress(network, eoa.RoleManager),
		Unpauser:        eoa.MustAddress(network, eoa.RoleManager),
		XToken:          xtkn,
		Token:           canon.Address,
	}

	addr, receipt, err := proxy.Deploy(ctx, backend, proxy.DeployParams{
		Network:     network,
		Create3Salt: salt,
		DeployImpl: func(txOpts *bind.TransactOpts, backend *ethbackend.Backend) (common.Address, *ethtypes.Transaction, error) {
			addr, tx, _, err := bindings.DeployLockbox(txOpts, backend)
			return addr, tx, err
		},
		PackInitCode: func(impl common.Address) ([]byte, error) {
			return packLockboxInitCode(cfg, impl)
		},
	})
	if err != nil {
		return errors.Wrap(err, "deploy")
	}

	log.Info(ctx, "Lockbox deployed", "addr", addr, "tx", maybeTxHash(receipt), "xtoken", xtoken.Symbol(), "chain", chain.Name)

	return nil
}

// packLockboxInitCode packs the initialization code for the lockbox contract proxy.
func packLockboxInitCode(cfg LockboxConfig, impl common.Address) ([]byte, error) {
	lockboxAbi, err := bindings.LockboxMetaData.GetAbi()
	if err != nil {
		return nil, errors.Wrap(err, "get abi")
	}

	proxyAbi, err := bindings.TransparentUpgradeableProxyMetaData.GetAbi()
	if err != nil {
		return nil, errors.Wrap(err, "get proxy abi")
	}

	initializer, err := lockboxAbi.Pack("initialize", cfg.Admin, cfg.Pauser, cfg.Unpauser, cfg.Token, cfg.XToken)
	if err != nil {
		return nil, errors.Wrap(err, "encode initializer")
	}

	return contracts.PackInitCode(proxyAbi, bindings.TransparentUpgradeableProxyBin, impl, cfg.ProxyAdminOwner, initializer)
}
