package service

import (
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/uni"
	"github.com/omni-network/omni/scripts/trade/users"
	"github.com/omni-network/omni/scripts/trade/users/db"
)

func userFromDB(u db.User) (users.User, error) {
	addr, err := uni.ParseAddress(u.Address)
	if err != nil {
		return users.User{}, errors.Wrap(err, "parse address")
	}

	return users.User{
		ID:      u.ID,
		PrivyID: u.PrivyID,
		Address: addr,
	}, nil
}

func usersFromDB(ul []db.User) ([]users.User, error) {
	resp := make([]users.User, 0, len(ul))
	for _, u := range ul {
		user, err := userFromDB(u)
		if err != nil {
			return nil, errors.Wrap(err, "convert user from db")
		}

		resp = append(resp, user)
	}

	return resp, nil
}
