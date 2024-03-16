package netman

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/forkjoin"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/txmgr"
	"github.com/omni-network/omni/test/e2e/types"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
)

const (

	// privKeyHex1 of pre-funded anvil account 1.
	privKeyHex1 = "0x59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d"

	// Fifth contract address of privKeyHex0 (ProxyAdmin, FeeOracleV1Impl, FeeOracleV1Proxy, PortalImpl come first).
	privatePortalAddr = "0xDc64a140Aa3E981100a9becA4E685f962f0cF6C9"
)

//nolint:gochecknoglobals // Static mapping.
var (
	privateRelayerKey = mustHexToKey(privKeyHex1)
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
}

func NewManager(testnet types.Testnet, backends ethbackend.Backends, relayerKeyFile string) (Manager, error) {
	if testnet.OnlyMonitor {
		if testnet.Name != netconf.Testnet {
			return nil, errors.New("the AVS contract is currently only deployed to testnet")
		}

		return &manager{
			backends: backends,
		}, nil
	}

	// Create partial portals. This will be updated by Deploy*Portals.
	portals := make(map[uint64]Portal)

	// Private chains have deterministic deploy height and addresses.
	privateChainDeployInfo := DeployInfo{
		DeployHeight:  0,
		PortalAddress: common.HexToAddress(privatePortalAddr),
	}

	// Just use the first omni evm instance for now.
	omniEVM := testnet.OmniEVMs[0]
	portals[omniEVM.Chain.ID] = Portal{
		Chain:      omniEVM.Chain,
		DeployInfo: privateChainDeployInfo,
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
			omniChainID: omniEVM.Chain.ID,
			relayerKey:  privateRelayerKey,
			backends:    backends,
		}, nil
	case netconf.Staging:
		relayerKey, err := crypto.LoadECDSA(relayerKeyFile)
		if err != nil {
			return nil, errors.Wrap(err, "read relayer key file", "path", relayerKeyFile)
		}

		return &manager{
			portals:     portals,
			omniChainID: omniEVM.Chain.ID,
			relayerKey:  relayerKey,
			backends:    backends,
		}, nil
	default:
		return nil, errors.New("unknown network")
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

		_, txOpts, backend, err := m.backends.BindOpts(ctx, portal.Chain.ID)
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
	deployFunc := func(ctx context.Context, portal Portal) (*deployResult, error) {
		log.Debug(ctx, "Deploying to", "chain", portal.Chain.Name)

		_, txOpts, backend, err := m.backends.BindOpts(ctx, portal.Chain.ID)
		if err != nil {
			return nil, errors.Wrap(err, "deploy opts", "chain", portal.Chain.Name)
		}

		height, err := backend.BlockNumber(ctx)
		if err != nil {
			return nil, errors.Wrap(err, "get block number", "chain", portal.Chain.Name)
		}

		addr, contract, err := deployOmniContracts(
			ctx, txOpts, backend, valSetID, validators,
		)
		if err != nil {
			return nil, errors.Wrap(err, "deploy public omni contracts", "chain", portal.Chain.Name)
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
	deployFunc := func(ctx context.Context, portal Portal) (*bindings.OmniPortal, error) {
		chain := portal.Chain.Name
		_, txOpts, backend, err := m.backends.BindOpts(ctx, portal.Chain.ID)
		if err != nil {
			return nil, errors.Wrap(err, "deploy opts", "chain", chain)
		}

		addr, contract, err := deployOmniContracts(ctx, txOpts, backend, valSetID, validators)
		if err != nil {
			return nil, errors.Wrap(err, "deploy private omni contracts", "chain", chain)
		} else if addr != portal.DeployInfo.PortalAddress {
			return nil, errors.New("deployed address does not match existing address",
				"expected", portal.DeployInfo.PortalAddress.Hex(),
				"actual", addr.Hex(),
				"chain", chain)
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

func (m *manager) fundPrivateRelayer(ctx context.Context) error {
	if privateRelayerKey.Equal(m.relayerKey) {
		return nil // No need to fund relayer if key is private.
	}

	relayerAddr := crypto.PubkeyToAddress(m.relayerKey.PublicKey)

	for _, portal := range m.portals {
		if portal.Chain.IsPublic {
			continue // We use relayer key for public chain, it should already be funded.
		}

		addr, _, backend, err := m.backends.BindOpts(ctx, portal.Chain.ID)
		if err != nil {
			return errors.Wrap(err, "deploy opts")
		}

		tx, _, err := backend.Send(ctx, addr, txmgr.TxCandidate{
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
