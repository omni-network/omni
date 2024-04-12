package resolvers_test

import (
	"context"
	"testing"

	"github.com/omni-network/omni/explorer/db"
	"github.com/omni-network/omni/explorer/graphql/app"
	"github.com/omni-network/omni/explorer/graphql/resolvers"

	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/gqltesting"
)

func TestSupportedChains(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	test := createGqlTest(t)
	t.Cleanup(func() {
		if err := test.Client.Close(); err != nil {
			t.Error(err)
		}
	})
	db.CreateTestChain(t, ctx, test.Client, 1)
	db.CreateTestChain(t, ctx, test.Client, 2)

	gqltesting.RunTests(t, []*gqltesting.Test{
		{
			Context: ctx,
			Schema:  graphql.MustParseSchema(app.Schema, &resolvers.Query{BlocksResolver: test.Resolver}, test.Opts...),
			Query: `
				{
					supportedchains{
						ChainID
						Name
					}
				}
			`,
			ExpectedResult: `
				{
					"supportedchains":[
						{
							"ChainID":"0x1",
							"Name":"test-chain1"
						},
						{
							"ChainID":"0x2",
							"Name":"test-chain2"
						}
					]
				}
			`,
		},
	})
}
