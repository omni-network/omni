package data

import (
	"fmt"
	"strconv"

	"github.com/omni-network/omni/lib/errors"
)

// Implements the graphql.Marshaler interface for the Long type using base-10 numbers.
type Long uint64

func (Long) ImplementsGraphQLType(name string) bool {
	return name == "Long"
}

func (l *Long) UnmarshalGraphQL(input any) error {
	switch input := input.(type) {
	case string:
		value, err := strconv.ParseUint(input, 10, 64)
		if err != nil {
			return errors.Wrap(err, "parsing Long")
		}
		*l = Long(value)
		return nil
	default:
		return errors.New("cannot unmarshal Long scalar type from", "type", fmt.Sprintf("%T", input))
	}
}

func (l Long) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%d"`, l)), nil
}
