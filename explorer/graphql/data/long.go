package data

import (
	"fmt"
	"strconv"
)

// Implements the graphql.Marshaler interface for the Long type using base-10 numbers.
type Long uint64

func (l Long) ImplementsGraphQLType(name string) bool {
	return name == "Long"
}

func (l *Long) UnmarshalGraphQL(input interface{}) error {
	switch input := input.(type) {
	case string:
		value, err := strconv.ParseUint(input, 10, 64)
		if err != nil {
			return err
		}
		*l = Long(value)
		return nil
	default:
		return fmt.Errorf("cannot unmarshal Long scalar type from %T", input)
	}
}

func (l Long) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%d"`, l)), nil
}
