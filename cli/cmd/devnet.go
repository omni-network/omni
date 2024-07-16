package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"time"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/e2e/app"
	"github.com/omni-network/omni/e2e/manifests"
	"github.com/omni-network/omni/halo/genutil/evm/predeploys"
	"github.com/omni-network/omni/lib/anvil"
	"github.com/omni-network/omni/lib/buildinfo"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/txmgr"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"

	"github.com/spf13/cobra"
)

func newDevnetCmds() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "devnet",
		Short: "Local devnet commands",
		Args:  cobra.NoArgs,
	}

	cmd.AddCommand(
		newDevnetFundCmd(),
		newDevnetAVSAllow(),
		newDevnetStartCmd(),
		newDevnetInfoCmd(),
		newDevnetCleanCmd(),
	)

	return cmd
}

func newDevnetFundCmd() *cobra.Command {
	var cfg devnetFundConfig

	cmd := &cobra.Command{
		Use:   "fund",
		Short: "Fund a local devnet account with 1 ETH",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			return devnetFund(cmd.Context(), cfg)
		},
	}

	bindDevnetFundConfig(cmd, &cfg)

	return cmd
}

func newDevnetAVSAllow() *cobra.Command {
	var cfg devnetAllowConfig

	cmd := &cobra.Command{
		Use:   "avs-allow",
		Short: "Add an operator to the omni AVS allow list",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			return devnetAllow(cmd.Context(), cfg)
		},
	}

	bindDevnetAVSAllowConfig(cmd, &cfg)

	return cmd
}

func newDevnetStartCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "start",
		Short: "Build and deploy a local docker compose devnet with 2 anvil nodes and a halo node",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return deployDevnet(cmd.Context())
		},
	}
}

func newDevnetInfoCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "info",
		Short: "Display portal addresses and RPC URLs for the deployed devnet",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return printDevnetInfo(cmd.Context())
		},
	}
}

func newDevnetCleanCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "clean",
		Short: "Cleans (deletes) previously preserved devnet files and directories",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return cleanupDevnet(cmd.Context())
		},
	}
}

func cleanupDevnet(ctx context.Context) error {
	def, err := devnetDefinition(ctx)
	if err != nil {
		return err
	}

	if err := app.CleanInfra(ctx, def); err != nil {
		return err
	}

	return app.CleanupDir(ctx, def.Testnet.Dir)
}

func printDevnetInfo(ctx context.Context) error {
	// It contains correct portal addrs and external (localhost) RPCs.
	network, endpoints, err := loadDevnetNetwork(ctx)
	if err != nil {
		return errors.Wrap(err, "load internal network")
	}

	type info struct {
		ChainID       uint64         `json:"chain_id"`
		ChainName     string         `json:"chain_name"`
		PortalAddress common.Address `json:"portal_address"`
		RPCURL        string         `json:"rpc_url"`
	}

	var infos []info
	for _, chain := range network.EVMChains() {
		rpc, err := endpoints.ByNameOrID(chain.Name, chain.ID)
		if err != nil {
			return err
		}

		infos = append(infos, info{
			ChainID:       chain.ID,
			ChainName:     chain.Name,
			PortalAddress: chain.PortalAddress,
			RPCURL:        rpc,
		})
	}

	// Marshal and print the final combined JSON output
	jsonOutput, err := json.MarshalIndent(infos, "", "  ")
	if err != nil {
		return errors.Wrap(err, "marshal infos")
	}
	fmt.Println(string(jsonOutput))

	return nil
}

func devnetDefinition(ctx context.Context) (app.Definition, error) {
	manifestFile, err := writeTempFile(manifests.Devnet0())
	if err != nil {
		return app.Definition{}, err
	}

	defCfg := app.DefaultDefinitionConfig(ctx)
	defCfg.ManifestFile = manifestFile

	imgTag, ok := buildinfo.GitCommit()
	if !ok {
		imgTag = "main"
	}

	defCfg.OmniImgTag = imgTag

	def, err := app.MakeDefinition(ctx, defCfg, "devnet")
	if err != nil {
		return app.Definition{}, err
	}

	def.Testnet.Name = "devnet0"
	def.Testnet.Dir, err = homeDir(netconf.Devnet)
	if err != nil {
		return app.Definition{}, err
	}

	return def, nil
}

func loadDevnetNetwork(ctx context.Context) (netconf.Network, xchain.RPCEndpoints, error) {
	devnetPath, err := homeDir(netconf.Devnet)
	if err != nil {
		return netconf.Network{}, nil, err
	}

	endpointsFile := filepath.Join(devnetPath, "endpoints.json")
	if _, err := os.Stat(endpointsFile); os.IsNotExist(err) {
		return netconf.Network{}, nil, &CliError{
			Msg:     "failed to load ~/.omni/devnet/endpoints.json",
			Suggest: "Have you run `omni devnet start` yet?",
		}
	}

	var endpoints xchain.RPCEndpoints
	if bz, err := os.ReadFile(endpointsFile); err != nil {
		return netconf.Network{}, nil, errors.Wrap(err, "read endpoints file")
	} else if err := json.Unmarshal(bz, &endpoints); err != nil {
		return netconf.Network{}, nil, errors.Wrap(err, "unmarshal endpoints file")
	}

	netID := netconf.Devnet

	portalReg, err := makePortalRegistry(netID, endpoints)
	if err != nil {
		return netconf.Network{}, nil, errors.Wrap(err, "make portal registry")
	}

	network, err := netconf.AwaitOnChain(log.WithNoopLogger(ctx), netID, portalReg, endpoints.Keys())
	if err != nil {
		return netconf.Network{}, nil, &CliError{
			Msg:     "failed to check on-chain registry",
			Suggest: "Have you run `omni devnet start` yet?",
		}
	}

	return network, endpoints, nil
}

