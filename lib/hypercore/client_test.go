package hypercore_test

import (
	"flag"
	"os"
	"strings"
	"testing"

	"github.com/omni-network/omni/lib/hypercore"

	"github.com/ethereum/go-ethereum/crypto"

	"github.com/stretchr/testify/require"
)

var integration = flag.Bool("integration", false, "run integration tests")

func TestUseBigBlocks(t *testing.T) {
	t.Parallel()

	if !*integration {
		t.Skip("Skipping integration test. Use -integration flag to run")
	}

	// Private key needs atleast 1 USDC on Hypercore to run this test
	pkHex := os.Getenv("TEST_PRIVATE_KEY")
	if pkHex == "" {
		t.Skip("TEST_PRIVATE_KEY environment variable not set, skipping integration test")
	}

	pk, err := crypto.HexToECDSA(strings.TrimPrefix(pkHex, "0x"))
	require.NoError(t, err)

	signer := hypercore.NewPrivateKeySigner(pk)
	client := hypercore.NewClient(signer)

	err = client.UseBigBlocks(t.Context())
	require.NoError(t, err)
}
