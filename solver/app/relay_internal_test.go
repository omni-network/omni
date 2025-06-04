package app

import (
	"math/big"
	"testing"

	"github.com/omni-network/omni/contracts/bindings"

	"github.com/ethereum/go-ethereum/common"

	"github.com/stretchr/testify/require"
)

func TestValidateGaslessOrder(t *testing.T) {
	t.Parallel()

	validOrder := bindings.IERC7683GaslessCrossChainOrder{
		OriginSettler: common.HexToAddress("0x1234567890123456789012345678901234567890"),
		User:          common.HexToAddress("0xabcdefabcdefabcdefabcdefabcdefabcdefabcd"),
		Nonce:         big.NewInt(1),
		OriginChainId: big.NewInt(1),
		OpenDeadline:  uint32(1000),
		FillDeadline:  uint32(2000),
		OrderDataType: common.HexToHash("0x1234"),
		OrderData:     []byte("test order data"),
	}

	tests := []struct {
		name    string
		order   bindings.IERC7683GaslessCrossChainOrder
		wantErr bool
		errMsg  string
	}{
		{
			name:    "valid order",
			order:   validOrder,
			wantErr: false,
		},
		{
			name: "zero user address",
			order: func() bindings.IERC7683GaslessCrossChainOrder {
				o := validOrder
				o.User = common.Address{}

				return o
			}(),
			wantErr: true,
			errMsg:  "user address cannot be zero",
		},
		{
			name: "zero origin settler",
			order: func() bindings.IERC7683GaslessCrossChainOrder {
				o := validOrder
				o.OriginSettler = common.Address{}

				return o
			}(),
			wantErr: true,
			errMsg:  "origin settler cannot be zero",
		},
		{
			name: "zero nonce",
			order: func() bindings.IERC7683GaslessCrossChainOrder {
				o := validOrder
				o.Nonce = big.NewInt(0)

				return o
			}(),
			wantErr: true,
			errMsg:  "nonce must be positive",
		},
		{
			name: "invalid fill deadline",
			order: func() bindings.IERC7683GaslessCrossChainOrder {
				o := validOrder
				o.FillDeadline = 500

				return o
			}(),
			wantErr: true,
			errMsg:  "fill deadline must be after open deadline",
		},
		{
			name: "empty order data",
			order: func() bindings.IERC7683GaslessCrossChainOrder {
				o := validOrder
				o.OrderData = []byte{}

				return o
			}(),
			wantErr: true,
			errMsg:  "order data cannot be empty",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			err := validateGaslessOrder(tt.order)
			if tt.wantErr {
				require.Error(t, err)
				require.Contains(t, err.Error(), tt.errMsg)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
