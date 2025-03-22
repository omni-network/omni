package ethclient

import (
	"context"
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

	h1 := &types.Header{Number: bi.N(1)}
	h2 := &types.Header{Number: bi.N(2), ParentHash: h1.Hash()}
	h3 := &types.Header{Number: bi.N(3), ParentHash: h2.Hash()}

	testCl := &testClient{headers: []*types.Header{h1, h2, h3}}
	cache := newHeaderCache(testCl)
	cache.limit = 2

	// Fetch h1 by hash, ensure queried
	h, err := cache.HeaderByHash(ctx, h1.Hash())
	require.NoError(t, err)
	require.Equal(t, h1, h)
	require.Equal(t, 1, testCl.headerByHash)

	// Fetch h1 by hash again, ensure cached
	h, err = cache.HeaderByHash(ctx, h1.Hash())
	require.NoError(t, err)
	require.Equal(t, h1, h)
	require.Equal(t, 1, testCl.headerByHash)

	// Fetch h1 by number, ensure queried
	h, err = cache.HeaderByNumber(ctx, h1.Number)
	require.NoError(t, err)
	require.Equal(t, h1, h)
	require.Equal(t, 1, testCl.headerByNumber)

	// Fetch h2 by number, ensure queried
	h, err = cache.HeaderByNumber(ctx, h2.Number)
	require.NoError(t, err)
	require.Equal(t, h2, h)
	require.Equal(t, 2, testCl.headerByNumber)

	// Fetch h2 by hash, ensure cached
	h, err = cache.HeaderByHash(ctx, h2.Hash())
	require.NoError(t, err)
	require.Equal(t, h2, h)
	require.Equal(t, 1, testCl.headerByHash)

	// Fetch h3 by type, ensure queried
	h, err = cache.HeaderByType(ctx, HeadLatest)
	require.NoError(t, err)
	require.Equal(t, h3, h)
	require.Equal(t, 1, testCl.headerByType)

	// Fetch h3 by type again, ensure queried
	h, err = cache.HeaderByType(ctx, HeadLatest)
	require.NoError(t, err)
	require.Equal(t, h3, h)
	require.Equal(t, 2, testCl.headerByType)

	// Fetch h3 by number, ensure queried
	h, err = cache.HeaderByNumber(ctx, h3.Number)
	require.NoError(t, err)
	require.Equal(t, h3, h)
	require.Equal(t, 3, testCl.headerByNumber)

	// Fetch h3 by hash, ensure cached
	h, err = cache.HeaderByHash(ctx, h3.Hash())
	require.NoError(t, err)
	require.Equal(t, h3, h)
	require.Equal(t, 2, testCl.headerByHash)

	// Ensure h1 pruned, so queried when fetched by hash
	h, err = cache.HeaderByHash(ctx, h1.Hash())
	require.NoError(t, err)
	require.Equal(t, h1, h)
	require.Equal(t, 3, testCl.headerByHash)
}

type testClient struct {
	Client
	headerByHash   int
	headerByType   int
	headerByNumber int
	headers        []*types.Header
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
