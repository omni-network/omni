package netman

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/test/e2e/types"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
)

const (
	// privKeyHex0 of pre-funded anvil account 0.
	privKeyHex0 = "0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
	// privKeyHex1 of pre-funded anvil account 1.
	privKeyHex1 = "0x59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d"

	// Second contract address of privKeyHex0 (first is FeeOracleV1 @ 0x5FbDB2315678afecb367f032d93F642f64180aa3).
	privatePortalAddr = "0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512"
)

//nolint:gochecknoglobals // Static mapping.
var (
	privateDeployKey  = mustHexToKey(privKeyHex0)
	privateRelayerKey = mustHexToKey(privKeyHex1)
)

// Manager abstract logic to deploy and bootstrap a extNetwork.
type Manager interface {
	// DeployPublicPortals deploys portals to public chains, like arb-goerli.
	DeployPublicPortals(ctx context.Context) error

	// DeployInfo returns the deployed network information.
	// Note that the private chains has to be deterministic, since this is called before deploying private portals.
	DeployInfo() map[types.EVMChain]DeployInfo

	// DeployPrivatePortals deploys portals to private (docker) chains.
	DeployPrivatePortals(ctx context.Context) error

	// Portals returns the deployed portals from both public and private chains.
	Portals() map[uint64]Portal

	// RelayerKey returns the relayer private key hex.
	RelayerKey() *ecdsa.PrivateKey
}

func NewManager(testnet types.Testnet, deployKeyFile string,
	relayerKeyFile string,
) (Manager, error) {
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
		RPCURL:     omniEVM.ExternalRPC,
		DeployInfo: privateChainDeployInfo,
	}
	// Add all portals
	for _, anvil := range testnet.AnvilChains {
		portals[anvil.Chain.ID] = Portal{
			Chain:      anvil.Chain,
			RPCURL:     anvil.ExternalRPC,
			DeployInfo: privateChainDeployInfo,
		}
	}
	// Add all public chains
	for _, public := range testnet.PublicChains {
		portals[public.Chain.ID] = Portal{
			Chain:  public.Chain,
			RPCURL: public.RPCAddress,
			// Public chain deploy height and address will be updated by DeployPublicPortals.
		}
	}

	// Instantiate all clients
	for chainID := range portals {
		portal := portals[chainID]
		client, err := ethclient.Dial(portals[chainID].RPCURL)
		if err != nil {
			return nil, errors.Wrap(err, "dial rpc", "chain", portal.Chain.Name, "url", portal.RPCURL)
		}
		portal.Client = client
		portals[chainID] = portal
	}

	switch testnet.Network {
	case netconf.Devnet:
		if deployKeyFile != "" || relayerKeyFile != "" {
			return nil, errors.New("deploy and relayer keys not supported in devnet")
		}

		return &manager{
			portals:    portals,
			relayerKey: privateRelayerKey,
		}, nil
	case netconf.Staging:
		deployKey, err := crypto.LoadECDSA(deployKeyFile)
		if err != nil {
			return nil, errors.Wrap(err, "read deploy key file", "path", deployKeyFile)
		}
		relayerKey, err := crypto.LoadECDSA(relayerKeyFile)
		if err != nil {
			return nil, errors.Wrap(err, "read relayer key file", "path", relayerKeyFile)
		}

		return &manager{
			portals:         portals,
			publicDeployKey: deployKey,
			relayerKey:      relayerKey,
		}, nil
	default:
		return nil, errors.New("unknown extNetwork")
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
	RPCURL     string
	DeployInfo DeployInfo
	Client     *ethclient.Client
	Contract   *bindings.OmniPortal
	txOpts     *bind.TransactOpts
}

// TxOpts returns transaction options using the deploy key.
func (p Portal) TxOpts(ctx context.Context) *bind.TransactOpts {
	clone := *p.txOpts
	clone.Context = ctx

	return &clone
}

var _ Manager = (*manager)(nil)

type manager struct {
	portals         map[uint64]Portal // Note that this is mutable, Portals are updated by Deploy*Portals.
	publicDeployKey *ecdsa.PrivateKey
	relayerKey      *ecdsa.PrivateKey
}

