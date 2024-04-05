package resolvers_test

import (
	"context"
	"testing"

	"github.com/omni-network/omni/explorer/db"
	"github.com/omni-network/omni/explorer/db/ent"
	"github.com/omni-network/omni/explorer/graphql/app"
	d "github.com/omni-network/omni/explorer/graphql/data"
	"github.com/omni-network/omni/explorer/graphql/resolvers"

	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/gqltesting"
)

type gqlTest struct {
	Client   *ent.Client
	Opts     []graphql.SchemaOpt
	Provider *d.Provider
	Resolver resolvers.BlocksResolver
}

func createGqlTest(t *testing.T) *gqlTest {
	t.Helper()
	client := db.CreateTestEntClient(t)
	provider := &d.Provider{
		EntClient: client,
	}
	br := resolvers.BlocksResolver{
		BlocksProvider: provider,
	}

	opts := []graphql.SchemaOpt{
		graphql.UseFieldResolvers(),
		graphql.UseStringDescriptions(),
	}

	return &gqlTest{
		Client:   client,
		Provider: provider,
		Resolver: br,
		Opts:     opts,
	}
}

func TestXBlockQuery(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	test := createGqlTest(t)
	t.Cleanup(func() {
		if err := test.Client.Close(); err != nil {
			t.Error(err)
		}
	})
	block := db.CreateTestBlock(t, ctx, test.Client, 0)
	db.CreateXMsg(t, ctx, test.Client, block, 2, 0)
	db.CreateReceipt(t, ctx, test.Client, block, 2, 0)

	gqltesting.RunTests(t, []*gqltesting.Test{
		{
			Context: ctx,
			Schema:  graphql.MustParseSchema(app.Schema, &resolvers.Query{BlocksResolver: test.Resolver}, test.Opts...),
			Query: `
				{
					xblock(sourceChainID: 1, height: 0){
						SourceChainID
						BlockHeight
						BlockHash
						Messages{
							SourceChainID
						}
						Receipts{
							SourceChainID
						}
					}
				}
			`,
			ExpectedResult: `
				{
					"xblock":
					{
						"BlockHash":"0x0000000000000000000000000103176f1b2d62675e370103176f1b2d62675e37",
						"BlockHeight":"0x0",
						"SourceChainID":"0x1",
						"Messages":[
							{
								"SourceChainID":"0x1"
							}
						],
						"Receipts":[
							{
								"SourceChainID":"0x1"
							}
						]
					}
				}
			`,
		},
	})
}

func TestXBlocksQuery(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	test := createGqlTest(t)
	t.Cleanup(func() {
		if err := test.Client.Close(); err != nil {
			t.Error(err)
		}
	})
	db.CreateTestBlocks(t, ctx, test.Client, 2)

	gqltesting.RunTests(t, []*gqltesting.Test{
		{
			Context: ctx,
			Schema:  graphql.MustParseSchema(app.Schema, &resolvers.Query{BlocksResolver: test.Resolver}, test.Opts...),
			Query: `
				{
					xblockrange(from: 0, to: 2){
						SourceChainID
						BlockHeight
						BlockHash
					}
				}
			`,
			ExpectedResult: `
				{
					"xblockrange":[
					{
						"BlockHash":"0x0000000000000000000000000103176f1b2d62675e370103176f1b2d62675e37",
						"BlockHeight":"0x0",
						"SourceChainID":
						"0x1"
					},
					{
						"BlockHash":"0x0000000000000000000000000103176f1b2d62675e370103176f1b2d62675e37",
						"BlockHeight":"0x1",
						"SourceChainID":
						"0x1"
					}]
				}
			`,
		},
	})
}

func TestXBlocksCount(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	test := createGqlTest(t)
	t.Cleanup(func() {
		if err := test.Client.Close(); err != nil {
			t.Error(err)
		}
	})
	db.CreateTestBlocks(t, ctx, test.Client, 2)

	gqltesting.RunTests(t, []*gqltesting.Test{
		{
			Context: ctx,
			Schema:  graphql.MustParseSchema(app.Schema, &resolvers.Query{BlocksResolver: test.Resolver}, test.Opts...),
			Query: `
				{
					xblockcount
				}
			`,
			ExpectedResult: `
				{
					"xblockcount": "0x2"
				}
			`,
		},
	})
}
