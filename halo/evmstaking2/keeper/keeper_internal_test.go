package keeper

import (
	context "context"
	"testing"

	"github.com/omni-network/omni/halo/evmstaking2/types"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/netconf"
	evmenginetypes "github.com/omni-network/omni/octane/evmengine/types"

	k1 "github.com/cometbft/cometbft/crypto/secp256k1"

	"github.com/ethereum/go-ethereum/common"

	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdktestutil "github.com/cosmos/cosmos-sdk/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/stretchr/testify/require"
)

//nolint:paralleltest // Asserting insertion ids of sequential writes
func TestInsertAndDeleteEVMEvents(t *testing.T) {
	tests := []struct {
		name       string
		event      evmenginetypes.EVMEvent
		insertedID uint64
		height     int64
	}{
		{
			name: "Insert event with address [1,2,3]",
			event: evmenginetypes.EVMEvent{
				Address: []byte{1, 2, 3},
			},
			insertedID: 1,
			height:     0,
		},
		{
			name: "Insert event with address [2,3,4]",
			event: evmenginetypes.EVMEvent{
				Address: []byte{2, 3, 4},
			},
			insertedID: 2,
			height:     1,
		},
	}

	submissionDelay := int64(5)

	keeper, ctx := setupKeeper(t, submissionDelay, nil, nil, nil, nil, nil)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctx = ctx.WithBlockHeight(test.height)

			err := keeper.Deliver(ctx, common.Hash{}, test.event)
			require.NoError(t, err)

			event, err := keeper.eventsTable.Get(ctx, test.insertedID)
			require.NoError(t, err)
			require.Equal(t, test.event.Address, event.GetEvent().Address)
		})
	}

	// Make sure no submission happens for heights in the range 2 to 4
	for h := int64(2); h < keeper.submissionDelay; h++ {
		ctx = ctx.WithBlockHeight(h)
		err := keeper.EndBlock(ctx)
		require.NoError(t, err)
	}

	// All events are present because we did not deliver them.
	for _, test := range tests {
		found, err := keeper.eventsTable.Has(ctx, test.insertedID)
		require.NoError(t, err)
		require.True(t, found)
	}

	// Now "execute" block number `submissionDelay`
	err := keeper.EndBlock(ctx.WithBlockHeight(submissionDelay))
	require.NoError(t, err)

	// All events are deleted now
	for _, test := range tests {
		found, err := keeper.eventsTable.Has(ctx, test.insertedID)
		require.NoError(t, err)
		require.False(t, found)
	}
}

func TestDeliverDelegate(t *testing.T) {
	t.Parallel()

	submissionDelay := int64(3)
	ethClientMock, err := ethclient.NewEngineMock(
		ethclient.WithPortalRegister(netconf.SimnetNetwork()),
		ethclient.WithMockSelfDelegation(k1.GenPrivKey().PubKey(), 1),
	)
	require.NoError(t, err)

	msgServer := &msgServerMock{}

	keeper, ctx := setupKeeper(t, submissionDelay, ethClientMock, authKeeperMock{}, bankKeeperMock{}, stakingKeeperMock{}, msgServer)

	var hash common.Hash
	events, err := keeper.Prepare(ctx, hash)
	require.NoError(t, err)

	require.Len(t, events, 1)

	for _, event := range events {
		err := keeper.Deliver(ctx, hash, event)
		require.NoError(t, err)
	}

	// Make sure the event was persisted.
	insertedID := uint64(1)
	found, err := keeper.eventsTable.Has(ctx, insertedID)
	require.NoError(t, err)
	require.True(t, found)

	ctx = ctx.WithBlockHeight(submissionDelay)
	err = keeper.EndBlock(ctx)
	require.NoError(t, err)

	// Make sure the event was deleted.
	found, err = keeper.eventsTable.Has(ctx, insertedID)
	require.NoError(t, err)
	require.False(t, found)

	// Assert that the message was delivered to the msg server.
	require.Len(t, msgServer.delegateMsgBuffer, 1)
}

func setupKeeper(
	t *testing.T,
	submissionDelay int64,
	ethCl ethclient.EngineClient,
	aKeeper types.AuthKeeper,
	bKeeper types.BankKeeper,
	sKeeper types.StakingKeeper,
	msgServer types.StakingMsgServer,
) (*Keeper, sdk.Context) {
	t.Helper()

	key := storetypes.NewKVStoreKey(types.ModuleName)
	storeSvc := runtime.NewKVStoreService(key)
	ctx := sdktestutil.DefaultContext(key, storetypes.NewTransientStoreKey("test_key"))
	ctx = ctx.WithBlockHeight(1)
	ctx = ctx.WithChainID(netconf.Simnet.Static().OmniConsensusChainIDStr())

	k, err := NewKeeper(
		storeSvc,
		ethCl,
		aKeeper,
		bKeeper,
		sKeeper,
		msgServer,
		submissionDelay,
	)
	require.NoError(t, err, "new keeper")

	return k, ctx
}

type stakingKeeperMock struct{}

func (stakingKeeperMock) GetValidator(context.Context, sdk.ValAddress) (stypes.Validator, error) {
	return stypes.Validator{}, nil
}

type authKeeperMock struct{}

func (authKeeperMock) HasAccount(context.Context, sdk.AccAddress) bool {
	return true
}

func (authKeeperMock) NewAccountWithAddress(context.Context, sdk.AccAddress) sdk.AccountI {
	return nil
}

func (authKeeperMock) SetAccount(context.Context, sdk.AccountI) {}

type bankKeeperMock struct{}

func (bankKeeperMock) MintCoins(context.Context, string, sdk.Coins) error {
	return nil
}

func (bankKeeperMock) SendCoinsFromModuleToAccount(context.Context, string, sdk.AccAddress, sdk.Coins) error {
	return nil
}

type msgServerMock struct {
	createValidatorMsgBuffer []*stypes.MsgCreateValidator
	delegateMsgBuffer        []*stypes.MsgDelegate
}

func (srv *msgServerMock) CreateValidator(_ context.Context, msg *stypes.MsgCreateValidator) (*stypes.MsgCreateValidatorResponse, error) {
	srv.createValidatorMsgBuffer = append(srv.createValidatorMsgBuffer, msg)
	return nil, nil //nolint:nilnil // API requires nil-nil return
}

func (srv *msgServerMock) Delegate(_ context.Context, msg *stypes.MsgDelegate) (*stypes.MsgDelegateResponse, error) {
	srv.delegateMsgBuffer = append(srv.delegateMsgBuffer, msg)
	return nil, nil //nolint:nilnil // API requires nil-nil return
}
