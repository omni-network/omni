package app

import (
	"context"

	"github.com/omni-network/omni/lib/tracer"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

// startTrace returns a context and span rooted to orderID.
func startTrace(ctx context.Context, srcChain string, orderID OrderID, opts ...trace.SpanStartOption) (context.Context, trace.Span) {
	var traceID trace.TraceID
	copy(traceID[:], orderID[:])

	ctx, span := tracer.Start(tracer.RootedCtx(ctx, traceID), "solver/order", opts...)

	span.SetAttributes(attribute.String("src_chain", srcChain))
	span.SetAttributes(attribute.String("order_id", orderID.Hex()))

	return ctx, span
}
