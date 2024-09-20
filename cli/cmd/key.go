package cmd

import (
	"encoding/hex"
	"os"
	"path"
	"strings"

	"github.com/omni-network/omni/halo/attest/voter"
	halocmd "github.com/omni-network/omni/halo/cmd"
	halocfg "github.com/omni-network/omni/halo/config"
	libcmd "github.com/omni-network/omni/lib/cmd"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"

	k1 "github.com/cometbft/cometbft/crypto/secp256k1"
	"github.com/cometbft/cometbft/privval"

	"github.com/ethereum/go-ethereum/crypto"

	"github.com/spf13/cobra"
)

func newCreateConsensusKeyCmd() *cobra.Command {
	var home string
	cmd := &cobra.Command{
		Use:   "create-consensus-key",
		Short: "Create new CometBFT consensus private key and state files",
		Long: "Create new CometBFT consensus private key and state files " +
			"used for P2P consensus and xchain attestation. It is created in the default " +
			"cometBFT paths: `<home>/config/priv_validator_key.json` " +
			"and `<home>/data/priv_validator_state.json` " +
			"and `<home>/data/voter_state.json`",
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			// Initialize comet config.
			cmtCfg := halocmd.DefaultCometConfig(home)

			// Initialize halo config.
			haloCfg := halocfg.DefaultConfig()
			haloCfg.HomeDir = home

			keyFile := cmtCfg.PrivValidatorKeyFile()
			stateFile := cmtCfg.PrivValidatorStateFile()
			voterFile := haloCfg.VoterStateFile()

			for _, file := range []string{keyFile, stateFile, voterFile} {
				if err := os.MkdirAll(path.Dir(file), 0o755); err != nil {
					return errors.Wrap(err, "ensure dir")
				}
				if err := ensureNotExist(file); err != nil {
					return &CliError{Msg: "existing file found: " + file}
				}
			}

			filePV := privval.NewFilePV(k1.GenPrivKey(), keyFile, stateFile)
			pubkey, err := filePV.GetPubKey()
			if err != nil {
				return errors.Wrap(err, "pubkey")
			}
			pubkeyHex := hex.EncodeToString(pubkey.Bytes())

			// CometBFT panics instead of error :(
			err = func() (err error) {
				defer func() {
					if r := recover(); r != nil {
						if e, ok := r.(error); ok {
							err = errors.Wrap(e, "save private validator key files")
						} else {
							err = errors.New("save private validator key files", "err", r)
						}
					}
				}()

				filePV.Save()

				return nil
			}()
			if err != nil {
				return err
			}

			if err := voter.GenEmptyStateFile(voterFile); err != nil {
				return err
			}

			ctx := cmd.Context()
			log.Info(ctx, "Created consensus voter state file", "path", voterFile)
			log.Info(ctx, "Created consensus private validator state file", "path", stateFile)
			log.Info(ctx, "Created consensus private key", "path", keyFile, "pubkey", pubkeyHex)
			log.Info(ctx, "🚧 Remember to backup the private key if the node is a validator 🚧")

			return nil
		},
	}

	libcmd.BindHomeFlag(cmd.Flags(), &home)

	return cmd
}

func newCreateOperatorKeyCmd() *cobra.Command {
	cfg := defaultCreateKeyConfig()
	cmd := &cobra.Command{
		Use:   "create-operator-key",
		Short: "Create new operator key",
		Long:  `Create new operator EOA private key used to fund and sign operator staking transactions`,
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			if err := cfg.Verify(); err != nil {
				return errors.Wrap(err, "verify flags")
			}

			privKey, err := crypto.GenerateKey()
			if err != nil {
				return errors.Wrap(err, "generate key")
			}

			addr := crypto.PubkeyToAddress(privKey.PublicKey).Hex()
			privKeyfile := strings.Replace(cfg.PrivateKeyFile, "{ADDRESS}", addr, 1)

			if err := crypto.SaveECDSA(privKeyfile, privKey); err != nil {
				return errors.Wrap(err, "save key")
			}

			log.Info(cmd.Context(), "🎉 Created operator private key",
				"type", cfg.Type,
				"file", privKeyfile,
				"address", addr,
			)
			log.Info(cmd.Context(), "🚧 Remember to backup this key 🚧")

			return nil
		},
	}

	bindCreateKeyConfig(cmd, &cfg)

	return cmd
}

type keyType string

const (
	KeyTypeInsecure keyType = "insecure"
)

func (t keyType) Verify() error {
	if t != KeyTypeInsecure {
		return errors.New("invalid key type")
	}

	return nil
}

type createKeyConfig struct {
	Type           keyType
	PrivateKeyFile string
}

func (c createKeyConfig) Verify() error {
	if err := c.Type.Verify(); err != nil {
		return errors.Wrap(err, "verify --type")
	}

	if c.PrivateKeyFile == "" {
		return errors.New("required flag --output-file not set")
	}

	return nil
}

func defaultCreateKeyConfig() createKeyConfig {
	return createKeyConfig{
		Type:           KeyTypeInsecure,
		PrivateKeyFile: "./operator-private-key-{ADDRESS}",
	}
}

func ensureNotExist(file string) error {
	_, err := os.Stat(file)
	if err == nil {
		return errors.New("file exists", "path", file)
	} else if !os.IsNotExist(err) {
		return errors.Wrap(err, "unexpected error")
	}

	return nil
}
