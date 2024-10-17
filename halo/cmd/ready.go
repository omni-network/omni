package cmd

import (
	"context"
	"net/http"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"

	"github.com/spf13/cobra"
)

type readyConfig struct {
	MonitoringURL string
}

func defaultReadyConfig() readyConfig {
	return readyConfig{
		MonitoringURL: "http://localhost:26660",
	}
}

func newReadyCmd() *cobra.Command {
	cfg := defaultReadyConfig()

	cmd := &cobra.Command{
		Use:   "ready",
		Short: "Query remote node for readiness",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			err := queryReady(cmd.Context(), cfg)
			if err != nil {
				return errors.Wrap(err, "ready")
			}

			return nil
		},
	}

	bindReadyFlags(cmd, &cfg)

	return cmd
}

// queryReady calls halo's /ready endpoint and returns nil if the status is ready
// or an error otherwise.
func queryReady(ctx context.Context, cfg readyConfig) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, cfg.MonitoringURL, nil)
	if err != nil {
		return errors.Wrap(err, "http request creation")
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return errors.Wrap(err, "http request")
	}
	defer resp.Body.Close()

	if resp.StatusCode/100 != 2 {
		return errors.New("node not ready")
	}

	log.Info(ctx, "Node ready")

	return nil
}
