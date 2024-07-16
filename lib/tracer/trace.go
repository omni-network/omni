// Package tracer provides a global OpenTelemetry tracer.
package tracer

import (
	"context"
	"io"
	"sync"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/netconf"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
	"go.opentelemetry.io/otel/trace"
	"go.opentelemetry.io/otel/trace/noop"
)

var (
	// tracer is the global app level tracer, it defaults to a noop tracer.
	tracer   = noop.NewTracerProvider().Tracer("")
	tracerMu sync.RWMutex
)

// Start creates a span and a context.Context containing the newly-created span from the global tracer.
// See go.opentelemetry.io/otel/trace#Start for more details.
//
//nolint:spancheck // False positive.
func Start(ctx context.Context, spanName string, opts ...trace.SpanStartOption) (context.Context, trace.Span) {
	tracerMu.RLock()
	defer tracerMu.RUnlock()

	return tracer.Start(ctx, spanName, opts...)
}

// RootedCtx returns a copy of the parent context containing a tracing span context
// rooted to the trace ID. All spans started from the context will be rooted to the trace ID.
func RootedCtx(ctx context.Context, traceID trace.TraceID) context.Context {
	return trace.ContextWithSpanContext(ctx, trace.NewSpanContext(trace.SpanContextConfig{
		TraceID: traceID,
	}))
}

// Identifiers defines the global tracer attributes
// That uniquely identifies the network/service/instance
// that produced each trace.
type Identifiers struct {
	Network  netconf.ID
	Service  string // halo/relayer/monitor
	Instance string // validator01/seed01
}

// Init initializes the global tracer via the option(s) defaulting to a noop tracer. It returns a shutdown function.
func Init(ctx context.Context, ids Identifiers, cfg Config, opts ...func(*options)) (func(context.Context) error, error) {
	tracerMu.Lock()
	defer tracerMu.Unlock()

	cfgOpt, err := cfg.toOpts()
	if err != nil {
		return nil, err
	}

	var o options
	for _, opt := range append(opts, cfgOpt) {
		opt(&o)
	}

	if o.exporterFunc == nil {
		return func(context.Context) error {
			return nil
		}, nil
	}

	exporter, err := o.exporterFunc(ctx)
	if err != nil {
		return nil, err
	}

	tp, err := newTraceProvider(exporter, ids)
	if err != nil {
		return nil, err
	}

	// Set globals
	otel.SetTracerProvider(tp)
	tracer = tp.Tracer("")

	return tp.Shutdown, nil
}

type options struct {
	exporterFunc func(context.Context) (sdktrace.SpanExporter, error)
}

// WithStdOut returns an option to configure an OpenTelemetry exporter for tracing
// telemetry to be written to an output destination as JSON.
func WithStdOut(w io.Writer) func(*options) {
	return func(o *options) {
		o.exporterFunc = func(context.Context) (sdktrace.SpanExporter, error) {
			ex, err := stdouttrace.New(stdouttrace.WithWriter(w))
			if err != nil {
				return nil, errors.Wrap(err, "stdout trace exporter")
			}

			return ex, nil
		}
	}
}

// WithOTLP returns an option to configure an OpenTelemetry tracing exporter for Jaeger.
func WithOTLP(endpoint string, headers map[string]string) func(*options) {
	return func(o *options) {
		o.exporterFunc = func(ctx context.Context) (sdktrace.SpanExporter, error) {
			opts := []otlptracehttp.Option{
				otlptracehttp.WithEndpointURL(endpoint),
				otlptracehttp.WithHeaders(headers),
			}

			ex, err := otlptracehttp.New(ctx, opts...)
			if err != nil {
				return nil, errors.Wrap(err, "otlp exporter")
			}

			return ex, nil
		}
	}
}

func newTraceProvider(exp sdktrace.SpanExporter, ids Identifiers) (*sdktrace.TracerProvider, error) {
	r, err := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(semconv.SchemaURL,
			semconv.ServiceName(ids.Service),
			semconv.ServiceInstanceID(ids.Instance),
			semconv.DeploymentEnvironment(ids.Network.String()),
		))
	if err != nil {
		return nil, errors.Wrap(err, "merge resource")
	}

	return sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithBatcher(exp),
		sdktrace.WithResource(r),
	), nil
}

// AddEvent adds an event to the span in the given context with the specified name and attributes.
// See go.opentelemetry.io/otel/trace#Span.AddEvent for more details.
func AddEvent(ctx context.Context, name string, attrs ...attribute.KeyValue) {
	trace.SpanFromContext(ctx).AddEvent(name, trace.WithAttributes(attrs...))
}
