package e2e_test

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"sync"
	"testing"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/e2e/app"
	"github.com/omni-network/omni/e2e/docker"
	"github.com/omni-network/omni/e2e/types"
	"github.com/omni-network/omni/e2e/vmcompose"
	"github.com/omni-network/omni/halo/genutil/evm/predeploys"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/feature"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"

	rpchttp "github.com/cometbft/cometbft/rpc/client/http"
	rpctypes "github.com/cometbft/cometbft/rpc/core/types"
	e2e "github.com/cometbft/cometbft/test/e2e/pkg"
	cmttypes "github.com/cometbft/cometbft/types"

	"github.com/ethereum/go-ethereum/common"

	"github.com/stretchr/testify/require"
)

//nolint:gochecknoglobals // This was copied from cometbft/test/e2e/test/e2e_test.go
var (
	depsCache      = map[string]NetworkDeps{}
	depsCacheMtx   = sync.Mutex{}
	blocksCache    = map[string][]*cmttypes.Block{}
	blocksCacheMtx = sync.Mutex{}
)

// Portal is a struct that contains the chain, client and contract for a portal.
type Portal struct {
	Chain    netconf.Chain
	Client   ethclient.Client
	Contract *bindings.OmniPortal
}

type NetworkDeps struct {
	Testnet      types.Testnet
	Backends     ethbackend.Backends
	Network      netconf.Network
	RPCEndpoints xchain.RPCEndpoints
	SolverAddr   string
}

func (d NetworkDeps) OmniBackend() (*ethbackend.Backend, error) {
	return d.Backends.Backend(d.Network.ID.Static().OmniExecutionChainID)
}

func (d NetworkDeps) L1Backend() (*ethbackend.Backend, error) {
	chainID, ok := d.Network.EthereumChain()
	if !ok {
		return nil, errors.New("no ethereum chain")
	}

	return d.Backends.Backend(chainID.ID)
}

type testFunc struct {
	TestNode    func(*testing.T, netconf.Network, *e2e.Node, []Portal)
	TestPortal  func(*testing.T, netconf.Network, Portal, []Portal)
	TestOmniEVM func(*testing.T, ethclient.Client)
	TestNetwork func(context.Context, *testing.T, NetworkDeps)
	skipFunc    func(types.Manifest) bool
}

func testNode(t *testing.T, fn func(*testing.T, netconf.Network, *e2e.Node, []Portal)) {
	t.Helper()
	test(t, testFunc{TestNode: fn})
}

func testPortal(t *testing.T, fn func(*testing.T, netconf.Network, Portal, []Portal)) {
	t.Helper()
	test(t, testFunc{TestPortal: fn})
}

func testOmniEVM(t *testing.T, fn func(*testing.T, ethclient.Client)) {
	t.Helper()
	test(t, testFunc{TestOmniEVM: fn})
}

func testNetwork(t *testing.T, fn func(context.Context, *testing.T, NetworkDeps)) {
	t.Helper()
	test(t, testFunc{TestNetwork: fn})
}

func maybeTestNetwork(
	t *testing.T,
	skipFunc func(types.Manifest) bool,
	fn func(context.Context, *testing.T, NetworkDeps),
) {
	t.Helper()
	test(t, testFunc{TestNetwork: fn, skipFunc: skipFunc})
}

