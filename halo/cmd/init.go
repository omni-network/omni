package cmd

import (
	"context"
	"math/rand/v2"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/omni-network/omni/halo/attest/voter"
	halocfg "github.com/omni-network/omni/halo/config"
	"github.com/omni-network/omni/halo/genutil"
	libcmd "github.com/omni-network/omni/lib/cmd"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"

	cmtconfig "github.com/cometbft/cometbft/config"
	k1 "github.com/cometbft/cometbft/crypto/secp256k1"
	cmtos "github.com/cometbft/cometbft/libs/os"
	"github.com/cometbft/cometbft/p2p"
	"github.com/cometbft/cometbft/p2p/pex"
	"github.com/cometbft/cometbft/privval"
	rpchttp "github.com/cometbft/cometbft/rpc/client/http"

	"github.com/ethereum/go-ethereum/common"

	"github.com/spf13/cobra"
)

const maxPeers = 5 // Limit the amount of peers to add to the address book.

// InitConfig is the config for the init command.
type InitConfig struct {
	HomeDir       string
	Moniker       string
	Network       netconf.ID
	TrustedSync   bool
	AddrBook      bool
	HaloCfgFunc   func(*halocfg.Config)
	CometCfgFunc  func(*cmtconfig.Config)
	Force         bool
	Clean         bool
	ExecutionHash common.Hash
}

func (c InitConfig) Verify() error {
	return c.Network.Verify()
}

func (c InitConfig) HaloCfg(cfg *halocfg.Config) {
	if c.HaloCfgFunc != nil {
		c.HaloCfgFunc(cfg)
	}
}

func (c InitConfig) CometCfg(cfg *cmtconfig.Config) {
	if c.CometCfgFunc != nil {
		c.CometCfgFunc(cfg)
	}
}

