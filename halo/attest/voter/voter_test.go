package voter_test

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"testing"
	"time"

	"github.com/omni-network/omni/halo/attest/types"
	"github.com/omni-network/omni/halo/attest/voter"
	vtypes "github.com/omni-network/omni/halo/valsync/types"
	"github.com/omni-network/omni/lib/k1util"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"

	k1 "github.com/cometbft/cometbft/crypto/secp256k1"

	"github.com/ethereum/go-ethereum/common"

	fuzz "github.com/google/gofuzz"
	"github.com/stretchr/testify/require"
)

func TestRunner(t *testing.T) {
	t.Parallel()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	path := filepath.Join(t.TempDir(), "state.json")
	err := voter.GenEmptyStateFile(path)
	require.NoError(t, err)

	pk := k1.GenPrivKey()
	require.NoError(t, err)

	const (
		chain1     = 1
		conf       = xchain.ConfFinalized
		isVal      = true
		isNotVal   = false
		returnsOk  = true
		returnsErr = false
	)

	network := testNetwork(chain1)
	prov := make(stubProvider)
	backoff := new(testBackOff)
	deps := &mockDeps{}
	v := voter.LoadVoterForT(t, pk, path, prov, deps, network, backoff.BackOff)

	// callback is a helper function that calls the callback and asserts the error.
	callback := func(t *testing.T, sub sub, height uint64, isVal, ok bool) {
		t.Helper()
		setIsVal(t, v, pk, isVal)
		err := sub.callback(ctx, xchain.Block{
			BlockHeader: xchain.BlockHeader{
				ChainID:     chain1,
				BlockHeight: height,
			},
			Msgs: []xchain.Msg{{}}, // Non-empty XBlock should always be attested to
		})

		if ok {
			require.NoError(t, err)
		} else {
			require.Error(t, err)
			sub.result <- err // Unblock the synchronous caller
		}
	}

	// Start
	v.Start(ctx)

	// Assert it restarts a few (3) times, while not validator
	require.Eventually(t, func() bool {
		return backoff.Count() > 3
	}, time.Second, time.Millisecond)

	// Enable IsValidator, this will start a subscription
	setIsVal(t, v, pk, true)

	sub := <-prov // Get the subscription
	require.EqualValues(t, chain1, sub.req.ChainID)
	require.EqualValues(t, 0, sub.req.Height)

	callback(t, sub, 0, isVal, returnsOk) // Callback block 0 (in window)
	callback(t, sub, 1, isVal, returnsOk) // Callback block 1 (after window)

	v.TrimBehind(minByChain(network, chain1, 3)) // Set window to 3
	callback(t, sub, 2, isVal, returnsErr)       // Callback block 2 (before window) (triggers reset of worker)

	// Assert it reset
	sub = <-prov // Get the new subscription
	require.EqualValues(t, chain1, sub.req.ChainID)
	require.EqualValues(t, 2, sub.req.Height) // Assert it starts from 2 this time

	v.TrimBehind(minByChain(network, chain1, 2)) // Set window to 2
	callback(t, sub, 2, isVal, returnsOk)        // Callback block 2

	callback(t, sub, 3, isNotVal, returnsErr) // Callback block 3, but not validator anymore (triggers reset of worker)

	const newHeight = 55
	const newOffset = 77
	deps.SetHeightAndOffset(newHeight, newOffset) // Set a new latest attestation height

	// Enable IsValidator again
	setIsVal(t, v, pk, true)

	// Assert it reset
	sub = <-prov // Get the new subscription
	require.EqualValues(t, chain1, sub.req.ChainID)
	require.EqualValues(t, newHeight+1, sub.req.Height) // Assert it starts from newHeight+1 this time

	cancel()
	v.WaitDone()
}

