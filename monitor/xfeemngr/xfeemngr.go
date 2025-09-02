package xfeemngr

import (
	"context"
	"crypto/ecdsa"
	"time"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/tokenpricer/coingecko"
	"github.com/omni-network/omni/lib/tokens"
	"github.com/omni-network/omni/lib/xchain"
	"github.com/omni-network/omni/monitor/xfeemngr/gasprice"
	"github.com/omni-network/omni/monitor/xfeemngr/ticker"
	"github.com/omni-network/omni/monitor/xfeemngr/tokenprice"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/crypto"
)

type Manager struct {
	gprice  gasprice.Buffer
	tprice  tokenprice.Buffer
	oracles map[uint64]feeOracle
}

type Config struct {
	RPCEndpoints    xchain.RPCEndpoints
	CoinGeckoAPIKey string
}

const (
	// Interval at which to sync buffers with on-chain.
	feeOracleSyncInterval = 20 * time.Minute

	// Update token price if live price is 10% different on-chain.
	tokenPriceBufferThreshold = 0.1

	// Check live token prices every 30 seconds.
	tokenPriceBufferSyncInterval = 30 * time.Second

	// Check live gas prices every 30 seconds.
	gasPriceBufferSyncInterval = 30 * time.Second

	maxSaneGasPrice     = uint64(500_000_000_000)
	maxSaneNativePerEth = float64(1_000_000)
	maxSaneEthPerNative = float64(1)
)

var chainSyncOverrides = map[uint64]time.Duration{
	// override for ethereum mainnet, to reduce spend
	evmchain.IDEthereum: 90 * time.Minute,

	// overide on holesky too, to test overide on omega
	evmchain.IDHolesky: 90 * time.Minute,

	// override for sepolia, to reduce spend
	evmchain.IDSepolia: 90 * time.Minute,
}

func Start(
	ctx context.Context,
	network netconf.Network,
	cfg Config,
	privKeyPath string,
	ethClients map[uint64]ethclient.Client,
) error {
	log.Info(ctx, "Starting fee manager", "endpoint", cfg.RPCEndpoints)

	privKey, err := crypto.LoadECDSA(privKeyPath)
	if err != nil {
		return errors.Wrap(err, "load private key")
	}

	toSync, err := chainsToSync(network)
	if err != nil {
		return err
	}

	cgCl := coingecko.New(coingecko.WithAPIKey(cfg.CoinGeckoAPIKey))

	gprice, err := gasprice.NewBuffer(
		makeGasPricers(ethClients),
		ticker.New(gasPriceBufferSyncInterval))
	if err != nil {
		return errors.Wrap(err, "new gas price buffer")
	}

	tprice := tokenprice.NewBuffer(cgCl,
		[]tokens.Asset{tokens.NOM, tokens.ETH},
		tokenPriceBufferThreshold,
		ticker.New(tokenPriceBufferSyncInterval))

	oracles, err := makeOracles(network, toSync, ethClients, privKey, gprice, tprice)
	if err != nil {
		return err
	}

	m := Manager{
		gprice:  gprice,
		tprice:  tprice,
		oracles: oracles,
	}

	startMonitoring(ctx, network, ethClients)

	m.start(ctx)

	return nil
}

func (m *Manager) start(ctx context.Context) {
	ctx = log.WithCtx(ctx, "process", "xfeemngr")

	// stream gas and token prices into buffers
	m.gprice.Stream(ctx)
	m.tprice.Stream(ctx)

	// start oracle sync
	for _, oracle := range m.oracles {
		oracle.syncForever(ctx)
	}
}

// makeGasPricers makes a map chainID to ethereum.GasPricer for the given network / endpoints.
// This map is required by gasprice.Buffer.
func makeGasPricers(ethClients map[uint64]ethclient.Client) map[uint64]ethereum.GasPricer {
	pricers := make(map[uint64]ethereum.GasPricer)
	for chainID, ethClient := range ethClients {
		pricers[chainID] = ethClient
	}

	return pricers
}

// makeOracles makes a map chainID to feeOracle for each chain in the network.
func makeOracles(network netconf.Network, toSync []evmchain.Metadata, ethClients map[uint64]ethclient.Client,
	pk *ecdsa.PrivateKey, gprice gasprice.Buffer, tprice tokenprice.Buffer) (map[uint64]feeOracle, error) {
	oracles := make(map[uint64]feeOracle)

	for _, chain := range network.EVMChains() {
		ethCl, ok := ethClients[chain.ID]
		if !ok {
			return nil, errors.New("eth client not found", "chain", chain.ID)
		}

		oracle, err := makeOracle(chain, toSync, ethCl, pk, gprice, tprice)
		if err != nil {
			return nil, errors.Wrap(err, "make oracle", "chain", chain.Name)
		}

		oracles[chain.ID] = oracle
	}

	return oracles, nil
}

// chainsToSync returns a list of evmchain.Metadata to sync on each fee oracle.
// This includes all evm chains in the network, and their "postsTo" chains.
func chainsToSync(network netconf.Network) ([]evmchain.Metadata, error) {
	var toSync []evmchain.Metadata

	// avoid dups - some chains have same postsTo, and they may be in the network
	set := make(map[uint64]bool)

	// add all chains in network
	for _, chain := range network.EVMChains() {
		meta, ok := evmchain.MetadataByID(chain.ID)
		if !ok {
			return nil, errors.New("chain metadata not found", "chain", chain.ID)
		}

		toSync = append(toSync, meta)
		set[meta.ChainID] = true
	}

	// add all "postsTo" chains
	for _, chain := range toSync {
		if chain.PostsTo == 0 || set[chain.PostsTo] {
			continue
		}

		meta, ok := evmchain.MetadataByID(chain.PostsTo)
		if !ok {
			return nil, errors.New("chain metadata not found", "chain", meta.ChainID)
		}

		toSync = append(toSync, meta)
		set[meta.ChainID] = true
	}

	return toSync, nil
}
