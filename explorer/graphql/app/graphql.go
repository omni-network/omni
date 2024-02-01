package app

import (
	"net/http"

	"github.com/omni-network/omni/explorer/graphql/data"
	"github.com/omni-network/omni/explorer/graphql/resolvers"
	"github.com/omni-network/omni/lib/log"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"

	_ "embed"
)

//go:embed schema.graphql
var Schema string

//go:embed index.html
var graphiql []byte

// GraphQL returns a new graphql handler. We use the relay handler to create the graphql handler.
func GraphQL(provider data.Provider) http.Handler {
	br := resolvers.BlocksResolver{
		BlocksProvider: provider,
	}

	opts := []graphql.SchemaOpt{
		graphql.UseFieldResolvers(),
		graphql.UseStringDescriptions(),
	}
	s := graphql.MustParseSchema(Schema, &resolvers.Query{BlocksResolver: br}, opts...)

	return &relay.Handler{Schema: s}
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html; charset=utf-8")
	_, err := w.Write(graphiql)
	if err != nil {
		log.Warn(r.Context(), "graphql home err", err)
	}
}
