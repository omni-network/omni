package cmd

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/omni-network/omni/halo/app"
	"github.com/omni-network/omni/halo/attest"
	libcmd "github.com/omni-network/omni/lib/cmd"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"

	cmtconfig "github.com/cometbft/cometbft/config"
	k1 "github.com/cometbft/cometbft/crypto/secp256k1"
	cmtos "github.com/cometbft/cometbft/libs/os"
	cmtrand "github.com/cometbft/cometbft/libs/rand"
	"github.com/cometbft/cometbft/p2p"
	"github.com/cometbft/cometbft/privval"
	"github.com/cometbft/cometbft/types"
	cmttime "github.com/cometbft/cometbft/types/time"

	"github.com/spf13/cobra"
)

// newInitCmd returns a new cobra command that initializes the files and folders required by halo.
func newInitCmd() *cobra.Command {
	var (
		homeDir = app.DefaultHomeDir
		force   bool
	)

	cmd := &cobra.Command{
		Use:   "init",
		Short: "Initializes required halo files and directories",
		Long: `Initializes required halo files and directories.

Ensures all the following files and directories exist:
  <home>/                            # Halo home directory
  ├── config                         # Config directory
  │   ├── config.toml                # CometBFT configuration
  │   ├── genesis.json               # Omni chain genesis file
  │   ├── halo.toml                  # Halo configuration
  │   ├── node_key.json              # Node P2P identity key
  │   └── priv_validator_key.json    # CometBFT private validator key (back this up and keep it safe)
  ├── data                           # Data directory
  │   ├── priv_validator_state.json  # CometBFT private validator state (slashing protection)
  │   ├── snapshots                  # Snapshot directory
  │   └── xattestations_state.json   # Cross chain attestation state (slashing protection)

Existing files are not overwritten.
The home directory should only contain subdirectories, no files, use --force to ignore this check.
`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return initFiles(cmd.Context(), homeDir, force)
		},
	}

	libcmd.BindHomeFlag(cmd.Flags(), &homeDir)
	cmd.Flags().BoolVar(&force, "force", false, "Force initialization even if home directory contains files")

	return cmd
}

// initFiles initializes the files and folders required by halo.
// If no other genesis file exists, it will create a devnet genesis file.
//
//nolint:revive // Force bool is ok.
func initFiles(ctx context.Context, homeDir string, force bool) error {
	log.Info(ctx, "Initializing files and directories", "home", homeDir, "force", force)

	// Quick sanity check if --home contains files (it should only contain dirs).
	// This prevents accidental initialization in wrong current dir.
	if !force {
		files, _ := os.ReadDir(homeDir) // Ignore error, we'll just assume it's empty.
		for _, file := range files {
			if file.IsDir() { // Ignore directories
				continue
			}

			return errors.New("home directory contains unexpected file(s), use --force to initialize anyway",
				"home", homeDir, "example_file", file.Name())
		}
	}

	// Initialize default configs.
	comet := defaultCometConfig(homeDir)
	cfg := app.DefaultHaloConfig()
	cfg.HomeDir = homeDir

	// Folders
	folders := []struct {
		Name string
		Path string
	}{
		{"home", homeDir},
		{"data", filepath.Join(homeDir, cmtconfig.DefaultDataDir)},
		{"config", filepath.Join(homeDir, cmtconfig.DefaultConfigDir)},
		{"comet db", comet.DBDir()},
		{"snapshot", cfg.SnapshotDir()},
		{"app db", cfg.AppStateDir()},
	}
	for _, folder := range folders {
		if cmtos.FileExists(folder.Path) {
			// Dir exists, just skip
			continue
		} else if err := cmtos.EnsureDir(folder.Path, 0o755); err != nil {
			return errors.Wrap(err, "create folder")
		}
		log.Info(ctx, "Generated folder", "reason", folder.Name, "path", folder.Path)
	}

	// Setup comet config
	cmtConfigFile := filepath.Join(homeDir, cmtconfig.DefaultConfigDir, cmtconfig.DefaultConfigFileName)
	if cmtos.FileExists(cmtConfigFile) {
		log.Info(ctx, "Found comet config file", "path", cmtConfigFile)
	} else {
		cmtconfig.WriteConfigFile(cmtConfigFile, &comet) // This panics on any error :(
		log.Info(ctx, "Generated default comet config file", "path", cmtConfigFile)
	}

	// Setup halo config
	haloConfigFile := cfg.ConfigFile()
	if cmtos.FileExists(haloConfigFile) {
		log.Info(ctx, "Found halo config file", "path", haloConfigFile)
	} else if err := app.WriteConfigTOML(cfg); err != nil {
		return err
	} else {
		log.Info(ctx, "Generated default halo config file", "path", haloConfigFile)
	}

	// Setup comet private validator
	var pv *privval.FilePV
	privValKeyFile := comet.PrivValidatorKeyFile()
	privValStateFile := comet.PrivValidatorStateFile()
	if cmtos.FileExists(privValKeyFile) {
		pv = privval.LoadFilePV(privValKeyFile, privValStateFile) // This hard exits on any error.
		log.Info(ctx, "Found private validator",
			"key_file", privValKeyFile,
			"state_file", privValStateFile,
		)
	} else {
		pv = privval.NewFilePV(k1.GenPrivKey(), privValKeyFile, privValStateFile)
		pv.Save()
		log.Info(ctx, "Generated private validator",
			"key_file", privValKeyFile,
			"state_file", privValStateFile)
	}

	// Setup node key
	nodeKeyFile := comet.NodeKeyFile()
	if cmtos.FileExists(nodeKeyFile) {
		log.Info(ctx, "Found node key", "path", nodeKeyFile)
	} else if _, err := p2p.LoadOrGenNodeKey(nodeKeyFile); err != nil {
		return errors.Wrap(err, "load or generate node key")
	} else {
		log.Info(ctx, "Generated node key", "path", nodeKeyFile)
	}

	// Setup genesis file
	genFile := comet.GenesisFile()
	if cmtos.FileExists(genFile) {
		log.Info(ctx, "Found genesis file", "path", genFile)
	} else {
		// Create a devnet genesis file with this node as single validator.
		genDoc := types.GenesisDoc{
			ChainID:         fmt.Sprintf("dev-%s", cmtrand.Str(6)),
			GenesisTime:     cmttime.Now(),
			ConsensusParams: types.DefaultConsensusParams(),
		}
		pubKey, err := pv.GetPubKey()
		if err != nil {
			return errors.Wrap(err, "get public key")
		}

		const nonZeroPower = 10 // Use any non-zero power for this single validator.
		genDoc.Validators = []types.GenesisValidator{{
			Address: pubKey.Address(),
			PubKey:  pubKey,
			Power:   nonZeroPower,
		}}

		if err := genDoc.SaveAs(genFile); err != nil {
			return errors.Wrap(err, "save genesis file")
		}
		log.Info(ctx, "Generated devnet genesis file", "path", genFile)
	}

	// Attest state
	attStateFile := cfg.AttestStateFile()
	if cmtos.FileExists(attStateFile) {
		log.Info(ctx, "Found attest state file", "path", attStateFile)
	} else if err := attest.GenEmptyStateFile(attStateFile); err != nil {
		return err
	} else {
		log.Info(ctx, "Generated attest state file", "path", attStateFile)
	}

	return nil
}
