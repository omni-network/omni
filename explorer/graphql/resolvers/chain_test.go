package resolvers_test

import (
	"context"
	"testing"

	"github.com/omni-network/omni/explorer/db/testutil"
	"github.com/omni-network/omni/explorer/graphql/app"

	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/gqltesting"
)

func TestSupportedChains(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	test := createGqlTest(t)
	testutil.CreateTestChain(t, ctx, test.Client, 1)
	testutil.CreateTestChain(t, ctx, test.Client, 2)

	gqltesting.RunTests(t, []*gqltesting.Test{
		{
			Context: ctx,
			Schema:  graphql.MustParseSchema(app.Schema, test.Resolver, test.Opts...),
			Query: `
				{
					supportedChains{
						id
						chainID
						displayID
						name
						logoUrl
					}
				}
			`,
			ExpectedResult: `{
				"supportedChains": [
					{
						"chainID": "0x673",
						"displayID": "1651",
						"id": "Y2hhaW46MTY1MQ==",
						"logoUrl": "https://chainlist.org/unknown-logo.png",
						"name": "Omni Ephemeral"
					},
					{
						"chainID": "0x674",
						"displayID": "1652",
						"id": "Y2hhaW46MTY1Mg==",
						"logoUrl": "https://chainlist.org/unknown-logo.png",
						"name": "Mock L1 Fast"
					},
					{
						"chainID": "0x675",
						"displayID": "1653",
						"id": "Y2hhaW46MTY1Mw==",
						"logoUrl": "https://chainlist.org/unknown-logo.png",
						"name": "Mock L1 Slow"
					},
					{
						"chainID": "0x676",
						"displayID": "1654",
						"id": "Y2hhaW46MTY1NA==",
						"logoUrl": "https://chainlist.org/unknown-logo.png",
						"name": "Mock L2"
					}
				]
			}
			`,
		},
	})
}
