package provider

import (
	"context"
	"testing"
	"time"

	ptypes "github.com/omni-network/omni/halo/portal/types"
	rtypes "github.com/omni-network/omni/halo/registry/types"
	"github.com/omni-network/omni/lib/cchain"
	"github.com/omni-network/omni/lib/tutil"

	ctypes "github.com/cometbft/cometbft/rpc/core/types"
	"github.com/cometbft/cometbft/types"

	fuzz "github.com/google/gofuzz"
	"github.com/stretchr/testify/require"
)

//go:generate go test . -golden -clean

func TestXBlock(t *testing.T) {
	t.Parallel()
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

	timestamp := time.Unix(1712312027, 0).UTC()

	valFunc := func(ctx context.Context, h uint64, _ bool) (valSetResponse, bool, error) {
		require.EqualValues(t, height, h)
		var resp []cchain.PortalValidator
		f.Fuzz(&resp)

		return valSetResponse{
			ValSetID:   height,
			Validators: resp,
		}, true, nil
	}
	networkFunc := func(ctx context.Context, h uint64, _ bool) (*rtypes.NetworkResponse, bool, error) {
		require.EqualValues(t, height, h)
		var resp rtypes.NetworkResponse
		f.Fuzz(&resp)

		return &resp, true, nil
	}
	chainFunc := func(ctx context.Context) (uint64, error) {
		return 77, nil
	}
	headerFunc := func(ctx context.Context, h *int64) (*ctypes.ResultHeader, error) {
		return &ctypes.ResultHeader{
			Header: &types.Header{
				Time: timestamp,
			},
		}, nil
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

	prov := Provider{valset: valFunc, networkFunc: networkFunc, chainID: chainFunc, header: headerFunc, portalBlock: portalBlockFunc}

	block, ok, err := prov.XBlock(context.Background(), height, false)
	require.NoError(t, err)
	require.True(t, ok)
	tutil.RequireGoldenJSON(t, block)
}
