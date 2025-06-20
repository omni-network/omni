package users

import (
	"github.com/omni-network/omni/lib/errors"

	"github.com/google/uuid"
)

func (r RequestCreate) Validate() error {
	if r.PrivyID == "" {
		return errors.New("privy_id empty")
	}
	if r.Address.IsZero() {
		return errors.New("address empty")
	}
	if r.ID == uuid.Nil {
		return errors.New("id empty")
	}

	return nil
}
