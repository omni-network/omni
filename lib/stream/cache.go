package stream

import (
	"math"
	"sync"
)

// NewNopCache returns a cache that does nothing.
func NewNopCache[E any]() Cache[E] {
	return new(nopCache[E])
}

type nopCache[E any] struct{}

func (nopCache[E]) Get(uint64) []E  { return nil }
func (nopCache[E]) Set(uint64, []E) {}

// NewCache returns a new map based cache using the provided limit.
func NewCache[E any](limit int) Cache[E] {
	return &mapCache[E]{
		limit: limit,
		elems: make(map[uint64]E),
		floor: math.MaxUint64,
	}
}

// mapCache implements Cache storing elements in a map.
// It will only retain the highest `limit` elements.
// It assumes elements are roughly sequential and that heights are not sparse.
type mapCache[E any] struct {
	limit int // Immutable

	// Mutable fields
	mu    sync.RWMutex
	elems map[uint64]E
	floor uint64 // floor is equaled to or lower than the lowest height in elems. It is MaxUint64 if elems is empty.
}

// Get returns all strictly sequential elements in the cache from the provided height (inclusive).
func (c *mapCache[E]) Get(from uint64) []E {
	c.mu.RLock()
	defer c.mu.RUnlock()

	var resp []E
	for height := from; ; height++ {
		elem, ok := c.elems[height]
		if !ok {
			break
		}

		resp = append(resp, elem)
	}

	return resp
}

// Set adds the provided elements to the cache starting from the provided height (inclusive).
// Elements are assumed to be in strictly sequential order.
func (c *mapCache[E]) Set(from uint64, elems []E) {
	c.mu.Lock()
	defer c.mu.Unlock()

	for i, elem := range elems {
		height := from + uint64(i)

		if len(c.elems) >= c.limit && height < c.floor {
			// Don't add too low elements when cache is full, they will just be deleted below.
			continue
		}

		c.elems[height] = elem

		if height < c.floor {
			c.floor = height // Update floor
		}
	}

	for len(c.elems) > c.limit {
		delete(c.elems, c.floor)
		c.floor++
	}
}
