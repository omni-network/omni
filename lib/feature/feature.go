package feature

import (
	"context"
	"slices"
	"sync"

	"github.com/omni-network/omni/lib/log"
)

const (
	// FlagFuzzOctane enables fuzz testing of octane.
	FlagFuzzOctane Flag = "fuzz-octane"
)

// enabledFlags holds all globally enabled feature flags. The reason for having it is that
// we want to use feature flags across the entire code base. However, CometBFT doesn't allow
// specifying a root context, and the Cosmos SDK doesn't actually use the context provided in ABCI,
// so that feature flags cannot be shared with modules via context.
var enabledFlags sync.Map

var allFlags = map[Flag]bool{
	FlagFuzzOctane: true,
}

// Flag is a feature flag.
type Flag string

// Enabled returns true if the flag is enabled in the context or globally.
func (f Flag) Enabled(ctx context.Context) bool {
	return enabled(ctx, f)
}

type key struct{}

// WithFlags returns a copy of the context with the given flags enabled.
// Note that this should only be called once on app startup.
// Multiple calls will overwrite the existing flags.
func WithFlags(ctx context.Context, flags Flags) context.Context {
	var enabled []Flag
	for _, f := range flags.Typed() {
		if !allFlags[f] {
			// Don't error, just log, this ensures flags can be safely removed.
			log.Warn(ctx, "Ignoring unknown feature flag", nil, "flag", f)
			continue
		}
		enabled = append(enabled, f)
	}

	if len(enabled) == 0 {
		return ctx
	}

	log.Info(ctx, "Enabling feature flags", "flags", enabled)

	return context.WithValue(ctx, key{}, enabled)
}

// WithFlag is a convenience function for testing to enable single flags.
func WithFlag(ctx context.Context, flag Flag) context.Context {
	return WithFlags(ctx, Flags{string(flag)})
}

// SetGlobals enables all given flags globally.
func SetGlobals(flags Flags) {
	for _, flag := range flags {
		enabledFlags.Store(Flag(flag), true)
	}
}

// enabled returns true if the given flag is enabled globally or in the context.
func enabled(ctx context.Context, flag Flag) bool {
	if _, ok := enabledFlags.Load(flag); ok {
		return true
	}

	flags, ok := ctx.Value(key{}).([]Flag)
	if !ok {
		return false
	}

	return slices.Contains(flags, flag)
}
