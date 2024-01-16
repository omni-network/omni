//go:build verify_logs

package log

import (
	"log/slog"
	"regexp"
	"time"
)

const badKey = "!BADKEY"

var snakeRegex = regexp.MustCompile(`^[a-z0-9_]+$`)

// verifyAttrs verifies that the log attribute keys are valid and in snake case.
// It is only enabled via the verify_logs build tag to avoid performance penalty.
func verifyAttrs(attrs []any) {
	r := slog.NewRecord(time.Now(), 0, "", 0)
	r.Add(attrs...)
	r.Attrs(func(attr slog.Attr) bool {
		if attr.Key == "" {
			panic("empty log attribute key")
		} else if attr.Key == badKey {
			panic("bad log attribute key:" + attr.String())
		} else if !snakeRegex.MatchString(attr.Key) {
			panic("log attribute key not snake case: " + attr.String())
		}

		return true
	})
}
