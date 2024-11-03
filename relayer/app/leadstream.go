package relayer

import (
	"context"
	"sync"
	"time"

	"github.com/omni-network/omni/lib/cchain"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/umath"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/tidwall/btree"
)

// errStop is returned when the leader should stop streaming.
var errStop = errors.New("stop")

// getLimit returns the limit of attestations to cache per chain version.
func getLimit(network netconf.ID) int {
	if network == netconf.Devnet {
		return 20 // Tiny buffer in devnet
	}

	return 10_000 // 10k atts & 1KB per attestation & 10 chain versions ~= 100MB
}

// getBackoff returns the duration to backoff before querying the cache again.
func getBackoff(network netconf.ID) time.Duration {
	if network == netconf.Simnet {
		return time.Millisecond // No backoff in tests
	}

	return time.Second // Default 1 second blocks otherwise
}

// leadChaosTimeout returns a function that returns true if the leader chaos timeout has been reached.
// This ensures we rotate leaders after a certain time (and test leader rotation).
func leadChaosTimeout(network netconf.ID) func() bool {
	t0 := time.Now()
	return func() bool {
		duration := time.Hour // Default 1 hour timeout
		if network == netconf.Devnet {
			duration = time.Second * 10
		} else if network == netconf.Staging {
			duration = time.Minute
		}

		return time.Since(t0) > duration
	}
}

// leader tracks a worker actively streaming attestations and adding them to the cache.
// It "locks" a range of offsets as leader of that range.
type leader struct {
	mu     sync.RWMutex
	from   uint64
	latest uint64
	delete func()
}

// contains returns true if the offset is within the leader's range.
func (s *leader) contains(offset uint64) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.from <= offset && offset <= s.latest
}

// IncRange increases the range to the provided height plus one
// since the leader will move on the that height next.
func (s *leader) IncRange(height uint64) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.latest = height + 1
}

// Delete removes the leader from the buffer.
func (s *leader) Delete() {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.delete()
}

func newAttestBuffer(limit int) *attestBuffer {
	return &attestBuffer{
		limit:   limit,
		leaders: make(map[*leader]struct{}),
		atts:    btree.NewMap[uint64, xchain.Attestation](32), // Degree of 32 is a good default
	}
}

// attestBuffer is an attestations cache being populated by leaders.
// The goal is to avoid overlapping streams.
type attestBuffer struct {
	limit   int
	mu      sync.RWMutex
	leaders map[*leader]struct{}
	atts    *btree.Map[uint64, xchain.Attestation]
}

// Add adds an attestation to the cache if the cache is not full or
// if the attestation is not too old.
// It returns true if it was added and an existing key was replaced.
//
//nolint:nonamedreturns // Name for clarify of API.
func (b *attestBuffer) Add(att xchain.Attestation) (replaced bool) {
	b.mu.Lock()
	defer b.mu.Unlock()

	defer func() {
		// Maybe trim if we grew over limit
		for b.atts.Len() > b.limit {
			height, _, ok := b.atts.PopMin()
			if ok && height == att.AttestOffset {
				replaced = false // Not added, so not replaced
			}
		}
	}()

	_, replaced = b.atts.Set(att.AttestOffset, att)

	return replaced
}

// Get returns the attestation at the provided offset or false.
func (b *attestBuffer) Get(offset uint64) (xchain.Attestation, bool) {
	b.mu.RLock()
	defer b.mu.RUnlock()

	return b.atts.Get(offset)
}

// containsStreamerUnsafe returns true if any leader contains the offset.
// It is unsafe since it assumes the lock is held.
func (b *attestBuffer) containsStreamerUnsafe(offset uint64) bool {
	for s := range b.leaders {
		if s.contains(offset) {
			return true
		}
	}

	return false
}

