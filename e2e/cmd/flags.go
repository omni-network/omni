//nolint:lll // Long lines are easier to read for flag descriptions.
package cmd

import (
	"github.com/omni-network/omni/e2e/app"
	"github.com/omni-network/omni/e2e/app/admin"
	"github.com/omni-network/omni/e2e/app/agent"
	"github.com/omni-network/omni/e2e/app/key"
	"github.com/omni-network/omni/e2e/types"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func bindDefFlags(flags *pflag.FlagSet, cfg *app.DefinitionConfig) {
	bindPromFlags(flags, &cfg.AgentSecrets)
	flags.StringVarP(&cfg.ManifestFile, "manifest-file", "f", cfg.ManifestFile, "path to manifest file")
	flags.StringVar(&cfg.InfraProvider, "infra", cfg.InfraProvider, "infrastructure provider: docker, vmcompose")
	flags.StringVar(&cfg.InfraDataFile, "infra-file", cfg.InfraDataFile, "infrastructure data file (not required for docker provider)")
	flags.StringVar(&cfg.DeployKeyFile, "deploy-key", cfg.DeployKeyFile, "path to deploy private key file")
	flags.StringVar(&cfg.FireAPIKey, "fireblocks-api-key", cfg.FireAPIKey, "FireBlocks api key")
	flags.StringVar(&cfg.FireKeyPath, "fireblocks-key-path", cfg.FireKeyPath, "FireBlocks RSA private key path")
	flags.StringVar(&cfg.OmniImgTag, "omni-image-tag", cfg.OmniImgTag, "Omni docker images tag (halo, relayer). Defaults to working dir git commit.")
	flags.StringToStringVar(&cfg.RPCOverrides, "rpc-overrides", cfg.RPCOverrides, "Public chain rpc overrides: '<chain1>=<url1>,<url2>'")
	flags.StringVar(&cfg.TracingEndpoint, "tracing-endpoint", cfg.TracingEndpoint, "Tracing endpoint")
	flags.StringVar(&cfg.TracingHeaders, "tracing-headers", cfg.TracingHeaders, "Tracing headers")
}

func bindE2EFlags(flags *pflag.FlagSet, cfg *app.E2ETestConfig) {
	flags.BoolVar(&cfg.Preserve, "preserve", cfg.Preserve, "preserve infrastructure after test")
}

func bindPromFlags(flags *pflag.FlagSet, cfg *agent.Secrets) {
	flags.StringVar(&cfg.URL, "prom-url", cfg.URL, "prometheus url (only required if prometheus==true)")
	flags.StringVar(&cfg.User, "prom-user", cfg.User, "prometheus user")
	flags.StringVar(&cfg.Pass, "prom-password", cfg.Pass, "prometheus password")
}

func bindDeployFlags(flags *pflag.FlagSet, cfg *app.DeployConfig) {
	flags.Uint64Var(&cfg.PingPongN, "ping-pong", cfg.PingPongN, "Number of ping pongs messages to send. 0 disables it")
}

func bindServiceFlags(flags *pflag.FlagSet, cfg *types.ServiceConfig) {
	flags.StringVar(&cfg.Regexp, "services", cfg.Regexp, "Regexp applied to services per VM. Any match results in the VM being upgraded/restarted (all services on that VM are upgraded/restarted, not only matching services)")
}

func bindCreate3DeployFlags(flags *pflag.FlagSet, cfg *app.Create3DeployConfig) {
	flags.Uint64Var(&cfg.ChainID, "chain-id", cfg.ChainID, "chain id of the chain to deploy to")
}

func bindKeyCreateFlags(cmd *cobra.Command, cfg *key.UploadConfig) {
	cmd.Flags().StringVar(&cfg.Name, "name", cfg.Name, "key name: either node name or eoa account type")
	cmd.Flags().StringVar((*string)(&cfg.Type), "type", string(cfg.Type), "key type: validator, p2p_execution, p2p_consensus, eoa")

	_ = cmd.MarkFlagRequired("name")
	_ = cmd.MarkFlagRequired("type")
}

func bindPortalAdminFlags(flags *pflag.FlagSet, cfg *admin.PortalAdminConfig) {
	flags.StringVar(&cfg.Chain, "chain", cfg.Chain, "Run admin command on a specific chain (\"--chain=all\" for all chains)")
}
