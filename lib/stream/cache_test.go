package stream_test

import (
	"math/rand/v2"
	"sync"
	"testing"

	"github.com/omni-network/omni/lib/stream"

	"github.com/stretchr/testify/require"
)

//go:generate go test . -run=TestFloor -count=100

func TestFloor(t *testing.T) {
	t.Parallel()

	const total = 100
	toAdd := make(map[int]struct{})
	for i := range total {
		toAdd[i] = struct{}{}
	}

	cache := stream.NewCache[int](total)
	// Add in random order
	for i := range toAdd {
		cache.Set(uint64(i), []int{i})
		stream.EnsureFloor(t, cache)
	}
	// Add in random order
	for i := range toAdd {
		i += total
		cache.Set(uint64(i), []int{i})
		stream.EnsureFloor(t, cache)
	}
}

//go:generate go test . -run=TestCache -count=100

func TestCache(t *testing.T) {
	t.Parallel()

	const limit = 100
	cache := stream.NewCache[int](limit)

	// Helper function ensuring length of returned elements
	requireGetLen := func(t *testing.T, from uint64, l int) {
		t.Helper()
		require.Len(t, cache.Get(from), l)
	}

	var wg sync.WaitGroup
	added := make(chan []int)

	// Ensure it is empty
	requireGetLen(t, 0, 0)
	requireGetLen(t, 1, 0)
	requireGetLen(t, limit, 0)

	// Concurrently add elements to the cache
	const writers = 10
	for i := range writers {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			count := limit / writers
			start := count * i

			var elems []int
			for j := 0; j < count; j++ {
				elems = append(elems, start+j)
			}

			cache.Set(uint64(start), elems)
			added <- elems
		}(i)
	}

	// Wait for all elements to be added
	go func() {
		wg.Wait()
		close(added)
	}()

	for elems := range added {
		// Ensure these are all added
		for i, elem := range elems {
			got := cache.Get(uint64(elem))
			// Ensure they are in order
			for j, e := range got {
				if j == 0 {
					continue
				}
				require.Equal(t, e, got[j-1]+1)
			}

			for _, expect := range elems[i:] {
				require.Contains(t, got, expect)
			}
		}
	}

	// Ensure the cache is full
	requireGetLen(t, 0, limit)
	cache.Set(limit, []int{limit, limit + 1}) // Add two more elements
	requireGetLen(t, 0, 0)                    // Ensure the oldest element is removed
	requireGetLen(t, 1, 0)                    // Ensure the second-oldest element is removed
	requireGetLen(t, 2, limit)                // Ensure the second-oldest element is removed

	// Add a new element with a gap at limit+2
	cache.Set(limit+3, []int{limit + 3})
	requireGetLen(t, 2, 0)       // Ensure the third-oldest element is removed
	requireGetLen(t, 3, limit-1) // The rest should still be there (before gap)
	requireGetLen(t, limit+3, 1) // The new element should be there (after gap
}

func Benchmark(b *testing.B) {
	const limit = 1000

	benchmarks := []struct {
		name  string
		cache stream.Cache[int]
	}{
		{
			name:  "nop",
			cache: stream.NewNopCache[int](),
		},
		//{
		//	name:  "slice",
		//	cache: stream.NewSliceCache[int](limit),
		// },
		{
			name:  "map",
			cache: stream.NewCache[int](limit),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			var wg sync.WaitGroup
			r := rand.New(rand.NewPCG(uint64(b.N), uint64(b.N)))

			// Concurrently run Get and Set operations
			for i := 0; i < b.N; i++ {
				wg.Add(2)

				// Perform Set operations concurrently
				go func() {
					defer wg.Done()
					height := r.Uint64N(limit)
					elems := make([]int, 10)
					for j := range elems {
						elems[j] = rand.IntN(100)
					}
					bm.cache.Set(height, elems)
				}()

				// Perform Get operations concurrently
				go func() {
					defer wg.Done()
					_ = bm.cache.Get(rand.Uint64N(limit))
				}()
			}
		})
	}
}
