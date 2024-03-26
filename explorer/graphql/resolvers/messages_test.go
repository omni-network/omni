package resolvers_test

import (
	"context"
	"testing"

	"github.com/omni-network/omni/explorer/graphql/app"
	"github.com/omni-network/omni/explorer/graphql/resolvers"

	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/gqltesting"
)

func TestXMsgCount(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	test := createGqlTest(t)
	t.Cleanup(func() {
		if err := test.Client.Close(); err != nil {
			t.Error(err)
		}
	})
	createTestBlocks(ctx, t, test.Client, 2)

	gqltesting.RunTests(t, []*gqltesting.Test{
		{
			Context: ctx,
			Schema:  graphql.MustParseSchema(app.Schema, &resolvers.Query{BlocksResolver: test.Resolver}, test.Opts...),
			Query: `
				{
					xmsgcount
				}
			`,
			ExpectedResult: `
				{
					"xmsgcount": "0x2"
				}
			`,
		},
	})
}
