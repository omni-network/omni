package tokens_test

import (
	"testing"

	"github.com/omni-network/omni/lib/tokens"

	"github.com/stretchr/testify/require"
)

// TestDecimals sanity checks token decimals, ensuring we do not add invalid ones.
func TestDecimals(t *testing.T) {
	t.Parallel()

	for _, asset := range tokens.UniqueAssets() {
		if asset == tokens.USDC {
			require.Equal(t, uint(6), asset.Decimals)
		} else {
			require.Equal(t, uint(18), asset.Decimals)
		}
	}
}
