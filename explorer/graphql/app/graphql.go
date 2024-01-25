package app

import (
	_ "embed"
	"net/http"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/omni-network/omni/explorer/graphql/data"
	"github.com/omni-network/omni/explorer/graphql/resolvers"
)

//go:embed graphql.schema
var schema string

//go:embed index.html
var graphiql []byte

func GraphQL(provider data.Provider) http.Handler {
	// dummy hard-coded data
	br := resolvers.BlocksResolver{
		BlocksProvider: provider,
	}

	opts := []graphql.SchemaOpt{
		graphql.UseFieldResolvers(),
		graphql.UseStringDescriptions(),
	}
	s := graphql.MustParseSchema(schema, &resolvers.Query{BlocksResolver: br}, opts...)
	return &relay.Handler{Schema: s}
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html; charset=utf-8")
	w.Write(graphiql)
}
