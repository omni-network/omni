package contract

import (
	"context"
	"math/big"
	"sync"
)

type MockFeeOracleV1 struct {
	mu           sync.Mutex
	gasPriceOn   map[uint64]*big.Int
	toNativeRate map[uint64]*big.Int
}

var _ FeeOracleV1 = (*MockFeeOracleV1)(nil)

func NewMockFeeOracleV1() *MockFeeOracleV1 {
	return &MockFeeOracleV1{
		gasPriceOn:   make(map[uint64]*big.Int),
		toNativeRate: make(map[uint64]*big.Int),
	}
}

func (m *MockFeeOracleV1) SetGasPriceOn(_ context.Context, destChainID uint64, gasPrice *big.Int) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.gasPriceOn[destChainID] = gasPrice

	return nil
}

func (m *MockFeeOracleV1) SetToNativeRate(_ context.Context, destChainID uint64, rate *big.Int) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.toNativeRate[destChainID] = rate

	return nil
}

func (m *MockFeeOracleV1) GasPriceOn(_ context.Context, destChainID uint64) (*big.Int, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	gasPrice, ok := m.gasPriceOn[destChainID]
	if !ok {
		return big.NewInt(0), nil
	}

	return gasPrice, nil
}

func (m *MockFeeOracleV1) ToNativeRate(_ context.Context, destChainID uint64) (*big.Int, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	rate, ok := m.toNativeRate[destChainID]
	if !ok {
		return big.NewInt(0), nil
	}

	return rate, nil
}
