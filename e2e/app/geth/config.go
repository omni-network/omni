package geth

import (
	"encoding/hex"
	"encoding/json"
	"os"
	"path/filepath"
	"time"

	"github.com/omni-network/omni/e2e/types"
	"github.com/omni-network/omni/lib/errors"

	"github.com/ethereum/go-ethereum/core"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/eth/downloader"
)

// WriteAllConfig writes all the geth config files for all omniEVMs.
func WriteAllConfig(testnet types.Testnet, genesis core.Genesis) error {
	gethGenesisBz, err := json.MarshalIndent(genesis, "", "  ")
	if err != nil {
		return errors.Wrap(err, "marshal genesis")
	}

	gethConfigFiles := func(evm types.OmniEVM) map[string][]byte {
		return map[string][]byte{
			"genesis.json":   gethGenesisBz,
			"geth/nodekey":   []byte(hex.EncodeToString(ethcrypto.FromECDSA(evm.NodeKey))), // Nodekey is hex encoded
			"geth/jwtsecret": []byte(evm.JWTSecret),
		}
	}

	for _, evm := range testnet.OmniEVMs {
		for file, data := range gethConfigFiles(evm) {
			path := filepath.Join(testnet.Dir, evm.InstanceName, file)
			if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
				return errors.Wrap(err, "mkdir", "path", path)
			}
			if err := os.WriteFile(path, data, 0o644); err != nil {
				return errors.Wrap(err, "write geth config")
			}
		}

		conf := Config{
			Moniker:      evm.InstanceName,
			IsArchive:    evm.IsArchive,
			ChainID:      evm.Chain.ChainID,
			BootNodes:    evm.Peers, // TODO(corver): Use seed nodes once available.
			TrustedNodes: evm.Peers,
		}
		if err := WriteConfigTOML(conf, filepath.Join(testnet.Dir, evm.InstanceName, "config.toml")); err != nil {
			return errors.Wrap(err, "write geth config")
		}
	}

	return nil
}

// WriteConfigTOML writes the geth config to a toml file.
func WriteConfigTOML(conf Config, path string) error {
	bz, err := tomlSettings.Marshal(MakeGethConfig(conf))
	if err != nil {
		return errors.Wrap(err, "marshal toml")
	}

	if err := os.WriteFile(path, bz, 0o644); err != nil {
		return errors.Wrap(err, "write toml")
	}

	return nil
}

// MakeGethConfig returns the full omni geth config for the provided custom config.
func MakeGethConfig(conf Config) FullConfig {
	cfg := defaultGethConfig()
	cfg.Eth.NetworkId = conf.ChainID
	cfg.Node.DataDir = "/geth" // Mount inside docker container
	cfg.Node.IPCPath = "/geth/geth.ipc"

	// Use syncmode=full. Since default "snap" sync has race condition on startup. Where engineAPI newPayload fails
	// if snapsync has not completed. Should probably wait for snapsync to complete before starting engineAPI?
	cfg.Eth.SyncMode = downloader.FullSync

	// Disable pruning for archive nodes.
	// Note that runtime flags are also required for archive nodes, specifically:
	//   --gcmode==archive
	//   --state.scheme=hash
	// This will be deprecated once new state.scheme=path support archive nodes.
	// See https://blog.ethereum.org/2023/09/12/geth-v1-13-0.
	cfg.Eth.NoPruning = conf.IsArchive
	cfg.Eth.Preimages = conf.IsArchive // Geth auto-enables this when NoPruning is set.

	// Ethereum has slow block building times (2~4s), but we need fast times (<1s).
	// Use 500ms so blocks are built in less than 1s.
	cfg.Eth.Miner.Recommit = 500 * time.Millisecond

	// Set the bootnodes and trusted nodes.
	cfg.Node.UserIdent = conf.Moniker
	cfg.Node.P2P.DiscoveryV4 = true // TODO(corver): Switch to v5.
	cfg.Node.P2P.BootstrapNodesV5 = conf.BootNodes
	cfg.Node.P2P.BootstrapNodes = conf.BootNodes
	cfg.Node.P2P.TrustedNodes = conf.TrustedNodes

	// Bind listen addresses to all interfaces inside the container.
	const allInterfaces = "0.0.0.0"
	cfg.Node.AuthAddr = allInterfaces
	cfg.Node.HTTPHost = allInterfaces
	cfg.Node.WSHost = allInterfaces
	cfg.Node.P2P.ListenAddr = allInterfaces + ":30303"

	// Add eth module
	cfg.Node.HTTPModules = append(cfg.Node.HTTPModules, "eth")
	cfg.Node.WSModules = append(cfg.Node.WSModules, "eth")

	if conf.IsArchive {
		cfg.Node.HTTPModules = append(cfg.Node.HTTPModules, "debug")
		cfg.Node.WSModules = append(cfg.Node.WSModules, "debug")
	}

	// Allow all incoming connections.
	cfg.Node.HTTPVirtualHosts = []string{"*"}
	cfg.Node.AuthVirtualHosts = []string{"*"}
	cfg.Node.GraphQLVirtualHosts = []string{"*"}
	cfg.Node.WSOrigins = []string{"*"}
	cfg.Node.HTTPCors = []string{"*"}
	cfg.Node.GraphQLCors = []string{"*"}

	return cfg
}
