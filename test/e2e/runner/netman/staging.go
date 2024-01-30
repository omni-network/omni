package netman

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
)

const arbGoerli = "arb_goerli"

// Staging returns the default e2e Staging network configuration.
// Note that arb_goerli portal address is defined by deployer flag used at runtime.
func defaultStaging() netconf.Network {
	return netconf.Network{
		Name: netconf.Devnet,
		Chains: []netconf.Chain{
			{
				ID:            1, // From static/geth_genesis.json
				Name:          "omni_evm",
				RPCURL:        "http://localhost:8545",
				AuthRPCURL:    "http://localhost:8551",
				PortalAddress: "0x5FbDB2315678afecb367f032d93F642f64180aa3",
				IsOmni:        true,
			},
			{
				ID:            421613, // From https://chainlist.org/?search=arbitrum+goerli&testnets=true
				Name:          arbGoerli,
				RPCURL:        "https://arbitrum-goerli.publicnode.com",
				PortalAddress: "", // Defined at runtime
				DeployHeight:  0,  // Defined at runtime
			},
		},
	}
}

type stagingManager struct {
	network    netconf.Network
	portals    map[uint64]Portal
	deployKey  *ecdsa.PrivateKey
	relayerKey *ecdsa.PrivateKey
}

func (m *stagingManager) DeployPublicPortals(ctx context.Context) error {
	// Log provided key balances for public chains (just FYI).
	for _, chain := range m.network.Chains {
		if !publicChains[chain.Name] {
			continue // Only log public chain balances.
		}
		if err := logBalance(ctx, chain, m.deployKey, "deploy_key"); err != nil {
			return err
		}
		if err := logBalance(ctx, chain, m.relayerKey, "relayer_key"); err != nil {
			return err
		}
	}

	portals, network, err := deployPublicContracts(ctx, m.network, m.deployKey)
	if err != nil {
		return err
	}
	m.network = network
	m.portals = portals

	return nil
}

func (m *stagingManager) DeployPrivatePortals(ctx context.Context) error {
	portals, err := deployPrivateContracts(ctx, m.network, privKey0)
	if err != nil {
		return err
	}

	m.portals = mergePortals(m.portals, portals)

	return m.fundPrivateRelayer(ctx, privKey1)
}

func (m *stagingManager) HostNetwork() netconf.Network {
	return m.network
}

func (m *stagingManager) DockerNetwork() netconf.Network {
	return dockerNetwork(m.network)
}

func (m *stagingManager) RelayerKey() (*ecdsa.PrivateKey, error) {
	return m.relayerKey, nil
}

func (m *stagingManager) Portals() map[uint64]Portal {
	return m.portals
}

func (m *stagingManager) AdditionalService() []string {
	return additionalServices(m.network)
}

func (m *stagingManager) relayerAddr() common.Address {
	return ethcrypto.PubkeyToAddress(m.relayerKey.PublicKey)
}

func (m *stagingManager) fundPrivateRelayer(ctx context.Context, fromPrivKey *ecdsa.PrivateKey) error {
	fromAddr := ethcrypto.PubkeyToAddress(fromPrivKey.PublicKey)

	for _, chain := range m.network.Chains {
		if publicChains[chain.Name] {
			continue // We use relayer key for public chain, it should already be funded.
		}

		ethCl := m.portals[chain.ID].Client

		nonce, err := ethCl.PendingNonceAt(ctx, fromAddr)
		if err != nil {
			return errors.Wrap(err, "get nonce")
		}

		price, err := ethCl.SuggestGasPrice(ctx)
		if err != nil {
			return errors.Wrap(err, "get gas price")
		}

		to := m.relayerAddr()

		txData := types.LegacyTx{
			Nonce:    nonce,
			GasPrice: price,
			Gas:      100_000, // 100k is fine
			To:       &to,
			Value:    new(big.Int).Mul(big.NewInt(10), big.NewInt(params.Ether)), // 10 ETH
		}

		signer := types.LatestSignerForChainID(big.NewInt(int64(chain.ID)))
		tx, err := types.SignNewTx(fromPrivKey, signer, &txData)
		if err != nil {
			return errors.Wrap(err, "sign tx")
		}

		if err := ethCl.SendTransaction(ctx, tx); err != nil {
			return errors.Wrap(err, "send tx")
		}
	}

	return nil
}
