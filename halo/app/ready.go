package app

import (
	"encoding/json"
	"io"
	"sync"

	"github.com/omni-network/omni/lib/errors"
)

// Consist of all flags describing the current readiness status of a halo node.
type ReadyResponse struct {
	mu                 sync.RWMutex
	ConsensusSynced    bool   `json:"consensus_synced"`      // Unhealthy if false
	ConsensusP2PPeers  int    `json:"consensus_p_2_p_peers"` // Unhealthy if 0
	ExecutionConnected bool   `json:"execution_connected"`   // Unhealthy if false
	ExecutionSynced    bool   `json:"execution_synced"`      // Unhealthy if false
	ExecutionP2PPeers  uint64 `json:"execution_p_2_p_peers"` // Unhealthy if 0
}

// Indicates if the current status is considered as healthy.
func (r *ReadyResponse) Healthy() bool {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return r.ConsensusSynced && r.ConsensusP2PPeers > 0 && r.ExecutionConnected &&
		r.ExecutionSynced && r.ExecutionP2PPeers > 0
}

// Marshals the strucxt to the specified writer.
func (r *ReadyResponse) Serialize(w io.Writer) error {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if err := json.NewEncoder(w).Encode(r); err != nil {
		return errors.Wrap(err, "serialization failed")
	}

	return nil
}

// Sets flag indicating the consensus is synced.
func (r *ReadyResponse) SetConsensusSynced(value bool) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.ConsensusSynced = value
}

// Sets number of consensus peers.
func (r *ReadyResponse) SetConsensusP2PPeers(value int) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.ConsensusP2PPeers = value
}

// Sets flag indicating the execution engine is connected.
func (r *ReadyResponse) SetExecutionConnected(value bool) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.ExecutionConnected = value
}

// Sets flag indicating the execution engine is synced.
func (r *ReadyResponse) SetExecutionSynced(value bool) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.ExecutionSynced = value
}

// Sets number of execution engine peers.
func (r *ReadyResponse) SetExecutionP2PPeers(value uint64) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.ExecutionP2PPeers = value
}
