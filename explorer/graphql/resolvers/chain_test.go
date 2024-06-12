package resolvers_test

import (
	"context"
	"testing"

	"github.com/omni-network/omni/explorer/db/testutil"
	"github.com/omni-network/omni/explorer/graphql/app"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/gqltesting"
)

func TestSupportedChains(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	devnet := createGqlTest(t, netconf.Devnet)
	testutil.CreateTestChain(t, ctx, devnet.Client, 1)
	testutil.CreateTestChain(t, ctx, devnet.Client, 2)
	staging := createGqlTest(t, netconf.Staging)
	testnet := createGqlTest(t, netconf.Omega)

	gqltesting.RunTests(t, []*gqltesting.Test{
		{
			Context: ctx,
			Schema:  graphql.MustParseSchema(app.Schema, devnet.Resolver, devnet.Opts...),
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
						"logoUrl": "https://icons.llamao.fi/icons/chains/rsz_optimism.jpg",
						"name": "Mock L1 Fast"
					},
					{
						"chainID": "0x676",
						"displayID": "1654",
						"id": "Y2hhaW46MTY1NA==",
						"logoUrl": "https://icons.llamao.fi/icons/chains/rsz_arbitrum.jpg",
						"name": "Mock L2"
					},
					{
						"chainID": "0x677",
						"displayID": "1655",
						"id": "Y2hhaW46MTY1NQ==",
						"logoUrl": "https://icons.llamao.fi/icons/chains/rsz_optimism.jpg",
						"name": "Mock Op"
					},
					{
						"chainID": "0x678",
						"displayID": "1656",
						"id": "Y2hhaW46MTY1Ng==",
						"logoUrl": "https://icons.llamao.fi/icons/chains/rsz_arbitrum.jpg",
						"name": "Mock Arb"
					}
				]
			}
			`,
		},
		{
			Context: ctx,
			Schema:  graphql.MustParseSchema(app.Schema, staging.Resolver, staging.Opts...),
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
						"chainID": "0x675",
						"displayID": "1653",
						"id": "Y2hhaW46MTY1Mw==",
						"logoUrl": "https://chainlist.org/unknown-logo.png",
						"name": "Mock L1 Slow"
					},
					{
						"chainID": "0xaa37dc",
						"displayID": "11155420",
						"id": "Y2hhaW46MTExNTU0MjA=",
						"logoUrl": "https://icons.llamao.fi/icons/chains/rsz_optimism.jpg",
						"name": "Op Sepolia"
					}
				]
			}
			`,
		},
		{
			Context: ctx,
			Schema:  graphql.MustParseSchema(app.Schema, testnet.Resolver, testnet.Opts...),
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
						"chainID": "0xa4",
						"displayID": "164",
						"id": "Y2hhaW46MTY0",
						"logoUrl": "https://chainlist.org/unknown-logo.png",
						"name": "Omni Omega"
					},
					{
						"chainID": "0x4268",
						"displayID": "17000",
						"id": "Y2hhaW46MTcwMDA=",
						"logoUrl": "https://icons.llamao.fi/icons/chains/rsz_ethereum.jpg",
						"name": "Holesky"
					},
					{
						"chainID": "0x66eee",
						"displayID": "421614",
						"id": "Y2hhaW46NDIxNjE0",
						"logoUrl": "https://icons.llamao.fi/icons/chains/rsz_arbitrum.jpg",
						"name": "Arb Sepolia"
					},
					{
						"chainID": "0xaa37dc",
						"displayID": "11155420",
						"id": "Y2hhaW46MTExNTU0MjA=",
						"logoUrl": "https://icons.llamao.fi/icons/chains/rsz_optimism.jpg",
						"name": "Op Sepolia"
					}
				]
			}
			`,
		},
	})
}
