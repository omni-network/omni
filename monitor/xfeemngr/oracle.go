package xfeemngr

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/tokens"
	"github.com/omni-network/omni/monitor/xfeemngr/contract"
	"github.com/omni-network/omni/monitor/xfeemngr/gasprice"
	"github.com/omni-network/omni/monitor/xfeemngr/ticker"
	"github.com/omni-network/omni/monitor/xfeemngr/tokenprice"
)

type feeOracle struct {
	chain  evmchain.Metadata   // source chain
	tick   ticker.Ticker       // ticker to sync fee oracle
	toSync []evmchain.Metadata // chains to sync on fee oracle
	gprice *gasprice.Buffer    // gas price buffer
	tprice *tokenprice.Buffer  // token price buffer

	getContract func(context.Context) (contract.FeeOracleV1, error)
}

func makeOracle(chain netconf.Chain, toSync []evmchain.Metadata, ethCl ethclient.Client,
	pk *ecdsa.PrivateKey, gprice *gasprice.Buffer, tprice *tokenprice.Buffer) (feeOracle, error) {
	chainmeta, ok := evmchain.MetadataByID(chain.ID)
	if !ok {
		return feeOracle{}, errors.New("chain metadata not found", "chain", chain.ID)
	}

	syncInterval := feeOracleSyncInterval
	override, ok := chainSyncOverrides[chain.ID]
	if ok {
		syncInterval = override
	}

	getContract := func() func(context.Context) (contract.FeeOracleV1, error) {
		var c *contract.BoundFeeOracleV1

		return func(ctx context.Context) (contract.FeeOracleV1, error) {
			if c != nil {
				return c, nil
			}

			bound, err := contract.New(ctx, chain, ethCl, pk)
			if err != nil {
				return nil, errors.Wrap(err, "new bound fee oracle")
			}

			c = bound

			return bound, nil
		}
	}()

	return feeOracle{
		chain:       chainmeta,
		tick:        ticker.New(ticker.WithInterval(syncInterval)),
		toSync:      toSync,
		gprice:      gprice,
		tprice:      tprice,
		getContract: getContract,
	}, nil
}

// syncForever syncs the on-chain gas price and token conversion rates with their respective buffers, forever.
func (o feeOracle) syncForever(ctx context.Context) {
	ctx = log.WithCtx(ctx, "component", "feeOracle", "chain", o.chain.Name)
	log.Info(ctx, "Starting fee oracle sync")
	o.tick.Go(ctx, o.syncOnce)
}

// syncOnce syncs the on-chain gas price and token conversion rates with their respective buffers, once.
func (o feeOracle) syncOnce(ctx context.Context) {
	ctx = log.WithCtx(ctx, "src_token", o.chain.NativeToken)

	for _, dest := range o.toSync {
		err := o.syncGasPrice(ctx, dest)
		if err != nil {
			log.Error(ctx, "Failed to sync gas price", err, "dest_chain", dest.ChainID)
		}

		err = o.syncToNativeRate(ctx, dest)
		if err != nil {
			log.Error(ctx, "Failed to sync conversion rate", err, "dest_chain", dest.ChainID)
		}

		err = o.correctPostsTo(ctx, dest)
		if err != nil {
			log.Error(ctx, "Failed to correct postsTo chain", err, "dest_chain", dest.ChainID)
		}
	}
}

