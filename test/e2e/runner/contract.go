package main

import (
	"context"
	"math/big"
	"strings"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	// anvilPrivKeyHex of pre-funded anvil account 0.
	anvilPrivKeyHex = "0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
)

var (
	// defaultNetworkis the  devnet network configuration.
	//nolint:gochecknoglobals // This is predefined at this point.
	defaultNetwork = netconf.Network{
		Name: netconf.Devnet,
		Chains: []netconf.Chain{
			{
				ID:            100,
				Name:          "chain_a",
				RPCURL:        "http://localhost:6545",
				PortalAddress: "0x5FbDB2315678afecb367f032d93F642f64180aa3",
			},
		},
	}
)

// defaultServices returns the default additional docker-compose services to start.
func defaultServices() []string {
	resp := make([]string, 0, len(defaultNetwork.Chains))
	for _, chain := range defaultNetwork.Chains {
		resp = append(resp, chain.Name)
	}

	return resp
}

func DeployContracts(ctx context.Context) error {
	log.Info(ctx, "Deploying portal contracts")

	for _, chain := range defaultNetwork.Chains {
		ethClient, err := ethclient.Dial(chain.RPCURL)
		if err != nil {
			return errors.Wrap(err, "dial chain")
		}

		txOpts, err := newTxOpts(anvilPrivKeyHex, chain.ID)
		if err != nil {
			return err
		}

		addr, _, _, err := bindings.DeployOmniPortal(txOpts, ethClient)
		if err != nil {
			return errors.Wrap(err, "deploy portal")
		} else if addr.Hex() != chain.PortalAddress {
			return errors.New("portal address mismatch")
		}
	}

	return nil
}

func newTxOpts(privKeyHex string, chainID uint64) (*bind.TransactOpts, error) {
	pk, err := crypto.HexToECDSA(strings.TrimPrefix(privKeyHex, "0x"))
	if err != nil {
		return nil, errors.Wrap(err, "parse private key")
	}

	txOpts, err := bind.NewKeyedTransactorWithChainID(
		pk,
		big.NewInt(int64(chainID)),
	)
	if err != nil {
		return nil, errors.Wrap(err, "keyed tx ops")
	}

	return txOpts, nil
}
