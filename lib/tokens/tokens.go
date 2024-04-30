package tokens

type Token string

const (
	OMNI Token = "OMNI"
	ETH  Token = "ETH"
)

var (
	coingeckoIDs = map[Token]string{
		OMNI: "omni-network",
		ETH:  "ethereum",
	}
)

func (t Token) String() string {
	return string(t)
}

func (t Token) CoingeckoID() string {
	return coingeckoIDs[t]
}

func FromCoingeckoID(id string) (Token, bool) {
	for t, i := range coingeckoIDs {
		if i == id {
			return t, true
		}
	}

	return "", false
}

func MustFromCoingeckoID(id string) Token {
	t, ok := FromCoingeckoID(id)
	if !ok {
		panic("unknown coingecko id: " + id)
	}

	return t
}
