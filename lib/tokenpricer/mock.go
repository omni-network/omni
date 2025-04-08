package tokenpricer

import (
	"context"
	"sync"

	"github.com/omni-network/omni/lib/tokenmeta"
)

type Mock struct {
	mu     sync.RWMutex
	prices map[tokenmeta.Meta]float64
}

var _ Pricer = (*Mock)(nil)

func NewMock(prices map[tokenmeta.Meta]float64) *Mock {
	cloned := make(map[tokenmeta.Meta]float64)
	for k, v := range prices {
		cloned[k] = v
	}

	return &Mock{prices: cloned}
}

func (m *Mock) Price(_ context.Context, tkn tokenmeta.Meta) (float64, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return m.prices[tkn], nil
}

func (m *Mock) Prices(_ context.Context, tkns ...tokenmeta.Meta) (map[tokenmeta.Meta]float64, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	resp := make(map[tokenmeta.Meta]float64)
	for _, t := range tkns {
		resp[t] = m.prices[t]
	}

	return resp, nil
}

func (m *Mock) SetPrice(token tokenmeta.Meta, price float64) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.prices[token] = price
}