func TestVoteWindow(t *testing.T) {
	t.Parallel()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	path := filepath.Join(t.TempDir(), "state.json")
	err := voter.GenEmptyStateFile(path)
	require.NoError(t, err)

	pk := k1.GenPrivKey()
	const chain1 = 1

	network := testNetwork(chain1)
	backoff := new(testBackOff)
	prov := make(stubProvider)
	deps := &mockDeps{}
	v := voter.LoadVoterForT(t, pk, path, prov, deps, network, backoff.BackOff)
	require.NoError(t, err)
	setIsVal(t, v, pk, true)

	v.Start(ctx)
	expectSubscriptions(t, prov, chain1, 0)

	w := &wrappedVoter{v: v, f: fuzz.New().NilChance(0).NumElements(1, 64), consensusChainID: network.ID.Static().OmniConsensusChainIDUint64()}

	// Add 1,2,3
	w.Add(t, chain1, 1)
	w.Add(t, chain1, 2)
	w.Add(t, chain1, 3)

	// Ensure all are available
	w.Available(t, chain1, 1, true)
	w.Available(t, chain1, 2, true)
	w.Available(t, chain1, 3, true)

	// Propose 1
	w.Propose(t, chain1, 1)

	// Trim behind 3 (deletes 1 and 2)
	l := w.v.TrimBehind(minByChain(network, chain1, 3))
	require.EqualValues(t, 2, l)

	w.Available(t, chain1, 1, false)
	w.Available(t, chain1, 2, false)
	w.Available(t, chain1, 3, true)

	// Trim behind 4 (deletes 3)
	l = w.v.TrimBehind(minByChain(network, chain1, 4))
	require.EqualValues(t, 1, l)

	w.Available(t, chain1, 1, false)
	w.Available(t, chain1, 2, false)
	w.Available(t, chain1, 3, false)

	// Ensure latest by chain not trimmed.
	latest, ok := w.v.LatestByChain(xchain.ChainVersion{ID: chain1, ConfLevel: xchain.ConfFinalized})
	require.True(t, ok)
	require.EqualValues(t, 3, latest.AttestHeader.AttestOffset)
}

func TestVoter(t *testing.T) {
	t.Parallel()
	fuzzer := fuzz.New().NilChance(0).NumElements(1, 64)

	path := filepath.Join(t.TempDir(), "state.json")
	err := voter.GenEmptyStateFile(path)
	require.NoError(t, err)

	pk := k1.GenPrivKey()

	const (
		chain1 = 1
		chain2 = 2
		chain3 = 3
	)

	network := testNetwork(chain1, chain2, chain3)

	// reloadVoter reloads the voter from disk. Asserting it starts streaming from the given heights.
	cancel := context.CancelFunc(func() {})
	reloadVoter := func(t *testing.T, from1, from2 uint64) *wrappedVoter {
		t.Helper()

		p := make(stubProvider)
		backoff := new(testBackOff)
		v := voter.LoadVoterForT(t, pk, path, p, stubDeps{}, network, backoff.BackOff)
		require.NoError(t, err)
		setIsVal(t, v, pk, true)

		cancel()
		var ctx context.Context
		ctx, cancel = context.WithCancel(context.Background())

		v.Start(ctx)

		expectSubscriptions(t, p,
			chain1, from1,
			chain2, from2,
			chain3, 0,
		)

		return &wrappedVoter{v: v, f: fuzzer, consensusChainID: network.ID.Static().OmniConsensusChainIDUint64()}
	}

	v := reloadVoter(t, 0, 0)

	// Add 1, 2, 3 (and assert they are available)
	v.Add(t, 1, 1)
	v.Add(t, 1, 2)
	v.Add(t, 1, 3)

	// Reload
	v = reloadVoter(t, 4, 0)

	// Add noise
	v.Add(t, 2, 1)

	// Assert all are still available
	v.Available(t, 1, 1, true)
	v.Available(t, 1, 2, true)
	v.Available(t, 1, 3, true)
	v.Available(t, 2, 1, true)

	// Reload
	v = reloadVoter(t, 4, 2)

	// Propose and commit 3 only
	v.Propose(t, 1, 3)
	v.Commit(t, 1, 3)

	// Assert 1, 2 are available
	v.Available(t, 1, 1, true)
	v.Available(t, 1, 2, true)
	v.Available(t, 1, 3, false)

	// Reload
	v = reloadVoter(t, 4, 2)

	// Propose 1
	v.Propose(t, 1, 1)
	v.Available(t, 1, 1, false)

	// Propose 2 (resets 1)
	v.Propose(t, 1, 2)
	v.Available(t, 1, 1, true)
	v.Available(t, 1, 2, false)

	// Commit 1 (resets 2)
	v.Commit(t, 1, 1)
	v.Available(t, 1, 1, false)
	v.Available(t, 1, 2, true)

	// Reload
	v = reloadVoter(t, 4, 2)

	// Commit 2 and noise
	v.Commit(t, 1, 2)
	v.Commit(t, 2, 1)

	// All committed
	bz, err := os.ReadFile(path)
	require.NoError(t, err)
	var stateJSON map[string]any
	require.NoError(t, json.Unmarshal(bz, &stateJSON))

	require.Len(t, stateJSON, 4)
	require.Empty(t, stateJSON["available"])
	require.Empty(t, stateJSON["proposed"])
	require.Len(t, stateJSON["committed"], 2) // One per chain
	require.Len(t, stateJSON["latest"], 2)    // One per chain

	v.AddErr(t, 1, 3)
	v.AddErr(t, 1, 2)
	v.AddErr(t, 1, 5)
}

