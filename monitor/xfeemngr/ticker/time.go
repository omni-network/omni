package ticker

import (
	"context"
	"time"
)

var _ Ticker = TimeTicker{}

type TimeTicker struct {
	interval time.Duration
}

func New(interval time.Duration) TimeTicker {
	return TimeTicker{interval: interval}
}

func (t TimeTicker) Go(ctx context.Context, f func(context.Context)) {
	tick := func() {
		interval := t.interval
		ticker := time.NewTicker(interval)
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
			case <-ticker.C:
				func() {
					ctx, cancel := context.WithTimeout(ctx, interval)
					defer cancel()
					f(ctx)
				}()
			}
		}
	}

	go tick()
}
