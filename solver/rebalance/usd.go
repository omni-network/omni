package rebalance

import (
	"fmt"
	"math/big"
	"strconv"

	"github.com/omni-network/omni/lib/bi"
)

// usdDecimals is the number of decimals used for USD amounts.
// It is set to 6 to match the USDC token's decimal precision.
const usdDecimals = 6

// formatUSD returns a string representation of the base unit amount assuming 6
// decimals, with "USD" appended.
func formatUSD(n *big.Int) string {
	if n == nil {
		return "nil"
	}

	return fmt.Sprintf("%s %s",
		strconv.FormatFloat(bi.ToF64(n, usdDecimals), 'f', -1, 64), // Use FormatFloat 'f' instead of %f since it avoids trailing zeros
		"USD",
	)
}
