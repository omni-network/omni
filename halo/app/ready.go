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
	ConsensusHeight    int64   `json:"consensus_height"`              // Unhealthy if 0
	ConsensusRunning   bool    `json:"consensus_running"`             // Unhealthy if false
	ExecutionConnected bool    `json:"execution_connected"`           // Unhealthy if false
	ExecutionSynced    bool    `json:"execution_synced"`              // Unhealthy if false
	ExecutionP2PPeers  *uint64 `json:"execution_p2p_peers,omitempty"` // Unhealthy if 0 (ignored if not available)
	ExecutionHeight    uint64  `json:"execution_height"`              // Unhealthy if 0
}

// ready indicates if the current status is considered as healthy.
func (r *readinessStatus) ready() bool {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return r.readyUnsafe()
}

// readyUnsafe is the same as ready but without locking.
// It is unsafe since it assumes the lock is held.
func (r *readinessStatus) readyUnsafe() bool {
	return r.ConsensusSynced &&
		r.ConsensusP2PPeers > 0 &&
		r.ConsensusHeight > 0 &&
		r.ConsensusRunning &&
		r.ExecutionConnected &&
		r.ExecutionSynced &&
		r.ExecutionHeight > 0 &&
		(r.ExecutionP2PPeers == nil || *r.ExecutionP2PPeers > 0)
}

// serialize marshals the struct to the specified writer.
// It also returns the ready status.
func (r *readinessStatus) serialize(w io.Writer) (bool, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if err := json.NewEncoder(w).Encode(r); err != nil {
		return false, errors.Wrap(err, "readiness status serialization")
	}

	return r.readyUnsafe(), nil
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

// setExecutionHeight sets the latest block height of the execution engine.
func (r *readinessStatus) setExecutionHeight(value uint64) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.ExecutionHeight = value
}

// setConsensusHeight sets the latest block height of the consensus engine.
func (r *readinessStatus) setConsensusHeight(value int64) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.ConsensusHeight = value
}

// setConsensusRunning sets the flag indicating the consensus engine is running.
func (r *readinessStatus) setConsensusRunning(value bool) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.ConsensusRunning = value
}
