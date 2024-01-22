package log

import (
	"encoding/hex"
	"log/slog"
)

// Hex7 is a convenience function for a hex-encoded log attribute limited to first 7 chars for brevity.
// Note this is NOT 0x prefixed.
func Hex7(key string, value []byte) slog.Attr {
	h := hex.EncodeToString(value)

	const maxLen = 7
	if len(h) > maxLen {
		h = h[:maxLen]
	}

	return slog.String(key, h)
}
