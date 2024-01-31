package netman

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"slices"
	"strings"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
)

func deployPublicContracts(ctx context.Context, network netconf.Network, privKey *ecdsa.PrivateKey,
) (map[uint64]Portal, netconf.Network, error) {
	log.Info(ctx, "Deploying public portal contracts")

	resp := make(map[uint64]Portal)
	for i, chain := range network.Chains {
		if !publicChains[chain.Name] {
			continue // Only public chains are deployed here.
		} else if chain.PortalAddress != "" {
			return nil, netconf.Network{}, errors.New("public portal address already set")
		}

		portal, addr, height, err := deployContract(ctx, chain, privKey)
		if err != nil {
			return nil, netconf.Network{}, err
		}

		resp[chain.ID] = portal
		network.Chains[i].PortalAddress = addr.Hex()
		network.Chains[i].DeployHeight = height

		log.Info(ctx, "Deployed public portal contract", "chain", chain.Name, "address", addr.Hex(), "height", height)
	}

	return resp, network, nil
}

func deployPrivateContracts(ctx context.Context, network netconf.Network, privKey *ecdsa.PrivateKey,
) (map[uint64]Portal, error) {
	log.Info(ctx, "Deploying private portal contracts")

	resp := make(map[uint64]Portal)
	for _, chain := range network.Chains {
		if publicChains[chain.Name] {
			continue // Public chains are already deployed.
		}

		portal, addr, _, err := deployContract(ctx, chain, privKey)
		if err != nil {
			return nil, err
		} else if addr.Hex() != chain.PortalAddress {
			return nil, errors.New("private portal address mismatch",
				"chain", chain.Name,
				"expect", chain.PortalAddress,
				"actual", addr.Hex(),
			)
		}

		resp[chain.ID] = portal
	}

	return resp, nil
}

func deployContract(ctx context.Context, chain netconf.Chain, privKey *ecdsa.PrivateKey,
) (Portal, common.Address, uint64, error) {
	ethClient, err := ethclient.Dial(chain.RPCURL)
	if err != nil {
		return Portal{}, common.Address{}, 0, errors.Wrap(err, "dial chain")
	}

	height, err := ethClient.BlockNumber(ctx)
	if err != nil {
		return Portal{}, common.Address{}, 0, errors.Wrap(err, "get block number")
	}

	txOpts, err := newTxOpts(ctx, privKey, chain.ID)
	if err != nil {
		return Portal{}, common.Address{}, 0, err
	}

	addr, _, _, err := bindings.DeployOmniPortal(txOpts, ethClient)
	if err != nil {
		return Portal{}, common.Address{}, 0, errors.Wrap(err, "deploy portal")
	}

	contract, err := bindings.NewOmniPortal(addr, ethClient)
	if err != nil {
		return Portal{}, common.Address{}, 0, errors.Wrap(err, "create portal contract")
	}

	return Portal{
		Chain:    chain,
		Client:   ethClient,
		Contract: contract,
		txOpts:   txOpts,
	}, addr, height, nil
}

func newTxOpts(ctx context.Context, privKey *ecdsa.PrivateKey, chainID uint64) (*bind.TransactOpts, error) {
	txOpts, err := bind.NewKeyedTransactorWithChainID(privKey, big.NewInt(int64(chainID)))
	if err != nil {
		return nil, errors.Wrap(err, "keyed tx ops")
	}

	txOpts.Context = ctx

	return txOpts, nil
}

// writeNetworkConfig writes the network config (adjusted for intra-docker networking) to the given path.
func dockerNetwork(network netconf.Network) netconf.Network {
	// Clone the network since we need to change the RPC URLs for intra-docker networking.
	clone := netconf.Network{
		Name:   network.Name,
		Chains: slices.Clone(network.Chains),
	}

	for i, chain := range clone.Chains {
		if publicChains[chain.Name] {
			continue // Public chains RPCs remain unchanged.
		}
		clone.Chains[i].RPCURL = fmt.Sprintf("http://%v:8545", chain.Name)
		clone.Chains[i].AuthRPCURL = fmt.Sprintf("http://%v:8551", chain.Name)
	}

	return clone
}

func mergePortals(a, b map[uint64]Portal) map[uint64]Portal {
	for k, v := range b {
		a[k] = v
	}

	return a
}

// additionalServices returns additional (to halo) docker-compose services to start.
func additionalServices(network netconf.Network) []string {
	resp := make([]string, 0, len(network.Chains))
	for _, chain := range network.Chains {
		if publicChains[chain.Name] {
			continue // Cannot start public chains (like arb_goerli)
		}
		resp = append(resp, chain.Name)
	}

	resp = append(resp, "relayer")

	return resp
}

func logBalance(ctx context.Context, chain netconf.Chain, privkey *ecdsa.PrivateKey, name string) error {
	cl, err := ethclient.Dial(chain.RPCURL)
	if err != nil {
		return errors.Wrap(err, "dial chain")
	}

	addr := crypto.PubkeyToAddress(privkey.PublicKey)

	b, err := cl.BalanceAt(ctx, addr, nil)
	if err != nil {
		return errors.Wrap(err, "get balance")
	}

	bf, _ := b.Float64()
	bf /= params.Ether

	log.Info(ctx, "Provided public chain key balance",
		"chain", chain.Name,
		"address", addr.Hex(),
		"balance", fmt.Sprintf("%.2f", bf),
		"key_name", name,
	)

	return nil
}

func mustHexToKey(privKeyHex string) *ecdsa.PrivateKey {
	privKey, err := crypto.HexToECDSA(strings.TrimPrefix(privKeyHex, "0x"))
	if err != nil {
		panic(err)
	}

	return privKey
}
