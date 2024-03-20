package app

import (
	"context"
	"fmt"
	"net"
	"path/filepath"
	"strings"
	"time"

	"github.com/omni-network/omni/e2e/docker"
	"github.com/omni-network/omni/e2e/netman"
	"github.com/omni-network/omni/e2e/types"
	"github.com/omni-network/omni/e2e/vmcompose"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/netconf"

	k1 "github.com/cometbft/cometbft/crypto/secp256k1"
	e2e "github.com/cometbft/cometbft/test/e2e/pkg"
	"github.com/cometbft/cometbft/test/e2e/pkg/exec"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/p2p/enode"

	"github.com/BurntSushi/toml"
)

// DefinitionConfig is the configuration required to create a full Definition.
type DefinitionConfig struct {
	ManifestFile  string
	InfraProvider string

	// Secrets (not required for devnet)
	DeployKeyFile  string
	RelayerKeyFile string
	RPCOverrides   map[string]string

	InfraDataFile string // Not required for docker provider
	OmniImgTag    string // OmniImgTag is the docker image tag used for halo and relayer.
}

// DefaultDefinitionConfig returns a default configuration for a Definition.
func DefaultDefinitionConfig() DefinitionConfig {
	defaultTag := "main"
	if out, err := exec.CommandOutput(context.Background(), "git", "rev-parse", "--short", "HEAD"); err == nil {
		defaultTag = strings.TrimSpace(string(out))
	}

	return DefinitionConfig{
		InfraProvider: docker.ProviderName,
		OmniImgTag:    defaultTag,
	}
}

// Definition defines a e2e network. All (sub)commands of the e2e cli requires a definition operate.
// Armed with a definition, a e2e network can be deployed, started, tested, stopped, etc.
type Definition struct {
	Manifest types.Manifest
	Testnet  types.Testnet // Note that testnet is the cometBFT term.
	Infra    types.InfraProvider
	Netman   netman.Manager
	Backends ethbackend.Backends
}

func MakeDefinition(cfg DefinitionConfig) (Definition, error) {
	manifest, err := LoadManifest(cfg.ManifestFile)
	if err != nil {
		return Definition{}, errors.Wrap(err, "loading manifest")
	}

	var infd types.InfrastructureData
	switch cfg.InfraProvider {
	case docker.ProviderName:
		infd, err = docker.NewInfraData(manifest)
	case vmcompose.ProviderName:
		infd, err = vmcompose.LoadData(cfg.InfraDataFile)
	default:
		return Definition{}, errors.New("unknown infra provider", "provider", cfg.InfraProvider)
	}
	if err != nil {
		return Definition{}, errors.Wrap(err, "loading infrastructure data")
	}

	testnet, err := TestnetFromManifest(manifest, infd, cfg)
	if err != nil {
		return Definition{}, errors.Wrap(err, "loading testnet")
	}

	var backends ethbackend.Backends
	if !testnet.OnlyMonitor {
		backends, err = ethbackend.NewBackends(testnet, cfg.DeployKeyFile)
		if err != nil {
			return Definition{}, errors.Wrap(err, "new backends")
		}
	}

	netman, err := netman.NewManager(testnet, backends, cfg.RelayerKeyFile)
	if err != nil {
		return Definition{}, errors.Wrap(err, "get network")
	}

	var infp types.InfraProvider
	switch cfg.InfraProvider {
	case docker.ProviderName:
		infp = docker.NewProvider(testnet, infd, cfg.OmniImgTag)
	case vmcompose.ProviderName:
		infp = vmcompose.NewProvider(testnet, infd, cfg.OmniImgTag)
	default:
		return Definition{}, errors.New("unknown infra provider", "provider", cfg.InfraProvider)
	}

	return Definition{
		Manifest: manifest,
		Testnet:  testnet,
		Infra:    infp,
		Backends: backends,
		Netman:   netman,
	}, nil
}

func adaptCometTestnet(testnet *e2e.Testnet, imgTag string) *e2e.Testnet {
	testnet.Dir = runsDir(testnet.File)
	testnet.VoteExtensionsEnableHeight = 1
	testnet.UpgradeVersion = "omniops/halo:" + imgTag
	for i := range testnet.Nodes {
		testnet.Nodes[i] = adaptNode(testnet.Nodes[i], imgTag)
	}

	return testnet
}