// MaybeNewLeader returns a new leader if no other leader is already streaming the offset or false.
func (b *attestBuffer) MaybeNewLeader(offset uint64) (*leader, bool) {
	b.mu.RLock()
	contains := b.containsStreamerUnsafe(offset)
	b.mu.RUnlock()
	if contains {
		return nil, false
	}

	b.mu.Lock()
	defer b.mu.Unlock()

	// Double check in case another goroutine added a leader
	for s := range b.leaders {
		if s.contains(offset) {
			return nil, false
		}
	}

	s := &leader{
		from:   offset,
		latest: offset,
	}
	s.delete = func() {
		b.mu.Lock()
		defer b.mu.Unlock()

		delete(b.leaders, s)
	}

	b.leaders[s] = struct{}{}

	return s, true
}

// attestStreamer is a function that streams attestations from a specific offset.
// It abstracts cchain.Provider.StreamAttestations.
type attestStreamer func(ctx context.Context, chainVer xchain.ChainVersion, attestOffset uint64, workerName string, callback cchain.ProviderCallback) error

// newLeaderStreamer returns a new attestStreamer that avoids multiple overlapping streaming queries
// by selecting a leader to query each range of offsets.
func newLeaderStreamer(upstream attestStreamer, network netconf.ID) attestStreamer {
	var buffers sync.Map // map[xchain.ChainVer]*attestBuffer

	return func(ctx context.Context, chainVer xchain.ChainVersion, fromOffset uint64, workerName string, callback cchain.ProviderCallback) error {
		anyBuffer, _ := buffers.LoadOrStore(chainVer, newAttestBuffer(getLimit(network)))
		buffer := anyBuffer.(*attestBuffer) //nolint:revive,forcetypeassert // Type is known

		name := netconf.ChainVersionNamer(network)(chainVer)

		// Track the offset of the last attestation we "processed"
		prevOffset, ok := umath.Subtract(fromOffset, 1)
		if !ok {
			return errors.New("attest from offset zero [BUG]", "from", fromOffset)
		}

		// lead blocks and streams attestations from the provided height using the provided leader.
		// It populates the cache with fetched attestations.
		// It returns nil if it detects an overlap with another leader or on chaos timeout.
		lead := func(l *leader, from uint64) error {
			defer l.Delete()
			log.Debug(ctx, "Starting attest stream", "chain", name, "offset", from, "worker", workerName)
			timeout := leadChaosTimeout(network)

			err := upstream(ctx, chainVer, from, workerName, func(ctx context.Context, att xchain.Attestation) error {
				l.IncRange(att.AttestOffset)
				replaced := buffer.Add(att)

				if err := callback(ctx, att); err != nil {
					return err
				}

				prevOffset = att.AttestOffset

				if replaced {
					// Another leader already cached this att, so switch to reading from cache
					log.Debug(ctx, "Stopping overlapping attest stream", "chain", name, "offset", att.AttestOffset, "worker", workerName)
					return errStop
				} else if timeout() {
					log.Debug(ctx, "Stopping timed-out attest stream", "chain", name, "offset", att.AttestOffset, "worker", workerName)
					return errStop
				}

				return nil
			})
			if errors.Is(err, errStop) {
				return nil
			}

			return err
		}

		timer := time.NewTimer(0)
		defer timer.Stop()

		// Loop until the context is closed or error is encountered
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-timer.C:
				// Check if the cache is populated
				if att, ok := buffer.Get(prevOffset + 1); ok {
					// Got it, process the attestation
					if err := callback(ctx, att); err != nil {
						return err
					}

					prevOffset = att.AttestOffset
					timer.Reset(0) // Immediately go to next

					continue
				}

				// Cache isn't populated, check if we need to start streaming as a new leader
				if l, ok := buffer.MaybeNewLeader(prevOffset + 1); ok {
					if err := lead(l, prevOffset+1); err != nil {
						return err
					}

					// Leader stopped gracefully, so try cache immediately (with probably updated prevOffset)
					timer.Reset(0)

					continue
				}

				// Otherwise, wait a bit, and try the same offset again
				timer.Reset(getBackoff(network))
			}
		}
	}
}
