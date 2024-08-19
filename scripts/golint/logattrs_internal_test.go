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
}
`

	pkgLog := `package log

import (
	"context"
)

func Debug(context.Context, string, ...any) {}
func Info(context.Context, string, ...any) {}
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
