package tokenpricer

import (
	"context"
	"sync"

	"github.com/omni-network/omni/lib/tokens"
)

type Mock struct {
	mu     sync.RWMutex
	prices map[tokens.Asset]float64
}

var _ Pricer = (*Mock)(nil)

func NewMock(prices map[tokens.Asset]float64) *Mock {
	cloned := make(map[tokens.Asset]float64)
	for k, v := range prices {
		cloned[k] = v
	}

	return &Mock{prices: cloned}
}

func (m *Mock) USDPrice(_ context.Context, tkn tokens.Asset) (float64, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return m.prices[tkn], nil
}

func (m *Mock) USDPrices(_ context.Context, tkns ...tokens.Asset) (map[tokens.Asset]float64, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	resp := make(map[tokens.Asset]float64)
	for _, t := range tkns {
		resp[t] = m.prices[t]
	}

	return resp, nil
}

func (m *Mock) SetPrice(token tokens.Asset, price float64) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.prices[token] = price
}
