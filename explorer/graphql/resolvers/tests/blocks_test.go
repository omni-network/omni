package resolvers_tests

import (
	"testing"
	"time"

	gql "github.com/omni-network/omni/explorer/graphql/app"
	resolvers "github.com/omni-network/omni/explorer/graphql/resolvers"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/gqltesting"
	"go.uber.org/mock/gomock"
)

func TestXBlockQuery(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	sourceChainID := hexutil.Big(*hexutil.MustDecodeBig("0x4b2"))
	blockHeight := hexutil.Big(*hexutil.MustDecodeBig("0x1"))
	blockHash := common.Hash{}
	blockHash.SetBytes([]byte{1, 3, 23, 111, 27, 45, 98, 103, 94, 55, 1, 3, 23, 111, 27, 45, 98, 103, 94, 55})

	mockXBlock := &resolvers.XBlock{
		UUID:          graphql.ID("0x1"),
		SourceChainID: sourceChainID,
		BlockHeight:   blockHeight,
		BlockHash:     blockHash,
		Timestamp:     graphql.Time{Time: time.Now()},
		Messages:      nil,
	}
	mockProvider := NewMockBlocksProvider(ctrl)
	mockProvider.EXPECT().XBlock(uint64(1234), uint64(0)).Return(mockXBlock, true, nil)

	br := resolvers.BlocksResolver{
		BlocksProvider: mockProvider,
	}

	r := &resolvers.Query{
		BlocksResolver: br,
	}

	schema := graphql.MustParseSchema(gql.Schema, r)

	gqltesting.RunTests(t, []*gqltesting.Test{
		{
			Schema: schema,
			Query: `
				{
					xblock(sourceChainID: 1234, height: 0){
						UUID
						SourceChainID
						BlockHeight
						BlockHash
					}
			}
			`,
			ExpectedResult: `
				{
					"UUID": "0x1",
					"SourceChainId": "0x4bf",
					"BlockHeight": "0x1",
					"BlockHash": "0x0103176f1b2d62675e370103176f1b2d62675e37",
				}
			`,
		},
	})
}
