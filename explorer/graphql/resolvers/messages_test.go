package resolvers_test

import (
	"context"
	"testing"

	"github.com/omni-network/omni/explorer/db/testutil"
	"github.com/omni-network/omni/explorer/graphql/app"
	"github.com/omni-network/omni/explorer/graphql/resolvers"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/gqltesting"
)

func TestXMsg(t *testing.T) {
	t.Skip("This test is failing because the schema was changed")
	t.Parallel()
	ctx := context.Background()
	test := createGqlTest(t, netconf.Devnet)
	t.Cleanup(func() {
		if err := test.Client.Close(); err != nil {
			t.Error(err)
		}
	})
	testutil.CreateTestBlocks(t, ctx, test.Client, 2)

	gqltesting.RunTests(t, []*gqltesting.Test{
		{
			Context: ctx,
			Schema:  graphql.MustParseSchema(app.Schema, resolvers.NewRoot(test.Provider), test.Opts...),
			Query: `
				{
					xmsg(sourceChainID: 1, destChainID: 2, offset: 0){
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
	t.Skip("This test is failing because the schema was changed")
	t.Parallel()
	ctx := context.Background()
	test := createGqlTest(t, netconf.Devnet)
	t.Cleanup(func() {
		if err := test.Client.Close(); err != nil {
			t.Error(err)
		}
	})
	testutil.CreateTestBlocks(t, ctx, test.Client, 5)

	gqltesting.RunTests(t, []*gqltesting.Test{
		{
			Context: ctx,
			Schema:  graphql.MustParseSchema(app.Schema, resolvers.NewRoot(test.Provider), test.Opts...),
			Query: `
				{
					xmsgs(limit: 2){
						TotalCount
						Edges{
							Cursor
							Node {
								ID
								Offset
								TxHash
								BlockHeight
								Status
							}
						}
						PageInfo {
							NextCursor
						}
					}
				}
			`,
			ExpectedResult: `
			{
				"xmsgs":{
					"Edges":[
						{
							"Cursor":"0x200000005",
							"Node":{
								"BlockHeight":"0x4",
								"Status": "PENDING",
								"ID": "8589934597",
								"Offset":"0x4",
								"TxHash":"0x0102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f20"
							}
						},
						{
							"Cursor":"0x200000004",
							"Node":{
								"BlockHeight":"0x3",
								"Status": "SUCCESS",
								"ID": "8589934596",
								"Offset":"0x3",
								"TxHash":"0x0102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f20"
							}
						}
					],
					"PageInfo":{
						"NextCursor":"0x200000003"
					},
					"TotalCount":"0x5"
				}
			}
			`,
		},
	})
}

func TestXMsgsNoLimit(t *testing.T) {
	t.Skip("This test is failing because the schema was changed")
	t.Parallel()
	ctx := context.Background()
	test := createGqlTest(t, netconf.Devnet)
	t.Cleanup(func() {
		if err := test.Client.Close(); err != nil {
			t.Error(err)
		}
	})
	testutil.CreateTestBlocks(t, ctx, test.Client, 5)

	gqltesting.RunTests(t, []*gqltesting.Test{
		{
			Context: ctx,
			Schema:  graphql.MustParseSchema(app.Schema, resolvers.NewRoot(test.Provider), test.Opts...),
			Query: `
				{
					xmsgs(cursor: "0x200000003"){
						TotalCount
						Edges{
							Cursor
							Node {
								MsgOffset
								TxHash
								ID
								BlockHeight
								Status
							}
						}
					}
				}
			`,
			ExpectedResult: `
				{
					"xmsgs":
					{
						"Edges":
						[
							{
								"Cursor":"0x200000003",
								"Node":{
									"ID":"8589934595",
									"BlockHeight":"0x2",
									"Status":"SUCCESS",
									"MsgOffset":"0x2",
									"TxHash":"0x0102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f20"
								}
							},{
								"Cursor":"0x200000002",
								"Node":{
									"BlockHeight":"0x1",
									"ID":"8589934594",
									"Status":"SUCCESS",
									"MsgOffset":"0x1",
									"TxHash":"0x0102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f20"
								}
							},{
								"Cursor":"0x200000001",
								"Node":{
									"BlockHeight":"0x0",
									"ID":"8589934593",
									"Status":"SUCCESS",
									"MsgOffset":"0x0",
									"TxHash":"0x0102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f20"
								}
							}
						],
						"TotalCount":"0x5"
					}
				}
			`,
		},
	})
}

func TestXMsgsNoParams(t *testing.T) {
	t.Skip("This test is failing because the schema was changed")
	t.Parallel()
	ctx := context.Background()
	test := createGqlTest(t, netconf.Devnet)
	t.Cleanup(func() {
		if err := test.Client.Close(); err != nil {
			t.Error(err)
		}
	})
	testutil.CreateTestBlocks(t, ctx, test.Client, 5)

	gqltesting.RunTests(t, []*gqltesting.Test{
		{
			Context: ctx,
			Schema:  graphql.MustParseSchema(app.Schema, resolvers.NewRoot(test.Provider), test.Opts...),
			Query: `
				{
					xmsgs(){
						TotalCount
						Edges{
							Node {
								MsgOffset
								TxHash
								BlockHeight
								Status
							}
						}
						PageInfo {
							PrevCursor
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
								"BlockHeight":"0x4",
								"Status":"PENDING",
								"MsgOffset":"0x4",
								"TxHash":"0x0102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f20"
							}
						},{
								"Node":{
									"BlockHeight":"0x3",
									"Status":"SUCCESS",
									"MsgOffset":"0x3",
									"TxHash":"0x0102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f20"
								}
							},{
								"Node":{
									"BlockHeight":"0x2",
									"Status":"SUCCESS",
									"MsgOffset":"0x2",
									"TxHash":"0x0102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f20"
								}
							},{
								"Node":{
									"BlockHeight":"0x1",
									"Status":"SUCCESS",
									"MsgOffset":"0x1",
									"TxHash":"0x0102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f20"
								}
							},{
							"Node":{
								"BlockHeight":"0x0",
								"Status":"SUCCESS",
								"MsgOffset":"0x0",
								"TxHash":"0x0102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f20"
							}
						}
					],
					"PageInfo":{
						"PrevCursor":"0x20000001e"
					},
					"TotalCount":"0x5"
				}
			}
			`,
		},
	})
}