// test runs tests for testnet nodes. The callback functions are respectively given a
// single node to test, and a single portal to test, running as a subtest in parallel with other subtests.
//
// The testnet manifest must be given as the envvar E2E_MANIFEST. If not set,
// these tests are skipped so that they're not picked up during normal unit
// test runs. If E2E_NODE is also set, only the specified node is tested,
// otherwise all nodes are tested.
func test(t *testing.T, testFunc testFunc) {
	t.Helper()
	ctx := t.Context()

	deps := loadEnv(t)
	testnet := deps.Testnet
	nodes := testnet.Nodes

	if testFunc.skipFunc != nil && testFunc.skipFunc(testnet.Manifest) {
		t.Skip("Skipping test")
		return
	}

	ctx = feature.WithFlags(ctx, testnet.Manifest.FeatureFlags)

	if name := os.Getenv(app.EnvE2ENode); name != "" {
		node := testnet.LookupNode(name)
		require.NotNil(t, node, "node %q not found in testnet %q", name, testnet.Name)
		nodes = []*e2e.Node{node}
	}

	portals := makePortals(t, deps.Network, deps.Backends)
	for _, node := range nodes {
		if node.Stateless() {
			continue
		} else if testFunc.TestNode == nil {
			continue
		}

		t.Run(node.Name, func(t *testing.T) {
			t.Parallel()
			testFunc.TestNode(t, deps.Network, node, portals)
		})
	}

	if testFunc.TestPortal != nil {
		for _, portal := range portals {
			t.Run(portal.Chain.Name, func(t *testing.T) {
				t.Parallel()
				testFunc.TestPortal(t, deps.Network, portal, portals)
			})
		}
	}

	if testFunc.TestOmniEVM != nil {
		omniBackend, err := deps.Backends.Backend(deps.Network.ID.Static().OmniExecutionChainID)
		require.NoError(t, err)

		t.Run(omniBackend.Name(), func(t *testing.T) {
			t.Parallel()
			testFunc.TestOmniEVM(t, omniBackend)
		})
	}

	if testFunc.TestNetwork != nil {
		t.Run("network", func(t *testing.T) {
			t.Parallel()
			testFunc.TestNetwork(ctx, t, deps)
		})
	}
}

// makePortals creates a portal struct for each chain in the network.
func makePortals(t *testing.T, network netconf.Network, backends ethbackend.Backends) []Portal {
	t.Helper()
	resp := make([]Portal, 0, len(network.EVMChains()))
	for _, chain := range network.EVMChains() {
		if _, ok := evmchain.MetadataByID(chain.ID); !ok {
			t.Log("Skipping mock chain", chain.Name)
			continue
		}

		backend, err := backends.Backend(chain.ID)
		require.NoError(t, err)

		// create our Omni Portal Contract
		contract, err := bindings.NewOmniPortal(chain.PortalAddress, backend)
		require.NoError(t, err)
		require.NotNil(t, contract, "contract is nil")

		resp = append(resp, Portal{
			Chain:    chain,
			Client:   backend,
			Contract: contract,
		})
	}

	return resp
}

// loadEnv loads the testnet and network based on env vars.
func loadEnv(t *testing.T) NetworkDeps {
	t.Helper()

	manifestFile := os.Getenv(app.EnvE2EManifest)
	if manifestFile == "" {
		t.Skip(app.EnvE2EManifest + " not set, not an end-to-end test run")
	}
	if !filepath.IsAbs(manifestFile) {
		require.Fail(t, app.EnvE2EManifest+" must be an absolute path", "got", manifestFile)
	}

	ifdType := os.Getenv(app.EnvInfraType)
	ifdFile := os.Getenv(app.EnvInfraFile)
	if ifdType != docker.ProviderName && ifdFile == "" {
		require.Fail(t, app.EnvInfraFile+" not set while INFRASTRUCTURE_TYPE="+ifdType)
	} else if ifdType != docker.ProviderName && !filepath.IsAbs(ifdFile) {
		require.Fail(t, app.EnvInfraFile+" must be an absolute path", "got", ifdFile)
	}

	depsCacheMtx.Lock()
	defer depsCacheMtx.Unlock()
	if deps, ok := depsCache[manifestFile]; ok {
		return deps
	}

	m, err := app.LoadManifest(manifestFile)
	require.NoError(t, err)
	feature.SetGlobals(m.FeatureFlags)

	var ifd types.InfrastructureData
	switch ifdType {
	case docker.ProviderName:
		ifd, err = docker.NewInfraData(m)
	case vmcompose.ProviderName:
		ifd, err = vmcompose.LoadData(t.Context(), ifdFile)
	default:
		require.Fail(t, "unsupported infrastructure type", ifdType)
	}
	require.NoError(t, err)

	cfg := app.DefinitionConfig{
		ManifestFile: manifestFile,
	}
	testnet, err := app.TestnetFromManifest(t.Context(), m, ifd, cfg)
	require.NoError(t, err)

	endpointsFile := os.Getenv(app.EnvE2ERPCEndpoints)
	if endpointsFile == "" {
		t.Fatalf(app.EnvE2ERPCEndpoints + " not set")
	}
	bz, err := os.ReadFile(endpointsFile)
	require.NoError(t, err)
	var endpoints xchain.RPCEndpoints
	require.NoError(t, json.Unmarshal(bz, &endpoints))

	portalReg, err := makePortalRegistry(testnet.Network, endpoints)
	require.NoError(t, err)

	network, err := netconf.AwaitOnExecutionChain(t.Context(), testnet.Network, portalReg, endpoints.Keys())
	require.NoError(t, err)

	backends, err := ethbackend.BackendsFromTestnet(t.Context(), testnet)
	require.NoError(t, err)

	deps := NetworkDeps{
		Testnet:      testnet,
		Backends:     backends,
		Network:      network,
		RPCEndpoints: endpoints,
		SolverAddr:   testnet.SolverExternalAddr,
	}
	depsCache[manifestFile] = deps

	return deps
}

