package provider

import (
	"context"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/cchain"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

//nolint:gochecknoglobals // Static ABI types
var portalABI = mustGetABI(bindings.OmniPortalMetaData)

// XBlock returns the valsync XBlock at the given height, or false if not available, or an error.
// The height is equivalent to the validator set id.
func (p Provider) XBlock(ctx context.Context, height uint64) (xchain.Block, bool, error) {
	vals, ok, err := p.ValidatorSet(ctx, height)
	if err != nil {
		return xchain.Block{}, false, err
	} else if !ok {
		return xchain.Block{}, false, nil
	}

	chainID, err := p.chainID(ctx)
	if err != nil {
		return xchain.Block{}, false, errors.Wrap(err, "get chain ID")
	}

	portalVals, err := toPortalVals(vals)
	if err != nil {
		return xchain.Block{}, false, errors.Wrap(err, "convert validators")
	}

	data, err := portalABI.Pack("addValidatorSet", height, portalVals)
	if err != nil {
		return xchain.Block{}, false, errors.Wrap(err, "pack validators")
	}

	// Return a mostly stubbed xchain.Block with the encoded validators.
	return xchain.Block{
		BlockHeader: xchain.BlockHeader{
			SourceChainID: chainID,
			BlockHeight:   height,
		},
		Msgs: []xchain.Msg{{
			MsgID: xchain.MsgID{
				StreamID: xchain.StreamID{
					SourceChainID: chainID,
				},
				StreamOffset: height,
			},
			Data: data,
		}},
	}, true, nil
}

// mustGetABI returns the metadata's ABI as an abi.ABI type.
// It panics on error.
func mustGetABI(metadata *bind.MetaData) *abi.ABI {
	abi, err := metadata.GetAbi()
	if err != nil {
		panic(err)
	}

	return abi
}

// toPortalVals converts a slice of cchain.Validator to a slice of bindings.Validator.
func toPortalVals(vals []cchain.Validator) ([]bindings.Validator, error) {
	resp := make([]bindings.Validator, 0, len(vals))
	for _, val := range vals {
		if err := val.Verify(); err != nil {
			return nil, err
		}

		resp = append(resp, bindings.Validator{
			Addr:  val.Address,
			Power: uint64(val.Power),
		})
	}

	return resp, nil
}
