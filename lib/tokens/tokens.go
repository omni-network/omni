package tokens

import (
	"encoding/json"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/uni"

	"github.com/ethereum/go-ethereum/common"

	"github.com/gagliardetto/solana-go"

	_ "embed"
)

// NativeAddr is the "address" of the native token; the zero address.
var NativeAddr common.Address

type ChainClass string

const (
	ClassDevent  ChainClass = "devnet"
	ClassTestnet ChainClass = "testnet"
	ClassMainnet ChainClass = "mainnet"
)

// Token represents a deployed instance of an Asset on a specific blockchain.
// It includes ERC20 and native and solana tokens.
type Token struct {
	Asset
	ChainID    uint64
	ChainClass ChainClass
	Address    common.Address   // zero if native eth or solana
	SolAddress solana.PublicKey `json:"SolAddress,omitempty"` // zero if native sol or erc20
	IsMock     bool
}

func (t Token) IsNative() bool {
	return t.UniAddress().IsZero()
}

func (t Token) UniAddress() uni.Address {
	if uni.IsSolChain(t.ChainID) {
		return uni.SolAddress(t.SolAddress)
	}

	return uni.EthAddress(t.Address)
}

func (t Token) IsSol() bool {
	return t.UniAddress().IsSol()
}

func (t Token) IsEth() bool {
	return t.UniAddress().IsEth()
}

func (t Token) Is(asset Asset) bool {
	return t.Asset == asset
}

func (t Token) IsOMNI() bool {
	return t.Is(OMNI)
}

//go:embed tokens.json
var tokenJSON []byte

var tokens = func() []Token {
	// Load tokens from the embedded JSON file.
	// This ensures that this lib/tokens package is self-contained
	// and doesn't depend on external libraries which result in
	// cyclic dependencies.
	var tkns []Token
	if err := json.Unmarshal(tokenJSON, &tkns); err != nil { //nolint:musttag // Tags not required
		panic(errors.Wrap(err, "unmarshal tokens"))
	}

	return tkns
}()

func All() []Token {
	return tokens
}

// BySymbol returns the token with the given symbol and chain ID.
func BySymbol(chainID uint64, symbol string) (Token, bool) {
	for _, t := range tokens {
		if t.ChainID == chainID && t.Symbol == symbol {
			return t, true
		}
	}

	return Token{}, false
}

// ByAsset returns the token with the given Asset and chain ID.
func ByAsset(chainID uint64, asset Asset) (Token, bool) {
	for _, t := range tokens {
		if t.ChainID == chainID && t.Is(asset) {
			return t, true
		}
	}

	return Token{}, false
}

// Native is an alias for ByAddress with the native token address.
func Native(chainID uint64) (Token, bool) {
	return ByAddress(chainID, NativeAddr)
}

// ByAddress returns the token with the given address and chain ID.
func ByAddress(chainID uint64, addr common.Address) (Token, bool) {
	for _, t := range tokens {
		if t.ChainID == chainID && t.UniAddress().EqualsEth(addr) {
			return t, true
		}
	}

	return Token{}, false
}

// ByUniAddress returns the token with the given universal address and chain ID.
func ByUniAddress(chainID uint64, addr uni.Address) (Token, bool) {
	for _, t := range tokens {
		if t.ChainID == chainID && t.UniAddress().Equals(addr) {
			return t, true
		}
	}

	return Token{}, false
}

func ByChain(chainID uint64) []Token {
	var tkns []Token

	for _, t := range tokens {
		if t.ChainID == chainID {
			tkns = append(tkns, t)
		}
	}

	return tkns
}
