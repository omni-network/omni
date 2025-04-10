package tokenpricer

import (
	"context"
	"sync"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/tokens"
)

type Mock struct {
	mu     sync.RWMutex
	prices map[pair]float64
}

var _ Pricer = (*Mock)(nil)

// NewDevnetMock only supports OMNI/wstETH/ETH/USDC swaps.
func NewDevnetMock() *Mock {
	m := &Mock{prices: make(map[pair]float64)}
	m.SetPrice(tokens.WSTETH, tokens.USDC, 4000)
	m.SetPrice(tokens.ETH, tokens.USDC, 3000)
	m.SetPrice(tokens.OMNI, tokens.USDC, 5)

	return m
}

func NewUSDMock(prices map[tokens.Asset]float64) *Mock {
	cloned := make(map[pair]float64)
	for k, v := range prices {
		cloned[pair{Base: k, Quote: tokens.USDC}] = v
	}

	return &Mock{prices: cloned}
}

func (m *Mock) Price(_ context.Context, base, quote tokens.Asset) (float64, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	price, ok := m.prices[pair{Base: base, Quote: quote}]
	if ok {
		return price, nil
	}

	// Try via USDC
	usdBase, ok := m.prices[pair{Base: base, Quote: tokens.USDC}]
	if !ok {
		return 0, errors.New("mock price not found", "base", base, "quote", quote)
	}

	usdQuote, ok := m.prices[pair{Base: quote, Quote: tokens.USDC}]
	if !ok {
		return 0, errors.New("mock price not found", "base", base, "quote", quote)
	}

	return usdBase / usdQuote, nil
}

func (m *Mock) USDPrice(_ context.Context, tkn tokens.Asset) (float64, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	price, ok := m.prices[pair{Base: tkn, Quote: tokens.USDC}]
	if !ok {
		return 0, errors.New("mock usd price not found", "token", tkn)
	}

	return price, nil
}

func (m *Mock) USDPrices(_ context.Context, tkns ...tokens.Asset) (map[tokens.Asset]float64, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	resp := make(map[tokens.Asset]float64)
	for _, t := range tkns {
		var ok bool
		resp[t], ok = m.prices[pair{Base: t, Quote: tokens.USDC}]
		if !ok {
			return nil, errors.New("mock price not found", "token", t)
		}
	}

	return resp, nil
}

func (m *Mock) SetUSDPrice(token tokens.Asset, price float64) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.prices[pair{Base: token, Quote: tokens.USDC}] = price
}

func (m *Mock) SetPrice(base, quote tokens.Asset, price float64) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.prices[pair{Base: base, Quote: quote}] = price
	m.prices[pair{Base: quote, Quote: base}] = 1 / price
}
