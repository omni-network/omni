package resolvers

import "context"

type Query struct {
	BlocksResolver
}

func (Query) Hello(_ context.Context, args struct{ Name string }) string {
	return "Hello, " + args.Name + "!"
}