// syncGasPrice sets the on-chain gas price to the buffered gas price, if they differ.
func (o feeOracle) syncGasPrice(ctx context.Context, dest evmchain.Metadata) error {
	ctx = log.WithCtx(ctx, "chainId", dest.ChainID, "chain", dest.Name)

	buffered := o.gprice.GasPrice(dest.ChainID)

	if buffered == 0 {
		return nil
	}

	if buffered > maxSaneGasPrice {
		log.Warn(ctx, "Buffered gas price exceeds sane max", errors.New("unexpected gas price"), "buffered", buffered, "max_sane", maxSaneGasPrice)
		buffered = maxSaneGasPrice
	}

	c, err := o.getContract(ctx)
	if err != nil {
		return errors.Wrap(err, "get contract")
	}

	onChain, err := c.GasPriceOn(ctx, dest.ChainID)
	if err != nil {
		return errors.Wrap(err, "gas price on")
	}

	guageGasPrice(o.chain, dest, onChain.Uint64())

	shielded := withGasPriceShield(buffered)

	log.Info(ctx, "Syncing gas price", "buffered", buffered, "shielded", shielded, "dest_chain", onChain.Uint64())
	// if on chain gas price is within epsilon of buffered + GasPriceShield, do nothing
	// The shield helps keep on-chain gas prices higher than live gas prices
	if inEpsilon(float64(onChain.Uint64()), float64(shielded), 0.001) {
		return nil
	}

	err = c.SetGasPriceOn(ctx, dest.ChainID, new(big.Int).SetUint64(shielded))
	if err != nil {
		return errors.Wrap(err, "set gas price on")
	}

	// if on chain update successful, update gauge
	guageGasPrice(o.chain, dest, buffered)

	return nil
}

func (o feeOracle) correctPostsTo(ctx context.Context, dest evmchain.Metadata) error {
	c, err := o.getContract(ctx)
	if err != nil {
		return errors.Wrap(err, "get contract")
	}

	postsTo, err := c.PostsTo(ctx, dest.ChainID)
	if err != nil {
		return errors.Wrap(err, "postsTo")
	}

	// If postsTo is correct, do nothing
	// Either metadata.PostsTo == onchain postsTo
	// Or     metadata.PostsTo == 0, then chain "postsTo" itself, and on-chain postTo should be self
	if (dest.PostsTo == postsTo) || (dest.PostsTo == 0 && postsTo == dest.ChainID) {
		return nil
	}

	// if not correct, correct via BulkSetFeeParams (there is not a single setter for postsTo)
	// use current onchain gas price and conversion rate

	gasPrice, err := c.GasPriceOn(ctx, dest.ChainID)
	if err != nil {
		return errors.Wrap(err, "gas price on")
	}

	rate, err := c.ToNativeRate(ctx, dest.ChainID)
	if err != nil {
		return errors.Wrap(err, "conversion rate on")
	}

	params := []bindings.IFeeOracleV1ChainFeeParams{
		{
			ChainId:      dest.ChainID,
			GasPrice:     gasPrice,
			ToNativeRate: rate,
			PostsTo:      dest.PostsTo,
		},
	}

	err = c.BulkSetFeeParams(ctx, params)
	if err != nil {
		return errors.Wrap(err, "bulk set fee params")
	}

	log.Info(ctx, "Corrected postsTo", "dest_chain", dest.Name, "correct", dest.PostsTo, "actual", postsTo)

	return nil
}

// guageGasPrice updates the gas price gauge for the given chain.
func guageGasPrice(src, dest evmchain.Metadata, price uint64) {
	onChainGasPrice.WithLabelValues(src.Name, dest.Name).Set(float64(price))
}

