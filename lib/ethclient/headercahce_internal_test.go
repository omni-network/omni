package ethclient

import (
	"context"
	"encoding/json"
	"math/big"
	"testing"

	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/stretchr/testify/require"
)

func TestHeaderCache(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	fuzzHeader := func(height int, parent *types.Header) *types.Header {
		var resp types.Header
		NewFuzzer(0).Fuzz(&resp)
		resp.Number = bi.N(height)
		if parent != nil {
			resp.ParentHash = parent.Hash()
		}

		return &resp
	}

	h1 := fuzzHeader(1, nil)
	h2 := fuzzHeader(2, h1)
	h3 := fuzzHeader(3, h2)

	testCl := &testClient{headers: []*types.Header{h1, h2, h3}}
	cache, err := newHeaderCache(testCl)
	require.NoError(t, err)
	cache.limit = 2

	// Fetch h1 by hash, ensure queried
	h, err := cache.HeaderByHash(ctx, h1.Hash())
	require.NoError(t, err)
	headersEqual(t, h1, h)
	require.Equal(t, 1, testCl.headerByHash)

	// Fetch h1 by hash again, ensure cached
	h, err = cache.HeaderByHash(ctx, h1.Hash())
	require.NoError(t, err)
	headersEqual(t, h1, h)
	require.Equal(t, 1, testCl.headerByHash)

	// Fetch h1 by number, ensure cached
	h, err = cache.HeaderByNumber(ctx, h1.Number)
	require.NoError(t, err)
	headersEqual(t, h1, h)
	require.Equal(t, 0, testCl.headerByNumber)

	// Fetch h2 by number, ensure queried
	h, err = cache.HeaderByNumber(ctx, h2.Number)
	require.NoError(t, err)
	headersEqual(t, h2, h)
	require.Equal(t, 1, testCl.headerByNumber)

	// Fetch h2 by hash, ensure cached
	h, err = cache.HeaderByHash(ctx, h2.Hash())
	require.NoError(t, err)
	headersEqual(t, h2, h)
	require.Equal(t, 1, testCl.headerByHash)

	// Fetch h3 by type, ensure queried
	h, err = cache.HeaderByType(ctx, HeadLatest)
	require.NoError(t, err)
	headersEqual(t, h3, h)
	require.Equal(t, 1, testCl.headerByType)

	// Fetch h3 by type again, ensure queried
	h, err = cache.HeaderByType(ctx, HeadLatest)
	require.NoError(t, err)
	headersEqual(t, h3, h)
	require.Equal(t, 2, testCl.headerByType)

	// Fetch h3 by number, ensure cached
	h, err = cache.HeaderByNumber(ctx, h3.Number)
	require.NoError(t, err)
	headersEqual(t, h3, h)
	require.Equal(t, 1, testCl.headerByNumber)

	// Fetch h3 by hash, ensure cached
	h, err = cache.HeaderByHash(ctx, h3.Hash())
	require.NoError(t, err)
	headersEqual(t, h3, h)
	require.Equal(t, 1, testCl.headerByHash)

	// Ensure h1 pruned, so queried when fetched by hash
	h, err = cache.HeaderByHash(ctx, h1.Hash())
	require.NoError(t, err)
	headersEqual(t, h1, h)
	require.Equal(t, 2, testCl.headerByHash)
}

type testClient struct {
	Client
	headerByHash   int
	headerByType   int
	headerByNumber int
	headers        []*types.Header
}

func (c *testClient) Name() string {
	return "test"
}

func (c *testClient) HeaderByType(_ context.Context, typ HeadType) (*types.Header, error) {
	if typ != HeadLatest {
		return nil, errors.New("invalid head type")
	}

	c.headerByType++

	return c.headers[len(c.headers)-1], nil
}

func (c *testClient) HeaderByNumber(_ context.Context, num *big.Int) (*types.Header, error) {
	c.headerByNumber++
	for _, h := range c.headers {
		if bi.EQ(h.Number, num) {
			return h, nil
		}
	}

	return nil, errors.New("invalid number")
}

func (c *testClient) HeaderByHash(_ context.Context, hash common.Hash) (*types.Header, error) {
	c.headerByHash++
	for _, h := range c.headers {
		if h.Hash() == hash {
			return h, nil
		}
	}

	return nil, errors.New("invalid hash")
}

func headersEqual(t *testing.T, h1, h2 *types.Header) {
	t.Helper()
	bz1, err := json.MarshalIndent(h1, "", " ")
	require.NoError(t, err)
	bz2, err := json.MarshalIndent(h2, "", " ")
	require.NoError(t, err)
	require.Equal(t, string(bz1), string(bz2))
}
