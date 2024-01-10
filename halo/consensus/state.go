package consensus

import (
	"encoding/json"
	"sync"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/xchain"

	abci "github.com/cometbft/cometbft/abci/types"

	"github.com/ethereum/go-ethereum/crypto"
)

const pubKeyLen = 33

type validator struct {
	PubKey [pubKeyLen]byte
	Power  int64
}
type state struct {
	mu                 sync.Mutex
	ExecutionStateRoot [32]byte
	valSetID           uint64
	validators         []validator
	pendingAggs        []xchain.AggAttestation
	approvedAggs       []xchain.AggAttestation
}

// InitChainState sets consensus validators.
// It returns an error if the validators are already set.
// This implies validators can only be set once and never change.
func (s *state) InitChainState(req *abci.RequestInitChain) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.valSetID > 0 {
		return errors.New("chain already initialized")
	}

	if len(req.AppStateBytes) > 0 {
		return errors.New("app state bytes must be empty")
	}

	s.valSetID++ // Update from 0 to 1.

	for _, v := range req.Validators {
		if len(v.PubKey.GetEd25519()) > 0 {
			return errors.New("ed25519 keys not supported")
		}

		pk := v.PubKey.GetSecp256K1()
		if len(pk) != pubKeyLen {
			return errors.New("invalid pubkey length")
		}

		s.validators = append(s.validators, validator{
			PubKey: [pubKeyLen]byte(pk),
			Power:  v.Power,
		})
	}

	return nil
}

// AppHash returns the application hash.
func (s *state) AppHash() ([]byte, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	buf, err := json.Marshal(appHashJSON{
		ValSetID:     s.valSetID,
		Validators:   s.validators,
		PendingAggs:  s.pendingAggs,
		ApprovedAggs: s.approvedAggs,
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal app hash json")
	}

	return crypto.Keccak256(buf), nil
}

// AddAttestations adds the provided aggregates to the state, moving pending to approved if applicable.
func (s *state) AddAttestations(aggregates []xchain.AggAttestation) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Merge the new aggregates with existing, either pending or approved, add non-matching to remaining.
	var remaining []xchain.AggAttestation
	s.pendingAggs, remaining = mergeAggregates(s.pendingAggs, aggregates)
	s.approvedAggs, remaining = mergeAggregates(s.approvedAggs, remaining)

	// Add remaining non-matching to pending.
	s.pendingAggs = sortAggregates(append(s.pendingAggs, remaining...))

	// Check which pending are newly approved, and which are still pending.
	var stillPending, newApproved []xchain.AggAttestation
	for _, agg := range s.pendingAggs {
		if isApproved(agg, s.validators) {
			newApproved = append(newApproved, agg)
		} else {
			stillPending = append(stillPending, agg)
		}
	}

	// Update pending and approved.
	s.pendingAggs = sortAggregates(stillPending)
	s.approvedAggs = sortAggregates(append(s.approvedAggs, newApproved...))
}

// ApprovedAggregates returns a copy of the approved aggregates.
// For testing purposes only.
func (s *state) ApprovedAggregates() []xchain.AggAttestation {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Return a copy of the approved aggregates.
	aggs := make([]xchain.AggAttestation, len(s.approvedAggs))
	copy(aggs, s.approvedAggs)

	return aggs
}

// appHashJSON is the JSON representation of the state used to calculate app hash.
type appHashJSON struct {
	ExecutionStateRoot [32]byte                `json:"execution_state_root"`
	ValSetID           uint64                  `json:"validator_set_id"`
	Validators         []validator             `json:"validators"`
	PendingAggs        []xchain.AggAttestation `json:"pending_aggregate_attestations"`
	ApprovedAggs       []xchain.AggAttestation `json:"approved_aggregate_attestations"`
}
