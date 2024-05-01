package ticker

import (
	"context"
	"sync"
)

var _ Ticker = &MockTicker{}

type MockTicker struct {
	mu  sync.RWMutex
	fns []func(context.Context)
}

func NewMock() *MockTicker {
	return &MockTicker{fns: []func(context.Context){}}
}

func (t *MockTicker) Go(_ context.Context, f func(context.Context)) {
	t.mu.Lock()
	defer t.mu.Unlock()

	t.fns = append(t.fns, f)
}

func (t *MockTicker) Tick(ctx context.Context) {
	t.mu.RLock()
	defer t.mu.RUnlock()

	for _, f := range t.fns {
		f(ctx)
	}
}
