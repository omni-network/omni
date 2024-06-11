package anvil

import (
	"bytes"
	"context"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"text/template"
	"time"

	"github.com/omni-network/omni/e2e/app/static"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/log"

	_ "embed"
)

// Start starts a genesis anvil node and returns an ethclient and a stop function or an error.
// The dir parameter is the location of the docker compose.
// If useLogProxy is true, all requests are routed via a reserve proxy that logs all requests, which will be printed
// at stop.
func Start(ctx context.Context, dir string, chainID uint64) (ethclient.Client, string, func(), error) {
	ctx, cancel := context.WithTimeout(ctx, time.Minute) // Allow 1 minute for edge case of pulling images.
	defer cancel()
	if !composeDown(ctx, dir) {
		return nil, "", nil, errors.New("failure to clean up previous anvil instance")
	}

	// Ensure ports are available
	port, err := getAvailablePort()
	if err != nil {
		return nil, "", nil, errors.Wrap(err, "get available port")
	}

	if err := writeComposeFile(dir, chainID, port); err != nil {
		return nil, "", nil, errors.Wrap(err, "write compose file")
	}

	if err := writeAnvilState(dir); err != nil {
		return nil, "", nil, errors.Wrap(err, "write anvil state")
	}

	log.Info(ctx, "Starting anvil")

	out, err := execCmd(ctx, dir, "docker", "compose", "up", "-d", "--remove-orphans")
	if err != nil {
		return nil, "", nil, errors.Wrap(err, "docker compose up: "+out)
	}

	endpoint := "http://localhost:" + port

	ethCl, err := ethclient.Dial("anvil", endpoint)
	if err != nil {
		return nil, "", nil, errors.Wrap(err, "new eth client")
	}

	stop := func() { //nolint:contextcheck // Fresh context required for stopping.
		// Fresh stop context since above context might be canceled.
		stopCtx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()
		composeDown(stopCtx, dir)
	}

	// Wait up to 10 secs for RPC to be available
	const retry = 10
	for i := 0; i < retry; i++ {
		if i == retry-1 {
			stop()
			return nil, "", nil, errors.New("wait for RPC timed out")
		}

		select {
		case <-ctx.Done():
			return nil, "", nil, errors.Wrap(ctx.Err(), "timeout")
		case <-time.After(time.Millisecond * 500):
		}

		_, err := ethCl.BlockNumber(ctx)
		if err == nil {
			break
		}

		log.Warn(ctx, "Anvil: waiting for RPC to be available", err)
	}

	log.Info(ctx, "Anvil: RPC is available", "addr", endpoint)

	return ethCl, endpoint, stop, nil
}

// composeDown runs docker-compose down in the provided directory.
func composeDown(ctx context.Context, dir string) bool {
	if _, err := os.Stat(dir + "/compose.yaml"); os.IsNotExist(err) {
		return true
	}

	out, err := execCmd(ctx, dir, "docker", "compose", "down")
	if err != nil {
		log.Error(ctx, "Error: docker compose down", err, "out", out)
		return false
	}

	log.Debug(ctx, "Anvil: docker compose down: ok")

	return true
}

func execCmd(ctx context.Context, dir string, cmd string, args ...string) (string, error) {
	c := exec.CommandContext(ctx, cmd, args...)
	c.Dir = dir

	out, err := c.CombinedOutput()
	if err != nil {
		return string(out), errors.Wrap(err, "exec", "out", string(out))
	}

	return string(out), nil
}

func writeAnvilState(dir string) error {
	anvilStateFile := filepath.Join(dir, "state.json")
	if err := os.WriteFile(anvilStateFile, static.GetDevnetElAnvilState(), 0o644); err != nil {
		return errors.Wrap(err, "write anvil state")
	}

	return nil
}

//go:embed compose.yaml.tmpl
var composeTpl []byte

func writeComposeFile(dir string, chainID uint64, port string) error {
	tpl, err := template.New("").Parse(string(composeTpl))
	if err != nil {
		return errors.Wrap(err, "parse compose template")
	}

	var buf bytes.Buffer
	err = tpl.Execute(&buf, struct {
		ChainID uint64
		Port    string
	}{
		ChainID: chainID,
		Port:    port,
	})
	if err != nil {
		return errors.Wrap(err, "execute compose template")
	}

	err = os.WriteFile(dir+"/compose.yaml", buf.Bytes(), 0644)
	if err != nil {
		return errors.Wrap(err, "write compose file")
	}

	return nil
}

func getAvailablePort() (string, error) {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		return "", errors.Wrap(err, "resolve addr")
	}

	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return "", errors.Wrap(err, "listen")
	}
	defer l.Close()

	_, port, err := net.SplitHostPort(l.Addr().String())
	if err != nil {
		return "", errors.Wrap(err, "split host port")
	}

	return port, nil
}
