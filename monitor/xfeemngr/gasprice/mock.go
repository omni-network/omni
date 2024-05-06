package gasprice

import (
	"context"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum"
)

type MockPricer struct {
	mu    sync.RWMutex
	price uint64
}

var _ ethereum.GasPricer = (*MockPricer)(nil)

func NewMockPricer(price uint64) *MockPricer {
	return &MockPricer{
		mu:    sync.RWMutex{},
		price: price,
	}
}

func (m *MockPricer) SuggestGasPrice(_ context.Context) (*big.Int, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return big.NewInt(int64(m.price)), nil
}

func (m *MockPricer) SetPrice(price uint64) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.price = price
}

func (m *MockPricer) Price() uint64 {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return m.price
}
