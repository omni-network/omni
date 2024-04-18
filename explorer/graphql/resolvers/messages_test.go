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

func TestXMsgCount(t *testing.T) {
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

func TestXMsgRange(t *testing.T) {
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
					xmsgrange(from: 0, to: 2){
						SourceMessageSender
						TxHash
						BlockHeight
						BlockHash
					}
				}
			`,
			ExpectedResult: `
				{
					"xmsgrange":[{
						"BlockHash":"0x0000000000000000000000000103176f1b2d62675e370103176f1b2d62675e37",
						"BlockHeight":"0x0",
						"SourceMessageSender":"0x0102030405060708090a0b0c0d0e0f1011121314",
						"TxHash":"0x0102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f20"
					},
					{
						"BlockHash":"0x0000000000000000000000000103176f1b2d62675e370103176f1b2d62675e37",
						"BlockHeight":"0x1",
						"SourceMessageSender":"0x0102030405060708090a0b0c0d0e0f1011121314",
						"TxHash":"0x0102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f20"
					}]
				}
			`,
		},
	})
}

func TestXMsg(t *testing.T) {
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
					xmsg(sourceChainID: 1, destChainID: 2, streamOffset: 0){
						SourceMessageSender
						TxHash
						BlockHash
						Block {
							BlockHeight
						}
						Receipts {
							SourceChainID
						}
					}
				}
			`,
			ExpectedResult: `
			{
				"xmsg":{
					"BlockHash":"0x0000000000000000000000000103176f1b2d62675e370103176f1b2d62675e37",
					"Block":{
						"BlockHeight":"0x0"
					},
					"SourceMessageSender":"0x0102030405060708090a0b0c0d0e0f1011121314",
					"TxHash":"0x0102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f20",
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

// TODO (DAN): Fix tests (why does our auto increment id start super high? Add test for cursor out of range, negative cursor, negative limit

func TestXMsgsNoCursor(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	test := createGqlTest(t)
	t.Cleanup(func() {
		if err := test.Client.Close(); err != nil {
			t.Error(err)
		}
	})
	db.CreateTestBlocks(t, ctx, test.Client, 5)

	gqltesting.RunTests(t, []*gqltesting.Test{
		{
			Context: ctx,
			Schema:  graphql.MustParseSchema(app.Schema, &resolvers.Query{BlocksResolver: test.Resolver}, test.Opts...),
			Query: `
				{
					xmsgs(limit: 2){
						TotalCount
						Edges{
							Node {
								ID
								StreamOffset
								TxHash
								Block {
									BlockHeight
								}
								Receipts {
									Success
								}
							}
						}
						PageInfo {
							StartCursor
						}
					}
				}
			`,
			ExpectedResult: `
			{
				"xmsgs":{
					"Edges":[
						{
							"Node":{
								"Block":{
									"BlockHeight":"0x0"
								},
								"Receipts":[
									{
										"Success":{"Set":true,"Value":true}
									}
								],
								"ID": "8589934593",
								"StreamOffset":"0x0",
								"TxHash":"0x0102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f20"}
							},{
							"Node":{
								"Block":{
									"BlockHeight":"0x1"
								},
								"Receipts":[
									{
										"Success":{"Set":true,"Value":true}
									}
								],
								"ID": "8589934594",
								"StreamOffset":"0x1",
								"TxHash":"0x0102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f20"
							}
						}
					],
					"PageInfo":{
						"StartCursor":"0x200000003"
					},
					"TotalCount":"0x5"
				}
			}
			`,
		},
	})
}

func TestXMsgsNoLimit(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	test := createGqlTest(t)
	t.Cleanup(func() {
		if err := test.Client.Close(); err != nil {
			t.Error(err)
		}
	})
	db.CreateTestBlocks(t, ctx, test.Client, 5)

	gqltesting.RunTests(t, []*gqltesting.Test{
		{
			Context: ctx,
			Schema:  graphql.MustParseSchema(app.Schema, &resolvers.Query{BlocksResolver: test.Resolver}, test.Opts...),
			Query: `
				{
					xmsgs(cursor: "0x200000003"){
						TotalCount
						Edges{
							Node {
								StreamOffset
								TxHash
								ID
								Block {
									BlockHeight
								}
								Receipts {
									Success
								}
							}
						}
						PageInfo {
							StartCursor
						}
					}
				}
			`,
			ExpectedResult: `
			{
				"xmsgs":{
					"Edges":[
						{
							"Node":{
								"Block":{
									"BlockHeight":"0x0"
								},"Receipts":[
									{
										"Success":{"Set":true,"Value":true}
									}
								],
								"ID":"8589934593",
								"StreamOffset":"0x0",
								"TxHash":"0x0102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f20"
							}
						}
					],
					"PageInfo":{
						"StartCursor":"0x200000002"
					},
					"TotalCount":"0x5"
				}
			}
			`,
		},
	})
}

func TestXMsgsCursorOffset(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	test := createGqlTest(t)
	t.Cleanup(func() {
		if err := test.Client.Close(); err != nil {
			t.Error(err)
		}
	})
	db.CreateTestBlocks(t, ctx, test.Client, 5)

	gqltesting.RunTests(t, []*gqltesting.Test{
		{
			Context: ctx,
			Schema:  graphql.MustParseSchema(app.Schema, &resolvers.Query{BlocksResolver: test.Resolver}, test.Opts...),
			Query: `
				{
					xmsgs(){
						TotalCount
						Edges{
							Node {
								StreamOffset
								TxHash
								Block {
									BlockHeight
								}
								Receipts {
									Success
								}
							}
						}
						PageInfo {
							StartCursor
						}
					}
				}
			`,
			ExpectedResult: `
			{
				"xmsgs":{
					"Edges":[
						{
							"Node":{
								"Block":{
									"BlockHeight":"0x0"
								},"Receipts":[
									{
										"Success":{"Set":true,"Value":true}
									}
								],
								"StreamOffset":"0x0",
								"TxHash":"0x0102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f20"
							}
						}
					],
					"PageInfo":{
						"StartCursor":"0x200000002"
					},
					"TotalCount":"0x5"
				}
			}
			`,
		},
	})
}
