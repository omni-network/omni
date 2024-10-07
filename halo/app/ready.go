package app

import (
	"encoding/json"
	"io"
	"sync"

	"github.com/omni-network/omni/lib/errors"
)

// Consist of all flags describing the current readiness status of a halo node.
type readinessStatus struct {
	mu     sync.RWMutex
	status flags
}

type flags struct {
	ConsensusSynced    bool    `json:"consensus_synced"`                // Unhealthy if false
	ConsensusP2PPeers  int     `json:"consensus_p_2_p_peers"`           // Unhealthy if 0
	ExecutionConnected bool    `json:"execution_connected"`             // Unhealthy if false
	ExecutionSynced    bool    `json:"execution_synced"`                // Unhealthy if false
	ExecutionP2PPeers  *uint64 `json:"execution_p_2_p_peers,omitempty"` // Unhealthy if 0
}

// Indicates if the current status is considered as healthy.
// If the current halo instance is using an EVM RPC endpoint,
// the peers count is not taken into the account.
func (r *readinessStatus) ready() bool {
	r.mu.RLock()
	defer r.mu.RUnlock()

	ready := r.status.ConsensusSynced && r.status.ConsensusP2PPeers > 0 && r.status.ExecutionConnected &&
		r.status.ExecutionSynced

	peers := r.status.ExecutionP2PPeers
	if peers != nil {
		return ready && *peers > 0
	}

	return ready
}

// Marshals the struct to the specified writer.
func (r *readinessStatus) serialize(w io.Writer) error {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if err := json.NewEncoder(w).Encode(r.status); err != nil {
		return errors.Wrap(err, "serialization failed")
	}

	return nil
}

// Sets flag indicating the consensus is synced.
func (r *readinessStatus) setConsensusSynced(value bool) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.status.ConsensusSynced = value
}

// Sets number of consensus peers.
func (r *readinessStatus) setConsensusP2PPeers(value int) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.status.ConsensusP2PPeers = value
}

// Sets flag indicating the execution engine is connected.
func (r *readinessStatus) setExecutionConnected(value bool) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.status.ExecutionConnected = value
}

// Sets flag indicating the execution engine is synced.
func (r *readinessStatus) setExecutionSynced(value bool) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.status.ExecutionSynced = value
}

// Sets number of execution engine peers.
func (r *readinessStatus) setExecutionP2PPeers(value uint64) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.status.ExecutionP2PPeers = &value
}
