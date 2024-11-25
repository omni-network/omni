package feeoraclev2

import (
	"context"
	"math/big"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/tokens"
	"github.com/omni-network/omni/monitor/xfeemngr/gasprice"

	"github.com/ethereum/go-ethereum/params"
)

// feeParams returns the fee parameters for the given destination chains.
func feeParams(ctx context.Context, destChainIDs []uint64, backends ethbackend.Backends,
) ([]bindings.IFeeOracleV2FeeParams, error) {
	var resp []bindings.IFeeOracleV2FeeParams
	for _, destChainID := range destChainIDs {
		destChain, ok := evmchain.MetadataByID(destChainID)
		if !ok {
			return nil, errors.New("meta by chain id", "dest_chain", destChain.Name)
		}

		ps, err := destFeeParams(ctx, destChain, backends)
		if err != nil {
			return nil, err
		}

		resp = append(resp, ps)
	}

	return resp, nil
}

// destFeeParams returns the fee parameters for the given destination chain.
func destFeeParams(ctx context.Context, destChain evmchain.Metadata, backends ethbackend.Backends,
) (bindings.IFeeOracleV2FeeParams, error) {
	gasToken, ok := destChain.NativeToken.GasTokenID()
	if !ok {
		return bindings.IFeeOracleV2FeeParams{}, errors.New("dest chain gas token", "dest_chain", destChain.Name, "token", destChain.NativeToken)
	}

	// Get execution gas price, defaulting to 1 Gwei if any error occurs.
	var gasPrice *big.Int
	backend, err := backends.Backend(destChain.ChainID)
	if err != nil {
		log.Warn(ctx, "Failed getting exec backend, using default 1 Gwei", err, "dest_chain", destChain.Name)
		gasPrice = big.NewInt(params.GWei)
	} else {
		gasPrice, err = backend.SuggestGasPrice(ctx)
		if err != nil {
			log.Warn(ctx, "Failed fetching exec gas price, using default 1 Gwei", err, "dest_chain", destChain.Name)
			gasPrice = big.NewInt(params.GWei)
		}
	}

	// Use the chain's own ID if PostsTo is not set (i.e., is 0)
	chainForDataCost := destChain.ChainID
	if destChain.PostsTo != 0 {
		chainForDataCost = destChain.PostsTo
	}

	dataCostID, ok := evmchain.DataCostID(chainForDataCost)
	if !ok {
		return bindings.IFeeOracleV2FeeParams{}, errors.New("data cost id", "dest_chain", destChain.Name, "posts_to", destChain.PostsTo)
	}

	return bindings.IFeeOracleV2FeeParams{
		GasToken:     gasToken,
		BaseGasLimit: 100_000,
		ChainId:      destChain.ChainID,
		GasPrice:     gasprice.Tier(gasPrice.Uint64()),
		DataCostId:   dataCostID,
	}, nil
}

// dataCostParams returns the data cost parameters for the given destination chains.
func dataCostParams(ctx context.Context, destChainIDs []uint64, backends ethbackend.Backends,
) ([]bindings.IFeeOracleV2DataCostParams, error) {
	var resp []bindings.IFeeOracleV2DataCostParams
	for _, destChainID := range destChainIDs {
		destChain, ok := evmchain.MetadataByID(destChainID)
		if !ok {
			return nil, errors.New("meta by chain id", "dest_chain", destChain.Name)
		}

		ps, err := destDataCostParams(ctx, destChain, backends)
		if err != nil {
			return nil, err
		}

		resp = append(resp, ps)
	}

	return resp, nil
}

