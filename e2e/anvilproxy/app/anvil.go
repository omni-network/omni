package app

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/log"
)

type anvilConfig struct {
	Config
	Port int
}

type anvilInstance struct {
	Cfg  anvilConfig
	stop func()
	Cmd  *exec.Cmd
	Out  bytes.Buffer
}

func (i anvilInstance) URL() string {
	return fmt.Sprintf("http://127.0.0.1:%d", i.Cfg.Port)
}

func (i anvilInstance) Height(ctx context.Context) (uint64, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	ethCl, err := ethclient.Dial("proxy", i.URL())
	if err != nil {
		return 0, errors.Wrap(err, "dial ethclient")
	}

	h, err := ethCl.BlockNumber(ctx)
	if err != nil {
		return 0, errors.Wrap(err, "get block number")
	}

	return h, nil
}

func (i anvilInstance) AwaitHeight(in context.Context, min uint64, timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(in, timeout)
	defer cancel()

	for ctx.Err() == nil {
		if i.Cmd.ProcessState != nil && i.Cmd.ProcessState.Exited() {
			return errors.New("process exited", "state", i.Cmd.ProcessState.String())
		}

		h, err := i.Height(ctx)
		if err == nil && h >= min {
			return nil
		}

		time.Sleep(time.Second / 3)
	}

	if in.Err() != nil {
		return errors.Wrap(ctx.Err(), "parent context canceled awaiting up")
	}

	return errors.Wrap(ctx.Err(), "timeout awaiting up")
}

//nolint:govet // Context is canceled when instance stopped.
func startAnvil(ctx context.Context, cfg anvilConfig) (anvilInstance, error) {
	ctx, cancel := context.WithCancel(ctx)

	args := []string{
		"--port", strconv.Itoa(cfg.Port),
		"--chain-id", strconv.FormatUint(cfg.ChainID, 10),
		"--block-time", strconv.FormatUint(cfg.BlockTimeSecs, 10),
		"--slots-in-an-epoch", strconv.FormatUint(cfg.SlotsInEpoch, 10),
	}
	if cfg.LoadState != "" {
		args = append(args, "--load-state", cfg.LoadState)
	}
	if cfg.Silent {
		args = append(args, "--silent")
	}

	dir, err := os.MkdirTemp("", "")
	if err != nil {
		return anvilInstance{}, errors.Wrap(err, "create temp dir")
	}

	log.Info(ctx, "Starting anvil", "command", strings.Join(args, " "))

	var out bytes.Buffer
	cmd := exec.CommandContext(ctx, "anvil", args...)
	cmd.Stdout = &out
	cmd.Stderr = &out
	cmd.Dir = dir
	if err := cmd.Start(); err != nil {
		return anvilInstance{}, errors.Wrap(err, "start anvil", "out", out.String())
	}

	resp := anvilInstance{
		stop: func() {
			log.Debug(ctx, "Stopping anvil", "port", cfg.Port)
			cancel()
			_ = cmd.Process.Kill()
			_ = cmd.Wait()
		},
		Cfg: cfg,
		Out: out,
		Cmd: cmd,
	}

	if err := resp.AwaitHeight(ctx, 0, time.Second*5); err != nil {
		return resp, errors.Wrap(err, "await anvil startup", "out", out.String())
	}

	go func() {
		err := logLines(ctx, &out, fmt.Sprint(cfg.Port))
		if err != nil {
			log.Error(ctx, "Failed logging lines", err)
		}
	}()

	return resp, nil
}

func newPortProvider() func() int {
	next := 9000 - 1
	return func() int {
		next++
		return next
	}
}

// logLines blocks, reading lines from the reader and logging them.
// It returns when the reader is closed; io.EOF.
func logLines(ctx context.Context, buf *bytes.Buffer, prefix string) error {
	for ctx.Err() == nil {
		line, err := buf.ReadString('\n')
		if errors.Is(err, io.EOF) {
			// Just backoff
			time.Sleep(time.Millisecond * 100)
		} else if err != nil {
			return errors.Wrap(err, "read line")
		} else {
			fmt.Println(prefix + ": " + strings.TrimSuffix(line, "\n")) //nolint:forbidigo // Logging this doesn't make sense.
		}
	}

	return nil
}
