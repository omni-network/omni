package relayer

import (
	"context"

	ethlog "github.com/ethereum/go-ethereum/log"
	"github.com/omni-network/omni/lib/log"
	"golang.org/x/exp/slog"
)

var _ ethlog.Logger = (*ethLogger)(nil)

type ethLogger struct {
	ctx   context.Context //nolint:containedctx // This is a wrapper around the omni logger which is context based.
	level slog.Level
}

func (e ethLogger) With(ctx ...interface{}) ethlog.Logger {
	return ethLogger{
		ctx:   log.WithCtx(e.ctx, ctx...),
		level: e.level,
	}
}

func (e ethLogger) New(ctx ...interface{}) ethlog.Logger {
	return e.With(ctx...)
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
	keyVals, err := splitOutError(ctx)
	log.Warn(e.ctx, msg, err, keyVals...)
}

func (e ethLogger) Error(msg string, ctx ...interface{}) {
	keyVals, err := splitOutError(ctx)
	log.Error(e.ctx, msg, err, keyVals...)
}

func (e ethLogger) Crit(msg string, ctx ...interface{}) {
	// I don't want to do os.exit here
	keyVals, err := splitOutError(ctx)
	log.Error(e.ctx, msg, err, keyVals...)
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

func (e ethLogger) Enabled(_ context.Context, level slog.Level) bool {
	return e.level < level
}

// splitOutError splits the keyvals into a slice of keyvals without the error and the error.
func splitOutError(keyvals []any) ([]any, error) {
	var remaining []any
	var err error
	for i := 0; i < len(keyvals); i += 2 {
		if keyErr, ok := keyvals[i+1].(error); ok {
			err = keyErr
		} else {
			remaining = append(remaining, keyvals[i], keyvals[i+1])
		}
	}

	return remaining, err
}