func expectSubscriptions(t *testing.T, prov stubProvider, chainOffsets ...uint64) {
	t.Helper()

	expected := make(map[uint64]uint64)
	for i := 0; i < len(chainOffsets); i += 2 {
		chainID := chainOffsets[i]
		offset := chainOffsets[i+1]
		expected[chainID] = offset
	}

	l := len(expected)
	for i := 0; i < l; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		select {
		case <-ctx.Done():
			require.Fail(t, "timed out waiting for subscription")
		case next := <-prov:
			_, ok := expected[next.req.ChainID]
			require.True(t, ok)
			delete(expected, next.req.ChainID)
		}
		cancel()
	}
	require.Empty(t, expected)
}

var _ xchain.Provider = make(stubProvider)
var _ types.VoterDeps = stubDeps{}
var _ types.VoterDeps = new(mockDeps)

type wrappedVoter struct {
	v                *voter.Voter
	f                *fuzz.Fuzzer
	consensusChainID uint64
}

func (w *wrappedVoter) Add(t *testing.T, chainID, offset uint64) {
	t.Helper()
	var block xchain.Block
	w.f.Fuzz(&block)

	attHeader := xchain.AttestHeader{
		ConsensusChainID: w.consensusChainID,
		ChainVersion:     xchain.NewChainVersion(chainID, xchain.ConfFinalized),
		AttestOffset:     offset,
	}

	block.BlockHeader = xchain.BlockHeader{
		ChainID: chainID,

		BlockHash: common.Hash{},
	}

	err := w.v.Vote(attHeader, block, false)
	require.NoError(t, err)
}

func (w *wrappedVoter) Propose(t *testing.T, chainID, offset uint64) {
	t.Helper()

	header := &types.AttestHeader{
		SourceChainId:    chainID,
		ConsensusChainId: w.consensusChainID,
		ConfLevel:        uint32(xchain.ConfFinalized),
		AttestOffset:     offset,
	}

	err := w.v.SetProposed([]*types.AttestHeader{header})
	require.NoError(t, err)
}

func (w *wrappedVoter) Commit(t *testing.T, chainID, offset uint64) {
	t.Helper()

	header := &types.AttestHeader{
		SourceChainId:    chainID,
		ConsensusChainId: w.consensusChainID,
		ConfLevel:        uint32(xchain.ConfFinalized),
		AttestOffset:     offset,
	}

	err := w.v.SetCommitted([]*types.AttestHeader{header})
	require.NoError(t, err)
}

// Available asserts the given block is available.
func (w *wrappedVoter) Available(t *testing.T, chainID, offset uint64, ok bool) {
	t.Helper()

	var found bool
	for _, att := range w.v.GetAvailable() {
		if att.AttestHeader.SourceChainId == chainID && att.AttestHeader.AttestOffset == offset && att.AttestHeader.ConfLevel == uint32(xchain.ConfFinalized) {
			found = true
			break
		}
	}

	require.Equal(t, ok, found)
}

// AddErr adds a fuzzed block to the voter and asserts an error is returned.
func (w *wrappedVoter) AddErr(t *testing.T, chainID, offset uint64) {
	t.Helper()
	var block xchain.Block
	w.f.Fuzz(&block)
	block.BlockHeader = xchain.BlockHeader{
		ChainID: chainID,
	}

	attHeader := xchain.AttestHeader{
		ConsensusChainID: w.consensusChainID,
		ChainVersion:     xchain.NewChainVersion(chainID, xchain.ConfFinalized),
		AttestOffset:     offset,
	}

	err := w.v.Vote(attHeader, block, false)
	require.Error(t, err)
}

