package tokens

import (
	"math/big"

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

func FromCoingeckoID(id string) (Token, bool) {
	for _, t := range all {
		if t.CoingeckoID == id {
			return t, true
		}
	}

	return Token{}, false
}

// ToPrimaryF64 converts the token amount to a float64 value in the primary unit (e.g. ether, dollar, etc.).
func ToPrimaryF64(token Token, amount *big.Int) float64 {
	if token.Decimals == 6 {
		return toDec6F64(amount)
	}

	return bi.ToEtherF64(amount)
}

func toDec6F64(value *big.Int) float64 {
	f, _ := value.Float64()
	return f / 1e6
}
