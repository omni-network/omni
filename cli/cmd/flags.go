package cmd

import (
	"github.com/omni-network/omni/lib/netconf"

	"github.com/spf13/cobra"
)

func bindRegConfig(cmd *cobra.Command, cfg *RegConfig) {
	bindAVSAddress(cmd, &cfg.AVSAddr)

	const flagConfig = "config-file"
	cmd.Flags().StringVar(&cfg.ConfigFile, flagConfig, cfg.ConfigFile, "Path to the Eigen-Layer yaml configuration file")
	_ = cmd.MarkFlagRequired(flagConfig)
}

func bindInitConfig(cmd *cobra.Command, cfg *initConfig) {
	netconf.BindFlag(cmd.Flags(), &cfg.Network)
	cmd.Flags().StringVar(&cfg.Moniker, "moniker", "", "Human-readable node name used in p2p networking")
	cmd.Flags().StringVar(&cfg.Home, "home", "", "Home directory. If empty, defaults to: $HOME/.omni/<network>/")
	cmd.Flags().BoolVar(&cfg.Clean, "clean", cfg.Clean, "Delete contents of home directory")
}

func bindAVSAddress(cmd *cobra.Command, addr *string) {
	cmd.Flags().StringVar(addr, "avs-address", *addr, "Optional address of the Omni AVS contract")
}

func bindDeveloperForgeProjectConfig(cmd *cobra.Command, cfg *developerForgeProjectConfig) {
	cmd.Flags().StringVar(&cfg.templateName, "template", defaultTemplate, "Name of the forge template repo to use found in the omni-network github organization")
}

func bindDevnetAVSAllowConfig(cmd *cobra.Command, cfg *devnetAllowConfig) {
	bindRPCURL(cmd, &cfg.RPCURL)
	bindAVSAddress(cmd, &cfg.AVSAddr)

	const flagOperator = "operator"
	cmd.Flags().StringVar(&cfg.OperatorAddr, flagOperator, cfg.OperatorAddr, "Operator address to allow")
	_ = cmd.MarkFlagRequired(flagOperator)
}

func bindDevnetFundConfig(cmd *cobra.Command, d *devnetFundConfig) {
	bindRPCURL(cmd, &d.RPCURL)

	const flagAddress = "address"
	cmd.Flags().StringVar(&d.Address, flagAddress, d.Address, "Address to fund")
	_ = cmd.MarkFlagRequired(flagAddress)
}

func bindRPCURL(cmd *cobra.Command, rpcURL *string) {
	const flagRPCURL = "rpc-url"
	cmd.Flags().StringVar(rpcURL, flagRPCURL, *rpcURL, "URL of the eth-json RPC server")
	_ = cmd.MarkFlagRequired(flagRPCURL)
}
