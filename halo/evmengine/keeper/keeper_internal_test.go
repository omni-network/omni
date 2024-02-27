package keeper

import (
	"context"
	"testing"

	cmttypes "github.com/cometbft/cometbft/types"
	"github.com/ethereum/go-ethereum/common"

	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"
	"github.com/omni-network/omni/halo/comet"
	"github.com/stretchr/testify/require"
)

func TestKeeper_isNextProposer(t *testing.T) {
	ctx, storeService := setupCtxStore(t)
	cdc := getCodec(t)
	txConfig := authtx.NewTxConfig(cdc, nil)

	mockEngine, err := newMockEngineAPI()
	require.NoError(t, err)

	ap := mockAddressProvider{
		address: common.BytesToAddress([]byte("test")), // todo valid address
	}
	keeper := NewKeeper(cdc, storeService, &mockEngine, txConfig, ap)
	keeper.cmtAPI = mockCometAPI{}

	keeper.isNextProposer(ctx)
}

type mockCometAPI struct {
	comet.API
}

func (m mockCometAPI) Validators(ctx context.Context, height int64) (*cmttypes.ValidatorSet, bool, error) {
	return nil, false, nil
}
