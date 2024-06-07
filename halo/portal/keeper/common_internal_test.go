package keeper

import (
	"testing"

	"github.com/omni-network/omni/halo/portal/types"

	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdktestutil "github.com/cosmos/cosmos-sdk/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func SetupKeeper(t *testing.T) (Keeper, sdk.Context) {
	t.Helper()

	key := storetypes.NewKVStoreKey(types.ModuleName)
	storeSvc := runtime.NewKVStoreService(key)
	ctx := sdktestutil.DefaultContext(key, storetypes.NewTransientStoreKey("test_key"))
	ctx = ctx.WithBlockHeight(1)

	k, err := NewKeeper(storeSvc)
	require.NoError(t, err, "new keeper")

	return k, ctx
}
