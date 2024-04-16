package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/e2e/app"
	"github.com/omni-network/omni/e2e/manifests"
	"github.com/omni-network/omni/lib/buildinfo"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/txmgr"

	"github.com/ethereum/go-ethereum/common"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"

	"github.com/spf13/cobra"
)

const (
	// privKeyHex0 of pre-funded anvil account 0.
	privKeyHex0 = "0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
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
	)

	return cmd
}

func newDevnetFundCmd() *cobra.Command {
	var cfg devnetFundConfig

	cmd := &cobra.Command{
		Use:   "fund",
		Short: "Fund a local devnet account with 1 ETH",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
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
		RunE: func(cmd *cobra.Command, args []string) error {
			return devnetAllow(cmd.Context(), cfg)
		},
	}

	bindDevnetAVSAllowConfig(cmd, &cfg)

	return cmd
}

func newDevnetStartCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "start",
		Short: "Build and deploy a local dev environment with 2 anvil nodes and a halo node using Docker",
		RunE: func(cmd *cobra.Command, args []string) error {
			return deployDevnet(cmd.Context())
		},
	}
}

func newDevnetInfoCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "info",
		Short: "Display Portal Addresses and RPC URLs for the deployed dev environment",
		RunE: func(cmd *cobra.Command, args []string) error {
			return getDevnetInfo(cmd.Context())
		},
	}
}

func getDevnetInfo(ctx context.Context) error {
	// Fetch portal data from the helper function
	portals, err := getDevnetPortalsFromNetworkJSON(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to fetch portal data: %w")
	}

	// Fetch RPC data from the helper function
	rpcs, err := getDevnetRPCsFromManifestDef(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to fetch RPC data: %w")
	}

	// Convert RPCs slice to a map for easier lookup
	rpcMap := make(map[uint64]string)
	for _, rpc := range rpcs {
		rpcMap[rpc.ID] = rpc.RPCURL
	}

	// Prepare combined info structure
	combinedInfo := make([]struct {
		ChainID       int    `json:"chain_id"`
		ServiceName   string `json:"service_name"`
		PortalAddress string `json:"portal_address"`
		RPCURL        string `json:"rpc_url"`
	}, 0)

	// Merge data based on chain ID
	for _, portal := range portals {
		rpcURL, ok := rpcMap[uint64(portal.ChainID)]
		if !ok {
			fmt.Printf("Missing RPC URL for chain_id %d\n", portal.ChainID)
			continue
		}
		combinedInfo = append(combinedInfo, struct {
			ChainID       int    `json:"chain_id"`
			ServiceName   string `json:"service_name"`
			PortalAddress string `json:"portal_address"`
			RPCURL        string `json:"rpc_url"`
		}{
			ChainID:       portal.ChainID,
			ServiceName:   portal.ServiceName,
			PortalAddress: portal.PortalAddress,
			RPCURL:        rpcURL,
		})
	}

	// Marshal and print the final combined JSON output
	jsonOutput, err := json.MarshalIndent(combinedInfo, "", "  ")
	if err != nil {
		return errors.Wrap(err, "failed to marshal combined info into JSON: %w")
	}
	fmt.Println(string(jsonOutput))

	return nil
}

type chainInfo struct {
	ID     uint64
	Name   string
	RPCURL string
	Portal string
}

func getDevnetRPCsFromManifestDef(ctx context.Context) ([]chainInfo, error) {
	manifestContent := manifests.Devnet1()
	tempManifestPath := writeTempManifest(manifestContent)
	defer os.Remove(tempManifestPath)

	//nolint:contextcheck // The function does not support context passing, ignoring.
	defCfg := app.DefaultDefinitionConfig()
	defCfg.ManifestFile = tempManifestPath
	defCfg.OmniImgTag = buildinfo.Version()
	def, err := app.MakeDefinition(ctx, defCfg, "deploy")
	if err != nil {
		return nil, err
	}

	var chains []chainInfo

	omniEVM := def.Testnet.OmniEVMs[1]
	chains = append(chains, chainInfo{
		ID:     omniEVM.Chain.ID,
		Name:   omniEVM.Chain.Name,
		RPCURL: omniEVM.ExternalRPC,
	})

	for _, anvil := range def.Testnet.AnvilChains {
		chains = append(chains, chainInfo{
			ID:     anvil.Chain.ID,
			Name:   anvil.Chain.Name,
			RPCURL: anvil.ExternalRPC,
		})
	}

	return chains, nil
}

