package keeper

import (
	"testing"

	ptypes "github.com/omni-network/omni/halo/portal/types"
	"github.com/omni-network/omni/halo/registry/types"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/tutil"
	"github.com/omni-network/omni/lib/xchain"

	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdktestutil "github.com/cosmos/cosmos-sdk/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	fuzz "github.com/google/gofuzz"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

func TestAddPortal(t *testing.T) {
	t.Parallel()
	fuzzer := fuzz.New().NumElements(1, 8).NilChance(0)
	dupChainIDs := make(map[uint64]bool)
	newPortal := func() *Portal {
		var p Portal
		for {
			fuzzer.Fuzz(&p)
			p.Address = tutil.RandomAddress().Bytes()
			if dupChainIDs[p.GetChainId()] {
				continue
			} else if err := p.Verify(); err != nil {
				t.Log("retrying due to invalid fuzz")
				continue
			}

			break
		}
		dupChainIDs[p.GetChainId()] = true

		return &p
	}

	ctx, keeper, emitPortal := setupKeeper(t)

	ensurePortals := func(t *testing.T, portals ...*Portal) {
		t.Helper()
		latestPortals, err := keeper.getLatestPortals(ctx)
		require.NoError(t, err)

		cmpOpts := cmp.Options{cmpopts.IgnoreUnexported(Portal{})}
		if !cmp.Equal(portals, latestPortals, cmpOpts) {
			t.Error(cmp.Diff(portals, latestPortals, cmpOpts))
		}

		for _, p := range portals {
			ok, err := keeper.SupportedChain(ctx, p.GetChainId())
			require.NoError(t, err)
			require.True(t, ok)
		}
	}

	portals, err := keeper.getLatestPortals(ctx)
	require.NoError(t, err)
	require.Empty(t, portals)
	ok, err := keeper.SupportedChain(ctx, 99)
	require.NoError(t, err)
	require.False(t, ok)

	ctx = ctx.WithBlockHeight(1)

	// Add a new portal (in block 1)
	p1 := newPortal()
	require.NoError(t, keeper.addPortal(ctx, p1))
	ensurePortals(t, p1)

	// Add a second portal (also in block 1)
	p2 := newPortal()
	require.NoError(t, keeper.addPortal(ctx, p2))
	ensurePortals(t, p1, p2)

	ctx = ctx.WithBlockHeight(2)

	// Add p1 with an additional shard (in block 2), thats fine
	p1.ShardIds = append(p1.ShardIds, p1.GetShardIds()[0]+1) // Note there is a non-zero probability of this adding a duplicate shard that will cause the test to fail.
	require.NoError(t, keeper.addPortal(ctx, p1))
	ensurePortals(t, p1, p2)

	// Add p1 again (in block 2), that errors
	require.Error(t, keeper.addPortal(ctx, p1))

	// Add p1 with a different address, that errors
	p3 := proto.Clone(p1).(*Portal)
	p3.Address = tutil.RandomAddress().Bytes()
	require.Error(t, keeper.addPortal(ctx, p3))
	ensurePortals(t, p1, p2)

	// Add p1 with a different deploy height, that errors
	p4 := proto.Clone(p1).(*Portal)
	p4.DeployHeight++
	require.Error(t, keeper.addPortal(ctx, p4))
	ensurePortals(t, p1, p2)

	// Add p1 again (in block 2), that errors
	require.Error(t, keeper.addPortal(ctx, p1))

	// We added portals in two blocks.
	require.Len(t, emitPortal.emittedIDs, 2)
	require.EqualValues(t, []uint64{1, 2}, emitPortal.emittedIDs)
}

func setupKeeper(t *testing.T) (sdk.Context, Keeper, *testEmitPortal) {
	t.Helper()

	key := storetypes.NewKVStoreKey(types.ModuleName)
	storeSvc := runtime.NewKVStoreService(key)
	ctx := sdktestutil.DefaultContext(key, storetypes.NewTransientStoreKey("test_key"))
	ctx = ctx.WithBlockHeight(1)
	ctx = ctx.WithChainID(netconf.Simnet.Static().OmniConsensusChainIDStr())

	emitPortal := new(testEmitPortal)
	k, err := NewKeeper(emitPortal, storeSvc, nil, netconf.ChainNamer(netconf.Simnet))
	require.NoError(t, err, "new keeper")

	return ctx, k, emitPortal
}

var _ ptypes.EmitPortal = &testEmitPortal{}

type testEmitPortal struct {
	emittedIDs []uint64
}

func (t *testEmitPortal) EmitMsg(_ sdk.Context, typ ptypes.MsgType, msgTypeID uint64, destChainID uint64, shardID xchain.ShardID) (uint64, error) {
	if typ != ptypes.MsgTypeNetwork {
		return 0, errors.New("invalid message type")
	} else if destChainID != xchain.BroadcastChainID {
		return 0, errors.New("invalid destination chain id")
	} else if shardID != xchain.ShardBroadcast0 {
		return 0, errors.New("invalid shard id")
	}

	t.emittedIDs = append(t.emittedIDs, msgTypeID)

	return uint64(len(t.emittedIDs)), nil
}
