package tokens

type Token struct {
	Symbol      string
	Name        string
	CoingeckoID string
}

var (
	OMNI = Token{
		Symbol:      "OMNI",
		Name:        "Omni Network",
		CoingeckoID: "omni-network",
	}

	ETH = Token{
		Symbol:      "ETH",
		Name:        "Ether",
		CoingeckoID: "ethereum",
	}

	STETH = Token{
		Symbol:      "stETH",
		Name:        "Lido Staked Ether",
		CoingeckoID: "lido-staked-ether",
	}

	WSTETH = Token{
		Symbol:      "wstETH",
		Name:        "Wrapped Staked Ether",
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
