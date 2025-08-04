// Only run this if -race=false, since CosmosSDK has known data races when doing gRPC queries.
//go:build !race

//nolint:paralleltest // CosmosSDK dependency prevents parallel execution
package app_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"os"
	"testing"
	"time"

	haloapp "github.com/omni-network/omni/halo/app"
	magellan2 "github.com/omni-network/omni/halo/app/upgrades/magellan"
	"github.com/omni-network/omni/halo/app/upgrades/static"
	halocmd "github.com/omni-network/omni/halo/cmd"
	halocfg "github.com/omni-network/omni/halo/config"
	"github.com/omni-network/omni/lib/cchain"
	cprovider "github.com/omni-network/omni/lib/cchain/provider"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/tutil"
	"github.com/omni-network/omni/lib/xchain"

	rpchttp "github.com/cometbft/cometbft/rpc/client/http"
	"github.com/cometbft/cometbft/types"

	db "github.com/cosmos/cosmos-db"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	"github.com/stretchr/testify/require"
)

func TestSmoke(t *testing.T) {
	ctx, cancel := context.WithCancel(t.Context())
	defer cancel()

	ctx, err := log.Init(ctx, log.Config{Color: log.ColorForce, Level: "debug", Format: log.FormatConsole})
	require.NoError(t, err)

	cfg := setupSimnet(t)

	encConf, err := haloapp.ClientEncodingConfig(ctx, netconf.Simnet)
	require.NoError(t, err)

	// Start the server async
	async, stopfunc, err := haloapp.Start(ctx, cfg)
	require.NoError(t, err)
	go func() {
		tutil.RequireNoError(t, <-async)
	}()

	// Connect to the server.
	cl, err := rpchttp.New(cfg.Comet.RPC.ListenAddress, "/websocket")
	require.NoError(t, err)

	cprov := cprovider.NewABCI(cl, netconf.Simnet)
	cprovGRPC, err := cprovider.NewGRPC(cfg.SDKGRPC.Address, netconf.Simnet, encConf.InterfaceRegistry)
	require.NoError(t, err)

	// Wait until we get to block 3.
	const target = uint64(3)
	require.Eventually(t, func() bool {
		s, err := cl.Status(ctx)
		if err != nil {
			t.Log("Failed to get status: ", err)
			return false
		}

		return s.SyncInfo.LatestBlockHeight >= int64(target)
	}, time.Second*time.Duration(target*2), time.Millisecond*100)

	testReadyEndpoint(t, cfg)
	testAPI(t, cfg)
	testCProvider(t, ctx, cprov)
	testCProvider(t, ctx, cprovGRPC)
	testModuleParams(t, ctx, cprov)

	genSet, err := cl.Validators(ctx, int64Ptr(1), nil, nil)
	require.NoError(t, err)
	genSetHash := types.NewValidatorSet(genSet.Validators).Hash()

	// Wait for cometBFT validator set to change
	require.Eventually(t, func() bool {
		set, err := cl.Validators(ctx, nil, nil, nil)
		require.NoError(t, err)
		setHash := types.NewValidatorSet(set.Validators).Hash()

		return !bytes.Equal(genSetHash, setHash)
	}, time.Second*time.Duration(target*4), time.Millisecond*100)

	srcChain := netconf.Simnet.Static().OmniExecutionChainID
	chainVer := xchain.ChainVersion{ID: srcChain, ConfLevel: xchain.ConfFinalized}

	// Ensure all blocks are attested and approved.
	cprov.StreamAsync(ctx, chainVer, 1, "test", func(ctx context.Context, approved xchain.Attestation) error {
		// Sanity check we can fetch latest directly as well.
		att, ok, err := cprov.LatestAttestation(ctx, chainVer)
		tutil.RequireNoError(t, err)
		require.True(t, ok)
		require.Equal(t, srcChain, att.ChainID)

		require.Equal(t, srcChain, approved.ChainID)
		t.Logf("cprovider streamed approved block: %d", approved.AttestOffset)
		if approved.BlockHeight >= target {
			cancel()
		}

		return nil
	})

	<-ctx.Done()

	// Stop the server, with a fresh context
	require.NoError(t, stopfunc(t.Context()))
}

func testModuleParams(t *testing.T, ctx context.Context, cprov cchain.Provider) {
	t.Helper()

	sParamsResp, err := cprov.QueryClients().Slashing.Params(ctx, &slashingtypes.QueryParamsRequest{})
	require.NoError(t, err)
	expected := magellan2.SlashingParams()
	require.Equal(t, expected.String(), sParamsResp.Params.String())

	mParamsResp, err := cprov.QueryClients().Mint.Params(ctx, &minttypes.QueryParamsRequest{})
	require.NoError(t, err)
	require.Equal(t, magellan2.MintParams.String(), mParamsResp.Params.String())

	inflResponse, err := cprov.QueryClients().Mint.Inflation(ctx, &minttypes.QueryInflationRequest{})
	require.NoError(t, err)
	require.Equal(t, magellan2.MintParams.InflationMin.String(), inflResponse.Inflation.String())
}

