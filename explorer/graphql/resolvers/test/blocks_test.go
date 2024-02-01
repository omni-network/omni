package resolvers_test

import (
	"context"
	"testing"
	"time"

	"github.com/omni-network/omni/explorer/graphql/app"
	"github.com/omni-network/omni/explorer/graphql/resolvers"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/gqltesting"
	"go.uber.org/mock/gomock"
)

func TestXBlockQuery(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	sourceChainID := hexutil.Big(*hexutil.MustDecodeBig("0x4b2"))
	blockHeight := hexutil.Big(*hexutil.MustDecodeBig("0x1"))
	blockHashBytes := []byte{1, 3, 23, 111, 27, 45, 98, 103, 94, 55, 1, 3, 23, 111, 27, 45, 98, 103, 94, 55}
	blockHash := common.Hash{}
	blockHash.SetBytes(blockHashBytes)

	opts := []graphql.SchemaOpt{
		graphql.UseFieldResolvers(),
		graphql.UseStringDescriptions(),
	}

	mockXBlock := &resolvers.XBlock{
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

	gqltesting.RunTests(t, []*gqltesting.Test{
		{
			Context: ctx,
			Schema:  graphql.MustParseSchema(app.Schema, &resolvers.Query{BlocksResolver: br}, opts...),
			Query: `
				{
					xblock(sourceChainID: 1234, height: 0){
						SourceChainID
						BlockHeight
						BlockHash
					}
				}
			`,
			ExpectedResult: `
				{
					"xblock":
					{
						"BlockHash":"0x0000000000000000000000000103176f1b2d62675e370103176f1b2d62675e37",
						"BlockHeight":"0x1",
						"SourceChainID":"0x4b2"
					}
				}
			`,
		},
	})
}
