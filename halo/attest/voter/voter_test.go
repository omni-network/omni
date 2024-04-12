package voter_test

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"sync"
	"testing"
	"time"

	"github.com/omni-network/omni/halo/attest/types"
	"github.com/omni-network/omni/halo/attest/voter"
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
	const (
		chain1     = 1
		isVal      = true
		isNotVal   = false
		returnsOk  = true
		returnsErr = false
	)

	prov := make(stubProvider)
	deps := &mockDeps{isValChan: make(chan bool)}
	v := voter.LoadVoterForT(t, pk, path, prov, deps, map[uint64]string{chain1: ""})

	// callback is a helper function that calls the callback and asserts the error.
	callback := func(t *testing.T, sub sub, deps *mockDeps, height uint64, isVal, ok bool) {
		t.Helper()
		go func() { deps.isValChan <- isVal }() // IsValidator should be called first.
		err := sub.callback(ctx, xchain.Block{
			BlockHeader: xchain.BlockHeader{
				SourceChainID: chain1,
				BlockHeight:   height,
			},
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

	// Assert it restarts a few time, checking IsValidator
	deps.isValChan <- isNotVal
	deps.isValChan <- isNotVal

	// Enable IsValidator, this will start a subscription
	deps.isValChan <- isVal

	sub := <-prov // Get the subscription
	require.EqualValues(t, chain1, sub.chainID)
	require.EqualValues(t, 0, sub.height)

	callback(t, sub, deps, 0, isVal, returnsOk) // Callback block 0 (in window)
	callback(t, sub, deps, 1, isVal, returnsOk) // Callback block 1 (after window)

	deps.SetWindow(3)                            // Set window to 3
	callback(t, sub, deps, 2, isVal, returnsErr) // Callback block 2 (before window) (triggers reset of worker)

	// Assert it reset
	deps.isValChan <- isVal
	sub = <-prov // Get the new subscription
	require.EqualValues(t, chain1, sub.chainID)
	require.EqualValues(t, 2, sub.height) // Assert it starts from 2 this time

	deps.SetWindow(2)                           // Set window to 2
	callback(t, sub, deps, 2, isVal, returnsOk) // Callback block 2

	callback(t, sub, deps, 3, isNotVal, returnsErr) // Callback block 3, but not validator anymore (triggers reset of worker)

	// Assert it reset
	const newHeight = 99
	deps.SetHeight(newHeight) // Set a new latest attestation height
	deps.isValChan <- isVal
	sub = <-prov // Get the new subscription
	require.EqualValues(t, chain1, sub.chainID)
	require.EqualValues(t, newHeight+1, sub.height) // Assert it starts from newHeight+1 this time

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

	prov := make(stubProvider)
	deps := &mockDeps{isVal: true}
	v := voter.LoadVoterForT(t, pk, path, prov, deps, map[uint64]string{chain1: ""})
	require.NoError(t, err)

	v.Start(ctx)
	expectSubscriptions(t, prov, chain1, 0)

	w := &wrappedVoter{v: v, f: fuzz.New().NilChance(0).NumElements(1, 64)}

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
	l := w.v.TrimBehind(map[uint64]uint64{chain1: 3})
	require.EqualValues(t, 2, l)

	w.Available(t, chain1, 1, false)
	w.Available(t, chain1, 2, false)
	w.Available(t, chain1, 3, true)

	// Trim behind 4 (deletes 3)
	l = w.v.TrimBehind(map[uint64]uint64{chain1: 4})
	require.EqualValues(t, 1, l)

	w.Available(t, chain1, 1, false)
	w.Available(t, chain1, 2, false)
	w.Available(t, chain1, 3, false)

	// Ensure latest by chain not trimmed.
	latest, ok := w.v.LatestByChain(chain1)
	require.True(t, ok)
	require.EqualValues(t, 3, latest.BlockHeader.Height)
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

	// reloadVoter reloads the voter from disk. Asserting it starts streaming from the given heights.
	cancel := context.CancelFunc(func() {})
	reloadVoter := func(t *testing.T, from1, from2 uint64) *wrappedVoter {
		t.Helper()
		p := make(stubProvider)
		a := voter.LoadVoterForT(t, pk, path, p, stubDeps{}, map[uint64]string{chain1: "", chain2: "", chain3: ""})
		require.NoError(t, err)

		cancel()
		var ctx context.Context
		ctx, cancel = context.WithCancel(context.Background())

		a.Start(ctx)

		expectSubscriptions(t, p,
			chain1, from1,
			chain2, from2,
			chain3, 0,
		)

		return &wrappedVoter{v: a, f: fuzzer}
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

func expectSubscriptions(t *testing.T, prov stubProvider, chainHeights ...uint64) {
	t.Helper()

	expected := make(map[uint64]uint64)
	for i := 0; i < len(chainHeights); i += 2 {
		chainID := chainHeights[i]
		height := chainHeights[i+1]
		expected[chainID] = height
	}

	l := len(expected)
	for i := 0; i < l; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		select {
		case <-ctx.Done():
			require.Fail(t, "timed out waiting for subscription")
		case next := <-prov:
			h, ok := expected[next.chainID]
			require.True(t, ok)
			require.EqualValues(t, h, next.height)
			delete(expected, next.chainID)
		}
		cancel()
	}
	require.Empty(t, expected)
}

var _ xchain.Provider = make(stubProvider)
var _ types.VoterDeps = stubDeps{}
var _ types.VoterDeps = new(mockDeps)

type wrappedVoter struct {
	v *voter.Voter
	f *fuzz.Fuzzer
}

func (w *wrappedVoter) Add(t *testing.T, chainID, height uint64) {
	t.Helper()
	var block xchain.Block
	w.f.Fuzz(&block)
	block.BlockHeader = xchain.BlockHeader{
		SourceChainID: chainID,
		BlockHeight:   height,
	}

	err := w.v.Vote(block, false)
	require.NoError(t, err)
}

func (w *wrappedVoter) Propose(t *testing.T, chainID, height uint64) {
	t.Helper()

	header := &types.BlockHeader{
		ChainId: chainID,
		Height:  height,
	}

	err := w.v.SetProposed([]*types.BlockHeader{header})
	require.NoError(t, err)
}

func (w *wrappedVoter) Commit(t *testing.T, chainID, height uint64) {
	t.Helper()

	header := &types.BlockHeader{
		ChainId: chainID,
		Height:  height,
	}

	err := w.v.SetCommitted([]*types.BlockHeader{header})
	require.NoError(t, err)
}

// Available asserts the given block is available.
func (w *wrappedVoter) Available(t *testing.T, chainID, height uint64, ok bool) {
	t.Helper()

	var found bool
	for _, att := range w.v.GetAvailable() {
		if att.BlockHeader.ChainId == chainID && att.BlockHeader.Height == height {
			found = true
			break
		}
	}

	require.Equal(t, ok, found)
}

// AddErr adds a fuzzed block to the voter and asserts an error is returned.
func (w *wrappedVoter) AddErr(t *testing.T, chainID, height uint64) {
	t.Helper()
	var block xchain.Block
	w.f.Fuzz(&block)
	block.BlockHeader = xchain.BlockHeader{
		SourceChainID: chainID,
		BlockHeight:   height,
	}

	err := w.v.Vote(block, false)
	require.Error(t, err)
}

type mockDeps struct {
	mu        sync.Mutex
	isValChan chan bool // If non-nil, each IsValidator call will read from this.
	isVal     bool      // Else IsValidator will return this.
	window    uint64
	height    uint64
}

func (m *mockDeps) SetIsValidator(is bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.isVal = is
}

func (m *mockDeps) SetWindow(w uint64) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.window = w
}

func (m *mockDeps) SetHeight(h uint64) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.height = h
}

func (m *mockDeps) IsValidator(context.Context, common.Address) bool {
	if m.isValChan != nil {
		return <-m.isValChan
	}

	m.mu.Lock()
	defer m.mu.Unlock()

	return m.isVal
}

func (m *mockDeps) WindowCompare(_ context.Context, _ uint64, height uint64) (int, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if height < m.window {
		return -1, nil
	} else if height > m.window {
		return 1, nil
	}

	return 0, nil
}

func (m *mockDeps) LatestAttestationHeight(context.Context, uint64) (uint64, bool, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	return m.height, m.height > 0, nil
}

type stubDeps struct{}

func (a stubDeps) IsValidator(context.Context, common.Address) bool {
	return true
}

func (stubDeps) WindowCompare(context.Context, uint64, uint64) (int, error) {
	return 0, nil
}

func (stubDeps) LatestAttestationHeight(context.Context, uint64) (uint64, bool, error) {
	return 0, false, nil
}

type sub struct {
	chainID  uint64
	height   uint64
	callback xchain.ProviderCallback
	result   chan error
}

var _ xchain.Provider = make(stubProvider)

type stubProvider chan sub

func (p stubProvider) StreamBlocks(ctx context.Context, chainID uint64, fromHeight uint64, callback xchain.ProviderCallback) error {
	result := make(chan error)
	p <- sub{chainID, fromHeight, callback, result}

	select {
	case <-ctx.Done():
		return ctx.Err()
	case err := <-result:
		return err
	}
}

func (stubProvider) StreamAsync(context.Context, uint64, uint64, xchain.ProviderCallback) error {
	panic("unexpected")
}

func (stubProvider) GetBlock(context.Context, uint64, uint64) (xchain.Block, bool, error) {
	panic("unexpected")
}

func (stubProvider) GetSubmittedCursor(context.Context, uint64, uint64) (xchain.StreamCursor, bool, error) {
	panic("unexpected")
}

func (stubProvider) GetEmittedCursor(context.Context, uint64, uint64) (xchain.StreamCursor, bool, error) {
	panic("unexpected")
}
