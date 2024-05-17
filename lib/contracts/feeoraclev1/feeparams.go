package feeoraclev1

import (
	"context"
	"math/big"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/tokens"
	"github.com/omni-network/omni/monitor/xfeemngr"

	"github.com/ethereum/go-ethereum/params"
)

func feeParams(ctx context.Context, srcChainID uint64, destChainIDs []uint64, backends ethbackend.Backends, pricer tokens.Pricer,
) ([]bindings.IFeeOracleV1ChainFeeParams, error) {
	params := make([]bindings.IFeeOracleV1ChainFeeParams, len(destChainIDs))

	// used cached pricer, to avoid multiple price requests for same token
	pricer = tokens.NewCachedPricer(pricer)

	srcChain, ok := evmchain.MetadataByID(srcChainID)
	if !ok {
		return nil, errors.New("meta by chain id", "chain_id", srcChainID)
	}

	for i, destChainID := range destChainIDs {
		ps, err := destFeeParams(ctx, srcChain, destChainID, backends, pricer)
		if err != nil {
			return nil, err
		}

		params[i] = ps
	}

	return params, nil
}

// feeParams returns the fee parameters for the given source token and destination chains.
func destFeeParams(ctx context.Context, srcChain evmchain.Metadata, destChainID uint64, backends ethbackend.Backends, pricer tokens.Pricer,
) (bindings.IFeeOracleV1ChainFeeParams, error) {
	destChain, ok := evmchain.MetadataByID(destChainID)
	if !ok {
		return bindings.IFeeOracleV1ChainFeeParams{}, errors.New("meta by chain id", "chain_id", destChainID)
	}

	backend, err := backends.Backend(destChainID)
	if err != nil {
		return bindings.IFeeOracleV1ChainFeeParams{}, errors.Wrap(err, "get backend", "chain_id", destChainID)
	}

	// conversion rate from "dest token" to "src token"
	// ex if dest chain is ETH, and src chain is OMNI, we need to know the rate of ETH to OMNI.
	toNativeRate, err := conversionRate(ctx, pricer, destChain.NativeToken, srcChain.NativeToken)
	if err != nil {
		return bindings.IFeeOracleV1ChainFeeParams{}, err
	}

	gasPrice, err := backend.SuggestGasPrice(ctx)
	if err != nil {
		log.Warn(ctx, "Failed fetching gas price, using default 1 Gwei", err, "chain_id", destChainID)
		gasPrice = big.NewInt(params.GWei)
	}

	return bindings.IFeeOracleV1ChainFeeParams{
		ChainId:      destChainID,
		ToNativeRate: rateToNumerator(toNativeRate),
		GasPrice:     withGasPriceShield(gasPrice),
	}, nil
}

// conversionRate returns the conversion rate C from token F to token T, where C = price(F) / price(T).
// Ex. We want to convert from ETH to OMNI. We need to know the what X OMNI = 1 ETH.
// If the price of OMNI is 10, the price of ETH is 1000. The conversion rate C is price(ETH) / price(OMNI) = 1000 / 10 = 100.
func conversionRate(ctx context.Context, pricer tokens.Pricer, from, to tokens.Token) (float64, error) {
	if from == to {
		return 1, nil
	}

	prices, err := pricer.Price(ctx, from, to)
	if err != nil {
		return 0, errors.Wrap(err, "get price", "ids", "from", from, "to", to)
	}

	return prices[from] / prices[to], nil
}

// conversionRateDenom matches the CONVERSION_RATE_DENOM on the FeeOracleV1 contract.
// This denominator helps convert between token amounts in solidity, in which there are no floating point numbers.
//
//	ex. (amt A) * (rate R) / CONVERSION_RATE_DENOM = (amt B)
var conversionRateDenom = big.NewInt(1_000_000)

// rateToNumerator translates a float rate (ex 0.1) to numerator / CONVERSION_RATE_DENOM (ex 100_000).
// This rate-as-numerator representation is used in FeeOracleV1 contracts.
func rateToNumerator(r float64) *big.Int {
	denom := new(big.Float).SetInt64(conversionRateDenom.Int64())
	numer := new(big.Float).SetFloat64(r)
	norm, _ := new(big.Float).Mul(numer, denom).Int(nil)

	return norm
}

// withGasPriceShield returns the gas price with an added xfeemngr.GasPriceShield pct offset.
func withGasPriceShield(gasPrice *big.Int) *big.Int {
	gasPriceF := float64(gasPrice.Uint64())
	return new(big.Int).SetUint64(uint64(gasPriceF + (xfeemngr.GasPriceShield * gasPriceF)))
}
