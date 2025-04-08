package tokenprice

import (
	"context"
	"sync"

	"github.com/omni-network/omni/lib/tokenmeta"
)

type MockBuffer struct {
	mu     sync.RWMutex
	prices map[tokenmeta.Meta]float64
}

var _ Buffer = (*MockBuffer)(nil)

func NewMockBuffer() *MockBuffer {
	return &MockBuffer{
		mu:     sync.RWMutex{},
		prices: make(map[tokenmeta.Meta]float64),
	}
}

func (b *MockBuffer) SetPrice(token tokenmeta.Meta, price float64) {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.prices[token] = price
}

func (b *MockBuffer) Price(token tokenmeta.Meta) float64 {
	b.mu.RLock()
	defer b.mu.RUnlock()

	return b.prices[token]
}

func (*MockBuffer) Stream(context.Context) {
	// no-op
}
