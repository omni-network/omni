package netman

import (
	"context"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/e2e/types"
	"github.com/omni-network/omni/lib/contracts"
	"github.com/omni-network/omni/lib/contracts/feeoraclev1"
	"github.com/omni-network/omni/lib/contracts/portal"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/forkjoin"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/common"
)

// Manager abstract logic to deploy and bootstrap a network.
// TODO(corver): No need to split public vs private portal deploys anymore.
type Manager interface {
	// DeployPublicPortals deploys portals to public chains, like arb-goerli.
	DeployPublicPortals(ctx context.Context, valSetID uint64, validators []bindings.Validator) error

	// DeployInfo returns the deployed network information by chain ID.
	// Note that the private chains has to be deterministic, since this is called before deploying private portals.
	DeployInfo() map[uint64]DeployInfo

	// DeployPrivatePortals deploys portals to private (docker) chains.
	DeployPrivatePortals(ctx context.Context, valSetID uint64, validators []bindings.Validator) error

	// Portals returns the deployed portals from both public and private chains.
	Portals() map[uint64]Portal
}

func NewManager(testnet types.Testnet, backends ethbackend.Backends) (Manager, error) {
	if testnet.OnlyMonitor {
		if !netconf.IsAny(testnet.Network, netconf.Omega, netconf.Mainnet) {
			return nil, errors.New("monitor-only only supported for testnet and mainnet")
		}

		return &manager{
			backends: backends,
		}, nil
	}

	network := testnet.Network

	// Create partial portals. This will be updated by Deploy*Portals.
	portals := make(map[uint64]Portal)

	// Private chains have deterministic deploy height and addresses.
	privateChainDeployInfo := DeployInfo{
		DeployHeight:  0,
		PortalAddress: contracts.Portal(network),
	}

	if testnet.HasOmniEVM() {
		// Just use the first omni evm instance for now.
		omniEVM := testnet.OmniEVMs[0]
		portals[omniEVM.Chain.ChainID] = Portal{
			Chain:      omniEVM.Chain,
			DeployInfo: privateChainDeployInfo,
		}
	}

	// Add all portals
	for _, anvil := range testnet.AnvilChains {
		portals[anvil.Chain.ChainID] = Portal{
			Chain:      anvil.Chain,
			DeployInfo: privateChainDeployInfo,
		}
	}
	// Add all public chains
	for _, public := range testnet.PublicChains {
		// Public chain deploy height and address may be statically populated.
		staticDeploy, _ := testnet.Network.Static().PortalDeployment(public.Chain().ChainID)
		portals[public.Chain().ChainID] = Portal{
			Chain: public.Chain(),
			DeployInfo: DeployInfo{
				PortalAddress: staticDeploy.Address,
				DeployHeight:  staticDeploy.DeployHeight,
			},
		}
	}

	switch testnet.Network {
	case netconf.Devnet:
		return &manager{
			portals:     portals,
			omniChainID: netconf.Devnet.Static().OmniExecutionChainID,
			backends:    backends,
			network:     netconf.Devnet,
		}, nil
	case netconf.Staging:
		return &manager{
			portals:     portals,
			omniChainID: netconf.Staging.Static().OmniExecutionChainID,
			backends:    backends,
			network:     netconf.Staging,
		}, nil
	case netconf.Omega:
		return &manager{
			portals:     portals,
			omniChainID: netconf.Omega.Static().OmniExecutionChainID,
			backends:    backends,
			network:     netconf.Omega,
		}, nil
	default:
		return nil, errors.New("unknown network", "network", network)
	}
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
	portals     map[uint64]Portal // Note that this is mutable, Portals are updated by Deploy*Portals.
	omniChainID uint64
	backends    ethbackend.Backends
	network     netconf.ID
}

func (m *manager) DeployInfo() map[uint64]DeployInfo {
	resp := make(map[uint64]DeployInfo)
	for _, portal := range m.portals {
		resp[portal.Chain.ChainID] = portal.DeployInfo
	}

	return resp
}

