package netman

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/e2e/types"
	"github.com/omni-network/omni/lib/anvil"
	"github.com/omni-network/omni/lib/contracts/portal"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/forkjoin"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/txmgr"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
)

//nolint:gochecknoglobals // Static mapping.
var (
	// this must be different than netman.Operator(), otherwise the txmgr nonce
	// get's out of sync, because the relayer TXMgr and the operator TXMgr are running
	// in separate services.
	privateRelayerKey = anvil.DevPrivateKey5()

	// fbDev is the address of the fireblocks "dev" account.
	fbDev = common.HexToAddress("0x7a6cF389082dc698285474976d7C75CAdE08ab7e")
)

// Manager abstract logic to deploy and bootstrap a network.
type Manager interface {
	// DeployPublicPortals deploys portals to public chains, like arb-goerli.
	DeployPublicPortals(ctx context.Context, valSetID uint64, validators []bindings.Validator) error

	// DeployInfo returns the deployed network information.
	// Note that the private chains has to be deterministic, since this is called before deploying private portals.
	DeployInfo() map[types.EVMChain]DeployInfo

	// DeployPrivatePortals deploys portals to private (docker) chains.
	DeployPrivatePortals(ctx context.Context, valSetID uint64, validators []bindings.Validator) error

	// Portals returns the deployed portals from both public and private chains.
	Portals() map[uint64]Portal

	// RelayerKey returns the relayer private key hex.
	RelayerKey() *ecdsa.PrivateKey

	// Operator returns the address of the account that operates the network.
	Operator() common.Address
}

