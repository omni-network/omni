package relayer

import (
	"context"

	ethlog "github.com/ethereum/go-ethereum/log"
	"github.com/omni-network/omni/lib/log"
	"golang.org/x/exp/slog"
)

var _ ethlog.Logger = (*ethLogger)(nil)

type ethLogger struct {
	ctx context.Context //nolint:containedctx // This is a wrapper around the omni logger which is context based.
}

func (e ethLogger) With(ctx ...interface{}) ethlog.Logger {
	return ethLogger{
		ctx: log.WithCtx(e.ctx, ctx...),
	}
}

func (e ethLogger) New(ctx ...interface{}) ethlog.Logger {
	return ethLogger{
		ctx: log.WithCtx(e.ctx, ctx...),
	}
}

func (e ethLogger) Log(level slog.Level, msg string, ctx ...interface{}) {
	e.Write(level, msg, ctx...)
}

func (e ethLogger) Trace(msg string, ctx ...interface{}) {
	log.Debug(e.ctx, msg, ctx...)
}

func (e ethLogger) Debug(msg string, ctx ...interface{}) {
	log.Debug(e.ctx, msg, ctx...)
}

func (e ethLogger) Info(msg string, ctx ...interface{}) {
	log.Info(e.ctx, msg, ctx...)
}

func (e ethLogger) Warn(msg string, ctx ...interface{}) {
	log.Warn(e.ctx, msg, nil, ctx...)
}

func (e ethLogger) Error(msg string, ctx ...interface{}) {
	log.Error(e.ctx, msg, nil, ctx...)
}

func (e ethLogger) Crit(msg string, ctx ...interface{}) {
	// I don't want to do os.exit here
	log.Error(e.ctx, msg, nil, ctx...)
}

func (e ethLogger) Write(level slog.Level, msg string, attrs ...any) {
	switch level {
	case slog.LevelInfo:
		e.Info(msg, attrs...)
	case slog.LevelWarn:
		e.Warn(msg, attrs...)
	case slog.LevelError:
		e.Error(msg, attrs...)
	case slog.LevelDebug:
		e.Debug(msg, attrs...)
	}
}

func (e ethLogger) Enabled(ctx context.Context, level slog.Level) bool {
	//TODO implement me
	panic("implement me")
}
