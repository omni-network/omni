package cctpgen

import (
	"context"
	"math/big"
	"time"

	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/cctp"
	cctpdb "github.com/omni-network/omni/lib/cctp/db"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/tokens"
	"github.com/omni-network/omni/lib/tokens/tokenutil"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	cosmosdb "github.com/cosmos/cosmos-db"
)

// Start starts bridging usdc via cctp. Metrics managed by lib/cctp.
func Start(
	ctx context.Context,
	network netconf.Network,
	clients map[uint64]ethclient.Client,
	privKeyPath string,
	dbDir string,
) error {
	if network.ID == netconf.Devnet {
		// Devnet is not supported.
		return nil
	}

	log.Info(ctx, "Starting cctpgen")

	network = trimNetwork(network)

	privKey, err := crypto.LoadECDSA(privKeyPath)
	if err != nil {
		return errors.Wrap(err, "load private key")
	}

	sender := crypto.PubkeyToAddress(privKey.PublicKey)
	recipient := sender
	minter := sender

	backends, err := ethbackend.BackendsFromClients(clients, privKey)
	if err != nil {
		return errors.Wrap(err, "create backends")
	}

	db, err := newDB(dbDir)
	if err != nil {
		return errors.Wrap(err, "create db")
	}

	err = cctp.MintAuditForever(ctx, db, newCCTPClient(network.ID), network, backends, recipient, minter)
	if err != nil {
		return errors.Wrap(err, "mint audit forever")
	}

	go doSendsForever(ctx, db, network.ID, backends, sender)

	return nil
}

func newCCTPClient(networkID netconf.ID) cctp.Client {
	api := cctp.TestnetAPI
	if networkID == netconf.Mainnet {
		api = cctp.MainnetAPI
	}

	return cctp.NewClient(api)
}

// newDB returns a new CCTP DB instance based on the given directory.
func newDB(dbDir string) (*cctpdb.DB, error) {
	if dbDir == "" {
		memDB := cosmosdb.NewMemDB()
		return cctpdb.New(memDB)
	}

	var err error
	lvlDB, err := cosmosdb.NewGoLevelDB("cctpgen", dbDir, nil)
	if err != nil {
		return nil, errors.Wrap(err, "new golevel db")
	}

	return cctpdb.New(lvlDB)
}

// doSendsForever continuously bridges USDC between chains.
func doSendsForever(ctx context.Context, db *cctpdb.DB, networkID netconf.ID, backends ethbackend.Backends, bridger common.Address) {
	interval := 30 * time.Minute
	retryInterval := 1 * time.Minute

	timer := time.NewTimer(0)
	defer timer.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-timer.C:
			timer.Reset(interval)

			err := doSendsOnce(ctx, db, networkID, backends, bridger)
			if err != nil {
				log.Warn(ctx, "CCTP sends failed (will retry)", err)
				timer.Reset(retryInterval)
			}
		}
	}
}

func doSendsOnce(
	ctx context.Context,
	db *cctpdb.DB,
	networkID netconf.ID,
	backends ethbackend.Backends,
	bridger common.Address,
) error {
	sends := getSends(networkID)

	// make sure we have enough balance for all sends
	for _, send := range sends {
		backend, err := backends.Backend(send.SrcChain)
		if err != nil {
			return errors.Wrap(err, "backend")
		}

		balance, err := tokenutil.BalanceOfAsset(ctx, backend, tokens.USDC, bridger)
		if err != nil {
			return errors.Wrap(err, "balance of")
		}

		if bi.LT(balance, send.Amount) {
			return errors.New("insufficient balance")
		}
	}

	// do sends
	for _, send := range sends {
		backend, err := backends.Backend(send.SrcChain)
		if err != nil {
			return errors.Wrap(err, "get backend")
		}

		_, err = cctp.SendUSDC(ctx, db, networkID, backend, cctp.SendUSDCArgs{
			Sender:      bridger,
			Recipient:   bridger,
			SrcChainID:  send.SrcChain,
			DestChainID: send.DestChain,
			Amount:      send.Amount,
		})
		if err != nil {
			return errors.Wrap(err, "send usdc")
		}
	}

	return nil
}

// Send represents a single send operation.
type Send struct {
	SrcChain  uint64
	DestChain uint64
	Amount    *big.Int
}

// getSends returns list of sends based on the network ID.
func getSends(networkID netconf.ID) []Send {
	if networkID == netconf.Mainnet {
		return []Send{
			{evmchain.IDArbitrumOne, evmchain.IDBase, bi.Dec6(1)},     // Arbitrum One -> Base
			{evmchain.IDBase, evmchain.IDEthereum, bi.Dec6(1)},        // Base -> Ethereum
			{evmchain.IDEthereum, evmchain.IDOptimism, bi.Dec6(1)},    // Ethereum -> Optimism
			{evmchain.IDOptimism, evmchain.IDArbitrumOne, bi.Dec6(1)}, // Optimism -> Arbitrum One
		}
	}

	// omega / staging
	return []Send{
		{evmchain.IDArbSepolia, evmchain.IDBaseSepolia, bi.Dec6(1)}, // Arbitrum Sepolia -> Base Sepolia
		{evmchain.IDBaseSepolia, evmchain.IDOpSepolia, bi.Dec6(1)},  // Base Sepolia -> Optimism Sepolia
		{evmchain.IDOpSepolia, evmchain.IDArbSepolia, bi.Dec6(1)},   // Optimism Sepolia -> Arbitrum Sepolia
	}
}

// trimNetwork to only include chains with CCTP support.
func trimNetwork(network netconf.Network) netconf.Network {
	trimmed := netconf.Network{ID: network.ID}

	for _, chain := range network.Chains {
		if cctp.IsSupportedChain(chain.ID) {
			trimmed.Chains = append(trimmed.Chains, chain)
		}
	}

	return trimmed
}
