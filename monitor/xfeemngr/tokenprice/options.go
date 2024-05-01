package tokenprice

import (
	"time"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/tokens"
	"github.com/omni-network/omni/monitor/xfeemngr/ticker"
)

type Opts struct {
	thresholdPct float64
	ticker       ticker.Ticker
	tokens       []tokens.Token
}

func defaultOpts() *Opts {
	return &Opts{
		thresholdPct: 0.1,
		ticker:       ticker.New(ticker.WithInterval(30 * time.Second)),
	}
}

func makeOpts(opts []func(*Opts)) (*Opts, error) {
	o := defaultOpts()
	for _, opt := range opts {
		opt(o)
	}

	if len(o.tokens) == 0 {
		return nil, errors.New("no tokens provided")
	}

	return o, nil
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

func WithTokens(t ...tokens.Token) func(*Opts) {
	return func(o *Opts) {
		o.tokens = t
	}
}