// syncToNativeRate sets the on-chain conversion rate to the buffered conversion rate, if they differ.
func (o feeOracle) syncToNativeRate(ctx context.Context, dest evmchain.Metadata) error {
	ctx = log.WithCtx(ctx, "dest_chain", dest.Name, "dest_token", dest.NativeToken)

	srcPrice := o.tprice.Price(o.chain.NativeToken)
	destPrice := o.tprice.Price(dest.NativeToken)

	if srcPrice == 0 || destPrice == 0 {
		return nil
	}

	// bufferedRate "source token per destination token" is "USD per dest" / "USD per src"
	bufferedRate := destPrice / srcPrice

	log.Info(ctx, "Syncing native token rate", "source_price", srcPrice, "destination_price", destPrice, "buffered_rate", bufferedRate)
	if o.chain.NativeToken == tokens.OMNI && dest.NativeToken == tokens.ETH && bufferedRate > maxSaneOmniPerEth {
		log.Warn(ctx, "Buffered omni-per-eth exceeds sane max", errors.New("unexpected conversion rate"), "buffered", bufferedRate, "max_sane", maxSaneOmniPerEth)
		bufferedRate = maxSaneOmniPerEth
	}

	if o.chain.NativeToken == tokens.ETH && dest.NativeToken == tokens.OMNI && bufferedRate > maxSaneEthPerOmni {
		log.Warn(ctx, "Buffered eth-per-omni exceeds sane max", errors.New("unexpected conversion rate"), "buffered", bufferedRate, "max_sane", maxSaneEthPerOmni)
		bufferedRate = maxSaneEthPerOmni
	}

	bufferedNumer := rateToNumerator(bufferedRate)

	c, err := o.getContract(ctx)
	if err != nil {
		return errors.Wrap(err, "get contract")
	}

	onChainNumer, err := c.ToNativeRate(ctx, dest.ChainID)
	if err != nil {
		return errors.Wrap(err, "conversion rate on")
	}

	onChainRate := numeratorToRate(onChainNumer)
	guageRate(o.chain, dest, onChainRate)

	// compare on chain and buffered rates within epsilon, with epsilon < 1 / rateDenom
	// such that epsilon is more precise than on chain rates
	if inEpsilon(onChainRate, bufferedRate, 1.0/float64(rateDenom*10)) {
		return nil
	}

	// if bufferred rate is less than we can represent on chain, use smallest representable rate - 1/CONVERSION_RATE_DENOM
	if bufferedRate < 1.0/float64(rateDenom) {
		log.Warn(ctx, "Buffered rate too small, setting minimum on chain", errors.New("conversion rate < min repr"), "buffered", bufferedRate)
		bufferedNumer = big.NewInt(1)
	}

	err = c.SetToNativeRate(ctx, dest.ChainID, bufferedNumer)
	if err != nil {
		return errors.Wrap(err, "set to native rate")
	}

	// if on chain update successful, update gauge
	guageRate(o.chain, dest, bufferedRate)

	return nil
}

// guageRate updates the conversion rate gauge for the given source and destination chains.
func guageRate(src, dest evmchain.Metadata, rate float64) {
	onChainConversionRate.WithLabelValues(src.Name, dest.Name, src.NativeToken.String(), dest.NativeToken.String()).Set(rate)
}

// rateDenom matches FeeOracleV1.CONVERSION_RATE_DENOM
// This denominator helps convert between token amounts in solidity, in which there are no floating point numbers.
//
//	ex. (amt A) * (rate R) / CONVERSION_RATE_DENOM = (amt B)
const rateDenom = 1_000_000

// rateToNumerator translates a float rate (ex 0.1) to numerator / CONVERSION_RATE_DENOM (ex 100_000).
// This rate-as-numerator representation is used in FeeOracleV1 contracts.
func rateToNumerator(r float64) *big.Int {
	denom := new(big.Float).SetUint64(rateDenom)
	numer := new(big.Float).SetFloat64(r)
	norm, _ := new(big.Float).Mul(numer, denom).Int(nil)

	return norm
}

// numeratorToRate translates a rate numerator / CONVERSION_RATE_DENOM to a float rate.
// It is the inverse of rateToNumerator. We use non-numerator rates in metrics and logs.
func numeratorToRate(n *big.Int) float64 {
	denom := new(big.Float).SetUint64(rateDenom)
	numer := new(big.Float).SetInt(n)
	rate, _ := new(big.Float).Quo(numer, denom).Float64()

	return rate
}

// inEpsilon returns true if a and b are within epsilon of each other.
func inEpsilon(a, b, epsilon float64) bool {
	diff := a - b

	return diff < epsilon && diff > -epsilon
}

// withGasPriceShield returns the gas price with an added GasPriceShield pct offset.
func withGasPriceShield(gasPrice uint64) uint64 {
	gasPriceF := float64(gasPrice)
	return uint64(gasPriceF + gasPriceF*GasPriceShield)
}
