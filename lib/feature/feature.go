package feature

import (
	"context"

	"github.com/omni-network/omni/lib/log"
)

const (
	// FlagEVMStakingModule enables the wip EVM Staking Module feature.
	FlagEVMStakingModule Flag = "evm-staking-module"
)

var allFlags = map[Flag]bool{
	FlagEVMStakingModule: true,
}

// Flag is a feature flag.
type Flag string

// Enabled returns true if the flag is enabled in the context.
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

// enabled returns true if the given flag is enabled in the context.
func enabled(ctx context.Context, flag Flag) bool {
	flags, ok := ctx.Value(key{}).([]Flag)
	if !ok {
		return false
	}

	for _, f := range flags {
		if f == flag {
			return true
		}
	}

	return false
}
