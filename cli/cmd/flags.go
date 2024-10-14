package cmd

import (
	"github.com/omni-network/omni/lib/netconf"

	"github.com/spf13/cobra"
)

const (
	flagPrivateKeyFile   = "private-key-file"
	flagConsPubKeyHex    = "consensus-pubkey-hex"
	flagSelfDelegation   = "self-delegation"
	flagDelegationAmount = "amount"
	flagNetwork          = "network"
	flagConfig           = "config-file"
	flagOperator         = "operator"
	flagRPCURL           = "rpc-url"
	flagAddress          = "address"
	flagType             = "type"
)

func bindRegConfig(cmd *cobra.Command, cfg *RegConfig) {
	bindAVSAddress(cmd, &cfg.AVSAddr)

	cmd.Flags().StringVar(&cfg.ConfigFile, flagConfig, cfg.ConfigFile, "Path to the Eigen-Layer yaml configuration file")
	_ = cmd.MarkFlagRequired(flagConfig)
}

func bindInitConfig(cmd *cobra.Command, cfg *initConfig) {
	netconf.BindFlag(cmd.Flags(), &cfg.Network)
	cmd.Flags().StringVar(&cfg.Moniker, "moniker", "", "Human-readable node name used in p2p networking")
	cmd.Flags().StringVar(&cfg.Home, "home", "", "Home directory. If empty, defaults to: $HOME/.omni/<network>/")
	cmd.Flags().BoolVar(&cfg.Clean, "clean", cfg.Clean, "Delete contents of home directory")
	cmd.Flags().BoolVar(&cfg.Archive, "archive", cfg.Archive, "Enable archive mode. Note this requires more disk space")
	cmd.Flags().BoolVar(&cfg.Debug, "debug", cfg.Debug, "Configure nodes with debug log level")
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

	cmd.Flags().StringVar(&cfg.OperatorAddr, flagOperator, cfg.OperatorAddr, "Operator address to allow")
	_ = cmd.MarkFlagRequired(flagOperator)
}

func bindDevnetFundConfig(cmd *cobra.Command, d *devnetFundConfig) {
	bindRPCURL(cmd, &d.RPCURL)

	cmd.Flags().StringVar(&d.Address, flagAddress, d.Address, "Address to fund")
	_ = cmd.MarkFlagRequired(flagAddress)
}

func bindRPCURL(cmd *cobra.Command, rpcURL *string) {
	cmd.Flags().StringVar(rpcURL, flagRPCURL, *rpcURL, "URL of the eth-json RPC server")
	_ = cmd.MarkFlagRequired(flagRPCURL)
}

func bindUnjailConfig(cmd *cobra.Command, cfg *unjailConfig) {
	netconf.BindFlag(cmd.Flags(), &cfg.Network)
	bindPrivateKeyFile(cmd, &cfg.PrivateKeyFile)
	_ = cmd.MarkFlagRequired(flagNetwork)
}

func bindDelegateValConfig(cmd *cobra.Command, cfg *delegateValConfig) {
	netconf.BindFlag(cmd.Flags(), &cfg.network)
	bindPrivateKeyFile(cmd, &cfg.privateKeyFile)

	cmd.Flags().Uint64Var(&cfg.amount, flagDelegationAmount, cfg.amount, "increase delegation amount in OMNI (minimum 1 OMNI)")

	bindRPCs(cmd, &cfg.evmRPC, &cfg.consensusRPC)

	_ = cmd.MarkFlagRequired(flagConsPubKeyHex)
	_ = cmd.MarkFlagRequired(flagDelegationAmount)
	_ = cmd.MarkFlagRequired(flagNetwork)
}

func bindCreateValConfig(cmd *cobra.Command, cfg *createValConfig) {
	netconf.BindFlag(cmd.Flags(), &cfg.network)
	bindPrivateKeyFile(cmd, &cfg.privateKeyFile)

	cmd.Flags().StringVar(&cfg.consensusPubKeyHex, flagConsPubKeyHex, cfg.consensusPubKeyHex, "Hex-encoded validator consensus public key")
	cmd.Flags().Uint64Var(&cfg.selfDelegation, flagSelfDelegation, cfg.selfDelegation, "Self-delegation amount in OMNI (minimum 100 OMNI)")

	bindRPCs(cmd, &cfg.evmRPC, &cfg.consensusRPC)

	_ = cmd.MarkFlagRequired(flagConsPubKeyHex)
	_ = cmd.MarkFlagRequired(flagSelfDelegation)
	_ = cmd.MarkFlagRequired(flagNetwork)
}

func bindRPCs(cmd *cobra.Command, evmRPC *string, consensusRPC *string) {
	cmd.Flags().StringVar(evmRPC, "evm-rpc", "", "Omni EVM execution RPC API endpoint")
	cmd.Flags().StringVar(consensusRPC, "consensus-rpc", "", "Omni consensus RPC API endpoint")
}

func bindPrivateKeyFile(cmd *cobra.Command, privateKeyFile *string) {
	cmd.Flags().StringVar(privateKeyFile, flagPrivateKeyFile, *privateKeyFile, "Path to the private key file")
	_ = cmd.MarkFlagRequired(flagPrivateKeyFile)
}

func bindCreateKeyConfig(cmd *cobra.Command, cfg *createKeyConfig) {
	cmd.Flags().StringVar((*string)(&cfg.Type), flagType, string(cfg.Type), "Type of key to create")
	cmd.Flags().StringVar(&cfg.PrivateKeyFile, "output-file", cfg.PrivateKeyFile, "Path to output private key file. Note that '{ADDRESS}' will be replaced with the address")
}
