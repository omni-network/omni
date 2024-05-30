package app

import (
	"os"
	"path/filepath"
	"testing"

	halocfg "github.com/omni-network/omni/halo/config"
	"github.com/omni-network/omni/lib/k1util"

	cmtconfig "github.com/cometbft/cometbft/config"
	"github.com/cometbft/cometbft/privval"
	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"

	"github.com/ethereum/go-ethereum/crypto"

	"github.com/stretchr/testify/require"
)

func TestLoadPrivVal(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name         string
		cmtPrivval   bool
		cmtPrivState bool
		err          bool
	}{
		{
			name:         "comet privval and state",
			cmtPrivval:   true,
			cmtPrivState: true,
		},
		{
			name:         "no files",
			cmtPrivval:   false,
			cmtPrivState: false,
			err:          true,
		},
		{
			name:         "comet privval only",
			cmtPrivval:   true,
			cmtPrivState: false,
			err:          true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			homeDir := t.TempDir()

			// Define the file paths
			cmtPrivvalFile := filepath.Join(homeDir, "config", "priv_validator_key.json")
			cmtPrivStateFile := filepath.Join(homeDir, "data", "priv_validator_state.json")

			// Ensure the config and data directories exist
			require.NoError(t, os.Mkdir(filepath.Dir(cmtPrivvalFile), 0755))
			require.NoError(t, os.Mkdir(filepath.Dir(cmtPrivStateFile), 0755))

			// Generate the expected private key
			privKey, err := crypto.GenerateKey()
			require.NoError(t, err)

			// Convert the private key to a comet private key
			cmtPrivKey, err := k1util.StdPrivKeyToComet(privKey)
			require.NoError(t, err)

			// Write the comet privval file (with non-zero state)
			key := privval.NewFilePV(cmtPrivKey, cmtPrivvalFile, cmtPrivStateFile)
			err = key.SignVote("chain", &cmtproto.Vote{
				Type: cmtproto.PrecommitType,
			})
			require.NoError(t, err)
			key.Save()

			// Remove the files if they are not needed
			if !tt.cmtPrivval {
				require.NoError(t, os.Remove(cmtPrivvalFile))
			}
			if !tt.cmtPrivState {
				require.NoError(t, os.Remove(cmtPrivStateFile))
			}

			// Setup the config
			cfg := Config{
				Config: halocfg.Config{
					HomeDir: homeDir,
				},
				Comet: cmtconfig.Config{
					BaseConfig: cmtconfig.BaseConfig{
						RootDir:            homeDir,
						PrivValidatorKey:   "config/priv_validator_key.json",
						PrivValidatorState: "data/priv_validator_state.json",
					},
				},
			}

			// Run the test
			pv, err := loadPrivVal(cfg)

			// Assert the results
			if tt.err {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.NotNil(t, pv)
				require.True(t, pv.Key.PrivKey.Equals(cmtPrivKey))
			}
		})
	}
}