func (m *manager) DeployPublicPortals(ctx context.Context, valSetID uint64, validators []bindings.Validator,
) error {
	log.Info(ctx, "Deploying public portal contracts")

	// Define a forkjoin work function that will deploy the omni contracts for each chain
	deployFunc := func(ctx context.Context, p Portal) (*deployResult, error) {
		log.Debug(ctx, "Deploying to", "chain", p.Chain.Name)

		backend, err := m.backends.Backend(p.Chain.ChainID)
		if err != nil {
			return nil, errors.Wrap(err, "deploy opts", "chain", p.Chain.Name)
		}

		addr, deployHeight, err := m.deployIfNeeded(ctx, p.Chain, backend, valSetID, validators)
		if err != nil {
			return nil, errors.Wrap(err, "deploy public portals", "chain", p.Chain.Name)
		}

		contract, err := bindings.NewOmniPortal(addr, backend)
		if err != nil {
			return nil, errors.Wrap(err, "bind contract", "chain", p.Chain.Name)
		}

		return &deployResult{
			Contract: contract,
			Addr:     addr,
			Height:   deployHeight,
		}, nil
	}

	fork, join, cancel := forkjoin.New(ctx, deployFunc)
	defer cancel()
	for chainID := range m.portals {
		portal := m.portals[chainID]

		if !portal.Chain.IsPublic {
			continue // Only public chains are deployed here.
		}

		fork(portal)
	}

	for res := range join() {
		if res.Err != nil {
			return errors.Wrap(res.Err, "fork join")
		}

		portal := m.portals[res.Input.Chain.ChainID]

		portal.DeployInfo = DeployInfo{
			PortalAddress: res.Output.Addr,
			DeployHeight:  res.Output.Height,
		}
		portal.Contract = res.Output.Contract

		m.portals[res.Input.Chain.ChainID] = portal
	}

	return nil
}

func (m *manager) DeployPrivatePortals(ctx context.Context, valSetID uint64, validators []bindings.Validator,
) error {
	log.Info(ctx, "Deploying private portal contracts")

	// Define a forkjoin work function that will deploy the omni contracts for each chain
	deployFunc := func(ctx context.Context, p Portal) (deployResult, error) {
		backend, err := m.backends.Backend(p.Chain.ChainID)
		if err != nil {
			return deployResult{}, errors.Wrap(err, "deploy opts", "chain", p.Chain.Name)
		}

		addr, height, err := m.deployIfNeeded(ctx, p.Chain, backend, valSetID, validators)
		if err != nil {
			return deployResult{}, errors.Wrap(err, "deploy private portals", "chain", p.Chain.Name)
		} else if addr != p.DeployInfo.PortalAddress {
			return deployResult{}, errors.New("deployed address does not match existing address",
				"expected", p.DeployInfo.PortalAddress.Hex(),
				"actual", addr.Hex(),
				"chain", p.Chain.Name)
		}

		contract, err := bindings.NewOmniPortal(addr, backend)
		if err != nil {
			return deployResult{}, errors.Wrap(err, "bind contract", "chain", p.Chain.Name)
		}

		return deployResult{
			Contract: contract,
			Addr:     addr,
			Height:   height,
		}, nil
	}

	// Start the forkjoin
	fork, join, cancel := forkjoin.New(ctx, deployFunc)
	defer cancel()
	for chainID := range m.portals {
		portal := m.portals[chainID]
		if portal.Chain.IsPublic {
			continue // Public chains are already deployed.
		}

		fork(portal)
	}

	// Join the results
	for res := range join() {
		if res.Err != nil {
			return errors.Wrap(res.Err, "fork join")
		}

		// Update the portal with the deployed contract, height and address
		portal := m.portals[res.Input.Chain.ChainID]

		portal.Contract = res.Output.Contract
		portal.DeployInfo = DeployInfo{
			PortalAddress: res.Output.Addr,
			DeployHeight:  res.Output.Height,
		}

		m.portals[res.Input.Chain.ChainID] = portal
	}

	return nil
}

func (m *manager) Portals() map[uint64]Portal {
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

	feeOracle, _, err := feeoraclev1.Deploy(ctx, m.network, chain.ChainID, m.destChainIDs(chain.ChainID), m.backends)
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

// destChainIDs returns all configured destination chain ids for a given source.
func (m *manager) destChainIDs(srcChainID uint64) []uint64 {
	var destChainIDs []uint64

	for id := range m.Portals() {
		if id != srcChainID {
			destChainIDs = append(destChainIDs, id)
		}
	}

	return destChainIDs
}

type deployResult struct {
	Contract *bindings.OmniPortal
	Addr     common.Address
	Height   uint64
}
