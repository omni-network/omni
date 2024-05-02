package ticker

import (
	"context"
)

type Ticker interface {
	Go(ctx context.Context, f func(context.Context))
}
