package tokenpricer

import (
	"context"
	"sync"

	"github.com/omni-network/omni/lib/tokens"
)

type Mock struct {
	mu     sync.RWMutex
	prices map[pair]float64
}

var _ Pricer = (*Mock)(nil)

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

	return m.prices[pair{Base: base, Quote: quote}], nil
}

func (m *Mock) USDPrice(_ context.Context, tkn tokens.Asset) (float64, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return m.prices[pair{Base: tkn, Quote: tokens.USDC}], nil
}

func (m *Mock) USDPrices(_ context.Context, tkns ...tokens.Asset) (map[tokens.Asset]float64, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	resp := make(map[tokens.Asset]float64)
	for _, t := range tkns {
		resp[t] = m.prices[pair{Base: t, Quote: tokens.USDC}]
	}

	return resp, nil
}

func (m *Mock) SetUSDPrice(token tokens.Asset, price float64) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.prices[pair{Base: token, Quote: tokens.USDC}] = price
}
