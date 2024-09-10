package cmd

import (
	"os"
	"path"

	libcmd "github.com/omni-network/omni/lib/cmd"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"

	"github.com/cometbft/cometbft/config"
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
			"cometBFT paths: `<home>/config/priv_validator_key.json` and `<home>/data/priv_validator_state.json`",
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			cfg := config.DefaultConfig()
			cfg.RootDir = home

			filePV := privval.NewFilePV(k1.GenPrivKey(), cfg.PrivValidatorKeyFile(), cfg.PrivValidatorStateFile())

			if err := os.MkdirAll(path.Dir(cfg.PrivValidatorKeyFile()), 0o755); err != nil {
				return errors.Wrap(err, "ensure config dir")
			}
			if err := os.MkdirAll(path.Dir(cfg.PrivValidatorStateFile()), 0o755); err != nil {
				return errors.Wrap(err, "ensure data dir")
			}
			if err := ensureNotExist(cfg.PrivValidatorKeyFile()); err != nil {
				return &CliError{Msg: "existing private key file found: " + cfg.PrivValidatorKeyFile()}
			}
			if err := ensureNotExist(cfg.PrivValidatorStateFile()); err != nil {
				return &CliError{Msg: "existing state file found: " + cfg.PrivValidatorKeyFile()}
			}

			// CometBFT panics instead of error :(
			err := func() (err error) {
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

			log.Info(cmd.Context(), "🎉 Created consensus private key",
				"home", home,
				"private_key", cfg.PrivValidatorKeyFile(),
				"state_file", cfg.PrivValidatorStateFile(),
			)
			log.Info(cmd.Context(), "🚧 Remember to backup the private key if the node is a validator 🚧")

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
			if err := crypto.SaveECDSA(cfg.PrivateKeyFile, privKey); err != nil {
				return errors.Wrap(err, "save key")
			}

			log.Info(cmd.Context(), "🎉 Created operator private key",
				"type", cfg.Type,
				"file", cfg.PrivateKeyFile,
				"address", crypto.PubkeyToAddress(privKey.PublicKey).Hex(),
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
		PrivateKeyFile: "./operator-private-key",
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
