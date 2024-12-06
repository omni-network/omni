package avs

import (
	"context"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

var (
	delegationManager = common.HexToAddress("0x39053D51B77DC0d36036Fc1fCc8Cb819df8Ef37A")
	avsDirectory      = common.HexToAddress("0x135DDa560e946695d6f155dACaFC6f1F25C1F5AF")
)

type implDeploymentConfig struct {
	Deployer          common.Address
	DelegationManager common.Address
	AVSDirectory      common.Address
}

func isEmpty(addr common.Address) bool {
	return addr == common.Address{}
}

func (cfg implDeploymentConfig) Validate() error {
	if isEmpty(cfg.Deployer) {
		return errors.New("deployer is zero")
	}
	if isEmpty(cfg.DelegationManager) {
		return errors.New("delegation manager is zero")
	}
	if isEmpty(cfg.AVSDirectory) {
		return errors.New("avs directory is zero")
	}

	return nil
}

// DeployImpl deploys the OmniAVS implementation contract without upgrading its proxy
// The proxy is managed by a manual multisig external from e2e, and will be upgraded manually.
func DeployImpl(ctx context.Context, network netconf.ID, backend *ethbackend.Backend) (common.Address, *ethtypes.Receipt, error) {
	cfg := implDeploymentConfig{
		Deployer:          eoa.MustAddress(network, eoa.RoleDeployer),
		DelegationManager: delegationManager,
		AVSDirectory:      avsDirectory,
	}
	if err := cfg.Validate(); err != nil {
		return common.Address{}, nil, errors.Wrap(err, "validate cfg")
	}

	txOpts, err := backend.BindOpts(ctx, cfg.Deployer)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "bind opts")
	}

	impl, tx, _, err := bindings.DeployOmniAVS(txOpts, backend, cfg.DelegationManager, cfg.AVSDirectory)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "deploy omni avs impl")
	}

	receipt, err := backend.WaitMined(ctx, tx)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "wait mined")
	}

	return impl, receipt, nil
}
