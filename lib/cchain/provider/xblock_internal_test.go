package provider

import (
	"context"
	"testing"

	ptypes "github.com/omni-network/omni/halo/portal/types"
	rtypes "github.com/omni-network/omni/halo/registry/types"
	"github.com/omni-network/omni/lib/cchain"
	"github.com/omni-network/omni/lib/tutil"

	fuzz "github.com/google/gofuzz"
	"github.com/stretchr/testify/require"
)

//go:generate go test . -golden -clean

func setupTest(t *testing.T) (uint64, func(ctx context.Context, h uint64, _ bool) (*rtypes.NetworkResponse, bool, error), valsetFunc, chainIDFunc, portalBlockFunc) {
	t.Helper()
	f := fuzz.NewWithSeed(99).NilChance(0).Funcs(
		// Fuzz valid validators.
		func(e *cchain.PortalValidator, c fuzz.Continue) {
			c.FuzzNoCustom(e)
			if e.Power < 0 {
				e.Power = -e.Power
			} else if e.Power == 0 {
				e.Power = 1
			}
		})

	var height uint64
	f.Fuzz(&height)

	valFunc := func(ctx context.Context, h uint64, _ bool) (valSetResponse, bool, error) {
		require.Equal(t, height, h)
		var resp []cchain.PortalValidator
		f.Fuzz(&resp)

		return valSetResponse{
			ValSetID:   height,
			Validators: resp,
		}, true, nil
	}
	networkFunc := func(ctx context.Context, h uint64, _ bool) (*rtypes.NetworkResponse, bool, error) {
		require.Equal(t, height, h)
		var resp rtypes.NetworkResponse
		f.Fuzz(&resp)

		return &resp, true, nil
	}
	chainFunc := func(ctx context.Context) (uint64, error) {
		return 77, nil
	}
	portalBlockFunc := func(ctx context.Context, h uint64, _ bool) (*ptypes.BlockResponse, bool, error) {
		var valSetMsg ptypes.Msg
		f.Fuzz(&valSetMsg)
		valSetMsg.Type = uint32(ptypes.MsgTypeValSet)
		valSetMsg.MsgTypeId = h

		var networkMsg ptypes.Msg
		f.Fuzz(&networkMsg)
		networkMsg.Type = uint32(ptypes.MsgTypeNetwork)
		networkMsg.MsgTypeId = h

		return &ptypes.BlockResponse{
			Id:            h,
			CreatedHeight: 123456,
			Msgs:          []ptypes.Msg{valSetMsg, networkMsg},
		}, true, nil
	}

	return height, networkFunc, valFunc, chainFunc, portalBlockFunc
}

// TestXBlock ensures we receive expected xblock response from provider.
func TestXBlock(t *testing.T) {
	t.Parallel()

	height, networkFunc, valFunc, chainFunc, portalBlockFunc := setupTest(t)
	prov := Provider{valset: valFunc, networkFunc: networkFunc, chainID: chainFunc, portalBlock: portalBlockFunc}

	block, ok, err := prov.XBlock(t.Context(), height, false)
	require.NoError(t, err)
	require.True(t, ok)
	tutil.RequireGoldenJSON(t, block)
}

// TestXBlock_MaliciousResponse ensures that the provider will return an error if any of the msgs in the msgs slice
// of the block response is nil.
func TestXBlock_MaliciousResponse(t *testing.T) {
	t.Parallel()

	portalBlockFunc := func(ctx context.Context, h uint64, _ bool) (*ptypes.BlockResponse, bool, error) {
		return &ptypes.BlockResponse{
			Id:            h,
			CreatedHeight: 123456,
			Msgs:          []ptypes.Msg{}, // set empty msgs
		}, true, nil
	}

	height, networkFunc, valFunc, chainFunc, _ := setupTest(t)
	prov := Provider{valset: valFunc, networkFunc: networkFunc, chainID: chainFunc, portalBlock: portalBlockFunc}
	_, ok, err := prov.XBlock(t.Context(), height, false)
	require.Error(t, err)
	require.Equal(t, "unexpected empty block [BUG]", err.Error())
	require.False(t, ok)
}
