package commands

import (
	"os"
	"path/filepath"
	"strconv"
	"testing"

	"github.com/cometbft/cometbft/types"
	config2 "github.com/omni-network/omni/pkg/config"
	"gotest.tools/assert"
)

func TestInitKeysFilesConfig(t *testing.T) {
	t.Run("init file", func(t *testing.T) {
		// create a temp directory and throwaway after the test
		tmpDir := makeTempDIr(t)
		defer removeTempDir(t, tmpDir)

		// prepare the command
		rootCmd, err := NewRootCommand()
		if err != nil {
			t.Fatal(err)
		}
		rootCmd.AddCommand(InitKeyFileConfigCmd())

		// execute the command
		rootCmd.SetArgs([]string{"init", "--home-dir", tmpDir})
		err = rootCmd.Execute()
		if err != nil {
			t.Fatal(err)
		}

		// check if the files are created
		pvFileName := filepath.Join(config.cometBFTConfig.RootDir, config.cometBFTConfig.PrivValidatorKey)
		gFileName := filepath.Join(config.cometBFTConfig.RootDir, config.cometBFTConfig.Genesis)
		nkFileName := filepath.Join(config.cometBFTConfig.RootDir, config.cometBFTConfig.NodeKey)

		assert.Equal(t, true, isFilePresent(pvFileName), "private validator file is not present")
		assert.Equal(t, true, isFilePresent(gFileName), "genesis  file is not present")
		assert.Equal(t, true, isFilePresent(nkFileName), "node key  file is not present")

		// check if the chain id in the genesis file is correct
		genDoc, err := types.GenesisDocFromFile(gFileName)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, true, genDoc.ChainID == strconv.Itoa(int(config.omniConfig.ChainConfig.NetworkID)), "invalid chain id")
	})

	t.Run("mainnet config", func(t *testing.T) {
		// create a temp directory and throwaway after the test
		tmpDir := makeTempDIr(t)
		defer removeTempDir(t, tmpDir)

		// prepare the command
		rootCmd, err := NewRootCommand()
		if err != nil {
			t.Fatal(err)
		}
		rootCmd.AddCommand(InitKeyFileConfigCmd())

		// execute the command
		rootCmd.SetArgs([]string{"init", "--chain-id", "mainnet", "--home-dir", tmpDir})
		err = rootCmd.Execute()
		if err != nil {
			t.Fatal(err)
		}

		// check if the chain id in the genesis file is correct
		gFileName := filepath.Join(config.cometBFTConfig.RootDir, config.cometBFTConfig.Genesis)
		genDoc, err := types.GenesisDocFromFile(gFileName)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, true, genDoc.ChainID == strconv.Itoa(int(config2.Mainnet.NetworkID)), "invalid mainnet chain id")
	})

	t.Run("localnet config", func(t *testing.T) {
		// create a temp directory and throwaway after the test
		tmpDir := makeTempDIr(t)
		defer removeTempDir(t, tmpDir)

		// prepare the command
		rootCmd, err := NewRootCommand()
		if err != nil {
			t.Fatal(err)
		}
		rootCmd.AddCommand(InitKeyFileConfigCmd())

		// execute the command
		rootCmd.SetArgs([]string{"init", "--chain-id", "99999", "--portal-addr", "0xabcdefghijklmnop", "--home-dir", tmpDir})
		err = rootCmd.Execute()
		if err != nil {
			t.Fatal(err)
		}

		// check if the chain id in the genesis file is correct
		gFileName := filepath.Join(config.cometBFTConfig.RootDir, config.cometBFTConfig.Genesis)
		genDoc, err := types.GenesisDocFromFile(gFileName)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, true, genDoc.ChainID == "99999", "invalid chain id")
		assert.Equal(t, true, config.omniConfig.ChainConfig.OmniPortalAddress == "0xabcdefghijklmnop", "invalid portal address")
	})
}

func isFilePresent(fileName string) bool {
	fi, err := os.Stat(fileName)
	if err != nil {
		return false
	}
	if fi.IsDir() {
		return false
	}
	return true
}
