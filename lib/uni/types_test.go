package uni_test

import (
	"encoding/json"
	"testing"

	"github.com/omni-network/omni/lib/tutil"
	"github.com/omni-network/omni/lib/uni"

	"github.com/ethereum/go-ethereum/common"

	"github.com/gagliardetto/solana-go"
	"github.com/stretchr/testify/require"
)

//go:generate go test . -golden

func TestAddressJSON(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		addr uni.Address
	}{
		{
			name: "zero",
			addr: uni.Address{},
		},
		{
			name: "evm",
			addr: uni.MustHexToAddress("0x012345678901234567890123456789abcdeabcde"),
		},
		{
			name: "evm_zero",
			addr: uni.EVMAddress(common.Address{}),
		},
		{
			name: "svm",
			addr: uni.MustBase58ToAddress("H9eDjRw7UKFDzzts4jK7hmc2chKL2z4cuA8bU1xDnuRA"),
		},
		{
			name: "svm_zero",
			addr: uni.SVMAddress(solana.PublicKey{}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			bz, err := tt.addr.MarshalJSON()
			require.NoError(t, err)

			var addr uni.Address
			err = json.Unmarshal(bz, &addr)
			require.NoError(t, err)

			require.Equal(t, tt.addr, addr)

			tutil.RequireGoldenJSON(t, tt.addr)
		})
	}
}

func TestInvalidJSON(t *testing.T) {
	t.Parallel()
	testInvalidJSON(t, `""`)
	testInvalidJSON(t, `"ab"`)
	testInvalidJSON(t, `"12"`)
	testInvalidJSON(t, `"0"`)
	testInvalidJSON(t, `"0x"`)
	testInvalidJSON(t, `"0x00"`)
	testInvalidJSON(t, `"0x012345678901234567890123456789abcdeabcdz"`)    // Invald hex
	testInvalidJSON(t, `"0x012345678901234567890123456789abcdeabcdef"`)   // Too long for EVM address
	testInvalidJSON(t, `"0x012345678901234567890123456789abcdeabcd"`)     // Too short for EVM address
	testInvalidJSON(t, `"H9eDjRw7UKFDzzts4jK7hmc2chKL2z4cuA8bU1xDnuR0"`)  // Invalid SVM address
	testInvalidJSON(t, `"H9eDjRw7UKFDzzts4jK7hmc2chKL2z4cuA8bU1xDnuRA1"`) // Too long for SVM address
	testInvalidJSON(t, `"H9eDjRw7UKFDzzts4jK7hmc2chKL2z4cuA8bU1xDnu"`)    // Too short for SVM address
}

func testInvalidJSON(t *testing.T, jsonStr string) {
	t.Helper()

	var addr uni.Address
	err := json.Unmarshal([]byte(jsonStr), &addr)
	require.Error(t, err, "expected error for invalid JSON: %s", jsonStr)
}
