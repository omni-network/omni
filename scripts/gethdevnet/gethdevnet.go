package gethdevnet

import (
	"context"
	"fmt"
	"net"
	"os/exec"
	"path"
	"time"

	"github.com/omni-network/omni/lib/engine"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
)

const (
	gethPort     = 8551
	logProxyPort = 9551
)

// StartGenesisGeth starts a genesis geth node and returns an auth rpc client and a stop function or an error.
// The dir parameter is the location of this source code directory.
// If useLogProxy is true, all requests are routed via a reserve proxy that logs all requests, which will be printed
// at stop.
func StartGenesisGeth(ctx context.Context, dir string, useLogProxy bool) (engine.Client, func(), error) {
	if !composeDown(dir) { //nolint:contextcheck // Composedown uses a fresh context
		return engine.Client{}, nil, errors.New("failure to clean up previous geth instance")
	}

	// Ensure ports are available
	if !isPortFree(gethPort) || !isPortFree(logProxyPort) {
		// Try to stop any running docker instances
		return engine.Client{}, nil, errors.New("ports 8551 and 9551 must be free")
	}

	log.Info(ctx, "Starting geth")

	// Start geth
	out, err := execCmd(ctx, dir, "docker", "compose", "up", "-d", "--remove-orphans")
	if err != nil {
		return engine.Client{}, nil, errors.Wrap(err, "docker compose up: "+out)
	}

	jwtByes, err := engine.LoadJWTHexFile(path.Join(dir, "execution", "jwtsecret"))
	if err != nil {
		return engine.Client{}, nil, errors.Wrap(err, "read jwt file")
	}

	port := gethPort
	if useLogProxy {
		port = logProxyPort
	}
	endpoint := fmt.Sprintf("http://localhost:%d", port)

	engCl, err := engine.NewClient(ctx, endpoint, jwtByes)
	if err != nil {
		return engine.Client{}, nil, errors.Wrap(err, "new engine client")
	}

	stop := func() {
		if useLogProxy {
			// Use a new context to ensure we always print the logs
			out, err = execCmd(context.Background(), dir, "docker", "compose", "logs", "logproxy")
			if err == nil {
				fmt.Printf("====== logproxy: docker compose logs =======\n%s\n", out)
			}
		}
		composeDown(dir)
	}

	// Wait up to 30 secs for geth RPC to be available
	const retry = 30
	for i := 0; i < retry; i++ {
		if i == retry-1 {
			stop()
			return engine.Client{}, nil, errors.New("geth: wait for RPC timed out")
		}

		select {
		case <-ctx.Done():
			return engine.Client{}, nil, ctx.Err()
		case <-time.After(time.Millisecond * 500): //nolint:gomnd // 500ms is common wait time
		}

		_, err := engCl.BlockNumber(ctx)
		if err == nil {
			break
		}

		log.Warn(ctx, "Geth: waiting for RPC to be available", err)
	}

	log.Info(ctx, "Geth: RPC is available", "addr", endpoint)

	return engCl, stop, nil
}

// composeDown runs docker-compose down in the provided directory.
func composeDown(dir string) bool {
	ctx := context.Background() // Use a new context to ensure we always stop

	out, err := execCmd(ctx, dir, "docker", "compose", "down")
	if err != nil {
		log.Error(ctx, "Error: docker compose down", err, "out", out)
		return false
	}

	log.Debug(ctx, "Geth: docker compose down: ok")

	return true
}

func execCmd(ctx context.Context, dir string, cmd string, args ...string) (string, error) {
	c := exec.CommandContext(ctx, cmd, args...)
	c.Dir = dir

	out, err := c.CombinedOutput()
	if err != nil {
		return string(out), fmt.Errorf("exec %s %s: %w", cmd, args, err)
	}

	return string(out), nil
}

func isPortFree(port int) bool {
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return false
	}
	_ = l.Close()

	return true
}
