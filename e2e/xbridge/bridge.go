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
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"

	"golang.org/x/sync/errgroup"
)

type BridgeConfig struct {
	ProxyAdminOwner common.Address
	Admin           common.Address
	Authorizer      common.Address
	Pauser          common.Address
	Unpauser        common.Address
	OmniPortal      common.Address
	Token           common.Address
	Lockbox         common.Address
}

func (cfg BridgeConfig) Validate() error {
	if isEmpty(cfg.Admin) {
		return errors.New("admin is zero")
	}
	if isEmpty(cfg.Authorizer) {
		return errors.New("authorizer is zero")
	}
	if isEmpty(cfg.Pauser) {
		return errors.New("pauser is zero")
	}
	if isEmpty(cfg.Unpauser) {
		return errors.New("unpauser is zero")
	}
	if isEmpty(cfg.OmniPortal) {
		return errors.New("omni portal is zero")
	}
	if isEmpty(cfg.Token) {
		return errors.New("token is zero")
	}

	return nil
}

func BridgeAddr(ctx context.Context, network netconf.ID, xtoken types.XToken) (common.Address, error) {
	salt, err := BridgeSalt(ctx, network, xtoken)
	if err != nil {
		return common.Address{}, errors.Wrap(err, "salt")
	}

	return contracts.Create3Address(network, salt), nil
}

func BridgeSalt(ctx context.Context, network netconf.ID, xtoken types.XToken) (string, error) {
	net := network.String()

	if network == netconf.Staging {
		v, err := contracts.StagingID(ctx)
		if err != nil {
			return "", errors.Wrap(err, "staging id")
		}
		net = v
	}

	return fmt.Sprintf("%s-%s-bridge", net, xtoken.Symbol()), nil
}

