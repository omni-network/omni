package tokenpricer

import (
	"context"
	"math/big"
	"sync"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/tokens"
)

type Mock struct {
	mu     sync.RWMutex
	prices map[pair]*big.Rat
}

var _ Pricer = (*Mock)(nil)

// NewDevnetMock only supports OMNI/wstETH/ETH/USDC/NOM swaps.
func NewDevnetMock() *Mock {
	m := &Mock{prices: make(map[pair]*big.Rat)}
	m.SetPrice(tokens.WSTETH, tokens.USDC, big.NewRat(4000, 1)) // 1 WSTETH = 4000 USDC
	m.SetPrice(tokens.ETH, tokens.USDC, big.NewRat(3000, 1))    // 1 ETH = 3000 USDC
	m.SetPrice(tokens.OMNI, tokens.USDC, big.NewRat(5, 1))      // 1 OMNI = 5 USDC
	m.SetPrice(tokens.NOM, tokens.USDC, big.NewRat(5, 75))      // 1 NOM = 5/75 USDC = OMNI/75

	return m
}

func NewUSDMock(prices map[tokens.Asset]float64) *Mock {
	cloned := make(map[pair]*big.Rat)
	for k, v := range prices {
		cloned[pair{Base: k, Quote: tokens.USDC}] = new(big.Rat).SetFloat64(v)
	}

	return &Mock{prices: cloned}
}

// Price returns the price of the base asset denominated in the quote asset.
// Note that for canonical solver prices, base=deposit and quote=expense.
func (m *Mock) Price(_ context.Context, base, quote tokens.Asset) (*big.Rat, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	price, ok := m.prices[pair{Base: base, Quote: quote}]
	if ok {
		return price, nil
	}

	// Try via USDC
	usdBase, ok := m.prices[pair{Base: base, Quote: tokens.USDC}]
	if !ok {
		return nil, errors.New("mock price not found", "base", base, "quote", quote)
	}

	usdQuote, ok := m.prices[pair{Base: quote, Quote: tokens.USDC}]
	if !ok {
		return nil, errors.New("mock price not found", "base", base, "quote", quote)
	}

	return new(big.Rat).Quo(usdBase, usdQuote), nil
}

func (m *Mock) USDPrice(ctx context.Context, tkn tokens.Asset) (float64, error) {
	prices, err := m.USDPrices(ctx, tkn)
	if err != nil {
		return 0, errors.Wrap(err, "get price")
	}

	return prices[tkn], nil
}

func (m *Mock) USDPriceRat(ctx context.Context, tkn tokens.Asset) (*big.Rat, error) {
	prices, err := m.USDPricesRat(ctx, tkn)
	if err != nil {
		return nil, errors.Wrap(err, "get price")
	}

	return prices[tkn], nil
}

// USDPrices returns the price of each coin in USD.
func (m *Mock) USDPrices(ctx context.Context, assets ...tokens.Asset) (map[tokens.Asset]float64, error) {
	rats, err := m.USDPricesRat(ctx, assets...)
	if err != nil {
		return nil, errors.Wrap(err, "get price")
	}

	prices := make(map[tokens.Asset]float64)
	for asset, price := range rats {
		f, _ := price.Float64()
		prices[asset] = f
	}

	return prices, nil
}

func (m *Mock) USDPricesRat(_ context.Context, assets ...tokens.Asset) (map[tokens.Asset]*big.Rat, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	resp := make(map[tokens.Asset]*big.Rat)
	for _, t := range assets {
		var ok bool
		resp[t], ok = m.prices[pair{Base: t, Quote: tokens.USDC}]
		if !ok {
			return nil, errors.New("mock usd price not found", "token", t)
		}
	}

	return resp, nil
}

func (m *Mock) SetUSDPrice(token tokens.Asset, price float64) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.prices[pair{Base: token, Quote: tokens.USDC}] = new(big.Rat).SetFloat64(price)
}

func (m *Mock) SetPrice(base, quote tokens.Asset, price *big.Rat) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.prices[pair{Base: base, Quote: quote}] = price
	m.prices[pair{Base: quote, Quote: base}] = new(big.Rat).Inv(price)

	m.prices[pair{Base: base, Quote: base}] = big.NewRat(1, 1)
	m.prices[pair{Base: quote, Quote: quote}] = big.NewRat(1, 1)
}
