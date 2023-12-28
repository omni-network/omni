// Copyright 2023 The Omni Authors. All rights reserved.
// This file is part of the omni library.
//
// The omni library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The omni library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.

// You should have received a copy of the GNU Lesser General Public License
// along with the omni library. If not, see <http://www.gnu.org/licenses/>.

package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/0x6flab/namegenerator"
)

const (
	MainnetChainId       = 8880         // TODO: Need to check if this id is available
	MainnetPortalAddress = "0x1234...." // TODO: Please fill this after deployment of the contract
	TestnetChainId       = 8980         // TODO: Need to check if this id is available
	TestnetPortalAddress = "0x5678..."  // TODO: Please fill this after deployment of the contract

	DefaultOmniRootDir    = ".omni"  // default omni root directory
	DefaultOmniConfigDir  = "config" // default omni config directory
	DefaultOmniDataDir    = "data"   // default omni data directory
	DefaultOmniLogDir     = "log"    // default omni log directory
	DefaultConfigFileName = "config" // default omni configuration file (without .yaml ext)
	DefaultLogLevel       = "info"   // default log level for omni node
	DefaultLogWriter      = "stdout" // log writer (console, file)
	DefaultDirPerm        = 0o700    // default directory permission for all the above directories

	NativeTokenSymbol   = "OMNI" // the native tracker symbol of omni protocol
	LayerOneTokenSymbol = "ETH"  // the native tracker symbol of ethereum layer 1 protocol
)

type OmniConfig struct {
	NodeName           string          `yaml:"nodename"`       // a human-readable name for this node
	LogLevel           string          `yaml:"loglevel"`       // the log level for logging events
	LogWriter          string          `yaml:"logwriter"`      // where to write the log
	HomeDirectory      string          `yaml:"homedir"`        // the home directory which has omni and cometBFT root directory
	OmniRootDir        string          `yaml:"rootdir"`        // omni specific root directory
	OmniConfigDir      string          `yaml:"configdir"`      // omni specific config directory
	OmniDataDir        string          `yaml:"datadir"`        // omni specific data directory
	OmniLogDir         string          `yaml:"logdir"`         // omni node's log directory
	OmniConfigFileName string          `yaml:"configfilename"` // config file name of omni
	ChainConfig        OmniChainConfig `yaml:"chainconfig"`
}

type OmniChainConfig struct {
	NetworkID          uint64                  `yaml:"networkid"`       // the network id of the omni chain
	OmniPortalAddress  string                  `yaml:"portaladdress"`   // the portal address of the omni chain
	RollupChainsConfig []SupportedChainsConfig `yaml:"supportedchains"` // configurations related to rollup chains
}

type SupportedChainsConfig struct {
	ChainName     string `yaml:"name"`          // name of the rollup chain
	ChainID       uint64 `yaml:"id"`            // chain id of the rollup chain
	TokenSymbol   string `yaml:"tokensymbol"`   // the token tracker of the chain
	PortalAddress string `yaml:"portaladdress"` // the address of the portal contract deployed in the chain
	RPCUrl        string `yaml:"rpcurl"`        // the rpc url for this chain
}

var (
	// Mainnet default values set for the omni chain
	Mainnet = OmniChainConfig{
		NetworkID:         MainnetChainId,
		OmniPortalAddress: MainnetPortalAddress,

		// for every supported rollup chain, add its configuration
		RollupChainsConfig: []SupportedChainsConfig{
			{
				ChainName:     "Arbitrum",
				ChainID:       1111,                     // TODO: set proper chain id
				TokenSymbol:   "ARB",                    // TODO: set proper symbol name
				PortalAddress: "0x1234...",              // TODO: set proper portal address
				RPCUrl:        "https://aaaaa.com:8080", // TODO: set proper rpc url
			},
			{
				ChainName:     "Optimism",
				ChainID:       1111,                    // TODO: set proper chain id
				TokenSymbol:   "OPT",                   // TODO: set proper symbol name
				PortalAddress: "0x1234...",             // TODO: set proper portal address
				RPCUrl:        "https://bbbb.com:8080", // TODO: set proper rpc url
			},
		},
	}

	// Testnet default values set for the omni chain
	Testnet = OmniChainConfig{
		NetworkID:         TestnetChainId,
		OmniPortalAddress: TestnetPortalAddress,

		// for every supported rollup chain, add its configuration
		RollupChainsConfig: []SupportedChainsConfig{
			{
				ChainName:     "Arbitrum",
				ChainID:       1111,                     // TODO: set proper chain id
				TokenSymbol:   "ARB",                    // TODO: set proper symbol name
				PortalAddress: "0x1234...",              // TODO: set proper portal address
				RPCUrl:        "https://aaaaa.com:8080", // TODO: set proper rpc url
			},
			{
				ChainName:     "Optimism",
				ChainID:       1111,                    // TODO: set proper chain id
				TokenSymbol:   "OPT",                   // TODO: set proper symbol name
				PortalAddress: "0x1234...",             // TODO: set proper portal address
				RPCUrl:        "https://bbbb.com:8080", // TODO: set proper rpc url
			},
		},
	}
)

// LoadDefaultConfig constructs the default configuration for the omni chain
func LoadDefaultConfig() *OmniConfig {
	generator := namegenerator.NewNameGenerator()
	randomNodeName := generator.Generate()

	homeDir, err := os.UserHomeDir()
	if err != nil {
		homeDir = "" // TODO: not passing error up. can we safely assume we will have a home dir????
	}

	omniRootDir := filepath.Join(homeDir, DefaultOmniRootDir)
	omniConfigDir := filepath.Join(omniRootDir, DefaultOmniConfigDir)
	omniDataDir := filepath.Join(omniRootDir, DefaultOmniDataDir)
	omniLogDir := filepath.Join(omniRootDir, DefaultOmniLogDir)

	return &OmniConfig{
		NodeName:           randomNodeName,
		LogLevel:           DefaultLogLevel,
		LogWriter:          DefaultLogWriter,
		HomeDirectory:      homeDir,
		OmniRootDir:        omniRootDir,
		OmniConfigDir:      omniConfigDir,
		OmniDataDir:        omniDataDir,
		OmniLogDir:         omniLogDir,
		OmniConfigFileName: DefaultConfigFileName,
		ChainConfig:        Mainnet,
	}
}

// EnsureDirectories make sure the required directories are present before we start
// using them
func EnsureDirectories(ocfg *OmniConfig) error {
	// ensure that the root directory is present, otherwise create it
	err := os.MkdirAll(ocfg.OmniRootDir, DefaultDirPerm)
	if err != nil {
		return fmt.Errorf("could not create directory %q: %w", ocfg.OmniRootDir, err)
	}

	// ensure that the omni config directory is present, otherwise create it
	err = os.MkdirAll(ocfg.OmniConfigDir, DefaultDirPerm)
	if err != nil {
		return fmt.Errorf("could not create directory %q: %w", ocfg.OmniConfigDir, err)
	}

	// ensure that the omni data directory is present, otherwise create it
	err = os.MkdirAll(ocfg.OmniDataDir, DefaultDirPerm)
	if err != nil {
		return fmt.Errorf("could not create directory %q: %w", ocfg.OmniDataDir, err)
	}

	// ensure that the omni log directory is present, otherwise create it
	err = os.MkdirAll(ocfg.OmniLogDir, DefaultDirPerm)
	if err != nil {
		return fmt.Errorf("could not create directory %q: %w", ocfg.OmniLogDir, err)
	}

	return nil
}