func deployBridges(ctx context.Context, network netconf.Network, backends ethbackend.Backends, xtoken types.XToken) error {
	var eg errgroup.Group

	for _, chain := range network.EVMChains() {
		eg.Go(func() error {
			backend, err := backends.Backend(chain.ID)
			if err != nil {
				return errors.Wrap(err, "get backend", "chain", chain.Name)
			}

			err = deployBridge(ctx, network.ID, chain, backend, xtoken)
			if err != nil {
				return errors.Wrap(err, "deploy bridge", "chain", chain.Name)
			}

			err = configureRoutes(ctx, network, chain, backend, xtoken)
			if err != nil {
				return errors.Wrap(err, "set routes", "chain", chain.Name)
			}

			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		return errors.Wrap(err, "deploy all")
	}

	return nil
}

// deployBridge deploys a new bridge contract and returns the address and receipt.
func deployBridge(
	ctx context.Context,
	network netconf.ID,
	chain netconf.Chain,
	backend *ethbackend.Backend,
	xtoken types.XToken,
) error {
	addrs, err := contracts.GetAddresses(ctx, network)
	if err != nil {
		return errors.Wrap(err, "get addrs")
	}

	token, err := xtoken.Address(ctx, network)
	if err != nil {
		return errors.Wrap(err, "xtoken address")
	}

	canon, err := xtoken.Canonical(ctx, network)
	if err != nil {
		return errors.Wrap(err, "canonical")
	}

	salt, err := BridgeSalt(ctx, network, xtoken)
	if err != nil {
		return errors.Wrap(err, "salt")
	}

	deploy := func(cfg BridgeConfig) error {
		gasLimits := xtoken.GasLimits()

		addr, receipt, err := proxy.Deploy(ctx, backend, proxy.DeployParams{
			Network:     network,
			Create3Salt: salt,
			DeployImpl: func(txOpts *bind.TransactOpts, backend *ethbackend.Backend) (common.Address, *ethtypes.Transaction, error) {
				addr, tx, _, err := bindings.DeployBridge(txOpts, backend, gasLimits.BridgeNoLockbox, gasLimits.BridgeWithLockbox)
				return addr, tx, err
			},
			PackInitCode: func(impl common.Address) ([]byte, error) {
				return packBridgeInitCode(cfg, impl)
			},
		})

		if err != nil {
			return err
		}

		log.Info(ctx, "Bridge deployed", "addr", addr, "tx", maybeTxHash(receipt), "xtoken", xtoken.Symbol(), "chain", chain.Name)

		return nil
	}

	cfg := BridgeConfig{
		ProxyAdminOwner: eoa.MustAddress(network, eoa.RoleUpgrader),
		Admin:           eoa.MustAddress(network, eoa.RoleManager),
		Authorizer:      eoa.MustAddress(network, eoa.RoleManager),
		Pauser:          eoa.MustAddress(network, eoa.RoleManager),
		Unpauser:        eoa.MustAddress(network, eoa.RoleManager),
		OmniPortal:      addrs.Portal,
		Token:           token,
	}

	// lockbox only required on chain with canonical deployment
	if canon.ChainID != chain.ID {
		return deploy(cfg)
	}

	lockbock, err := LockboxAddr(ctx, network, xtoken)
	if err != nil {
		return errors.Wrap(err, "lockbox address")
	}

	cfg.Lockbox = lockbock

	return deploy(cfg)
}

func configureRoutes(
	ctx context.Context,
	network netconf.Network,
	chain netconf.Chain,
	backend *ethbackend.Backend,
	xtoken types.XToken,
) error {
	addr, err := BridgeAddr(ctx, network.ID, xtoken)
	if err != nil {
		return errors.Wrap(err, "bridge addr")
	}

	txOpts, err := backend.BindOpts(ctx, eoa.MustAddress(network.ID, eoa.RoleManager))
	if err != nil {
		return errors.Wrap(err, "bind opts")
	}

	canon, err := xtoken.Canonical(ctx, network.ID)
	if err != nil {
		return errors.Wrap(err, "canonical")
	}

	var destChainIDs []uint64
	var routes []bindings.IBridgeRoute
	for _, dest := range network.EVMChains() {
		if dest.ID == chain.ID {
			continue
		}

		destChainIDs = append(destChainIDs, dest.ID)
		routes = append(routes, bindings.IBridgeRoute{
			Bridge:     addr,
			HasLockbox: dest.ID == canon.ChainID,
		})
	}

	bridge, err := bindings.NewBridge(addr, backend)
	if err != nil {
		return errors.Wrap(err, "new bridge")
	}

	tx, err := bridge.ConfigureRoutes(txOpts, destChainIDs, routes)
	if err != nil {
		return errors.Wrap(err, "set destinations")
	}

	_, err = backend.WaitMined(ctx, tx)
	if err != nil {
		return errors.Wrap(err, "wait mined")
	}

	tx, err = bridge.AuthorizeRoutes(txOpts, destChainIDs)
	if err != nil {
		return errors.Wrap(err, "authorize routes")
	}

	receipt, err := backend.WaitMined(ctx, tx)
	if err != nil {
		return errors.Wrap(err, "wait mined")
	}

	log.Info(ctx, "Routes set", "bridge", addr, "xtoken", xtoken.Symbol(), "chain", chain.Name, "bridge", addr, "tx", maybeTxHash(receipt))

	return nil
}

// packBridgeInitCode packs the initialization code for the bridge contract proxy.
func packBridgeInitCode(cfg BridgeConfig, impl common.Address) ([]byte, error) {
	bridgeAbi, err := bindings.BridgeMetaData.GetAbi()
	if err != nil {
		return nil, errors.Wrap(err, "get abi")
	}

	proxyAbi, err := bindings.TransparentUpgradeableProxyMetaData.GetAbi()
	if err != nil {
		return nil, errors.Wrap(err, "get proxy abi")
	}

	initializer, err := bridgeAbi.Pack("initialize", cfg.Admin, cfg.Authorizer, cfg.Pauser, cfg.Unpauser, cfg.OmniPortal, cfg.Token, cfg.Lockbox)
	if err != nil {
		return nil, errors.Wrap(err, "encode initializer")
	}

	return contracts.PackInitCode(proxyAbi, bindings.TransparentUpgradeableProxyBin, impl, cfg.ProxyAdminOwner, initializer)
}
