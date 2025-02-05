package solana

import (
	"bytes"
	"context"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"text/template"
	"time"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"

	solana "github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"

	_ "embed"
)

// Start starts a genesis solana node and returns a client, a funded private key and a stop function or an error.
// The dir parameter is the location of the docker compose.
func Start(ctx context.Context, dir string) (*rpc.Client, solana.PrivateKey, func(), error) {
	ctx, cancel := context.WithTimeout(ctx, time.Minute) // Allow 1 minute for edge case of pulling images.
	defer cancel()

	if !composeDown(ctx, dir) {
		return nil, nil, nil, errors.New("failure to clean up previous solana instance")
	}

	// Ensure ports are available
	port, err := getAvailablePort()
	if err != nil {
		return nil, nil, nil, errors.Wrap(err, "get available port")
	}

	if err := writeComposeFile(dir, port, "stable"); err != nil {
		return nil, nil, nil, errors.Wrap(err, "write compose file")
	}

	log.Info(ctx, "Starting solana")

	out, err := execCmd(ctx, dir, "docker", "compose", "up", "-d", "--remove-orphans")
	if err != nil {
		return nil, nil, nil, errors.Wrap(err, "docker compose up: "+out)
	}

	endpoint := "http://localhost:" + port

	cl := rpc.New(endpoint)

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
			return nil, nil, nil, errors.New("wait for RPC timed out")
		}

		select {
		case <-ctx.Done():
			return nil, nil, nil, errors.Wrap(ctx.Err(), "timeout")
		case <-time.After(time.Millisecond * 500):
		}

		_, err := cl.GetBlockHeight(ctx, rpc.CommitmentFinalized)
		if err == nil {
			break
		}

		if i > retry/2 {
			log.Warn(ctx, "Solana: waiting for RPC to be available", err)
		}
	}

	privKey, err := solana.PrivateKeyFromSolanaKeygenFile(filepath.Join(dir, "id.json"))
	if err != nil {
		return nil, nil, nil, errors.Wrap(err, "get private key")
	}

	log.Info(ctx, "Solana: RPC is available", "addr", endpoint)

	return cl, privKey, stop, nil
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

	log.Debug(ctx, "Solana: docker compose down: ok")

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

//go:embed compose.yaml.tmpl
var composeTpl []byte

func writeComposeFile(dir string, port, version string) error {
	tpl, err := template.New("").Parse(string(composeTpl))
	if err != nil {
		return errors.Wrap(err, "parse compose template")
	}

	var buf bytes.Buffer
	err = tpl.Execute(&buf, struct {
		Port    string
		Version string
	}{
		Port:    port,
		Version: version,
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
