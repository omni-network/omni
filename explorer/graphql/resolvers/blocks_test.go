package resolvers_test

import (
	"context"
	"testing"

	"github.com/omni-network/omni/explorer/graphql/app"
	d "github.com/omni-network/omni/explorer/graphql/data"
	"github.com/omni-network/omni/explorer/graphql/resolvers"

	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/gqltesting"
)

func TestXBlockQuery(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	client := resolvers.CreateTestEntClient(t)
	resolvers.CreateTestBlock(t, ctx, client)
	defer client.Close()

	opts := []graphql.SchemaOpt{
		graphql.UseFieldResolvers(),
		graphql.UseStringDescriptions(),
	}

	provider := &d.Provider{
		EntClient: client,
	}

	br := resolvers.BlocksResolver{
		BlocksProvider: provider,
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
						"BlockHeight":"0x0",
						"SourceChainID":"0x4d2"
					}
				}
			`,
		},
	})
}
