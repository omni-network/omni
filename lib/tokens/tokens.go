package tokens

import (
	"fmt"
	"math/big"
	"strconv"

	"github.com/omni-network/omni/lib/bi"
)

type Token struct {
	Symbol      string
	Name        string
	Decimals    uint
	CoingeckoID string
}

var (
	OMNI = Token{
		Symbol:      "OMNI",
		Name:        "Omni Network",
		Decimals:    18,
		CoingeckoID: "omni-network",
	}

	ETH = Token{
		Symbol:      "ETH",
		Name:        "Ether",
		Decimals:    18,
		CoingeckoID: "ethereum",
	}

	USDC = Token{
		Symbol:      "USDC",
		Name:        "USD Coin",
		Decimals:    6,
		CoingeckoID: "usdc",
	}

	STETH = Token{
		Symbol:      "stETH",
		Name:        "Lido Staked Ether",
		Decimals:    18,
		CoingeckoID: "lido-staked-ether",
	}

	WSTETH = Token{
		Symbol:      "wstETH",
		Name:        "Wrapped Staked Ether",
		Decimals:    18,
		CoingeckoID: "wrapped-steth",
	}

	all = []Token{OMNI, ETH, STETH, WSTETH}
)

func All() []Token {
	return all
}

func (t Token) String() string {
	return t.Symbol
}

// F64ToAmt returns the float64 in primary units as a big.Int in base units.
// E.g. 1.1 (ether) is returned as 1.1 * 10^18 (wei).
func (t Token) F64ToAmt(n float64) *big.Int {
	if t.Decimals == 6 {
		return bi.Dec6(n)
	}

	return bi.Ether(n)
}

// AmtToF64 returns amount in base units as a float64 in primary units.
// E.g. 1.1 * 10^18 (wei) is returned as 1.1 (ether).
func (t Token) AmtToF64(n *big.Int) float64 {
	return bi.ToF64(n, t.Decimals)
}

// FormatAmt returns a string representation of the provided base unit amount
// in primary units with the token symbol appended. Note all decimals are shown.
func (t Token) FormatAmt(n *big.Int) string {
	if n == nil {
		return "nil"
	}

	return fmt.Sprintf("%s %s",
		strconv.FormatFloat(t.AmtToF64(n), 'f', -1, 64), // Use FormatFloat 'f' instead of %f since it avoids trailing zeros
		t.Symbol,
	)
}
