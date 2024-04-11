package relayer

import (
	"context"
	"fmt"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/buildinfo"
	cprovider "github.com/omni-network/omni/lib/cchain/provider"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	xprovider "github.com/omni-network/omni/lib/xchain/provider"

	"github.com/cometbft/cometbft/rpc/client"
	"github.com/cometbft/cometbft/rpc/client/http"

	ethcrypto "github.com/ethereum/go-ethereum/crypto"
)

func Run(ctx context.Context, cfg Config) error {
	log.Info(ctx, "Starting relayer")

	buildinfo.Instrument(ctx)

	network, err := netconf.Load(cfg.NetworkFile)
	if err != nil {
		return err
	}

	rpcClientPerChain, err := initializeRPCClients(network.EVMChains())
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

	cprov := cprovider.NewABCIProvider(tmClient, network.ID, network.ChainNamesByIDs())
	xprov := xprovider.New(network, rpcClientPerChain, cprov)

	state, ok, err := LoadCursors(cfg.StateFile)
	if err != nil {
		return err
	} else if !ok {
		state = NewEmptyState(cfg.StateFile)
	}

	for _, destChain := range network.EVMChains() {
		// Setup sender
		sendProvider := func() (SendFunc, error) {
			sender, err := NewSender(destChain, rpcClientPerChain[destChain.ID], *privateKey,
				network.ChainNamesByIDs())
			if err != nil {
				return nil, err
			}

			return sender.SendTransaction, nil
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
			state,
			awaitValSet)

		go worker.Run(ctx)
	}

	startMonitoring(ctx, network, xprov, cprov, ethcrypto.PubkeyToAddress(privateKey.PublicKey), rpcClientPerChain)

	select {
	case <-ctx.Done():
		log.Info(ctx, "Shutdown detected, stopping...")
		return nil
	case err := <-serveMonitoring(cfg.MonitoringAddr):
		return err
	}
}

func newClient(tmNodeAddr string) (client.Client, error) {
	c, err := http.New(fmt.Sprintf("tcp://%s", tmNodeAddr), "/websocket")
	if err != nil {
		return nil, errors.Wrap(err, "new tendermint client")
	}

	return c, nil
}

func initializeRPCClients(chains []netconf.Chain) (map[uint64]ethclient.Client, error) {
	rpcClientPerChain := make(map[uint64]ethclient.Client)
	for _, chain := range chains {
		c, err := ethclient.Dial(chain.Name, chain.RPCURL)
		if err != nil {
			return nil, errors.Wrap(err, "dial rpc", "chain_id", chain.ID, "rpc_url", chain.RPCURL)
		}
		rpcClientPerChain[chain.ID] = c
	}

	return rpcClientPerChain, nil
}