// deployDevnet initializes and deploys the devnet network using the e2e app.
func deployDevnet(ctx context.Context) error {
	def, err := devnetDefinition(ctx)
	if err != nil {
		return err
	}

	_, err = app.Deploy(ctx, def, app.DefaultDeployConfig())

	return err
}

func writeTempFile(content []byte) (string, error) {
	f, err := os.CreateTemp("", "")
	if err != nil {
		return "", errors.Wrap(err, "create temp file")
	}
	defer f.Close()

	if _, err := f.Write(content); err != nil {
		return "", errors.Wrap(err, "write temp manifest")
	}

	return f.Name(), nil
}

type devnetAllowConfig struct {
	OperatorAddr string
	RPCURL       string
	AVSAddr      string
}

func devnetAllow(ctx context.Context, cfg devnetAllowConfig) error {
	if !common.IsHexAddress(cfg.OperatorAddr) {
		return errors.New("invalid operator address", "address", cfg.OperatorAddr)
	}

	avsOwner, backend, err := devnetBackend(ctx, cfg.RPCURL)
	if err != nil {
		return err
	}

	chainID, err := backend.ChainID(ctx)
	if err != nil {
		return errors.Wrap(err, "get chain id")
	}

	avsAddress, err := avsAddressOrDefault(cfg.AVSAddr, chainID)
	if err != nil {
		return err
	}

	omniAVS, err := bindings.NewOmniAVS(avsAddress, backend)
	if err != nil {
		return errors.Wrap(err, "omni avs")
	}

	txOpts, err := backend.BindOpts(ctx, avsOwner)
	if err != nil {
		return err
	}

	tx, err := omniAVS.AddToAllowlist(txOpts, common.HexToAddress(cfg.OperatorAddr))
	if err != nil {
		return errors.Wrap(err, "add to allowlist")
	}

	if _, err := backend.WaitMined(ctx, tx); err != nil {
		return errors.Wrap(err, "wait mined")
	}

	log.Info(ctx, "Operator added to Omni AVS allow list", "address", cfg.OperatorAddr)

	return nil
}

type devnetFundConfig struct {
	Address string
	RPCURL  string
}

func devnetFund(ctx context.Context, cfg devnetFundConfig) error {
	if !common.IsHexAddress(cfg.Address) {
		return errors.New("invalid ETH address", "address", cfg.Address)
	}

	funder, backend, err := devnetBackend(ctx, cfg.RPCURL)
	if err != nil {
		return err
	}

	addr := common.HexToAddress(cfg.Address)
	tx, _, err := backend.Send(ctx, funder, txmgr.TxCandidate{
		To:       &addr,
		GasLimit: 100_000,
		Value:    big.NewInt(params.Ether),
	})
	if err != nil {
		return errors.Wrap(err, "send tx")
	} else if _, err := backend.WaitMined(ctx, tx); err != nil {
		return errors.Wrap(err, "wait mined")
	}

	balance, err := backend.EtherBalanceAt(ctx, addr)
	if err != nil {
		return errors.Wrap(err, "get balance")
	}

	log.Info(ctx, "Account funded", "address", cfg.Address, "balance", fmt.Sprintf("%.2f ETH", balance))

	return nil
}

// devnetBackend returns a backend populated with the default anvil account 0 private key.
func devnetBackend(ctx context.Context, rpcURL string) (common.Address, *ethbackend.Backend, error) {
	ethCl, err := ethclient.Dial("", rpcURL)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "dial eth client", "url", rpcURL)
	}

	chainID, err := ethCl.ChainID(ctx)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "get chain id")
	}

	backend, err := ethbackend.NewBackend("", chainID.Uint64(), time.Second, ethCl)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "create backend")
	}

	funderAddr, err := backend.AddAccount(anvil.DevPrivateKey0())
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "add account")
	}

	return funderAddr, backend, nil
}

func homeDir(network netconf.ID) (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", errors.Wrap(err, "user home dir")
	}

	return filepath.Join(homeDir, ".omni", network.String()), nil
}

func makePortalRegistry(network netconf.ID, endpoints xchain.RPCEndpoints) (*bindings.PortalRegistry, error) {
	meta := netconf.MetadataByID(network, network.Static().OmniExecutionChainID)
	rpc, err := endpoints.ByNameOrID(meta.Name, meta.ChainID)
	if err != nil {
		return nil, err
	}

	ethCl, err := ethclient.Dial(meta.Name, rpc)
	if err != nil {
		return nil, err
	}

	resp, err := bindings.NewPortalRegistry(common.HexToAddress(predeploys.PortalRegistry), ethCl)
	if err != nil {
		return nil, errors.Wrap(err, "create portal registry")
	}

	return resp, nil
}
