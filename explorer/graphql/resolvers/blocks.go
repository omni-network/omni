package resolvers

import (
	"context"
	"errors"
	"fmt"

	"github.com/graph-gophers/graphql-go"
)

type XBlock struct {
	ID   graphql.ID
	Name string
}

type BlocksResolver struct {
	Blocks []XBlock
}

type XBlocksArgs struct {
	From int32
	To   int32
}

func (b *BlocksResolver) XBlocks(ctx context.Context, args XBlocksArgs) ([]XBlock, error) {
	var res []XBlock
	if args.From < 0 {
		return res, errors.New("negative index from index")
	}

	if int(args.To) > len(b.Blocks) {
		return res, fmt.Errorf("max length greater than lenght: lenght %d", len(b.Blocks))
	}

	res = b.Blocks[args.From:args.To]
	return res, nil
}