// fetchBlockChain fetches a complete, up-to-date block history from
// the freshest testnet archive node.
func fetchBlockChain(ctx context.Context, t *testing.T) []*cmttypes.Block {
	t.Helper()

	deps := loadEnv(t)
	testnet := deps.Testnet

	// Find the freshest archive node
	var (
		client *rpchttp.HTTP
		status *rpctypes.ResultStatus
	)
	for _, node := range testnet.ArchiveNodes() {
		c, err := node.Client()
		require.NoError(t, err)
		s, err := c.Status(ctx)
		require.NoError(t, err)
		if status == nil || s.SyncInfo.LatestBlockHeight > status.SyncInfo.LatestBlockHeight {
			client = c
			status = s
		}
	}
	require.NotNil(t, client, "couldn't find an archive node")

	// Fetch blocks. Look for existing block history in the block cache, and
	// extend it with any new blocks that have been produced.
	blocksCacheMtx.Lock()
	defer blocksCacheMtx.Unlock()

	from := status.SyncInfo.EarliestBlockHeight
	to := status.SyncInfo.LatestBlockHeight
	blocks, ok := blocksCache[testnet.Name]
	if !ok {
		blocks = make([]*cmttypes.Block, 0, to-from+1)
	}
	if len(blocks) > 0 {
		from = blocks[len(blocks)-1].Height + 1
	}

	for h := from; h <= to; h++ {
		resp, err := client.Block(ctx, &(h))
		require.NoError(t, ctx.Err(), "Timeout fetching all blocks: %d of %d", h, to)
		require.NoError(t, err)
		require.NotNil(t, resp.Block)
		require.Equal(t, h, resp.Block.Height, "unexpected block height %v", resp.Block.Height)
		blocks = append(blocks, resp.Block)
	}
	require.NotEmpty(t, blocks, "blockchain does not contain any blocks")
	blocksCache[testnet.Name] = blocks

	return blocks
}

func makePortalRegistry(network netconf.ID, endpoints xchain.RPCEndpoints) (*bindings.PortalRegistry, error) {
	meta := netconf.MetadataByID(network, network.Static().OmniExecutionChainID)
	rpc, err := endpoints.ByNameOrID(meta.Name, meta.ChainID)
	if err != nil {
		return nil, err
	}

	ethCl, err := ethclient.Dial(meta.Name, rpc)
	if err != nil {
		return nil, err
	}

	resp, err := bindings.NewPortalRegistry(common.HexToAddress(predeploys.PortalRegistry), ethCl)
	if err != nil {
		return nil, errors.Wrap(err, "create portal registry")
	}

	return resp, nil
}
