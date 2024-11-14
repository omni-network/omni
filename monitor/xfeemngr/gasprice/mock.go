package gasprice

import (
	"context"
	"math/big"
	"sync"

	"github.com/omni-network/omni/lib/umath"

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

	return umath.NewBigInt(m.price), nil
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

type MockBuffer struct {
	mu     sync.RWMutex
	prices map[uint64]uint64
}

var _ Buffer = (*MockBuffer)(nil)

func NewMockBuffer() *MockBuffer {
	return &MockBuffer{
		mu:     sync.RWMutex{},
		prices: make(map[uint64]uint64),
	}
}

func (b *MockBuffer) SetGasPrice(chainID uint64, price uint64) {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.prices[chainID] = price
}

func (b *MockBuffer) GasPrice(chainID uint64) uint64 {
	b.mu.RLock()
	defer b.mu.RUnlock()

	return b.prices[chainID]
}

func (*MockBuffer) Stream(context.Context) {
	// no-op
}
