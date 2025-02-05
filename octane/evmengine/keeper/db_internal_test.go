package keeper

import (
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
	keeper, err := NewKeeper(cdc, storeService, &mockEngine, txConfig, ap, frp, evmLogProc)
	require.NoError(t, err)

	addr1 := common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")
	addr2 := common.HexToAddress("0x70997970C51812dc3A010C7d01b50e0d17dc79C8")

	inputs := []struct {
		addr   common.Address
		height uint64
		amount uint64
		expID  uint64
	}{
		{addr1, 1, 777, 1},
		{addr2, 2, 8888, 2},
		{addr1, 100, 9999999, 3},
	}

	for _, in := range inputs {
		id, err := keeper.insertWithdrawal(ctx, in.addr, in.height, in.amount)
		require.NoError(t, err)
		require.Equal(t, in.expID, id)
	}

	withdrawals, err := keeper.getWithdrawals(ctx)
	require.NoError(t, err)
	require.Len(t, withdrawals, 3)

	for i, in := range inputs {
		require.Equal(t, in.expID, withdrawals[i].GetId())
		require.Equal(t, in.addr.Bytes(), withdrawals[i].GetAddress())
		require.Equal(t, in.amount, withdrawals[i].GetAmountGwei())
		require.Equal(t, in.height, withdrawals[i].GetCreatedHeight())
	}
}
