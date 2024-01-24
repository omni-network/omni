package app

import (
	_ "embed"
	"net/http"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/omni-network/omni/explorer/graphql/resolvers"
)

//go:embed graphql.schema
var schema string

//go:embed index.html
var graphiql []byte

func GraphQL() http.Handler {
	// dummy hard-coded data
	br := resolvers.BlocksResolver{
		Blocks: []resolvers.XBlock{
			{
				ID:   graphql.ID("id1"),
				Name: "name1",
			},
			{
				ID:   graphql.ID("id2"),
				Name: "name2",
			},
			{
				ID:   graphql.ID("id3"),
				Name: "name3",
			},
			{
				ID:   graphql.ID("id4"),
				Name: "name4",
			},
			{
				ID:   graphql.ID("id5"),
				Name: "name5",
			},
			{
				ID:   graphql.ID("id6"),
				Name: "name6",
			},
			{
				ID:   graphql.ID("id7"),
				Name: "name7",
			},
		},
	}

	s := graphql.MustParseSchema(schema, &resolvers.Query{BlocksResolver: br}, graphql.UseFieldResolvers())
	return &relay.Handler{Schema: s}
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html; charset=utf-8")
	w.Write(graphiql)
}
