package cmd

import (
	"context"
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

// deployDevnetNetwork initializes and deploys the devnet network using the e2e app.
func deployDevnet(ctx context.Context) error {
	manifestContent := manifests.Devnet1()

	tempManifestPath := writeTempManifest(manifestContent)
	defer os.Remove(tempManifestPath)

	//nolint:contextcheck // The function does not support context passing, ignoring.
	defCfg := app.DefaultDefinitionConfig()
	defCfg.ManifestFile = tempManifestPath
	defCfg.OmniImgTag = buildinfo.Version()
	def, err := app.MakeDefinition(ctx, defCfg, "deploy") // holds dir var
	if err != nil {
		return err
	}

	// Retrieve the home directory from the environment variable
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return errors.Wrap(err, "failed to get user home directory")
	}
	def.Testnet.Dir = filepath.Join(homeDir, ".omni_devnet") // Use filepath to correctly handle paths

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
