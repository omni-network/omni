package keeper

import (
	"context"
	"strings"
	"testing"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/halo/evmredenom"
	"github.com/omni-network/omni/halo/evmstaking/testutil"
	"github.com/omni-network/omni/halo/evmstaking/types"
	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/k1util"
	"github.com/omni-network/omni/lib/netconf"
	evmengkeeper "github.com/omni-network/omni/octane/evmengine/keeper"
	etypes "github.com/omni-network/omni/octane/evmengine/types"

	k1 "github.com/cometbft/cometbft/crypto/secp256k1"

	"github.com/ethereum/go-ethereum/common"

	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdktestutil "github.com/cosmos/cosmos-sdk/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

//nolint:paralleltest // Asserting insertion ids of sequential writes
func TestInsertAndDeliverEVMEvents(t *testing.T) {
	tests := []struct {
		name       string
		event      etypes.EVMEvent
		insertedID uint64
		height     int64
	}{
		{
			name: "Insert event with address [1,2,3]",
			event: etypes.EVMEvent{
				Address: []byte{1, 2, 3},
			},
			insertedID: 1,
			height:     0,
		},
		{
			name: "Insert event with address [2,3,4]",
			event: etypes.EVMEvent{
				Address: []byte{2, 3, 4},
			},
			insertedID: 2,
			height:     1,
		},
	}

	deliverInterval := int64(5)

	keeper, ctx := setupKeeper(t, deliverInterval, nil)

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
	for h := int64(2); h < keeper.deliverInterval; h++ {
		ctx = ctx.WithBlockHeight(h)
		err := keeper.EndBlock(ctx)
		require.NoError(t, err)
	}

	// All events are present because we did not deliver them yet.
	for _, test := range tests {
		assertContains(t, ctx, keeper, test.insertedID)
	}

	// Now "execute" block number `deliverInterval`
	err := keeper.EndBlock(ctx.WithBlockHeight(deliverInterval))
	require.NoError(t, err)

	// All events are deleted now
	for _, test := range tests {
		assertNotContains(t, ctx, keeper, test.insertedID)
	}
}

func TestDeliveryWithBrokenServer(t *testing.T) {
	t.Parallel()

	deliverInterval := int64(3)
	ethStake := int64(7)
	privKey := k1.GenPrivKey()

	ethClientMock, err := ethclient.NewEngineMock(
		ethclient.WithPortalRegister(netconf.SimnetNetwork()),
		ethclient.WithMockValidatorCreation(privKey.PubKey()),
		ethclient.WithMockSelfDelegation(privKey.PubKey(), ethStake),
	)
	require.NoError(t, err)

	ctrl := gomock.NewController(t)
	sServerMock := testutil.NewMockStakingMsgServer(ctrl)
	err = errors.New("unconditional error")
	sServerMock.EXPECT().CreateValidator(gomock.Any(), gomock.Any()).AnyTimes().Return(nil, err)
	sServerMock.EXPECT().Delegate(gomock.Any(), gomock.Any()).AnyTimes().Return(nil, err)

	keeper, ctx := setupKeeper(t, deliverInterval, sServerMock)

	events, err := getStakingEvents(ctx, ethClientMock, keeper)
	require.NoError(t, err)

	expectDelegates := 1
	expectCreates := 1
	expectTotalEvents := expectDelegates + expectCreates

	require.Len(t, events, expectTotalEvents)

	for _, event := range events {
		err := keeper.parseAndDeliver(ctx, &event)
		require.Contains(t, err.Error(), "unconditional error")
	}
}

func TestDeliveryOfInvalidEvents(t *testing.T) {
	t.Parallel()

	deliverInterval := int64(3)
	ethStake := int64(7)
	privKey := k1.GenPrivKey()

	ethClientMock, err := ethclient.NewEngineMock(
		ethclient.WithPortalRegister(netconf.SimnetNetwork()),
		ethclient.WithMockValidatorCreation(privKey.PubKey()),
		ethclient.WithMockSelfDelegation(privKey.PubKey(), ethStake),
	)
	require.NoError(t, err)

	ctrl := gomock.NewController(t)
	sServerMock := testutil.NewMockStakingMsgServer(ctrl)
	sServerMock.EXPECT().CreateValidator(gomock.Any(), gomock.Any()).AnyTimes().Return(nil, nil)
	sServerMock.EXPECT().Delegate(gomock.Any(), gomock.Any()).AnyTimes().Return(nil, nil)

	keeper, ctx := setupKeeper(t, deliverInterval, sServerMock)

	events, err := getStakingEvents(ctx, ethClientMock, keeper)
	require.NoError(t, err)

	expectDelegates := 1
	expectCreates := 1
	expectTotalEvents := expectDelegates + expectCreates

	require.Len(t, events, expectTotalEvents)

	// Break the address for both events and make sure parsing fails
	for _, event := range events {
		event.Address = []byte{}
		err := keeper.parseAndDeliver(ctx, &event)
		require.Contains(t, err.Error(), "invalid address length")
	}

	// Break the topics for both events and make sure parsing fails
	for _, event := range events {
		event.Topics = [][]byte{}
		err := keeper.parseAndDeliver(ctx, &event)
		require.Contains(t, err.Error(), "empty topics")
	}

	createValEvent := events[0]
	// Break the data for the create validator event
	createValEvent.Data = []byte{}
	err = keeper.parseAndDeliver(ctx, &createValEvent)
	require.Contains(t, err.Error(), "create validator: pubkey to cosmos")

	// Deliver the event so that we can test delegation
	err = keeper.parseAndDeliver(ctx, &events[0])
	require.NoError(t, err)

	// Can't add same validator twice (this relies on sKeeper stub working correctly)
	err = keeper.parseAndDeliver(ctx, &events[0])
	require.Contains(t, err.Error(), "create validator: validator already exists")

	delegateEvent := events[1]
	// Break the data for the delegate event
	delegateEvent.Data = []byte{}
	err = keeper.parseAndDeliver(ctx, &delegateEvent)
	require.Contains(t, err.Error(), "stake amount missing")
}

func TestHappyPathDelivery(t *testing.T) {
	t.Parallel()

	deliverInterval := int64(3)
	ethStake := int64(8 * evmredenom.Factor)

	privKey := k1.GenPrivKey()

	ethClientMock, err := ethclient.NewEngineMock(
		ethclient.WithPortalRegister(netconf.SimnetNetwork()),
		ethclient.WithMockValidatorCreation(privKey.PubKey()),
		ethclient.WithMockSelfDelegation(privKey.PubKey(), ethStake),
	)
	require.NoError(t, err)

	var delegateMsgBuffer []*stypes.MsgDelegate
	var createValidatorMsgBuffer []*stypes.MsgCreateValidator

	ctrl := gomock.NewController(t)
	sServerMock := testutil.NewMockStakingMsgServer(ctrl)
	sServerMock.EXPECT().
		CreateValidator(gomock.Any(), gomock.Any()).
		AnyTimes().Do(func(ctx context.Context, msg *stypes.MsgCreateValidator) {
		createValidatorMsgBuffer = append(createValidatorMsgBuffer, msg)
	}).
		Return(new(stypes.MsgCreateValidatorResponse), nil)
	sServerMock.EXPECT().
		Delegate(gomock.Any(), gomock.Any()).
		AnyTimes().Do(func(ctx context.Context, msg *stypes.MsgDelegate) {
		delegateMsgBuffer = append(delegateMsgBuffer, msg)
	}).
		Return(new(stypes.MsgDelegateResponse), nil)

	keeper, ctx := setupKeeper(t, deliverInterval, sServerMock)

	events, err := getStakingEvents(ctx, ethClientMock, keeper)
	require.NoError(t, err)

	expectDelegates := 1
	expectCreates := 1
	expectTotalEvents := expectDelegates + expectCreates

	require.Len(t, events, expectTotalEvents)

	for _, event := range events {
		err := keeper.Deliver(ctx, common.Hash{}, event)
		require.NoError(t, err)
	}

	// Make sure the events were persisted.
	for id := 1; id <= expectTotalEvents; id++ {
		assertContains(t, ctx, keeper, uint64(id))
	}

	ctx = ctx.WithBlockHeight(deliverInterval)
	err = keeper.EndBlock(ctx)
	require.NoError(t, err)

	// Make sure the events were deleted.
	for id := 1; id <= expectTotalEvents; id++ {
		assertNotContains(t, ctx, keeper, uint64(id))
	}

	// Assert that the message was delivered to the msg server.
	require.Len(t, delegateMsgBuffer, 1)
	msg := delegateMsgBuffer[0]
	// Sanity check of addresses
	require.Len(t, msg.DelegatorAddress, 45)
	require.Len(t, msg.ValidatorAddress, 52)
	require.True(t, strings.HasPrefix(msg.DelegatorAddress, "cosmos"), msg.DelegatorAddress)
	require.True(t, strings.HasPrefix(msg.ValidatorAddress, "cosmosvaloper"), msg.ValidatorAddress)

	stake := evmredenom.ToStakeCoin(bi.Ether(ethStake))
	require.Equal(t, msg.Amount, stake)

	require.Len(t, createValidatorMsgBuffer, 1)
	msg2 := createValidatorMsgBuffer[0]
	// Sanity check of addresses
	require.Len(t, msg2.ValidatorAddress, 52)
	require.True(t, strings.HasPrefix(msg2.ValidatorAddress, "cosmosvaloper"), msg.ValidatorAddress)

	createdAmount := bi.Ether(1 * evmredenom.Factor)
	expected := evmredenom.ToStakeCoin(createdAmount)
	require.Equal(t, expected.String(), msg2.Value.String())
}

func TestNonSelfDelegationEventDelivery(t *testing.T) {
	t.Parallel()

	deliverInterval := int64(3)
	ethStake := int64(1)

	privKey := k1.GenPrivKey()
	delegatorPrivKey := k1.GenPrivKey()
	delegatorAddr, err := k1util.PubKeyToAddress(delegatorPrivKey.PubKey())
	if err != nil {
		panic(errors.Wrap(err, "pubkey to address"))
	}

	ethClientMock, err := ethclient.NewEngineMock(
		ethclient.WithPortalRegister(netconf.SimnetNetwork()),
		ethclient.WithMockValidatorCreation(privKey.PubKey()),
		ethclient.WithMockDelegation(privKey.PubKey(), delegatorAddr, ethStake),
	)
	require.NoError(t, err)

	var delegateMsgBuffer []*stypes.MsgDelegate
	var createValidatorMsgBuffer []*stypes.MsgCreateValidator

	ctrl := gomock.NewController(t)
	sServerMock := testutil.NewMockStakingMsgServer(ctrl)
	sServerMock.EXPECT().
		CreateValidator(gomock.Any(), gomock.Any()).
		AnyTimes().Do(func(ctx context.Context, msg *stypes.MsgCreateValidator) {
		createValidatorMsgBuffer = append(createValidatorMsgBuffer, msg)
	}).
		Return(new(stypes.MsgCreateValidatorResponse), nil)
	sServerMock.EXPECT().
		Delegate(gomock.Any(), gomock.Any()).
		AnyTimes().Do(func(ctx context.Context, msg *stypes.MsgDelegate) {
		delegateMsgBuffer = append(delegateMsgBuffer, msg)
	}).
		Return(new(stypes.MsgDelegateResponse), nil)

	keeper, ctx := setupKeeper(t, deliverInterval, sServerMock)

	events, err := getStakingEvents(ctx, ethClientMock, keeper)
	require.NoError(t, err)

	expectDelegates := 1
	expectCreates := 1
	expectTotalEvents := expectDelegates + expectCreates

	require.Len(t, events, expectTotalEvents)

	for _, event := range events {
		err := keeper.Deliver(ctx, common.Hash{}, event)
		require.NoError(t, err)
	}

	// Make sure the events were persisted.
	for id := 1; id <= expectTotalEvents; id++ {
		assertContains(t, ctx, keeper, uint64(id))
	}

	ctx = ctx.WithBlockHeight(deliverInterval)
	err = keeper.EndBlock(ctx)
	require.NoError(t, err)

	// Make sure the events were deleted.
	for id := 1; id <= expectTotalEvents; id++ {
		assertNotContains(t, ctx, keeper, uint64(id))
	}

	// Assert that the message was delivered to the msg server.
	require.Len(t, delegateMsgBuffer, 1)
	msg := delegateMsgBuffer[0]
	// Sanity check of addresses
	require.Len(t, msg.DelegatorAddress, 45)
	require.Len(t, msg.ValidatorAddress, 52)
	require.True(t, strings.HasPrefix(msg.DelegatorAddress, "cosmos"), msg.DelegatorAddress)
	require.True(t, strings.HasPrefix(msg.ValidatorAddress, "cosmosvaloper"), msg.ValidatorAddress)
	expected := evmredenom.ToStakeCoin(bi.Ether(ethStake))
	require.Equal(t, expected, msg.Amount)
}

func TestUndelegationEventDelivery(t *testing.T) {
	t.Parallel()

	deliverInterval := int64(3)
	ethStake := int64(1)

	privKey := k1.GenPrivKey()
	delegatorPrivKey := k1.GenPrivKey()
	delegatorAddr, err := k1util.PubKeyToAddress(delegatorPrivKey.PubKey())
	if err != nil {
		panic(errors.Wrap(err, "pubkey to address"))
	}

	ethClientMock, err := ethclient.NewEngineMock(
		ethclient.WithPortalRegister(netconf.SimnetNetwork()),
		ethclient.WithMockValidatorCreation(privKey.PubKey()),
		ethclient.WithMockUndelegation(privKey.PubKey(), delegatorAddr, ethStake),
	)
	require.NoError(t, err)

	var undelegateMsgBuffer []*stypes.MsgUndelegate
	var createValidatorMsgBuffer []*stypes.MsgCreateValidator

	ctrl := gomock.NewController(t)
	sServerMock := testutil.NewMockStakingMsgServer(ctrl)
	sServerMock.EXPECT().
		CreateValidator(gomock.Any(), gomock.Any()).
		AnyTimes().Do(func(ctx context.Context, msg *stypes.MsgCreateValidator) {
		createValidatorMsgBuffer = append(createValidatorMsgBuffer, msg)
	}).
		Return(new(stypes.MsgCreateValidatorResponse), nil)

	sServerMock.EXPECT().
		Undelegate(gomock.Any(), gomock.Any()).
		AnyTimes().Do(func(ctx context.Context, msg *stypes.MsgUndelegate) {
		undelegateMsgBuffer = append(undelegateMsgBuffer, msg)
	}).
		Return(new(stypes.MsgUndelegateResponse), nil)

	keeper, ctx := setupKeeper(t, deliverInterval, sServerMock)

	events, err := getStakingEvents(ctx, ethClientMock, keeper)
	require.NoError(t, err)

	expectUndelegates := 1
	expectCreates := 1
	expectTotalEvents := expectUndelegates + expectCreates

	require.Len(t, events, expectTotalEvents)

	for _, event := range events {
		err := keeper.Deliver(ctx, common.Hash{}, event)
		require.NoError(t, err)
	}

	// Make sure the events were persisted.
	for id := 1; id <= expectTotalEvents; id++ {
		assertContains(t, ctx, keeper, uint64(id))
	}

	ctx = ctx.WithBlockHeight(deliverInterval)
	err = keeper.EndBlock(ctx)
	require.NoError(t, err)

	// Make sure the events were deleted.
	for id := 1; id <= expectTotalEvents; id++ {
		assertNotContains(t, ctx, keeper, uint64(id))
	}

	// Assert that the message was delivered to the msg server.
	require.Len(t, undelegateMsgBuffer, 1)
	msg := undelegateMsgBuffer[0]
	// Sanity check of addresses
	require.Len(t, msg.DelegatorAddress, 45)
	require.Len(t, msg.ValidatorAddress, 52)
	require.True(t, strings.HasPrefix(msg.DelegatorAddress, "cosmos"), msg.DelegatorAddress)
	require.True(t, strings.HasPrefix(msg.ValidatorAddress, "cosmosvaloper"), msg.ValidatorAddress)
	expected := evmredenom.ToStakeCoin(bi.Ether(ethStake))
	require.Equal(t, expected, msg.Amount)
}

func TestEditValidator(t *testing.T) {
	t.Parallel()

	privKey := k1.GenPrivKey()
	valAddr, err := k1util.PubKeyToAddress(privKey.PubKey())
	require.NoError(t, err)

	params := &bindings.StakingEditValidatorParams{
		Moniker:                  "moniker",
		Identity:                 "identity",
		Website:                  "https://website",
		SecurityContact:          "https://contract",
		Details:                  "details",
		CommissionRatePercentage: 99,
		MinSelfDelegation:        bi.N(11),
	}
	ethClientMock, err := ethclient.NewEngineMock(
		ethclient.WithMockEditValidator(privKey.PubKey(), params),
	)
	require.NoError(t, err)

	var msgBuffer []*stypes.MsgEditValidator

	ctrl := gomock.NewController(t)
	sServerMock := testutil.NewMockStakingMsgServer(ctrl)
	sServerMock.EXPECT().
		EditValidator(gomock.Any(), gomock.Any()).
		Times(1).
		DoAndReturn(
			func(ctx context.Context, msg *stypes.MsgEditValidator) (*stypes.MsgEditValidatorResponse, error) {
				msgBuffer = append(msgBuffer, msg)
				return new(stypes.MsgEditValidatorResponse), nil
			},
		)

	keeper, ctx := setupKeeper(t, 1, sServerMock)

	events, err := getStakingEvents(ctx, ethClientMock, keeper)
	require.NoError(t, err)
	require.Len(t, events, 1)

	for _, event := range events {
		err := keeper.Deliver(ctx, common.Hash{}, event)
		require.NoError(t, err)
	}

	ctx = ctx.WithBlockHeight(1)
	err = keeper.EndBlock(ctx)
	require.NoError(t, err)

	require.Len(t, msgBuffer, 1)
	msg := msgBuffer[0]

	require.Equal(t, msg.ValidatorAddress, sdk.ValAddress(valAddr.Bytes()).String())
	require.Equal(t, msg.Description.Moniker, params.Moniker)
	require.Equal(t, msg.Description.Identity, params.Identity)
	require.Equal(t, msg.Description.Website, params.Website)
	require.Equal(t, msg.Description.SecurityContact, params.SecurityContact)
	require.Equal(t, msg.Description.Details, params.Details)
	require.Equal(t, msg.MinSelfDelegation.String(), params.MinSelfDelegation.String())
	require.Equal(t, msg.CommissionRate.TruncateInt64(), int64(params.CommissionRatePercentage))
}

func assertContains(t *testing.T, ctx context.Context, keeper *Keeper, eventID uint64) {
	t.Helper()
	found, err := keeper.eventsTable.Has(ctx, eventID)
	require.NoError(t, err)
	require.True(t, found)
}

func assertNotContains(t *testing.T, ctx context.Context, keeper *Keeper, eventID uint64) {
	t.Helper()
	found, err := keeper.eventsTable.Has(ctx, eventID)
	require.NoError(t, err)
	require.False(t, found)
}

func setupKeeper(
	t *testing.T,
	deliverInterval int64,
	sServer types.StakingMsgServer,
) (*Keeper, sdk.Context) {
	t.Helper()

	key := storetypes.NewKVStoreKey(types.ModuleName)
	storeSvc := runtime.NewKVStoreService(key)
	ctx := sdktestutil.DefaultContext(key, storetypes.NewTransientStoreKey("test_key"))
	ctx = ctx.WithBlockHeight(1)
	ctx = ctx.WithChainID(netconf.Simnet.Static().OmniConsensusChainIDStr())

	ctrl := gomock.NewController(t)

	authKeeperMock := testutil.NewMockAuthKeeper(ctrl)
	authKeeperMock.EXPECT().HasAccount(gomock.Any(), gomock.Any()).AnyTimes().Return(true)
	authKeeperMock.EXPECT().NewAccountWithAddress(gomock.Any(), gomock.Any()).AnyTimes().Return(nil)

	bKeeperMock := testutil.NewMockWrappedBankKeeper(ctrl)
	bKeeperMock.EXPECT().MintCoins(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(nil)
	bKeeperMock.EXPECT().SendCoinsFromModuleToAccountNoWithdrawal(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(nil)

	sKeeperMock := testutil.NewMockStakingKeeper(ctrl)
	var seenValidators map[string]bool
	sKeeperMock.EXPECT().GetValidator(gomock.Any(), gomock.Any()).AnyTimes().
		DoAndReturn(func(_ context.Context, addr sdk.ValAddress) (stypes.Validator, error) {
			if seenValidators == nil {
				seenValidators = make(map[string]bool)
			}

			hexAddr := string(addr)

			if _, found := seenValidators[hexAddr]; found {
				return stypes.Validator{}, nil
			}
			seenValidators[hexAddr] = true

			return stypes.Validator{}, errors.New("validator does not exist")
		})

	k, err := NewKeeper(
		storeSvc,
		authKeeperMock,
		bKeeperMock,
		sKeeperMock,
		sServer,
		deliverInterval,
	)
	require.NoError(t, err, "new keeper")

	return k, ctx
}

// getStakingEvents returns the staking events from the mock engine client.
func getStakingEvents(ctx context.Context, cl ethclient.EngineClient, keeper *Keeper) ([]etypes.EVMEvent, error) {
	return evmengkeeper.FetchProcEvents(ctx, cl, common.Hash{}, keeper)
}
