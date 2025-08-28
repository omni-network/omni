package tokenpricer

import (
	"math/big"
	"testing"

	"github.com/omni-network/omni/lib/tokens"

	"github.com/stretchr/testify/require"
)

func TestDevnetPricer(t *testing.T) {
	t.Parallel()

	pricer := NewDevnetMock()

	price, err := pricer.Price(t.Context(), tokens.NOM, tokens.ETH)
	require.NoError(t, err)
	require.Equal(t, big.NewRat(5.0, 3000.0*75.0), price)
}
