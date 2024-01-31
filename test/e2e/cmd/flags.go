package cmd

import (
	"github.com/omni-network/omni/test/e2e/app"

	"github.com/spf13/pflag"
)

func bindDefFlags(flags *pflag.FlagSet, cfg *app.DefinitionConfig) {
	flags.StringVarP(&cfg.ManifestFile, "manifest-file", "f", cfg.ManifestFile, "path to manifest file")
	flags.StringVar(&cfg.InfraProvider, "infra", cfg.InfraProvider, "infrastructure provider")
	flags.StringVar(&cfg.DeployKeyFile, "deploy-key", cfg.DeployKeyFile, "path to deploy private key file")
	flags.StringVar(&cfg.RelayerKeyFile, "relayer-key", cfg.RelayerKeyFile, "path to relayer private key file")
}

func bindE2EFlags(flags *pflag.FlagSet, cfg *app.E2ETestConfig) {
	flags.BoolVar(&cfg.Preserve, "preserve", cfg.Preserve, "preserve infrastructure after test")
}
