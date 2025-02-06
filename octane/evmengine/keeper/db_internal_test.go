package keeper

import (
	"context"
	"testing"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/k1util"
	"github.com/omni-network/omni/lib/tutil"

	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"

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
	keeper, err := NewKeeper(cdc, storeService, &mockEngine, txConfig, ap, frp, evmLogProc)
	require.NoError(t, err)

	addr1 := sdk.AccAddress(tutil.RandomAddress().Bytes())
	addr2 := sdk.AccAddress(tutil.RandomAddress().Bytes())

	inputs := []struct {
		addr   sdk.AccAddress
		height uint64
		amount uint64
		expID  uint64
	}{
		{addr1, 1, 777, 1},
		{addr2, 2, 8888, 2},
		{addr1, 100, 9999999, 3},
	}

	for _, in := range inputs {
		ctx = ctx.WithBlockHeight(int64(in.height))
		err := keeper.InsertWithdrawal(ctx, in.addr, in.amount)
		require.NoError(t, err)
	}

	withdrawals, err := getAllWithdrawals(ctx, keeper)
	require.NoError(t, err)
	require.Len(t, withdrawals, 3)

	for i, in := range inputs {
		require.Equal(t, in.expID, withdrawals[i].GetId())
		require.Equal(t, in.addr.Bytes(), withdrawals[i].GetAddress())
		require.Equal(t, in.amount, withdrawals[i].GetAmountGwei())
		require.Equal(t, in.height, withdrawals[i].GetCreatedHeight())
	}

	withdrawalsByAddr, err := keeper.getWithdrawalsByAddress(ctx, addr1)
	require.NoError(t, err)
	require.Len(t, withdrawalsByAddr, 2)
	require.Equal(t, addr1.Bytes(), withdrawalsByAddr[0].GetAddress())

	withdrawalsByAddr, err = keeper.getWithdrawalsByAddress(ctx, addr2)
	require.NoError(t, err)
	require.Len(t, withdrawalsByAddr, 1)
	require.Equal(t, addr2.Bytes(), withdrawalsByAddr[0].GetAddress())
}

// getAllWithdrawals returns all stored withdrawals.
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
