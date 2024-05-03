package ticker

import (
	"context"
	"sync"
)

var _ Ticker = &MockTicker{}

type fnWithCtx struct {
	//nolint:containedctx // Mock object, needed to keep ctx passed to Go with the function
	ctx context.Context
	fn  func(context.Context)
}

type MockTicker struct {
	mu  sync.RWMutex
	fns []fnWithCtx
}

func NewMock() *MockTicker {
	return &MockTicker{fns: []fnWithCtx{}}
}

func (t *MockTicker) Go(ctx context.Context, f func(context.Context)) {
	t.mu.Lock()
	defer t.mu.Unlock()

	t.fns = append(t.fns, fnWithCtx{ctx: ctx, fn: f})
}

func (t *MockTicker) Tick() {
	t.mu.RLock()
	defer t.mu.RUnlock()

	for _, f := range t.fns {
		f.fn(f.ctx)
	}
}