func NewManager(testnet types.Testnet, backends ethbackend.Backends, relayerKeyFile string) (Manager, error) {
	if testnet.OnlyMonitor {
		if testnet.Network != netconf.Testnet {
			return nil, errors.New("the AVS contract is currently only deployed to testnet")
		}

		return &manager{
			backends: backends,
		}, nil
	}

	network := testnet.Network

	privPortalAddr, found := portal.AddrForNetwork(network)
	if !found {
		return nil, errors.New("unknown network", "network", network)
	}

	// Create partial portals. This will be updated by Deploy*Portals.
	portals := make(map[uint64]Portal)

	// Private chains have deterministic deploy height and addresses.
	privateChainDeployInfo := DeployInfo{
		DeployHeight:  0,
		PortalAddress: privPortalAddr,
	}

	if testnet.HasOmniEVM() {
		// Just use the first omni evm instance for now.
		omniEVM := testnet.OmniEVMs[0]
		portals[omniEVM.Chain.ID] = Portal{
			Chain:      omniEVM.Chain,
			DeployInfo: privateChainDeployInfo,
		}
	}

	// Add all portals
	for _, anvil := range testnet.AnvilChains {
		portals[anvil.Chain.ID] = Portal{
			Chain:      anvil.Chain,
			DeployInfo: privateChainDeployInfo,
		}
	}
	// Add all public chains
	for _, public := range testnet.PublicChains {
		portals[public.Chain.ID] = Portal{
			Chain: public.Chain,
			// Public chain deploy height and address will be updated by DeployPublicPortals.
		}
	}

	switch testnet.Network {
	case netconf.Devnet:
		if relayerKeyFile != "" {
			return nil, errors.New("relayer keys not supported in devnet")
		}

		return &manager{
			portals:     portals,
			omniChainID: netconf.Devnet.Static().OmniExecutionChainID,
			relayerKey:  privateRelayerKey,
			backends:    backends,
			network:     netconf.Devnet,
			operator:    anvil.DevAccount4(),
		}, nil
	case netconf.Staging:
		relayerKey, err := crypto.LoadECDSA(relayerKeyFile)
		if err != nil {
			return nil, errors.Wrap(err, "read relayer key file", "path", relayerKeyFile)
		}

		return &manager{
			portals:     portals,
			omniChainID: netconf.Staging.Static().OmniExecutionChainID,
			relayerKey:  relayerKey,
			backends:    backends,
			network:     netconf.Staging,
			operator:    fbDev,
		}, nil
	case netconf.Testnet:
		// no testnet relayer
		return &manager{
			portals:     portals,
			omniChainID: netconf.Testnet.Static().OmniExecutionChainID,
			backends:    backends,
			network:     netconf.Testnet,
			operator:    fbDev,
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
	relayerKey  *ecdsa.PrivateKey
	backends    ethbackend.Backends
	network     netconf.ID
	operator    common.Address
}

func (m *manager) DeployInfo() map[types.EVMChain]DeployInfo {
	resp := make(map[types.EVMChain]DeployInfo)
	for _, portal := range m.portals {
		resp[portal.Chain] = portal.DeployInfo
	}

	return resp
}

func (m *manager) DeployPublicPortals(ctx context.Context, valSetID uint64, validators []bindings.Validator,
) error {
	// Log provided key balances for public chains (just FYI).
	for _, portal := range m.portals {
		if !portal.Chain.IsPublic {
			continue // Only log public chain balances.
		}

		txOpts, backend, err := m.backends.BindOpts(ctx, portal.Chain.ID, m.operator)
		if err != nil {
			return errors.Wrap(err, "deploy opts", "chain", portal.Chain.Name)
		}

		if err := logBalance(ctx, backend, portal.Chain.Name, txOpts.From, "deploy_key"); err != nil {
			return err
		}

		relayerAddr := crypto.PubkeyToAddress(m.relayerKey.PublicKey)
		if err := logBalance(ctx, backend, portal.Chain.Name, relayerAddr, "relayer_key"); err != nil {
			return err
		}
	}

	log.Info(ctx, "Deploying public portal contracts")

	// Define a forkjoin work function that will deploy the omni contracts for each chain
	deployFunc := func(ctx context.Context, p Portal) (*deployResult, error) {
		log.Debug(ctx, "Deploying to", "chain", p.Chain.Name)

		backend, err := m.backends.Backend(p.Chain.ID)
		if err != nil {
			return nil, errors.Wrap(err, "deploy opts", "chain", p.Chain.Name)
		}

		height, err := backend.BlockNumber(ctx)
		if err != nil {
			return nil, errors.Wrap(err, "block number", "chain", p.Chain.Name)
		}

		addr, _, err := portal.DeployIfNeeded(ctx, m.network, backend, valSetID, validators)
		if err != nil {
			return nil, errors.Wrap(err, "deploy public omni contracts", "chain", p.Chain.Name)
		}

		contract, err := bindings.NewOmniPortal(addr, backend)
		if err != nil {
			return nil, errors.Wrap(err, "bind contract", "chain", p.Chain.Name)
		}

		return &deployResult{
			Contract: contract,
			Addr:     addr,
			Height:   height,
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

		portal := m.portals[res.Input.Chain.ID]

		portal.DeployInfo = DeployInfo{
			PortalAddress: res.Output.Addr,
			DeployHeight:  res.Output.Height,
		}
		portal.Contract = res.Output.Contract

		m.portals[res.Input.Chain.ID] = portal
		log.Info(ctx, "Deployed public portal contract", "chain", portal.Chain.Name, "address", res.Output.Addr.Hex(), "height", res.Output.Height)
	}

	return nil
}

func (m *manager) DeployPrivatePortals(ctx context.Context, valSetID uint64, validators []bindings.Validator,
) error {
	log.Info(ctx, "Deploying private portal contracts")

	// Define a forkjoin work function that will deploy the omni contracts for each chain
	deployFunc := func(ctx context.Context, p Portal) (*bindings.OmniPortal, error) {
		chain := p.Chain.Name
		backend, err := m.backends.Backend(p.Chain.ID)
		if err != nil {
			return nil, errors.Wrap(err, "deploy opts", "chain", chain)
		}

		addr, _, err := portal.DeployIfNeeded(ctx, m.network, backend, valSetID, validators)

		if err != nil {
			return nil, errors.Wrap(err, "deploy private omni contracts", "chain", chain)
		} else if addr != p.DeployInfo.PortalAddress {
			return nil, errors.New("deployed address does not match existing address",
				"expected", p.DeployInfo.PortalAddress.Hex(),
				"actual", addr.Hex(),
				"chain", chain)
		}

		contract, err := bindings.NewOmniPortal(addr, backend)
		if err != nil {
			return nil, errors.Wrap(err, "bind contract", "chain", chain)
		}

		return contract, nil
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

		// Update the portal with the deployed contract
		portal := m.portals[res.Input.Chain.ID]
		portal.Contract = res.Output
		m.portals[res.Input.Chain.ID] = portal
	}

	return m.fundPrivateRelayer(ctx)
}

func (m *manager) Portals() map[uint64]Portal {
	return m.portals
}

func (m *manager) RelayerKey() *ecdsa.PrivateKey {
	return m.relayerKey
}

func (m *manager) Operator() common.Address {
	return m.operator
}

func (m *manager) fundPrivateRelayer(ctx context.Context) error {
	if privateRelayerKey.Equal(m.relayerKey) {
		return nil // No need to fund relayer if key is private.
	}

	relayerAddr := crypto.PubkeyToAddress(m.relayerKey.PublicKey)

	for _, portal := range m.portals {
		if portal.Chain.IsPublic {
			continue // We use relayer key for public chain, it should already be funded.
		}

		backend, err := m.backends.Backend(portal.Chain.ID)
		if err != nil {
			return errors.Wrap(err, "backend", "chain", portal.Chain.Name)
		}

		pctx := log.WithCtx(ctx, "chain", portal.Chain.Name)

		bal, err := backend.BalanceAt(pctx, m.operator, nil)
		if err != nil {
			return err
		}
		b, _ := bal.Float64()
		log.Info(pctx, "Funding relayer operator balance", "balance", b/params.Ether, "operator", m.operator.Hex())
		// TODO(corver): Remove this debug log

		tx, _, err := backend.Send(pctx, m.operator, txmgr.TxCandidate{
			To:       &relayerAddr,
			GasLimit: 100_000,                                                    // 100k is fine,
			Value:    new(big.Int).Mul(big.NewInt(10), big.NewInt(params.Ether)), // 10 ETH
		})
		if err != nil {
			return errors.Wrap(err, "send ether")
		} else if _, err := backend.WaitMined(ctx, tx); err != nil {
			return errors.Wrap(err, "wait mined")
		}
	}

	return nil
}

type deployResult struct {
	Contract *bindings.OmniPortal
	Addr     common.Address
	Height   uint64
}