func adaptNode(node *e2e.Node, tag string) *e2e.Node {
	node.Version = "omniops/halo:" + tag
	node.PrivvalKey = k1.GenPrivKey()

	return node
}

// runsDir returns the runs directory for a given manifest file.
// E.g. /path/to/manifests/manifest.toml > /path/to/runs/manifest.
func runsDir(manifestFile string) string {
	resp := strings.TrimSuffix(manifestFile, filepath.Ext(manifestFile))
	return strings.Replace(resp, "manifests", "runs", 1)
}

// LoadManifest loads a manifest from disk.
func LoadManifest(path string) (types.Manifest, error) {
	manifest := types.Manifest{}
	_, err := toml.DecodeFile(path, &manifest)
	if err != nil {
		return manifest, errors.Wrap(err, "decode manifest")
	}

	return manifest, nil
}

func MonitorOnlyTestnet(manifest types.Manifest, infd types.InfrastructureData, cfg DefinitionConfig) (types.Testnet, error) {
	publics, err := publicChains(manifest, cfg)
	if err != nil {
		return types.Testnet{}, err
	}

	cmtTestnet, err := monitorCometTestnet(manifest.Manifest, cfg.ManifestFile, infd.InfrastructureData)
	if err != nil {
		return types.Testnet{}, errors.Wrap(err, "testnet from manifest")
	}

	return types.Testnet{
		Network:      manifest.Network,
		Testnet:      cmtTestnet,
		PublicChains: publics,
		OnlyMonitor:  true,
	}, nil
}

// monitorCometTestnet returns a bare minimum instance of *e2e.Testnet. It doesn't have any nodes or chain details setup.
func monitorCometTestnet(manifest e2e.Manifest, file string, ifd e2e.InfrastructureData) (*e2e.Testnet, error) {
	dir := strings.TrimSuffix(file, filepath.Ext(file))

	_, ipNet, err := net.ParseCIDR(ifd.Network)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("invalid IP network address %q", ifd.Network))
	}

	testnet := &e2e.Testnet{
		Name:         filepath.Base(dir),
		File:         file,
		Dir:          runsDir(file),
		IP:           ipNet,
		InitialState: manifest.InitialState,
		Prometheus:   manifest.Prometheus,
	}

	return testnet, nil
}

