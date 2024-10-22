package relayer

import (
	"context"
	"github.com/omni-network/omni/lib/cchain"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/umath"
	"github.com/omni-network/omni/lib/xchain"
	"sync"
	"time"
)

const attestBufLimit = 1000

func newAttestBuffer() *attestBuffer {
	return &attestBuffer{
		streamers: make(map[uint64]bool),
	}
}

type attestBuffer struct {
	mu        sync.RWMutex
	streamers map[uint64]bool
	atts      []*xchain.Attestation
}

func (b *attestBuffer) startOffset() uint64 {
	b.mu.RLock()
	defer b.mu.RUnlock()

	if len(b.atts) == 0 {
		return 0
	}

	return b.atts[0].AttestOffset
}

func (b *attestBuffer) endOffset() uint64 {
	b.mu.RLock()
	defer b.mu.RUnlock()

	if len(b.atts) == 0 {
		return 0
	}

	return b.atts[len(b.atts)-1].AttestOffset
}

func (b *attestBuffer) Add(att xchain.Attestation) bool {
	b.mu.Lock()
	defer b.mu.Unlock()
	defer func() {
		// Maybe trim if we grew over limit
		if len(b.atts) > attestBufLimit {
			b.atts = b.atts[len(b.atts)-attestBufLimit:]
		}
	}()

	start := b.startOffset()
	end := b.endOffset()

	if umath.SubtractOrZero(end, att.AttestOffset) > attestBufLimit {
		// Don't add atts that are too far behind, just ignore.
		return true
	}

	// If this is the first att, just add it.
	if len(b.atts) == 0 {
		b.atts = append(b.atts, &att)

		return true
	}

	if att.AttestOffset < start {
		// Add to the front.
		growLen := int(start) - int(att.AttestOffset)
		grow := make([]*xchain.Attestation, growLen)
		grow[0] = &att
		b.atts = append(grow, b.atts...)

		return true
	} else if att.AttestOffset > end {
		// Add to the back.
		growLen := int(att.AttestOffset) - int(end)
		grow := make([]*xchain.Attestation, growLen)
		grow[len(grow)-1] = &att
		b.atts = append(b.atts, grow...)

		return true
	}

	// Add to middle if not already present
	index := int(att.AttestOffset) - int(start)
	if b.atts[index] != nil {
		return false
	}

	b.atts[index] = &att

	return true
}

func (b *attestBuffer) ShouldStreamFrom(offset uint64) bool {
	b.mu.Lock()
	defer b.mu.Unlock()

	if b.streamers[offset] {
		return false
	}

	b.streamers[offset] = true

	return true
}

func (b *attestBuffer) Get(offset uint64) (xchain.Attestation, bool, error) {
	b.mu.RLock()
	defer b.mu.RUnlock()

	index, ok := umath.Subtract(offset, b.startOffset())
	if !ok {
		return xchain.Attestation{}, false, errors.New("attestation already trimmed")
	}

	if index >= uint64(len(b.atts)) {
		return xchain.Attestation{}, false, nil
	}

	return *b.atts[index], true, nil
}

type attestStreamer func(ctx context.Context, chainVer xchain.ChainVersion, attestOffset uint64, workerName string, callback cchain.ProviderCallback) error

func newAttestStreamer(cprov cchain.Provider) attestStreamer {
	var buffers sync.Map // map[xchain.ChainVer]*attestBuffer

	return func(ctx context.Context, chainVer xchain.ChainVersion, attestOffset uint64, workerName string, callback cchain.ProviderCallback) error {
		anyBuffer, _ := buffers.LoadOrStore(chainVer, newAttestBuffer())
		buffer := anyBuffer.(*attestBuffer)

		prevOffset := attestOffset
		streamCtx, cancel := context.WithCancel(ctx)
		if buffer.ShouldStreamFrom(attestOffset) {
			log.Debug(streamCtx, "Starting attest stream", "chain", chainVer, "offset", attestOffset, "worker", workerName)
			cprov.StreamAsync(streamCtx, chainVer, attestOffset, workerName, func(ctx context.Context, att xchain.Attestation) error {
				if !buffer.Add(att) {
					// Another stream already cached this att, so switch to reading from cache
					cancel()
					log.Debug(streamCtx, "Stopping overlapping attest stream", "chain", chainVer, "offset", att.AttestOffset, "worker", workerName)
					return ctx.Err()
				}

				prevOffset = att.AttestOffset

				return callback(ctx, att)
			})
		} else {
			cancel()
		}

		<-streamCtx.Done()
		log.Debug(ctx, "Streaming from attest buffer", "chain", chainVer, "prev_offset", prevOffset, "worker", workerName)
		for ctx.Err() == nil {
			att, ok, err := buffer.Get(prevOffset + 1)
			if err != nil {
				return err
			} else if !ok {
				time.Sleep(time.Second)
				continue
			}

			if err := callback(ctx, att); err != nil {
				return err
			}

			prevOffset++
		}

		return ctx.Err()
	}
}
