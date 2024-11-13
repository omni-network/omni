package relayer

import (
	"context"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/halo/genutil/evm/predeploys"
	"github.com/omni-network/omni/lib/buildinfo"
	cprovider "github.com/omni-network/omni/lib/cchain/provider"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"
	xprovider "github.com/omni-network/omni/lib/xchain/provider"
	"github.com/omni-network/omni/relayer/app/cursor"

	"github.com/cometbft/cometbft/rpc/client"
	"github.com/cometbft/cometbft/rpc/client/http"

	"github.com/ethereum/go-ethereum/common"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"

	dbm "github.com/cosmos/cosmos-db"
)

func Run(ctx context.Context, cfg Config) error {
	log.Info(ctx, "Starting relayer")

	db, err := initializeDB(ctx, cfg)
	if err != nil {
		return err
	}

	buildinfo.Instrument(ctx)

	// Start metrics first, so app is "up"
	monitorChan := serveMonitoring(cfg.MonitoringAddr)

	portalReg, err := makePortalRegistry(cfg.Network, cfg.RPCEndpoints)
	if err != nil {
		return err
	}

	network, err := netconf.AwaitOnExecutionChain(ctx, cfg.Network, portalReg, cfg.RPCEndpoints.Keys())
	if err != nil {
		return err
	}

	rpcClientPerChain, err := initializeRPCClients(network.EVMChains(), cfg.RPCEndpoints)
	if err != nil {
		return err
	}

	privateKey, err := ethcrypto.LoadECDSA(cfg.PrivateKey)
	if err != nil {
		return errors.Wrap(err, "failed to load private key")
	}

	tmClient, err := newClient(cfg.HaloURL)
	if err != nil {
		return err
	}

	cprov := cprovider.NewABCI(tmClient, network.ID)
	xprov := xprovider.New(network, rpcClientPerChain, cprov)

	pricer := newTokenPricer(ctx)
	pnl := newPnlLogger(network.ID, pricer)

	cursors, err := cursor.NewCursors(db, xprov, cfg.ConfirmInterval)
	if err != nil {
		return err
	}
	go cursors.Monitor(ctx, network)

	for _, destChain := range network.EVMChains() {
		// Setup send provider
		sendProvider := func() (SendAsync, error) {
			sender, err := NewSender(
				network.ID,
				destChain,
				rpcClientPerChain[destChain.ID],
				*privateKey,
				network.ChainVersionNames(),
				pnl.log,
			)
			if err != nil {
				return nil, err
			}

			return sender.SendAsync, nil
		}

		// Setup validator set awaiter
		portal, err := bindings.NewOmniPortal(destChain.PortalAddress, rpcClientPerChain[destChain.ID])
		if err != nil {
			return errors.Wrap(err, "create portal contract")
		}
		awaitValSet := newValSetAwaiter(portal, destChain.BlockPeriod)

		// Start worker
		worker := NewWorker(
			destChain,
			network,
			cprov,
			xprov,
			CreateSubmissions,
			sendProvider,
			awaitValSet,
			cursors,
		)

		go worker.Run(ctx)
	}

	select {
	case <-ctx.Done():
		log.Info(ctx, "Shutdown detected, stopping...")
		return nil
	case err := <-monitorChan:
		return err
	}
}

func newClient(tmNodeAddr string) (client.Client, error) {
	c, err := http.New("tcp://"+tmNodeAddr, "/websocket")
	if err != nil {
		return nil, errors.Wrap(err, "new tendermint client")
	}

	return c, nil
}

func initializeRPCClients(chains []netconf.Chain, endpoints xchain.RPCEndpoints) (map[uint64]ethclient.Client, error) {
	rpcClientPerChain := make(map[uint64]ethclient.Client)
	for _, chain := range chains {
		rpc, err := endpoints.ByNameOrID(chain.Name, chain.ID)
		if err != nil {
			return nil, err
		}
		c, err := ethclient.Dial(chain.Name, rpc)
		if err != nil {
			return nil, errors.Wrap(err, "dial rpc", "chain_name", chain.Name, "chain_id", chain.ID, "rpc_url", rpc)
		}
		rpcClientPerChain[chain.ID] = c
	}

	return rpcClientPerChain, nil
}

func initializeDB(ctx context.Context, cfg Config) (dbm.DB, error) {
	var db dbm.DB
	if cfg.DBDir == "" {
		log.Warn(ctx, "No --db-dir provided, using in-memory DB", nil)
		return dbm.NewMemDB(), nil
	}
	var err error
	db, err = dbm.NewGoLevelDB("indexer", cfg.DBDir, nil)
	if err != nil {
		return nil, errors.Wrap(err, "new golevel db")
	}

	return db, nil
}

func makePortalRegistry(network netconf.ID, endpoints xchain.RPCEndpoints) (*bindings.PortalRegistry, error) {
	meta := netconf.MetadataByID(network, network.Static().OmniExecutionChainID)
	rpc, err := endpoints.ByNameOrID(meta.Name, meta.ChainID)
	if err != nil {
		return nil, err
	}

	ethCl, err := ethclient.Dial(meta.Name, rpc)
	if err != nil {
		return nil, err
	}

	resp, err := bindings.NewPortalRegistry(common.HexToAddress(predeploys.PortalRegistry), ethCl)
	if err != nil {
		return nil, errors.Wrap(err, "create portal registry")
	}

	return resp, nil
}