// chainJSONInfo is a structure to hold the simplified chain information.
type chainJSONInfo struct {
	ChainID       int    `json:"chain_id"`
	ServiceName   string `json:"service_name"`
	PortalAddress string `json:"portal_address"`
}

func getDevnetPortalsFromNetworkJSON(_ context.Context) ([]chainJSONInfo, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get user home directory: %w")
	}
	devnetPath := filepath.Join(homeDir, ".omni", "devnet")
	if _, err := os.Stat(devnetPath); os.IsNotExist(err) {
		return nil, errors.Wrap(err, "no config files detected. Have you run `omni devnet start` yet?")
	}
	jsonFilePath := filepath.Join(devnetPath, "validator01", "config", "network.json")

	// Read the JSON file using os.ReadFile
	jsonData, err := os.ReadFile(jsonFilePath)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read JSON file: %w")
	}

	// Parse the JSON data
	var data struct {
		Chains []struct {
			ID            int    `json:"id"`
			Name          string `json:"name"`
			PortalAddress string `json:"portal_address"`
		} `json:"chains"`
	}

	if err := json.Unmarshal(jsonData, &data); err != nil {
		return nil, errors.Wrap(err, "failed to parse JSON data: %w")
	}

	// Filter and collect the necessary chain info
	var results []chainJSONInfo
	for _, chain := range data.Chains {
		if chain.PortalAddress != "" {
			results = append(results, chainJSONInfo{
				ChainID:       chain.ID,
				ServiceName:   chain.Name,
				PortalAddress: chain.PortalAddress,
			})
		}
	}

	return results, nil
}

// deployDevnetNetwork initializes and deploys the devnet network using the e2e app.
func deployDevnet(ctx context.Context) error {
	manifestContent := manifests.Devnet1()

	tempManifestPath := writeTempManifest(manifestContent)
	defer os.Remove(tempManifestPath)

	//nolint:contextcheck // The function does not support context passing, ignoring.
	defCfg := app.DefaultDefinitionConfig()
	defCfg.ManifestFile = tempManifestPath
	defCfg.OmniImgTag = buildinfo.Version()
	def, err := app.MakeDefinition(ctx, defCfg, "deploy")
	if err != nil {
		return err
	}

	// Retrieve the home directory from the environment variable
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return errors.Wrap(err, "failed to get user home directory")
	}
	def.Testnet.Dir = filepath.Join(homeDir, ".omni", "devnet") // Use filepath to correctly handle paths

	deployCfg := app.DefaultDeployConfig()
	_, err = app.Deploy(ctx, def, deployCfg)

	return err
}

func writeTempManifest(content []byte) string {
	tempFile, err := os.CreateTemp("", "devnet_manifest_*.toml")
	if err != nil {
		panic(fmt.Errorf("failed to create temp manifest file: %w", err))
	}
	defer tempFile.Close()

	if _, err := tempFile.Write(content); err != nil {
		panic(fmt.Errorf("failed to write to temp manifest file: %w", err))
	}

	return tempFile.Name()
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

	b, err := backend.BalanceAt(ctx, addr, nil)
	if err != nil {
		return errors.Wrap(err, "get balance")
	}

	bf, _ := b.Float64()
	bf /= params.Ether

	log.Info(ctx, "Account funded", "address", cfg.Address, "balance", fmt.Sprintf("%.2f ETH", bf))

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

	funderPrivKey, err := ethcrypto.HexToECDSA(strings.TrimPrefix(privKeyHex0, "0x"))
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "parse private key")
	}

	backend, err := ethbackend.NewBackend("", chainID.Uint64(), time.Second, ethCl)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "create backend")
	}

	funderAddr, err := backend.AddAccount(funderPrivKey)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "add account")
	}

	return funderAddr, backend, nil
}