func testAPI(t *testing.T, cfg haloapp.Config) {
	t.Helper()

	u, err := url.Parse(cfg.SDKAPI.Address)
	require.NoError(t, err)

	base := "http://" + u.Host

	for _, path := range []string{
		"/",
		"/status",
		"/cosmos/staking/v1beta1/validators",
		"/cosmos/slashing/v1beta1/signing_infos",
	} {
		_, err = http.Get(base + path)
		require.NoError(t, err)
	}
}

func testReadyEndpoint(t *testing.T, cfg haloapp.Config) {
	t.Helper()

	base := "http://0.0.0.0" + cfg.Comet.Instrumentation.PrometheusListenAddr

	response, err := http.Get(base + "/ready")
	require.NoError(t, err)
	defer response.Body.Close()

	type status struct {
		ConsensusSynced bool `json:"consensus_synced"`
		ExecutionSynced bool `json:"execution_synced"`
	}

	var readyResponse status
	err = json.NewDecoder(response.Body).Decode(&readyResponse)
	// We only check that the endpoint returns a parsable response.
	require.NoError(t, err)
}

func testCProvider(t *testing.T, ctx context.Context, cprov cprovider.Provider) {
	t.Helper()

	// Ensure it doesn't error for unknown chains.
	_, ok, err := cprov.LatestAttestation(ctx, xchain.ChainVersion{})
	require.NoError(t, err)
	require.False(t, ok)

	// Ensure getting latest xblock.
	xblock, ok, err := cprov.XBlock(ctx, 0, true)
	tutil.RequireNoError(t, err)
	require.True(t, ok)
	require.Len(t, xblock.Msgs, 2)
	for i, msg := range xblock.Msgs {
		require.Equal(t, xchain.ShardBroadcast0, msg.ShardID)
		require.Equal(t, xchain.BroadcastChainID, msg.DestChainID)
		require.EqualValues(t, i, msg.LogIndex)
	}

	// Ensure getting latest xblock.
	xblock2, ok, err := cprov.XBlock(ctx, xblock.BlockHeight, false)
	tutil.RequireNoError(t, err)
	require.True(t, ok)
	require.Equal(t, xblock, xblock2)

	// Ensure it doesn't error for unknown validator sets.
	_, ok, err = cprov.PortalValidatorSet(ctx, 33)
	require.NoError(t, err)
	require.False(t, ok)

	exec, cons, err := cprov.GenesisFiles(ctx)
	require.NoError(t, err)
	require.Nil(t, exec)
	require.NotNil(t, cons)

	require.Eventually(t, func() bool {
		_, ok, err := cprov.CurrentPlannedPlan(ctx)
		require.NoError(t, err)

		return ok
	}, time.Second*5, time.Millisecond*100)

	resp, err := cprov.QueryClients().Mint.Inflation(ctx, &minttypes.QueryInflationRequest{})
	require.NoError(t, err)
	require.Equal(t, "0.110000000000000000", resp.Inflation.String())
}

func setupSimnet(t *testing.T) haloapp.Config {
	t.Helper()
	homeDir := t.TempDir()

	cmtCfg := halocmd.DefaultCometConfig(homeDir)
	cmtCfg.DBBackend = string(db.MemDBBackend)
	cmtCfg.P2P.ListenAddress = tutil.RandomListenAddress(t) // Avoid port clashes
	cmtCfg.RPC.ListenAddress = tutil.RandomListenAddress(t) // Avoid port clashes
	cmtCfg.Instrumentation.Prometheus = false

	haloCfg := halocfg.DefaultConfig()
	haloCfg.HomeDir = homeDir
	haloCfg.Network = netconf.Simnet
	haloCfg.BackendType = string(db.MemDBBackend)
	haloCfg.EVMBuildDelay = time.Millisecond
	haloCfg.EngineEndpoint = "dummy"
	haloCfg.EngineJWTFile = "dummy"
	haloCfg.RPCEndpoints = map[string]string{"omni_evm": "dummy"}
	haloCfg.SDKAPI.Enable = true

	cfg := haloapp.Config{
		Config: haloCfg,
		Comet:  cmtCfg,
	}

	executionGenesis, err := ethclient.MockGenesisBlock()
	tutil.RequireNoError(t, err)

	err = halocmd.InitFiles(log.WithNoopLogger(t.Context()), halocmd.InitConfig{
		HomeDir:        homeDir,
		Network:        netconf.Simnet,
		ExecutionHash:  executionGenesis.Hash(),
		GenesisUpgrade: static.LatestUpgrade(),
	})
	tutil.RequireNoError(t, err)

	// CometBFT doesn't shutdown cleanly. It leaves goroutines running that write to disk.
	// The test sometimes fails with: TempDir RemoveAll cleanup: unlinkat ... directory not empty
	// Manually retry deleting everything a few times. This should prevent to test from flapping.
	t.Cleanup(func() {
		for i := 0; i < 5; i++ {
			err := os.RemoveAll(homeDir)
			if err == nil {
				break
			}
			time.Sleep(time.Millisecond * 500)
		}
	})

	return cfg
}

func int64Ptr(i int64) *int64 {
	return &i
}
