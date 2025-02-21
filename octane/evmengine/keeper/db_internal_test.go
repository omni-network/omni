package keeper

import (
	"context"
	"testing"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/k1util"
	"github.com/omni-network/omni/lib/tutil"

	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"

	"github.com/ethereum/go-ethereum/common"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"
	"github.com/stretchr/testify/require"
)

func TestKeeper_withdrawalsPersistence(t *testing.T) {
	t.Parallel()

	cdc := getCodec(t)
	txConfig := authtx.NewTxConfig(cdc, nil)

	mockEngine, err := newMockEngineAPI(2)
	require.NoError(t, err)
	cmtAPI := newMockCometAPI(t, nil)
	header := cmtproto.Header{Height: 1, AppHash: tutil.RandomHash().Bytes()}
	nxtAddr, err := k1util.PubKeyToAddress(cmtAPI.validatorSet.Validators[1].PubKey)
	require.NoError(t, err)

	ctx, storeService := setupCtxStore(t, &header)
	ctx = ctx.WithExecMode(sdk.ExecModeFinalize)

	ap := mockAddressProvider{
		address: nxtAddr,
	}
	frp := newRandomFeeRecipientProvider()
	evmLogProc := mockEventProc{deliverErr: errors.New("test error")}
	maxWithdrawalsPerBlock := uint64(4)
	keeper, err := NewKeeper(cdc, storeService, &mockEngine, txConfig, ap, frp, maxWithdrawalsPerBlock, evmLogProc)
	require.NoError(t, err)

	addr1 := tutil.RandomAddress()
	addr2 := tutil.RandomAddress()
	addr3 := tutil.RandomAddress()

	type testCase struct {
		addr   common.Address
		height uint64
		amount uint64
		expID  uint64
	}

	inputs := []testCase{
		{addr1, 1, 777, 1},
		{addr2, 2, 8888, 2},
		{addr1, 100, 9999999, 3},
		{addr3, 120, 10000000, 4},
	}

	for _, in := range inputs {
		ctx = ctx.WithBlockHeight(int64(in.height))
		err := keeper.InsertWithdrawal(ctx, in.addr, in.amount)
		require.NoError(t, err)
	}

	withdrawals, err := getAllWithdrawals(ctx, keeper)
	require.NoError(t, err)
	require.Len(t, withdrawals, len(inputs))

	matchesTestCase := func(w *Withdrawal, in testCase) {
		require.Equal(t, in.expID, w.GetId())
		require.Equal(t, in.addr.Bytes(), w.GetAddress())
		require.Equal(t, in.amount, w.GetAmountGwei())
		require.Equal(t, in.height, w.GetCreatedHeight())
	}

	for i, in := range inputs {
		matchesTestCase(withdrawals[i], in)
	}

	withdrawalsByAddr, err := keeper.listWithdrawalsByAddress(ctx, addr1)
	require.NoError(t, err)
	require.Len(t, withdrawalsByAddr, 2)
	require.Equal(t, addr1.Bytes(), withdrawalsByAddr[0].GetAddress())

	withdrawalsByAddr, err = keeper.listWithdrawalsByAddress(ctx, addr2)
	require.NoError(t, err)
	require.Len(t, withdrawalsByAddr, 1)
	require.Equal(t, addr2.Bytes(), withdrawalsByAddr[0].GetAddress())

	withdrawalsByAddr, err = keeper.listWithdrawalsByAddress(ctx, tutil.RandomAddress())
	require.NoError(t, err)
	require.Empty(t, withdrawalsByAddr)

	// make sure we have no withdrawals below and at height 1
	withdrawalsByHeight, err := keeper.EligibleWithdrawals(ctx.WithBlockHeight(0))
	require.NoError(t, err)
	require.Empty(t, withdrawalsByHeight)

	withdrawalsByHeight, err = keeper.EligibleWithdrawals(ctx.WithBlockHeight(1))
	require.NoError(t, err)
	require.Empty(t, withdrawalsByHeight)

	// make sure we have exactly one withdrawal below height 2
	withdrawalsByHeight, err = keeper.EligibleWithdrawals(ctx.WithBlockHeight(2))
	require.NoError(t, err)
	require.Len(t, withdrawalsByHeight, 1)
	matchesTestCase(withdrawalsByHeight[0], inputs[0])

	// under height 50 we only have 2 withdrawals
	withdrawalsByHeight, err = keeper.EligibleWithdrawals(ctx.WithBlockHeight(50))
	require.NoError(t, err)
	require.Len(t, withdrawalsByHeight, 2)
	matchesTestCase(withdrawalsByHeight[0], inputs[0])
	matchesTestCase(withdrawalsByHeight[1], inputs[1])

	// under height 1000 we get all of them
	withdrawalsByHeight, err = keeper.EligibleWithdrawals(ctx.WithBlockHeight(1000))
	require.NoError(t, err)
	require.Len(t, withdrawalsByHeight, 4)
	matchesTestCase(withdrawalsByHeight[0], inputs[0])
	matchesTestCase(withdrawalsByHeight[1], inputs[1])
	matchesTestCase(withdrawalsByHeight[2], inputs[2])
	matchesTestCase(withdrawalsByHeight[3], inputs[3])

	// under height 1000 we get the first 2 if we limit the output by 2
	keeper.maxWithdrawalsPerBlock /= 2
	withdrawalsByHeight, err = keeper.EligibleWithdrawals(ctx.WithBlockHeight(1000))
	require.NoError(t, err)
	require.Len(t, withdrawalsByHeight, int(keeper.maxWithdrawalsPerBlock))
	matchesTestCase(withdrawalsByHeight[0], inputs[0])
	matchesTestCase(withdrawalsByHeight[1], inputs[1])
}

// getAllWithdrawals returns all withdrawals in the keeper DB.
func getAllWithdrawals(ctx context.Context, k *Keeper) ([]*Withdrawal, error) {
	iter, err := k.withdrawalTable.List(ctx, WithdrawalIdIndexKey{})
	if err != nil {
		return nil, errors.Wrap(err, "list withdrawals")
	}
	defer iter.Close()

	var withdrawals []*Withdrawal

	for iter.Next() {
		val, err := iter.Value()
		if err != nil {
			return nil, errors.Wrap(err, "get withdrawal")
		}
		withdrawals = append(withdrawals, val)
	}

	return withdrawals, nil
}
