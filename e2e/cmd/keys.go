package cmd

import (
	"bytes"
	"fmt"
	"io"
	"reflect"

	"github.com/omni-network/omni/e2e/app"
	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/e2e/app/key"
	"github.com/omni-network/omni/e2e/types"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/common"

	"github.com/naoina/toml"
	"github.com/spf13/cobra"
)

func newKeyCreate(def *app.Definition) *cobra.Command {
	cfg := key.UploadConfig{}

	cmd := &cobra.Command{
		Use:   "key-create",
		Short: "Creates a private key in GCP secret manager for a node in a manifest",
		RunE: func(cmd *cobra.Command, _ []string) error {
			if def.Testnet.Network == netconf.Simnet || def.Testnet.Network == netconf.Devnet {
				return errors.New("cannot create keys for simnet or devnet")
			}

			cfg.Network = def.Testnet.Network

			if err := verifyKeyNodeType(*def, cfg); err != nil {
				return errors.Wrap(err, "verify cfg", "name", cfg.Name, "type", cfg.Type)
			}

			_, err := key.UploadNew(cmd.Context(), cfg)

			return err
		},
	}

	bindKeyCreateFlags(cmd, &cfg)

	return cmd
}

func newKeyCreateAll(def *app.Definition) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "key-create-all",
		Short: "Creates (if required) all private keys in GCP secret manager for nodes in a manifest",
		RunE: func(cmd *cobra.Command, _ []string) error {
			if def.Testnet.Network == netconf.Simnet || def.Testnet.Network == netconf.Devnet {
				return errors.New("cannot create keys for simnet or devnet")
			}

			ctx := cmd.Context()
			network := def.Testnet.Network
			manifest := def.Manifest
			if manifest.Keys == nil {
				manifest.Keys = make(map[string]map[key.Type]string)
			}

			for _, cfg := range allKeyConfigs(*def) {
				log.Debug(ctx, "Checking key", "name", cfg.Name, "type", cfg.Type)

				// Check if key already defined
				if cfg.Type == key.EOA {
					_, ok := eoa.Address(network, eoa.Role(cfg.Name))
					if ok {
						log.Info(ctx, "Skipping eoa key already defined", "name", cfg.Name)
						continue
					}
				} else {
					_, ok := def.Manifest.Keys[cfg.Name][cfg.Type]
					if ok {
						log.Info(ctx, "Skipping secret key already in manifest", "name", cfg.Name, "type", cfg.Type)
						continue
					}
				}

				if err := verifyKeyNodeType(*def, cfg); err != nil {
					return errors.Wrap(err, "verify cfg", "name", cfg.Name, "type", cfg.Type)
				}

				privkey, err := key.UploadNew(ctx, cfg)
				if err != nil {
					return errors.Wrap(err, "upload key", "name", cfg.Name, "type", cfg.Type)
				}

				addr, err := privkey.Addr()
				if err != nil {
					return err
				}

				if cfg.Type != key.EOA {
					if _, ok := manifest.Keys[cfg.Name]; !ok {
						manifest.Keys[cfg.Name] = make(map[key.Type]string)
					}
					manifest.Keys[cfg.Name][cfg.Type] = addr
				}
			}

			if err := printKeysToml(manifest.Keys, cmd.OutOrStdout()); err != nil {
				return errors.Wrap(err, "write keys toml")
			}

			return nil
		},
	}

	return cmd
}

// verifyKeyNodeType checks if the node exists in the manifest and if the key type is allowed for the node.
func verifyKeyNodeType(def app.Definition, cfg key.UploadConfig) error {
	if err := cfg.Type.Verify(); err != nil {
		return err
	}

	if cfg.Type != key.EOA {
		// Non-EOA keys must be not already defined in the manifest.
		_, ok := def.Manifest.Keys[cfg.Name][cfg.Type]
		if ok {
			return errors.New("key already exists in manifest")
		}
	}

	if cfg.Type == key.EOA {
		eoaRole := eoa.Role(cfg.Name)
		if err := eoaRole.Verify(); err != nil {
			return errors.Wrap(err, "verifying name as eoa type")
		}

		account, ok := eoa.AccountForRole(def.Testnet.Network, eoaRole)
		if !ok {
			return errors.New("eoa account not found")
		}

		if account.Type != eoa.TypeSecret {
			return errors.New("cannot create eoa key for non secret account")
		}

		if account.Address != (common.Address{}) {
			return errors.New("cannot create eoa key already defined", "addr", account.Address.Hex())
		}

		return nil
	}

	for _, node := range def.Testnet.Nodes {
		if node.Name == cfg.Name {
			if cfg.Type == key.P2PExecution {
				return errors.New("cannot create execution key for halo node")
			}

			return nil
		}
	}

	for _, evm := range def.Testnet.OmniEVMs {
		if evm.InstanceName == cfg.Name {
			if cfg.Type != key.P2PExecution {
				return errors.New("only execution keys allowed for evm nodes")
			}

			return nil
		}
	}

	return errors.New("node not found")
}

// allKeyConfigs returns all key configurations for a given manifest.
func allKeyConfigs(def app.Definition) []key.UploadConfig {
	var cfgs []key.UploadConfig

	// We need a relayer and monitor EOA key
	cfgs = append(cfgs,
		key.UploadConfig{
			Network: def.Testnet.Network,
			Name:    string(eoa.RoleRelayer),
			Type:    key.EOA,
		}, key.UploadConfig{
			Network: def.Testnet.Network,
			Name:    string(eoa.RoleMonitor),
			Type:    key.EOA,
		},
	)

	// All evm's need execution P2P keys
	for _, evm := range def.Testnet.OmniEVMs {
		cfgs = append(cfgs, key.UploadConfig{
			Network: def.Testnet.Network,
			Name:    evm.InstanceName,
			Type:    key.P2PExecution,
		})
	}

	for _, node := range def.Testnet.Nodes {
		// All halo's need consensus P2P keys
		cfgs = append(cfgs, key.UploadConfig{
			Network: def.Testnet.Network,
			Name:    node.Name,
			Type:    key.P2PConsensus,
		})

		// Validators need validator keys
		if node.Mode == types.ModeValidator {
			cfgs = append(cfgs, key.UploadConfig{
				Network: def.Testnet.Network,
				Name:    node.Name,
				Type:    key.Validator,
			})
		}
	}

	return cfgs
}

// printKeysToml prints the keys to as toml.
func printKeysToml(keys map[string]map[key.Type]string, w io.Writer) error {
	keysOnly := struct {
		Keys map[string]map[key.Type]string `toml:"keys"`
	}{
		Keys: keys,
	}
	var buf bytes.Buffer

	tomlSettings := toml.Config{
		NormFieldName: func(_ reflect.Type, key string) string {
			return key
		},
		FieldToKey: func(_ reflect.Type, field string) string {
			return field
		},
		MissingField: func(rt reflect.Type, field string) error {
			return errors.New("field not defined", "field", field, "type", rt.String())
		},
	}
	err := tomlSettings.NewEncoder(&buf).Encode(keysOnly)
	if err != nil {
		return errors.Wrap(err, "encode manifest")
	}

	_, _ = fmt.Fprint(w, buf.String())

	return nil
}
