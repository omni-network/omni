package anvil_test

import (
	"context"
	"testing"
	"time"

	"github.com/omni-network/omni/lib/anvil"
	"github.com/omni-network/omni/lib/tutil"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"

	"cosmossdk.io/math"
	"github.com/stretchr/testify/require"
)

const (
	chainName   = "test"
	chainID     = 99
	blockPeriod = time.Second
)

func TestFundAccounts(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	ethCl, stop, err := anvil.Start(ctx, tutil.TempDir(t), chainID)
	require.NoError(t, err)
	defer stop()

	accounts := []common.Address{
		common.HexToAddress("0x111"),
		common.HexToAddress("0x222"),
		common.HexToAddress("0x333"),
	}

	amt := math.NewInt(100).MulRaw(params.Ether).BigInt() // 100 ETH
	err = anvil.FundAccounts(ctx, ethCl, amt, accounts...)
	require.NoError(t, err)

	for _, account := range accounts {
		balance, err := ethCl.BalanceAt(ctx, account, nil)
		require.NoError(t, err)
		require.Equal(t, amt, balance)
	}
}
