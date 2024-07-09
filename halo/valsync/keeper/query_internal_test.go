package keeper

import (
	"testing"

	types1 "github.com/omni-network/omni/halo/attest/types"
	ptypes "github.com/omni-network/omni/halo/portal/types"
	"github.com/omni-network/omni/halo/valsync/types"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/ethereum/go-ethereum/crypto"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestValidatorSet(t *testing.T) {
	t.Parallel()

	set1 := newSet(t, 1, 2)
	set2 := newSet(t, 1, 2, 3)
	set3 := newSet(t, 3, 4)

	sets12 := [][]*Validator{set1, set2}
	sets123 := [][]*Validator{set1, set2, set3}

	tests := []struct {
		name        string
		populate    [][]*Validator
		expectation expectation
		req         *types.ValidatorSetRequest
		resp        *types.ValidatorSetResponse
	}{
		{
			name:        "query_set_1_of_3",
			populate:    sets123,
			expectation: defaultExpectation(),
			req:         newReqID(1),
			resp:        newResp(1, set1),
		},
		{
			name:        "query_set_3_of_3",
			populate:    sets123,
			expectation: defaultExpectation(),
			req:         newReqID(3),
			resp:        newResp(3, set3),
		},
		{
			name:        "query_latest_of_2",
			populate:    sets12,
			expectation: defaultExpectation(),
			req:         newReqLatest(),
			resp:        newResp(2, set2),
		},
		{
			name:        "query_latest_of_3",
			populate:    sets123,
			expectation: defaultExpectation(),
			req:         newReqLatest(),
			resp:        newResp(3, set3),
		},
		{
			name:        "query_approved_2_of_2",
			populate:    sets12,
			expectation: approvedExpectation(),
			req:         newReqLatest(),
			resp:        newActivatedResp(2, set2),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			keeper, sdkCtx := setupKeeper(t, tt.expectation)

			for i, set := range tt.populate {
				sdkCtx = sdkCtx.WithBlockHeight(int64(i + 1))
				_, err := keeper.insertValidatorSet(sdkCtx, clone(set), i == 0)
				require.NoError(t, err)
			}

			_, err := keeper.processAttested(sdkCtx)
			require.NoError(t, err)

			resp, err := keeper.ValidatorSet(sdkCtx, tt.req)
			require.NoError(t, err)

			if !cmp.Equal(resp, tt.resp, cmpOpts) {
				t.Error(cmp.Diff(resp, tt.resp, cmpOpts))
			}
		})
	}
}

func approvedExpectation() expectation {
	return func(ctx sdk.Context, m mocks) {
		blockOffset := uint64(99)

		m.portal.EXPECT().EmitMsg(
			gomock.Any(),
			ptypes.MsgTypeValSet,
			gomock.Any(),
			xchain.BroadcastChainID,
			xchain.ShardBroadcast0,
		).AnyTimes().
			Return(blockOffset, nil)

		m.aKeeper.EXPECT().ListAttestationsFrom(
			gomock.Any(),
			netconf.Simnet.Static().OmniConsensusChainIDUint64(),
			uint32(xchain.ConfFinalized),
			blockOffset,
			uint64(1), // Only query for 1 attestation.
		).AnyTimes().
			Return([]*types1.Attestation{{}}, nil)

		m.subscriber.EXPECT().UpdateValidatorSet(gomock.Any()).Do(
			func(set *types.ValidatorSetResponse) {
				if set.Id == 0 {
					panic("validator set id is zero")
				}
				if set.ActivatedHeight < set.CreatedHeight {
					panic("validator set activated height is less than created")
				}
				if len(set.Validators) == 0 {
					panic("validator set is empty")
				}
				for _, val := range set.Validators {
					if val.Power == 0 {
						panic("zero power validator in active set")
					}
				}
			},
		).AnyTimes()
	}
}

func defaultExpectation() expectation {
	return func(ctx sdk.Context, m mocks) {
		blockOffset := uint64(99)

		m.portal.EXPECT().EmitMsg(
			gomock.Any(),
			ptypes.MsgTypeValSet,
			gomock.Any(),
			xchain.BroadcastChainID,
			xchain.ShardBroadcast0,
		).AnyTimes().
			Return(blockOffset, nil)

		m.aKeeper.EXPECT().ListAttestationsFrom(
			gomock.Any(),
			netconf.Simnet.Static().OmniConsensusChainIDUint64(),
			uint32(xchain.ConfFinalized),
			blockOffset,
			uint64(1), // Only query for 1 attestation.
		).AnyTimes().
			Return(nil, nil)
	}
}

func newReqID(id uint64) *types.ValidatorSetRequest {
	return &types.ValidatorSetRequest{
		Id: id,
	}
}

func newReqLatest() *types.ValidatorSetRequest {
	return &types.ValidatorSetRequest{
		Id:     0,
		Latest: true,
	}
}

func newActivatedResp(id uint64, set []*Validator) *types.ValidatorSetResponse {
	resp := newResp(id, set)
	resp.ActivatedHeight = id + 2

	return resp
}

func newResp(id uint64, set []*Validator) *types.ValidatorSetResponse {
	var vals []*types.Validator
	for _, v := range set {
		vals = append(vals, &types.Validator{
			ConsensusPubkey: v.GetPubKey(),
			Power:           v.GetPower(),
		})
	}

	return &types.ValidatorSetResponse{
		Id:              id,
		CreatedHeight:   id,
		ActivatedHeight: 0,
		Validators:      vals,
	}
}

func newSet(t *testing.T, pubkeys ...int) []*Validator {
	t.Helper()

	var resp []*Validator
	for _, pubkey := range pubkeys {
		var pk [32]byte
		pk[0] = byte(pubkey)
		priv, err := crypto.ToECDSA(pk[:])
		require.NoError(t, err)

		resp = append(resp, &Validator{
			PubKey: crypto.CompressPubkey(&priv.PublicKey),
			Power:  1,
		})
	}

	return resp
}

func clone(set []*Validator) []*Validator {
	var resp []*Validator
	for _, v := range set {
		resp = append(resp, &Validator{
			PubKey: v.GetPubKey(),
			Power:  v.GetPower(),
		})
	}

	return resp
}

var cmpOpts = cmp.Options{cmpopts.IgnoreUnexported(
	Validator{},
	types.Validator{},
	types.ValidatorSetRequest{},
	types.ValidatorSetResponse{},
)}
