package relayer

import (
	"context"
	"log/slog"

	"github.com/omni-network/omni/lib/log"

	ethlog "github.com/ethereum/go-ethereum/log"

	eslog "golang.org/x/exp/slog"
)

var _ ethlog.Logger = (*ethLogger)(nil)

type ethLogger struct {
	log *slog.Logger
}

// WrapLogger wraps the logger inside the context and returns a new logger compatible with the
// Ethereum logger interface.
func WrapLogger(ctx context.Context) ethlog.Logger {
	return &ethLogger{
		log: log.GetLogger(ctx),
	}
}

func (e ethLogger) With(ctx ...any) ethlog.Logger {
	return ethLogger{
		log: e.log.With(ctx...),
	}
}

func (e ethLogger) New(ctx ...any) ethlog.Logger {
	return &ethLogger{
		log: e.log.With(ctx...),
	}
}

func (e ethLogger) Log(level eslog.Level, msg string, ctx ...any) {
	e.Write(level, msg, ctx...)
}

func (e ethLogger) Trace(msg string, ctx ...any) {
	e.log.Debug(msg, ctx...)
}

func (e ethLogger) Debug(msg string, ctx ...any) {
	e.log.Debug(msg, ctx...)
}

func (e ethLogger) Info(msg string, ctx ...any) {
	e.log.Info(msg, ctx...)
}

func (e ethLogger) Warn(msg string, ctx ...any) {
	e.log.Warn(msg, ctx...)
}

func (e ethLogger) Error(msg string, ctx ...any) {
	e.log.Error(msg, ctx...)
}

func (e ethLogger) Crit(msg string, ctx ...any) {
	// I don't want to do os.exit here
	e.log.Error(msg, ctx...)
}

func (e ethLogger) Write(level eslog.Level, msg string, attrs ...any) {
	switch level {
	case eslog.LevelInfo:
		e.log.Info(msg, attrs...)
	case eslog.LevelWarn:
		e.log.Warn(msg, attrs...)
	case eslog.LevelError:
		e.log.Error(msg, attrs...)
	case eslog.LevelDebug:
		e.log.Debug(msg, attrs...)
	}
}

func (ethLogger) Enabled(context.Context, eslog.Level) bool {
	return true
}