func (m *manager) DeployInfo() map[types.EVMChain]DeployInfo {
	resp := make(map[types.EVMChain]DeployInfo)
	for _, portal := range m.portals {
		resp[portal.Chain] = portal.DeployInfo
	}

	return resp
}

func (m *manager) DeployPublicPortals(ctx context.Context) error {
	// Log provided key balances for public chains (just FYI).
	for _, portal := range m.portals {
		if !portal.Chain.IsPublic {
			continue // Only log public chain balances.
		}
		if err := logBalance(ctx, portal.Client, m.publicDeployKey, "deploy_key"); err != nil {
			return err
		}
		if err := logBalance(ctx, portal.Client, m.relayerKey, "relayer_key"); err != nil {
			return err
		}
	}

	log.Info(ctx, "Deploying public portal contracts")

	for chainID := range m.portals {
		portal := m.portals[chainID]

		if !portal.Chain.IsPublic {
			continue // Only public chains are deployed here.
		}

		height, err := portal.Client.BlockNumber(ctx)
		if err != nil {
			return errors.Wrap(err, "get block number")
		}

		addr, contract, txops, err := deployContract(ctx, chainID, portal.Client, m.publicDeployKey)
		if err != nil {
			return errors.Wrap(err, "deploy public portal contract")
		}

		portal.DeployInfo = DeployInfo{
			PortalAddress: addr,
			DeployHeight:  height,
		}
		portal.Contract = contract
		portal.txOpts = txops

		m.portals[chainID] = portal
		log.Info(ctx, "Deployed public portal contract", "chain", portal.Chain.Name, "address", addr.Hex(), "height", height)
	}

	return nil
}

func (m *manager) DeployPrivatePortals(ctx context.Context) error {
	log.Info(ctx, "Deploying private portal contracts")

	for chainID := range m.portals {
		portal := m.portals[chainID]
		if portal.Chain.IsPublic {
			continue // Public chains are already deployed.
		}

		addr, contract, txops, err := deployContract(ctx, chainID, portal.Client, privateDeployKey)
		if err != nil {
			return errors.Wrap(err, "deploy public portal contract")
		} else if addr != portal.DeployInfo.PortalAddress {
			return errors.New("deployed address does not match existing address",
				"expected", portal.DeployInfo.PortalAddress.Hex(),
				"actual", addr.Hex())
		}

		portal.Contract = contract
		portal.txOpts = txops

		m.portals[chainID] = portal
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
	fromKey := privateRelayerKey
	toKey := m.relayerKey

	if fromKey.Equal(toKey) {
		return nil // No need to fund relayer if key is private.
	}

	fromAddr := crypto.PubkeyToAddress(fromKey.PublicKey)
	toAddr := crypto.PubkeyToAddress(toKey.PublicKey)

	for _, portal := range m.portals {
		if portal.Chain.IsPublic {
			continue // We use relayer key for public chain, it should already be funded.
		}

		ethCl := portal.Client

		nonce, err := ethCl.PendingNonceAt(ctx, fromAddr)
		if err != nil {
			return errors.Wrap(err, "get nonce")
		}

		price, err := ethCl.SuggestGasPrice(ctx)
		if err != nil {
			return errors.Wrap(err, "get gas price")
		}

		txData := ethtypes.LegacyTx{
			Nonce:    nonce,
			GasPrice: price,
			Gas:      100_000, // 100k is fine
			To:       &toAddr,
			Value:    new(big.Int).Mul(big.NewInt(10), big.NewInt(params.Ether)), // 10 ETH
		}

		signer := ethtypes.LatestSignerForChainID(big.NewInt(int64(portal.Chain.ID)))
		tx, err := ethtypes.SignNewTx(fromKey, signer, &txData)
		if err != nil {
			return errors.Wrap(err, "sign tx")
		}

		if err := ethCl.SendTransaction(ctx, tx); err != nil {
			return errors.Wrap(err, "send tx")
		}
	}

	return nil
}
