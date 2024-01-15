package app

import (
	"context"

	"github.com/omni-network/omni/lib/log"

	cmtlog "github.com/cometbft/cometbft/libs/log"
)

var _ cmtlog.Logger = (*cmtLogger)(nil)

// cmtLogger implements cmtlog.Logger by using the omni logging pattern.
type cmtLogger struct {
	ctx context.Context //nolint:containedctx // This is a wrapper around the omni logger which is context based.
}

func newCmtLogger(ctx context.Context) cmtLogger {
	return cmtLogger{
		ctx: ctx,
	}
}

func (c cmtLogger) Debug(msg string, keyvals ...any) {
	log.Debug(c.ctx, msg, keyvals...)
}

func (c cmtLogger) Info(msg string, keyvals ...any) {
	log.Info(c.ctx, msg, keyvals...)
}

func (c cmtLogger) Error(msg string, keyvals ...any) {
	log.Error(c.ctx, msg, nil, keyvals...)
}

func (c cmtLogger) With(keyvals ...any) cmtlog.Logger { //nolint:ireturn // This signature is required by interface.
	return cmtLogger{
		ctx: log.WithCtx(c.ctx, keyvals...),
	}
}
