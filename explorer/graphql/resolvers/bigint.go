package resolvers

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
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
// time scalar as an input
func (b *BigInt) UnmarshalGraphQL(input interface{}) error {
	switch input := input.(type) {
	case string:
		_, ok := b.Int.SetString(input, 10)
		if !ok {
			return errors.New("failed to parse big int")
		}
		return nil
	default:
		return fmt.Errorf("wrong type for BigInt: %T", input)
	}
}

// MarshalJSON is a custom marshaler for BigInt.
//
// This function will be called whenever you
// query for fields that use the UInt64 type.
func (b *BigInt) MarshalJSON() ([]byte, error) {
	return json.Marshal(b)
}
