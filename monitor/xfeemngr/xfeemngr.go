package xfeemngr

import (
	"context"
	"crypto/ecdsa"
	"time"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/tokens"
	"github.com/omni-network/omni/lib/tokens/coingecko"
	"github.com/omni-network/omni/lib/xchain"
	"github.com/omni-network/omni/monitor/xfeemngr/gasprice"
	"github.com/omni-network/omni/monitor/xfeemngr/ticker"
	"github.com/omni-network/omni/monitor/xfeemngr/tokenprice"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/crypto"
)

type Manager struct {
	gprice  *gasprice.Buffer
	tprice  *tokenprice.Buffer
	oracles map[uint64]feeOracle
	ticker  ticker.Ticker
}

const (
	// feeOracleSyncInterval is the interval at which fee oracles syncs buffered gas and token prices with FeeOracle deployments.
	feeOracleSyncInterval = 5 * time.Minute

	// tokenPriceBufferThreshold is the pct threshold at which a new token price is buffered.
	tokenPriceBufferThreshold = 0.1

	// gasPriceBufferThreshold is the pct threshold at which a new gas price is buffered.
	gasPriceBufferThreshold = 0.1

	// GasPriceShield is the pct offset above the buffered gas price the oracle sets on chain
	// Setting shield == buffer threshold ensures that on chain gas price is always at least as high as the live gas price.
	GasPriceShield = 0.1

	// maxSaneGasPrice is the maximum sane gas price in gwei to post for any destination chain. Set to 500 gwei.
	maxSaneGasPrice = uint64(500_000_000_000)

	// maxSaneOmniPerEth is the maximum sane conversion rate of omni to eth.
	maxSaneOmniPerEth = float64(1_000_000)

	// maxSaneEthPerOmni is the maximum sane conversion rate of eth to omni.
	maxSaneEthPerOmni = float64(1)
)

func Start(ctx context.Context, network netconf.Network, endpoints xchain.RPCEndpoints, privKeyPath string) error {
	privKey, err := crypto.LoadECDSA(privKeyPath)
	if err != nil {
		return errors.Wrap(err, "load private key")
	}

	gasPricers, err := makeGasPricers(network, endpoints)
	if err != nil {
		return err
	}

	gprice := gasprice.NewBuffer(gasPricers, gasprice.WithThresholdPct(gasPriceBufferThreshold))
	tprice := tokenprice.NewBuffer(coingecko.New(), tokens.OMNI, tokens.ETH, tokenprice.WithThresholdPct(tokenPriceBufferThreshold))

	oracles, err := makeOracles(ctx, network, endpoints, privKey, gprice, tprice)
	if err != nil {
		return err
	}

	m := Manager{
		gprice:  gprice,
		tprice:  tprice,
		oracles: oracles,
		ticker:  ticker.New(ticker.WithInterval(feeOracleSyncInterval)),
	}

	if err := startMonitoring(ctx, network, endpoints); err != nil {
		return errors.Wrap(err, "start monitoring")
	}

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
		oracle.syncForever(ctx, m.ticker)
	}
}

// makeGasPricers makes a map chainID to ethereum.GasPricer for the given network / endpoints.
// This map is required by gasprice.Buffer.
func makeGasPricers(network netconf.Network, endpoints xchain.RPCEndpoints) (map[uint64]ethereum.GasPricer, error) {
	pricers := make(map[uint64]ethereum.GasPricer)

	for _, chain := range network.EVMChains() {
		rpc, err := endpoints.ByNameOrID(chain.Name, chain.ID)
		if err != nil {
			return nil, errors.Wrap(err, "rpc endpoint")
		}

		client, err := ethclient.Dial(chain.Name, rpc)
		if err != nil {
			return nil, errors.Wrap(err, "dial client")
		}

		pricers[chain.ID] = client
	}

	return pricers, nil
}

// makeOracles makes a map chainID to feeOracle for each chain in the network.
func makeOracles(ctx context.Context, network netconf.Network, endpoints xchain.RPCEndpoints,
	pk *ecdsa.PrivateKey, gprice *gasprice.Buffer, tprice *tokenprice.Buffer) (map[uint64]feeOracle, error) {
	oracles := make(map[uint64]feeOracle)

	for _, chain := range network.EVMChains() {
		oracle, err := makeOracle(ctx, chain, network, endpoints, pk, gprice, tprice)
		if err != nil {
			return nil, errors.Wrap(err, "make oracle", "chain", chain.Name)
		}

		oracles[chain.ID] = oracle
	}

	return oracles, nil
}
