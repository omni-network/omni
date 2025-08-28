package feeoraclev1

import (
	"context"
	"math/big"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/tokenpricer"
	"github.com/omni-network/omni/lib/tokens"
	"github.com/omni-network/omni/monitor/xfeemngr/gasprice"
)

func feeParams(ctx context.Context, srcChainID uint64, destChainIDs []uint64, backends ethbackend.Backends, pricer tokenpricer.Pricer,
) ([]bindings.IFeeOracleV1ChainFeeParams, error) {
	// used cached pricer, to avoid multiple price requests for same token
	pricer = tokenpricer.NewCached(pricer)

	srcChain, ok := evmchain.MetadataByID(srcChainID)
	if !ok {
		return nil, errors.New("meta by chain id", "chain_id", srcChainID)
	}

	var resp []bindings.IFeeOracleV1ChainFeeParams
	for _, destChainID := range destChainIDs {
		destChain, ok := evmchain.MetadataByID(destChainID)
		if !ok {
			return nil, errors.New("meta by chain id", "dest_chain", destChain.Name)
		}

		ps, err := destFeeParams(ctx, srcChain, destChain, backends, pricer)
		if err != nil {
			return nil, err
		}

		resp = append(resp, ps)

		// Add postsTo rates if not already present
		if destChain.PostsTo != 0 && !contains(resp, destChain.PostsTo) {
			resp = append(resp, bindings.IFeeOracleV1ChainFeeParams{
				ChainId:      destChain.PostsTo,
				PostsTo:      destChain.PostsTo,
				GasPrice:     bi.Gwei(1),
				ToNativeRate: rateToNumerator(1), // Stub native rates for now
			})
		}
	}

	return resp, nil
}

// feeParams returns the fee parameters for the given source token and destination chains.
func destFeeParams(ctx context.Context, srcChain evmchain.Metadata, destChain evmchain.Metadata, backends ethbackend.Backends, pricer tokenpricer.Pricer,
) (bindings.IFeeOracleV1ChainFeeParams, error) {
	backend, err := backends.Backend(destChain.ChainID)
	if err != nil {
		return bindings.IFeeOracleV1ChainFeeParams{}, errors.Wrap(err, "get backend", "dest_chain", destChain.Name)
	}

	// conversion rate from "dest token" to "src token"
	// ex if dest chain is ETH, and src chain is NOM, we need to know the rate of ETH to NOM.
	toNativeRate, err := conversionRate(ctx, pricer, destChain.NativeToken, srcChain.NativeToken)
	if err != nil {
		if srcChain.NativeToken == destChain.NativeToken {
			toNativeRate = 1 // 1 ETH = 1 ETH || 1 NOM = 1 NOM
		} else if destChain.NativeToken == tokens.NOM {
			toNativeRate = 0.000033 // 1 NOM = 0.000033 ETH
		} else {
			toNativeRate = 30000 // 1 ETH = 30000 NOM
		}
		log.Warn(ctx, "Failed fetching conversion rate, using default", err, "dest_chain", destChain.Name, "src_chain", srcChain.Name, "to_native_rate", toNativeRate)
	}

	gasPrice, err := backend.SuggestGasPrice(ctx)
	if err != nil {
		log.Warn(ctx, "Failed fetching gas price, using default 1 Gwei", err, "dest_chain", destChain.Name)
		gasPrice = bi.Gwei(1)
	}

	postsTo := destChain.PostsTo

	// if not configured, chain "posts to" itself
	if postsTo == 0 {
		postsTo = destChain.ChainID
	}

	return bindings.IFeeOracleV1ChainFeeParams{
		ChainId:      destChain.ChainID,
		PostsTo:      postsTo,
		ToNativeRate: rateToNumerator(toNativeRate),
		GasPrice:     gasprice.Tier(gasPrice),
	}, nil
}

// conversionRate returns the conversion rate C from token F to token T, where C = price(F) / price(T).
// Ex. We want to convert from ETH to NOM. We need to know the what X NOM = 1 ETH.
// If the price of NOM is 10, the price of ETH is 1000. The conversion rate C is price(ETH) / price(NOM) = 1000 / 10 = 100.
func conversionRate(ctx context.Context, pricer tokenpricer.Pricer, from, to tokens.Asset) (float64, error) {
	if from == to {
		return 1, nil
	}

	prices, err := pricer.USDPrices(ctx, from, to)
	if err != nil {
		return 0, errors.Wrap(err, "get price", "from", from, "to", to)
	}

	has := func(t tokens.Asset) bool {
		p, ok := prices[t]
		return ok && p > 0
	}
	if !has(to) {
		return 0, errors.New("missing to token price", "to", to)
	} else if !has(from) {
		return 0, errors.New("missing from token price", "from", from)
	}

	return prices[from] / prices[to], nil
}

// conversionRateDenom matches the CONVERSION_RATE_DENOM on the FeeOracleV1 contract.
// This denominator helps convert between token amounts in solidity, in which there are no floating point numbers.
//
//	ex. (amt A) * (rate R) / CONVERSION_RATE_DENOM = (amt B)
var conversionRateDenom = bi.N(1_000_000)

// rateToNumerator translates a float rate (ex 0.1) to numerator / CONVERSION_RATE_DENOM (ex 100_000).
// This rate-as-numerator representation is used in FeeOracleV1 contracts.
func rateToNumerator(r float64) *big.Int {
	denom := new(big.Float).SetInt64(conversionRateDenom.Int64())
	numer := new(big.Float).SetFloat64(r)
	norm, _ := new(big.Float).Mul(numer, denom).Int(nil)

	return norm
}

func contains(params []bindings.IFeeOracleV1ChainFeeParams, chainID uint64) bool {
	for _, param := range params {
		if param.ChainId == chainID {
			return true
		}
	}

	return false
}
