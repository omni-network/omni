package resolvers

import (
	"context"
	"encoding/json"
	"math/big"
	"strconv"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
)

// BigInt represents a bigint number in GraphQL since many browsers only work with int32.
type BigInt struct {
	big.Int
}

// ImplementsGraphQLType maps this custom Go type
// to the graphql scalar type in the schema.
func (BigInt) ImplementsGraphQLType(name string) bool {
	return name == "BigInt"
}

// UnmarshalGraphQL is a custom unmarshaler for BigInt.
//
// This function will be called whenever you use the
// time scalar as an input.
func (b *BigInt) UnmarshalGraphQL(input any) error {
	switch input := input.(type) {
	case string:

		_, ok := b.Int.SetString(input, 10)
		if !ok {
			return errors.New("failed to parse big int from string")
		}

		return nil
	case int32:

		_, ok := b.Int.SetString(strconv.Itoa(int(input)), 10)
		if !ok {
			return errors.New("failed to parse big int from int32 %v", input)
		}

		return nil
	default:
		return errors.New("wrong type for BigInt: %T", input)
	}
}

// MarshalJSON is a custom marshaler for BigInt.
//
// This function will be called whenever you
// query for fields that use the UInt64 type.
func (b *BigInt) MarshalJSON() ([]byte, error) {
	res, err := json.Marshal(b)
	if err != nil {
		log.Error(context.Background(), "failed to marshal big int to json", err)
		return nil, errors.Wrap(err, "failed to marshal big into to json %v", &b)
	}

	return res, nil
}
