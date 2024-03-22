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
	"github.com/omni-network/omni/lib/tutil"

	fuzz "github.com/google/gofuzz"
	"github.com/stretchr/testify/require"
)

//go:generate go test . -golden -clean

func TestRunCmd(t *testing.T) { //nolint:paralleltest,tparallel // RunCmd modifies global state via setMonikerForT
	setMonikerForT(t)

	tests := []struct {
		Name  string
		Args  []string
		Files map[string][]byte
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
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()

			cmd := newRunCmd("run", func(_ context.Context, actual app.Config) error {
				tutil.RequireGoldenJSON(t, actual)

				return nil
			})

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

func TestTomlConfig(t *testing.T) {
	t.Parallel()
	dir := t.TempDir()
	var chars = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	// Create a fuzzer with small uint64s and ansi strings (toml struggles with large numbers and UTF8).
	fuzzer := fuzz.New().Funcs(
		func(i *uint64, c fuzz.Continue) {
			*i = uint64(rand.Intn(1_000_000))
		},
		func(s *string, c fuzz.Continue) {
			for i := 0; i < rand.Intn(64); i++ {
				*s += string(chars[rand.Intn(len(chars))])
			}
		},
	)

	var expect halocfg.Config
	fuzzer.Fuzz(&expect)
	expect.HomeDir = dir

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
	require.NoError(t, rootCmd.Execute())
}

// slice is a convenience function for creating string slice literals.
func slice(strs ...string) []string {
	return strs
}
