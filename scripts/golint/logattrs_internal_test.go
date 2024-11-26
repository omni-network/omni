package main

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestLogAttrs(t *testing.T) {
	t.Parallel()

	pkgMain := `package main

import (
	"context"
	"main/log"
	"main/errors"
	"log/slog"
	"strings"
)

func main() {
	ctx := context.Background()

	log.Debug(ctx, "Good, no attrs")
	log.Debug(ctx, "Good, valid attrs", "foo", "bar", "snake_case", "qux")
	log.Debug(ctx, "Bad, invalid attrs", "foo", "bar", "camelCase", "qux") // want "log/error attribute key must be snake_case"

	err := errors.New("fine", "foo", "bar", "snake_case", "qux")
	err = errors.Wrap(err, "bad", "foo",
		"bar", "camelCase", "qux") // want "log/error attribute key must be snake_case"
	log.Warn(ctx, "no extra attrs", err)
	log.Warn(ctx, "good attrs", err, "foo", "bar", "snake_case", "Values don't matter'")
	log.Error(ctx, "bad attrs", err, "camelCase", "qux") // want "log/error attribute key must be snake_case"

	log.InfoErr(ctx, "Good, no attrs", err)
	log.InfoErr(ctx, "Good, valid attrs", err, "foo", "bar", "snake_case", "qux")
	log.DebugErr(ctx, "Bad, invalid attrs", err, "foo", "bar", "camelCase", "qux") // want "log/error attribute key must be snake_case"
	log.DebugErr(ctx, "Bad, nil error", nil, "foo", "bar") // want "info/debug-err called with nil error"
	log.InfoErr(ctx, "Bad, nil error", nil) // want "info/debug-err called with nil error"

	log.Debug(ctx, "Bad, err attr", "error", err) // want "error attributes not allowed"
	log.Info(ctx, "Bad, err attr", "foo", "bar", "err", err) // want "error attributes not allowed"
	log.Warn(ctx, "Bad, err attr", nil, "foo", "bar", "err", err) // want "error attributes not allowed"
    err = errors.New("bad, err attr", "foo", "bar", "err", err) // want "error attributes not allowed"

	err = errors.Wrap(err, "wrap", nil) // want "bad log/error attribute key"
	log.Debug(ctx, "Bad attr key", err) // want "bad log/error attribute key"
	log.Info(ctx, "Bad attr key", "key", "value", -1) // want "bad log/error attribute key"
	log.Warn(ctx, "Bad attr key", nil, struct{}{}) // want "bad log/error attribute key"


	log.Debug(ctx, "slog attr ok", slog.String("notChecked", ""))
	log.Info(ctx, "Ignore []any...", []any{1,2,"three"}...)
	var key string
	log.Info(ctx, "Ignore string vars", key, "value")

	tup := struct{ Key string }{""}
	log.Info(ctx, "Ignore string fields", tup.Key, "value")
	log.Info(ctx, "Ignore function calls", strings.ToUpper("ok"), "value")
}
`

	pkgLog := `package log

import (
	"context"
)

func Debug(context.Context, string, ...any) {}
func DebugErr(context.Context, string, error, ...any) {}
func Info(context.Context, string, ...any) {}
func InfoErr(context.Context, string, error, ...any) {}
func Warn(context.Context, string, error, ...any) {}
func Error(context.Context, string, error, ...any) {}
`

	pkgError := `package errors

func Wrap(err error, _ string, _ ...any) error { return err }
func New(string, ...any) error { return nil }
`

	dir, cleanup, err := analysistest.WriteFiles(map[string]string{
		"main/main.go":          pkgMain,
		"main/log/log.go":       pkgLog,
		"main/errors/errors.go": pkgError,
	})
	if err != nil {
		t.Fatal(err)
	}
	defer cleanup()
	_ = analysistest.Run(t, dir, logAttrsAnalyzer, "main")
}
