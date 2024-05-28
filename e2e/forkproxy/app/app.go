package app

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"

	"github.com/ethereum/go-ethereum/ethclient"

	"golang.org/x/sync/errgroup"
)

type Config struct {
	ListenAddr    string
	ChainID       uint64
	LoadState     string
	BlockTimeSecs uint64
	Silent        bool
	EnableForking bool
	SlotsInEpoch  uint64
}

func (c Config) FinalizeDepth() uint64 {
	return c.SlotsInEpoch * 2
}

func (c Config) MaxForkDepth() uint64 {
	return c.FinalizeDepth() / 2
}

func DefaultConfig() Config {
	return Config{
		ListenAddr:    "0.0.0.0:8545",
		ChainID:       1337,
		LoadState:     "",
		Silent:        true,
		BlockTimeSecs: 1,
		SlotsInEpoch:  32,
	}
}

func Run(ctx context.Context, cfg Config) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	portProvider := newPortProvider()
	rootCfg := anvilConfig{
		Config: cfg,
		Port:   portProvider(),
	}
	root, err := startAnvil(ctx, rootCfg)
	if err != nil {
		return errors.Wrap(err, "start root anvil")
	}

	proxy, err := newProxy(root)
	if err != nil {
		return errors.Wrap(err, "new proxy")
	}

	httpServer := &http.Server{
		Addr:              cfg.ListenAddr,
		ReadHeaderTimeout: 30 * time.Second,
		IdleTimeout:       30 * time.Second,
		WriteTimeout:      30 * time.Second,
		Handler:           proxy,
	}

	log.Info(ctx, "Starting forkproxy server", "address", cfg.ListenAddr)

	var eg errgroup.Group
	eg.Go(func() error {
		return forkForever(ctx, cfg, portProvider, proxy)
	})
	eg.Go(func() error {
		defer cancel() // Cancel the app context if serving fails.

		// ListenAndServe always returns an error.
		return errors.Wrap(httpServer.ListenAndServe(), "serve")
	})
	eg.Go(func() error {
		<-ctx.Done()
		log.Info(ctx, "Shutdown detected, stopping server")

		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := httpServer.Shutdown(shutdownCtx); err != nil { //nolint:contextcheck // Explicit new shutdown context.
			return errors.Wrap(err, "server shutdown")
		}

		proxy.getInstance().stop()

		return nil
	})

	if err := eg.Wait(); errors.Is(err, http.ErrServerClosed) {
		return nil // No error on shutdown.
	} else if err != nil {
		return errors.Wrap(err, "run server")
	}

	return nil
}

func forkForever(ctx context.Context, cfg Config, newPort func() int, proxy *proxy) error {
	if !cfg.EnableForking {
		<-ctx.Done()
		return nil
	}
	if cfg.BlockTimeSecs != 1 {
		return errors.New("forking only supported for 1s block times")
	}

	for ctx.Err() == nil {
		nextForkDepth := 1 + rand.Intn(int(cfg.MaxForkDepth()-1))       //nolint:gosec // Not a problem
		sleepSecs := (nextForkDepth * int(cfg.BlockTimeSecs)) * 12 / 10 // Wait 20% longer

		select {
		case <-ctx.Done():
			return nil
		case <-time.After(time.Duration(sleepSecs) * time.Second):
		}

		current := proxy.getInstance()
		height, err := current.Height(ctx)
		if err != nil {
			return errors.Wrap(err, "get height")
		} else if int(height) <= nextForkDepth {
			log.Warn(ctx, "Fork depth too deep", nil, "height", height, "fork_depth", nextForkDepth)
			continue
		}

		forkHeight := height - uint64(nextForkDepth)

		log.Info(ctx, "Forking chain", "fork_depth", nextForkDepth, "current_height", height, "forked_height", forkHeight)

		// Configure and start fork
		nextCfg := anvilConfig{
			Config:       cfg,
			Port:         newPort(),
			ForkInstance: current,
			ForkHeight:   forkHeight,
		}
		next, err := startAnvil(ctx, nextCfg)
		if err != nil {
			return errors.Wrap(err, "start anvil")
		}
		// Wait for it come up
		if err := next.AwaitHeight(ctx, forkHeight+1, time.Second*5); err != nil {
			return errors.Wrap(err, "await up")
		}
		// Replace proxy target
		if err := proxy.setTarget(next); err != nil {
			return errors.Wrap(err, "set target")
		}
	}

	return nil
}

func newProxy(instance anvilInstance) (*proxy, error) {
	resp := new(proxy)
	if err := resp.setTarget(instance); err != nil {
		return nil, err
	}

	return resp, nil
}

type proxy struct {
	mu       sync.RWMutex
	instance *anvilInstance
	target   *url.URL
}

func (p *proxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	httputil.NewSingleHostReverseProxy(p.getTarget()).ServeHTTP(w, r)
}

func (p *proxy) getTarget() *url.URL {
	p.mu.RLock()
	defer p.mu.RUnlock()

	return p.target
}

func (p *proxy) getInstance() *anvilInstance {
	p.mu.Lock()
	defer p.mu.Unlock()

	return p.instance
}

func (p *proxy) setTarget(target anvilInstance) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	u, err := url.Parse(target.URL())
	if err != nil {
		return errors.Wrap(err, "parse target url")
	}

	if p.instance != nil {
		p.instance.stop()
	}

	p.instance = &target
	p.target = u

	return nil
}

type anvilConfig struct {
	Config
	Port         int
	ForkHeight   uint64
	ForkInstance *anvilInstance
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

	ethCl, err := ethclient.DialContext(ctx, i.URL())
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
	if cfg.ForkInstance != nil {
		args = append(args,
			"--fork-url", cfg.ForkInstance.URL(),
			"--fork-block-number", strconv.FormatUint(cfg.ForkHeight, 10),
		)
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
func logLines(ctx context.Context, r io.Reader, prefix string) error {
	scanner := bufio.NewReader(r)
	for ctx.Err() == nil {
		line, err := scanner.ReadString('\n')
		if errors.Is(err, io.EOF) {
			// Just backoff
			time.Sleep(time.Millisecond * 10)
		} else if err != nil {
			return errors.Wrap(err, "read line")
		} else {
			fmt.Println(prefix + ": " + line) //nolint:forbidigo // Logging this doesn't make sense.
		}
	}

	return nil
}
