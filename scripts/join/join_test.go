package join_test

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"math"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"time"

	clicmd "github.com/omni-network/omni/cli/cmd"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/tutil"
	"github.com/omni-network/omni/lib/umath"

	rpchttp "github.com/cometbft/cometbft/rpc/client/http"

	"github.com/stretchr/testify/require"
	"golang.org/x/sync/errgroup"
)

var (
	logsFile    = flag.String("logs_file", "join_test.log", "File to write docker logs to")
	network     = flag.String("network", "omega", "Network to join (default: omega)")
	integration = flag.Bool("integration", false, "Run integration tests")
)

// TestJoinNetwork starts a local node (using omni operator init-nodes)
// and waits for it to sync.
//
//nolint:paralleltest // Parallel tests not supported since we start docker containers.
func TestJoinNetwork(t *testing.T) {
	if !*integration {
		t.Skip("skipping integration test")
	}

	const (
		timeout     = time.Hour * 10
		minDuration = time.Minute * 10
	)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	home := t.TempDir()
	logsPath, err := filepath.Abs(*logsFile)
	require.NoError(t, err)
	output, err := os.OpenFile(logsPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	require.NoError(t, err)

	networkID := netconf.ID(*network)
	cfg := clicmd.InitConfig{
		Network: networkID,
		Home:    home,
		Moniker: t.Name(),
		HaloTag: getGitCommit7(t),
	}

	require.NoError(t, ensureHaloImage(cfg.HaloTag))

	log.Info(ctx, "Exec: omni operator init-nodes", "network", networkID)
	require.NoError(t, clicmd.InitNodes(log.WithNoopLogger(ctx), cfg))

	t0 := time.Now()

	var eg errgroup.Group
	eg.Go(func() error {
		defer cancel() // Stop other goroutine

		// Start the nodes.
		log.Info(ctx, "Exec: docker compose up", "logs_file", logsPath)
		cmd := exec.CommandContext(ctx, "docker", "compose", "up")
		cmd.Stderr = output
		cmd.Stdout = output
		cmd.Dir = home
		err := cmd.Run()
		if err == nil || ctx.Err() != nil {
			return nil // Docker compose didn't error
		}

		return errors.Wrap(err, "docker compose up early exit")
	})
	eg.Go(func() error {
		defer cancel() // Stop other goroutine

		// Monitor the progress until synced.
		cmtCl, err := rpchttp.New("http://localhost:26657", "/websocket")
		require.NoError(t, err)
		ethCl, err := ethclient.Dial("omni_evm", "http://localhost:8545")
		require.NoError(t, err)

		ticker := time.NewTicker(time.Second * 30)
		defer ticker.Stop()

		timeoutCtx, timeoutCancel := context.WithTimeout(ctx, timeout)
		defer timeoutCancel()

		var t1 time.Time // When target execution height identified

		for {
			select {
			case <-ctx.Done():
				return nil
			case <-timeoutCtx.Done():
				return errors.New("timed out waiting for sync", "duration", "duration", time.Since(t0).Truncate(time.Second))
			case <-ticker.C:
				haloStatus, err := retry(ctx, haloStatus)
				require.NoError(t, err)
				stats, err := retry(ctx, getContainerStats)
				require.NoError(t, err)

				// CometBFT RPC errors while syncing, so best effort fetch
				var haloSynced bool
				haloHeight := int64(-1) // Indicates failed fetch
				if haloResult, err := cmtCl.Status(ctx); err == nil {
					haloSynced = !haloResult.SyncInfo.CatchingUp
					haloHeight = haloResult.SyncInfo.LatestBlockHeight
				}

				execProgress, execSyncing, err := retryOk(ctx, ethCl.ProgressIfSyncing)
				require.NoError(t, err)

				execSynced := true
				var execHeight, execTarget uint64
				if execSyncing {
					execSynced = execProgress.Done()
					execHeight = execProgress.CurrentBlock
					execTarget = execProgress.HighestBlock
				}

				if t1.IsZero() && execTarget > 0 {
					t1 = time.Now()
				}

				eta, bps := estimate(execHeight, execTarget, time.Since(t1))

				log.Info(ctx, "Status",
					"cstatus", haloStatus,
					"csynced", haloSynced,
					"cheight", haloHeight,
					"ccpu", stats.HaloCPU,
					"crx", stats.HaloRX,
					"esynced", execSynced,
					"eheight", execHeight,
					"etarget", execTarget,
					"ecpu", stats.EVMCPU,
					"erx", stats.EVMRX,
					"duration", time.Since(t0).Truncate(10*time.Second),
					"eta", eta,
					"bps", bps,
				)

				if haloStatus == "healthy" &&
					execSynced &&
					haloSynced &&
					time.Since(t0) > minDuration {
					log.Info(ctx, "Synced 🎉", "duration", time.Since(t0).Truncate(time.Second))

					return nil
				}
			}
		}
	})

	if err := eg.Wait(); err != nil {
		printLogsTail(t, logsPath)
		tutil.RequireNoError(t, err)
	}
}

// estimate returns the estimated time until the target height is reached, and blocks per second given
// current height and time since start.
func estimate(height uint64, target uint64, since time.Duration) (time.Duration, float64) {
	remaining, ok := umath.Subtract(target, height)
	if !ok || height == 0 {
		return 0, 0
	}

	eta := time.Duration(float64(since) * float64(remaining) / float64(height))

	bps := float64(height) / since.Seconds()

	return eta.Truncate(time.Second), math.Round(bps)
}

func printLogsTail(t *testing.T, path string) {
	t.Helper()
	bz, err := os.ReadFile(path)
	require.NoError(t, err)

	lines := strings.Split(string(bz), "\n")
	const n = 50
	if len(lines) > n {
		lines = lines[len(lines)-n:]
	}

	fmt.Println(strings.Join(lines, "\n"))
}

func retryOk[R any](ctx context.Context, fn func(context.Context) (R, bool, error)) (R, bool, error) {
	type commaOK struct {
		R  R
		Ok bool
	}

	resp, err := retry[commaOK](ctx, func(ctx context.Context) (commaOK, error) {
		r, ok, err := fn(ctx)
		return commaOK{R: r, Ok: ok}, err
	})
	if err != nil {
		var zero R
		return zero, false, err
	}

	return resp.R, resp.Ok, nil
}

func retry[R any](ctx context.Context, fn func(context.Context) (R, error)) (R, error) {
	const retry = 10
	for i := 1; ; i++ {
		r, err := fn(ctx)
		if err == nil {
			return r, nil
		}

		if i >= retry {
			return r, err
		}

		log.Warn(ctx, "Failed attempt (will retry)", err, "attempt", i)
		time.Sleep(time.Second * time.Duration(i))
	}
}

func haloStatus(ctx context.Context) (string, error) {
	out, err := exec.CommandContext(ctx, "docker", "inspect", "--format={{json .State.Health.Status }}", "halo").CombinedOutput()
	if err != nil {
		return "", errors.Wrap(err, "docker inspect")
	}

	return strings.Trim(strings.TrimSpace(string(out)), "\""), nil
}

func ensureHaloImage(haloTag string) error {
	out, err := exec.Command("docker", "images", "-q", "omniops/halovisor:"+haloTag).CombinedOutput()
	if err != nil {
		return errors.Wrap(err, "docker images")
	} else if strings.TrimSpace(string(out)) == "" {
		return errors.New("omniops/halovisor:" + haloTag + " image not found locally (make build-docker?)")
	}

	return nil
}

func getGitCommit7(t *testing.T) string {
	t.Helper()

	out, err := exec.Command("git", "rev-parse", "HEAD").CombinedOutput()
	require.NoError(t, err)

	return string(out)[0:7]
}

type stats struct {
	HaloCPU string
	HaloRX  string
	EVMCPU  string
	EVMRX   string
}

type statsJSON struct {
	BlockIO   string `json:"BlockIO"`
	CPUPerc   string `json:"CPUPerc"`
	Container string `json:"Container"`
	ID        string `json:"ID"`
	MemPerc   string `json:"MemPerc"`
	MemUsage  string `json:"MemUsage"`
	Name      string `json:"Name"`
	NetIO     string `json:"NetIO"`
	PIDs      string `json:"PIDs"`
}

// getContainerStats returns halo and omni_evm CPU and network RX stats.
func getContainerStats(ctx context.Context) (stats, error) {
	out, err := exec.CommandContext(ctx, "docker", "stats", "--format=json", "--no-stream", "--no-trunc").CombinedOutput()
	if err != nil {
		return stats{}, errors.Wrap(err, "docker stats")
	}

	var resp stats
	for _, line := range bytes.Split(out, []byte("\n")) {
		line = bytes.TrimSpace(line)
		if string(line) == "" {
			continue
		}

		var s statsJSON
		err := json.Unmarshal(line, &s)
		if err != nil {
			return stats{}, errors.Wrap(err, "json unmarshal")
		}

		cpu := strings.TrimSpace(s.CPUPerc)
		rx := strings.TrimSpace(strings.Split(s.NetIO, "/")[0])
		if s.Name == "halo" {
			resp.HaloCPU = cpu
			resp.HaloRX = rx
		} else if s.Name == "omni_evm" {
			resp.EVMCPU = cpu
			resp.EVMRX = rx
		}
	}

	return resp, nil
}
