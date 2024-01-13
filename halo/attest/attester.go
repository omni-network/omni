package attest

import (
	"context"
	"path/filepath"
	"sync"
	"testing"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/cometbft/cometbft/crypto"
)

// FIXME(corver): Attester drops available but not confirmed attestations on restart.

var _ Service = (*Attester)(nil)

// NewAttester returns a new attester.
func NewAttester(ctx context.Context, state State, privKey crypto.PrivKey, provider xchain.Provider, chains []uint64,
) (*Attester, error) {
	if len(privKey.PubKey().Bytes()) != 33 {
		return nil, errors.New("invalid private key")
	}

	a := newInternal(state, privKey)

	// Subscribe to provider
	for chainID, fromHeight := range fromHeights(state, chains) {
		err := provider.Subscribe(ctx, chainID, fromHeight, a.Attest)
		if err != nil {
			return nil, err
		}
	}

	return a, nil
}

// NewAttesterForT returns a new attester with empty state for testing.
func NewAttesterForT(t *testing.T, privKey crypto.PrivKey) *Attester {
	t.Helper()

	state := &FileState{
		atts: make(map[uint64]xchain.Attestation),
		path: filepath.Join(t.TempDir(), "state.json"),
	}

	return newInternal(state, privKey)
}

// newInternal returns a new attester. It is used by exported constructors.
func newInternal(state State, privKey crypto.PrivKey) *Attester {
	return &Attester{
		privKey: privKey,
		state:   state,
	}
}

// Attester implements the Service interface.
// It is responsible for creating and keeping track of attestations.
type Attester struct {
	privKey crypto.PrivKey
	state   State

	mu        sync.Mutex
	available []xchain.Attestation
	proposed  []xchain.Attestation
}

// Attest creates an attestation for the given block and adds it to the internal state.
func (a *Attester) Attest(_ context.Context, block *xchain.Block) error {
	att, err := CreateAttestation(a.privKey, *block)
	if err != nil {
		return err
	}

	if err := a.state.Add(att); err != nil {
		return err
	}

	a.addAvailable(att)

	return nil
}

// GetAvailable returns all the available attestations.
func (a *Attester) GetAvailable() []xchain.Attestation {
	a.mu.Lock()
	defer a.mu.Unlock()

	atts := make([]xchain.Attestation, len(a.available))
	copy(atts, a.available)

	return atts
}

// SetProposed sets the attestations as proposed.
func (a *Attester) SetProposed(headers []xchain.BlockHeader) {
	proposed := make(map[xchain.BlockHeader]bool)
	for _, header := range headers {
		proposed[header] = true
	}

	a.mu.Lock()
	defer a.mu.Unlock()

	var newAvailable, newProposed []xchain.Attestation
	for _, att := range append(a.available, a.proposed...) {
		// If proposed, move it to the proposed list.
		if proposed[att.BlockHeader] {
			newProposed = append(newProposed, att)
		} else { // Otherwise, keep or move it to in the available list.
			newAvailable = append(newAvailable, att)
		}
	}

	a.available = newAvailable
	a.proposed = newProposed
}

// SetCommitted sets the attestations as committed.
func (a *Attester) SetCommitted(headers []xchain.BlockHeader) {
	confirmed := make(map[xchain.BlockHeader]bool)
	for _, header := range headers {
		confirmed[header] = true
	}

	a.mu.Lock()
	defer a.mu.Unlock()

	var newAvailable []xchain.Attestation
	for _, att := range append(a.available, a.proposed...) {
		// If not confirmed, keep/move it to available.
		if !confirmed[att.BlockHeader] {
			newAvailable = append(newAvailable, att)
		}
	}

	a.available = newAvailable
	a.proposed = nil
}

func (a *Attester) LocalPubKey() [33]byte {
	bz := a.privKey.PubKey().Bytes()
	return [33]byte(bz)
}

// addAvailable adds the attestation to the available list.
func (a *Attester) addAvailable(att xchain.Attestation) {
	a.mu.Lock()
	defer a.mu.Unlock()

	a.available = append(a.available, att)
}

// fromHeights returns a map of heights to start streaming from (inclusive) by chain ID.
func fromHeights(state State, chains []uint64) map[uint64]uint64 {
	resp := make(map[uint64]uint64)

	// Initialize all chains to 0
	for _, chain := range chains {
		resp[chain] = 0
	}

	// Set heights for existing chains from state
	for _, att := range state.Get() {
		resp[att.SourceChainID] = att.BlockHeight + 1
	}

	return resp
}
