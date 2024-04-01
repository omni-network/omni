//nolint:lll // Long lines are easier to read for flag descriptions.
package cmd

import (
	"github.com/omni-network/omni/e2e/app"
	"github.com/omni-network/omni/e2e/app/agent"
	"github.com/omni-network/omni/e2e/app/key"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func bindDefFlags(flags *pflag.FlagSet, cfg *app.DefinitionConfig) {
	flags.StringVarP(&cfg.ManifestFile, "manifest-file", "f", cfg.ManifestFile, "path to manifest file")
	flags.StringVar(&cfg.InfraProvider, "infra", cfg.InfraProvider, "infrastructure provider: docker, vmcompose")
	flags.StringVar(&cfg.InfraDataFile, "infra-file", cfg.InfraDataFile, "infrastructure data file (not required for docker provider)")
	flags.StringVar(&cfg.DeployKeyFile, "deploy-key", cfg.DeployKeyFile, "path to deploy private key file")
	flags.StringVar(&cfg.RelayerKeyFile, "relayer-key", cfg.RelayerKeyFile, "path to relayer private key file")
	flags.StringVar(&cfg.FireAPIKey, "fireblocks-api-key", cfg.FireAPIKey, "FireBlocks api key")
	flags.StringVar(&cfg.FireKeyPath, "fireblocks-key-path", cfg.FireKeyPath, "FireBlocks RSA private key path")
	flags.StringVar(&cfg.OmniImgTag, "omni-image-tag", cfg.OmniImgTag, "Omni docker images tag (halo, relayer). Defaults to working dir git commit.")
	flags.StringVar(&cfg.ExplorerImageTag, "explorer-image-tag", cfg.ExplorerImageTag, "Explorer docker images tag (indexer, graphql, ui). Defaults to working dir git commit.")
	flags.StringVar(&cfg.IndexerDBConn, "indexer-db-conn-url", cfg.IndexerDBConn, "Indexer database connection url")
	flags.StringToStringVar(&cfg.RPCOverrides, "rpc-overrides", cfg.RPCOverrides, "Pubilc chain rpc overrides: '<chain1>=<url1>'")
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
	bindPromFlags(flags, &cfg.AgentSecrets)
	flags.Uint64Var(&cfg.PingPongN, "ping-pong", cfg.PingPongN, "Number of ping pongs messages to send. 0 disables it")
}

func bindCreate3DeployFlags(flags *pflag.FlagSet, cfg *app.Create3DeployConfig) {
	flags.Uint64Var(&cfg.ChainID, "chain-id", cfg.ChainID, "chain id of the chain to deploy to")
}

func bindKeyCreateFlags(cmd *cobra.Command, cfg *key.UploadConfig) {
	cmd.Flags().StringVar(&cfg.NodeName, "node", cfg.NodeName, "node name")
	cmd.Flags().StringVar((*string)(&cfg.Type), "type", string(cfg.Type), "key type: validator, p2p_execution, p2p_consensus")

	_ = cmd.MarkFlagRequired("node")
	_ = cmd.MarkFlagRequired("type")
}
