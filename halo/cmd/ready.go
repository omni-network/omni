package cmd

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"

	cmtcfg "github.com/cometbft/cometbft/config"

	"github.com/spf13/cobra"
)

func newReadyCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ready",
		Short: "Assert the readiness of the halo node",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			err := assertReady(cmd.Context())
			if err != nil {
				return errors.Wrap(err, "ready failed")
			}

			return nil
		},
	}

	return cmd
}

// assertReady calls halo's /ready endpoint and returns nil if the status is ready
// or an error otherwise.
func assertReady(ctx context.Context) error {
	cfg := cmtcfg.DefaultConfig()
	url := fmt.Sprintf("http://0.0.0.0%v/ready", cfg.Instrumentation.PrometheusListenAddr)

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return errors.Wrap(err, "http request creation")
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return errors.Wrap(err, "http request")
	}
	defer resp.Body.Close()

	if resp.StatusCode < 400 {
		log.Info(ctx, "The node is ready")
		return nil
	}

	return errors.New("the node is not ready yet")
}
