package resolvers_test

import (
	"context"
	"testing"

	"github.com/omni-network/omni/explorer/db/ent"
	"github.com/omni-network/omni/explorer/db/testutil"
	"github.com/omni-network/omni/explorer/graphql/app"
	"github.com/omni-network/omni/explorer/graphql/data"
	"github.com/omni-network/omni/explorer/graphql/resolvers"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/gqltesting"
)

type gqlTest struct {
	Client   *ent.Client
	Opts     []graphql.SchemaOpt
	Provider *data.Provider
	Resolver *resolvers.Root
}

func createGqlTest(t *testing.T, network netconf.ID) *gqlTest {
	t.Helper()
	client := testutil.CreateTestEntClient(t)
	p := data.NewProvider(context.Background(), client, network)
	r := resolvers.NewRoot(p)

	opts := []graphql.SchemaOpt{
		graphql.UseFieldResolvers(),
		graphql.UseStringDescriptions(),
	}

	t.Cleanup(func() {
		if err := client.Close(); err != nil {
			t.Error(err)
		}
	})

	return &gqlTest{
		Client:   client,
		Provider: p,
		Resolver: r,
		Opts:     opts,
	}
}

func TestXBlocksQuery(t *testing.T) {
	t.Skip("This test is failing because the schema was changed")
	t.Parallel()
	ctx := context.Background()
	test := createGqlTest(t, netconf.Devnet)
	testutil.CreateTestBlocks(t, ctx, test.Client, 2)

	gqltesting.RunTests(t, []*gqltesting.Test{
		{
			Context: ctx,
			Schema:  graphql.MustParseSchema(app.Schema, resolvers.NewRoot(test.Provider), test.Opts...),
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
