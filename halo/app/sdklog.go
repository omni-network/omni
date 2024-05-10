package app

import (
	"context"

	"github.com/omni-network/omni/lib/log"

	sdklog "cosmossdk.io/log"
)

var _ sdklog.Logger = sdkLogger{}

// dropCosmosDebugs is a map of cosmosSDK debug messages that should be dropped.
// These are super noisy and not useful.
//
//nolint:gochecknoglobals // Static mapping
var dropCosmosDebugs = map[string]bool{
	"recursiveRemove": true,
	"BATCH SAVE":      true,
	"SAVE TREE":       true,
}

// sdkLogger implements sdklog.Logger by using the omni logging pattern.
// Comet log level is controlled separately in config.toml, since comet logs are very noisy.
type sdkLogger struct {
	ctx context.Context //nolint:containedctx // This is a wrapper around the omni logger which is context based.
}

func newSDKLogger(ctx context.Context) sdkLogger {
	return sdkLogger{
		ctx: log.WithSkip(ctx, 4), // Skip this logger.
	}
}

func (c sdkLogger) Debug(msg string, keyvals ...any) {
	if dropCosmosDebugs[msg] {
		return
	}
	log.Debug(c.ctx, msg, keyvals...)
}

func (c sdkLogger) Info(msg string, keyvals ...any) {
	log.Info(c.ctx, msg, keyvals...)
}

func (c sdkLogger) Warn(msg string, keyvals ...any) {
	keyvals, err := splitOutError(keyvals)

	log.Warn(c.ctx, msg, err, keyvals...)
}

func (c sdkLogger) Error(msg string, keyvals ...any) {
	keyvals, err := splitOutError(keyvals)

	log.Error(c.ctx, msg, err, keyvals...)
}

func (c sdkLogger) With(keyVals ...any) sdklog.Logger {
	return sdkLogger{
		ctx: log.WithCtx(c.ctx, keyVals...),
	}
}

func (c sdkLogger) Impl() any {
	return c
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
