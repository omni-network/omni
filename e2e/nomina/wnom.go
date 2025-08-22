package nomina

import (
	"context"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/halo/genutil/evm/predeploys"
	"github.com/omni-network/omni/lib/contracts/proxy"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/common"
)

func isWnomTokenDeployed(ctx context.Context, backend *ethbackend.Backend) (bool, error) {
	wnom := common.HexToAddress(predeploys.WNomina)

	impl, err := proxy.Impl(ctx, backend, wnom)
	if err != nil {
		return false, errors.Wrap(err, "impl")
	}

	if impl == (common.Address{}) {
		return false, nil
	}

	return true, nil
}

func deployWnomTokenIfNeeded(ctx context.Context, network netconf.Network, backends ethbackend.Backends) error {
	omni, ok := network.OmniEVMChain()
	if !ok {
		return errors.New("no omni chain")
	}

	backend, err := backends.Backend(omni.ID)
	if err != nil {
		return errors.Wrap(err, "get backend")
	}

	deployed, err := isWnomTokenDeployed(ctx, backend)
	if err != nil {
		return errors.Wrap(err, "is wnom token deployed")
	}

	if deployed {
		return nil
	}

	return deployWnomToken(ctx, network.ID, backend)
}

func deployWnomToken(ctx context.Context, network netconf.ID, backend *ethbackend.Backend) error {
	wnom := common.HexToAddress(predeploys.WNomina)
	deployer := eoa.MustAddress(network, eoa.RoleDeployer)
	upgrader := eoa.MustAddress(network, eoa.RoleUpgrader)

	admin, err := proxy.Admin(ctx, backend, wnom)
	if err != nil {
		return errors.Wrap(err, "admin")
	}

	proxyAdmin, err := bindings.NewProxyAdmin(admin, backend)
	if err != nil {
		return errors.Wrap(err, "new proxy admin")
	}

	deployerTxOpts, err := backend.BindOpts(ctx, deployer)
	if err != nil {
		return errors.Wrap(err, "bind opts")
	}

	impl, tx, _, err := bindings.DeployWNomina(deployerTxOpts, backend)
	if err != nil {
		return errors.Wrap(err, "deploy wnomina")
	}

	_, err = backend.WaitMined(ctx, tx)
	if err != nil {
		return errors.Wrap(err, "deploy wait mined")
	}

	upgraderTxOpts, err := backend.BindOpts(ctx, upgrader)
	if err != nil {
		return errors.Wrap(err, "bind opts")
	}

	tx, err = proxyAdmin.UpgradeAndCall(upgraderTxOpts, wnom, impl, nil)
	if err != nil {
		return errors.Wrap(err, "upgrade wnomina")
	}

	_, err = backend.WaitMined(ctx, tx)
	if err != nil {
		return errors.Wrap(err, "upgrade wait mined")
	}

	tx, err = proxyAdmin.RenounceOwnership(upgraderTxOpts)
	if err != nil {
		return errors.Wrap(err, "renounce ownership")
	}

	_, err = backend.WaitMined(ctx, tx)
	if err != nil {
		return errors.Wrap(err, "renounce wait mined")
	}

	log.Info(ctx, "WNomina token deployed", "address", wnom, "impl", impl, "network", network)

	return nil
}
