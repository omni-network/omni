package stream

import (
	"strconv"
	"time"

	"github.com/omni-network/omni/lib/errors"

	"github.com/maypok86/otter"
)

// Cacher defines the cache type for caching stream items.
// Stream uses height as stream cursor so our cache implementation
// relies on the height for calculating the key.
//
// The key implementation can vary based on the context and is up
// to the implementor.
type Cacher[K any, E any] interface {
	// Get cached item at the height (stream cursor)
	// as well as a boolean indicating a cache hit or miss.
	// If item is not found nil and false is returned.
	Get(height uint64) ([]E, bool)
	// Set item in the cache at the height and return true.
	// If item is failed to store, returns false.
	Set(height uint64, value []E) bool
}

const cacheSize = 100_000

type Opt[E any] func(*Cache[E])

// WithKeyFunc sets a custom key function for the Cache.
func WithKeyFunc[E any](keyFunc func(uint64) string) Opt[E] {
	return func(c *Cache[E]) {
		c.key = keyFunc
	}
}

// WithTTL sets a custom key function for the Cache.
func WithTTL[E any](ttl time.Duration) Opt[E] {
	return func(c *Cache[E]) {
		c.ttl = ttl
	}
}

// Cache is a default cache implementation that uses
// otter Cache as an underlying cache. It assumes keys
// as strings.
// note: the key is string type for POC, in production we would probably want to do
// variable length encoding (e.g. VLQ or similar/simpler) due to the nature of the
// keys, for example chain ID is a small set of keys we support, and heights, although
// they can reach 64 bytes they are growing slowly so no need to keep the key big initially.
type Cache[E any] struct {
	cache *otter.Cache[string, []E]
	key   func(uint64) string
	ttl   time.Duration
}

func NewCache[E any](opts ...Opt[E]) (*Cache[E], error) {
	cache := &Cache[E]{
		// default implementation for key
		key: func(height uint64) string {
			return strconv.FormatUint(height, 10)
		},
	}

	for _, opt := range opts {
		opt(cache)
	}

	builder := otter.MustBuilder[string, []E](cacheSize)

	if cache.ttl != 0 {
		builder.WithTTL(cache.ttl)
	}

	otterCache, err := builder.Build()
	if err != nil {
		return nil, errors.Wrap(err, "cache build [BUG]", "size", cacheSize)
	}
	cache.cache = &otterCache

	return cache, nil
}

func (c *Cache[E]) Get(height uint64) ([]E, bool) {
	// todo report hit/miss metrics
	return c.cache.Get(c.key(height))
}

func (c *Cache[E]) Set(height uint64, value []E) bool {
	return c.cache.Set(c.key(height), value)
}
