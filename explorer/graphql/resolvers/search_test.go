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

func TestSearchQueryBlock(t *testing.T) {
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
					search(query: "0x0000000000000000000000000103176f1b2d62675e370103176f1b2d62675e37"){
						TxHash
						BlockHeight
						SourceChainID
						Type
					}
				}
			`,
			ExpectedResult: `
				{
					"search":{
						"BlockHeight":"0x0",
						"SourceChainID":"0x1",
						"TxHash":"0x0000000000000000000000000000000000000000000000000000000000000000",
						"Type":"BLOCK"
					}
				}
			`,
		},
	})
}

func TestSearchQueryMessage(t *testing.T) {
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
					search(query: "0x0102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f20"){
						TxHash
						BlockHeight
						SourceChainID
						Type
					}
				}
			`,
			ExpectedResult: `
				{
					"search":{
						"BlockHeight":"0x0",
						"SourceChainID":"0x0",
						"TxHash":"0x0102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f20",
						"Type":"MESSAGE"
					}
				}
			`,
		},
	})
}

func TestSearchQueryReceipt(t *testing.T) {
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
					search(query: "0x0102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f21"){
						TxHash
						BlockHeight
						SourceChainID
						Type
					}
				}
			`,
			ExpectedResult: `
				{
					"search":{
						"BlockHeight":"0x0",
						"SourceChainID":"0x0",
						"TxHash":"0x0102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f21",
						"Type":"RECEIPT"
					}
				}
			`,
		},
	})
}
