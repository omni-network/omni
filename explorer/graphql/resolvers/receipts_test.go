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

func TestXReceiptCount(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	test := createGqlTest(t)
	t.Cleanup(func() {
		if err := test.Client.Close(); err != nil {
			t.Error(err)
		}
	})
	db.CreateTestBlocks(t, ctx, test.Client, 3)

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
					"xreceiptcount": "0x2"
				}
			`,
		},
	})
}

func TestReceipt(t *testing.T) {
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
					xreceipt(sourceChainID: 1, destChainID: 2, streamOffset: 0){
						TxHash
						Block {
							BlockHeight
						}
						Messages {
							SourceMessageSender
						}
					}
				}
			`,
			ExpectedResult: `
			{
				"xreceipt":{
					"TxHash":"0x0102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f21",
					"Block":{
						"BlockHeight":"0x0"
					},
					"Messages":[
						{
							"SourceMessageSender":"0x0102030405060708090a0b0c0d0e0f1011121314"
						}
					]
				}
			}
			`,
		},
	})
}
