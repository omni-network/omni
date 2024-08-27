package cmd

import (
	"encoding/hex"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"

	"github.com/ethereum/go-ethereum/crypto"

	"github.com/spf13/cobra"
)

func newCreateKeyCmd() *cobra.Command {
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
				"pubkey", hex.EncodeToString(crypto.CompressPubkey(&privKey.PublicKey)),
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
