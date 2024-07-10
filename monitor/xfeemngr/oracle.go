package xfeemngr

import (
	"context"
	"crypto/ecdsa"
	"math/big"

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
	contract contract.FeeOracleV1
	chain    evmchain.Metadata   // source chain
	dests    []evmchain.Metadata // destination chains
	gprice   *gasprice.Buffer    // gas price buffer
	tprice   *tokenprice.Buffer  // token price buffer
}

func makeOracle(ctx context.Context, chain netconf.Chain, network netconf.Network, ethClients map[uint64]ethclient.Client,
	pk *ecdsa.PrivateKey, gprice *gasprice.Buffer, tprice *tokenprice.Buffer) (feeOracle, error) {
	chainmeta, ok := evmchain.MetadataByID(chain.ID)
	if !ok {
		return feeOracle{}, errors.New("chain metadata not found", "chain", chain.ID)
	}

	destChains, err := makeDestChains(chain.ID, network)
	if err != nil {
		return feeOracle{}, errors.Wrap(err, "make dest chains")
	}

	ethCl, ok := ethClients[chain.ID]
	if !ok {
		return feeOracle{}, errors.New("eth client not found", "chain", chain.ID)
	}

	bound, err := contract.New(ctx, chain, ethCl, pk)
	if err != nil {
		return feeOracle{}, errors.Wrap(err, "new bound fee oracle")
	}

	return feeOracle{
		chain:    chainmeta,
		dests:    destChains,
		contract: bound,
		gprice:   gprice,
		tprice:   tprice,
	}, nil
}

// syncForever syncs the on-chain gas price and token conversion rates with their respective buffers, forever.
func (o feeOracle) syncForever(ctx context.Context, tick ticker.Ticker) {
	ctx = log.WithCtx(ctx, "component", "feeOracle", "chain", o.chain.Name)
	log.Info(ctx, "Starting fee oracle sync")
	tick.Go(ctx, o.syncOnce)
}

// syncOnce syncs the on-chain gas price and token conversion rates with their respective buffers, once.
func (o feeOracle) syncOnce(ctx context.Context) {
	ctx = log.WithCtx(ctx, "srcChainID", o.chain.ChainID, "srcToken", o.chain.NativeToken)

	for _, dest := range o.dests {
		err := o.syncGasPrice(ctx, dest)
		if err != nil {
			log.Error(ctx, "Failed to sync gas price", err, "destChainID", dest.ChainID)
		}

		err = o.syncToNativeRate(ctx, dest)
		if err != nil {
			log.Error(ctx, "Failed to sync conversion rate", err, "destChainID", dest.ChainID)
		}
	}
}

// syncGasPrice sets the on-chain gas price to the buffered gas price, if they differ.
func (o feeOracle) syncGasPrice(ctx context.Context, dest evmchain.Metadata) error {
	ctx = log.WithCtx(ctx, "destChainID", dest.ChainID)

	buffered := o.gprice.GasPrice(dest.ChainID)

	if buffered == 0 {
		return nil
	}

	if buffered > maxSaneGasPrice {
		log.Warn(ctx, "Buffered gas price exceeds sane max", errors.New("unexpected gas price"), "buffered", buffered, "maxSane", maxSaneGasPrice)
		buffered = maxSaneGasPrice
	}

	onChain, err := o.contract.GasPriceOn(ctx, dest.ChainID)
	if err != nil {
		return errors.Wrap(err, "gas price on")
	}

	guageGasPrice(o.chain, dest, onChain.Uint64())

	shielded := withGasPriceShield(buffered)

	// if on chain gas price is within epsilon of buffered + GasPriceShield, do nothing
	// The shield helps keep on-chain gas prices higher than live gas prices
	if inEpsilon(float64(onChain.Uint64()), float64(shielded), 0.001) {
		return nil
	}

	err = o.contract.SetGasPriceOn(ctx, dest.ChainID, new(big.Int).SetUint64(shielded))
	if err != nil {
		return errors.Wrap(err, "set gas price on")
	}

	// if on chain update successful, update gauge
	guageGasPrice(o.chain, dest, buffered)

	return nil
}

// guageGasPrice updates the gas price gauge for the given chain.
func guageGasPrice(src, dest evmchain.Metadata, price uint64) {
	onChainGasPrice.WithLabelValues(src.Name, dest.Name).Set(float64(price))
}

// syncToNativeRate sets the on-chain conversion rate to the buffered conversion rate, if they differ.
func (o feeOracle) syncToNativeRate(ctx context.Context, dest evmchain.Metadata) error {
	ctx = log.WithCtx(ctx, "destChainID", dest.ChainID, "destToken", dest.NativeToken)

	srcPrice := o.tprice.Price(o.chain.NativeToken)
	destPrice := o.tprice.Price(dest.NativeToken)

	if srcPrice == 0 || destPrice == 0 {
		return nil
	}

	// bufferedRate "source token per destination token" is "USD per dest" / "USD per src"
	bufferedRate := destPrice / srcPrice

	if o.chain.NativeToken == tokens.OMNI && dest.NativeToken == tokens.ETH && bufferedRate > maxSaneOmniPerEth {
		log.Warn(ctx, "Buffered omni-per-eth exceeds sane max", errors.New("unexpected conversion rate"), "buffered", bufferedRate, "maxSane", maxSaneOmniPerEth)
		bufferedRate = maxSaneOmniPerEth
	}

	if o.chain.NativeToken == tokens.ETH && dest.NativeToken == tokens.OMNI && bufferedRate > maxSaneEthPerOmni {
		log.Warn(ctx, "Buffered eth-per-omni exceeds sane max", errors.New("unexpected conversion rate"), "buffered", bufferedRate, "maxSane", maxSaneEthPerOmni)
		bufferedRate = maxSaneEthPerOmni
	}

	bufferedNumer := rateToNumerator(bufferedRate)

	onChainNumer, err := o.contract.ToNativeRate(ctx, dest.ChainID)
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

	err = o.contract.SetToNativeRate(ctx, dest.ChainID, bufferedNumer)
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

// makeDestChains generates a list of destination chains, excluding the source chain.
func makeDestChains(srcChainID uint64, network netconf.Network) ([]evmchain.Metadata, error) {
	chains := network.EVMChains()
	destChains := make([]evmchain.Metadata, 0, len(chains)-1)

	var foundSrc bool
	for _, chain := range chains {
		if chain.ID == srcChainID {
			foundSrc = true
			continue
		}

		meta, ok := evmchain.MetadataByID(chain.ID)
		if !ok {
			return nil, errors.New("chain metadata not found", "chain", chain.ID)
		}

		destChains = append(destChains, meta)
	}

	if !foundSrc {
		return nil, errors.New("source chain not in network", "chain", srcChainID)
	}

	return destChains, nil
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
