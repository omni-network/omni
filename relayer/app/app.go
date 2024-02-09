package relayer

import (
	"context"
	"fmt"

	cprovider "github.com/omni-network/omni/lib/cchain/provider"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/gitinfo"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	xprovider "github.com/omni-network/omni/lib/xchain/provider"

	"github.com/cometbft/cometbft/rpc/client"
	"github.com/cometbft/cometbft/rpc/client/http"

	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func Run(ctx context.Context, cfg Config) error {
	log.Info(ctx, "Starting relayer")

	commit, timestamp := gitinfo.Get()
	log.Info(ctx, "Version info", "git_commit", commit, "git_timestamp", timestamp)

	network, err := netconf.Load(cfg.NetworkFile)
	if err != nil {
		return err
	}

	rpcClientPerChain, err := initializeRPCClients(network.Chains)
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

	cprov := cprovider.NewABCIProvider(tmClient)
	xprov := xprovider.New(network, rpcClientPerChain)

	for _, destChain := range network.Chains {
		sendProvider := func() (SendFunc, error) {
			sender, err := NewOpSender(ctx, destChain, rpcClientPerChain[destChain.ID], *privateKey,
				network.ChainNamesByIDs())
			if err != nil {
				return nil, err
			}

			return sender.SendTransaction, nil
		}

		worker := NewWorker(destChain, network,
			cprov,
			xprov,
			CreateSubmissions,
			sendProvider)

		go worker.Run(ctx)
	}

	startMonitoring(ctx, network, xprov, ethcrypto.PubkeyToAddress(privateKey.PublicKey), rpcClientPerChain)

	<-ctx.Done()
	log.Info(ctx, "Shutdown detected, stopping...")

	return nil
}

func newClient(tmNodeAddr string) (client.Client, error) {
	c, err := http.New(fmt.Sprintf("tcp://%s", tmNodeAddr), "/websocket")
	if err != nil {
		return nil, errors.Wrap(err, "new tendermint client")
	}

	return c, nil
}

func initializeRPCClients(chains []netconf.Chain) (map[uint64]*ethclient.Client, error) {
	rpcClientPerChain := make(map[uint64]*ethclient.Client)
	for _, chain := range chains {
		c, err := ethclient.Dial(chain.RPCURL)
		if err != nil {
			return nil, errors.Wrap(err, "dial rpc", "chain_id", chain.ID, "rpc_url", chain.RPCURL)
		}
		rpcClientPerChain[chain.ID] = c
	}

	return rpcClientPerChain, nil
}
