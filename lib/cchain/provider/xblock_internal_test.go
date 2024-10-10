package provider

import (
	"context"
	"testing"
	"time"

	ptypes "github.com/omni-network/omni/halo/portal/types"
	"github.com/omni-network/omni/lib/cchain"
	"github.com/omni-network/omni/lib/tutil"

	ctypes "github.com/cometbft/cometbft/rpc/core/types"
	"github.com/cometbft/cometbft/types"

	fuzz "github.com/google/gofuzz"
	"github.com/stretchr/testify/require"
)

//go:generate go test . -golden -clean

func setupTest(t *testing.T) (uint64, valsetFunc, chainIDFunc, headerFunc, portalBlockFunc) {
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
		var valSetMsg *ptypes.Msg
		f.Fuzz(&valSetMsg)
		valSetMsg.Type = uint32(ptypes.MsgTypeValSet)
		valSetMsg.MsgTypeId = h

		return &ptypes.BlockResponse{
			Id:            h,
			CreatedHeight: 123456,
			Msgs:          []*ptypes.Msg{valSetMsg},
		}, true, nil
	}

	return height, valFunc, chainFunc, headerFunc, portalBlockFunc
}

// TestXBlock ensures we receive expected xblock response from provider.
func TestXBlock(t *testing.T) {
	t.Parallel()

	height, valFunc, chainFunc, headerFunc, portalBlockFunc := setupTest(t)
	prov := Provider{valset: valFunc, chainID: chainFunc, header: headerFunc, portalBlock: portalBlockFunc}

	block, ok, err := prov.XBlock(context.Background(), height, false)
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
			Msgs:          []*ptypes.Msg{nil, nil, nil}, // set msgs to nil
		}, true, nil
	}

	height, valFunc, chainFunc, headerFunc, _ := setupTest(t)
	prov := Provider{valset: valFunc, chainID: chainFunc, header: headerFunc, portalBlock: portalBlockFunc}
	_, ok, err := prov.XBlock(context.Background(), height, false)
	require.Errorf(t, err, "unexpected nil msg in block response msgs slice possible malicious response [BUG]")
	require.False(t, ok)
}
