package netman

import (
	"context"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/e2e/types"
	"github.com/omni-network/omni/lib/contracts/feeoraclev1"
	"github.com/omni-network/omni/lib/contracts/portal"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/expbackoff"
	"github.com/omni-network/omni/lib/forkjoin"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/common"
)

// Manager abstract logic to deploy portal contracts.
type Manager interface {
	// DeployPortals deploys portals to all chains.
	DeployPortals(ctx context.Context, valSetID uint64, validators []bindings.Validator) error

	// Portals returns the deployed portals from both public and private chains.
	Portals() map[uint64]Portal
}

func NewManager(testnet types.Testnet, backends ethbackend.Backends) (Manager, error) {
	return &manager{
		backends: backends,
		network:  testnet.Network,
		chains:   testnet.EVMChains(),
	}, nil
}

// DeployInfo contains the deployed portal address and height.
type DeployInfo struct {
	PortalAddress common.Address
	DeployHeight  uint64
}

// Portal contains all deployed portal information and state.
type Portal struct {
	Chain      types.EVMChain
	DeployInfo DeployInfo
	Contract   *bindings.OmniPortal
}

var _ Manager = (*manager)(nil)

type manager struct {
	chains   []types.EVMChain
	portals  map[uint64]Portal // This is nil until populated by DeployPortals.
	backends ethbackend.Backends
	network  netconf.ID
}

func (m *manager) DeployPortals(ctx context.Context, valSetID uint64, validators []bindings.Validator,
) error {
	if m.portals != nil {
		return errors.New("portals already deployed")
	}

	log.Info(ctx, "Deploying portal contracts")

	// Define a forkjoin work function that will deploy the omni contracts for each chain
	deployFunc := func(ctx context.Context, chain types.EVMChain) (Portal, error) {
		ctx = log.WithCtx(ctx, "chain", chain.Name)
		log.Debug(ctx, "Deploying portal")

		backend, err := m.backends.Backend(chain.ChainID)
		if err != nil {
			return Portal{}, errors.Wrap(err, "deploy opts")
		}

		addr, deployHeight, err := m.deployIfNeeded(ctx, chain, backend, valSetID, validators)
		if err != nil {
			return Portal{}, errors.Wrap(err, "deploy public portals")
		}

		contract, err := bindings.NewOmniPortal(addr, backend)
		if err != nil {
			return Portal{}, errors.Wrap(err, "bind contract")
		}

		return Portal{
			Chain: chain,
			DeployInfo: DeployInfo{
				PortalAddress: addr,
				DeployHeight:  deployHeight,
			},
			Contract: contract,
		}, nil
	}

	results, cancel := forkjoin.NewWithInputs(ctx, deployFunc, m.chains)
	defer cancel()

	m.portals = make(map[uint64]Portal)
	for res := range results {
		if res.Err != nil {
			return errors.Wrap(res.Err, "fork join", "chain", res.Input.Name)
		}

		m.portals[res.Output.Chain.ChainID] = res.Output
	}

	return nil
}

func (m *manager) Portals() map[uint64]Portal {
	if m.portals == nil {
		panic("portals not deployed yet")
	}

	return m.portals
}

// deployIfNeeded deploys a portal if it is not already deployed.
//
// In the case it is deployed, it:
//   - returns an error if the network is ephemeral
//   - returns an error if the deployment is not set in the static network static
//   - else, it returns the deployment address and height.
//
// In the case it is not deployed, it:
//   - returns an error if the deployment is set in the static network static
//   - else, it deploys the portal and returns the deployment address and height.
func (m *manager) deployIfNeeded(ctx context.Context, chain types.EVMChain, backend *ethbackend.Backend, valSetID uint64, validators []bindings.Validator,
) (common.Address, uint64, error) {
	isDeployed, addr, err := portal.IsDeployed(ctx, m.network, backend)
	if err != nil {
		return common.Address{}, 0, errors.Wrap(err, "is deployed", "chain", chain)
	}

	staticDeploy, hasStatic := m.network.Static().PortalDeployment(chain.ChainID)

	// for ephemeral networks, require that the portal is not already deployed
	if isDeployed && m.network.IsEphemeral() {
		return common.Address{}, 0, errors.New("ephemeral portal already deployed", "network", m.network, "chain", chain.Name, "address", addr.Hex())
	}

	// if the portal is deployed, but not set in the network static, return an error
	if isDeployed && !hasStatic {
		return common.Address{}, 0, errors.New("portal deployed, but not set in network static", "chain", chain.Name, "address", addr.Hex())
	}

	// if the portal is not deployed, but set in the network static, return an error
	if !isDeployed && hasStatic {
		return common.Address{}, 0, errors.New("portal not deployed, but set in network static", "chain", chain.Name)
	}

	// if the static deployment is set, return it
	if hasStatic {
		return staticDeploy.Address, staticDeploy.DeployHeight, nil
	}

	// Deploying fee oracle sporadically fails during gas estimation. Just retry a few times.
	var feeOracle common.Address
	err = expbackoff.Retry(ctx, func() error {
		feeOracle, _, err = feeoraclev1.Deploy(ctx, m.network, chain.ChainID, m.chainIDs(), m.backends)
		return err
	})
	if err != nil {
		return common.Address{}, 0, errors.Wrap(err, "deploy fee oracle", "chain", chain.Name)
	}

	// at this point, we need to deploy the portal
	addr, receipt, err := portal.Deploy(ctx, m.network, backend, feeOracle, valSetID, validators)
	if err != nil {
		return common.Address{}, 0, errors.Wrap(err, "deploy public omni contracts", "chain", chain.Name)
	} else if receipt == nil {
		return common.Address{}, 0, errors.New("no receipt", "chain", chain.Name)
	}
	log.Info(ctx, "Deployed portal", "chain", chain.Name, "address", addr.Hex(), "height", receipt.BlockNumber.Uint64())

	return addr, receipt.BlockNumber.Uint64(), nil
}

// chainIDs returns all chain ids.
func (m *manager) chainIDs() []uint64 {
	ids := make([]uint64, 0, len(m.chains))
	for _, chain := range m.chains {
		ids = append(ids, chain.ChainID)
	}

	return ids
}
