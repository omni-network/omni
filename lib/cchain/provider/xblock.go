package provider

import (
	"context"

	"github.com/omni-network/omni/contracts/bindings"
	ptypes "github.com/omni-network/omni/halo/portal/types"
	"github.com/omni-network/omni/halo/registry/types"
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

	dataProviders := map[ptypes.MsgType]func(ctx context.Context, msg *ptypes.Msg) ([]byte, error){
		ptypes.MsgTypeValSet:  p.msgValSetData,
		ptypes.MsgTypeNetwork: p.msgNetworkData,
	}

	var msgs []xchain.Msg
	for _, msg := range block.Msgs {
		if msg.ShardID() == xchain.ShardBroadcast0 && msg.StreamOffset == 1 && msg.MsgType() != ptypes.MsgTypeValSet {
			return xchain.Block{}, false, errors.New("initial broadcast message not genesis valset [BUG]", "type", msg.MsgType())
		}

		dataProvider, ok := dataProviders[msg.MsgType()]
		if !ok {
			return xchain.Block{}, false, errors.New("unexpected msg type [BUG]", "type", msg.MsgType())
		}

		data, err := dataProvider(ctx, msg)
		if err != nil {
			return xchain.Block{}, false, errors.Wrap(err, "get msg data", "type", msg.MsgType())
		}

		msgs = append(msgs, xchain.Msg{
			MsgID: xchain.MsgID{
				StreamID: xchain.StreamID{
					SourceChainID: chainID,
					DestChainID:   msg.DestChainId,
					ShardID:       msg.ShardID(),
				},
				StreamOffset: msg.StreamOffset,
			},
			Data: data,
		})
	}

	// Return a mostly stubbed xchain.Block with the encoded validators.
	return xchain.Block{
		BlockHeader: xchain.BlockHeader{
			ChainID:     chainID,
			BlockHeight: block.Id,
		},
		Msgs: msgs,
	}, true, nil
}

func (p Provider) msgValSetData(ctx context.Context, msg *ptypes.Msg) ([]byte, error) {
	valset, ok, err := p.valset(ctx, msg.MsgTypeId, false)
	if err != nil {
		return nil, errors.Wrap(err, "get valset")
	} else if !ok {
		return nil, errors.New("unexpected valset not found [BUG]")
	}

	portalVals, err := toPortalVals(valset.Validators)
	if err != nil {
		return nil, errors.Wrap(err, "convert validators")
	}

	data, err := portalABI.Pack("addValidatorSet", valset.ValSetID, portalVals)
	if err != nil {
		return nil, errors.Wrap(err, "pack validators")
	}

	return data, nil
}

func (p Provider) msgNetworkData(ctx context.Context, msg *ptypes.Msg) ([]byte, error) {
	network, ok, err := p.networkFunc(ctx, msg.MsgTypeId, false)
	if err != nil {
		return nil, errors.Wrap(err, "get network")
	} else if !ok {
		return nil, errors.New("unexpected network not found [BUG]")
	}

	data, err := portalABI.Pack("setNetwork", toPortalChains(network.Portals))
	if err != nil {
		return nil, errors.Wrap(err, "pack validators")
	}

	return data, nil
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

func toPortalChains(portals []*types.Portal) []bindings.XTypesChain {
	resp := make([]bindings.XTypesChain, 0, len(portals))
	for _, portal := range portals {
		resp = append(resp, bindings.XTypesChain{
			ChainId: portal.GetChainId(),
			Shards:  portal.GetShardIds(),
		})
	}

	return resp
}
