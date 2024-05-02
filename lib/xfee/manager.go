// xfee package provides a fee manager that manages xchain fee parameters across multiple chains.
package xfee

import (
	"context"
	"math/big"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/tokens"
)

// conversionRateDenom matches the CONVERSION_RATE_DENOM on the FeeOracleV1 contract.
// This denominator helps convert between token amounts in solidity, in which there are no floating point numbers.
//
//	ex. (amt A) * (rate R) / CONVERSION_RATE_DENOM = (amt B)
var conversionRateDenom = big.NewInt(1_000_000)

// Manager is a  fee manager, that managers xchain fee parameters across multiple chains.
type Manager struct {
	pricer   tokens.Pricer
	chainIDs []uint64
	backends ethbackend.Backends
}

func NewManager(chainIDs []uint64, pricer tokens.Pricer, backends ethbackend.Backends) *Manager {
	return &Manager{
		pricer:   pricer,
		chainIDs: chainIDs,
		backends: backends,
	}
}

// FeeParams returns the FeeOracleV1 fee parameters for a given source chain.
// This includes parameters per each destination chain, set in Manager.chainIds, excluding the source chain.
func (m Manager) FeeParams(ctx context.Context, srcChainID uint64) ([]bindings.IFeeOracleV1ChainFeeParams, error) {
	destChainIDs := m.destChainIds(srcChainID)
	params := make([]bindings.IFeeOracleV1ChainFeeParams, len(destChainIDs))

	srcMeta, ok := evmchain.MetadataByID(srcChainID)
	if !ok {
		return nil, errors.New("meta by chain id", "chain_id", srcChainID)
	}

	srcToken := srcMeta.NativeToken

	for i, destChainID := range destChainIDs {
		ps, err := m.destFeeParams(ctx, srcToken, destChainID)
		if err != nil {
			return nil, err
		}

		params[i] = ps
	}

	return params, nil
}

// feeParams returns the fee parameters for the given source token and destination chains.
func (m Manager) destFeeParams(ctx context.Context, srcToken tokens.Token, destChainID uint64) (bindings.IFeeOracleV1ChainFeeParams, error) {
	destMeta, ok := evmchain.MetadataByID(destChainID)
	if !ok {
		return bindings.IFeeOracleV1ChainFeeParams{}, errors.New("meta by chain id", "chain_id", destChainID)
	}

	backend, err := m.backends.Backend(destChainID)
	if err != nil {
		return bindings.IFeeOracleV1ChainFeeParams{}, errors.Wrap(err, "get backend", "chain_id", destChainID)
	}

	toNativeRate, err := conversionRate(ctx, m.pricer, srcToken, destMeta.NativeToken)
	if err != nil {
		return bindings.IFeeOracleV1ChainFeeParams{}, err
	}

	gasPrice, err := backend.SuggestGasPrice(ctx)
	if err != nil {
		return bindings.IFeeOracleV1ChainFeeParams{}, errors.Wrap(err, "get gas price", "chain_id", destChainID)
	}

	return bindings.IFeeOracleV1ChainFeeParams{
		ChainId:      destChainID,
		ToNativeRate: toNativeRate,
		GasPrice:     gasPrice,
	}, nil
}

// destChainIds returns the FeeManager.chainIDs excluding the given srcChainID.
func (m Manager) destChainIds(srcChainID uint64) []uint64 {
	destChainIDs := make([]uint64, len(m.chainIDs)-1)
	for i, chainID := range m.chainIDs {
		if chainID == srcChainID {
			continue
		}

		destChainIDs[i] = chainID
	}

	return destChainIDs
}

// conversionRate returns the conversion rate C such that.
func conversionRate(ctx context.Context, pricer tokens.Pricer, from, to tokens.Token) (*big.Int, error) {
	if from == to {
		return conversionRateDenom, nil
	}

	p, err := pricer.Price(ctx, from, to)
	if err != nil {
		return nil, errors.Wrap(err, "get price", "ids", "from", from, "to", to)
	}

	return rateFromPrices(p[from], p[to]), nil
}

// rateFromPrices returns the rate R such that Y FROM * R = X TO, such that Y
// and X have the same dollar value. R is normalized by conversionRateDenom.
func rateFromPrices(fromPrice, toPrice float64) *big.Int {
	r := new(big.Float).Quo(big.NewFloat(fromPrice), big.NewFloat(toPrice))
	denom := new(big.Float).SetInt64(conversionRateDenom.Int64())

	n, _ := r.Mul(r, denom).Int(nil)

	return n
}