type mockDeps struct {
	mu     sync.Mutex
	height uint64
	offset uint64
}

func (m *mockDeps) SetHeightAndOffset(height, offset uint64) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.height = height
	m.offset = offset
}

func (m *mockDeps) LatestAttestation(context.Context, xchain.ChainVersion) (xchain.Attestation, bool, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	return xchain.Attestation{
		AttestHeader: xchain.AttestHeader{
			AttestOffset: m.offset,
			ChainVersion: xchain.ChainVersion{ConfLevel: xchain.ConfFinalized},
		},
		BlockHeader: xchain.BlockHeader{
			BlockHeight: m.height,
		},
	}, m.height > 0, nil
}

type stubDeps struct{}

func (stubDeps) LatestAttestation(context.Context, xchain.ChainVersion) (xchain.Attestation, bool, error) {
	return xchain.Attestation{}, false, nil
}

type sub struct {
	req      xchain.ProviderRequest
	callback xchain.ProviderCallback
	result   chan error
}

var _ xchain.Provider = make(stubProvider)

type stubProvider chan sub

func (p stubProvider) StreamBlocks(ctx context.Context, req xchain.ProviderRequest, callback xchain.ProviderCallback) error {
	result := make(chan error)

	p <- sub{req, callback, result}

	select {
	case <-ctx.Done():
		return ctx.Err()
	case err := <-result:
		return err
	}
}

func (stubProvider) StreamAsync(context.Context, xchain.ProviderRequest, xchain.ProviderCallback) error {
	panic("unexpected")
}

func (stubProvider) GetBlock(context.Context, xchain.ProviderRequest) (xchain.Block, bool, error) {
	panic("unexpected")
}

func (stubProvider) GetSubmittedCursor(context.Context, xchain.StreamID) (xchain.SubmitCursor, bool, error) {
	panic("unexpected")
}

func (stubProvider) GetEmittedCursor(context.Context, xchain.EmitRef, xchain.StreamID) (xchain.EmitCursor, bool, error) {
	panic("unexpected")
}

func (stubProvider) ChainVersionHeight(context.Context, xchain.ChainVersion) (uint64, error) {
	panic("unexpected")
}

type testBackOff struct {
	mu      sync.Mutex
	backoff int
}

func (b *testBackOff) Count() int {
	b.mu.Lock()
	defer b.mu.Unlock()

	return b.backoff
}

func (b *testBackOff) BackOff() {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.backoff++
}

func setIsVal(t *testing.T, v *voter.Voter, pk k1.PrivKey, isVal bool) {
	t.Helper()

	cmtPubkey, err := k1util.PBPubKeyFromBytes(pk.PubKey().Bytes())
	require.NoError(t, err)

	vals := []*vtypes.Validator{{
		ConsensusPubkey: k1.GenPrivKey().PubKey().Bytes(),
		Power:           1,
	}}
	if isVal {
		vals = append(vals, &vtypes.Validator{
			ConsensusPubkey: cmtPubkey.GetSecp256K1(),
			Power:           1,
		})
	}

	err = v.UpdateValidatorSet(&vtypes.ValidatorSetResponse{
		Id:         uint64(time.Now().UnixNano()),
		Validators: vals,
	})
	require.NoError(t, err)
}

func testNetwork(chainIDs ...uint64) netconf.Network {
	var chains []netconf.Chain
	for _, id := range chainIDs {
		chains = append(chains, netconf.Chain{
			ID:     id,
			Name:   fmt.Sprintf("chain_%d", id),
			Shards: []xchain.ShardID{xchain.ShardFinalized0},
		})
	}

	return netconf.Network{
		ID:     netconf.Simnet,
		Chains: chains,
	}
}

//nolint:unparam // ChainID will change in future
func minByChain(network netconf.Network, chainID uint64, min uint64) map[xchain.ChainVersion]uint64 {
	chain, _ := network.Chain(chainID)

	return map[xchain.ChainVersion]uint64{
		chain.ChainVersions()[0]: min,
	}
}