//nolint:nosprintfhostport // Not an issue for non-critical e2e test code.
func TestnetFromManifest(manifest types.Manifest, infd types.InfrastructureData, cfg DefinitionConfig) (types.Testnet, error) {
	if manifest.OnlyMonitor {
		// Create a bare minimum comet testnet only with test di, prometheus and ipnet.
		// Otherwise e2e.NewTestnetFromManifest panics because there are no nodes set
		// in the only_monitor manifest.
		return MonitorOnlyTestnet(manifest, infd, cfg)
	}

	cmtTestnet, err := e2e.NewTestnetFromManifest(manifest.Manifest, cfg.ManifestFile, infd.InfrastructureData)
	if err != nil {
		return types.Testnet{}, errors.Wrap(err, "testnet from manifest")
	}

	var omniEVMS []types.OmniEVM
	for name, gcmode := range manifest.OmniEVMs() {
		inst, ok := infd.Instances[name]
		if !ok {
			return types.Testnet{}, errors.New("omni evm instance not found in infrastructure data")
		}

		nodeKey, err := crypto.GenerateKey()
		if err != nil {
			return types.Testnet{}, errors.Wrap(err, "generate node key")
		}

		en := enode.NewV4(&nodeKey.PublicKey, inst.IPAddress, 30303, 30303)

		internalIP := inst.IPAddress.String()
		if infd.Provider == docker.ProviderName {
			internalIP = name // For docker, we use container names
		}

		omniEVMS = append(omniEVMS, types.OmniEVM{
			Chain:           types.OmniEVMByNetwork(manifest.Network),
			InstanceName:    name,
			InternalIP:      inst.IPAddress,
			ExternalIP:      inst.ExtIPAddress,
			ProxyPort:       inst.Port,
			InternalRPC:     fmt.Sprintf("http://%s:8545", internalIP),
			InternalAuthRPC: fmt.Sprintf("http://%s:8551", internalIP),
			ExternalRPC:     fmt.Sprintf("http://%s:%d", inst.ExtIPAddress.String(), inst.Port),
			NodeKey:         nodeKey,
			Enode:           en,
			GcMode:          gcmode,
		})
	}

	// Second pass to mesh the bootnodes
	for i := range omniEVMS {
		var bootnodes []*enode.Node
		for j, bootEVM := range omniEVMS {
			if i == j {
				continue // Skip self
			}
			bootnodes = append(bootnodes, bootEVM.Enode)
		}
		omniEVMS[i].BootNodes = bootnodes
	}

	var anvils []types.AnvilChain
	for _, chain := range types.AnvilChainsByNames(manifest.AnvilChains) {
		inst, ok := infd.Instances[chain.Name]
		if !ok {
			return types.Testnet{}, errors.New("anvil chain instance not found in infrastructure data")
		}

		chain.IsAVSTarget = chain.Name == manifest.AVSTarget

		internalIP := inst.IPAddress.String()
		if infd.Provider == docker.ProviderName {
			internalIP = chain.Name // For docker, we use container names
		}
		if infd.Provider == vmcompose.ProviderName {
			chain.BlockPeriod = time.Second * 12 // Slow block times for anvils on long-lived VMs to reduce disk usage.
		}

		anvils = append(anvils, types.AnvilChain{
			Chain:       chain,
			InternalIP:  inst.IPAddress,
			ProxyPort:   inst.Port,
			LoadState:   "./anvil/state.json",
			InternalRPC: fmt.Sprintf("http://%s:8545", internalIP),
			ExternalRPC: fmt.Sprintf("http://%s:%d", inst.ExtIPAddress.String(), inst.Port),
		})
	}

	publics, err := publicChains(manifest, cfg)
	if err != nil {
		return types.Testnet{}, err
	}

	return types.Testnet{
		Network:      manifest.Network,
		Testnet:      adaptCometTestnet(cmtTestnet, cfg.OmniImgTag),
		OmniEVMs:     omniEVMS,
		AnvilChains:  anvils,
		PublicChains: publics,
	}, nil
}

func publicChains(manifest types.Manifest, cfg DefinitionConfig) ([]types.PublicChain, error) {
	var publics []types.PublicChain
	for _, name := range manifest.PublicChains {
		chain, err := types.PublicChainByName(name)
		if err != nil {
			return nil, errors.Wrap(err, "get public chain")
		}

		chain.IsAVSTarget = chain.Name == manifest.AVSTarget

		addr, ok := cfg.RPCOverrides[name]
		if !ok {
			addr = types.PublicRPCByName(name)
		}

		publics = append(publics, types.PublicChain{
			Chain:      chain,
			RPCAddress: addr,
		})
	}

	return publics, nil
}

// internalNetwork returns a internal intra-network netconf.Network from the testnet and deployInfo.
func internalNetwork(testnet types.Testnet, deployInfo map[types.EVMChain]netman.DeployInfo, evmPrefix string) netconf.Network {
	var chains []netconf.Chain

	// in monitor only mode there are no anvil chains and omni EVMs
	if !testnet.OnlyMonitor {
		omniEVM := omniEVMByPrefix(testnet, evmPrefix)
		chains = append(chains, netconf.Chain{
			ID:                omniEVM.Chain.ID,
			Name:              omniEVM.Chain.Name,
			RPCURL:            omniEVM.InternalRPC,
			AuthRPCURL:        omniEVM.InternalAuthRPC,
			PortalAddress:     deployInfo[omniEVM.Chain].PortalAddress.Hex(),
			DeployHeight:      deployInfo[omniEVM.Chain].DeployHeight,
			BlockPeriod:       omniEVM.Chain.BlockPeriod,
			FinalizationStrat: omniEVM.Chain.FinalizationStrat,
			IsOmni:            true,
		})

		// Add all anvil chains
		for _, anvil := range testnet.AnvilChains {
			chains = append(chains, netconf.Chain{
				ID:                anvil.Chain.ID,
				Name:              anvil.Chain.Name,
				RPCURL:            anvil.InternalRPC,
				PortalAddress:     deployInfo[anvil.Chain].PortalAddress.Hex(),
				DeployHeight:      deployInfo[anvil.Chain].DeployHeight,
				BlockPeriod:       anvil.Chain.BlockPeriod,
				FinalizationStrat: anvil.Chain.FinalizationStrat,
				IsEthereum:        anvil.Chain.IsAVSTarget,
			})
		}
	}

	// Add all public chains
	for _, public := range testnet.PublicChains {
		pc := netconf.Chain{
			ID:                public.Chain.ID,
			Name:              public.Chain.Name,
			RPCURL:            public.RPCAddress,
			PortalAddress:     deployInfo[public.Chain].PortalAddress.Hex(),
			DeployHeight:      deployInfo[public.Chain].DeployHeight,
			BlockPeriod:       public.Chain.BlockPeriod,
			FinalizationStrat: public.Chain.FinalizationStrat,
			IsEthereum:        public.Chain.IsAVSTarget,
			AVSContractAddr:   public.Chain.AVSContractAddress,
		}

		chains = append(chains, pc)
	}

	return netconf.Network{
		Name:   testnet.Network,
		Chains: chains,
	}
}

