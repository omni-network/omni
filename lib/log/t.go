package log

import (
	"strings"
	"testing"
)

// testWriter is a simple io.Writer that logs to the testing.TB.
// It is defined in t.go to mitigate t.Log noise which adds this filename to all logs.
type testWriter struct {
	t *testing.T
}

func (w testWriter) Write(b []byte) (int, error) {
	w.t.Log(strings.TrimSpace(string(b)))
	return len(b), nil
}
