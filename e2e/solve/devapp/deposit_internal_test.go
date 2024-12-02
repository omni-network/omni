package devapp

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"

	"github.com/stretchr/testify/require"
)

func TestPackUnpack(t *testing.T) {
	t.Parallel()

	dep := DepositArgs{
		OnBehalfOf: common.Address{},
		Amount:     big.NewInt(1),
	}

	packed, err := packDeposit(dep)
	require.NoError(t, err)

	dep2, err := unpackDeposit(packed)
	require.NoError(t, err)

	require.Equal(t, dep, dep2)
}
