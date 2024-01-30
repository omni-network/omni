package relayer

import (
	"context"

	ethlog "github.com/ethereum/go-ethereum/log"

	"golang.org/x/exp/slog"
)

var _ ethlog.Logger = (*ethLogger)(nil)

type ethLogger struct {
	log *slog.Logger
}

func (e ethLogger) With(ctx ...any) ethlog.Logger {
	return ethLogger{
		log: e.log.With(ctx...),
	}
}

// WrapLogger wraps an instance of [slog.Logger] returns a new logger compatiple with the Ethereum logger interaface.
func WrapLogger(l *slog.Logger) ethlog.Logger {
	return &ethLogger{
		log: l,
	}
}

func (e ethLogger) New(ctx ...any) ethlog.Logger {
	return &ethLogger{
		log: e.log.With(ctx...),
	}
}

func (e ethLogger) Log(level slog.Level, msg string, ctx ...any) {
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

func (e ethLogger) Write(level slog.Level, msg string, attrs ...any) {
	switch level {
	case slog.LevelInfo:
		e.log.Info(msg, attrs...)
	case slog.LevelWarn:
		e.log.Warn(msg, attrs...)
	case slog.LevelError:
		e.log.Error(msg, attrs...)
	case slog.LevelDebug:
		e.log.Debug(msg, attrs...)
	}
}

func (ethLogger) Enabled(context.Context, slog.Level) bool {
	return true
}
