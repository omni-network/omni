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

	ctx := context.Background()
	stop, err := tracer.Init(ctx, netconf.Simnet, cfg)
	require.NoError(t, err)
	defer func() {
		t.Log("Stopping tracer")
		require.NoError(t, stop(ctx))
		t.Log("Stopped tracer")
	}()

	ctx, span1 := tracer.StartChainHeight(ctx, "test_chain", 123, "root")
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

	var buf bytes.Buffer
	stop, err := tracer.Init(ctx, netconf.Simnet, tracer.Config{}, tracer.WithStdOut(&buf))
	require.NoError(t, err)

	ctx, span := tracer.Start(ctx, "root")
	inner(ctx)
	span.End()

	require.NoError(t, stop(ctx))

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
