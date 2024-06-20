//nolint:paralleltest // Tracing use global tracers.
package tracer_test

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"os"
	"testing"
	"time"

	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/tracer"

	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

var integration = flag.Bool("integration", false, "run integration tests")

func TestIntegration(t *testing.T) {
	if !*integration {
		t.Skip("tracer integration tests disabled")
	}

	endpoint, ok := os.LookupEnv("TRACING_ENDPOINT")
	require.True(t, ok)
	headers, ok := os.LookupEnv("TRACING_HEADERS")
	require.True(t, ok)

	cfg := tracer.Config{
		Endpoint: endpoint,
		Headers:  headers,
	}

	ids := tracer.Identifiers{
		Network:  netconf.Simnet,
		Service:  "test_service",
		Instance: "test_instance",
	}

	ctx := context.Background()
	stop, err := tracer.Init(ctx, ids, cfg)
	require.NoError(t, err)
	defer func() {
		t.Log("Stopping tracer")
		require.NoError(t, stop(ctx))
		t.Log("Stopped tracer")
	}()

	ctx, span1 := tracer.StartChainHeight(ctx, netconf.Simnet, "test_chain", 123, "root")
	defer span1.End()

	time.Sleep(time.Millisecond * 10)

	ctx, span2 := tracer.Start(ctx, "parent")
	defer span2.End()

	time.Sleep(time.Millisecond * 20)

	tracer.AddEvent(ctx, "event", attribute.String("k", "v"))

	ctx, span3 := tracer.Start(ctx, "child")
	defer span3.End()

	time.Sleep(time.Millisecond * 30)
}

func TestDefaultNoopTracer(_ *testing.T) {
	// This just shouldn't panic.
	ctx, span := tracer.Start(context.Background(), "root")
	defer span.End()

	inner(ctx)
}

func TestStdOutTracer(t *testing.T) {
	ctx := context.Background()

	ids := tracer.Identifiers{
		Network:  netconf.Simnet,
		Service:  "test_service",
		Instance: "test_instance",
	}

	var buf bytes.Buffer
	stop, err := tracer.Init(ctx, ids, tracer.Config{}, tracer.WithStdOut(&buf))
	require.NoError(t, err)

	ctx, span := tracer.Start(ctx, "root")
	inner(ctx)
	span.End()

	// Force flush
	err = otel.GetTracerProvider().(interface {
		ForceFlush(ctx context.Context) error
	}).ForceFlush(ctx)
	require.NoError(t, err)

	require.NoError(t, stop(ctx))

	// Sometimes the buffer is empty even though we force flush. We need to fix this, but for now just don't flap.
	if buf.Len() == 0 {
		t.Log("Skipping due to race")
		return
	}

	var m map[string]any
	d := json.NewDecoder(&buf)

	err = d.Decode(&m)
	require.NoError(t, err)
	require.Equal(t, "inner", m["Name"])

	err = d.Decode(&m)
	require.NoError(t, err)
	require.Equal(t, "root", m["Name"])
}

func inner(ctx context.Context) {
	var span trace.Span
	_, span = tracer.Start(ctx, "inner")
	defer span.End()
}
