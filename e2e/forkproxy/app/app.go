package app

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os/exec"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"

	"golang.org/x/sync/errgroup"
)

type Config struct {
	ListenAddr    string
	ChainID       uint64
	LoadState     string
	BlockTimeSecs uint64
	Silent        bool
}

func DefaultConfig() Config {
	return Config{
		ListenAddr:    "0.0.0.0:8545",
		ChainID:       1337,
		LoadState:     "",
		Silent:        true,
		BlockTimeSecs: 1,
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
	// TODO(corver): Start a goroutines that forks the instance every few seconds.
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

		proxy.stopInstance()

		return nil
	})

	if err := eg.Wait(); errors.Is(err, http.ErrServerClosed) {
		return nil // No error on shutdown.
	} else if err != nil {
		return errors.Wrap(err, "server")
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
	instance anvilInstance
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

func (p *proxy) stopInstance() {
	p.mu.Lock()
	defer p.mu.Unlock()

	p.instance.stop()
}

func (p *proxy) setTarget(target anvilInstance) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	u, err := url.Parse(target.URL())
	if err != nil {
		return errors.Wrap(err, "parse target url")
	}

	p.instance = target
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

//nolint:govet // Context is canceled when instance stopped.
func startAnvil(ctx context.Context, cfg anvilConfig) (anvilInstance, error) {
	ctx, cancel := context.WithCancel(ctx)

	args := []string{
		"--port", strconv.Itoa(cfg.Port),
		"--chain-id", strconv.FormatUint(cfg.ChainID, 10),
		"--block-time", strconv.FormatUint(cfg.BlockTimeSecs, 10),
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

	log.Info(ctx, "Starting anvil", "command", strings.Join(args, " "))

	var out bytes.Buffer
	cmd := exec.CommandContext(ctx, "anvil", args...)
	cmd.Stdout = &out
	cmd.Stderr = &out
	if err := cmd.Start(); err != nil {
		return anvilInstance{}, errors.Wrap(err, "start anvil", "out", out.String())
	}

	go func() {
		err := logLines(ctx, &out)
		if err != nil {
			log.Error(ctx, "Failed logging lines", err)
		}
	}()

	return anvilInstance{
		stop: func() {
			cancel()
			_ = cmd.Wait()
		},
		Cfg: cfg,
		Out: out,
		Cmd: cmd,
	}, nil
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
func logLines(ctx context.Context, r io.Reader) error {
	scanner := bufio.NewReader(r)
	for ctx.Err() == nil {
		line, err := scanner.ReadString('\n')
		if errors.Is(err, io.EOF) {
			// Just backoff
			time.Sleep(time.Millisecond * 10)
		} else if err != nil {
			return errors.Wrap(err, "read line")
		} else {
			fmt.Println(line) //nolint:forbidigo // Logging this doesn't make sense.
		}
	}

	return nil
}
