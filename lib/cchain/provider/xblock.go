package provider

import (
	"context"

	"github.com/omni-network/omni/contracts/bindings"
	ptypes "github.com/omni-network/omni/halo/portal/types"
	"github.com/omni-network/omni/lib/cchain"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

//nolint:gochecknoglobals // Static ABI types
var portalABI = mustGetABI(bindings.OmniPortalMetaData)

// XBlock returns the consensus XBlock at the given height/offset or latest, or false if not available, or an error.
func (p Provider) XBlock(ctx context.Context, height uint64, latest bool) (xchain.Block, bool, error) {
	block, ok, err := p.portalBlock(ctx, height, latest)
	if err != nil {
		return xchain.Block{}, false, err
	} else if !ok {
		return xchain.Block{}, false, nil
	} else if !latest && block.Id != height {
		return xchain.Block{}, false, errors.New("unexpected block height [BUG]")
	}

	chainID, err := p.chainID(ctx)
	if err != nil {
		return xchain.Block{}, false, errors.Wrap(err, "get chain ID")
	}

	var msgs []xchain.Msg
	for _, msg := range block.Msgs {
		switch ptypes.MsgType(msg.Type) {
		case ptypes.MsgTypeValSet:
			valset, ok, err := p.valset(ctx, msg.MsgTypeId, false)
			if err != nil {
				return xchain.Block{}, false, errors.Wrap(err, "get valset")
			} else if !ok {
				return xchain.Block{}, false, errors.New("unexpected valset not found [BUG]")
			}

			portalVals, err := toPortalVals(valset.Validators)
			if err != nil {
				return xchain.Block{}, false, errors.Wrap(err, "convert validators")
			}

			data, err := portalABI.Pack("addValidatorSet", valset.ValSetID, portalVals)
			if err != nil {
				return xchain.Block{}, false, errors.Wrap(err, "pack validators")
			}

			msgs = append(msgs, xchain.Msg{
				MsgID: xchain.MsgID{
					StreamID: xchain.StreamID{
						SourceChainID: chainID,
						DestChainID:   msg.DestChainId,
						ShardID:       xchain.ShardID(msg.ShardId),
					},
					StreamOffset: msg.StreamOffset,
				},
				Data: data,
			})
		default:
			return xchain.Block{}, false, errors.New("unexpected msg type [BUG]")
		}
	}

	// Return a mostly stubbed xchain.Block with the encoded validators.
	return xchain.Block{
		BlockHeader: xchain.BlockHeader{
			SourceChainID: chainID,
			ConfLevel:     xchain.ConfFinalized, // Hardcode ConfLevel for now.
			BlockOffset:   block.Id,
			BlockHeight:   block.Id,
		},
		Msgs: msgs,
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
