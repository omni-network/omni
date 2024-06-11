package tracer

import (
	"net/url"
	"strings"

	"github.com/omni-network/omni/lib/errors"

	"github.com/spf13/pflag"
)

// BindFlags binds the provided flags to the corresponding fields in the Config struct.
func BindFlags(flags *pflag.FlagSet, cfg *Config) {
	flags.StringVar(&cfg.Endpoint, "tracing-endpoint", cfg.Endpoint, "Tracing OTLP endpoint")
	flags.StringVar(&cfg.Headers, "tracing-headers", cfg.Headers, "Tracing OTLP headers")
}

// DefaultConfig returns the default empty configuration for OTLP tracing.
func DefaultConfig() Config {
	return Config{}
}

// Config defines OTLP config for grafana cloud.
// See https://grafana.com/docs/grafana-cloud/monitor-applications/application-observability/setup/quickstart/go/
type Config struct {
	Endpoint string // E.g. "https://otlp-gateway-prod-us-east-0.grafana.net/otlp"
	Headers  string // E.g. "Authorization=Basic NzQk..3O34"
}

func (c Config) toOpts() (func(*options), error) {
	if c.Endpoint == "" {
		return func(*options) {}, nil
	}

	headers := make(map[string]string)
	if len(c.Headers) > 0 {
		var err error
		headers, err = stringToHeader(c.Headers)
		if err != nil {
			return nil, err
		}
	}

	return WithOTLP(c.Endpoint, headers), nil
}

// stringToHeader converts a string of comma-separated header key-value pairs
// into a map[string]string.
// Copied from go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp@v1.24.0/internal/envconfig/envconfig.go:167.
func stringToHeader(value string) (map[string]string, error) {
	headersPairs := strings.Split(value, ",")
	headers := make(map[string]string)

	for _, header := range headersPairs {
		n, v, found := strings.Cut(header, "=")
		if !found {
			return nil, errors.New("missing '='", "header", header)
		}

		name, err := url.PathUnescape(n)
		if err != nil {
			return nil, errors.Wrap(err, "escape header key", "key", n)
		}

		trimmedName := strings.TrimSpace(name)
		value, err := url.PathUnescape(v)
		if err != nil {
			return nil, errors.Wrap(err, "escape header value", "value", v)
		}

		trimmedValue := strings.TrimSpace(value)

		headers[trimmedName] = trimmedValue
	}

	return headers, nil
}