// newInitCmd returns a new cobra command that initializes the files and folders required by halo.
func newInitCmd() *cobra.Command {
	// Default config flags
	cfg := InitConfig{
		HomeDir:      halocfg.DefaultHomeDir,
		Force:        false,
		HaloCfgFunc:  func(*halocfg.Config) {},
		CometCfgFunc: func(*cmtconfig.Config) {},
	}

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
  │   ├── snapshots                  # Snapshot directory
  │   ├── priv_validator_state.json  # CometBFT private validator state (slashing protection)
  │   └── voter_state.json           # Cross chain voter state (slashing protection)

Existing files are not overwritten, unless --clean is specified.
The home directory should only contain subdirectories, no files, use --force to ignore this check.
`,
		RunE: func(cmd *cobra.Command, _ []string) error {
			ctx := cmd.Context()
			if err := libcmd.LogFlags(ctx, cmd.Flags()); err != nil {
				return err
			}

			if err := cfg.Verify(); err != nil {
				return errors.Wrap(err, "verify flags")
			}

			return InitFiles(cmd.Context(), cfg)
		},
	}

	bindInitFlags(cmd.Flags(), &cfg)

	return cmd
}

// InitFiles initializes the files and folders required by halo.
// It ensures a network and genesis file is generated/downloaded for the provided network.
//
//nolint:gocognit,nestif,gocyclo,maintidx // This is just many sequential steps.
func InitFiles(ctx context.Context, initCfg InitConfig) error {
	if initCfg.Network == "" {
		return errors.New("required flag --network empty")
	}

	log.Info(ctx, "Initializing halo files and directories", "home", initCfg.HomeDir, "network", initCfg.Network)
	homeDir := initCfg.HomeDir
	network := initCfg.Network

	// Quick sanity check if --home contains files (it should only contain dirs).
	// This prevents accidental initialization in wrong current dir.
	if !initCfg.Force {
		files, _ := os.ReadDir(homeDir) // Ignore error, we'll just assume it's empty.
		for _, file := range files {
			if file.IsDir() { // Ignore directories
				continue
			}

			return errors.New("home directory contains unexpected file(s), use --force to initialize anyway",
				"home", homeDir, "example_file", file.Name())
		}
	}

	if initCfg.Clean {
		log.Info(ctx, "Deleting home directory, since --clean=true")
		if err := os.RemoveAll(homeDir); err != nil {
			return errors.Wrap(err, "remove home dir")
		}
	}

	// Initialize comet config.
	comet := DefaultCometConfig(homeDir)
	comet.Moniker = initCfg.Moniker
	initCfg.CometCfg(&comet)

	// Initialize halo config.
	cfg := halocfg.DefaultConfig()
	cfg.HomeDir = homeDir
	cfg.Network = network
	initCfg.HaloCfg(&cfg)

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

	// Add P2P seeds to comet config (persisted peers works better than seeds)
	if seeds := network.Static().ConsensusSeeds(); len(seeds) > 0 {
		comet.P2P.PersistentPeers = strings.Join(seeds, ",")
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

	// Connect to RPC server
	rpcServer := network.Static().ConsensusRPC()
	rpcCl, err := rpchttp.New(rpcServer, "/websocket")
	if err != nil {
		return errors.Wrap(err, "create rpc client")
	}

	if initCfg.TrustedSync {
		// Trusted state sync only supported for protected networks.
		height, hash, err := getTrustHeightAndHash(ctx, rpcCl)
		if err != nil {
			return errors.Wrap(err, "get trusted height")
		}

		comet.StateSync.Enable = true
		comet.StateSync.RPCServers = []string{rpcServer, rpcServer} // CometBFT requires two RPC servers. Duplicate our RPC for now.
		comet.StateSync.TrustHeight = height
		comet.StateSync.TrustHash = hash

		log.Info(ctx, "Trusted state-sync enabled", "height", height, "hash", hash, "rpc_endpoint", rpcServer)
	} else {
		log.Info(ctx, "Not initializing trusted state sync")
	}

	addrBookPath := filepath.Join(homeDir, cmtconfig.DefaultConfigDir, cmtconfig.DefaultAddrBookName)
	if initCfg.AddrBook && !cmtos.FileExists(addrBookPath) {
		// Populate address book with random public peers from the connected node.
		// This aids in bootstrapping the P2P network. Seed nodes don't work well on their pwn for some reason.

		peers, err := getPeers(ctx, rpcCl, maxPeers)
		if err != nil {
			return errors.Wrap(err, "get peers", "rpc", rpcServer)
		} else if len(peers) == 0 {
			return errors.New("no routable public peers found", "rpc", rpcServer)
		}

		addrBook := pex.NewAddrBook(addrBookPath, true)
		for _, peer := range peers {
			if err := addrBook.AddAddress(peer, peer); err != nil {
				return errors.Wrap(err, "add address")
			}
		}
		addrBook.Save()

		log.Info(ctx, "Populated comet address book", "path", addrBookPath, "peers", len(peers), "rpc_endpoint", rpcServer)
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
	} else if err := halocfg.WriteConfigTOML(cfg, log.DefaultConfig()); err != nil {
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
		log.Info(ctx, "Found cometBFT private validator",
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

	// Setup genesis file
	genFile := comet.GenesisFile()
	if cmtos.FileExists(genFile) {
		log.Info(ctx, "Found genesis file", "path", genFile)
	} else if network == netconf.Simnet {
		pubKey, err := pv.GetPubKey()
		if err != nil {
			return errors.Wrap(err, "get public key")
		}

		cosmosGen, err := genutil.MakeGenesis(network, time.Now(), initCfg.ExecutionHash, pubKey)
		if err != nil {
			return err
		}

		genDoc, err := cosmosGen.ToGenesisDoc()
		if err != nil {
			return errors.Wrap(err, "convert to genesis doc")
		}

		if err := genDoc.SaveAs(genFile); err != nil {
			return errors.Wrap(err, "save genesis file")
		}
		log.Info(ctx, "Generated simnet genesis file", "path", genFile)
	} else if len(network.Static().ConsensusGenesisJSON) > 0 {
		err := os.WriteFile(genFile, network.Static().ConsensusGenesisJSON, 0o644)
		if err != nil {
			return errors.Wrap(err, "save genesis file")
		}
		log.Info(ctx, "Generated well-known network genesis file", "path", genFile)
	} else {
		return errors.New("network genesis file not supported yet", "network", network)
	}

	// Vote state
	voterStateFile := cfg.VoterStateFile()
	if cmtos.FileExists(voterStateFile) {
		log.Info(ctx, "Found voter state file", "path", voterStateFile)
	} else if err := voter.GenEmptyStateFile(voterStateFile); err != nil {
		return err
	} else {
		log.Info(ctx, "Generated voter state file", "path", voterStateFile)
	}

	return nil
}

// getPeers returns up to max random public peer addresses of the connected node.
func getPeers(ctx context.Context, cl *rpchttp.HTTP, max int) ([]*p2p.NetAddress, error) {
	info, err := cl.NetInfo(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get net info")
	}

	// Shuffle to pick random list of max peers.
	peers := info.Peers
	rand.Shuffle(len(peers), func(i, j int) {
		peers[i], peers[j] = peers[j], peers[i]
	})

	var resp []*p2p.NetAddress
	for _, peer := range peers {
		info := peer.NodeInfo
		addr, err := p2p.NewNetAddressString(p2p.IDAddressString(info.ID(), info.ListenAddr))
		if err != nil {
			return nil, errors.Wrap(err, "parse net address", "addr", p2p.IDAddressString(info.ID(), info.ListenAddr))
		}
		if !addr.Routable() {
			continue // Drop non-routable (private) peers
		}

		log.Info(ctx, "Adding peer to comet address book", "addr", addr.String(), "moniker", peer.NodeInfo.Moniker)

		resp = append(resp, addr)

		if len(resp) >= max {
			break
		}
	}

	return resp, nil
}

func getTrustHeightAndHash(ctx context.Context, cl *rpchttp.HTTP) (int64, string, error) {
	latest, err := cl.Block(ctx, nil)
	if err != nil {
		return 0, "", errors.Wrap(err, "get latest block")
	}

	// Truncate height to last defaultSnapshotPeriod
	const defaultSnapshotPeriod int64 = 1000
	snapshotHeight := defaultSnapshotPeriod * (latest.Block.Height / defaultSnapshotPeriod)

	if snapshotHeight == 0 {
		return 0, "", errors.New("initial snapshot height not reached yet", "latest_height", latest.Block.Height, "target", defaultSnapshotPeriod)
	}

	b, err := cl.Block(ctx, &snapshotHeight)
	if err != nil {
		return 0, "", errors.Wrap(err, "get snapshot block")
	}

	return b.Block.Height, b.BlockID.Hash.String(), nil
}
