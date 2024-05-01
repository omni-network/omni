package xfeemngr

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"
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

func makeOracle(ctx context.Context, chain netconf.Chain, network netconf.Network, endpoints xchain.RPCEndpoints,
	pk *ecdsa.PrivateKey, gprice *gasprice.Buffer, tprice *tokenprice.Buffer) (feeOracle, error) {
	chainmeta, ok := evmchain.MetadataByID(chain.ID)
	if !ok {
		return feeOracle{}, errors.New("chain metadata not found", "chain", chain.ID)
	}

	destChains, err := makeDestChains(chain.ID, network)
	if err != nil {
		return feeOracle{}, errors.Wrap(err, "make dest chains")
	}

	bound, err := contract.New(ctx, chain, endpoints, pk)
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
	buffered := o.gprice.GasPrice(dest.ChainID)

	if buffered == 0 {
		return nil
	}

	onChain, err := o.contract.GasPriceOn(ctx, dest.ChainID)
	if err != nil {
		return errors.Wrap(err, "gas price on")
	}

	// if on chain matches buffered, return
	if onChain.Uint64() == buffered {
		return nil
	}

	err = o.contract.SetGasPriceOn(ctx, dest.ChainID, new(big.Int).SetUint64(buffered))
	if err != nil {
		return errors.Wrap(err, "set gas price on")
	}

	log.Debug(ctx, "Updated gas price", "destChainID", dest.ChainID, "old", onChain, "new", buffered)

	return nil
}

// syncToNativeRate sets the on-chain conversion rate to the buffered conversion rate, if they differ.
func (o feeOracle) syncToNativeRate(ctx context.Context, dest evmchain.Metadata) error {
	srcPrice := o.tprice.Price(o.chain.NativeToken)
	destPrice := o.tprice.Price(dest.NativeToken)

	if srcPrice == 0 || destPrice == 0 {
		return nil
	}

	buffered := rateFromPrices(destPrice, srcPrice)

	// if native tokens are the same, we should only ever have a 1:1 conversion rate
	// TODO: warn if on-chain or buffered rate is not 1:1, when it should be
	if o.chain.NativeToken == dest.NativeToken {
		buffered.Set(conversionRateDenom)
	}

	onChain, err := o.contract.ToNativeRate(ctx, dest.ChainID)
	if err != nil {
		return errors.Wrap(err, "conversion rate on")
	}

	// compare on chain and buffered rates within epsilon
	// use epsilon 1000, because the conversion rate is normalized by 1_000_000
	// this gets use 1000 / 1000_000 = 0.001 precision
	epsilon := big.NewInt(1000)
	if inEpsilon(onChain, buffered, epsilon) {
		return nil
	}

	err = o.contract.SetToNativeRate(ctx, dest.ChainID, buffered)
	if err != nil {
		return errors.Wrap(err, "set to native rate")
	}

	log.Debug(ctx, "Updated to-native rate", "destChainID", dest.ChainID, "old", onChain, "new", buffered)

	return nil
}

// makeDestChains generates a list of destination chains, excluding the source chain.
func makeDestChains(srcChainID uint64, network netconf.Network) ([]evmchain.Metadata, error) {
	destChains := make([]evmchain.Metadata, len(network.Chains)-1)
	for i, chain := range network.Chains {
		if chain.ID == srcChainID {
			continue
		}

		meta, ok := evmchain.MetadataByID(chain.ID)
		if !ok {
			return nil, errors.New("chain metadata not found", "chain", chain.ID)
		}

		destChains[i] = meta
	}

	return destChains, nil
}

// conversionRateDenom matches the CONVERSION_RATE_DENOM on the FeeOracleV1 contract.
// This denominator helps convert between token amounts in solidity, in which there are no floating point numbers.
//
//	ex. (amt A) * (rate R) / CONVERSION_RATE_DENOM = (amt B)
var conversionRateDenom = big.NewInt(1_000_000)

// rateFromPrices returns the rate R such that Y FROM * R = X TO, such that Y
// and X have the same dollar value. R is normalized by conversionRateDenom.
func rateFromPrices(fromPrice, toPrice float64) *big.Int {
	r := new(big.Float).Quo(big.NewFloat(fromPrice), big.NewFloat(toPrice))
	denom := new(big.Float).SetInt64(conversionRateDenom.Int64())

	n, _ := r.Mul(r, denom).Int(nil)

	return n
}

// inEpsilon returns true if a and b are within epsilon of each other.
func inEpsilon(a, b, epsilon *big.Int) bool {
	// if a - b < epsilon, then a == b
	return new(big.Int).Sub(a, b).CmpAbs(epsilon) == 0
}
