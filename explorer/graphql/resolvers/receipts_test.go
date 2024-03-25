package resolvers_test

import (
	"context"
	"testing"

	"github.com/omni-network/omni/explorer/db/ent"
	"github.com/omni-network/omni/explorer/graphql/app"
	"github.com/omni-network/omni/explorer/graphql/resolvers"

	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/gqltesting"
)

func TestXReceiptCount(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	test := createGqlTest(t)
	defer func(Client *ent.Client) {
		err := Client.Close()
		if err != nil {
			t.Error(err)
		}
	}(test.Client)
	resolvers.CreateTestBlocks(ctx, t, test.Client, 2)

	gqltesting.RunTests(t, []*gqltesting.Test{
		{
			Context: ctx,
			Schema:  graphql.MustParseSchema(app.Schema, &resolvers.Query{BlocksResolver: test.Resolver}, test.Opts...),
			Query: `
				{
					xreceiptcount
				}
			`,
			ExpectedResult: `
				{
					"xreceiptcount": "0x1"
				}
			`,
		},
	})
}
