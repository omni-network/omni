package tracer

import (
	"context"
	"encoding/binary"
	"hash/fnv"

	"github.com/omni-network/omni/lib/netconf"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

// StartChainHeight returns a context and span rooted to the network+network.Version+chain+height.
// This creates a new trace root and should generally only by xprovider or cprovider.
//
//nolint:spancheck // False positive.
func StartChainHeight(ctx context.Context, network netconf.ID, chain string, height uint64, spanName string, opts ...trace.SpanStartOption) (context.Context, trace.Span) {
	// Deterministic TraceID for network+network.Version+chain+height.
	// So all traces of the same block across all apps/instances of the same network are correlated.
	// Note this only works for protected networks with consistent versions.
	// Ephemeral network traces will not be correlated.

	h := fnv.New128a()
	_, _ = h.Write([]byte(network.String()))
	_, _ = h.Write([]byte(network.Static().Version))
	_, _ = h.Write([]byte(chain))
	_ = binary.Write(h, binary.BigEndian, height)

	var traceID trace.TraceID
	copy(traceID[:], h.Sum(nil))

	ctx, span := tracer.Start(RootedCtx(ctx, traceID), spanName, opts...)

	span.SetAttributes(attribute.String("chain", chain))
	span.SetAttributes(attribute.Int64("height", int64(height)))

	return ctx, span
}
