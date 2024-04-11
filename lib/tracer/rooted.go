package tracer

import (
	"context"
	"encoding/binary"
	"hash/fnv"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

// StartChainHeight returns a context and span rooted to the chain+height.
// This creates a new trace root and should generally only by xprovider or cprovider.
func StartChainHeight(ctx context.Context, chain string, height uint64, spanName string, opts ...trace.SpanStartOption) (context.Context, trace.Span) {
	// Deterministic TraceID for chain+height. So all traces across all instances are correlated.
	h := fnv.New128a()
	_, _ = h.Write([]byte(chain))
	_ = binary.Write(h, binary.BigEndian, height)

	var traceID trace.TraceID
	copy(traceID[:], h.Sum(nil))

	ctx, span := tracer.Start(RootedCtx(ctx, traceID), spanName, opts...)

	span.SetAttributes(attribute.String("chain", chain))
	span.SetAttributes(attribute.Int64("height", int64(height)))

	return ctx, span
}
