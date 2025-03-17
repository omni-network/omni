package solve

import (
	"context"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/lib/contracts"
	"github.com/omni-network/omni/lib/create3"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/tokens"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"golang.org/x/sync/errgroup"
)

type MockToken struct {
	tokens.Token
	ChainID   uint64
	NetworkID netconf.ID
}

var mocks = []MockToken{
	// staging mock wstETH
	{Token: tokens.WSTETH, ChainID: evmchain.IDBaseSepolia, NetworkID: netconf.Staging},

	// devnet L1 mock wstETH
	{Token: tokens.WSTETH, ChainID: evmchain.IDMockL1, NetworkID: netconf.Devnet},

	// devnet L2 mock wstETH
	{Token: tokens.WSTETH, ChainID: evmchain.IDMockL2, NetworkID: netconf.Devnet},
}

func MockTokens() []MockToken {
	return mocks
}

func (m MockToken) Address() common.Address {
	return create3.Address(
		contracts.Create3Factory(m.NetworkID),
		m.Create3Salt(),
		eoa.MustAddress(m.NetworkID, eoa.RoleDeployer),
	)
}

func (m MockToken) Create3Salt() string {
	return m.NetworkID.String() + "::mock::" + m.Symbol
}

// maybeDeployMockTokens deploys mintable ERC20 tokens on ephemeral networks, for solvernet testing.
func maybeDeployMockTokens(ctx context.Context, network netconf.Network, backends ethbackend.Backends) error {
	if !network.ID.IsEphemeral() {
		return nil
	}

	var eg errgroup.Group

	for _, mock := range mocks {
		eg.Go(func() error {
			if mock.NetworkID != network.ID {
				return nil
			}

			chain, ok := network.Chain(mock.ChainID)
			if !ok {
				return errors.New("mock token chain not in network", "chain", mock.ChainID)
			}

			backend, err := backends.Backend(chain.ID)
			if err != nil {
				return errors.Wrap(err, "get backend", "chain", chain.Name)
			}

			return maybeDeployMockToken(ctx, network.ID, backend, mock)
		})
	}

	if err := eg.Wait(); err != nil {
		return errors.Wrap(err, "deploy tokens")
	}

	return nil
}

func maybeDeployMockToken(ctx context.Context, network netconf.ID, backend *ethbackend.Backend, mock MockToken) error {
	deployer := eoa.MustAddress(network, eoa.RoleDeployer)

	txOpts, err := backend.BindOpts(ctx, deployer)
	if err != nil {
		return errors.Wrap(err, "bind opts")
	}

	factory, err := bindings.NewCreate3(contracts.Create3Factory(network), backend)
	if err != nil {
		return errors.Wrap(err, "new create3")
	}

	salt := mock.Create3Salt()

	addr, deployed, err := isDeployed(ctx, backend, factory, deployer, salt)
	if err != nil {
		return errors.Wrap(err, "is deployed")
	}

	if deployed {
		log.Info(ctx, "MockToken already deployed", "addr", addr, "salt", salt, "name", mock.Name, "symbol", mock.Symbol)
		return nil
	}

	abi, err := bindings.MockERC20MetaData.GetAbi()
	if err != nil {
		return errors.Wrap(err, "get abi")
	}

	initCode, err := contracts.PackInitCode(abi, bindings.MockERC20Bin, mock.Name, mock.Symbol)
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

	log.Info(ctx, "MockToken deployed", "addr", addr, "salt", salt, "name", mock.Name, "symbol", mock.Symbol)

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
