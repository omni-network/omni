package main

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/omni-network/omni/anchor/anchorinbox"
	libcmd "github.com/omni-network/omni/lib/cmd"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/svmutil"

	"github.com/gagliardetto/solana-go"
	"github.com/spf13/cobra"
)

func main() {
	var dir string
	cmd := libcmd.NewRootCmd("localsvm", "Local SVM server with inbox")
	cmd.RunE = func(cmd *cobra.Command, _ []string) error {
		ctx := cmd.Context()

		cfg := log.DefaultConfig()
		cfg.Level = log.LevelDebug
		ctx, err := log.Init(ctx, cfg)
		if err != nil {
			return err
		}

		return run(ctx, dir)
	}
	cmd.Flags().StringVar(&dir, "dir", "/tmp/svm", "Directory to use for SVM data")

	libcmd.Main(cmd)
}

func run(ctx context.Context, dir string) error {
	if err := os.RemoveAll(dir); err != nil && !os.IsNotExist(err) {
		return errors.Wrap(err, "remove SVM directory")
	}
	if err := os.MkdirAll(dir, 0755); err != nil {
		return errors.Wrap(err, "create SVM directory")
	}
	log.Info(ctx, "Using new temporary directory for SVM", "dir", dir)

	prog := anchorinbox.Program()

	cl, addr, privkey, cancel, err := svmutil.Start(ctx, dir, prog)
	if err != nil {
		return errors.Wrap(err, "start SVM client")
	}
	defer cancel()

	log.Info(ctx, "Creating USDC mint...")
	mintResp, err := svmutil.CreateMint(ctx, cl, privkey, "usdc", 6)
	if err != nil {
		return errors.Wrap(err, "create mint")
	}

	log.Info(ctx, "Deploying anchor inbox program...")
	_, err = svmutil.Deploy(ctx, cl, dir, prog)
	if err != nil {
		return errors.Wrap(err, "deploy anchor inbox program")
	}

	if err := dumpConfig(ctx, dir, addr, privkey, mintResp, prog); err != nil {
		return err
	}

	log.Info(ctx, "SVM node is running, press Ctrl+C to stop it")
	<-ctx.Done()
	log.Info(ctx, "Stopping SVM node...")

	return nil
}

func dumpConfig(
	ctx context.Context,
	dir string,
	addr string,
	authorityKey solana.PrivateKey,
	mint svmutil.CreateMintResp,
	program svmutil.Program) error {
	type mintConfig struct {
		Symbol           string `json:"symbol"`
		MintAccount      string `json:"mint_account"`
		Authority        string `json:"authority"`
		AuthTokenAccount string `json:"authority_token_account"`
	}
	type programConfig struct {
		Name    string `json:"name"`
		Account string `json:"account"`
	}

	config := struct {
		RPCAddress   string          `json:"rpc_address"`
		AuthorityKey string          `json:"authority_key"`
		Mints        []mintConfig    `json:"mints"`
		Programs     []programConfig `json:"programs"`
	}{
		RPCAddress:   addr,
		AuthorityKey: authorityKey.String(),
		Mints: []mintConfig{{
			Symbol:           mint.Symbol,
			MintAccount:      mint.MintAccount.String(),
			AuthTokenAccount: mint.AuthATA.String(),
			Authority:        mint.Authority.PublicKey().String(),
		}},
		Programs: []programConfig{{
			Name:    program.Name,
			Account: program.MustPublicKey().String(),
		}},
	}
	content, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return errors.Wrap(err, "marshal config")
	}
	configFile := filepath.Join(dir, "config.json")
	if err := os.WriteFile(configFile, content, 0644); err != nil {
		return errors.Wrap(err, "write config file")
	}

	log.Info(ctx, "Dumped SVM config file", "file", configFile)

	return nil
}
