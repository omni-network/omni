//nolint:lll // Long lines are easier to read for flag descriptions.
package cmd

import (
	"github.com/omni-network/omni/test/e2e/app"

	"github.com/spf13/pflag"
)

func bindDefFlags(flags *pflag.FlagSet, cfg *app.DefinitionConfig) {
	flags.StringVarP(&cfg.ManifestFile, "manifest-file", "f", cfg.ManifestFile, "path to manifest file")
	flags.StringVar(&cfg.InfraProvider, "infra", cfg.InfraProvider, "infrastructure provider: docker, vmcompose")
	flags.StringVar(&cfg.InfraDataFile, "infra-file", cfg.InfraDataFile, "infrastructure data file (not required for docker provider)")
	flags.StringVar(&cfg.DeployKeyFile, "deploy-key", cfg.DeployKeyFile, "path to deploy private key file")
	flags.StringVar(&cfg.RelayerKeyFile, "relayer-key", cfg.RelayerKeyFile, "path to relayer private key file")
	flags.StringVar(&cfg.OmniImgTag, "omni-image-tag", cfg.OmniImgTag, "Omni docker images tag (halo, relayer). Defaults to working dir git commit.")
	flags.StringToStringVar(&cfg.AnvilStateFiles, "anvil-state", cfg.AnvilStateFiles, "path to anvil state files to load into anvil chains: '<chain1>=<file1>'")
	flags.StringToStringVar(&cfg.RPCOverrides, "rpc-overrides", cfg.RPCOverrides, "Pubilc chain rpc overrides: '<chain1>=<url1>'")
}

func bindE2EFlags(flags *pflag.FlagSet, cfg *app.E2ETestConfig) {
	flags.BoolVar(&cfg.Preserve, "preserve", cfg.Preserve, "preserve infrastructure after test")
}

func bindPromFlags(flags *pflag.FlagSet, cfg *app.PromSecrets) {
	flags.StringVar(&cfg.URL, "prom-url", cfg.URL, "prometheus url (only required if prometheus==true)")
	flags.StringVar(&cfg.User, "prom-user", cfg.User, "prometheus user")
	flags.StringVar(&cfg.Pass, "prom-password", cfg.Pass, "prometheus password")
}

func bindDeployFlags(flags *pflag.FlagSet, cfg *app.DeployConfig) {
	bindPromFlags(flags, &cfg.PromSecrets)
	flags.StringVar(&cfg.EigenFile, "eigen-file", cfg.EigenFile, "path to json file defining eigenlayer deployments. Empty to skip AVS deployment")
}
