package comet

import (
	"context"
	"strings"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"

	cmtlog "github.com/cometbft/cometbft/libs/log"
)

var _ cmtlog.Logger = (*logger)(nil)

const (
	levelError = iota + 1
	levelInfo
	levelDebug
)

// levels maps strings to numbers for easy comparison.
//
//nolint:gochecknoglobals // Global is ok here.
var levels = map[string]int{
	"error": levelError,
	"info":  levelInfo,
	"debug": levelDebug,
}

// dropCometDebugs is a map of cometBFT debug messages that should be dropped.
// These are super noisy and not useful.
//
//nolint:gochecknoglobals // Static mapping
var dropCometDebugs = map[string]bool{
	"Read PacketMsg":       true,
	"Received bytes":       true,
	"Send":                 true,
	"Receive":              true,
	"Sending vote message": true,
	"Flush":                true,
	"setHasVote":           true,
	"TrySend":              true,
}

// logger implements cmtlog.Logger by using the omni logging pattern.
// Comet log level is controlled separately in config.toml, since comet logs are very noisy.
type logger struct {
	ctx   context.Context //nolint:containedctx // This is a wrapper around the omni logger which is context based.
	level int
}

func NewLogger(ctx context.Context, levelStr string) (cmtlog.Logger, error) {
	level, ok := levels[strings.ToLower(levelStr)]
	if !ok {
		return logger{}, errors.New("invalid comet log level", "level", levelStr)
	}

	return logger{
		ctx:   log.WithSkip(ctx, 4), // Skip this logger.
		level: level,
	}, nil
}

func (c logger) Debug(msg string, keyvals ...any) {
	if c.level < levelDebug {
		return
	} else if dropCometDebugs[msg] {
		return
	}

	log.Debug(c.ctx, msg, keyvals...)
}

func (c logger) Info(msg string, keyvals ...any) {
	if c.level < levelInfo {
		return
	}
	log.Info(c.ctx, msg, keyvals...)
}

func (c logger) Error(msg string, keyvals ...any) {
	if c.level < levelError {
		return
	}

	keyvals, err := splitOutError(keyvals)

	log.Error(c.ctx, msg, err, keyvals...)
}

func (c logger) With(keyvals ...any) cmtlog.Logger { //nolint:ireturn // This signature is required by interface.
	return logger{
		ctx:   log.WithCtx(c.ctx, keyvals...),
		level: c.level,
	}
}

// splitOutError splits the keyvals into a slice of keyvals without the error and the error.
func splitOutError(keyvals []any) ([]any, error) {
	var remaining []any
	var err error
	for i := 0; i < len(keyvals)-1; i += 2 {
		if keyErr, ok := keyvals[i+1].(error); ok {
			err = keyErr
		} else {
			remaining = append(remaining, keyvals[i], keyvals[i+1])
		}
	}

	return remaining, err
}
