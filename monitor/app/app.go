package monitor

import (
	"context"
	"net/http"
	"time"

	"github.com/omni-network/omni/lib/buildinfo"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/monitor/account"
	"github.com/omni-network/omni/monitor/avs"
	"github.com/omni-network/omni/monitor/loadgen"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Run starts the monitor service.
func Run(ctx context.Context, cfg Config) error {
	log.Info(ctx, "Starting monitor service")

	buildinfo.Instrument(ctx)

	network, err := netconf.Load(cfg.NetworkFile)
	if err != nil {
		return err
	}

	if err := avs.Monitor(ctx, network); err != nil {
		return errors.Wrap(err, "monitor AVS")
	}

	if err := account.Monitor(ctx, network); err != nil {
		return errors.Wrap(err, "monitor account balances")
	}

	if err := startLoadGen(ctx, cfg, network); err != nil {
		return errors.Wrap(err, "start load generator")
	}

	if err := startAVSSync(ctx, cfg, network); err != nil {
		return errors.Wrap(err, "start AVS sync")
	}

	select {
	case <-ctx.Done():
		log.Info(ctx, "Shutdown detected, stopping...")
		return nil
	case err := <-serveMonitoring(cfg.MonitoringAddr):
		return err
	}
}

// serveMonitoring starts a goroutine that serves the monitoring API. It
// returns a channel that will receive an error if the server fails to start.
func serveMonitoring(address string) <-chan error {
	errChan := make(chan error)
	go func() {
		mux := http.NewServeMux()
		mux.Handle("/metrics", promhttp.Handler())

		srv := &http.Server{
			Addr:              address,
			ReadHeaderTimeout: 5 * time.Second,
			IdleTimeout:       5 * time.Second,
			WriteTimeout:      5 * time.Second,
			Handler:           mux,
		}
		errChan <- errors.Wrap(srv.ListenAndServe(), "serve monitoring")
	}()

	return errChan
}

func startLoadGen(ctx context.Context, cfg Config, network netconf.Network) error {
	if err := loadgen.Start(ctx, network, cfg.LoadGen); err != nil {
		return errors.Wrap(err, "start load generator")
	}

	return nil
}
