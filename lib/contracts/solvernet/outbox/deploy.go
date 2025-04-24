package outbox

import (
	"context"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/lib/contracts"
	"github.com/omni-network/omni/lib/contracts/solvernet"
	"github.com/omni-network/omni/lib/create3"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/common"
)

type DeploymentConfig struct {
	Create3Factory  common.Address
	Create3Salt     string
	ProxyAdminOwner common.Address
	Owner           common.Address
	Solver          common.Address
	Portal          common.Address
	Inbox           common.Address
	Deployer        common.Address
	ExpectedAddr    common.Address
	Executor        common.Address
}

func (cfg DeploymentConfig) Validate() error {
	if (cfg.Create3Factory == common.Address{}) {
		return errors.New("create3 factory is zero")
	}
	if cfg.Create3Salt == "" {
		return errors.New("create3 salt is empty")
	}
	if (cfg.ProxyAdminOwner == common.Address{}) {
		return errors.New("proxy admin is zero")
	}
	if (cfg.Deployer == common.Address{}) {
		return errors.New("deployer is not set")
	}
	if (cfg.Owner == common.Address{}) {
		return errors.New("owner is not set")
	}
	if (cfg.Inbox == common.Address{}) {
		return errors.New("inbox is zero")
	}
	if (cfg.Portal == common.Address{}) {
		return errors.New("portal is zero")
	}
	if (cfg.Solver == common.Address{}) {
		return errors.New("solver is zero")
	}
	if (cfg.ExpectedAddr == common.Address{}) {
		return errors.New("expected address is zero")
	}
	if (cfg.Executor == common.Address{}) {
		return errors.New("executor is zero")
	}

	return nil
}

// Deploy idempotently deploys a new SolverNetOutbox contract and returns the address and receipt.
func Deploy(ctx context.Context, network netconf.Network, backend *ethbackend.Backend) (common.Address, *ethclient.Receipt, error) {
	addrs, err := contracts.GetAddresses(ctx, network.ID)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "get addresses")
	}

	isDeployed, err := contracts.IsDeployed(ctx, backend, addrs.SolverNetOutbox)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "is deployed")
	} else if isDeployed {
		return addrs.SolverNetOutbox, nil, nil
	}

	salts, err := contracts.GetSalts(ctx, network.ID)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "get salts")
	}

	cfg := DeploymentConfig{
		Create3Factory:  addrs.Create3Factory,
		Create3Salt:     salts.SolverNetOutbox,
		Owner:           eoa.MustAddress(network.ID, eoa.RoleManager),
		Deployer:        eoa.MustAddress(network.ID, eoa.RoleDeployer),
		ProxyAdminOwner: eoa.MustAddress(network.ID, eoa.RoleUpgrader),
		Solver:          eoa.MustAddress(network.ID, eoa.RoleSolver),
		Portal:          addrs.Portal,
		Inbox:           addrs.SolverNetInbox,
		ExpectedAddr:    addrs.SolverNetOutbox,
		Executor:        addrs.SolverNetExecutor,
	}

	return deploy(ctx, cfg, network, backend)
}

func deploy(ctx context.Context, cfg DeploymentConfig, network netconf.Network, backend *ethbackend.Backend) (common.Address, *ethclient.Receipt, error) {
	if err := cfg.Validate(); err != nil {
		return common.Address{}, nil, errors.Wrap(err, "validate config")
	}

	chainID, err := backend.ChainID(ctx)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "get chain id")
	}

	mailbox, err := solvernet.HyperlaneMailbox(chainID.Uint64())
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "get hyperlane mailbox")
	}

	txOpts, err := backend.BindOpts(ctx, cfg.Deployer)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "bind opts")
	}

	factory, err := bindings.NewCreate3(cfg.Create3Factory, backend)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "new create3")
	}

	salt := create3.HashSalt(cfg.Create3Salt)

	addr, err := factory.GetDeployed(nil, txOpts.From, salt)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "get deployed")
	} else if (cfg.ExpectedAddr != common.Address{}) && addr != cfg.ExpectedAddr {
		return common.Address{}, nil, errors.New("unexpected address", "expected", cfg.ExpectedAddr, "actual", addr)
	}

	impl, tx, _, err := bindings.DeploySolverNetOutbox(txOpts, backend, mailbox)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "deploy impl")
	}

	_, err = backend.WaitMined(ctx, tx)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "wait mined impl")
	}

	initCode, err := packInitCode(cfg, impl)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "pack init code")
	}

	tx, err = factory.DeployWithRetry(txOpts, salt, initCode) //nolint:contextcheck // Context is txOpts
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "deploy proxy")
	}

	receipt, err := backend.WaitMined(ctx, tx)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "wait mined proxy")
	}

	// setInboxes
	var chainIDs []uint64
	var inboxes []bindings.ISolverNetOutboxInboxConfig
	for _, chain := range network.EVMChains() {
		if chain.ID == chainID.Uint64() {
			continue
		}

		provider, err := solvernet.Provider(chain.ID)
		if err != nil {
			return common.Address{}, nil, errors.Wrap(err, "get provider")
		}

		if provider != 0 {
			chainIDs = append(chainIDs, chain.ID)
			inboxes = append(inboxes, bindings.ISolverNetOutboxInboxConfig{
				Inbox:    cfg.Inbox,
				Provider: provider,
			})
		}
	}

	txOpts, err = backend.BindOpts(ctx, cfg.Owner)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "bind opts")
	}

	inbox, err := bindings.NewSolverNetOutbox(addr, backend)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "new inbox")
	}

	tx, err = inbox.SetInboxes(txOpts, chainIDs, inboxes)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "set inboxes")
	}

	_, err = backend.WaitMined(ctx, tx)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "wait mined set inboxes")
	}

	return addr, receipt, nil
}

func packInitCode(cfg DeploymentConfig, impl common.Address) ([]byte, error) {
	outboxAbi, err := bindings.SolverNetOutboxMetaData.GetAbi()
	if err != nil {
		return nil, errors.Wrap(err, "get abi")
	}

	proxyAbi, err := bindings.TransparentUpgradeableProxyMetaData.GetAbi()
	if err != nil {
		return nil, errors.Wrap(err, "get proxy abi")
	}

	initializer, err := outboxAbi.Pack("initialize", cfg.Owner, cfg.Solver, cfg.Portal, cfg.Executor)
	if err != nil {
		return nil, errors.Wrap(err, "encode initializer")
	}

	return contracts.PackInitCode(proxyAbi, bindings.TransparentUpgradeableProxyBin, impl, cfg.ProxyAdminOwner, initializer)
}
