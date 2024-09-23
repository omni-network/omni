package keeper

import (
	"testing"

	k1 "github.com/cometbft/cometbft/crypto/secp256k1"

	"github.com/stretchr/testify/require"
)

func TestInsertInvalidValidators(t *testing.T) {
	t.Parallel()

	tests := []struct {
		Name string
		Val  *Validator
		Err  string
	}{
		{
			Name: "nil pubkey",
			Err:  "invalid pubkey",
			Val: &Validator{
				PubKey: nil,
			},
		},
		{
			Name: "invalid pubkey",
			Err:  "invalid pubkey",
			Val: &Validator{
				PubKey: []byte{1, 2, 3, 4, 5, 6, 7},
			},
		},
		{
			Name: "negative power",
			Err:  "negative power",
			Val: &Validator{
				PubKey: k1.GenPrivKey().PubKey().Bytes(),
				Power:  -1,
			},
		},
		{
			Name: "valset id",
			Err:  "valset id already set",
			Val: &Validator{
				PubKey:   k1.GenPrivKey().PubKey().Bytes(),
				ValsetId: 1,
			},
		},
		{
			Name: "nil validator",
			Err:  "nil validator",
			Val:  nil,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			keeper, sdkCtx := setupKeeper(t, defaultExpectation())

			sdkCtx = sdkCtx.WithBlockHeight(1)

			_, err := keeper.insertValidatorSet(sdkCtx, []*Validator{test.Val}, false)
			require.ErrorContains(t, err, test.Err)
		})
	}
}
