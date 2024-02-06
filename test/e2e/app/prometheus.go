package app

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"time"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/test/e2e/types"

	_ "embed"
)

type PromSecrets struct {
	URL  string
	User string
	Pass string
}

const promPort = 26660 // Default metrics port for all omni apps (from cometBFT)

//go:embed static/prometheus.yml.tmpl
var promConfigTmpl []byte

func writePrometheusConfig(ctx context.Context, testnet types.Testnet, secrets PromSecrets) error {
	bz, err := genPromConfig(ctx, testnet, secrets)
	if err != nil {
		return errors.Wrap(err, "generating prometheus config")
	}

	promFile := filepath.Join(testnet.Dir, "prometheus", "prometheus.yml")
	if err := os.MkdirAll(filepath.Dir(promFile), 0755); err != nil {
		return errors.Wrap(err, "creating prometheus dir")
	}

	if err := os.WriteFile(promFile, bz, 0644); err != nil {
		return errors.Wrap(err, "writing prometheus config")
	}

	return nil
}

func genPromConfig(ctx context.Context, testnet types.Testnet, secrets PromSecrets) ([]byte, error) {
	var nodeTargets []string
	for _, node := range testnet.Nodes {
		// Prometheus is always inside the same docker-compose, so use service names.
		nodeTargets = append(nodeTargets, fmt.Sprintf("%s:%d", node.Name, promPort))
	}

	network := testnet.Network
	if network == netconf.Devnet {
		network += "-" + time.Now().Format("20060102150405") // Add suffix to distinguish between devnets
	}

	if secrets.URL == "" {
		log.Warn(ctx, "Prometheus remote URL not set, metrics not being pushed to Grafana cloud", nil)
	}

	data := promTmplData{
		Network:        network,
		RemoteURL:      secrets.URL,
		RemoteUsername: secrets.User,
		RemotePassword: secrets.Pass,
		ScrapeConfigs: []promScrapConfig{
			{
				JobName:     "relayer",
				MetricsPath: "/metrics",
				targets:     []string{fmt.Sprintf("relayer:%d", promPort)},
			}, {
				JobName:     "halo",
				MetricsPath: "/metrics",
				targets:     nodeTargets,
			},
		},
	}

	t, err := template.New("").Parse(string(promConfigTmpl))
	if err != nil {
		return nil, errors.Wrap(err, "parsing template")
	}

	var bz bytes.Buffer
	if err := t.Execute(&bz, data); err != nil {
		return nil, errors.Wrap(err, "executing template")
	}

	return bz.Bytes(), nil
}

type promTmplData struct {
	Network        string            // Used a "network" label to all metrics
	RemoteURL      string            // URL to the Grafana cloud server
	RemoteUsername string            // Username to the Grafana cloud server
	RemotePassword string            // Password to the Grafana cloud server
	ScrapeConfigs  []promScrapConfig // List of scrape configs
}

type promScrapConfig struct {
	JobName     string
	MetricsPath string
	targets     []string
}

func (c promScrapConfig) Targets() string {
	return strings.Join(c.targets, ",")
}
