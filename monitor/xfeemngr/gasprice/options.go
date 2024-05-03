package gasprice

import (
	"time"

	"github.com/omni-network/omni/monitor/xfeemngr/ticker"
)

type Opts struct {
	thresholdPct float64
	ticker       ticker.Ticker
}

func defaultOpts() *Opts {
	return &Opts{
		thresholdPct: 0.1,
		ticker:       ticker.New(ticker.WithInterval(30 * time.Second)),
	}
}

func makeOpts(opts []func(*Opts)) *Opts {
	o := defaultOpts()
	for _, opt := range opts {
		opt(o)
	}

	return o
}

func WithTicker(t ticker.Ticker) func(*Opts) {
	return func(o *Opts) {
		o.ticker = t
	}
}

func WithThresholdPct(pct float64) func(*Opts) {
	return func(o *Opts) {
		o.thresholdPct = pct
	}
}