// externalNetwork returns a external e2e-app netconf.Network from the testnet and deployInfo.
func externalNetwork(testnet types.Testnet, deployInfo map[types.EVMChain]netman.DeployInfo) netconf.Network {
	var chains []netconf.Chain

	// Connect to a random omni evm
	omniEVM := random[types.OmniEVM](testnet.OmniEVMs)
	chains = append(chains, netconf.Chain{
		ID:                omniEVM.Chain.ID,
		Name:              omniEVM.Chain.Name,
		RPCURL:            omniEVM.ExternalRPC,
		PortalAddress:     deployInfo[omniEVM.Chain].PortalAddress.Hex(),
		DeployHeight:      deployInfo[omniEVM.Chain].DeployHeight,
		BlockPeriod:       omniEVM.Chain.BlockPeriod,
		FinalizationStrat: omniEVM.Chain.FinalizationStrat,
		IsOmni:            true,
	})

	// Add all anvil chains
	for _, anvil := range testnet.AnvilChains {
		chains = append(chains, netconf.Chain{
			ID:                anvil.Chain.ID,
			Name:              anvil.Chain.Name,
			RPCURL:            anvil.ExternalRPC,
			PortalAddress:     deployInfo[anvil.Chain].PortalAddress.Hex(),
			DeployHeight:      deployInfo[anvil.Chain].DeployHeight,
			BlockPeriod:       anvil.Chain.BlockPeriod,
			FinalizationStrat: anvil.Chain.FinalizationStrat,
			IsEthereum:        anvil.Chain.IsAVSTarget,
		})
	}

	// Add all public chains
	for _, public := range testnet.PublicChains {
		chains = append(chains, netconf.Chain{
			ID:                public.Chain.ID,
			Name:              public.Chain.Name,
			RPCURL:            public.RPCAddress,
			PortalAddress:     deployInfo[public.Chain].PortalAddress.Hex(),
			DeployHeight:      deployInfo[public.Chain].DeployHeight,
			BlockPeriod:       public.Chain.BlockPeriod,
			FinalizationStrat: public.Chain.FinalizationStrat,
			IsEthereum:        public.Chain.IsAVSTarget,
		})
	}

	for _, chain := range chains {
		if err := chain.FinalizationStrat.Verify(); err != nil {
			panic(err) // Ok to panic since this e2e and is build-time issue.
		}
	}

	return netconf.Network{
		Name:   testnet.Network,
		Chains: chains,
	}
}

// omniEVMByPrefix returns a omniEVM from the testnet with the given prefix.
// Or a random omniEVM if prefix is empty.
// Or the only omniEVM if there is only one.
func omniEVMByPrefix(testnet types.Testnet, prefix string) types.OmniEVM {
	if prefix == "" {
		return random(testnet.OmniEVMs)
	} else if len(testnet.OmniEVMs) == 1 {
		return testnet.OmniEVMs[0]
	}

	for _, evm := range testnet.OmniEVMs {
		if strings.HasPrefix(evm.InstanceName, prefix) {
			return evm
		}
	}

	panic("evm not found")
}

// random returns a random item from a slice.
func random[T any](items []T) T {
	return items[int(time.Now().UnixNano())%len(items)]
}
