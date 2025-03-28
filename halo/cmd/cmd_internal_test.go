package cmd

import (
	"bytes"
	"context"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/omni-network/omni/halo/app"
	halocfg "github.com/omni-network/omni/halo/config"
	libcmd "github.com/omni-network/omni/lib/cmd"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/tutil"

	fuzz "github.com/google/gofuzz"
	"github.com/stretchr/testify/require"
)

//go:generate go test . -golden -clean

func TestRunCmd(t *testing.T) {
	setMonikerForT(t)

	tests := []struct {
		Name  string
		Args  []string
		Files map[string][]byte
		Env   map[string]string
	}{
		{
			Name: "defaults",
			Args: slice("run"),
		},
		{
			Name: "flags",
			Args: slice("run", "--home=foo", "--engine-jwt-file=bar"),
		},
		{
			Name: "toml files",
			Args: slice("run", "--home=testinput/input1"),
		},
		{
			Name: "json files",
			Args: slice("run", "--home=testinput/input2"),
			Env:  map[string]string{"HALO_NETWORK": "GOTEST"},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			cmd := newRunCmd("run", func(_ context.Context, actual app.Config) error {
				tutil.RequireGoldenJSON(t, actual)

				return nil
			})

			for k, v := range test.Env {
				t.Setenv(k, v)
			}

			rootCmd := libcmd.NewRootCmd("halo", "", cmd)
			rootCmd.SetArgs(test.Args)
			require.NoError(t, rootCmd.Execute())
		})
	}
}

func TestCLIReference(t *testing.T) {
	t.Parallel()
	const root = "halo" // Use to identify root command (vs subcommands).

	tests := []struct {
		Command string
	}{
		{root},
		{"run"},
		{"init"},
		{"rollback"},
	}

	for _, test := range tests {
		t.Run(test.Command, func(t *testing.T) {
			t.Parallel()

			var args []string
			if test.Command != root {
				if strings.Contains(test.Command, "operator") {
					subCmd := strings.Split(test.Command, " ")
					args = append(args, subCmd...)
				} else {
					args = append(args, test.Command)
				}
			}
			args = append(args, "--help")

			cmd := New()
			cmd.SetArgs(args)

			var bz bytes.Buffer
			cmd.SetOut(&bz)

			require.NoError(t, cmd.Execute())

			tutil.RequireGoldenBytes(t, bz.Bytes())
		})
	}
}

//go:generate go test . -run=TestTomlConfig -count=100

func TestTomlConfig(t *testing.T) {
	t.Parallel()
	dir := t.TempDir()
	var chars = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	randomString := func() string {
		var resp string
		for i := 0; i < 1+rand.Intn(64); i++ {
			resp += string(chars[rand.Intn(len(chars))])
		}

		return resp
	}

	// Create a fuzzer with small uint64s and ansi strings (toml struggles with large numbers and UTF8).
	fuzzer := fuzz.New().NilChance(0).NumElements(1, 8).Funcs(
		func(i *uint64, c fuzz.Continue) {
			*i = uint64(rand.Intn(1_000_000))
		},
		func(i *uint, c fuzz.Continue) {
			*i = uint(rand.Intn(1_000_000))
		},
		func(s *string, c fuzz.Continue) {
			*s = randomString()
		},
		func(s *netconf.ID, c fuzz.Continue) {
			*s = netconf.ID(randomString())
		},
	)

	var expect halocfg.Config
	fuzzer.Fuzz(&expect)
	expect.HomeDir = dir

	// The Toml library converts map keys to lower case. So do this so expect==actual.
	for k := range expect.RPCEndpoints {
		expect.RPCEndpoints[strings.ToLower(randomString())] = randomString()
		delete(expect.RPCEndpoints, k)
	}

	// Ensure the <home>/config directory exists.
	require.NoError(t, os.Mkdir(filepath.Join(dir, "config"), 0o755))

	// Write the randomized config to disk.
	require.NoError(t, halocfg.WriteConfigTOML(expect, log.DefaultConfig()))

	// Create a run command that asserts the config is as expected.
	cmd := newRunCmd("run", func(_ context.Context, actual app.Config) error {
		require.Equal(t, expect, actual.Config)

		return nil
	})

	// Create and execute a root command that runs the run command.
	rootCmd := libcmd.NewRootCmd("halo", "", cmd)
	rootCmd.SetArgs([]string{"run", "--home=" + dir})
	tutil.RequireNoError(t, rootCmd.Execute())
}

func TestInvalidCmds(t *testing.T) {
	t.Parallel()
	tests := []struct {
		Name        string
		Args        []string
		ErrContains string
	}{
		{
			Name:        "no args",
			Args:        []string{},
			ErrContains: "no sub-command specified",
		},
		{
			Name:        "invalid args",
			Args:        []string{"invalid"},
			ErrContains: "unknown command \"invalid\" for \"halo\"",
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			rootCmd := libcmd.NewRootCmd("halo", "", New())
			libcmd.WrapRunE(rootCmd, func(ctx context.Context, err error) {})
			rootCmd.SetArgs(test.Args)
			err := rootCmd.Execute()
			require.Error(t, err)
			require.Contains(t, err.Error(), test.ErrContains)
		})
	}
}

//nolint:paralleltest // Global test moniker is used.
func TestDefaultCometConfig(t *testing.T) {
	setMonikerForT(t)

	home := t.TempDir()
	path := filepath.Join(home, "config", "config.toml")
	require.NoError(t, os.MkdirAll(filepath.Dir(path), 0o755))

	cfg := DefaultCometConfig(home)

	err := WriteCometConfig(path, &cfg)
	require.NoError(t, err)

	cfg2, err := parseCometConfig(t.Context(), home)
	require.NoError(t, err)

	cfg.StateSync.RPCServers = []string{} // Replace nil with empty slice for comparison.
	require.Equal(t, cfg, cfg2)

	bz, err := os.ReadFile(path)
	require.NoError(t, err)
	tutil.RequireGoldenBytes(t, bz, tutil.WithFilename("default_config.toml"))
}

// slice is a convenience function for creating string slice literals.
func slice(strs ...string) []string {
	return strs
}
