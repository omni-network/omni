package log

import (
	"io"
	"log/slog"

	"github.com/muesli/termenv"
)

// LoggersForT returns a map of loggers for testing.
func LoggersForT() map[string]func(io.Writer) *slog.Logger {
	testOpts := func(w io.Writer) func(*options) {
		return func(o *options) {
			o.Writer = w
			o.Test = true
			o.Level = slog.LevelDebug
			o.Color = termenv.Ascii
		}
	}

	resp := make(map[string]func(w io.Writer) *slog.Logger)
	for name, fn := range loggerFuncs {
		resp[name] = func(w io.Writer) *slog.Logger {
			return fn(testOpts(w))
		}
	}

	return resp
}