// destDataCostParams returns the data cost parameters for the given destination chain.
func destDataCostParams(ctx context.Context, destChain evmchain.Metadata, backends ethbackend.Backends,
) (bindings.IFeeOracleV2DataCostParams, error) {
	// Use the chain's own ID if PostsTo is not set (i.e., is 0)
	chainForDataCost := destChain.ChainID
	if destChain.PostsTo != 0 {
		chainForDataCost = destChain.PostsTo
	}

	postsToMetadata, ok := evmchain.MetadataByID(chainForDataCost)
	if !ok {
		return bindings.IFeeOracleV2DataCostParams{}, errors.New("posts to metadata", "dest_chain", destChain.Name, "posts_to", destChain.PostsTo)
	}

	gasToken, ok := postsToMetadata.NativeToken.GasTokenID()
	if !ok {
		return bindings.IFeeOracleV2DataCostParams{}, errors.New("posts to gas token", "posts_to", destChain.PostsTo, "token", postsToMetadata.NativeToken)
	}

	// Get data cost gas price, defaulting to 1 Gwei if any error occurs.
	var gasPrice *big.Int
	backend, err := backends.Backend(chainForDataCost)
	if err != nil {
		log.Warn(ctx, "Failed getting data cost backend, using default 1 Gwei", err, "dest_chain", destChain.Name, "posts_to", destChain.PostsTo)
		gasPrice = big.NewInt(params.GWei)
	} else {
		gasPrice, err = backend.SuggestGasPrice(ctx)
		if err != nil {
			log.Warn(ctx, "Failed fetching data cost gas price, using default 1 Gwei", err, "dest_chain", destChain.Name, "posts_to", destChain.PostsTo)
			gasPrice = big.NewInt(params.GWei)
		}
	}

	dataCostID, ok := evmchain.DataCostID(chainForDataCost)
	if !ok {
		return bindings.IFeeOracleV2DataCostParams{}, errors.New("data cost id", "dest_chain", destChain.Name, "posts_to", destChain.PostsTo)
	}

	return bindings.IFeeOracleV2DataCostParams{
		GasToken:       gasToken,
		BaseDataBuffer: 100,
		DataCostId:     dataCostID,
		GasPrice:       gasprice.Tier(gasPrice.Uint64()),
		GasPerByte:     16,
	}, nil
}

// nativeRateParams returns the native rate parameters for the given source chain.
func nativeRateParams(ctx context.Context, pricer tokens.Pricer, srcChainID uint64) ([]bindings.IFeeOracleV2NativeRateParams, error) {
	// used cached pricer, to avoid multiple price requests for same token
	pricer = tokens.NewCachedPricer(pricer)

	srcChain, ok := evmchain.MetadataByID(srcChainID)
	if !ok {
		return nil, errors.New("meta by chain id", "chain_id", srcChainID)
	}

	var resp []bindings.IFeeOracleV2NativeRateParams
	for token := range tokens.GasTokenIDs() {
		ps, err := destNativeRateParams(ctx, pricer, srcChain, token)
		if err != nil {
			return nil, err
		}

		resp = append(resp, ps)
	}

	return resp, nil
}

// destNativeRateParams returns the native rate parameters for the given source chain and destination token.
func destNativeRateParams(ctx context.Context, pricer tokens.Pricer, srcChain evmchain.Metadata, destToken tokens.Token,
) (bindings.IFeeOracleV2NativeRateParams, error) {
	// conversion rate from "dest token" to "src token"
	// ex if dest chain is ETH, and src chain is OMNI, we need to know the rate of ETH to OMNI.
	toNativeRate, err := conversionRate(ctx, pricer, srcChain.NativeToken, destToken)
	if err != nil {
		if destToken == srcChain.NativeToken {
			toNativeRate = 1 // 1 ETH = 1 ETH || 1 OMNI = 1 OMNI
		} else if destToken == tokens.OMNI {
			toNativeRate = 0.0025 // 1 OMNI = 0.0025 ETH
		} else {
			toNativeRate = 400 // 1 ETH = 400 OMNI
		}
		log.Warn(ctx, "Failed fetching conversion rate, using default", err, "src_chain", srcChain.Name, "src_token", srcChain.NativeToken, "dest_token", destToken, "to_native_rate", toNativeRate)
	}

	gasTokenID, ok := destToken.GasTokenID()
	if !ok {
		return bindings.IFeeOracleV2NativeRateParams{}, errors.New("dest token gas token id", "dest_token", destToken)
	}

	return bindings.IFeeOracleV2NativeRateParams{
		GasToken:   gasTokenID,
		NativeRate: rateToNumerator(toNativeRate),
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

	has := func(t tokens.Token) bool {
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

// conversionRateDenom matches the CONVERSION_RATE_DENOM on the FeeOracleV2 contract.
// This denominator helps convert between token amounts in solidity, in which there are no floating point numbers.
//
//	ex. (amt A) * (rate R) / CONVERSION_RATE_DENOM = (amt B)
var conversionRateDenom = big.NewInt(1_000_000)

// rateToNumerator translates a float rate (ex 0.1) to numerator / CONVERSION_RATE_DENOM (ex 100_000).
// This rate-as-numerator representation is used in FeeOracleV2 contracts.
func rateToNumerator(r float64) *big.Int {
	denom := new(big.Float).SetInt(conversionRateDenom)
	numer := new(big.Float).SetFloat64(r)
	norm, _ := new(big.Float).Mul(numer, denom).Int(nil)

	return norm
}
