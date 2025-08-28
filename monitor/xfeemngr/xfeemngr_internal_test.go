package xfeemngr

import (
	"context"
	"fmt"
	"math/big"
	"math/rand"
	"testing"
	"time"

	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/tokens"
	"github.com/omni-network/omni/monitor/xfeemngr/contract"
	"github.com/omni-network/omni/monitor/xfeemngr/gasprice"
	"github.com/omni-network/omni/monitor/xfeemngr/ticker"
	"github.com/omni-network/omni/monitor/xfeemngr/tokenprice"

	"github.com/stretchr/testify/require"
)

// TestStart is a length test that tests the xfeemngr.Manager lifecycle.
// If confirms that if gas price / token price buffers update, on chain oracles are updated.
func TestStart(t *testing.T) {
	t.Parallel()

	chainIDs := []uint64{1, 2, 3, 4, 5}

	// mock token prices / pricer
	initialTokenPrices := map[tokens.Asset]float64{
		tokens.OMNI: randTokenPrice(tokens.OMNI),
		tokens.NOM:  randTokenPrice(tokens.NOM),
		tokens.ETH:  randTokenPrice(tokens.ETH),
	}

	initialGasPrices := makeGasPrices(chainIDs)

	tick := ticker.NewMock()
	gpriceBuf := gasprice.NewMockBuffer()
	tpriceBuf := tokenprice.NewMockBuffer()

	for token, price := range initialTokenPrices {
		tpriceBuf.SetPrice(token, price)
	}

	for chainID, price := range initialGasPrices {
		gpriceBuf.SetGasPrice(chainID, price)
	}

	chains := makeChains(chainIDs)
	oracles := makeMockOracles(chains, tick, gpriceBuf, tpriceBuf)

	ctx := t.Context()

	mngr := Manager{
		gprice:  gpriceBuf,
		tprice:  tpriceBuf,
		oracles: oracles,
	}

	expect := func(tprices map[tokens.Asset]float64, gprices map[uint64]*big.Int) {
		for _, oracle := range oracles {
			src := oracle.chain

			for _, dest := range oracle.toSync {
				// check gas price
				c, err := oracle.getContract(ctx)
				require.NoError(t, err)
				gasprice, err := c.GasPriceOn(ctx, dest.ChainID)
				require.NoError(t, err)
				require.Equal(t, gprices[dest.ChainID].Uint64(), gasprice.Uint64(), "gas price")

				// check toNativeRate
				rate := tprices[dest.NativeToken] / tprices[src.NativeToken]

				// handle maximum rate case
				if src.NativeToken == tokens.OMNI && src.NativeToken != dest.NativeToken && rate > maxSaneNativePerEth {
					rate = maxSaneNativePerEth
				}
				if src.NativeToken == tokens.NOM && src.NativeToken != dest.NativeToken && rate > maxSaneNativePerEth {
					rate = maxSaneNativePerEth
				}
				if src.NativeToken == tokens.ETH && src.NativeToken != dest.NativeToken && rate > maxSaneEthPerNative {
					rate = maxSaneEthPerNative
				}

				numer := rateToNumerator(rate)

				if src.NativeToken == dest.NativeToken {
					//nolint:testifylint // should be exactly 1:1
					require.Equal(t, rate, float64(1), "expect 1:1 rate for same tokens")
				}

				// handle minimum rate case
				if rate < 1.0/float64(rateDenom) {
					numer = bi.One() // Use minimum representable rate
				}

				onChainNumer, err := mustGetContract(t, oracle).ToNativeRate(ctx, dest.ChainID)
				require.NoError(t, err)

				// allow variance of +-1 due to floating point rounding errors
				numerDiff := bi.Sub(numer, onChainNumer).Int64()
				require.True(t, numerDiff >= -1 && numerDiff <= 1, "conversion rate")
			}
		}
	}

	mngr.start(ctx)

	// tick once
	tick.Tick()

	// onchain should match initial
	expect(initialTokenPrices, initialGasPrices)

	// 10 steps
	gprices := initialGasPrices
	tprices := initialTokenPrices

	for i := 0; i < 10; i++ {
		// maybe update token buffer prices
		for token := range tprices {
			if randBool() {
				tpriceBuf.SetPrice(token, randTokenPrice(token))
				tprices[token] = tpriceBuf.Price(token)
			}
		}

		// maybe update gas buffer prices
		for chainID := range gprices {
			if randBool() {
				gpriceBuf.SetGasPrice(chainID, randGasPrice())
				gprices[chainID] = gpriceBuf.GasPrice(chainID)
			}
		}

		// tick
		tick.Tick()

		// onchain should match buffer
		expect(tprices, gprices)
	}
}

func makeMockOracles(chains []evmchain.Metadata, tick ticker.Ticker, gprice gasprice.Buffer, tprice tokenprice.Buffer) map[uint64]feeOracle {
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
	c, err := oracle.getContract(t.Context())
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
			token = tokens.NOM
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
func makeGasPrices(chainIDs []uint64) map[uint64]*big.Int {
	prices := make(map[uint64]*big.Int)

	for _, chainID := range chainIDs {
		prices[chainID] = randGasPrice()
	}

	return prices
}

// randGasPrice generates a random, reasonable gas price.
func randGasPrice() *big.Int {
	oneGwei := 1_000_000_000 // i gwei
	return bi.N(rand.Intn(oneGwei))
}

// randTokenPrice generates a random, reasonable token price.
func randTokenPrice(token tokens.Asset) float64 {
	// discriminate between ETH and other tokens (OMNI/NOM)
	// so that test conversion rates do not exceed sane maxes

	if token == tokens.ETH {
		return float64(rand.Intn(500)) + rand.Float64() + 3000
	}

	// For OMNI and NOM, use reasonable ranges
	return float64(rand.Intn(50)) + rand.Float64()
}

func randBool() bool {
	return rand.Intn(2) == 0
}
