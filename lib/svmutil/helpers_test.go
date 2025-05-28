package svmutil_test

import (
	"testing"

	"github.com/omni-network/omni/lib/svmutil"

	"github.com/ethereum/go-ethereum/crypto"

	"github.com/stretchr/testify/require"
)

func TestMapEVMKey(t *testing.T) {
	t.Parallel()

	const privHex = "289c2857d4598e37fb9647507e47a309d6133539bf21a8b9cb6df88fd5232032"
	evm, err := crypto.HexToECDSA(privHex)
	require.NoError(t, err)

	svm := svmutil.MapEVMKey(evm)

	evmAddr := crypto.PubkeyToAddress(evm.PublicKey).Hex()
	svmAddr := svm.PublicKey().String()

	require.Equal(t, "0x970E8128AB834E8EAC17Ab8E3812F010678CF791", evmAddr)
	require.Equal(t, "p6K5Ff9XkxxkUbriA9Kzrnn49H3cuGXEw1AiD4faNJyPQwW1Kxgqw8oN8nwQpAQYwfpF7G72tqyuXXXnvZ4dYYu", svm.String())
	require.Equal(t, "G8n9W2tGpcazbyrdQzniSKV6p4uFJArWCENhapDmHTx9", svmAddr)
}
