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
	xprovider "github.com/omni-network/omni/lib/xchain/provider"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	cosmosdb "github.com/cosmos/cosmos-db"
)

var (
	baseSepolia     = mustMeta(evmchain.IDBaseSepolia)
	opSepolia       = mustMeta(evmchain.IDArbSepolia)
	arbitrumSepolia = mustMeta(evmchain.IDOpSepolia)
)

// Start starts bridging usdc via cctp. Metrics managed by lib/cctp.
func Start(
	ctx context.Context,
	network netconf.Network,
	clients map[uint64]ethclient.Client,
	privKeyPath string,
	dbDir string,
) error {
	if network.ID != netconf.Omega && network.ID != netconf.Staging {
		// Only run for omega and staging
		return nil
	}

	log.Info(ctx, "Starting CCTP test process")

	chains := []evmchain.Metadata{
		baseSepolia,
		opSepolia,
		arbitrumSepolia,
	}

	xprov := xprovider.New(network, clients, nil)

	privKey, err := crypto.LoadECDSA(privKeyPath)
	if err != nil {
		return errors.Wrap(err, "load private key")
	}

	sender := crypto.PubkeyToAddress(privKey.PublicKey)
	recipient := sender
	minter := sender

	cctpClient := cctp.NewClient(cctp.TestnetAPI)

	backends, err := ethbackend.BackendsFromClients(clients, privKey)
	if err != nil {
		return errors.Wrap(err, "create backends")
	}

	db, err := newDB(dbDir)
	if err != nil {
		return errors.Wrap(err, "create db")
	}

	err = cctp.MintForever(ctx, db, cctpClient, backends, chains, minter)
	if err != nil {
		return errors.Wrap(err, "mint forever")
	}

	err = cctp.AuditForever(ctx, db, network.ID, xprov, clients, chains, recipient)
	if err != nil {
		return errors.Wrap(err, "audit forever")
	}

	cctp.MonitorForever(ctx, db)

	go doSendsForever(ctx, db, network.ID, backends, sender)

	return nil
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
	sends := []struct {
		srcChain  uint64
		destChain uint64
		amount    *big.Int
	}{
		{evmchain.IDArbSepolia, evmchain.IDBaseSepolia, bi.Dec6(1)}, // Arbitrum Sepolia -> Base Sepolia
		{evmchain.IDBaseSepolia, evmchain.IDOpSepolia, bi.Dec6(1)},  // Base Sepolia -> Optimism Sepolia
		{evmchain.IDOpSepolia, evmchain.IDArbSepolia, bi.Dec6(1)},   // Optimism Sepolia -> Arbitrum Sepolia
	}

	// make sure we have enough balance for all sends
	for _, send := range sends {
		backend, err := backends.Backend(send.srcChain)
		if err != nil {
			return errors.Wrap(err, "backend")
		}

		balance, err := tokenutil.BalanceOfAsset(ctx, backend, tokens.USDC, bridger)
		if err != nil {
			return errors.Wrap(err, "balance of")
		}

		if bi.LT(balance, send.amount) {
			return errors.New("insufficient balance")
		}
	}

	// do sends
	for _, send := range sends {
		backend, err := backends.Backend(send.srcChain)
		if err != nil {
			return errors.Wrap(err, "get backend")
		}

		_, err = cctp.SendUSDC(ctx, db, networkID, backend, cctp.SendUSDCArgs{
			Sender:      bridger,
			Recipient:   bridger,
			SrcChainID:  send.srcChain,
			DestChainID: send.destChain,
			Amount:      send.amount,
		})
		if err != nil {
			return errors.Wrap(err, "send usdc")
		}
	}

	return nil
}

// mustMeta returns the metadata for a chain ID, panicking if not found.
func mustMeta(chainID uint64) evmchain.Metadata {
	meta, ok := evmchain.MetadataByID(chainID)
	if !ok {
		panic("chain not found")
	}

	return meta
}
