package ethbackend

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"

	fuzz "github.com/google/gofuzz"
	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		txData  ethtypes.TxData
		fromHex string
	}{
		{
			name:   "legacy tx",
			txData: &ethtypes.LegacyTx{},
		},
		{
			name:   "dynamic fee tx",
			txData: &ethtypes.DynamicFeeTx{},
		},
		{
			name:    "legacy tx with zero prefix",
			txData:  &ethtypes.LegacyTx{},
			fromHex: "0x002985c832a67c0b31a05e909f443b641624da52",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			f := fuzz.New().NilChance(0).NumElements(1, 8)

			var from common.Address
			f.Fuzz(&from)
			f.Fuzz(test.txData)

			if test.fromHex != "" {
				from = common.HexToAddress(test.fromHex)
			}

			signer := backendStubSigner{}

			tx := ethtypes.NewTx(test.txData)
			tx2, err := tx.WithSignature(signer, from.Bytes())
			require.NoError(t, err)

			from2, err := signer.Sender(tx2)
			require.NoError(t, err)

			require.Equal(t, from, from2)
		})
	}
}
