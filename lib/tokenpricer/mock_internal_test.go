package tokenpricer

import (
	"testing"

	"github.com/omni-network/omni/lib/tokens"

	"github.com/stretchr/testify/require"
)

func TestDevnetPricer(t *testing.T) {
	t.Parallel()

	pricer := NewDevnetMock()

	price, err := pricer.Price(t.Context(), tokens.OMNI, tokens.ETH)
	require.NoError(t, err)
	require.InEpsilon(t, 5.0/3000.0, price, 0.0001)
}
