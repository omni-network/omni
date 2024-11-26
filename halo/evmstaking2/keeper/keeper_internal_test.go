package keeper

import (
	context "context"
	"errors"
	"strings"
	"testing"

	"github.com/omni-network/omni/halo/evmstaking2/types"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/netconf"
	etypes "github.com/omni-network/omni/octane/evmengine/types"

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

	keeper, ctx := setupKeeper(t, deliverInterval, nil, nil, nil, nil, nil)

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

	// All events are present because we did not deliver them.
	for _, test := range tests {
		found, err := keeper.eventsTable.Has(ctx, test.insertedID)
		require.NoError(t, err)
		require.True(t, found)
	}

	// Now "execute" block number `deliverInterval`
	err := keeper.EndBlock(ctx.WithBlockHeight(deliverInterval))
	require.NoError(t, err)

	// All events are deleted now
	for _, test := range tests {
		found, err := keeper.eventsTable.Has(ctx, test.insertedID)
		require.NoError(t, err)
		require.False(t, found)
	}
}

func TestDelivery(t *testing.T) {
	t.Parallel()

	deliverInterval := int64(3)
	ethStake := int64(7)

	ethClientMock, err := ethclient.NewEngineMock(
		ethclient.WithPortalRegister(netconf.SimnetNetwork()),
		ethclient.WithMockSelfDelegation(k1.GenPrivKey().PubKey(), ethStake),
		ethclient.WithMockValidatorCreation(k1.GenPrivKey().PubKey()),
	)
	require.NoError(t, err)

	sServer := new(msgServerStub)

	keeper, ctx := setupKeeper(t, deliverInterval, ethClientMock, new(authKeeperStub), new(bankKeeperStub), new(stakingKeeperStub), sServer)

	var hash common.Hash
	events, err := keeper.Prepare(ctx, hash)
	require.NoError(t, err)

	require.Len(t, events, 2)

	for _, event := range events {
		err := keeper.Deliver(ctx, hash, event)
		require.NoError(t, err)
	}

	// Make sure the events were persisted.
	for id := uint64(1); id < 3; id++ {
		found, err := keeper.eventsTable.Has(ctx, id)
		require.NoError(t, err)
		require.True(t, found)
	}

	ctx = ctx.WithBlockHeight(deliverInterval)
	err = keeper.EndBlock(ctx)
	require.NoError(t, err)

	// Make sure the events were deleted.
	for id := uint64(1); id < 3; id++ {
		found, err := keeper.eventsTable.Has(ctx, id)
		require.NoError(t, err)
		require.False(t, found)
	}

	// Assert that the message was delivered to the msg server.
	require.Len(t, sServer.delegateMsgBuffer, 1)
	msg := sServer.delegateMsgBuffer[0]
	// Sanity check of addresses
	require.Len(t, msg.DelegatorAddress, 45)
	require.Len(t, msg.ValidatorAddress, 52)
	require.True(t, strings.HasPrefix(msg.DelegatorAddress, "cosmos"), msg.DelegatorAddress)
	require.True(t, strings.HasPrefix(msg.ValidatorAddress, "cosmosvaloper"), msg.ValidatorAddress)
	oneEth := sdk.NewInt64Coin("stake", ethStake*1000000000000000000)
	require.Equal(t, msg.Amount, oneEth)
	require.Len(t, sServer.delegateMsgBuffer, 1)

	require.Len(t, sServer.createValidatorMsgBuffer, 1)
}

func setupKeeper(
	t *testing.T,
	deliverInterval int64,
	ethCl ethclient.EngineClient,
	aKeeper types.AuthKeeper,
	bKeeper types.BankKeeper,
	sKeeper types.StakingKeeper,
	sServer types.StakingMsgServer,
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
		sServer,
		deliverInterval,
	)
	require.NoError(t, err, "new keeper")

	return k, ctx
}

type stakingKeeperStub struct {
	// calls is the number of calls to GetValidator
	calls uint32
}

// GetValidator returns no errors on the first call, because it is called on delegation
// event delivery for the first time and the validator should be available.
// Second time it is called on a validator creation event and it should return an error
// on the pubkey of the new validator.
func (m *stakingKeeperStub) GetValidator(context.Context, sdk.ValAddress) (stypes.Validator, error) {
	if m.calls == 0 {
		m.calls++
		return stypes.Validator{}, nil
	}

	return stypes.Validator{}, errors.New("validator exists")
}

type authKeeperStub struct{}

func (authKeeperStub) HasAccount(context.Context, sdk.AccAddress) bool {
	return true
}

func (authKeeperStub) NewAccountWithAddress(context.Context, sdk.AccAddress) sdk.AccountI {
	return nil
}

func (authKeeperStub) SetAccount(context.Context, sdk.AccountI) {}

type bankKeeperStub struct{}

func (bankKeeperStub) MintCoins(context.Context, string, sdk.Coins) error {
	return nil
}

func (bankKeeperStub) SendCoinsFromModuleToAccount(context.Context, string, sdk.AccAddress, sdk.Coins) error {
	return nil
}

type msgServerStub struct {
	createValidatorMsgBuffer []*stypes.MsgCreateValidator
	delegateMsgBuffer        []*stypes.MsgDelegate
}

func (s *msgServerStub) CreateValidator(_ context.Context, msg *stypes.MsgCreateValidator) (*stypes.MsgCreateValidatorResponse, error) {
	s.createValidatorMsgBuffer = append(s.createValidatorMsgBuffer, msg)
	return new(stypes.MsgCreateValidatorResponse), nil
}

func (s *msgServerStub) Delegate(_ context.Context, msg *stypes.MsgDelegate) (*stypes.MsgDelegateResponse, error) {
	s.delegateMsgBuffer = append(s.delegateMsgBuffer, msg)
	return new(stypes.MsgDelegateResponse), nil
}
