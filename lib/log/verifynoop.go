//go:build !verify_logs

package log

// verifyAttrs is a noop default builds.
// This ensures no performance penalty for default builds.
func verifyAttrs([]any) {}
