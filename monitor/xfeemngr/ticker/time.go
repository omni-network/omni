package ticker

import (
	"context"
	"time"
)

var _ Ticker = TimeTicker{}

type TimeTicker struct {
	opts *Opts
}

type Opts struct {
	interval time.Duration
}

func New(opts ...func(*Opts)) TimeTicker {
	return TimeTicker{opts: makeOpts(opts)}
}

func WithInterval(interval time.Duration) func(*Opts) {
	return func(o *Opts) {
		o.interval = interval
	}
}

func defaultOpts() *Opts {
	return &Opts{
		interval: 30 * time.Second,
	}
}

func makeOpts(opts []func(*Opts)) *Opts {
	o := defaultOpts()
	for _, opt := range opts {
		opt(o)
	}

	return o
}

func (t TimeTicker) Go(ctx context.Context, f func(context.Context)) {
	tick := func() {
		interval := t.opts.interval
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
