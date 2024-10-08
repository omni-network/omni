package app

import (
	"encoding/json"
	"io"
	"sync"

	"github.com/omni-network/omni/lib/errors"
)

// readinessStatus contains all state related to halo readiness (health).
type readinessStatus struct {
	mu                 sync.RWMutex
	ConsensusSynced    bool    `json:"consensus_synced"`              // Unhealthy if false
	ConsensusP2PPeers  int     `json:"consensus_p2p_peers"`           // Unhealthy if 0
	ExecutionConnected bool    `json:"execution_connected"`           // Unhealthy if false
	ExecutionSynced    bool    `json:"execution_synced"`              // Unhealthy if false
	ExecutionP2PPeers  *uint64 `json:"execution_p2p_peers,omitempty"` // Unhealthy if 0
}

// ready indicates if the current status is considered as healthy.
// If the current halo instance is using an EVM RPC endpoint,
// the peers count is not taken into the account.
func (r *readinessStatus) ready() bool {
	r.mu.RLock()
	defer r.mu.RUnlock()

	ready := r.ConsensusSynced && r.ConsensusP2PPeers > 0 && r.ExecutionConnected &&
		r.ExecutionSynced

	peers := r.ExecutionP2PPeers
	if peers != nil {
		return ready && *peers > 0
	}

	return ready
}

// serialize marshals the struct to the specified writer.
func (r *readinessStatus) serialize(w io.Writer) error {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if err := json.NewEncoder(w).Encode(r); err != nil {
		return errors.Wrap(err, "serialization failed")
	}

	return nil
}

// setConsensusSynced sets the flag indicating the consensus is synced.
func (r *readinessStatus) setConsensusSynced(value bool) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.ConsensusSynced = value
}

// setConsensusP2PPeers sets the number of consensus peers.
func (r *readinessStatus) setConsensusP2PPeers(value int) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.ConsensusP2PPeers = value
}

// setExecutionConnected sets the flag indicating the execution engine is connected.
func (r *readinessStatus) setExecutionConnected(value bool) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.ExecutionConnected = value
}

// setExecutionSynced sets the flag indicating the execution engine is synced.
func (r *readinessStatus) setExecutionSynced(value bool) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.ExecutionSynced = value
}

// setExecutionP2PPeers sets the number of execution engine peers.
func (r *readinessStatus) setExecutionP2PPeers(value uint64) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.ExecutionP2PPeers = &value
}
