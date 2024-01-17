package cmd

import (
	"context"
	"path/filepath"
	"testing"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"

	cfg "github.com/cometbft/cometbft/config"

	"github.com/spf13/viper"
)

//nolint:gochecknoglobals // Overrides cometbft default moniker for testing.
var testMoniker string

// setMonikerForT sets the test moniker for the duration of the test.
// This is required for deterministic default cometbft config.
func setMonikerForT(t *testing.T) {
	t.Helper()
	testMoniker = "testmoniker"
	t.Cleanup(func() {
		testMoniker = ""
	})
}

// parseCometConfig parses the cometBFT config from disk and verifies it.
func parseCometConfig(ctx context.Context, homeDir string) (cfg.Config, error) {
	conf := cfg.DefaultConfig()

	if testMoniker != "" {
		conf.Moniker = testMoniker
	}

	const (
		file = "config" // CometBFT config files are named config.toml
		dir  = "config" // CometBFT config files are stored in the config directory
	)

	v := viper.New()
	v.SetConfigName(file)
	v.AddConfigPath(filepath.Join(homeDir, dir))

	// Attempt to read the cometBFT config file, gracefully ignoring errors
	// caused by a config file not being found. Return an error
	// if we cannot parse the config file.
	if err := v.ReadInConfig(); err != nil {
		// It's okay if there isn't a config file
		var cfgError viper.ConfigFileNotFoundError
		if ok := errors.As(err, &cfgError); !ok {
			return cfg.Config{}, errors.Wrap(err, "read comet config")
		}

		log.Warn(ctx, "No comet config.toml file found, using default config", nil)
	}

	err := v.Unmarshal(conf)
	if err != nil {
		return cfg.Config{}, errors.Wrap(err, "unmarshal comet config")
	}

	conf.RootDir = homeDir
	conf.SetRoot(conf.RootDir)
	if err := conf.ValidateBasic(); err != nil {
		return cfg.Config{}, errors.Wrap(err, "validate comet config")
	}

	if warnings := conf.CheckDeprecated(); len(warnings) > 0 {
		for _, warning := range warnings {
			log.Info(ctx, "Deprecated CometBFT config", "usage", warning)
		}
	}

	return *conf, nil
}
