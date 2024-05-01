package tokens

import (
	"context"
	"sync"
)

type MockPricer struct {
	mu     sync.RWMutex
	prices map[Token]float64
}

var _ Pricer = (*MockPricer)(nil)

func NewMockPricer(prices map[Token]float64) *MockPricer {
	cloned := make(map[Token]float64)
	for k, v := range prices {
		cloned[k] = v
	}

	return &MockPricer{prices: cloned}
}

func (m *MockPricer) Price(_ context.Context, tkns ...Token) (map[Token]float64, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	resp := make(map[Token]float64)
	for _, t := range tkns {
		resp[t] = m.prices[t]
	}

	return resp, nil
}

func (m *MockPricer) SetPrice(token Token, price float64) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.prices[token] = price
}
