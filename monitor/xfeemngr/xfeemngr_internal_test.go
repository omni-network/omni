package xfeemngr

import (
	"context"
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/tokens"
	"github.com/omni-network/omni/monitor/xfeemngr/contract"
	"github.com/omni-network/omni/monitor/xfeemngr/gasprice"
	"github.com/omni-network/omni/monitor/xfeemngr/ticker"
	"github.com/omni-network/omni/monitor/xfeemngr/tokenprice"

	"github.com/ethereum/go-ethereum"

	"github.com/stretchr/testify/require"
)

// TestStart is a length test that tests the xfeemngr.Manager lifecycle.
//
// It mocks the following external resources:
//   - gas pricers
//   - token pricer
//   - fee oracle contracts
//   - tickers
//
// The xfeemngr.Manager's job is to stream gas and token prices into their respective buffers,
// and to sync on chain FeeOracleV1 state with buffer values.
//
// The test has the following structure:
//   - Setup mocks, with initial gas and token pricers.
//   - Start the manager
//   - Tick the ticker, to potentially trigger buffer & on chain updates
//   - Assert on chain values are expected
func TestStart(t *testing.T) {
	t.Parallel()

	chainIDs := []uint64{1, 2, 3, 4, 5}

	// mock gas prices / pricers
	initialGasPrices := makeGasPrices(chainIDs)
	gasPricers := makeMockGasPricers(initialGasPrices)

	// mock token prices / pricer
	initialTokenPrices := map[tokens.Token]float64{
		tokens.OMNI: randTokenPrice(tokens.OMNI),
		tokens.ETH:  randTokenPrice(tokens.ETH),
	}

	tokenPricer := tokens.NewMockPricer(initialTokenPrices)

	// helper to get price of a token from the pricer
	priceOf := func(token tokens.Token) float64 {
		t.Helper()
		prices, err := tokenPricer.Price(context.Background(), token)
		require.NoError(t, err)

		price, ok := prices[token]
		require.True(t, ok)

		return price
	}

	tick := ticker.NewMock()

	gpriceThreshold := 0.1
	tpriceThreshold := 0.1

	gpriceBuf := gasprice.NewBuffer(toEthGasPricers(gasPricers), gasprice.WithThresholdPct(gpriceThreshold), gasprice.WithTicker(tick))
	tpriceBuf := tokenprice.NewBuffer(tokenPricer, tokens.OMNI, tokens.ETH, tokenprice.WithThresholdPct(tpriceThreshold), tokenprice.WithTicker(tick))

	chains := makeChains(chainIDs)
	oracles := makeMockOracles(chains, tick, gpriceBuf, tpriceBuf)

	// make sure all postsTo are zero, at start
	// this way when we check them later, we know they have been updated
	for _, oracle := range oracles {
		for _, dest := range oracle.toSync {
			postsTo, err := mustGetContract(t, oracle).PostsTo(context.Background(), dest.ChainID)
			require.NoError(t, err)
			require.Equal(t, uint64(0), postsTo)
		}
	}

	ctx := context.Background()

	mngr := Manager{
		gprice:  gpriceBuf,
		tprice:  tpriceBuf,
		oracles: oracles,
	}

	// start the manager
	mngr.start(ctx)

	// single tick should move fill gas price and token buffers
	// values should be read from buffers, and set on chain
	tick.Tick()

	// expect initial values to be set on chain
	expectInitials := func() {
		for _, oracle := range oracles {
			src := oracle.chain

			for _, dest := range oracle.toSync {
				// check gas price
				c, err := oracle.getContract(ctx)
				require.NoError(t, err)
				gasprice, err := c.GasPriceOn(ctx, dest.ChainID)
				require.NoError(t, err)
				require.Equal(t, withGasPriceShield(initialGasPrices[dest.ChainID]), gasprice.Uint64(), "initial gas price")

				// check to native rate
				// expect rate is float dest token per src token
				expectedRate := initialTokenPrices[dest.NativeToken] / initialTokenPrices[src.NativeToken]

				// numerator is the rate * conversionRateDenom, this is expected value on chain
				expectedNumer := rateToNumerator(expectedRate)

				if src.NativeToken == dest.NativeToken {
					//nolint:testifylint // 1:1 rate is expected
					require.Equal(t, float64(1), expectedRate, "expect 1:1 rate for same tokens")
				}

				onChainNumer, err := mustGetContract(t, oracle).ToNativeRate(ctx, dest.ChainID)
				require.NoError(t, err)
				require.Equal(t, expectedNumer.Uint64(), onChainNumer.Uint64(), "initial conversion rate")
			}
		}
	}

	expectInitials()

	// increase gas prices, but not above threshold
	for _, mock := range gasPricers {
		mock.SetPrice(mock.Price() + uint64(float64(mock.Price())*gpriceThreshold/2))
	}

	// increase token prices, but not above threshold
	for token, price := range initialTokenPrices {
		tokenPricer.SetPrice(token, price+(price*tpriceThreshold/2))
	}

	tick.Tick()

	// expect no change on chain
	expectInitials()

	// increase gas prices above threshold
	for _, mock := range gasPricers {
		mock.SetPrice(mock.Price() + uint64(float64(mock.Price())*gpriceThreshold)*2)
	}

	// increase token prices above threshold
	for token, price := range initialTokenPrices {
		tokenPricer.SetPrice(token, price+(price*tpriceThreshold)*2)
	}

	tick.Tick()

	// expect on chain gas prices and conversion rates to be updated
	for _, oracle := range oracles {
		src := oracle.chain

		for _, dest := range oracle.toSync {
			// check gas price
			gasprice, err := mustGetContract(t, oracle).GasPriceOn(ctx, dest.ChainID)
			require.NoError(t, err)
			require.Equal(t, withGasPriceShield(gasPricers[dest.ChainID].Price()), gasprice.Uint64(), "updated gas price")

			// check to native rate
			// expect rate is float dest token per src token
			expectedRate := priceOf(dest.NativeToken) / priceOf(src.NativeToken)

			// numerator is the rate * conversionRateDenom, this is expected value on chain
			expectedNumer := rateToNumerator(expectedRate)

			if src.NativeToken == dest.NativeToken {
				//nolint:testifylint // 1:1 rate is expected
				require.Equal(t, float64(1), expectedRate, "expect 1:1 rate for same tokens")
			}

			onChainNumer, err := mustGetContract(t, oracle).ToNativeRate(ctx, dest.ChainID)

			require.NoError(t, err)
			require.Equal(t, expectedNumer.Uint64(), onChainNumer.Uint64(), "updated conversion rate")
		}
	}

	// set gas prices back to initial
	for chainID, price := range initialGasPrices {
		gasPricers[chainID].SetPrice(price)
	}

	// set token prices back to initial
	for token, price := range initialTokenPrices {
		tokenPricer.SetPrice(token, price)
	}

	tick.Tick()

	// make sure all postsTo have been corrected
	for _, oracle := range oracles {
		for _, dest := range oracle.toSync {
			postsTo, err := mustGetContract(t, oracle).PostsTo(ctx, dest.ChainID)
			require.NoError(t, err)
			require.Equal(t, dest.PostsTo, postsTo)
		}
	}

	// expect on chain values to be back to initials
	expectInitials()
}

