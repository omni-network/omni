package flowgen

import (
	"context"

	"github.com/omni-network/omni/lib/log"

	"github.com/robfig/cron/v3"
)

// cronLogger is a logger for cron.Logger for omni logs.
type cronLogger struct {
	info  func(msg string, attrs ...any)
	error func(err error, msg string, attrs ...any)
}

var _ cron.Logger = cronLogger{}

func (l cronLogger) Info(msg string, args ...any) {
	l.info(msg, args...)
}

func (l cronLogger) Error(err error, msg string, args ...any) {
	l.error(err, msg, args...)
}

func newCronLogger(ctx context.Context) cron.Logger {
	return cronLogger{
		info:  func(msg string, attrs ...any) { log.Info(ctx, msg, attrs...) },
		error: func(err error, msg string, attrs ...any) { log.Error(ctx, msg, err, attrs...) },
	}
}
