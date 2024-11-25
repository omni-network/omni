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

var (
	gasTokenIDs = map[Token]uint8{
		OMNI: 1,
		ETH:  2,
	}
)

func (t Token) String() string {
	return string(t)
}

func (t Token) CoingeckoID() string {
	return coingeckoIDs[t]
}

func (t Token) GasTokenID() (uint8, bool) {
	id, ok := gasTokenIDs[t]
	return id, ok
}

func GasTokenIDs() map[Token]uint8 {
	result := make(map[Token]uint8, len(gasTokenIDs))
	for k, v := range gasTokenIDs {
		result[k] = v
	}

	return result
}

func FromCoingeckoID(id string) (Token, bool) {
	for t, i := range coingeckoIDs {
		if i == id {
			return t, true
		}
	}

	return "", false
}

func FromGasTokenID(id uint8) (Token, bool) {
	for t, i := range gasTokenIDs {
		if i == id {
			return t, true
		}
	}

	return "", false
}