func makeMockOracles(chains []evmchain.Metadata, tick ticker.Ticker, gprice *gasprice.Buffer, tprice *tokenprice.Buffer) map[uint64]feeOracle {
	oracles := make(map[uint64]feeOracle)

	for _, chain := range chains {
		mockContract := contract.NewMockFeeOracleV1()
		getContract := func(context.Context) (contract.FeeOracleV1, error) { return mockContract, nil }

		oracles[chain.ChainID] = feeOracle{
			chain:       chain,
			tick:        tick,
			toSync:      chains,
			gprice:      gprice,
			tprice:      tprice,
			getContract: getContract,
		}
	}

	return oracles
}

func mustGetContract(t *testing.T, oracle feeOracle) contract.FeeOracleV1 {
	t.Helper()
	c, err := oracle.getContract(context.Background())
	require.NoError(t, err)

	return c
}

// makeChains generates a list of mock chains from a list of chainIDs.
func makeChains(chainIDs []uint64) []evmchain.Metadata {
	chains := make([]evmchain.Metadata, 0, len(chainIDs)-1)

	for i, chainID := range chainIDs {
		// use ETH for all chains, except the first
		token := tokens.ETH
		if i == 0 {
			token = tokens.OMNI
		}

		// every even chian "postsTo" the last chain
		// every other chain "postsTo" itself
		postsTo := chainID
		if chainID%2 == 0 {
			postsTo = chainIDs[len(chainIDs)-1]
		}

		meta := evmchain.Metadata{
			ChainID:     chainID,
			Name:        "test-chain-" + fmt.Sprint(chainID),
			BlockPeriod: time.Second,
			NativeToken: token,
			PostsTo:     postsTo,
		}

		chains = append(chains, meta)
	}

	return chains
}

// makeGasPrices generates a map chainID -> gas price for each chainID.
func makeGasPrices(chainIDs []uint64) map[uint64]uint64 {
	prices := make(map[uint64]uint64)

	for _, chainID := range chainIDs {
		prices[chainID] = randGasPrice()
	}

	return prices
}

// makeMockGasPricers generates a map of mock gas pricers for n chains.
func makeMockGasPricers(prices map[uint64]uint64) map[uint64]*gasprice.MockPricer {
	mocks := make(map[uint64]*gasprice.MockPricer)
	for chainID, price := range prices {
		mocks[chainID] = gasprice.NewMockPricer(price)
	}

	return mocks
}

// toEthGasPricers  transform map[uint64]*gasprice.MockPricer to map[uint64]ethereum.GasPricer (interface).
func toEthGasPricers(mocks map[uint64]*gasprice.MockPricer) map[uint64]ethereum.GasPricer {
	pricers := make(map[uint64]ethereum.GasPricer)
	for chainID, mock := range mocks {
		pricers[chainID] = mock
	}

	return pricers
}

// randGasPrice generates a random, reasonable gas price.
func randGasPrice() uint64 {
	oneGwei := 1_000_000_000 // i gwei
	return uint64(rand.Float64() * float64(oneGwei))
}

// randTokenPrice generates a random, reasonable token price.
func randTokenPrice(token tokens.Token) float64 {
	// discriminate between ETH and other tokens (OMNI)
	// so that test omni-per-eth conversion rates do not exceed maxOmniPerEth

	if token == tokens.ETH {
		return float64(rand.Intn(500)) + rand.Float64() + 3000
	}

	return float64(rand.Intn(50)) + rand.Float64()
}
