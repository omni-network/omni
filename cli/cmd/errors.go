package cmd

import (
	"fmt"
	"strings"
)

// cliError is a custom error type for CLI usage that adds a helpful suggestion.
// It wraps the Msg like a normal error
// It doesn't support attributes yet.
type cliError struct {
	Msg     string
	Suggest string
	Attrs   []any // Attributes not yet implemented.
}

func (e *cliError) Error() string {
	var sb strings.Builder
	_, _ = sb.WriteString(e.Msg)
	if e.Suggest != "" {
		_, _ = sb.WriteString("\n")
		_, _ = sb.WriteString("🤔 " + e.Suggest)
	}

	return sb.String()
}

func (e *cliError) Wrap(msg string, attrs ...any) error {
	e.Msg = fmt.Sprintf("%s: %s", msg, e.Msg)
	e.Attrs = append(e.Attrs, attrs...)

	return e
}
