package tokens

import (
	"fmt"
	"math/big"
	"sort"
	"strconv"

	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/errors"
)

// Asset represents the canonical definition of a token,
// independent of any specific blockchain.
type Asset struct {
	Symbol      string
	Name        string
	Decimals    uint
	CoingeckoID string // See `Info > API ID` on asset's coingecko page.
}

var (
	OMNI = Asset{
		Symbol:      "OMNI",
		Name:        "Omni Network",
		Decimals:    18,
		CoingeckoID: "omni-network",
	}

	ETH = Asset{
		Symbol:      "ETH",
		Name:        "Ether",
		Decimals:    18,
		CoingeckoID: "ethereum",
	}

	WETH = Asset{
		Symbol:      "WETH",
		Name:        "Wrapped Ether",
		Decimals:    18,
		CoingeckoID: "weth",
	}

	STETH = Asset{
		Symbol:      "stETH",
		Name:        "Lido Staked Ether",
		Decimals:    18,
		CoingeckoID: "staked-ether",
	}

	WSTETH = Asset{
		Symbol:      "wstETH",
		Name:        "Wrapped Staked Ether",
		Decimals:    18,
		CoingeckoID: "wrapped-steth",
	}

	USDC = Asset{
		Symbol:      "USDC",
		Name:        "USD Coin",
		Decimals:    6,
		CoingeckoID: "usd-coin",
	}

	USDT = Asset{
		Symbol:      "USDT",
		Name:        "Tether",
		Decimals:    6,
		CoingeckoID: "tether",
	}

	USDT0 = Asset{
		Symbol:      "USDT0",
		Name:        "Tether Zero",
		Decimals:    6,
		CoingeckoID: "usdt0",
	}

	BNB = Asset{
		Symbol:      "BNB",
		Name:        "BNB",
		Decimals:    18,
		CoingeckoID: "binancecoin",
	}

	POL = Asset{
		Symbol:      "POL",
		Name:        "POL",
		Decimals:    18,
		CoingeckoID: "polygon-ecosystem-token",
	}

	HYPE = Asset{
		Symbol:      "HYPE",
		Name:        "Hyperliquid",
		Decimals:    18,
		CoingeckoID: "hyperliquid",
	}

	MNT = Asset{
		Symbol:      "MNT",
		Name:        "Mantle",
		Decimals:    18,
		CoingeckoID: "mantle",
	}

	BERA = Asset{
		Symbol:      "BERA",
		Name:        "Berachain",
		Decimals:    18,
		CoingeckoID: "berachain-bera",
	}

	NOM = Asset{
		Symbol:      "NOM",
		Name:        "Nomina",
		Decimals:    18,
		CoingeckoID: "nomina",
	}

	PLUME = Asset{
		Symbol:      "PLUME",
		Name:        "Plume",
		Decimals:    18,
		CoingeckoID: "plume",
	}

	SOL = Asset{
		Symbol:      "SOL",
		Name:        "Solana",
		Decimals:    9,
		CoingeckoID: "solana",
	}

	METH = Asset{
		Symbol:      "mETH",
		Name:        "mETH",
		Decimals:    18,
		CoingeckoID: "mantle-staked-ether",
	}
)

func AssetBySymbol(symbol string) (Asset, error) {
	for _, t := range tokens {
		if t.Symbol == symbol {
			return t.Asset, nil
		}
	}

	return Asset{}, errors.New("unknown asset", "symbol", symbol)
}

// UniqueAssets returns all unique assets of all tokens.
func UniqueAssets() []Asset {
	uniq := make(map[Asset]bool)

	for _, t := range tokens {
		uniq[t.Asset] = true
	}

	var resp []Asset
	for a := range uniq {
		resp = append(resp, a)
	}

	sort.Slice(resp, func(i, j int) bool {
		return resp[i].Symbol < resp[j].Symbol
	})

	return resp
}

func (a Asset) String() string {
	return a.Symbol
}

// F64ToAmt returns the float64 in primary units as a big.Int in base units.
// E.g. 1.1 (ether) is returned as 1.1 * 10^18 (wei).
func (a Asset) F64ToAmt(n float64) *big.Int {
	if a.Decimals == 6 {
		return bi.Dec6(n)
	}

	return bi.Ether(n)
}

// AmtToF64 returns amount in base units as a float64 in primary units.
// E.g. 1.1 * 10^18 (wei) is returned as 1.1 (ether).
func (a Asset) AmtToF64(n *big.Int) float64 {
	return bi.ToF64(n, a.Decimals)
}

// FormatAmt returns a string representation of the provided base unit amount
// in primary units with the token symbol appended. Note all decimals are shown.
func (a Asset) FormatAmt(n *big.Int) string {
	if n == nil {
		return "nil"
	}

	return fmt.Sprintf("%s %s",
		strconv.FormatFloat(a.AmtToF64(n), 'f', -1, 64), // Use FormatFloat 'f' instead of %f since it avoids trailing zeros
		a.Symbol,
	)
}
